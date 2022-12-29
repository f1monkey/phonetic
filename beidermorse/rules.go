package beidermorse

import "regexp"

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
	pattern      string
	leftContext  *regexp.Regexp
	rightContext *regexp.Regexp
	phonetic     string
}

type finalRules struct {
	approx finalRule
	exact  finalRule
}

type finalRule struct {
	first  []rule
	second map[uint64][]rule
}

type langRule struct {
	pattern *regexp.Regexp
	langs   uint64
	accept  bool
}
