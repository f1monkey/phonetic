package beidermorsesep

import (
	"fmt"
	"testing"

	"github.com/f1monkey/phonetic/beidermorse/common"
	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode_Ru_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_Ru_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_En_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_En_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		accuracy common.Accuracy
		input    string
		expected []string
	}{
		{
			accuracy: common.Approx,
			input:    "orange",
			expected: []string{
				"uranzi",
				"uranz",
				"uranS",
				"uranzi",
				"uranz",
				"uranhi",
				"uranh",
			},
		},
		{
			accuracy: common.Exact,
			input:    "orange",
			expected: []string{
				"oranZe",
				"oranS",
				"orandZe",
				"oranxe",
			},
		},
		{
			accuracy: common.Exact,
			input:    "van der orange",
			expected: []string{
				"vander",
				"bander",
				"oranZe",
				"oranS",
				"orandZe",
				"oranxe",
				"vanderoranZe",
				"vanderoranS",
				"vanderorandZe",
				"vanderoranxe",
				"banderorandZe",
				"banderoranxe",
			},
		},
		{
			accuracy: common.Approx,
			input:    "test",
			expected: []string{
				"tist",
				"tst",
				"tist",
				"tst",
			},
		},
		{
			accuracy: common.Exact,
			input:    "test",
			expected: []string{
				"test",
				"teSt",
			},
		},

		// this behavior is different from the behavior of the original library
		{
			accuracy: common.Exact,
			input:    "апельсин",
			expected: []string{"апельсин"},
		},
		{
			accuracy: common.Approx,
			input:    "апельсин",
			expected: []string{
				"апельсин",
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
