package bmpm

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

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
			expected:  249,
		},
		{
			word:      "апельсин",
			langRules: genLangRules,
			allLangs:  uint64(genAll),
			expected:  4,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := detectLang(c.word, c.langRules, c.allLangs)
			require.Equal(t, c.expected, result)
		})
	}
}
