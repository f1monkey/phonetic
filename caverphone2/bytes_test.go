package caverphone2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_bytesReplacePrefix(b *testing.B) {
	src := []byte("orange")
	from := []byte("ora")
	to := []byte("q")

	for i := 0; i < b.N; i++ {
		bytesReplacePrefix(src, from, to)
	}
}

func Test_bytesReplacePrefix(t *testing.T) {
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
			result := bytesReplacePrefix(c.src, c.from, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_bytesReplaceSuffix(b *testing.B) {
	src := []byte("orange")
	from := []byte("nge")
	to := []byte("q")

	for i := 0; i < b.N; i++ {
		bytesReplaceSuffix(src, from, to)
	}
}

func Test_bytesReplaceSuffix(t *testing.T) {
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
			result := bytesReplaceSuffix(c.src, c.from, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_bytesReplaceAll(b *testing.B) {
	src := []byte("orange")
	from := []byte("an")
	to := []byte("a")

	for i := 0; i < b.N; i++ {
		bytesReplaceAll(src, from, to)
	}
}

func Test_bytesReplaceAll(t *testing.T) {
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
			result := bytesReplaceAll(c.src, c.from, c.to)
			assert.Equal(t, c.expected, result)
		})
	}
}
