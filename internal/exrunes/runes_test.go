package exrunes

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Runes_Contains(t *testing.T) {
	cases := []struct {
		runes    Runes
		search   Runes
		expected bool
	}{
		{runes: Runes("orange"), search: Runes("ra"), expected: true},
		{runes: Runes("orange"), search: Runes("rag"), expected: false},
		{runes: Runes("orange"), search: Runes("ora"), expected: true},
		{runes: Runes("orange"), search: Runes("orange"), expected: true},
		{runes: Runes("orange"), search: Runes("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.Contains(c.search))
		})
	}
}

func Test_Runes_HasPrefix(t *testing.T) {
	cases := []struct {
		runes    Runes
		search   Runes
		expected bool
	}{
		{runes: Runes("orange"), search: Runes("ra"), expected: false},
		{runes: Runes("orange"), search: Runes("ora"), expected: true},
		{runes: Runes("orange"), search: Runes("orange"), expected: true},
		{runes: Runes("orange"), search: Runes("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.HasPrefix(c.search))
		})
	}
}

func Test_Runes_HasSuffix(t *testing.T) {
	cases := []struct {
		runes    Runes
		search   Runes
		expected bool
	}{
		{runes: Runes("orange"), search: Runes("ng"), expected: false},
		{runes: Runes("orange"), search: Runes("ge"), expected: true},
		{runes: Runes("orange"), search: Runes("orange"), expected: true},
		{runes: Runes("orange"), search: Runes("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.HasSuffix(c.search))
		})
	}
}

func Benchmark_Runes_ContainsAt(b *testing.B) {
	r1 := Runes("orange")
	r2 := Runes("ang")

	for i := 0; i < b.N; i++ {
		r1.ContainsAt(r2, 2)
	}
}

func Test_Runes_ContainsAt(t *testing.T) {
	cases := []struct {
		runes    Runes
		search   Runes
		index    int
		expected bool
	}{
		{runes: Runes("orange"), search: Runes("ra"), index: -1000, expected: true},
		{runes: Runes("orange"), search: Runes("ra"), index: 0, expected: false},
		{runes: Runes("orange"), search: Runes("ra"), index: 1, expected: true},
		{runes: Runes("orange"), search: Runes("ra"), index: 2, expected: false},
		{runes: Runes("orange"), search: Runes("ra"), index: 3, expected: false},
		{runes: Runes("orange"), search: Runes("ra"), index: 1000, expected: false},
		{runes: Runes("orange"), search: Runes("orange123"), index: 1000, expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, c.runes.ContainsAt(c.search, c.index))
		})
	}
}
