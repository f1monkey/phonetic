package beidermorse

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Benchmark_tokens_merge(b *testing.B) {
	t := tokens{
		{text: "k", langs: 1047288},
		{text: "ts", langs: 16392},
		{text: "dZ", langs: 524288},
	}

	for i := 0; i < b.N; i++ {
		t.merge(langsAny, tokens{
			{text: "O", langs: -1},
			{text: "P", langs: 16384},
		})
	}
}

func Test_tokens_merge(t *testing.T) {
	cases := []struct {
		tokens   tokens
		src      []tokens
		expected tokens
	}{
		{
			tokens: tokens{
				{text: "O", langs: -1},
				{text: "P", langs: 16384},
			},
			src: nil,
			expected: tokens{
				{text: "O", langs: -1},
				{text: "P", langs: 16384},
			},
		},
		{
			tokens: tokens{
				{text: "k", langs: 1047288},
				{text: "ts", langs: 16392},
				{text: "dZ", langs: 524288},
			},
			src: []tokens{
				{
					{text: "O", langs: -1},
					{text: "P", langs: 16384},
				},
			},
			expected: tokens{
				{text: "kO", langs: 1047288},
				{text: "kP", langs: 16384},
				{text: "tsO", langs: 16392},
				{text: "tsP", langs: 16384},
				{text: "dZO", langs: 524288},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := c.tokens.merge(langsAny, c.src...)
			require.Equal(t, c.expected, result)
		})
	}
}
