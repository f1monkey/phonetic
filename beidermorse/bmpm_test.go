package beidermorse

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_detectLang_en(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectLang("orange", Generic)
	}
}

func Benchmark_detectLang_ru(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectLang("апельсин", Generic)
	}
}

func Test_detectLang(t *testing.T) {
	cases := []struct {
		word     string
		mode     Mode
		expected languageID
	}{
		{
			word: "orange",
			mode: Generic,
			expected: languageID(
				gendutch |
					genenglish |
					genfrench |
					gengerman |
					gengreeklatin |
					genhungarian |
					genitalian |
					genlatvian |
					genpolish |
					genportuguese |
					genromanian |
					genrussian |
					genspanish |
					genturkish,
			),
		},
		{
			word:     "апельсин",
			mode:     Generic,
			expected: languageID(gencyrillic),
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := detectLang(c.word, c.mode)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_phonetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeTokens("test", Generic, Approx, 1, false)
	}
}

func Benchmark_applyRules(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyRules([]phonetic{
			{text: "orange", langs: 1047280},
		}, genRules[1], 1047280, false)
	}
}

func Test_applyRules(t *testing.T) {
	cases := []struct {
		name        string
		src         []phonetic
		lang        languageID
		rules       []rule
		ignoreLangs bool
		expected    []phonetic
	}{
		{
			name: "orange_approx_lang=1_stage=1",
			src: []phonetic{
				{text: "orange", langs: 1047280},
			},
			lang:  1047280,
			rules: genRules[1],
			expected: []phonetic{
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
			src: []phonetic{
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
			expected: []phonetic{
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
			src: []phonetic{
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
			expected: []phonetic{
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
			result := applyRules(c.src, c.rules, c.lang, c.ignoreLangs)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_mergeTokenGroups(b *testing.B) {
	r := [][]phonetic{
		{
			{text: "k", langs: 1047288},
			{text: "ts", langs: 16392},
			{text: "dZ", langs: 524288},
		},
		{
			{text: "O", langs: -1},
			{text: "P", langs: 16384},
		},
	}

	for i := 0; i < b.N; i++ {
		mergeTokenGroups(langsAny, r...)
	}
}

func Test_mergeTokenGroups(t *testing.T) {
	cases := []struct {
		src      [][]phonetic
		expected []phonetic
	}{
		{}, // empty
		{
			src: [][]phonetic{
				{
					{text: "O", langs: -1},
					{text: "P", langs: 16384},
				},
			},
			expected: []phonetic{
				{text: "O", langs: -1},
				{text: "P", langs: 16384},
			},
		},
		{
			src: [][]phonetic{
				{
					{text: "O", langs: -1},
					{text: "P", langs: 16384},
				},
			},
			expected: []phonetic{
				{text: "O", langs: -1},
				{text: "P", langs: 16384},
			},
		},
		{
			src: [][]phonetic{
				{
					{text: "k", langs: 1047288},
					{text: "ts", langs: 16392},
					{text: "dZ", langs: 524288},
				},
				{
					{text: "O", langs: -1},
					{text: "P", langs: 16384},
				},
			},
			expected: []phonetic{
				{text: "kO", langs: 1047288},
				{text: "kP", langs: 16384},
				{text: "tsO", langs: 16392},
				{text: "tsP", langs: 16384},
				{text: "dZO", langs: 524288},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := mergeTokenGroups(langsAny, c.src...)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_mergeLangResults(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeLangResults(1047288, 16384)
	}
}

func Test_mergeLangResults(t *testing.T) {
	cases := []struct {
		src      []languageID
		expected languageID
	}{
		{
			src:      []languageID{1, 128, 16384},
			expected: 0,
		},
		{
			src:      []languageID{4, 128, 16384},
			expected: 0,
		},
		{
			src:      []languageID{1047288, 16384},
			expected: 16384,
		},
		{
			src:      []languageID{1047288, 16384},
			expected: 16384,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := mergeLangResults(c.src...)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_prepareInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prepareInput("o'range orange", Generic)
	}
}

func Test_prepareInput(t *testing.T) {
	cases := []struct {
		input    string
		mode     Mode
		expected string
	}{
		{
			input:    "Orange",
			mode:     Generic,
			expected: "orange",
		},
		{
			input:    "de la orange",
			mode:     Generic,
			expected: "dela orange",
		},
		{
			input:    "van der orange",
			mode:     Generic,
			expected: "vander orange",
		},
		{
			input:    "van den orange",
			mode:     Generic,
			expected: "vanden orange",
		},
		{
			input:    "o'range orange orange",
			mode:     Generic,
			expected: "orange orangeorange",
		},
		{
			input:    "o'range orange orange",
			mode:     Sephardic,
			expected: "o rangeorangeorange",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s_%s", c.mode, c.input), func(t *testing.T) {
			result := prepareInput(c.input, c.mode)
			require.Equal(t, c.expected, result)
		})
	}
}
