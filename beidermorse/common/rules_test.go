package common

import (
	"regexp"
	"testing"

	"github.com/f1monkey/phonetic/internal/exrunes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Benchmark_Rules_Apply(b *testing.B) {
	t := &Tokens{
		Buf: func() *exrunes.Buffer {
			buf := exrunes.NewBuffer(100)
			buf.Add([]rune("orange"))
			return buf
		}(),
		Items: []Token{{ID: 0, Langs: 1047280}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genRules[1].Apply(t, 1047280, false)
	}
}

func Test_Rules_Apply(t *testing.T) {
	cases := []struct {
		name        string
		src         *Tokens
		lang        Lang
		rules       Rules
		ignoreLangs bool
		expected    []expectedToken
	}{
		{
			name: "test_approx_lang=1_stage=1",
			src: &Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("test"))

					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: 1047288},
				},
			},
			lang:  1047288,
			rules: genRules[1],
			expected: []expectedToken{
				{text: []rune("tESt"), langs: 128},
				{text: []rune("tEst"), langs: 1047288},
			},
			ignoreLangs: false,
		},
		{
			name: "test_approx_lang=1_stage=2",
			src: &Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("tESt"))
					buf.Add([]rune("tEst"))

					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: 128},
					{ID: 1, Langs: 1047288},
				},
			},

			lang:  1047288,
			rules: genFinalRules.Approx.First,
			expected: []expectedToken{
				{text: []rune("tEst"), langs: 128},
				{text: []rune("tEst"), langs: 1047288},
			},
			ignoreLangs: false,
		},
		{
			name: "test_approx_lang=1_stage=3",
			src: &Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("tEst"))
					buf.Add([]rune("tEst"))
					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: 128},
					{ID: 1, Langs: 1047288},
				},
			},
			lang:  1047288,
			rules: genFinalRules.Approx.Second[1],
			expected: []expectedToken{
				{text: []rune("tist"), langs: 1},
				{text: []rune("tYst"), langs: 1},
				{text: []rune("tis"), langs: 1},
				{text: []rune("tit"), langs: 1},
				{text: []rune("ti"), langs: 1},
			},
			ignoreLangs: true,
		},
		{
			name: "orange_approx_lang=1_stage=1",
			src: &Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("orange"))
					return buf
				}(),
				Items: []Token{{ID: 0, Langs: 1047280}},
			},
			lang:  1047280,
			rules: genRules[1],
			expected: []expectedToken{
				{text: []rune("OrAngE"), langs: 1047280},
				{text: []rune("OrAnxe"), langs: 262144},
				{text: []rune("OrAnhE"), langs: 131072},
				{text: []rune("OrAnje"), langs: 512},
				{text: []rune("OrAnZe"), langs: 32832},
				{text: []rune("OrAndZe"), langs: 331808},
				{text: []rune("PrAngE"), langs: 16384},
			},
			ignoreLangs: false,
		},
		{
			name: "orange_approx_lang=1_stage=2",
			src: &Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("OrAngE"))
					buf.Add([]rune("OrAnxe"))
					buf.Add([]rune("OrAnhE"))
					buf.Add([]rune("OrAnje"))
					buf.Add([]rune("OrAnZe"))
					buf.Add([]rune("OrAndZe"))
					buf.Add([]rune("PrAngE"))
					return buf
				}(),
				Items: []Token{

					{ID: 0, Langs: 1047280},
					{ID: 1, Langs: 262144},
					{ID: 2, Langs: 131072},
					{ID: 3, Langs: 512},
					{ID: 4, Langs: 32832},
					{ID: 5, Langs: 331808},
					{ID: 6, Langs: 16384},
				},
			},
			lang:  1047280,
			rules: genFinalRules.Approx.First,
			expected: []expectedToken{
				{text: []rune("OrAngE"), langs: 1047280},
				{text: []rune("OrAnxe"), langs: 262144},
				{text: []rune("OrAnE"), langs: 131072},
				{text: []rune("OrAnie"), langs: 512},
				{text: []rune("OrAnze"), langs: 32832},
				{text: []rune("OrAnze"), langs: 331808},
				{text: []rune("PrAngE"), langs: 16384},
			},
			ignoreLangs: false,
		},
		{
			name: "orange_stage=3_lang=1",
			src: &Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("OrAngE"))
					buf.Add([]rune("OrAnxe"))
					buf.Add([]rune("OrAnE"))
					buf.Add([]rune("OrAnie"))
					buf.Add([]rune("OrAnze"))
					buf.Add([]rune("OrAnze"))
					buf.Add([]rune("PrAngE"))
					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: 1047280},
					{ID: 1, Langs: 262144},
					{ID: 2, Langs: 131072},
					{ID: 3, Langs: 512},
					{ID: 4, Langs: 32832},
					{ID: 5, Langs: 331808},
					{ID: 6, Langs: 16384},
				},
			},
			lang:  1047280,
			rules: genFinalRules.Approx.Second[1],
			expected: []expectedToken{
				{text: []rune("orangi"), langs: 1},
				{text: []rune("oragi"), langs: 1},
				{text: []rune("orongi"), langs: 1},
				{text: []rune("orogi"), langs: 1},
				{text: []rune("orYngi"), langs: 1},
				{text: []rune("Yrangi"), langs: 1},
				{text: []rune("Yrongi"), langs: 1},
				{text: []rune("YrYngi"), langs: 1},
				{text: []rune("oranxi"), langs: 1},
				{text: []rune("oronxi"), langs: 1},
				{text: []rune("orani"), langs: 1},
				{text: []rune("oroni"), langs: 1},
				{text: []rune("oranii"), langs: 1},
				{text: []rune("oronii"), langs: 1},
				{text: []rune("oranzi"), langs: 1},
				{text: []rune("oronzi"), langs: 1},
				{text: []rune("urangi"), langs: 1},
				{text: []rune("urongi"), langs: 1},
			},
			ignoreLangs: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rr := c.rules.Apply(c.src, c.lang, c.ignoreLangs)

			result := make([]expectedToken, len(rr.Items))
			for i, t := range rr.Items {
				result[i] = expectedToken{text: c.src.Buf.Get(t.ID), langs: t.Langs}
			}

			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_Rule_ApplyTo(b *testing.B) {
	r := Rule{
		Pattern: []rune("ge"),
		Phonetic: []RuleToken{
			{
				Text:  []rune("gE"),
				Langs: -1,
			},
			{
				Text:  []rune("xe"),
				Langs: 262144,
			},
			{
				Text:  []rune("hE"),
				Langs: 131072,
			},
			{
				Text:  []rune("je"),
				Langs: 512,
			},
			{
				Text:  []rune("Ze"),
				Langs: 32832,
			},
			{
				Text:  []rune("dZe"),
				Langs: 331808,
			},
		},
	}

	s := []rune("orange")

	for i := 0; i < b.N; i++ {
		r.ApplyTo(s, 4)
	}
}

func Test_Rule_ApplyTo(t *testing.T) {
	r1 := Rule{
		Pattern: []rune("ge"),
		Phonetic: []RuleToken{
			{
				Text:  []rune("gE"),
				Langs: -1,
			},
			{
				Text:  []rune("xe"),
				Langs: 262144,
			},
			{
				Text:  []rune("hE"),
				Langs: 131072,
			},
			{
				Text:  []rune("je"),
				Langs: 512,
			},
			{
				Text:  []rune("Ze"),
				Langs: 32832,
			},
			{
				Text:  []rune("dZe"),
				Langs: 331808,
			},
		},
	}

	cases := []struct {
		name           string
		rule           Rule
		input          string
		position       int
		expectedResult []RuleToken
		expectedOffset int
	}{
		{
			name:     "applied",
			rule:     r1,
			input:    "orange",
			position: 4,

			expectedResult: []RuleToken{
				{Text: []rune("gE"), Langs: LangUnitialized},
				{Text: []rune("xe"), Langs: 262144},
				{Text: []rune("hE"), Langs: 131072},
				{Text: []rune("je"), Langs: 512},
				{Text: []rune("Ze"), Langs: 32832},
				{Text: []rune("dZe"), Langs: 331808},
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
			result, offset := c.rule.ApplyTo([]rune(c.input), c.position)
			assert.Equal(t, c.expectedResult, result)
			assert.Equal(t, c.expectedOffset, offset)
		})
	}
}

// partial copy of generic rules to use in tests
var genRules = map[Lang]Rules{
	1: {
		{
			Pattern: []rune("yna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("in"),
					Langs: 131072,
				},
				{
					Text:  []rune("ina"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ina"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("in"),
					Langs: 131072,
				},
				{
					Text:  []rune("ina"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("liova"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: 131072,
				},
				{
					Text:  []rune("lef"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lova"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: 131072,
				},
				{
					Text:  []rune("lef"),
					Langs: 131072,
				},
				{
					Text:  []rune("l"),
					Langs: 8,
				},
				{
					Text:  []rune("el"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("kova"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("kova"),
					Langs: -1,
				},
				{
					Text:  []rune("kof"),
					Langs: 131072,
				},
				{
					Text:  []rune("k"),
					Langs: 8,
				},
				{
					Text:  []rune("ek"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ova"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  []rune("of"),
					Langs: 131072,
				},
				{
					Text:  nil,
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ová"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("eva"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("eva"),
					Langs: -1,
				},
				{
					Text:  []rune("ef"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("aia"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("aja"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("aya"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lowa"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: 16384,
				},
				{
					Text:  []rune("l"),
					Langs: 16384,
				},
				{
					Text:  []rune("el"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("kowa"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("kova"),
					Langs: -1,
				},
				{
					Text:  []rune("kof"),
					Langs: 16384,
				},
				{
					Text:  []rune("k"),
					Langs: 16384,
				},
				{
					Text:  []rune("ek"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("owa"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  []rune("of"),
					Langs: 16384,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lowna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("lovna"),
					Langs: -1,
				},
				{
					Text:  []rune("levna"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 16384,
				},
				{
					Text:  []rune("el"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("kowna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("kovna"),
					Langs: -1,
				},
				{
					Text:  []rune("k"),
					Langs: 16384,
				},
				{
					Text:  []rune("ek"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("owna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ovna"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("lówna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("el"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kówna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ówna"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("a"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 16392,
				},
			},
		},
		{
			Pattern: []rune("pf"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("pf"),
					Langs: -1,
				},
				{
					Text:  []rune("p"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("que"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: 64,
				},
				{
					Text:  []rune("ke"),
					Langs: -1,
				},
				{
					Text:  []rune("kve"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bfpv]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: 32832,
				},
			},
		},
		{
			Pattern: []rune("ly"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("lj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("li"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("lj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lio"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lyo"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lt"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Suffix:           []rune("u"),
			},
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("lt"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("v"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: 128,
				},
				{
					Text:  []rune("b"),
					Langs: 262144,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ez"),
					Langs: 32768,
				},
				{
					Text:  []rune("eS"),
					Langs: 32768,
				},
				{
					Text:  []rune("eks"),
					Langs: -1,
				},
				{
					Text:  []rune("egz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[cs]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: 32768,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Suffix:           []rune("u"),
			},
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("tsk"),
					Langs: 16392,
				},
			},
		},
		{
			Pattern: []rune("cz"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
				{
					Text:  []rune("tsz"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("rh"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("dh"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("bh"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ph"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kh"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("x"),
					Langs: 131104,
				},
				{
					Text:  []rune("kh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lh"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("lh"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("nh"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("nh"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("ssch"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("xS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsch"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
				{
					Text:  []rune("sk"),
					Langs: 69632,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("sk"),
					Langs: 69632,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("ssh"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[äöü]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("sh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: 131104,
				},
				{
					Text:  []rune("sh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zh"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Z"),
					Langs: 131104,
				},
				{
					Text:  []rune("zh"),
					Langs: -1,
				},
				{
					Text:  []rune("tsh"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("chs"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ks"),
					Langs: 128,
				},
				{
					Text:  []rune("xs"),
					Langs: -1,
				},
				{
					Text:  []rune("tSs"),
					Langs: 131104,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 393248,
				},
				{
					Text:  []rune("k"),
					Langs: 69632,
				},
				{
					Text:  []rune("S"),
					Langs: 32832,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 393248,
				},
				{
					Text:  []rune("S"),
					Langs: 32832,
				},
			},
		},
		{
			Pattern: []rune("th"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("t"),
					Langs: 672,
				},
				{
					Text:  []rune("th"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gh"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("g"),
					Langs: 70144,
				},
				{
					Text:  []rune("gh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ouh"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aioe]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("v"),
					Langs: 64,
				},
				{
					Text:  []rune("uh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uh"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aioe]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("uh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: 66048,
				},
				{
					Text:  []rune("H"),
					Langs: 381024,
				},
			},
		},
		{
			Pattern: []rune("cia"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSa"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cią"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSom"),
					Langs: -1,
				},
				{
					Text:  []rune("tsom"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cią"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSon"),
					Langs: 16384,
				},
				{
					Text:  []rune("tson"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cię"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSem"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cię"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSen"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cie"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: 16384,
				},
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cio"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSo"),
					Langs: 16384,
				},
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ciu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tSu"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sci"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Si"),
					Langs: 4096,
				},
				{
					Text:  []rune("stsi"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZi"),
					Langs: 524288,
				},
				{
					Text:  []rune("tSi"),
					Langs: 81920,
				},
				{
					Text:  []rune("tS"),
					Langs: 65536,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: 4096,
				},
				{
					Text:  []rune("sts"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZ"),
					Langs: 524288,
				},
				{
					Text:  []rune("tS"),
					Langs: 81920,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("tsi"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZi"),
					Langs: 524288,
				},
				{
					Text:  []rune("tSi"),
					Langs: 81920,
				},
				{
					Text:  []rune("tS"),
					Langs: 65536,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cy"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZ"),
					Langs: 524288,
				},
				{
					Text:  []rune("tS"),
					Langs: 81920,
				},
				{
					Text:  []rune("k"),
					Langs: 512,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sç"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("stS"),
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ssz"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: 2048,
				},
				{
					Text:  []rune("sts"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("ssp"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: 128,
				},
				{
					Text:  []rune("sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sp"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: 128,
				},
				{
					Text:  []rune("sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sst"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("St"),
					Langs: 128,
				},
				{
					Text:  []rune("st"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("st"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("St"),
					Langs: 128,
				},
				{
					Text:  []rune("st"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sj"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sj"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sj"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("sj"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 16,
				},
				{
					Text:  []rune("sx"),
					Langs: 262144,
				},
				{
					Text:  []rune("sZ"),
					Langs: 589824,
				},
			},
		},
		{
			Pattern: []rune("sia"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Sa"),
					Langs: 16384,
				},
				{
					Text:  []rune("sa"),
					Langs: 16384,
				},
				{
					Text:  []rune("sja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sią"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Som"),
					Langs: 16384,
				},
				{
					Text:  []rune("som"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sią"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Son"),
					Langs: 16384,
				},
				{
					Text:  []rune("son"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("się"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Sem"),
					Langs: 16384,
				},
				{
					Text:  []rune("sem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("się"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Sen"),
					Langs: 16384,
				},
				{
					Text:  []rune("sen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
				{
					Text:  []rune("sje"),
					Langs: -1,
				},
				{
					Text:  []rune("Se"),
					Langs: 16384,
				},
				{
					Text:  []rune("zi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("So"),
					Langs: 16384,
				},
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("siu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Su"),
					Langs: 16384,
				},
				{
					Text:  []rune("sju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("si"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[äöëaáuiíoóeéêy]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Si"),
					Langs: 16384,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: 37056,
				},
			},
		},
		{
			Pattern: []rune("si"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Si"),
					Langs: 16384,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aáuíoóeéêy]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: 37056,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeouäöë]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: 32768,
				},
				{
					Text:  nil,
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("gue"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: 64,
				},
				{
					Text:  []rune("gve"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("g"),
					Langs: 64,
				},
				{
					Text:  []rune("gv"),
					Langs: 294912,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ao]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("guy"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gli"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("glI"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 4096,
				},
			},
		},
		{
			Pattern: []rune("gni"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("gnI"),
					Langs: -1,
				},
				{
					Text:  []rune("ni"),
					Langs: 4160,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("n"),
					Langs: 4160,
				},
				{
					Text:  []rune("nj"),
					Langs: 4160,
				},
				{
					Text:  []rune("gn"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggie"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("je"),
					Langs: 512,
				},
				{
					Text:  []rune("dZe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("j"),
					Langs: 512,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 4096,
				},
				{
					Text:  []rune("j"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("gge"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gE"),
					Langs: -1,
				},
				{
					Text:  []rune("xe"),
					Langs: 262144,
				},
				{
					Text:  []rune("gZe"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZe"),
					Langs: 331808,
				},
				{
					Text:  []rune("je"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("xi"),
					Langs: 262144,
				},
				{
					Text:  []rune("gZi"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZi"),
					Langs: 331808,
				},
				{
					Text:  []rune("i"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 4096,
				},
				{
					Text:  []rune("j"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("gie"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
				{
					Text:  []rune("gi"),
					Langs: 128,
				},
				{
					Text:  []rune("ji"),
					Langs: 64,
				},
				{
					Text:  []rune("dZe"),
					Langs: 4096,
				},
			},
		},
		{
			Pattern: []rune("gie"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
				{
					Text:  []rune("gi"),
					Langs: 128,
				},
				{
					Text:  []rune("dZe"),
					Langs: 4096,
				},
				{
					Text:  []rune("je"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: 512,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ge"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gE"),
					Langs: -1,
				},
				{
					Text:  []rune("xe"),
					Langs: 262144,
				},
				{
					Text:  []rune("Ze"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZe"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("xi"),
					Langs: 262144,
				},
				{
					Text:  []rune("Zi"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZi"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("ge"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("gE"),
					Langs: -1,
				},
				{
					Text:  []rune("xe"),
					Langs: 262144,
				},
				{
					Text:  []rune("hE"),
					Langs: 131072,
				},
				{
					Text:  []rune("je"),
					Langs: 512,
				},
				{
					Text:  []rune("Ze"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZe"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("xi"),
					Langs: 262144,
				},
				{
					Text:  []rune("hI"),
					Langs: 131072,
				},
				{
					Text:  []rune("i"),
					Langs: 512,
				},
				{
					Text:  []rune("Zi"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZi"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("d"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("g"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aouyei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aouei]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("h"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: 16,
				},
				{
					Text:  []rune("ix"),
					Langs: 262144,
				},
				{
					Text:  []rune("iZ"),
					Langs: 622656,
				},
			},
		},
		{
			Pattern: []rune("j"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aoeiuy]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 32,
				},
				{
					Text:  []rune("x"),
					Langs: 262144,
				},
				{
					Text:  []rune("Z"),
					Langs: 622656,
				},
			},
		},
		{
			Pattern: []rune("rz"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Suffix:           []rune("t"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: 16384,
				},
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rz"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("rz"),
					Langs: -1,
				},
				{
					Text:  []rune("rts"),
					Langs: 128,
				},
				{
					Text:  []rune("Z"),
					Langs: 16384,
				},
				{
					Text:  []rune("r"),
					Langs: 16384,
				},
				{
					Text:  []rune("rZ"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 160,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 131232,
				},
				{
					Text:  []rune("tS"),
					Langs: 160,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 131232,
				},
				{
					Text:  []rune("tz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Za"),
					Langs: 16384,
				},
				{
					Text:  []rune("za"),
					Langs: 16384,
				},
				{
					Text:  []rune("zja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Za"),
					Langs: 16384,
				},
				{
					Text:  []rune("zja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zią"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zom"),
					Langs: 16384,
				},
				{
					Text:  []rune("zom"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zią"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zon"),
					Langs: 16384,
				},
				{
					Text:  []rune("zon"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zię"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zem"),
					Langs: 16384,
				},
				{
					Text:  []rune("zem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zię"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zen"),
					Langs: 16384,
				},
				{
					Text:  []rune("zen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: 16384,
				},
				{
					Text:  []rune("ze"),
					Langs: 16384,
				},
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
				{
					Text:  []rune("Ze"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zo"),
					Langs: 16384,
				},
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ziu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zu"),
					Langs: 16384,
				},
				{
					Text:  []rune("zju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zi"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: 16384,
				},
				{
					Text:  []rune("zi"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
				{
					Text:  []rune("dzi"),
					Langs: 4096,
				},
				{
					Text:  []rune("tsi"),
					Langs: 4096,
				},
				{
					Text:  []rune("si"),
					Langs: 262144,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 128,
				},
				{
					Text:  []rune("ts"),
					Langs: 4096,
				},
				{
					Text:  []rune("S"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bdgv]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("dz"),
					Langs: 4096,
				},
				{
					Text:  []rune("Z"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ptckf]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 4096,
				},
				{
					Text:  []rune("S"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oue"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("oue"),
					Langs: -1,
				},
				{
					Text:  []rune("ve"),
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("eau"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ae"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Y"),
					Langs: 128,
				},
				{
					Text:  []rune("aje"),
					Langs: 131072,
				},
				{
					Text:  []rune("ae"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ão"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ãe"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ãi"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ea"),
					Langs: -1,
				},
				{
					Text:  []rune("ja"),
					Langs: 65536,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: 32,
				},
				{
					Text:  []rune("aje"),
					Langs: 131072,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("eu"),
					Langs: -1,
				},
				{
					Text:  []rune("Yj"),
					Langs: 128,
				},
				{
					Text:  []rune("ej"),
					Langs: 128,
				},
				{
					Text:  []rune("oj"),
					Langs: 128,
				},
				{
					Text:  []rune("Y"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ia"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: 128,
				},
				{
					Text:  []rune("e"),
					Langs: 16384,
				},
				{
					Text:  []rune("ije"),
					Langs: 131072,
				},
				{
					Text:  []rune("Q"),
					Langs: 16,
				},
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("io"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("iu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iy"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Y"),
					Langs: 128,
				},
				{
					Text:  []rune("oje"),
					Langs: 131072,
				},
				{
					Text:  []rune("u"),
					Langs: 16,
				},
				{
					Text:  []rune("oe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oi"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: 32,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: 576,
				},
				{
					Text:  []rune("au"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("où"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oy"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õe"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ua"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("va"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ue"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Q"),
					Langs: 128,
				},
				{
					Text:  []rune("uje"),
					Langs: 131072,
				},
				{
					Text:  []rune("ve"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ui"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("uj"),
					Langs: -1,
				},
				{
					Text:  []rune("vi"),
					Langs: -1,
				},
				{
					Text:  []rune("Y"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("uu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
				{
					Text:  []rune("Q"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("vo"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uy"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("uj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ya"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ye"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
				{
					Text:  []rune("ije"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yo"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("yu"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[áóéê]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[áóéê]$"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			LeftContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("EE"),
					Langs: 96,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("om"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ă"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: 65536,
				},
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ā"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("č"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ć"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("tS"),
					Langs: 16384,
				},
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ď"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("en"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ě"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ē"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ģ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ğ"),
			Phonetic: []RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("î"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ī"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ı"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 524288,
				},
				{
					Text:  nil,
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ķ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("t"),
					Langs: 8192,
				},
				{
					Text:  []rune("tj"),
					Langs: 8192,
				},
			},
		},
		{
			Pattern: []rune("ļ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ł"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("ñ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 262144,
				},
			},
		},
		{
			Pattern: []rune("ņ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 8192,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: 16384,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ô"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: 32768,
				},
				{
					Text:  []rune("Y"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("ò"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ř"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
				{
					Text:  []rune("rZ"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ś"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: 16384,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ţ"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ť"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  []rune("tj"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: 294912,
				},
			},
		},
		{
			Pattern: []rune("ū"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ů"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ù"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ý"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ż"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Z"),
					Langs: 16384,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ž"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ß"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
			Phonetic: []RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			RightContext: &Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			Phonetic: []RuleToken{
				{
					Text:  []rune("O"),
					Langs: -1,
				},
				{
					Text:  []rune("P"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("A"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("B"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZ"),
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: 65536,
				},
				{
					Text:  []rune("H"),
					Langs: 299072,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: 262144,
				},
				{
					Text:  []rune("Z"),
					Langs: 622656,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("O"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("V"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("w"),
					Langs: 48,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 294912,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 128,
				},
				{
					Text:  []rune("dz"),
					Langs: 4096,
				},
				{
					Text:  []rune("ts"),
					Langs: 4096,
				},
				{
					Text:  []rune("s"),
					Langs: 262144,
				},
			},
		},
	},
}

// partial copy of generic final rules to use in tests
var genFinalRules = FinalRules{
	Approx: FinalRule{
		First: Rules{
			{
				Pattern: []rune("h"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[fktSs]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("p"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("b"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pktSs]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("f"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("v"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pftSs]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("k"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("g"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("g"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("t"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("d"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("d"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("dZ"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("tS"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jnm"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("jm"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				LeftContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jI"),
				LeftContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("I"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[aA]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Suffix:           []rune("A"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("A"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("A"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("b"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("d"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("f"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("g"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("j"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("j"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("k"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("l"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("l"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("m"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("m"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("n"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("p"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("r"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("r"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("t"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("v"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("z"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("vanden"),
				LeftContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("vanden"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("vander"),
				LeftContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("vander"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("van"),
				LeftContext: &Matcher{
					MatchEmptyString: true,
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[bp]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("vam"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 16,
					},
				},
			},
			{
				Pattern: []rune("van"),
				LeftContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("van"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 16,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[bp]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("h"),
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("H"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("x"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sen"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[rmnl]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("zn"),
						Langs: -1,
					},
					{
						Text:  []rune("zon"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sen"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("sn"),
						Langs: -1,
					},
					{
						Text:  []rune("son"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sEn"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[rmnl]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("zn"),
						Langs: -1,
					},
					{
						Text:  []rune("zon"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sEn"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("sn"),
						Langs: -1,
					},
					{
						Text:  []rune("son"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("E"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("I"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Q"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Y"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("e"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("E"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("E"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("I"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("I"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Q"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("Q"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Y"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("Y"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEs"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("lEs"),
						Langs: -1,
					},
					{
						Text:  []rune("lz"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lE"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lE"),
						Langs: -1,
					},
					{
						Text:  []rune("l"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aue"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oue"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AvE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("AvE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ave"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("Ave"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("avE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("avE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ave"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ave"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OvE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("OvE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ove"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("Ove"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ovE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ovE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ove"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ove"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ea"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ea"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("EA"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ea"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("Ea"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("eA"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uje"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("iji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ije"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Iji"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjI"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ije"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjE"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ajo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ojo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ejo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uja"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ujo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ija"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ija"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjA"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ijo"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjO"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("j"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lYndEr"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lander"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lAndEr"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lAnder"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("landEr"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lender"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEndEr"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lendEr"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEnder"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("burk"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bUrk"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("burg"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bUrg"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Burk"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("BUrk"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Burg"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("BUrg"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[rmnl]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				LeftContext: &Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[rmnl]$"),
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dS"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				Phonetic: []RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[Lang]Rules{
			1: {
				{
					Pattern: []rune("mb"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("mb"),
							Langs: -1,
						},
						{
							Text:  []rune("b"),
							Langs: 512,
						},
					},
				},
				{
					Pattern: []rune("mp"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("mp"),
							Langs: -1,
						},
						{
							Text:  []rune("b"),
							Langs: 512,
						},
					},
				},
				{
					Pattern: []rune("ng"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("ng"),
							Langs: -1,
						},
						{
							Text:  []rune("g"),
							Langs: 512,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[fktSs]"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  []rune("f"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Prefix:           []rune("p"),
					},
					Phonetic: []RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  []rune("f"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[pktSs]"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("f"),
							Langs: -1,
						},
						{
							Text:  []rune("p"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Prefix:           []rune("f"),
					},
					Phonetic: []RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("f"),
							Langs: -1,
						},
						{
							Text:  []rune("p"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("B"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("b"),
							Langs: -1,
						},
						{
							Text:  []rune("v"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("V"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("v"),
							Langs: -1,
						},
						{
							Text:  []rune("b"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("t"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("t"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("g"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Suffix:           []rune("n"),
					},
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("g"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("k"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Suffix:           []rune("n"),
					},
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("k"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("p"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("r"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[Ee]$"),
					},
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("r"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("s"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("s"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("t"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[aeiouAEIOU]$"),
					},
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("t"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("s"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[aeiouAEIOU]$"),
					},
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("s"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("Q"),
							Langs: 128,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("D"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("Q"),
							Langs: 128,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("lEE"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("li"),
							Langs: -1,
						},
						{
							Text:  []rune("il"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("rEE"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("ri"),
							Langs: -1,
						},
						{
							Text:  []rune("ir"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("lE"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("li"),
							Langs: -1,
						},
						{
							Text:  []rune("il"),
							Langs: 32,
						},
						{
							Text:  []rune("lY"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("rE"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("ri"),
							Langs: -1,
						},
						{
							Text:  []rune("ir"),
							Langs: 32,
						},
						{
							Text:  []rune("rY"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("EE"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ea"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ou"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eu"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("e"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ai"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ai"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oi"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Oi"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ui"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ei"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ei"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iA"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("ia"),
							Langs: -1,
						},
						{
							Text:  []rune("io"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iA"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("ia"),
							Langs: -1,
						},
						{
							Text:  []rune("io"),
							Langs: -1,
						},
						{
							Text:  []rune("iY"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
						{
							Text:  []rune("D"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
						{
							Text:  nil,
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
						{
							Text:  nil,
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[DoiuQY]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sUts"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("suts"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Uts"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					LeftContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("a"),
					Phonetic: []RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
			},
		},
	},
}
