package beidermorse

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_detectLang(b *testing.B) {
	detectLang("orange", Generic)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		detectLang("orange", Generic)
	}
}

func Test_detectLang(t *testing.T) {
	cases := []struct {
		word     string
		mode     NameMode
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

// func Test_Phonetic(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		mode     NameMode
// 		input    string
// 		expected string
// 	}{
// 		{
// 			name:     "ru_generic",
// 			mode:     Generic,
// 			input:    "апельсин",
// 			expected: "",
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			result := Phonetic(c.input, c.mode)
// 			require.Equal(t, c.expected, result)
// 		})
// 	}
// }
