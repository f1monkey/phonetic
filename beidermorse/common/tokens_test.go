package common

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/f1monkey/phonetic/internal/exrunes"
	"github.com/stretchr/testify/require"
)

func Benchmark_Lang_Merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lang(1047288).merge(16384)
	}
}

func Test_Lang_Merge(t *testing.T) {
	cases := []struct {
		src      []Lang
		expected Lang
	}{
		{
			src:      []Lang{1, 128, 16384},
			expected: 0,
		},
		{
			src:      []Lang{4, 128, 16384},
			expected: 0,
		},
		{
			src:      []Lang{1047288, 16384},
			expected: 16384,
		},
		{
			src:      []Lang{1047288, 16384},
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

func Benchmark_Tokens_Merge(b *testing.B) {
	t1 := Tokens{
		{Text: exrunes.Runes("k"), Langs: 1047288},
		{Text: exrunes.Runes("ts"), Langs: 16392},
		{Text: exrunes.Runes("dZ"), Langs: 524288},
	}
	t2 := Tokens{
		{Text: exrunes.Runes("O"), Langs: -1},
		{Text: exrunes.Runes("P"), Langs: 16384},
	}

	for i := 0; i < b.N; i++ {
		t1.Merge(LangAny, t2)
	}
}

func Test_Tokens_Merge(t *testing.T) {
	cases := []struct {
		src      Tokens
		dst      Tokens
		expected Tokens
	}{
		{
			src: Tokens{
				{Text: exrunes.Runes("O"), Langs: -1},
				{Text: exrunes.Runes("P"), Langs: 16384},
			},
			dst: nil,
			expected: Tokens{
				{Text: exrunes.Runes("O"), Langs: -1},
				{Text: exrunes.Runes("P"), Langs: 16384},
			},
		},
		{
			src: Tokens{
				{Text: exrunes.Runes("k"), Langs: 1047288},
				{Text: exrunes.Runes("ts"), Langs: 16392},
				{Text: exrunes.Runes("dZ"), Langs: 524288},
			},
			dst: Tokens{
				{Text: exrunes.Runes("O"), Langs: -1},
				{Text: exrunes.Runes("P"), Langs: 16384},
			},
			expected: Tokens{
				{Text: exrunes.Runes("kO"), Langs: 1047288},
				{Text: exrunes.Runes("kP"), Langs: 16384},
				{Text: exrunes.Runes("tsO"), Langs: 16392},
				{Text: exrunes.Runes("tsP"), Langs: 16384},
				{Text: exrunes.Runes("dZO"), Langs: 524288},
			},
		},
		{
			src: Tokens{
				{Text: exrunes.Runes("t"), Langs: 128},
			},
			dst: Tokens{
				{Text: exrunes.Runes("i"), Langs: -1},
				{Text: exrunes.Runes("Y"), Langs: 128},
			},
			expected: Tokens{
				{Text: exrunes.Runes("ti"), Langs: 128},
				{Text: exrunes.Runes("tY"), Langs: 128},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := c.src.Merge(LangAny, c.dst)
			require.Equal(t, c.expected, result)
		})
	}
}

func Benchmark_Tokens_Deduplicate(b *testing.B) {
	src := Tokens{
		{Text: exrunes.Runes("foo"), Langs: 1},
		{Text: exrunes.Runes("bar"), Langs: 1},
		{Text: exrunes.Runes("foo"), Langs: 2},
		{Text: exrunes.Runes("foo"), Langs: 1},
		{Text: exrunes.Runes("foo"), Langs: 3},
	}

	for i := 0; i < b.N; i++ {
		src.Deduplicate()
	}
}

func Test_Tokens_Deduplicate(t *testing.T) {
	cases := []struct {
		src      Tokens
		expected Tokens
	}{
		{
			src: Tokens{
				{Text: exrunes.Runes("foo"), Langs: 1},
				{Text: exrunes.Runes("bar"), Langs: 1},
				{Text: exrunes.Runes("foo"), Langs: 2},
				{Text: exrunes.Runes("foo"), Langs: 1},
				{Text: exrunes.Runes("foo"), Langs: 3},
			},
			expected: Tokens{
				{Text: exrunes.Runes("foo"), Langs: 1},
				{Text: exrunes.Runes("bar"), Langs: 1},
				{Text: exrunes.Runes("foo"), Langs: 2},
				{Text: exrunes.Runes("foo"), Langs: 3},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := c.src.Deduplicate()
			require.Equal(t, c.expected, result)
		})
	}
}
