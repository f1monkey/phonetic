package beidermorse

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_detectLang(b *testing.B) {
	detectLang("orange", genLangRules, uint64(genAll))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		detectLang("orange", genLangRules, uint64(genAll))
	}
}

func Test_detectLang(t *testing.T) {
	cases := []struct {
		word      string
		langRules []langRule
		allLangs  uint64
		expected  uint64
	}{
		{
			word:      "orange",
			langRules: genLangRules,
			allLangs:  uint64(genAll),
			expected:  uint64(genany | genczech | gendutch | genenglish | genfrench | gengerman),
		},
		{
			word:      "апельсин",
			langRules: genLangRules,
			allLangs:  uint64(genAll),
			expected:  uint64(gencyrillic),
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := detectLang(c.word, c.langRules, c.allLangs)
			require.Equal(t, c.expected, result)
		})
	}
}
