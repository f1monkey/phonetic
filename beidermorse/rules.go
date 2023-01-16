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
	phoneticRules []phonetic
}

func (r rule) applyTo(input string, languageArg languageID, position int) (result []phonetic, applied bool, offset int) {
	patternLength := len([]rune(r.pattern))
	offset = 1

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

	result = mergeTokenGroups(languageArg, r.phoneticRules)
	if len(result) != 0 {
		applied = true
		offset = patternLength
	}

	return
}

type phonetic struct {
	text  string
	langs languageID
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
