package beidermorse

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runes_append(t *testing.T) {
	cases := []struct {
		r1       runes
		r2       runes
		expected runes
	}{
		{r1: nil, r2: runes("hello"), expected: runes("hello")},
		{r1: runes("hello"), r2: nil, expected: runes("hello")},
		{r1: runes("hello"), r2: runes("world"), expected: runes("helloworld")},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.r1.append(c.r2))
		})
	}
}

func Test_runes_contains(t *testing.T) {
	cases := []struct {
		runes    runes
		search   runes
		expected bool
	}{
		{runes: runes("orange"), search: runes("ra"), expected: true},
		{runes: runes("orange"), search: runes("rag"), expected: false},
		{runes: runes("orange"), search: runes("ora"), expected: true},
		{runes: runes("orange"), search: runes("orange"), expected: true},
		{runes: runes("orange"), search: runes("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.contains(c.search))
		})
	}
}

func Test_runes_hasPrefix(t *testing.T) {
	cases := []struct {
		runes    runes
		search   runes
		expected bool
	}{
		{runes: runes("orange"), search: runes("ra"), expected: false},
		{runes: runes("orange"), search: runes("ora"), expected: true},
		{runes: runes("orange"), search: runes("orange"), expected: true},
		{runes: runes("orange"), search: runes("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.hasPrefix(c.search))
		})
	}
}

func Test_runes_hasSuffix(t *testing.T) {
	cases := []struct {
		runes    runes
		search   runes
		expected bool
	}{
		{runes: runes("orange"), search: runes("ng"), expected: false},
		{runes: runes("orange"), search: runes("ge"), expected: true},
		{runes: runes("orange"), search: runes("orange"), expected: true},
		{runes: runes("orange"), search: runes("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.hasSuffix(c.search))
		})
	}
}

func Benchmark_runes_containsAt(b *testing.B) {
	r1 := runes("orange")
	r2 := runes("ang")

	for i := 0; i < b.N; i++ {
		r1.containsAt(r2, 2)
	}
}

func Test_runes_containsAt(t *testing.T) {
	cases := []struct {
		runes    runes
		search   runes
		index    int
		expected bool
	}{
		{runes: runes("orange"), search: runes("ra"), index: -1000, expected: true},
		{runes: runes("orange"), search: runes("ra"), index: 0, expected: false},
		{runes: runes("orange"), search: runes("ra"), index: 1, expected: true},
		{runes: runes("orange"), search: runes("ra"), index: 2, expected: false},
		{runes: runes("orange"), search: runes("ra"), index: 3, expected: false},
		{runes: runes("orange"), search: runes("ra"), index: 1000, expected: false},
		{runes: runes("orange"), search: runes("orange123"), index: 1000, expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.containsAt(c.search, c.index))
		})
	}
}
