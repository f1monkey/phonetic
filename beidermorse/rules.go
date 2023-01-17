package beidermorse

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Mode which name mode to use for matching
type Mode string

const (
	Generic   Mode = "generic"   // for general usage
	Ashkenazi Mode = "ashkenazi" // for ashkenazi names
	Sephardic Mode = "sephardic" // for sephardic names
)

func (m Mode) Valid() bool {
	return m == Generic || m == Ashkenazi || m == Sephardic
}

// Ruleset exact or approximate matching
type Ruleset string

const (
	Exact  Ruleset = "exact"  // exact matching rules
	Approx Ruleset = "approx" // approx matching (results in more tokens)
)

func (r Ruleset) Valid() bool {
	return r == Exact || r == Approx
}

type rule struct {
	pattern       string
	leftContext   *ruleMatcher
	rightContext  *ruleMatcher
	phoneticRules tokens
}

func (r rule) applyTo(input string, position int) (result []token, applied bool, offset int) {
	patternLength := len([]rune(r.pattern))
	offset = 1

	if len(r.phoneticRules) == 0 {
		return
	}

	if patternLength > utf8.RuneCountInString(input)-position || substr(input, position, patternLength) != r.pattern { // no match
		return
	}

	if r.rightContext != nil {
		if !r.rightContext.matches(substrFrom(input, position+patternLength)) {
			return
		}
	}

	if r.leftContext != nil {
		if !r.leftContext.matches(substrTo(input, position)) {
			return
		}
	}

	result = r.phoneticRules
	applied = true
	offset = patternLength

	return
}

type token struct {
	text  string
	langs languageID
}

type tokens []token

func (t tokens) merge(lang languageID, src ...tokens) tokens {
	if len(src) == 0 {
		return t
	}

	result := t
	i := 1
	for i < len(src)+1 { // use while-loop instead of recursion here
		newResult := make([]token, 0, len(result)*len(src[i-1]))
		for _, r1 := range result {
			for _, r2 := range src[i-1] {
				lang := lang.merge(r1.langs, r2.langs)
				if lang == langsInvalid {
					continue
				}

				newResult = append(newResult, token{
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

type finalRules struct {
	approx finalRule
	exact  finalRule
}

type finalRule struct {
	first  []rule
	second map[languageID][]rule
}

type langRule struct {
	match  ruleMatcher
	langs  languageID
	accept bool
}

type ruleMatcher struct {
	matchEmptyString bool
	pattern          *regexp.Regexp
	prefix           string
	suffix           string
	contains         string
}

func (r ruleMatcher) matches(str string) bool {
	if r.matchEmptyString && len(str) == 0 {
		return true
	}

	if r.contains != "" {
		return strings.Contains(str, r.contains)
	}

	if r.prefix != "" {
		return strings.HasPrefix(str, r.prefix)
	}

	if r.suffix != "" {
		return strings.HasSuffix(str, r.suffix)
	}

	if r.pattern != nil {
		return r.pattern.MatchString(str)
	}

	return false
}

type languageID int64

const (
	langsUnitialized languageID = -1
	langsInvalid     languageID = 0
	langsAny         languageID = 1
)

func (l languageID) merge(src ...languageID) languageID {
	if len(src) == 0 {
		return l
	}

	result := l
	for _, lang := range src {
		if result == langsUnitialized || result == langsAny {
			result = lang
			continue
		}
		result &= lang
	}

	return result
}
