package metaphone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode(b *testing.B) {
	e := NewEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		maxLen   int
	}{
		{"", "", 4},
		{"A", "A", 4},
		{"ж", "", 4},
		{"a", "A", 4},
		{"wr", "R", 4},
		{"Metaphone", "MTFN", 4},
		{"garbage", "KRBJ", 4},
		{"tech", "TX", 4},
		{"orange", "ORNJ", 4},
		{"Incomprehensibility", "INKM", 4},
		{"Incomprehensibility", "INKMPRHNSBLT", 100},
		{"Incomprehensibility", "INKMPRHNSBLT", -1},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			e := NewEncoder()
			e.maxLen = c.maxLen
			require.Equal(t, c.expected, e.Encode(c.input))
		})
	}
}

func Benchmark_filter(b *testing.B) {
	input := []rune("qwertyйцукен")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter(input)
	}
}

func Test_filter(t *testing.T) {
	cases := []struct {
		input    []rune
		expected []rune
	}{
		{nil, nil},
		{[]rune{}, []rune{}},
		{[]rune("q"), []rune("Q")},
		{[]rune("q123we"), []rune("QWE")},
		{[]rune("qwe1"), []rune("QWE")},
		{[]rune("1qwe1"), []rune("QWE")},
		{[]rune("1qweйцукен1"), []rune("QWE")},
	}

	for _, c := range cases {
		t.Run(string(c.input), func(t *testing.T) {
			require.Equal(t, c.expected, filter(c.input))
		})
	}
}

func Test_regionMatch(t *testing.T) {
	cases := []struct {
		input    []rune
		index    int
		chars    []rune
		expected bool
	}{
		{nil, 0, nil, false},
		{[]rune("HELLO"), 100, []rune("LL"), false},
		{[]rune("HELLO"), 2, []rune("LL"), true},
		{[]rune("HELLO"), 3, []rune("LO"), true},
		{[]rune("HELLO"), 4, []rune("LO"), false},
		{[]rune("HELL"), 0, []rune("HELL"), true},
	}

	for _, c := range cases {
		t.Run(string(c.input)+"_"+fmt.Sprint(c.index), func(t *testing.T) {
			require.Equal(t, c.expected, regionMatch(c.input, c.index, c.chars))
		})
	}
}

func Test_isVowel(t *testing.T) {
	cases := []struct {
		input    []rune
		index    int
		expected bool
	}{
		{nil, 0, false},
		{[]rune("AEIOU"), 100, false},
		{[]rune("BAB"), 1, true},
	}

	for _, c := range cases {
		t.Run(string(c.input)+"_"+fmt.Sprint(c.index), func(t *testing.T) {
			require.Equal(t, c.expected, isVowel(c.input, c.index))
		})
	}
}

func Test_isFrontv(t *testing.T) {
	cases := []struct {
		input    []rune
		index    int
		expected bool
	}{
		{nil, 0, false},
		{[]rune("EIY"), 100, false},
		{[]rune("AEA"), 1, true},
	}

	for _, c := range cases {
		t.Run(string(c.input)+"_"+fmt.Sprint(c.index), func(t *testing.T) {
			require.Equal(t, c.expected, isFrontv(c.input, c.index))
		})
	}
}

func Test_isVarson(t *testing.T) {
	cases := []struct {
		input    []rune
		index    int
		expected bool
	}{
		{nil, 0, false},
		{[]rune("CSPTG"), 100, false},
		{[]rune("ACA"), 1, true},
	}

	for _, c := range cases {
		t.Run(string(c.input)+"_"+fmt.Sprint(c.index), func(t *testing.T) {
			require.Equal(t, c.expected, isVarson(c.input, c.index))
		})
	}
}
