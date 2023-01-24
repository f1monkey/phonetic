package caverphone2

import (
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
	inputBytes = filter.ReplaceAll(inputBytes, nil)

	for i := range rules {
		if rules[i].regexp != nil {
			inputBytes = rules[i].regexp.ReplaceAll(inputBytes, rules[i].to)
		} else if rules[i].substr != nil {
			inputBytes = bytesReplaceAll(inputBytes, rules[i].substr, rules[i].to)
		} else if rules[i].prefix != nil {
			inputBytes = bytesReplacePrefix(inputBytes, rules[i].prefix, rules[i].to)
		} else if rules[i].suffix != nil {
			inputBytes = bytesReplaceSuffix(inputBytes, rules[i].suffix, rules[i].to)
		}
	}

	inputBytes = append(inputBytes, suffix...)

	return string(inputBytes[:len(suffix)])
}
