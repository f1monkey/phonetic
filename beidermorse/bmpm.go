package beidermorse

import (
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const langAny = 1

func redoLanguage(input string, mode Mode, ruleset Ruleset, concat bool) string {
	// we can do a better job of determining the language now that multiple names have been split
	languageArg := detectLang(input, mode)
	return phonetic(input, mode, ruleset, languageArg, concat)
}

func detectLang(word string, mode Mode) uint64 {
	var rules []langRule
	var all uint64
	switch mode {
	case Ashkenazi:
		rules = ashLangRules
		all = uint64(ashAll)
	case Sephardic:
		rules = sepLangRules
		all = uint64(sepAll)
	case Generic:
		rules = genLangRules
		all = uint64(genAll)
	}

	remaining := all
	for _, rule := range rules {
		if !regCache.get(rule.pattern).MatchString(word) {
			continue
		}

		if rule.accept {
			remaining &= rule.langs
		} else {
			remaining &= (all ^ rule.langs) % (remaining + 1)
		}

		if remaining == 0 {
			return remaining
		}
	}

	return remaining
}

func phonetic(input string, mode Mode, ruleset Ruleset, lang uint64, concat bool) string {
	// algorithm used here is as follows:
	//
	//   Before doing anything else:
	//    (1) replace leading:
	//         de<space>la<space> with dela<space>
	//         van<space>der<space> with vander<space>
	//         van<space>den<space> with vanden<space>
	//    (2) gen and ash: remove all apostrophes (i.e., X'Y ==> XY)
	//    (3) remove all spaces, apostrophes, and dashes except for the first one (i.e. X Y Z ==> X YZ)
	//    (4) convert remaining dashes and apostrophes (if any) to space (i.e. X'Y ==> X Y)

	//   if Exact:
	//     if <space> is present (i.e. X Y/
	//       X Y => XY
	//   if Approx or Hebrew:
	//     if <space> is present (i.e. X Y)
	//       if X in list (different lists for ash, sep, gen, see below)
	//         X Y => Y and XY
	//       else if X is not in list
	//         X Y => X, Y, and XY

	input = strings.TrimSpace(strings.ToLower(input))

	// remove spaces from within certain leading words

	list := []string{"de la", "van der", "van den"}
	for i := 0; i < len(list); i++ {
		target := list[i] + " "
		if substr(input, 0, utf8.RuneCountInString(target)) == target {
			target = list[i]
			input = strings.ReplaceAll(target, " ", "") + substrFrom(input, utf8.RuneCountInString(target))
		}
	}

	// for ash and gen -- remove all apostrophes

	if mode != Sephardic {
		input = strings.ReplaceAll(input, "'", "")
	}

	// remove all apostrophoes, dashes, and spaces except for the first one, replace first one with space

	list = []string{"'", "-", " "}
	for i := 0; i < len(list); i++ {
		target := list[i]

		if firstOne := strings.Index(input, target); firstOne != -1 {
			input = strings.ReplaceAll(input, target, "")                          // remove all occurences
			input = substr(input, 0, firstOne) + " " + substrFrom(input, firstOne) // replace first occurence with space
		}
	}

	if mode == Sephardic {

		list = []string{ // sephardi
			"abe", "aben", "abi", "abou", "abu", "al", "bar", "ben", "bou", "bu",
			"d", "da", "dal", "de", "del", "dela", "della", "des", "di",
			"el", "la", "le", "ibn", "ha",
		}

	} else if mode == Ashkenazi { // ashkenazi

		list = []string{
			"ben", "bar", "ha",
		}

	} else { // generic

		list = []string{
			"abe", "aben", "abi", "abou", "abu", "al", "bar", "ben", "bou", "bu",
			"d", "da", "dal", "de", "del", "dela", "della", "des", "di", "dos", "du",
			"el", "la", "le", "ibn", "van", "von", "ha", "vanden", "vander",
		}
	}

	// process a multiword name of form X Y

	if space := strings.Index(input, " "); space != -1 { // number of words is exactly two
		if concat { // exact matches
			// X Y => XY
			input = strings.ReplaceAll(input, " ", "") // concatenate the separate words of a name
		} else { // number of words is exactly two
			word1 := substr(input, 0, space)
			word2 := substrFrom(input, space+1)
			if inArray(list, word1) {
				// X Y => Y and XY
				results := redoLanguage(word2, mode, ruleset, concat)
				results += "-" + redoLanguage(word1+word2, mode, ruleset, concat)
				return results
			} else { // first word is not in list
				// X Y => X, Y, and XY
				results := redoLanguage(word1, mode, ruleset, concat)
				results += "-" + redoLanguage(word2, mode, ruleset, concat)
				results += "-" + redoLanguage(word1+word2, mode, ruleset, concat)
				return results
			}
		}
	}

	// at this point, input is only a single word

	inputLength := len([]rune(input))

	// apply language rules to map to phonetic alphabet
	rules, final1, final2 := getRules(mode, ruleset, lang)

	phonetic := ""
	var patternLength int
	for i := 0; i < inputLength; {
		found := false
		for _, rule := range rules {
			pattern := rule.pattern()
			patternLength = utf8.RuneCountInString(pattern)
			lcontext := rule.contextLeft()
			rcontext := rule.contextRight()

			// check to see if next sequence in input matches the string in the rule
			if patternLength > inputLength-i || substr(input, i, patternLength) != pattern { // no match
				continue
			}

			right := "^" + rcontext + ""
			left := "" + lcontext + "$" + ""

			// check that right context is satisfied
			if rcontext != "" {
				reg := regexp.MustCompile(right)
				if !reg.MatchString(substrFrom(input, i+patternLength)) {
					continue
				}
			}

			// check that left context is satisfied
			if lcontext != "" {
				reg := regexp.MustCompile(left)
				if !reg.MatchString(substr(input, 0, i)) {
					continue
				}
			}

			// // check to see if languageArg is one of the allowable ones (used only with "any" rules)
			// if (languageArg != "1") && (languagePos < len(rule)) {
			// 	language = rule[languagePos] // the required language(s) for this rule to apply
			// 	logical := rule[logicalPos]  // do we require ALL or ANY of the required languages
			// 	if logical == "ALL" {
			// 		// check to see if languageArg contains all the required languages
			// 		if (languageArg & language) != language {
			// 			continue
			// 		}
			// 	} else { // any
			// 		// check to see if languageArg contains at least one required language
			// 		if (languageArg & language) == 0 {
			// 			continue
			// 		}
			// 	}
			// }

			// check for incompatible attributes

			candidate := applyRuleIfCompatible(phonetic, rule.phonetic(), lang)
			if candidate == "" {
				continue
			}
			phonetic = candidate
			found = true
			break
		}
		if !found { // character in name that is not in table -- e.g., space
			patternLength = 1
		}
		i += patternLength
	}

	// apply final rules on phonetic-alphabet, doing a substitution of certain characters
	// finalRules1 are the common approx rules, finalRules2 are approx rules for specific language

	phonetic = applyFinalRules(phonetic, final1, lang, false) // apply common rules
	phonetic = applyFinalRules(phonetic, final2, lang, true)  // apply lang specific rules

	return phonetic

}

func applyFinalRules(phonetic string, finalRules []rule, languageArg uint64, strip bool) string {

	// optimization to save time

	if len(finalRules) == 0 {
		return phonetic
	}

	// expand the result

	phonetic = expand(phonetic)
	phoneticArray := strings.Split(phonetic, "|")

	var phonetic2 string
	for k := 0; k < len(phoneticArray); k++ {

		phonetic = phoneticArray[k]
		phonetic2 = ""

		phoneticx := normalizeLanguageAttributes(phonetic, true)
		for i := 0; i < utf8.RuneCountInString(phonetic); {
			found := false

			if substr(phonetic, i, 1) == "[" { // skip over language attribute
				attribStart := i
				i++
				for {
					vv := substr(phonetic, i, 1)
					i++
					if vv == "]" {
						attribEnd := i
						phonetic2 += substr(phonetic, attribStart, attribEnd-attribStart)
						break
					}
				}
				continue
			}

			var patternLength int
			for r := 0; r < len(finalRules); r++ {
				rule := finalRules[r]
				pattern := rule.pattern()
				patternLength = utf8.RuneCountInString(pattern)
				lcontext := rule.contextLeft()
				rcontext := rule.contextRight()

				right := "^" + rcontext
				left := lcontext + "$"

				// check to see if next sequence in phonetic matches the string in the rule
				if patternLength > utf8.RuneCountInString(phoneticx)-i || substr(phoneticx, i, patternLength) != pattern { // no match
					continue
				}

				// check that right context is satisfied
				if rcontext != "" {
					reg := regexp.MustCompile(right)
					if !reg.MatchString(substrFrom(phoneticx, i+patternLength)) {
						continue
					}
				}

				// check that left context is satisfied
				if lcontext != "" {
					reg := regexp.MustCompile(left)
					if !reg.MatchString(substr(phoneticx, 0, i)) {
						continue
					}
				}

				// // check to see if rule applies to languageArg (used only with "any" rules)
				// if (languageArg != "1") && (languagePos < len(rule)) {
				// 	language = rule[languagePos] // the required language(s) for this rule to apply
				// 	logical = rule[logicalPos]   // do we require ALL or ANY of the required languages
				// 	if logical == "ALL" {
				// 		// check to see if languageArg contains all the required languages
				// 		if (languageArg & language) != language {
				// 			continue
				// 		}
				// 	} else { // any
				// 		// check to see if languageArg contains at least one required language
				// 		if (languageArg & language) == 0 {
				// 			continue
				// 		}
				// 	}
				// }

				// check for incompatible attributes

				candidate := applyRuleIfCompatible(phonetic2, rule.phonetic(), languageArg)
				if candidate == "" {
					continue
				}
				phonetic2 = candidate

				found = true
				break
			}

			if !found { // character in name for which there is no subsitution in the table
				phonetic2 += substr(phonetic, i, 1)
				patternLength = 1
			}
			i += patternLength

		}
		phoneticArray[k] = expand(phonetic2)
	}
	phonetic = strings.Join(phoneticArray, "|")
	if strip {
		phonetic = normalizeLanguageAttributes(phonetic, true)
	}
	if strings.Index(phonetic, "|") != -1 {
		phonetic = "(" + removeDuplicateAlternates(phonetic) + ")"
	}
	return phonetic
}

func phoneticNumber(phonetic string, hash bool) string {
	if bracket := strings.Index(phonetic, "["); bracket != -1 {
		return substr(phonetic, 0, bracket)
	}
	return phonetic                                              // experimental !!!!
	phoneticLetters := "!bdfghjklmnNprsSt68vwzZxAa4oe5iI9uUEyQY" // true phonetic letters
	phoneticLetters += "1BCDEHJKLOTUVWX"                         // metaphonetic letters
	// dummy first letter, otherwise b would be treated as 0 and have no effect on result
	metaPhoneticLetters := "" // added letters to be used in finalxxx.php rules
	result := 0

	for i := 0; i < utf8.RuneCountInString(phonetic); i++ {
		var j int
		if substr(phonetic, i, 1) == "#" { // it's a meta phonetic letter
			if i == (utf8.RuneCountInString(phonetic) - 1) {
				panic("fatal error: invalid metaphonetic letter at position " + strconv.Itoa(i+1) + " in phonetic<br>")
			}
			i++
			j = strings.Index(metaPhoneticLetters, substr(phonetic, i, 1))
			if j != -1 {
				j += utf8.RuneCountInString(phoneticLetters)
			}
		} else {
			j = strings.Index(phoneticLetters, substr(phonetic, i, 1))
		}
		if j == -1 {
			panic("fatal error: invalid phonetic letter at position " + strconv.Itoa(i+1) + " in phonetic<br>")
		}
		result *= utf8.RuneCountInString(phoneticLetters) + utf8.RuneCountInString(metaPhoneticLetters)
		if hash {
			result = result & 0x7fffffff
		}
		result += j
	}
	r, err := hex.DecodeString(strconv.Itoa(result))
	if err != nil {
		panic(err)
	}

	return string(r)
}

func expand(phonetic string) string {
	altStart := strings.Index(phonetic, "(")
	if altStart == -1 {
		return normalizeLanguageAttributes(phonetic, false)
	}
	prefix := substr(phonetic, 0, altStart)
	altStart++ // get past the (
	altEnd := indexAt(phonetic, ")", altStart)
	altString := substr(phonetic, altStart, altEnd-altStart)
	altEnd++ // get past the )
	suffix := substrFrom(phonetic, altEnd)
	altArray := strings.Split(altString, "|")
	result := ""
	for i := 0; i < len(altArray); i++ {
		alt := altArray[i]
		alternate := expand(prefix + alt + suffix)
		if alternate != "" && alternate != "[0]" {
			if result != "" {
				result += "|"
			}
			result += alternate
		}
	}
	return result
}

func phoneticNumbersWithLeadingSpace(phonetic string) string {
	altStart := strings.Index(phonetic, "(")
	if altStart == -1 {
		return " " + phoneticNumber(phonetic, true)
	}
	prefix := substr(phonetic, 0, altStart)
	altStart++ // get past the (
	altEnd := indexAt(phonetic, ")", altStart)
	altString := substr(phonetic, altStart, altEnd-altStart)
	altEnd++ // get past the )
	suffix := substrFrom(phonetic, altEnd)
	altArray := strings.Split(altString, "|")
	result := ""
	for i := 0; i < len(altArray); i++ {
		alt := altArray[i]
		result += phoneticNumbersWithLeadingSpace(prefix + alt + suffix)
	}
	return result
}

func phoneticNumbers(phonetic string) string {
	phoneticArray := strings.Split(phonetic, "-") // for names with spaces in them
	result := ""
	for i := 0; i < len(phoneticArray); i++ {
		if i != 0 {
			result += " "
		}
		result += substrFrom(phoneticNumbersWithLeadingSpace(phoneticArray[i]), 1)
	}
	return result
}

func isPhoneticVowel(c string) bool {
	return (strings.Index("Aa4oe5iI9uUE", c) != -1)
}

func isAOTypeVowel(c string) bool {
	return (strings.Index("a4o59", c) != -1)
}

func isEITypeVowel(c string) bool {
	return (strings.Index("eiIy", c) != -1)
}

func isSZTypeConsonant(c string) bool {
	return (strings.Index("sSzZ", c) != -1)
}

func removeDuplicateAlternates(phonetic string) string {
	altString := phonetic
	altArray := strings.Split(altString, "|")

	result := "|"
	altcount := 0
	for i := 0; i < len(altArray); i++ {
		alt := altArray[i]
		if strings.Index(result, "|"+alt+"|") == -1 {
			result += alt + "|"
			altcount++
		}
	}

	result = substr(result, 1, utf8.RuneCountInString(result)-2) // remove leading and trailing |
	return result
}

func normalizeLanguageAttributes(text string, strip bool) string {
	// this is applied to a single alternative at a time -- not to a parenthisized list
	// it removes all embedded bracketed attributes, logically-ands them together, and places them at the end.

	// however if strip is true, this can indeed remove embedded bracketed attributes from a parenthesized list

	uninitialized := -1 // all 1's
	attrib := uninitialized
	for {
		bracketStart := strings.Index(text, "[")
		if bracketStart == -1 {
			break
		}
		bracketEnd := indexAt(text, "]", bracketStart)
		if bracketEnd == -1 {
			panic("fatal error: no closing square bracket: text=($text) strip=($strip)<br>")
		}

		v, err := strconv.Atoi(substr(text, bracketStart+1, bracketEnd-(bracketStart+1)))
		if err != nil {
			panic(err)
		}

		attrib &= v
		text = substr(text, 0, bracketStart) + substrFrom(text, bracketEnd+1)
	}
	if attrib == uninitialized || strip {
		return text
	} else if attrib == 0 {
		return "[0]" // means that the attributes were incompatible and there is no alternative here
	} else {
		return text + "[" + strconv.Itoa(attrib) + "]"
	}
}

func applyRuleIfCompatible(phonetic string, target string, languageArg uint64) string {

	// tests for compatible language rules
	// to do so, apply the rule, expand the results, and detect alternatives with incompatible attributes
	// then drop each alternative that has incompatible attributes and keep those that are compatible
	// if there are no compatible alternatives left, return false
	// otherwise return the compatible alternatives

	// apply the rule

	candidate := phonetic + target
	if strings.Index(candidate, "[") == -1 { // no attributes so we need test no further
		return candidate
	}

	// expand the result, converting incompatible attributes to [0]

	candidate = expand(candidate)
	candidateArray := strings.Split(candidate, "|")

	// drop each alternative that has incompatible attributes

	candidate = ""
	found := false

	for i := 0; i < len(candidateArray); i++ {
		thisCandidate := candidateArray[i]
		if languageArg != langAny {
			thisCandidate = normalizeLanguageAttributes(thisCandidate+"["+strconv.FormatUint(languageArg, 10)+"]", false)
		}
		if thisCandidate != "[0]" {
			//      if (candidate != "[0]") {
			found = true
			if candidate != "" {
				candidate += "|"
			}
			candidate += thisCandidate
		}
	}

	// return false if no compatible alternatives remain

	if !found {
		return ""
	}

	// return the result of applying the rule

	if strings.Index(candidate, "|") != -1 {
		candidate = "(" + candidate + ")"
	}
	return candidate

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
	lang uint64,
) (rules []rule, finalRules1 []rule, finalRules2 []rule) {
	switch mode {
	case Generic:
		rules = genRules[genLang(lang)]
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
