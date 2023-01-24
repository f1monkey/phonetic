package caverphone2

import "regexp"

type rule struct {
	substr []byte
	prefix []byte
	suffix []byte
	regexp *regexp.Regexp
	to     []byte
}

var filter = regexp.MustCompile("[^a-z]")

var rules = []rule{
	{
		suffix: []byte("e"),
		to:     []byte(""),
	},
	{
		prefix: []byte("cough"),
		to:     []byte("cou2f"),
	},
	{
		prefix: []byte("rough"),
		to:     []byte("rou2f"),
	},
	{
		prefix: []byte("tough"),
		to:     []byte("tou2f"),
	},
	{
		prefix: []byte("enough"),
		to:     []byte("enou2f"),
	},
	{
		prefix: []byte("trough"),
		to:     []byte("trou2f"),
	},
	{
		prefix: []byte("gn"),
		to:     []byte("2n"),
	},
	{
		suffix: []byte("mb"),
		to:     []byte("m2"),
	},
	{
		substr: []byte("cq"),
		to:     []byte("2q"),
	},
	{
		substr: []byte("ci"),
		to:     []byte("si"),
	},
	{
		substr: []byte("ce"),
		to:     []byte("se"),
	},
	{
		substr: []byte("cy"),
		to:     []byte("sy"),
	},
	{
		substr: []byte("tch"),
		to:     []byte("2ch"),
	},
	{
		substr: []byte("c"),
		to:     []byte("k"),
	},
	{
		substr: []byte("q"),
		to:     []byte("k"),
	},
	{
		substr: []byte("x"),
		to:     []byte("k"),
	},
	{
		substr: []byte("v"),
		to:     []byte("f"),
	},
	{
		substr: []byte("dg"),
		to:     []byte("2g"),
	},
	{
		substr: []byte("tio"),
		to:     []byte("sio"),
	},
	{
		substr: []byte("tia"),
		to:     []byte("sia"),
	},
	{
		substr: []byte("d"),
		to:     []byte("t"),
	},
	{
		substr: []byte("ph"),
		to:     []byte("fh"),
	},
	{
		substr: []byte("b"),
		to:     []byte("p"),
	},
	{
		substr: []byte("sh"),
		to:     []byte("s2"),
	},
	{
		substr: []byte("z"),
		to:     []byte("s"),
	},

	// ^[aeiou]
	{
		prefix: []byte("a"),
		to:     []byte("A"),
	},
	{
		prefix: []byte("e"),
		to:     []byte("A"),
	},
	{
		prefix: []byte("i"),
		to:     []byte("A"),
	},
	{
		prefix: []byte("o"),
		to:     []byte("A"),
	},
	{
		prefix: []byte("u"),
		to:     []byte("A"),
	},

	// [aeiou]
	{
		substr: []byte("a"),
		to:     []byte("3"),
	},
	{
		substr: []byte("e"),
		to:     []byte("3"),
	},
	{
		substr: []byte("i"),
		to:     []byte("3"),
	},
	{
		substr: []byte("o"),
		to:     []byte("3"),
	},
	{
		substr: []byte("u"),
		to:     []byte("3"),
	},

	{
		substr: []byte("j"),
		to:     []byte("y"),
	},
	{
		substr: []byte("y3"),
		to:     []byte("Y3"),
	},
	{
		prefix: []byte("y"),
		to:     []byte("A"),
	},
	{
		substr: []byte("y"),
		to:     []byte("3"),
	},
	{
		substr: []byte("3gh3"),
		to:     []byte("3kh3"),
	},
	{
		substr: []byte("gh"),
		to:     []byte("22"),
	},
	{
		substr: []byte("g"),
		to:     []byte("k"),
	},
	{
		regexp: regexp.MustCompile("s+"),
		to:     []byte("S"),
	},
	{
		regexp: regexp.MustCompile("t+"),
		to:     []byte("T"),
	},
	{
		regexp: regexp.MustCompile("p+"),
		to:     []byte("P"),
	},
	{
		regexp: regexp.MustCompile("k+"),
		to:     []byte("K"),
	},
	{
		regexp: regexp.MustCompile("f+"),
		to:     []byte("F"),
	},
	{
		regexp: regexp.MustCompile("m+"),
		to:     []byte("M"),
	},
	{
		regexp: regexp.MustCompile("n+"),
		to:     []byte("N"),
	},
	{
		substr: []byte("w3"),
		to:     []byte("W3"),
	},
	{
		substr: []byte("wh3"),
		to:     []byte("Wh3"),
	},
	{
		suffix: []byte("w"),
		to:     []byte("3"),
	},
	{
		substr: []byte("w"),
		to:     []byte("2"),
	},
	{
		prefix: []byte("h"),
		to:     []byte("A"),
	},
	{
		substr: []byte("h"),
		to:     []byte("2"),
	},
	{
		substr: []byte("r3"),
		to:     []byte("R3"),
	},
	{
		suffix: []byte("r"),
		to:     []byte("3"),
	},
	{
		substr: []byte("r"),
		to:     []byte("2"),
	},
	{
		substr: []byte("l3"),
		to:     []byte("L3"),
	},
	{
		suffix: []byte("l"),
		to:     []byte("3"),
	},
	{
		substr: []byte("l"),
		to:     []byte("2"),
	},
	{
		substr: []byte("2"),
		to:     []byte(""),
	},
	{
		suffix: []byte("3"),
		to:     []byte("A"),
	},
	{
		substr: []byte("3"),
		to:     []byte(""),
	},
}
