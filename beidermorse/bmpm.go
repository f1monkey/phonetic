package beidermorse

import (
	"math/bits"
	"strconv"
	"strings"
	"unicode/utf8"
)

const langAny = 1

func redoLanguage(input string, mode Mode, ruleset Ruleset, concat bool) []phonetic {
	// we can do a better job of determining the language now that multiple names have been split
	languageArg := detectLang(input, mode)
	return makeTokens(input, mode, ruleset, languageArg, concat)
}

func detectLang(word string, mode Mode) languageID {
	var rules []langRule
	var all languageID
	switch mode {
	case Ashkenazi:
		rules = ashLangRules
		all = languageID(ashAll)
	case Sephardic:
		rules = sepLangRules
		all = languageID(sepAll)
	case Generic:
		rules = genLangRules
		all = languageID(genAll)
	}

	remaining := all
	for _, rule := range rules {
		if !rule.match.matches(word) {
			continue
		}

		if rule.accept {
			remaining &= rule.langs
		} else {
			remaining &= ^rule.langs % (remaining + 1)
		}

		if remaining == 0 {
			return remaining
		}
	}

	return remaining
}

var firstList = []string{"de la", "van der", "van den"}

func makeTokens(input string, mode Mode, ruleset Ruleset, lang languageID, concat bool) []phonetic {
	input = prepareInput(input, mode) // returns at max two words

	rules, final1, final2, discards := getRules(mode, ruleset, lang)

	// process a multiword name of form X Y
	if space := strings.Index(input, " "); space != -1 { // number of words is exactly two
		if concat { // exact matches
			// X Y => XY
			input = strings.ReplaceAll(input, " ", "") // concatenate the separate words of a name
		} else { // number of words is exactly two
			word1 := substr(input, 0, space)
			word2 := substrFrom(input, space+1)
			if inArray(discards, word1) {
				// X Y => Y and XY
				results := redoLanguage(word2, mode, ruleset, concat)
				results = append(results, redoLanguage(word1+word2, mode, ruleset, concat)...)
				return results
			} else { // first word is not in list
				// X Y => X, Y, and XY
				results := redoLanguage(word1, mode, ruleset, concat)
				results = append(results, redoLanguage(word2, mode, ruleset, concat)...)
				results = append(results, redoLanguage(word1+word2, mode, ruleset, concat)...)
				return results
			}
		}
	}

	// apply main set of rules
	result := applyRules([]phonetic{{text: input, langs: lang}}, rules, lang, false)
	// apply common final rules
	result = applyRules(result, final1, lang, false)
	// apply lang specific final rules
	result = applyRules(result, final2, lang, true)

	return result
}

func prepareInput(input string, mode Mode) string {
	input = strings.TrimSpace(strings.ToLower(input))

	// remove spaces from within certain leading words
	for _, item := range firstList {
		target := item + " "
		if substr(input, 0, utf8.RuneCountInString(target)) == target {
			target = item
			input = strings.ReplaceAll(target, " ", "") + substrFrom(input, utf8.RuneCountInString(target))
		}
	}

	// for ash and gen -- remove all apostrophes
	if mode != Sephardic {
		input = strings.ReplaceAll(input, "'", "")
	}

	// remove all apostrophoes, dashes, and spaces except for the first one, replace first one with space
	list := []string{"'", "-", " "}
	for i := 0; i < len(list); i++ {
		target := list[i]

		if firstOne := strings.Index(input, target); firstOne != -1 {
			input = strings.ReplaceAll(input, target, "")                          // remove all occurences
			input = substr(input, 0, firstOne) + " " + substrFrom(input, firstOne) // replace first occurence with space
		}
	}

	return input
}

func applyRules(input []phonetic, rules []rule, languageArg languageID, ignoreLangs bool) []phonetic {
	if len(rules) == 0 {
		return input
	}

	var result []phonetic
	for _, token := range input {
		tokenRunes := []rune(token.text)
		newTokens := []phonetic{{text: "", langs: token.langs}}
		for j := 0; j < len(tokenRunes); {
			patternLength := 0
			found := false

			for _, r := range rules {
				patternLength = len([]rune(r.pattern))

				if patternLength > utf8.RuneCountInString(token.text)-j || substr(token.text, j, patternLength) != r.pattern { // no match
					continue
				}

				if r.rightContext != nil {
					if !r.rightContext.matches(substrFrom(token.text, j+patternLength)) {
						continue
					}
				}

				if r.leftContext != nil {
					if !r.leftContext.matches(substr(token.text, 0, j)) {
						continue
					}
				}

				if len(newTokens) == 0 {
					newTokens = mergeTokenGroups(languageArg, r.phoneticRules)
					if len(newTokens) != 0 {
						found = true
						break
					}
				} else {
					if rr := mergeTokenGroups(languageArg, newTokens, r.phoneticRules); len(rr) != 0 {
						newTokens = rr
						found = true
						break
					}
				}
			}

			if !found {
				for k := range newTokens {
					newTokens[k].text += string(tokenRunes[j])
				}
				j++
			} else {
				j += patternLength
			}
		}

		result = append(result, newTokens...)
	}

	if ignoreLangs {
		for i := range result {
			result[i].langs = langsAny
		}
	}

	return dedupTokens(result)
}

func dedupTokens(tokens []phonetic) []phonetic {
	result := make([]phonetic, 0, len(tokens))
	uniq := make(map[string]struct{}, len(tokens))

	for _, t := range tokens {
		key := strconv.Itoa(int(t.langs)) + "_" + t.text
		if _, ok := uniq[key]; ok {
			continue
		}
		uniq[key] = struct{}{}
		result = append(result, t)
	}

	return result
}

func mergeTokenGroups(langRestricton languageID, src ...[]phonetic) []phonetic {
	if len(src) == 0 {
		return nil
	}

	if len(src) == 1 {
		return src[0]
	}

	l := 0
	for _, r := range src {
		l += len(r)
	}

	result := src[0]
	i := 1
	for i < len(src) {
		newResult := make([]phonetic, 0, len(result)*len(src[i]))
		for _, r1 := range result {
			for _, r2 := range src[i] {
				lang := mergeLangResults(langRestricton, r1.langs, r2.langs)
				if lang == langsInvalid {
					continue
				}

				newResult = append(newResult, phonetic{
					text:  r1.text + r2.text,
					langs: lang,
				})
			}
		}
		result = newResult
		i++
	}

	return result
}

func mergeLangResults(src ...languageID) languageID {
	if len(src) == 0 {
		return langsInvalid
	}

	result := langsUnitialized
	for _, l := range src {
		if result == langsUnitialized || result == langsAny {
			result = l
			continue
		}
		result &= l
	}

	return result
}

func inArray[T comparable](data []T, value T) bool {
	for _, item := range data {
		if item == value {
			return true
		}
	}
	return false
}

func substr(s string, start int, length int) string {
	asRunes := []rune(s)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func substrFrom(s string, start int) string {
	return string([]rune(s)[start:])
}

func indexAt(s, sep string, n int) int {
	return n + strings.Index(s[n:], sep)
}

func getRules(
	mode Mode,
	ruleset Ruleset,
	lang languageID,
) (rules []rule, finalRules1 []rule, finalRules2 []rule, discards []string) {
	langCount := bits.OnesCount64(uint64(lang))
	if langCount > 1 {
		lang = 1 // any
	}

	switch mode {
	case Generic:
		rules = genRules[genLang(lang)]
		discards = genDiscards
		switch ruleset {
		case Approx:
			finalRules1 = genFinalRules.approx.first
			finalRules2 = genFinalRules.approx.second[lang]
		case Exact:
			finalRules1 = genFinalRules.exact.first
			finalRules2 = genFinalRules.exact.second[lang]
		}
	case Ashkenazi:
		rules = ashRules[ashLang(lang)]
		discards = ashDiscards
		switch ruleset {
		case Approx:
			finalRules1 = ashFinalRules.approx.first
			finalRules2 = ashFinalRules.approx.second[lang]
		case Exact:
			finalRules1 = ashFinalRules.exact.first
			finalRules2 = ashFinalRules.exact.second[lang]
		}
	case Sephardic:
		rules = sepRules[sepLang(lang)]
		discards = sepDiscards
		switch ruleset {
		case Approx:
			finalRules1 = sepFinalRules.approx.first
			finalRules2 = sepFinalRules.approx.second[lang]
		case Exact:
			finalRules1 = sepFinalRules.exact.first
			finalRules2 = sepFinalRules.exact.second[lang]
		}
	}
	return
}
