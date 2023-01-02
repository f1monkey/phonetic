package beidermorse

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_detectLang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectLang("orange", Generic)
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
	}

	for _, c := range cases {
		t.Run(c.phonetic, func(t *testing.T) {
			result := expand(c.phonetic)
			require.Equal(t, c.expected, result)
		})
	}
}
