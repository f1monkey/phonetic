package beidermorse

import (
	"regexp"
	"strings"
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

func (r rule) applyTo(input string, position int) (result []token, offset int) {
	patternLength := len([]rune(r.pattern))
	inputLength := len([]rune(input))
	offset = 1

	if len(r.phoneticRules) == 0 {
		return
	}

	if patternLength > inputLength-position || substr(input, position, patternLength) != r.pattern { // no match
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
	offset = patternLength

	return
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
