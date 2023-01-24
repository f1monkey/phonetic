package caverphone2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode(b *testing.B) {
	e := NewEncoder()
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "add",
			expected: "AT11111111",
		},
		{
			input:    "eat",
			expected: "AT11111111",
		},
		{
			input:    "hold",
			expected: "AT11111111",
		},
		{
			input:    "orange",
			expected: "ARNK111111",
		},
		{
			input:    "test",
			expected: "TST1111111",
		},
		{
			input:    "ready",
			expected: "RTA1111111",
		},
		{
			input:    "апельсин",
			expected: "1111111111",
		},
		{
			input:    "ብርቱካናማ",
			expected: "1111111111",
		},
		{
			input:    "کینو",
			expected: "1111111111",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			e := NewEncoder()
			result := e.Encode(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
