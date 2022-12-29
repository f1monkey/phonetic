package beidermorse

import (
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
		expected uint64
	}{
		{
			word:     "orange",
			mode:     Generic,
			expected: uint64(genany | genczech | gendutch | genenglish | genfrench | gengerman),
		},
		{
			word:     "апельсин",
			mode:     Generic,
			expected: uint64(gencyrillic),
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
		phonetic("test", Generic, Approx, 1, false)
	}
}
