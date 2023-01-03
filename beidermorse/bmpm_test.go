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
		expected int64
	}{
		{
			word: "orange",
			mode: Generic,
			expected: int64(
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
			expected: int64(gencyrillic),
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := detectLang(c.word, c.mode)

			fmt.Printf("%b\n", result)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_phonetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phonetic("test", Generic, Approx, 1, false)
	}
}

func Test_expand(t *testing.T) {
	cases := []struct {
		phonetic string
		expected string
	}{
		{
			phonetic: "(O|P[16384])",
			expected: "O|P[16384]",
		},
		{
			phonetic: "(k[1047288]|ts[16392]|dZ[524288])(O|P[16384])",
			expected: "kO[1047288]|kP[16384]|tsO[16392]|tsP[16384]|dZO[524288]",
		},
	}

	for _, c := range cases {
		t.Run(c.phonetic, func(t *testing.T) {
			result := expand(c.phonetic)
			require.Equal(t, c.expected, result)
		})
	}
}

func Test_normalizeLanguageAttributes(t *testing.T) {
	cases := []struct {
		phonetic string
		strip    bool
		expected string
	}{
		{
			phonetic: "tsYkYtsolo[128][16384]",
			strip:    false,
			expected: "[0]",
		},
		{
			phonetic: "tso",
			strip:    false,
			expected: "tso",
		},
		{
			phonetic: "tsokatsola[1047288][16384]",
			strip:    false,
			expected: "tsokatsola[16384]",
		},
		{
			phonetic: "tsokosu[64]lo",
			strip:    false,
			expected: "tsokosulo[64]",
		},
		{
			phonetic: "kOka[1047288]",
			strip:    true,
			expected: "kOka",
		},
		{
			phonetic: "kOdZa[524288]",
			strip:    true,
			expected: "kOdZa",
		},
	}

	for _, c := range cases {
		t.Run(c.phonetic, func(t *testing.T) {
			result := normalizeLanguageAttributes(c.phonetic, c.strip)
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
