package beidermorse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode_Ru_Approx(b *testing.B) {
	e := MustNewEncoder(WithNameMode(Generic), WithRuleset(Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

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
		{
			name:     "ru_ashkenazi_exact",
			mode:     Ashkenazi,
			ruleset:  Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			name:    "ru_ashkenazi_approx",
			mode:    Ashkenazi,
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
		{
			name:     "ru_sephardic_exact",
			mode:     Sephardic,
			ruleset:  Exact,
			input:    "апельсин",
			expected: []string{},
		},
		{
			name:     "ru_sephardic_approx",
			mode:     Sephardic,
			ruleset:  Approx,
			input:    "апельсин",
			expected: []string{},
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
