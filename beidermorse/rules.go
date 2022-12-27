package beidermorse

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
	patterns [4]string
}

func (r *rule) pattern() string {
	return r.patterns[0]
}

func (r *rule) contextLeft() string {
	return r.patterns[1]
}

func (r *rule) contextRight() string {
	return r.patterns[2]
}

func (r *rule) phonetic() string {
	return r.patterns[3]
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
	pattern string
	langs   uint64
	accept  bool
}
