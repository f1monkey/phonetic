package bmpm

type rule struct {
	patterns []string
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
	langs   int
	accept  bool
}
