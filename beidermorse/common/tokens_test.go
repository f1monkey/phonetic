package common

import (
	"fmt"
	"strconv"
	"testing"

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
		{Text: []rune("k"), Langs: 1047288},
		{Text: []rune("ts"), Langs: 16392},
		{Text: []rune("dZ"), Langs: 524288},
	}
	t2 := Tokens{
		{Text: []rune("O"), Langs: -1},
		{Text: []rune("P"), Langs: 16384},
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
				{Text: []rune("O"), Langs: -1},
				{Text: []rune("P"), Langs: 16384},
			},
			dst: nil,
			expected: Tokens{
				{Text: []rune("O"), Langs: -1},
				{Text: []rune("P"), Langs: 16384},
			},
		},
		{
			src: Tokens{
				{Text: []rune("k"), Langs: 1047288},
				{Text: []rune("ts"), Langs: 16392},
				{Text: []rune("dZ"), Langs: 524288},
			},
			dst: Tokens{
				{Text: []rune("O"), Langs: -1},
				{Text: []rune("P"), Langs: 16384},
			},
			expected: Tokens{
				{Text: []rune("kO"), Langs: 1047288},
				{Text: []rune("kP"), Langs: 16384},
				{Text: []rune("tsO"), Langs: 16392},
				{Text: []rune("tsP"), Langs: 16384},
				{Text: []rune("dZO"), Langs: 524288},
			},
		},
		{
			src: Tokens{
				{Text: []rune("t"), Langs: 128},
			},
			dst: Tokens{
				{Text: []rune("i"), Langs: -1},
				{Text: []rune("Y"), Langs: 128},
			},
			expected: Tokens{
				{Text: []rune("ti"), Langs: 128},
				{Text: []rune("tY"), Langs: 128},
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
		{Text: []rune("foo"), Langs: 1},
		{Text: []rune("bar"), Langs: 1},
		{Text: []rune("foo"), Langs: 2},
		{Text: []rune("foo"), Langs: 1},
		{Text: []rune("foo"), Langs: 3},
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
				{Text: []rune("foo"), Langs: 1},
				{Text: []rune("bar"), Langs: 1},
				{Text: []rune("foo"), Langs: 2},
				{Text: []rune("foo"), Langs: 1},
				{Text: []rune("foo"), Langs: 3},
			},
			expected: Tokens{
				{Text: []rune("foo"), Langs: 1},
				{Text: []rune("bar"), Langs: 1},
				{Text: []rune("foo"), Langs: 2},
				{Text: []rune("foo"), Langs: 3},
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
