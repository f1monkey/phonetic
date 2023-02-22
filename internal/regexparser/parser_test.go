package regexparser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	cases := []struct {
		pattern  string
		expected Result
		ok       bool
	}{
		{"^$", Result{}, true},
		{"^abc", Result{IsPrefix: true, Items: []string{"abc"}}, true},
		{"abc$", Result{IsSuffix: true, Items: []string{"abc"}}, true},
		{"^abc$", Result{IsExact: true, Items: []string{"abc"}}, true},
		{
			"[abc]foo",
			Result{
				Items: []string{"afoo", "bfoo", "cfoo"},
			},
			true,
		},
		{
			"[qwe]foo[zxc]",
			Result{
				Items: []string{"qfooz", "qfoox", "qfooc", "wfooz", "wfoox", "wfooc", "efooz", "efoox", "efooc"},
			},
			true,
		},
		{
			"(b|a|r)$",
			Result{
				IsSuffix: true,
				Items:    []string{"b", "a", "r"},
			},
			true,
		},
		{
			"(b|a|r)foo$",
			Result{
				IsSuffix: true,
				Items:    []string{"bfoo", "afoo", "rfoo"},
			},
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.pattern, func(t *testing.T) {
			result, ok := Parse(c.pattern)
			require.Equal(t, c.ok, ok)
			require.Equal(t, c.expected, result)
		})
	}
}
