package beidermorse

import (
	"fmt"
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
		mode     Mode
		ruleset  Ruleset
		input    string
		expected []string
	}{
		{
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
			mode:    Generic,
			ruleset: Exact,
			input:   "test",
			expected: []string{
				"teSt",
				"test",
			},
		},
		{
			mode:     Generic,
			ruleset:  Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
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
			mode:     Ashkenazi,
			ruleset:  Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
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
			mode:     Sephardic,
			ruleset:  Exact,
			input:    "апельсин",
			expected: []string{},
		},
		{
			mode:     Sephardic,
			ruleset:  Approx,
			input:    "апельсин",
			expected: []string{},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", c.mode, c.ruleset, c.input), func(t *testing.T) {
			e, err := NewEncoder(WithNameMode(c.mode), WithRuleset(c.ruleset))
			require.NoError(t, err)
			require.Equal(t, c.expected, e.Encode(c.input))
		})
	}
}
