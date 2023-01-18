package beidermorse

import (
	"math/bits"
	"strings"
)

func redoLanguage(input string, mode Mode, ruleset Ruleset, concat bool) tokens {
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

func makeTokens(input string, mode Mode, ruleset Ruleset, lang languageID, concat bool) tokens {
	input = prepareInput(input, mode) // returns at max two words

	mainRules, final1, final2, discards := getRules(mode, ruleset, lang)

	// process a multiword name of form X Y
	if space := strings.Index(input, " "); space != -1 { // number of words is exactly two
		if concat { // exact matches
			// X Y => XY
			input = strings.ReplaceAll(input, " ", "") // concatenate the separate words of a name
		} else { // number of words is exactly two
			word1 := substrTo(input, space)
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

	result := tokens{{text: input, langs: lang}}

	// apply main set of rules
	result = mainRules.apply(result, lang, false)
	// apply common final rules
	result = final1.apply(result, lang, false)
	// apply lang specific final rules
	result = final2.apply(result, lang, true)

	return result
}

func prepareInput(input string, mode Mode) string {
	input = strings.TrimSpace(strings.ToLower(input))

	// remove spaces from within certain leading words
	for _, item := range firstList {
		itemLen := len([]rune(item))
		target := item + " "
		if substrTo(input, itemLen+1) == target {
			target = item
			input = strings.ReplaceAll(target, " ", "") + substrFrom(input, itemLen)
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
			input = strings.ReplaceAll(input, target, "")                         // remove all occurences
			input = substrTo(input, firstOne) + " " + substrFrom(input, firstOne) // replace first occurence with space
		}
	}

	return input
}

func inArray[T comparable](data []T, value T) bool {
	for _, item := range data {
		if item == value {
			return true
		}
	}
	return false
}

func substrFrom(s string, start int) string {
	return string([]rune(s)[start:])
}

func substrTo(s string, end int) string {
	return string([]rune(s)[:end])
}

func getRules(
	mode Mode,
	ruleset Ruleset,
	lang languageID,
) (mainRules rules, finalRules1 rules, finalRules2 rules, discards []string) {
	langCount := bits.OnesCount64(uint64(lang))
	if langCount > 1 {
		lang = 1 // any
	}

	switch mode {
	case Generic:
		mainRules = genRules[genLang(lang)]
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
		mainRules = ashRules[ashLang(lang)]
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
		mainRules = sepRules[sepLang(lang)]
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
