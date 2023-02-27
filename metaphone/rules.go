package metaphone

type firstCharRule struct {
	secondRune  rune
	replaceWith rune
}

var firstCharRules = map[rune][]firstCharRule{
	// KN, GN, PN => N
	'K': {firstCharRule{'N', 'N'}},
	'G': {firstCharRule{'N', 'N'}},
	'P': {firstCharRule{'N', 'N'}},
	// AE => E
	'A': {firstCharRule{'E', 'E'}},
	// WR => R, WH => H
	'W': {
		firstCharRule{'R', 'R'},
		firstCharRule{'H', 'H'},
	},
	// X => S
	'X': {firstCharRule{0, 'S'}},
}

type ruleFunc func(runes []rune, index int, result []rune) ([]rune, int)

var rules = map[rune]ruleFunc{
	'A': vowelFunc,
	'B': func(runes []rune, index int, result []rune) ([]rune, int) {
		if isPreviousChar(runes, index, 'M') &&
			isLastChar(runes, index) {
			// B is silent if word ends in MB
			return result, index
		}
		return append(result, 'B'), index
	},
	'C': func(runes []rune, index int, result []rune) ([]rune, int) {
		if isPreviousChar(runes, index, 'S') &&
			!isLastChar(runes, index) &&
			isFrontv(runes, index+1) {
			// discard if SCI, SCE or SCY
			return result, index
		}
		if regionMatch(runes, index, Cregions[0]) {
			// "CIA" -> X
			return append(result, 'X'), index
		}
		if !isLastChar(runes, index) &&
			isFrontv(runes, index+1) {
			// CI,CE,CY -> S
			return append(result, 'S'), index
		}
		if isPreviousChar(runes, index, 'S') &&
			isNextChar(runes, index, 'H') {
			// SCH->SK
			return append(result, 'K'), index
		}
		if isNextChar(runes, index, 'H') { // detect CH
			if index == 0 &&
				len(runes) >= 3 &&
				isVowel(runes, 2) {
				// CH consonant -> K consonant
				result = append(result, 'K')
			} else {
				// CHvowel -> X
				result = append(result, 'X')

			}
		} else {
			result = append(result, 'K')
		}
		return result, index
	},
	'D': func(runes []rune, index int, result []rune) ([]rune, int) {
		if !isLastChar(runes, index+1) &&
			isNextChar(runes, index, 'G') &&
			isFrontv(runes, index+2) {
			// DGE DGI DGY -> J
			result = append(result, 'J')
			index += 2
		} else {
			result = append(result, 'T')
		}
		return result, index
	},
	'E': vowelFunc,
	'F': defaultFunc,
	'G': func(runes []rune, index int, result []rune) ([]rune, int) {
		if isLastChar(runes, index+1) &&
			isNextChar(runes, index, 'H') {
			// GH silent at end or before consonant
			return result, index
		}
		if !isLastChar(runes, index+1) &&
			isNextChar(runes, index, 'H') &&
			!isVowel(runes, index+2) {
			return result, index // silent G
		}
		if index > 0 &&
			(regionMatch(runes, index, Gregions[0]) ||
				regionMatch(runes, index, Gregions[1])) {
			return result, index // silent G
		}
		if !isLastChar(runes, index) &&
			isFrontv(runes, index+1) {
			result = append(result, 'J')
		} else {
			result = append(result, 'K')
		}
		return result, index
	},
	'H': func(runes []rune, index int, result []rune) ([]rune, int) {
		if isLastChar(runes, index) {
			return result, index // terminal H
		}
		if index > 0 && isVarson(runes, index-1) {
			return result, index
		}
		if isVowel(runes, index+1) {
			result = append(result, 'H') // Hvowel
		}
		return result, index
	},
	'I': vowelFunc,
	'J': defaultFunc,
	'K': func(runes []rune, index int, result []rune) ([]rune, int) {
		char := runes[index]
		if index > 0 { // not initial
			if !isPreviousChar(runes, index, 'C') {
				result = append(result, char)
			}
		} else {
			result = append(result, char) // initial K
		}

		return result, index
	},
	'L': defaultFunc,
	'M': defaultFunc,
	'N': defaultFunc,
	'O': vowelFunc,
	'P': func(runes []rune, index int, result []rune) ([]rune, int) {
		if isNextChar(runes, index, 'H') {
			// PH -> F
			result = append(result, 'F')
		} else {
			result = append(result, 'P')
		}
		return result, index
	},
	'Q': func(runes []rune, index int, result []rune) ([]rune, int) {
		return append(result, 'K'), index
	},
	'R': defaultFunc,
	'S': func(runes []rune, index int, result []rune) ([]rune, int) {
		if regionMatch(runes, index, Sregions[0]) ||
			regionMatch(runes, index, Sregions[1]) ||
			regionMatch(runes, index, Sregions[2]) {
			result = append(result, 'X')
		} else {
			result = append(result, 'S')
		}
		return result, index
	},
	'T': func(runes []rune, index int, result []rune) ([]rune, int) {
		if regionMatch(runes, index, Tregions[0]) ||
			regionMatch(runes, index, Tregions[1]) {
			return append(result, 'X'), index
		}
		if regionMatch(runes, index, Tregions[2]) {
			// Silent if in "TCH"
			return result, index
		}
		if regionMatch(runes, index, Tregions[3]) {
			result = append(result, '0') // theta
		} else {
			result = append(result, 'T')
		}
		return result, index
	},
	'U': vowelFunc,
	'V': func(runes []rune, index int, result []rune) ([]rune, int) {
		return append(result, 'F'), index
	},
	'W': func(runes []rune, index int, result []rune) ([]rune, int) {
		if !isLastChar(runes, index) &&
			isVowel(runes, index+1) {
			result = append(result, 'W')
		}
		return result, index
	},
	'X': func(runes []rune, index int, result []rune) ([]rune, int) {
		return append(result, 'K', 'S'), index
	},
	'Y': func(runes []rune, index int, result []rune) ([]rune, int) {
		if !isLastChar(runes, index) &&
			isVowel(runes, index+1) {
			result = append(result, 'Y')
		}
		return result, index
	},
	'Z': func(runes []rune, index int, result []rune) ([]rune, int) {
		return append(result, 'S'), index
	},
}

var Cregions = [][]rune{
	[]rune("CIA"),
}

var Gregions = [][]rune{
	[]rune("GN"),
	[]rune("GNED"),
}

var Sregions = [][]rune{
	[]rune("SH"),
	[]rune("SIO"),
	[]rune("SIA"),
}

var Tregions = [][]rune{
	[]rune("TIA"),
	[]rune("TIO"),
	[]rune("TCH"),
	[]rune("TH"),
}

var defaultFunc ruleFunc = func(runes []rune, index int, result []rune) ([]rune, int) {
	return append(result, runes[index]), index
}

// remove vowels if they are not at first positions
var vowelFunc ruleFunc = func(runes []rune, index int, result []rune) ([]rune, int) {
	if index == 0 {
		result = append(result, runes[index])
	}
	return result, index
}
