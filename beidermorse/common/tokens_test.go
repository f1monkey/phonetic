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

func Benchmark_Tokens_MergeRules(b *testing.B) {
	t2 := []RuleToken{
		{Text: []rune("O"), Langs: -1},
		{Text: []rune("P"), Langs: 16384},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		t1 := &Tokens{
			Buf: func() *exrunes.Buffer {
				buf := exrunes.NewBuffer(100)
				buf.Add([]rune("k"))
				buf.Add([]rune("ts"))
				buf.Add([]rune("dZ"))
				return buf
			}(),
		}
		t1.Items = make([]Token, 10)
		t1.Items = []Token{
			{ID: 0, Langs: 1047288},
			{ID: 1, Langs: 16392},
			{ID: 2, Langs: 524288},
		}
		b.StartTimer()
		t1.MergeRules(t2, LangAny)
	}
}

type expectedToken struct {
	text  []rune
	langs Lang
}

func Test_Tokens_MergeRuleTokens(t *testing.T) {
	cases := []struct {
		src      Tokens
		dst      []RuleToken
		expected []expectedToken
	}{
		{
			src: Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("O"))
					buf.Add([]rune("P"))
					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: -1},
					{ID: 1, Langs: 16384},
				}},
			dst: nil,
			expected: []expectedToken{
				{text: []rune("O"), langs: -1},
				{text: []rune("P"), langs: 16384},
			},
		},
		{
			src: Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("k"))
					buf.Add([]rune("ts"))
					buf.Add([]rune("dZ"))
					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: 1047288},
					{ID: 1, Langs: 16392},
					{ID: 2, Langs: 524288},
				}},
			dst: []RuleToken{
				{Text: []rune("O"), Langs: -1},
				{Text: []rune("P"), Langs: 16384},
			},
			expected: []expectedToken{
				{text: []rune("kO"), langs: 1047288},
				{text: []rune("kP"), langs: 16384},
				{text: []rune("tsO"), langs: 16392},
				{text: []rune("tsP"), langs: 16384},
				{text: []rune("dZO"), langs: 524288},
			},
		},
		{
			src: Tokens{
				Buf: func() *exrunes.Buffer {
					buf := exrunes.NewBuffer(100)
					buf.Add([]rune("t"))
					return buf
				}(),
				Items: []Token{
					{ID: 0, Langs: 128},
				},
			},
			dst: []RuleToken{
				{Text: []rune("i"), Langs: -1},
				{Text: []rune("Y"), Langs: 128},
			},
			expected: []expectedToken{
				{text: []rune("ti"), langs: 128},
				{text: []rune("tY"), langs: 128},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c.src.MergeRules(c.dst, LangAny)

			result := make([]expectedToken, len(c.src.Items))
			for i, t := range c.src.Items {
				result[i] = expectedToken{text: c.src.Buf.Get(t.ID), langs: t.Langs}
			}

			require.Equal(t, c.expected, result)
		})
	}
}

// func Benchmark_Tokens_Deduplicate(b *testing.B) {
// 	src := Tokens{
// 		{ID: []rune("foo"), Langs: 1},
// 		{ID: []rune("bar"), Langs: 1},
// 		{ID: []rune("foo"), Langs: 2},
// 		{ID: []rune("foo"), Langs: 1},
// 		{ID: []rune("foo"), Langs: 3},
// 	}

// 	for i := 0; i < b.N; i++ {
// 		src.Deduplicate()
// 	}
// }

// func Test_Tokens_Deduplicate(t *testing.T) {
// 	cases := []struct {
// 		src      Tokens
// 		expected Tokens
// 	}{
// 		{
// 			src: Tokens{
// 				{ID: []rune("foo"), Langs: 1},
// 				{ID: []rune("bar"), Langs: 1},
// 				{ID: []rune("foo"), Langs: 2},
// 				{ID: []rune("foo"), Langs: 1},
// 				{ID: []rune("foo"), Langs: 3},
// 			},
// 			expected: Tokens{
// 				{ID: []rune("foo"), Langs: 1},
// 				{ID: []rune("bar"), Langs: 1},
// 				{ID: []rune("foo"), Langs: 2},
// 				{ID: []rune("foo"), Langs: 3},
// 			},
// 		},
// 	}

// 	for i, c := range cases {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			result := c.src.Deduplicate()
// 			require.Equal(t, c.expected, result)
// 		})
// 	}
// }
