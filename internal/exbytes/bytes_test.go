package exbytes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_ReplacePrefix(b *testing.B) {
	src := []byte("orange")
	from := []byte("ora")
	to := []byte("q")

	for i := 0; i < b.N; i++ {
		ReplacePrefix(src, from, to)
	}
}

func Test_ReplacePrefix(t *testing.T) {
	cases := []struct {
		src      []byte
		from     []byte
		to       []byte
		expected []byte
	}{
		{src: []byte("orange"), from: []byte("ora"), to: []byte("qwe"), expected: []byte("qwenge")},
		{src: []byte("orange"), from: []byte("ora"), to: []byte("qw"), expected: []byte("qwnge")},
		{src: []byte("orange"), from: []byte("ora"), to: []byte(""), expected: []byte("nge")},
		{src: []byte("orange"), from: []byte("qwe"), to: []byte(""), expected: []byte("orange")},
		{src: []byte("orange"), from: []byte("ang"), to: []byte(""), expected: []byte("orange")},
		{src: []byte("orange"), from: []byte("orange1"), to: []byte(""), expected: []byte("orange")},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", string(c.src), string(c.from), string(c.to)), func(t *testing.T) {
			result := ReplacePrefix(c.src, c.from, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_ReplaceSuffix(b *testing.B) {
	src := []byte("orange")
	from := []byte("nge")
	to := []byte("q")

	for i := 0; i < b.N; i++ {
		ReplaceSuffix(src, from, to)
	}
}

func Test_ReplaceSuffix(t *testing.T) {
	cases := []struct {
		src      []byte
		from     []byte
		to       []byte
		expected []byte
	}{
		{src: []byte("orange"), from: []byte("nge"), to: []byte("qwe"), expected: []byte("oraqwe")},
		{src: []byte("orange"), from: []byte("nge"), to: []byte("qw"), expected: []byte("oraqw")},
		{src: []byte("orange"), from: []byte("nge"), to: []byte(""), expected: []byte("ora")},
		{src: []byte("orange"), from: []byte("qwe"), to: []byte(""), expected: []byte("orange")},
		{src: []byte("orange"), from: []byte("ang"), to: []byte(""), expected: []byte("orange")},
		{src: []byte("orange"), from: []byte("orange1"), to: []byte(""), expected: []byte("orange")},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", string(c.src), string(c.from), string(c.to)), func(t *testing.T) {
			result := ReplaceSuffix(c.src, c.from, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_ReplaceAll(b *testing.B) {
	src := []byte("orange")
	from := []byte("an")
	to := []byte("a")

	for i := 0; i < b.N; i++ {
		ReplaceAll(src, from, to)
	}
}

func Test_ReplaceAll(t *testing.T) {
	cases := []struct {
		src      []byte
		from     []byte
		to       []byte
		expected []byte
	}{
		{src: []byte("tool"), from: []byte("o"), to: []byte("b"), expected: []byte("tbbl")},
		{src: []byte("tool"), from: []byte("o"), to: []byte(""), expected: []byte("tl")},
		{src: []byte("orange"), from: []byte("ang"), to: []byte("a"), expected: []byte("orae")},
		{src: []byte("orange"), from: []byte("ora"), to: []byte("o"), expected: []byte("onge")},
		{src: []byte("orange"), from: []byte("ora"), to: []byte("or"), expected: []byte("ornge")},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", string(c.src), string(c.from), string(c.to)), func(t *testing.T) {
			result := ReplaceAll(c.src, c.from, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}

func Test_ReplaceSequence(t *testing.T) {
	cases := []struct {
		src      []byte
		sequence []byte
		to       []byte
		expected []byte
	}{
		{src: []byte("tool"), sequence: []byte("o"), to: []byte("b"), expected: []byte("tbl")},
		{src: []byte("abcqweqweabc"), sequence: []byte("qwe"), to: []byte(""), expected: []byte("abcabc")},
		{src: []byte("abcqweqweabc"), sequence: []byte("abc"), to: []byte(""), expected: []byte("qweqwe")},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", string(c.src), string(c.sequence), string(c.to)), func(t *testing.T) {
			result := ReplaceSequence(c.src, c.sequence, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}
