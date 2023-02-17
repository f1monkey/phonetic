package exrunes

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	cases := []struct {
		runes    []rune
		search   []rune
		expected bool
	}{
		{runes: []rune("orange"), search: []rune("ra"), expected: true},
		{runes: []rune("orange"), search: []rune("rag"), expected: false},
		{runes: []rune("orange"), search: []rune("ora"), expected: true},
		{runes: []rune("orange"), search: []rune("orange"), expected: true},
		{runes: []rune("orange"), search: []rune("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, Contains(c.runes, c.search))
		})
	}
}

func Test_HasPrefix(t *testing.T) {
	cases := []struct {
		runes    []rune
		search   []rune
		expected bool
	}{
		{runes: []rune("orange"), search: []rune("ra"), expected: false},
		{runes: []rune("orange"), search: []rune("ora"), expected: true},
		{runes: []rune("orange"), search: []rune("orange"), expected: true},
		{runes: []rune("orange"), search: []rune("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, HasPrefix(c.runes, c.search))
		})
	}
}

func Test_HasSuffix(t *testing.T) {
	cases := []struct {
		runes    []rune
		search   []rune
		expected bool
	}{
		{runes: []rune("orange"), search: []rune("ng"), expected: false},
		{runes: []rune("orange"), search: []rune("ge"), expected: true},
		{runes: []rune("orange"), search: []rune("orange"), expected: true},
		{runes: []rune("orange"), search: []rune("orange123"), expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, HasSuffix(c.runes, c.search))
		})
	}
}

func Benchmark_ContainsAt(b *testing.B) {
	r1 := []rune("orange")
	r2 := []rune("ang")

	for i := 0; i < b.N; i++ {
		ContainsAt(r1, r2, 2)
	}
}

func Test_ContainsAt(t *testing.T) {
	cases := []struct {
		runes    []rune
		search   []rune
		index    int
		expected bool
	}{
		{runes: []rune("orange"), search: []rune("ra"), index: -1000, expected: true},
		{runes: []rune("orange"), search: []rune("ra"), index: 0, expected: false},
		{runes: []rune("orange"), search: []rune("ra"), index: 1, expected: true},
		{runes: []rune("orange"), search: []rune("ra"), index: 2, expected: false},
		{runes: []rune("orange"), search: []rune("ra"), index: 3, expected: false},
		{runes: []rune("orange"), search: []rune("ra"), index: 1000, expected: false},
		{runes: []rune("orange"), search: []rune("orange123"), index: 1000, expected: false},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, ContainsAt(c.runes, c.search, c.index))
		})
	}
}

func Benchmark_Equal(b *testing.B) {
	r1 := []rune{'a', 'b', 'c', 'd'}
	r2 := []rune{'a', 'b', 'c'}
	for i := 0; i < b.N; i++ {
		Equal(r1, r2)
	}
}

func Test_Equal(t *testing.T) {
	cases := []struct {
		r1, r2   []rune
		expected bool
	}{
		{
			r1:       []rune{'a', 'b', 'c'},
			r2:       []rune{'a', 'b', 'c'},
			expected: true,
		},
		{
			r1:       []rune{'a', 'b', 'c', 'd'},
			r2:       []rune{'a', 'b', 'c'},
			expected: false,
		},
		{
			r1:       nil,
			r2:       nil,
			expected: true,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c.expected, Equal(c.r1, c.r2))
		})
	}
}
