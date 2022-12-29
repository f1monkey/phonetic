package beidermorse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		name     string
		mode     Mode
		ruleset  Ruleset
		input    string
		expected []string
	}{
		{
			name:     "ru_generic_exact",
			mode:     Generic,
			ruleset:  Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			name:    "ru_generic_approx",
			mode:    Generic,
			ruleset: Approx,
			input:   "апельсин",
			expected: []string{
				"apYlzn",
				"apilzn",
				"opYlzn",
				"opilzn",
				"aplzn",
				"oplzn",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e, err := NewEncoder(WithNameMode(c.mode), WithRuleset(c.ruleset))
			require.NoError(t, err)
			require.Equal(t, c.expected, e.Encode(c.input))
		})
	}
}
