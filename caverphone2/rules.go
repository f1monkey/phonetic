package caverphone2

import "regexp"

type rule struct {
	pattern     []byte
	regexp      *regexp.Regexp
	replaceWith []byte
}

var rules = []rule{
	{
		regexp:      regexp.MustCompile("[^a-z]"),
		replaceWith: []byte(""),
	},
	{
		regexp:      regexp.MustCompile("e$"),
		replaceWith: []byte(""),
	},
	{
		regexp:      regexp.MustCompile("^cough"),
		replaceWith: []byte("cou2f"),
	},
	{
		regexp:      regexp.MustCompile("^rough"),
		replaceWith: []byte("rou2f"),
	},
	{
		regexp:      regexp.MustCompile("^tough"),
		replaceWith: []byte("tou2f"),
	},
	{
		regexp:      regexp.MustCompile("^enough"),
		replaceWith: []byte("enou2f"),
	},
	{
		regexp:      regexp.MustCompile("^trough"),
		replaceWith: []byte("trou2f"),
	},
	{
		regexp:      regexp.MustCompile("^gn"),
		replaceWith: []byte("2n"),
	},
	{
		regexp:      regexp.MustCompile("mb$"),
		replaceWith: []byte("m2"),
	},
	{
		pattern:     []byte("cq"),
		replaceWith: []byte("2q"),
	},
	{
		pattern:     []byte("ci"),
		replaceWith: []byte("si"),
	},
	{
		pattern:     []byte("ce"),
		replaceWith: []byte("se"),
	},
	{
		pattern:     []byte("cy"),
		replaceWith: []byte("sy"),
	},
	{
		pattern:     []byte("tch"),
		replaceWith: []byte("2ch"),
	},
	{
		pattern:     []byte("c"),
		replaceWith: []byte("k"),
	},
	{
		pattern:     []byte("q"),
		replaceWith: []byte("k"),
	},
	{
		pattern:     []byte("x"),
		replaceWith: []byte("k"),
	},
	{
		pattern:     []byte("v"),
		replaceWith: []byte("f"),
	},
	{
		pattern:     []byte("dg"),
		replaceWith: []byte("2g"),
	},
	{
		pattern:     []byte("tio"),
		replaceWith: []byte("sio"),
	},
	{
		pattern:     []byte("tia"),
		replaceWith: []byte("sia"),
	},
	{
		pattern:     []byte("d"),
		replaceWith: []byte("t"),
	},
	{
		pattern:     []byte("ph"),
		replaceWith: []byte("fh"),
	},
	{
		pattern:     []byte("b"),
		replaceWith: []byte("p"),
	},
	{
		pattern:     []byte("sh"),
		replaceWith: []byte("s2"),
	},
	{
		pattern:     []byte("z"),
		replaceWith: []byte("s"),
	},
	{
		regexp:      regexp.MustCompile("^[aeiou]"),
		replaceWith: []byte("A"),
	},
	{
		regexp:      regexp.MustCompile("[aeiou]"),
		replaceWith: []byte("3"),
	},
	{
		pattern:     []byte("j"),
		replaceWith: []byte("y"),
	},
	{
		regexp:      regexp.MustCompile("^y3"),
		replaceWith: []byte("Y3"),
	},
	{
		regexp:      regexp.MustCompile("^y"),
		replaceWith: []byte("A"),
	},
	{
		pattern:     []byte("y"),
		replaceWith: []byte("3"),
	},
	{
		pattern:     []byte("3gh3"),
		replaceWith: []byte("3kh3"),
	},
	{
		pattern:     []byte("gh"),
		replaceWith: []byte("22"),
	},
	{
		pattern:     []byte("g"),
		replaceWith: []byte("k"),
	},
	{
		regexp:      regexp.MustCompile("s+"),
		replaceWith: []byte("S"),
	},
	{
		regexp:      regexp.MustCompile("t+"),
		replaceWith: []byte("T"),
	},
	{
		regexp:      regexp.MustCompile("p+"),
		replaceWith: []byte("P"),
	},
	{
		regexp:      regexp.MustCompile("k+"),
		replaceWith: []byte("K"),
	},
	{
		regexp:      regexp.MustCompile("f+"),
		replaceWith: []byte("F"),
	},
	{
		regexp:      regexp.MustCompile("m+"),
		replaceWith: []byte("M"),
	},
	{
		regexp:      regexp.MustCompile("n+"),
		replaceWith: []byte("N"),
	},
	{
		pattern:     []byte("w3"),
		replaceWith: []byte("W3"),
	},
	{
		pattern:     []byte("wh3"),
		replaceWith: []byte("Wh3"),
	},
	{
		regexp:      regexp.MustCompile("w$"),
		replaceWith: []byte("3"),
	},
	{
		pattern:     []byte("w"),
		replaceWith: []byte("2"),
	},
	{
		regexp:      regexp.MustCompile("^h"),
		replaceWith: []byte("A"),
	},
	{
		pattern:     []byte("h"),
		replaceWith: []byte("2"),
	},
	{
		pattern:     []byte("r3"),
		replaceWith: []byte("R3"),
	},
	{
		regexp:      regexp.MustCompile("r$"),
		replaceWith: []byte("3"),
	},
	{
		pattern:     []byte("r"),
		replaceWith: []byte("2"),
	},
	{
		pattern:     []byte("l3"),
		replaceWith: []byte("L3"),
	},
	{
		regexp:      regexp.MustCompile("l$"),
		replaceWith: []byte("3"),
	},
	{
		pattern:     []byte("l"),
		replaceWith: []byte("2"),
	},
	{
		pattern:     []byte("2"),
		replaceWith: []byte(""),
	},
	{
		regexp:      regexp.MustCompile("3$"),
		replaceWith: []byte("A"),
	},
	{
		pattern:     []byte("3"),
		replaceWith: []byte(""),
	},
}
