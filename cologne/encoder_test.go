package cologne

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode(b *testing.B) {
	e := NewEncoder()
	s := "Elon Musk"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Encode(s)
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		src      string
		expected string
	}{
		{src: "Breschnew", expected: "17863"},
		{src: "Elon Musk", expected: "05684"},
		{src: "aaaaaaaaaaa", expected: "0"},
		{src: "orange", expected: "0764"},
		{src: "light-grey", expected: "54247"},
		{src: "dark-grey", expected: "2747"},
		{src: "Müller-Lüdenscheidt", expected: "65752682"},
		{src: "Großtraktor", expected: "47827427"},
	}

	e := NewEncoder()
	for _, c := range cases {
		t.Run(c.src, func(t *testing.T) {
			result := e.Encode(c.src)
			require.Equal(t, c.expected, result)
		})
	}
}
