package bmpm

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
	second []secondFinalRule
}

type secondFinalRule struct {
	langs int
	rules []rule
}

type langRule struct {
	pattern string
	langs   uint64
	accept  bool
}
