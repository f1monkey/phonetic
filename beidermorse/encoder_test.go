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

func Benchmark_Encoder_Encode_Ru_Exact(b *testing.B) {
	e := MustNewEncoder(WithNameMode(Generic), WithRuleset(Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_En_Approx(b *testing.B) {
	e := MustNewEncoder(WithNameMode(Generic), WithRuleset(Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_En_Exact(b *testing.B) {
	e := MustNewEncoder(WithNameMode(Generic), WithRuleset(Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
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
			name:    "en_generic_approx",
			mode:    Generic,
			ruleset: Approx,
			input:   "orange",
			expected: []string{
				"orangi",
				"oragi",
				"orongi",
				"orogi",
				"orYngi",
				"Yrangi",
				"Yrongi",
				"YrYngi",
				"oranxi",
				"oronxi",
				"orani",
				"oroni",
				"oranii",
				"oronii",
				"oranzi",
				"oronzi",
				"urangi",
				"urongi",
			},
		},
		{
			name:    "en_generic_exact",
			mode:    Generic,
			ruleset: Exact,
			input:   "orange",
			expected: []string{
				"orange",
				"oranxe",
				"oranhe",
				"oranje",
				"oranZe",
				"orandZe",
			},
		},
		{
			name:    "en_generic_approx",
			mode:    Generic,
			ruleset: Approx,
			input:   "test",
			expected: []string{
				"tist",
				"tYst",
				"tis",
				"tit",
				"ti",
			},
		},
		{
			name:    "en_generic_exact",
			mode:    Generic,
			ruleset: Exact,
			input:   "test",
			expected: []string{
				"teSt",
				"test",
			},
		},
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
