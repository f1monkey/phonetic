package cologne

import (
	"unicode"

	"golang.org/x/exp/slices"
)

type Encoder struct{}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) Encode(input string) string {
	inputRunes := make([]rune, 0, len(input)*2)

	for _, r := range input {
		r = unicode.ToUpper(r)
		replacers, ok := replaceMap[r]
		if ok {
			inputRunes = append(inputRunes, replacers...)
		} else {
			inputRunes = append(inputRunes, r)
		}
	}

	result := make([]rune, 0, len(inputRunes)*2)

	// encode letter by letter from left to right according to the conversion table
	var prev, next rune
	for i := 0; i < len(inputRunes); i++ {
		if i < len(inputRunes)-1 {
			next = inputRunes[i+1]
		} else {
			next = 0
		}

		rr := inputRunes[i]
		applicableRules, ok := rules[rr]
		if !ok {
			continue
		}

		for _, ar := range applicableRules {
			if ar.matches(i, prev, next) {
				result = ar.apply(result)
				break
			}
		}

		prev = rr
	}

	// remove all digits occurring more than once next to each other
	for i := 1; i < len(result); i++ {
		if result[i-1] == result[i] {
			result = slices.Delete(result, i, i+1)
			i--
		}
	}

	// remove all code "0" except at the beginning
	for i := 1; i < len(result); i++ {
		if result[i] == '0' {
			result = slices.Delete(result, i, i+1)
			i--
		}
	}

	return string(result)
}
