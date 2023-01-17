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
	t := tokens{
		{text: "k", langs: 1047288},
		{text: "ts", langs: 16392},
		{text: "dZ", langs: 524288},
	}

	for i := 0; i < b.N; i++ {
		t.merge(langsAny, tokens{
			{text: "O", langs: -1},
			{text: "P", langs: 16384},
		})
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
				{text: "O", langs: -1},
				{text: "P", langs: 16384},
			},
			src: nil,
			expected: tokens{
				{text: "O", langs: -1},
				{text: "P", langs: 16384},
			},
		},
		{
			tokens: tokens{
				{text: "k", langs: 1047288},
				{text: "ts", langs: 16392},
				{text: "dZ", langs: 524288},
			},
			src: []tokens{
				{
					{text: "O", langs: -1},
					{text: "P", langs: 16384},
				},
			},
			expected: tokens{
				{text: "kO", langs: 1047288},
				{text: "kP", langs: 16384},
				{text: "tsO", langs: 16392},
				{text: "tsP", langs: 16384},
				{text: "dZO", langs: 524288},
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
		{text: "foo", langs: 1},
		{text: "bar", langs: 1},
		{text: "foo", langs: 2},
		{text: "foo", langs: 1},
		{text: "foo", langs: 3},
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
				{text: "foo", langs: 1},
				{text: "bar", langs: 1},
				{text: "foo", langs: 2},
				{text: "foo", langs: 1},
				{text: "foo", langs: 3},
			},
			expected: tokens{
				{text: "foo", langs: 1},
				{text: "bar", langs: 1},
				{text: "foo", langs: 2},
				{text: "foo", langs: 3},
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
