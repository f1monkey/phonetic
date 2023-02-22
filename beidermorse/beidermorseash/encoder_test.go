package beidermorseash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode_Ru_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_Ru_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_En_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_En_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		accuracy Accuracy
		input    string
		expected []string
	}{
		{
			accuracy: Approx,
			input:    "orange",
			expected: []string{
				"orangi",
				"orongi",
				"orYngi",
				"Yrangi",
				"Yrongi",
				"YrYngi",
				"oranzi",
				"oronzi",
				"orani",
				"oroni",
				"oranxi",
				"oronxi",
				"urangi",
				"urongi",
			},
		},
		{
			accuracy: Exact,
			input:    "orange",
			expected: []string{
				"orange",
				"orandZe",
				"oranhe",
				"oranxe",
			},
		},
		{
			accuracy: Exact,
			input:    "van der orange",
			expected: []string{
				"vander",
				"fander",
				"orange",
				"orandZe",
				"oranhe",
				"oranxe",
				"vanderorange",
				"vanderorandZe",
				"vanderoranhe",
				"vanderoranxe",
				"fanderorange",
			},
		},
		{
			accuracy: Approx,
			input:    "test",
			expected: []string{
				"tist",
				"tYst",
				"tinst",
				"tonst",
			},
		},
		{
			accuracy: Exact,
			input:    "test",
			expected: []string{
				"teSt",
				"test",
			},
		},
		{
			accuracy: Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			accuracy: Approx,
			input:    "апельсин",
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
		t.Run(fmt.Sprintf("%s-%s", c.accuracy, c.input), func(t *testing.T) {
			e, err := NewEncoder(WithAccuracy(c.accuracy))
			require.NoError(t, err)
			require.Equal(t, c.expected, e.Encode(c.input))
		})
	}
}
