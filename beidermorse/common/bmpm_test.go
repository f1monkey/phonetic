package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_prepareInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prepareInput("o'range orange", Generic)
	}
}

func Test_prepareInput(t *testing.T) {
	cases := []struct {
		input    string
		mode     Mode
		expected string
	}{
		{
			input:    "Orange",
			mode:     Generic,
			expected: "orange",
		},
		{
			input:    "de la orange",
			mode:     Generic,
			expected: "dela orange",
		},
		{
			input:    "van der orange",
			mode:     Generic,
			expected: "vander orange",
		},
		{
			input:    "van den orange",
			mode:     Generic,
			expected: "vanden orange",
		},
		{
			input:    "o'range orange orange",
			mode:     Generic,
			expected: "orange orangeorange",
		},
		{
			input:    "o'range orange orange",
			mode:     Sephardic,
			expected: "o rangeorangeorange",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s_%s", c.mode, c.input), func(t *testing.T) {
			result := prepareInput(c.input, c.mode)
			require.Equal(t, c.expected, result)
		})
	}
}
