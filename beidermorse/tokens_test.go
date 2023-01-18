package beidermorse

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_languageID_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		languageID(1047288).merge(16384)
	}
}

func Test_languageID_merge(t *testing.T) {
	cases := []struct {
		src      []languageID
		expected languageID
	}{
		{
			src:      []languageID{1, 128, 16384},
			expected: 0,
		},
		{
			src:      []languageID{4, 128, 16384},
			expected: 0,
		},
		{
			src:      []languageID{1047288, 16384},
			expected: 16384,
		},
		{
			src:      []languageID{1047288, 16384},
			expected: 16384,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := c.src[0].merge(c.src[1:]...)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_tokens_merge(b *testing.B) {
	t1 := tokens{
		{text: runes("k"), langs: 1047288},
		{text: runes("ts"), langs: 16392},
		{text: runes("dZ"), langs: 524288},
	}
	t2 := tokens{
		{text: runes("O"), langs: -1},
		{text: runes("P"), langs: 16384},
	}

	for i := 0; i < b.N; i++ {
		t1.merge(langsAny, t2)
	}
}

func Test_tokens_merge(t *testing.T) {
	cases := []struct {
		tokens   tokens
		src      []tokens
		expected tokens
	}{
		{
			tokens: tokens{
				{text: runes("O"), langs: -1},
				{text: runes("P"), langs: 16384},
			},
			src: nil,
			expected: tokens{
				{text: runes("O"), langs: -1},
				{text: runes("P"), langs: 16384},
			},
		},
		{
			tokens: tokens{
				{text: runes("k"), langs: 1047288},
				{text: runes("ts"), langs: 16392},
				{text: runes("dZ"), langs: 524288},
			},
			src: []tokens{
				{
					{text: runes("O"), langs: -1},
					{text: runes("P"), langs: 16384},
				},
			},
			expected: tokens{
				{text: runes("kO"), langs: 1047288},
				{text: runes("kP"), langs: 16384},
				{text: runes("tsO"), langs: 16392},
				{text: runes("tsP"), langs: 16384},
				{text: runes("dZO"), langs: 524288},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := c.tokens.merge(langsAny, c.src...)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_tokens_deduplicate(b *testing.B) {
	src := tokens{
		{text: runes("foo"), langs: 1},
		{text: runes("bar"), langs: 1},
		{text: runes("foo"), langs: 2},
		{text: runes("foo"), langs: 1},
		{text: runes("foo"), langs: 3},
	}

	for i := 0; i < b.N; i++ {
		src.deduplicate()
	}
}

func Test_tokens_deduplicate(t *testing.T) {
	cases := []struct {
		src      tokens
		expected tokens
	}{
		{
			src: tokens{
				{text: runes("foo"), langs: 1},
				{text: runes("bar"), langs: 1},
				{text: runes("foo"), langs: 2},
				{text: runes("foo"), langs: 1},
				{text: runes("foo"), langs: 3},
			},
			expected: tokens{
				{text: runes("foo"), langs: 1},
				{text: runes("bar"), langs: 1},
				{text: runes("foo"), langs: 2},
				{text: runes("foo"), langs: 3},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := c.src.deduplicate()
			require.Equal(t, c.expected, result)
		})
	}
}
