package beidermorse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_rules_applyTo(b *testing.B) {
	r := rule{
		pattern: "ge",
		phoneticRules: []token{
			{
				text:  "gE",
				langs: -1,
			},
			{
				text:  "xe",
				langs: 262144,
			},
			{
				text:  "hE",
				langs: 131072,
			},
			{
				text:  "je",
				langs: 512,
			},
			{
				text:  "Ze",
				langs: 32832,
			},
			{
				text:  "dZe",
				langs: 331808,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		r.applyTo("orange", 4)
	}
}

func Test_rules_applyTo(t *testing.T) {
	r1 := rule{
		pattern: "ge",
		phoneticRules: []token{
			{
				text:  "gE",
				langs: -1,
			},
			{
				text:  "xe",
				langs: 262144,
			},
			{
				text:  "hE",
				langs: 131072,
			},
			{
				text:  "je",
				langs: 512,
			},
			{
				text:  "Ze",
				langs: 32832,
			},
			{
				text:  "dZe",
				langs: 331808,
			},
		},
	}

	cases := []struct {
		name            string
		rule            rule
		input           string
		position        int
		expectedResult  []token
		expectedApplied bool
		expectedOffset  int
	}{
		{
			name:     "applied",
			rule:     r1,
			input:    "orange",
			position: 4,

			expectedResult: []token{
				{text: "gE", langs: langsUnitialized},
				{text: "xe", langs: 262144},
				{text: "hE", langs: 131072},
				{text: "je", langs: 512},
				{text: "Ze", langs: 32832},
				{text: "dZe", langs: 331808},
			},
			expectedApplied: true,
			expectedOffset:  2,
		},
		{
			name:     "not applied",
			rule:     r1,
			input:    "orange",
			position: 2,

			expectedResult:  nil,
			expectedApplied: false,
			expectedOffset:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, applied, offset := c.rule.applyTo(c.input, c.position)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedApplied, applied)
			assert.Equal(t, c.expectedOffset, offset)
		})
	}
}
