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
