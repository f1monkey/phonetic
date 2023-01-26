package cologne

var replaceMap = map[rune][]rune{
	'Ä': {'A'},
	'Ö': {'O'},
	'Ü': {'U'},
	'ß': {'S', 'S'},
}

type rule struct {
	initial   bool
	before    []rune
	notBefore []rune
	after     []rune
	notAfter  []rune
	result    []rune
}

func (r *rule) matches(pos int, prev rune, next rune) bool {
	if r.initial && pos != 0 {
		return false
	}

	if prev > 0 && len(r.after) > 0 {
		matched := false
		for _, rr := range r.after {
			if rr == prev {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	if prev > 0 && len(r.notAfter) > 0 {
		for _, rr := range r.notAfter {
			if rr == prev {
				return false
			}
		}
	}

	if next > 0 && len(r.before) > 0 {
		matched := false
		for _, rr := range r.before {
			if rr == next {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	if next > 0 && len(r.notBefore) > 0 {
		for _, rr := range r.notBefore {
			if rr == next {
				return false
			}
		}
	}

	return true
}

func (r *rule) apply(src []rune) []rune {
	return append(src, r.result...)
}

var rules = map[rune][]*rule{
	'A': {{result: []rune{'0'}}},
	'B': {{result: []rune{'1'}}},
	'C': {
		{initial: true, before: []rune{'A', 'H', 'K', 'L', 'O', 'Q', 'R', 'U', 'X'}, result: []rune{'4'}},
		{before: []rune{'A', 'H', 'K', 'O', 'Q', 'U', 'X'}, notAfter: []rune{'S', 'Z'}, result: []rune{'4'}},
		{after: []rune{'S', 'Z'}, result: []rune{'8'}},
		{initial: true, notBefore: []rune{'A', 'H', 'K', 'L', 'O', 'Q', 'R', 'U', 'X'}, result: []rune{'8'}},
		{notBefore: []rune{'A', 'H', 'K', 'O', 'Q', 'U', 'X'}, result: []rune{'8'}},
	},
	'D': {
		{notBefore: []rune{'C', 'S', 'Z'}, result: []rune{'2'}},
		{before: []rune{'C', 'S', 'Z'}, result: []rune{'8'}},
	},
	'E': {{result: []rune{'0'}}},
	'F': {{result: []rune{'3'}}},
	'G': {{result: []rune{'4'}}},
	'H': nil, // no rules
	'I': {{result: []rune{'0'}}},
	'J': {{result: []rune{'0'}}},
	'K': {{result: []rune{'4'}}},
	'L': {{result: []rune{'5'}}},
	'M': {{result: []rune{'6'}}},
	'N': {{result: []rune{'6'}}},
	'O': {{result: []rune{'0'}}},
	'P': {
		{before: []rune{'H'}, result: []rune{'3'}},
		{notBefore: []rune{'H'}, result: []rune{'1'}},
	},
	'Q': {{result: []rune{'4'}}},
	'R': {{result: []rune{'7'}}},
	'S': {{result: []rune{'8'}}},
	'T': {
		{notBefore: []rune{'C', 'S', 'Z'}, result: []rune{'2'}},
		{before: []rune{'C', 'S', 'Z'}, result: []rune{'8'}},
	},
	'U': {{result: []rune{'0'}}},
	'V': {{result: []rune{'3'}}},
	'W': {{result: []rune{'3'}}},
	'X': {
		{notAfter: []rune{'C', 'K', 'Q'}, result: []rune{'4', '8'}},
		{after: []rune{'C', 'K', 'Q'}, result: []rune{'8'}},
	},
	'Y': {{result: []rune{'0'}}},
	'Z': {{result: []rune{'8'}}},
}
