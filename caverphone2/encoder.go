package caverphone2

import (
	"strings"

	"github.com/f1monkey/phonetic/internal/exbytes"
	"golang.org/x/exp/slices"
)

var suffix = []byte("1111111111")

type Encoder struct{}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) Encode(input string) string {
	if len(input) == 0 {
		return string(suffix)
	}

	inputBytes := []byte(strings.ToLower(input))
	inputBytes = filterInput(inputBytes)

	for i := range rules {
		if rules[i].regexp != nil {
			inputBytes = rules[i].regexp.ReplaceAll(inputBytes, rules[i].to)
		} else if rules[i].substr != nil {
			inputBytes = exbytes.ReplaceAll(inputBytes, rules[i].substr, rules[i].to)
		} else if rules[i].prefix != nil {
			inputBytes = exbytes.ReplacePrefix(inputBytes, rules[i].prefix, rules[i].to)
		} else if rules[i].suffix != nil {
			inputBytes = exbytes.ReplaceSuffix(inputBytes, rules[i].suffix, rules[i].to)
		} else if rules[i].sequence != nil {
			inputBytes = exbytes.ReplaceSequence(inputBytes, rules[i].sequence, rules[i].to)
		}
	}

	inputBytes = append(inputBytes, suffix...)

	return string(inputBytes[:len(suffix)])
}

func filterInput(input []byte) []byte {
	for i := 0; i < len(input); i++ {
		if _, ok := availCharacters[input[i]]; !ok {
			input = slices.Delete(input, i, i+1)
			i--
		}
	}

	return input
}

var availCharacters = map[byte]struct{}{
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
}
