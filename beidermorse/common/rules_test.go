package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @todo fix this test

// func Benchmark_rules_apply(b *testing.B) {
// 	t := Tokens{
// 		{Text: Runes("orange"), Langs: 1047280},
// 	}
// 	for i := 0; i < b.N; i++ {
// 		genRules[1].apply(t, 1047280, false)
// 	}
// }

// func Test_Rules_Apply(t *testing.T) {
// 	cases := []struct {
// 		name        string
// 		src         Tokens
// 		lang        Lang
// 		rules       Rules
// 		ignoreLangs bool
// 		expected    Tokens
// 	}{
// 		{
// 			name: "test_approx_lang=1_stage=1",
// 			src: Tokens{
// 				{Text: Runes("test"), Langs: 1047288},
// 			},
// 			lang:  1047288,
// 			rules: genRules[1],
// 			expected: Tokens{
// 				{Text: Runes("tESt"), Langs: 128},
// 				{Text: Runes("tEst"), Langs: 1047288},
// 			},
// 			ignoreLangs: false,
// 		},
// 		{
// 			name: "test_approx_lang=1_stage=2",
// 			src: Tokens{
// 				{Text: Runes("tESt"), Langs: 128},
// 				{Text: Runes("tEst"), Langs: 1047288},
// 			},
// 			lang:  1047288,
// 			rules: genFinalRules.approx.first,
// 			expected: Tokens{
// 				{Text: Runes("tEst"), Langs: 128},
// 				{Text: Runes("tEst"), Langs: 1047288},
// 			},
// 			ignoreLangs: false,
// 		},
// 		{
// 			name: "test_approx_lang=1_stage=3",
// 			src: Tokens{
// 				{Text: Runes("tEst"), Langs: 128},
// 				{Text: Runes("tEst"), Langs: 1047288},
// 			},
// 			lang:  1047288,
// 			rules: genFinalRules.approx.second[1],
// 			expected: Tokens{
// 				{Text: Runes("tist"), Langs: 1},
// 				{Text: Runes("tYst"), Langs: 1},
// 				{Text: Runes("tis"), Langs: 1},
// 				{Text: Runes("tit"), Langs: 1},
// 				{Text: Runes("ti"), Langs: 1},
// 			},
// 			ignoreLangs: true,
// 		},
// 		{
// 			name: "orange_approx_lang=1_stage=1",
// 			src: Tokens{
// 				{Text: Runes("orange"), Langs: 1047280},
// 			},
// 			lang:  1047280,
// 			rules: genRules[1],
// 			expected: Tokens{
// 				{Text: Runes("OrAngE"), Langs: 1047280},
// 				{Text: Runes("OrAnxe"), Langs: 262144},
// 				{Text: Runes("OrAnhE"), Langs: 131072},
// 				{Text: Runes("OrAnje"), Langs: 512},
// 				{Text: Runes("OrAnZe"), Langs: 32832},
// 				{Text: Runes("OrAndZe"), Langs: 331808},
// 				{Text: Runes("PrAngE"), Langs: 16384},
// 			},
// 			ignoreLangs: false,
// 		},
// 		{
// 			name: "orange_approx_lang=1_stage=2",
// 			src: Tokens{
// 				{Text: Runes("OrAngE"), Langs: 1047280},
// 				{Text: Runes("OrAnxe"), Langs: 262144},
// 				{Text: Runes("OrAnhE"), Langs: 131072},
// 				{Text: Runes("OrAnje"), Langs: 512},
// 				{Text: Runes("OrAnZe"), Langs: 32832},
// 				{Text: Runes("OrAndZe"), Langs: 331808},
// 				{Text: Runes("PrAngE"), Langs: 16384},
// 			},
// 			lang:  1047280,
// 			rules: genFinalRules.approx.first,
// 			expected: Tokens{
// 				{Text: Runes("OrAngE"), Langs: 1047280},
// 				{Text: Runes("OrAnxe"), Langs: 262144},
// 				{Text: Runes("OrAnE"), Langs: 131072},
// 				{Text: Runes("OrAnie"), Langs: 512},
// 				{Text: Runes("OrAnze"), Langs: 32832},
// 				{Text: Runes("OrAnze"), Langs: 331808},
// 				{Text: Runes("PrAngE"), Langs: 16384},
// 			},
// 			ignoreLangs: false,
// 		},
// 		{
// 			name: "orange_stage=3_lang=1",
// 			src: Tokens{
// 				{Text: Runes("OrAngE"), Langs: 1047280},
// 				{Text: Runes("OrAnxe"), Langs: 262144},
// 				{Text: Runes("OrAnE"), Langs: 131072},
// 				{Text: Runes("OrAnie"), Langs: 512},
// 				{Text: Runes("OrAnze"), Langs: 32832},
// 				{Text: Runes("OrAnze"), Langs: 331808},
// 				{Text: Runes("PrAngE"), Langs: 16384},
// 			},
// 			lang:  1047280,
// 			rules: genFinalRules.approx.second[1],
// 			expected: Tokens{
// 				{Text: Runes("orangi"), Langs: 1},
// 				{Text: Runes("oragi"), Langs: 1},
// 				{Text: Runes("orongi"), Langs: 1},
// 				{Text: Runes("orogi"), Langs: 1},
// 				{Text: Runes("orYngi"), Langs: 1},
// 				{Text: Runes("Yrangi"), Langs: 1},
// 				{Text: Runes("Yrongi"), Langs: 1},
// 				{Text: Runes("YrYngi"), Langs: 1},
// 				{Text: Runes("oranxi"), Langs: 1},
// 				{Text: Runes("oronxi"), Langs: 1},
// 				{Text: Runes("orani"), Langs: 1},
// 				{Text: Runes("oroni"), Langs: 1},
// 				{Text: Runes("oranii"), Langs: 1},
// 				{Text: Runes("oronii"), Langs: 1},
// 				{Text: Runes("oranzi"), Langs: 1},
// 				{Text: Runes("oronzi"), Langs: 1},
// 				{Text: Runes("urangi"), Langs: 1},
// 				{Text: Runes("urongi"), Langs: 1},
// 			},
// 			ignoreLangs: true,
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			result := c.rules.Apply(c.src, c.lang, c.ignoreLangs)
// 			require.Equal(t, c.expected, result)
// 		})
// 	}
// }

func Benchmark_Rule_ApplyTo(b *testing.B) {
	r := Rule{
		Pattern: []rune("ge"),
		Phonetic: []Token{
			{
				Text:  Runes("gE"),
				Langs: -1,
			},
			{
				Text:  Runes("xe"),
				Langs: 262144,
			},
			{
				Text:  Runes("hE"),
				Langs: 131072,
			},
			{
				Text:  Runes("je"),
				Langs: 512,
			},
			{
				Text:  Runes("Ze"),
				Langs: 32832,
			},
			{
				Text:  Runes("dZe"),
				Langs: 331808,
			},
		},
	}

	s := Runes("orange")

	for i := 0; i < b.N; i++ {
		r.ApplyTo(s, 4)
	}
}

func Test_Rule_ApplyTo(t *testing.T) {
	r1 := Rule{
		Pattern: []rune("ge"),
		Phonetic: []Token{
			{
				Text:  Runes("gE"),
				Langs: -1,
			},
			{
				Text:  Runes("xe"),
				Langs: 262144,
			},
			{
				Text:  Runes("hE"),
				Langs: 131072,
			},
			{
				Text:  Runes("je"),
				Langs: 512,
			},
			{
				Text:  Runes("Ze"),
				Langs: 32832,
			},
			{
				Text:  Runes("dZe"),
				Langs: 331808,
			},
		},
	}

	cases := []struct {
		name           string
		rule           Rule
		input          string
		position       int
		expectedResult []Token
		expectedOffset int
	}{
		{
			name:     "applied",
			rule:     r1,
			input:    "orange",
			position: 4,

			expectedResult: []Token{
				{Text: Runes("gE"), Langs: LangUnitialized},
				{Text: Runes("xe"), Langs: 262144},
				{Text: Runes("hE"), Langs: 131072},
				{Text: Runes("je"), Langs: 512},
				{Text: Runes("Ze"), Langs: 32832},
				{Text: Runes("dZe"), Langs: 331808},
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
			result, offset := c.rule.ApplyTo(Runes(c.input), c.position)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedOffset, offset)
		})
	}
}
