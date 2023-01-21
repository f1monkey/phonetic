package caverphone2

import (
	"bytes"
	"strings"
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
	for i := range rules {
		if rules[i].regexp != nil {
			inputBytes = rules[i].regexp.ReplaceAll(inputBytes, rules[i].replaceWith)
		} else {
			inputBytes = bytes.ReplaceAll(inputBytes, rules[i].pattern, rules[i].replaceWith)
		}
	}

	inputBytes = append(inputBytes, suffix...)

	return string(inputBytes[:len(suffix)])
}
