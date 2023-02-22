package bmpm

import (
	"strings"

	"github.com/f1monkey/phonetic/internal/exrunes"
)

type DetectLangFunc = func(input string) Lang

func redoLanguage(input string, mode Mode, accuracy Accuracy, ruleset Ruleset, concat bool, buf *exrunes.Buffer) *Tokens {
	// we can do a better job of determining the language now that multiple names have been split
	languageArg := ruleset.DetectLang(input)
	return MakeTokens(input, mode, accuracy, ruleset, languageArg, concat, buf)
}

var firstList = []string{"de la", "van der", "van den"}

type Ruleset struct {
	Main       Rules
	Final1     Rules
	Final2     Rules
	Discards   []string
	DetectLang DetectLangFunc
}

func MakeTokens(input string, mode Mode, accuracy Accuracy, ruleset Ruleset, lang Lang, concat bool, buf *exrunes.Buffer) *Tokens {
	input = prepareInput(input, mode) // returns at max two words

	// process a multiword name of form X Y // @todo
	if space := strings.Index(input, " "); space != -1 { // number of words is exactly two
		if concat { // exact matches
			// X Y => XY
			input = strings.ReplaceAll(input, " ", "") // concatenate the separate words of a name
		} else { // number of words is exactly two
			word1 := substrTo(input, space)
			word2 := substrFrom(input, space+1)
			if inArray(ruleset.Discards, word1) {
				// X Y => Y and XY
				results := redoLanguage(word2, mode, accuracy, ruleset, concat, buf)
				results.Items = append(results.Items, redoLanguage(word1+word2, mode, accuracy, ruleset, concat, buf).Items...)
				return results
			} else { // first word is not in list
				// X Y => X, Y, and XY
				results := redoLanguage(word1, mode, accuracy, ruleset, concat, buf)
				results.Items = append(results.Items, redoLanguage(word2, mode, accuracy, ruleset, concat, buf).Items...)
				results.Items = append(results.Items, redoLanguage(word1+word2, mode, accuracy, ruleset, concat, buf).Items...)
				return results
			}
		}
	}

	result := &Tokens{Items: []Token{{
		ID:    buf.Add([]rune(input)),
		Langs: lang,
	}}, Buf: buf}

	// apply main set of rules
	result = ruleset.Main.Apply(result, lang, false)
	// apply common final rules
	result = ruleset.Final1.Apply(result, lang, false)
	// apply lang specific final rules
	result = ruleset.Final2.Apply(result, lang, true)

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
