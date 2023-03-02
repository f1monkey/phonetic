package metaphone

import (
	"unicode"

	"golang.org/x/exp/slices"
)

type Encoder struct {
	maxLen int
}

func NewEncoder(opts ...EncoderOption) *Encoder {
	result := &Encoder{
		maxLen: 4,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

// SetOption set encoder option. Method is not safe for concurrent usage
func (e *Encoder) SetOption(opt EncoderOption) {
	opt(e)
}

// EncoderOption func to provide encoder configuration parameter
type EncoderOption func(e *Encoder)

// WithMaxLen Set max length of the generated token
func WithMaxLen(maxLen int) EncoderOption {
	return func(e *Encoder) {
		e.maxLen = maxLen
	}
}

func (e *Encoder) Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	runes := filter([]rune(input))
	if len(runes) <= 1 {
		return string(runes)
	}
	runes = transformFirstCharacters(runes) // fix first characters in input

	maxLen := e.maxLen
	if maxLen < 0 {
		maxLen = len(input) * 3
	}
	result := make([]rune, 0, maxLen)

	index := 0
	for index < len(runes) && len(result) < maxLen {
		char := runes[index]

		// remove duplicate characters
		if char != 'C' && isPreviousChar(runes, index, char) {
			index++
			continue
		}

		if f, ok := rules[char]; ok {
			result, index = f(runes, index, result)
		}
		index++
	}

	if len(result) > maxLen {
		result = result[:maxLen]
	}

	return string(result)
}

func filter(input []rune) []rune {
	if len(input) == 0 {
		return input
	}

	for i := 0; i < len(input); i++ {
		if _, ok := allowedSymbols[input[i]]; !ok {
			input = slices.Delete(input, i, i+1)
			i--
			continue
		}
		input[i] = unicode.ToUpper(input[i])
	}

	return input
}

var allowedSymbols = map[rune]struct{}{
	'a': {},
	'b': {},
	'c': {},
	'd': {},
	'e': {},
	'f': {},
	'g': {},
	'h': {},
	'i': {},
	'j': {},
	'k': {},
	'l': {},
	'm': {},
	'n': {},
	'o': {},
	'p': {},
	'q': {},
	'r': {},
	's': {},
	't': {},
	'u': {},
	'v': {},
	'w': {},
	'x': {},
	'y': {},
	'z': {},
	'A': {},
	'B': {},
	'C': {},
	'D': {},
	'E': {},
	'F': {},
	'G': {},
	'H': {},
	'I': {},
	'J': {},
	'K': {},
	'L': {},
	'M': {},
	'N': {},
	'O': {},
	'P': {},
	'Q': {},
	'R': {},
	'S': {},
	'T': {},
	'U': {},
	'V': {},
	'W': {},
	'X': {},
	'Y': {},
	'Z': {},
}

func transformFirstCharacters(input []rune) []rune {
	if rule, ok := firstCharRules[input[0]]; ok {
		for _, r := range rule {
			if r.secondRune == 0 {
				input[0] = r.replaceWith
				return input
			} else if r.secondRune == input[1] {
				input = slices.Delete(input, 0, 1)
				input[0] = r.replaceWith
				return input
			}
		}
	}

	return input
}

func isPreviousChar(input []rune, index int, char rune) bool {
	if index-1 < 0 {
		return false
	}
	return input[index-1] == char
}

func isNextChar(input []rune, index int, char rune) bool {
	if index+1 < 0 {
		return false
	}

	if index+1 > len(input)-1 {
		return false
	}

	return input[index+1] == char
}

func isLastChar(input []rune, index int) bool {
	return index == len(input)-1
}

func regionMatch(input []rune, index int, chars []rune) bool {
	if len(input) == 0 || len(chars) == 0 {
		return false
	}

	j := 0
	for i := index; i < len(input) && j < len(chars); i++ {
		if input[i] != chars[j] {
			return false
		}
		j++
	}

	return j == len(chars)
}

var vowels = []rune("AEIOU")

func isVowel(input []rune, index int) bool {
	if index < 0 || index >= len(input) {
		return false
	}

	char := input[index]
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == char {
			return true
		}
	}
	return false
}

var frontv = []rune("EIY") // Variable used in Metaphone algorithm

func isFrontv(input []rune, index int) bool {
	if index < 0 || index >= len(input) {
		return false
	}

	char := input[index]
	for i := 0; i < len(frontv); i++ {
		if frontv[i] == char {
			return true
		}
	}
	return false
}

var varson = []rune("CSPTG") // Variable used in Metaphone algorithm

func isVarson(input []rune, index int) bool {
	if index < 0 || index >= len(input) {
		return false
	}

	char := input[index]
	for i := 0; i < len(varson); i++ {
		if varson[i] == char {
			return true
		}
	}
	return false
}
