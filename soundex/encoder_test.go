package soundex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode(b *testing.B) {
	s := "orange"
	e := NewEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Encode(s)
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: "Robert", expected: "R163"},
		{input: "Rupert", expected: "R163"},
		{input: "Rubin", expected: "R150"},
		{input: "Ashcraft", expected: "A261"},
		{input: "Ashcroft", expected: "A261"},
		{input: "Tymczak", expected: "T522"},
		{input: "Pfister", expected: "P236"},
		{input: "Honeyman", expected: "H555"},
		{input: "Апельсин", expected: "0000"},
		{input: "Orange", expected: "O652"},
	}

	e := NewEncoder()
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			require.Equal(t, c.expected, e.Encode(c.input))
		})
	}
}

func Benchmark_clean(b *testing.B) {
	data := []rune("ora1nge")
	for i := 0; i < b.N; i++ {
		clean(data)
	}
}

func Test_clean(t *testing.T) {
	cases := []struct {
		input    []rune
		expected []rune
	}{
		{input: []rune{}, expected: []rune{}},
		{input: []rune("a"), expected: []rune("A")},
		{input: []rune("orange"), expected: []rune("ORANGE")},
		{input: []rune("orange1"), expected: []rune("ORANGE")},
		{input: []rune("oran1ge"), expected: []rune("ORANGE")},
		{input: []rune("o1ran111ge"), expected: []rune("ORANGE")},
		{input: []rune("1orange"), expected: []rune("ORANGE")},
		{input: []rune("апельсин "), expected: []rune("АПЕЛЬСИН")},
	}

	for _, c := range cases {
		t.Run(string(c.input), func(t *testing.T) {
			require.Equal(t, c.expected, clean(c.input))
		})
	}
}
