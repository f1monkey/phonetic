package beidermorse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Benchmark_rules_apply(b *testing.B) {
	t := tokens{
		{text: runes("orange"), langs: 1047280},
	}
	for i := 0; i < b.N; i++ {
		genRules[1].apply(t, 1047280, false)
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
			name: "test_approx_lang=1_stage=1",
			src: tokens{
				{text: runes("test"), langs: 1047288},
			},
			lang:  1047288,
			rules: genRules[1],
			expected: tokens{
				{text: runes("tESt"), langs: 128},
				{text: runes("tEst"), langs: 1047288},
			},
			ignoreLangs: false,
		},
		{
			name: "test_approx_lang=1_stage=2",
			src: tokens{
				{text: runes("tESt"), langs: 128},
				{text: runes("tEst"), langs: 1047288},
			},
			lang:  1047288,
			rules: genFinalRules.approx.first,
			expected: tokens{
				{text: runes("tEst"), langs: 128},
				{text: runes("tEst"), langs: 1047288},
			},
			ignoreLangs: false,
		},
		{
			name: "test_approx_lang=1_stage=3",
			src: tokens{
				{text: runes("tEst"), langs: 128},
				{text: runes("tEst"), langs: 1047288},
			},
			lang:  1047288,
			rules: genFinalRules.approx.second[1],
			expected: tokens{
				{text: runes("tist"), langs: 1},
				{text: runes("tYst"), langs: 1},
				{text: runes("tis"), langs: 1},
				{text: runes("tit"), langs: 1},
				{text: runes("ti"), langs: 1},
			},
			ignoreLangs: true,
		},
		{
			name: "orange_approx_lang=1_stage=1",
			src: tokens{
				{text: runes("orange"), langs: 1047280},
			},
			lang:  1047280,
			rules: genRules[1],
			expected: tokens{
				{text: runes("OrAngE"), langs: 1047280},
				{text: runes("OrAnxe"), langs: 262144},
				{text: runes("OrAnhE"), langs: 131072},
				{text: runes("OrAnje"), langs: 512},
				{text: runes("OrAnZe"), langs: 32832},
				{text: runes("OrAndZe"), langs: 331808},
				{text: runes("PrAngE"), langs: 16384},
			},
			ignoreLangs: false,
		},
		{
			name: "orange_approx_lang=1_stage=2",
			src: tokens{
				{text: runes("OrAngE"), langs: 1047280},
				{text: runes("OrAnxe"), langs: 262144},
				{text: runes("OrAnhE"), langs: 131072},
				{text: runes("OrAnje"), langs: 512},
				{text: runes("OrAnZe"), langs: 32832},
				{text: runes("OrAndZe"), langs: 331808},
				{text: runes("PrAngE"), langs: 16384},
			},
			lang:  1047280,
			rules: genFinalRules.approx.first,
			expected: tokens{
				{text: runes("OrAngE"), langs: 1047280},
				{text: runes("OrAnxe"), langs: 262144},
				{text: runes("OrAnE"), langs: 131072},
				{text: runes("OrAnie"), langs: 512},
				{text: runes("OrAnze"), langs: 32832},
				{text: runes("OrAnze"), langs: 331808},
				{text: runes("PrAngE"), langs: 16384},
			},
			ignoreLangs: false,
		},
		{
			name: "orange_stage=3_lang=1",
			src: tokens{
				{text: runes("OrAngE"), langs: 1047280},
				{text: runes("OrAnxe"), langs: 262144},
				{text: runes("OrAnE"), langs: 131072},
				{text: runes("OrAnie"), langs: 512},
				{text: runes("OrAnze"), langs: 32832},
				{text: runes("OrAnze"), langs: 331808},
				{text: runes("PrAngE"), langs: 16384},
			},
			lang:  1047280,
			rules: genFinalRules.approx.second[1],
			expected: tokens{
				{text: runes("orangi"), langs: 1},
				{text: runes("oragi"), langs: 1},
				{text: runes("orongi"), langs: 1},
				{text: runes("orogi"), langs: 1},
				{text: runes("orYngi"), langs: 1},
				{text: runes("Yrangi"), langs: 1},
				{text: runes("Yrongi"), langs: 1},
				{text: runes("YrYngi"), langs: 1},
				{text: runes("oranxi"), langs: 1},
				{text: runes("oronxi"), langs: 1},
				{text: runes("orani"), langs: 1},
				{text: runes("oroni"), langs: 1},
				{text: runes("oranii"), langs: 1},
				{text: runes("oronii"), langs: 1},
				{text: runes("oranzi"), langs: 1},
				{text: runes("oronzi"), langs: 1},
				{text: runes("urangi"), langs: 1},
				{text: runes("urongi"), langs: 1},
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
		pattern: []rune("ge"),
		phoneticRules: []token{
			{
				text:  runes("gE"),
				langs: -1,
			},
			{
				text:  runes("xe"),
				langs: 262144,
			},
			{
				text:  runes("hE"),
				langs: 131072,
			},
			{
				text:  runes("je"),
				langs: 512,
			},
			{
				text:  runes("Ze"),
				langs: 32832,
			},
			{
				text:  runes("dZe"),
				langs: 331808,
			},
		},
	}

	s := runes("orange")

	for i := 0; i < b.N; i++ {
		r.applyTo(s, 4)
	}
}

func Test_rule_applyTo(t *testing.T) {
	r1 := rule{
		pattern: []rune("ge"),
		phoneticRules: []token{
			{
				text:  runes("gE"),
				langs: -1,
			},
			{
				text:  runes("xe"),
				langs: 262144,
			},
			{
				text:  runes("hE"),
				langs: 131072,
			},
			{
				text:  runes("je"),
				langs: 512,
			},
			{
				text:  runes("Ze"),
				langs: 32832,
			},
			{
				text:  runes("dZe"),
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
				{text: runes("gE"), langs: langsUnitialized},
				{text: runes("xe"), langs: 262144},
				{text: runes("hE"), langs: 131072},
				{text: runes("je"), langs: 512},
				{text: runes("Ze"), langs: 32832},
				{text: runes("dZe"), langs: 331808},
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
			result, offset := c.rule.applyTo(runes(c.input), c.position)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedOffset, offset)
		})
	}
}
