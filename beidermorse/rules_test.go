package beidermorse

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Benchmark_rules_apply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRules[1].apply(tokens{
			{text: "orange", langs: 1047280},
		}, 1047280, false)
	}
}

func Test_rules_apply(t *testing.T) {
	cases := []struct {
		name        string
		src         tokens
		lang        languageID
		rules       rules
		ignoreLangs bool
		expected    tokens
	}{
		{
			name: "orange_approx_lang=1_stage=1",
			src: tokens{
				{text: "orange", langs: 1047280},
			},
			lang:  1047280,
			rules: genRules[1],
			expected: tokens{
				{text: "OrAngE", langs: 1047280},
				{text: "OrAnxe", langs: 262144},
				{text: "OrAnhE", langs: 131072},
				{text: "OrAnje", langs: 512},
				{text: "OrAnZe", langs: 32832},
				{text: "OrAndZe", langs: 331808},
				{text: "PrAngE", langs: 16384},
			},
			ignoreLangs: false,
		},
		{
			name: "orange_approx_lang=1_stage=2",
			src: tokens{
				{text: "OrAngE", langs: 1047280},
				{text: "OrAnxe", langs: 262144},
				{text: "OrAnhE", langs: 131072},
				{text: "OrAnje", langs: 512},
				{text: "OrAnZe", langs: 32832},
				{text: "OrAndZe", langs: 331808},
				{text: "PrAngE", langs: 16384},
			},
			lang:  1047280,
			rules: genFinalRules.approx.first,
			expected: tokens{
				{text: "OrAngE", langs: 1047280},
				{text: "OrAnxe", langs: 262144},
				{text: "OrAnE", langs: 131072},
				{text: "OrAnie", langs: 512},
				{text: "OrAnze", langs: 32832},
				{text: "OrAnze", langs: 331808},
				{text: "PrAngE", langs: 16384},
			},
			ignoreLangs: false,
		},
		{
			name: "orange_stage=3_lang=1",
			src: tokens{
				{text: "OrAngE", langs: 1047280},
				{text: "OrAnxe", langs: 262144},
				{text: "OrAnE", langs: 131072},
				{text: "OrAnie", langs: 512},
				{text: "OrAnze", langs: 32832},
				{text: "OrAnze", langs: 331808},
				{text: "PrAngE", langs: 16384},
			},
			lang:  1047280,
			rules: genFinalRules.approx.second[1],
			expected: tokens{
				{text: "orangi", langs: 1},
				{text: "oragi", langs: 1},
				{text: "orongi", langs: 1},
				{text: "orogi", langs: 1},
				{text: "orYngi", langs: 1},
				{text: "Yrangi", langs: 1},
				{text: "Yrongi", langs: 1},
				{text: "YrYngi", langs: 1},
				{text: "oranxi", langs: 1},
				{text: "oronxi", langs: 1},
				{text: "orani", langs: 1},
				{text: "oroni", langs: 1},
				{text: "oranii", langs: 1},
				{text: "oronii", langs: 1},
				{text: "oranzi", langs: 1},
				{text: "oronzi", langs: 1},
				{text: "urangi", langs: 1},
				{text: "urongi", langs: 1},
			},
			ignoreLangs: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.rules.apply(c.src, c.lang, c.ignoreLangs)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_rule_applyTo(b *testing.B) {
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

func Test_rule_applyTo(t *testing.T) {
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
		name           string
		rule           rule
		input          string
		position       int
		expectedResult []token
		expectedOffset int
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
			expectedOffset: 2,
		},
		{
			name:     "not applied",
			rule:     r1,
			input:    "orange",
			position: 2,

			expectedResult: nil,
			expectedOffset: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, offset := c.rule.applyTo(c.input, c.position)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedOffset, offset)
		})
	}
}

func Test_containsAt(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		from     int
		expected bool
	}{
		{haystack: "апельсин", needle: "пе", from: 0, expected: false},
		{haystack: "апельсин", needle: "пе", from: 1, expected: true},
		{haystack: "апельсин", needle: "пе", from: 2, expected: false},
		{haystack: "апельсин", needle: "пе", from: 7, expected: false},
		{haystack: "апельсин", needle: "пе", from: 8, expected: false},
		{haystack: "апельсин", needle: "пе", from: 9, expected: false},

		{haystack: "orange", needle: "or", from: 0, expected: true},
		{haystack: "orange", needle: "or", from: 1, expected: false},
		{haystack: "orange", needle: "or", from: 5, expected: false},
		{haystack: "orange", needle: "or", from: 6, expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := containsAt(c.haystack, c.needle, c.from)
			require.Equal(t, c.expected, result)
		})
	}
}
