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
	phonetic      string
	phoneticRules []phoneticRule
}

type phoneticRule struct {
	text  string
	langs int64
}

type finalRules struct {
	approx finalRule
	exact  finalRule
}

type finalRule struct {
	first  []rule
	second map[int64][]rule
}

type langRule struct {
	match  ruleMatcher
	langs  int64
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
