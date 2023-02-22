// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorsesep

import (
	"github.com/f1monkey/phonetic/internal/bmpm"
	"regexp"
)

type Lang bmpm.Lang

const (
	Any Lang = 1 << iota
	French
	Hebrew
	Italian
	Portuguese
	Spanish
)

func (l Lang) String() string {
	switch l {
	case Any:
		return "any"
	case French:
		return "french"
	case Hebrew:
		return "hebrew"
	case Italian:
		return "italian"
	case Portuguese:
		return "portuguese"
	case Spanish:
		return "spanish"
	}
	return ""
}

const All = French +
	Hebrew +
	Italian +
	Portuguese +
	Spanish

var Rules = map[bmpm.Lang]bmpm.Rules{
	bmpm.Lang(Any): {
		{
			Pattern: []rune("ph"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gli"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gli"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("gni"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gni"),
					Langs: -1,
				},
				{
					Text:  []rune("ni"),
					Langs: 10,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: 10,
				},
				{
					Text:  []rune("nj"),
					Langs: 10,
				},
				{
					Text:  []rune("gn"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("gh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("dh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("bh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
				{
					Text:  []rune("bh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  []rune("th"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: 16,
				},
				{
					Text:  []rune("lh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: 16,
				},
				{
					Text:  []rune("nh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ig"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ig"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("ix"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tx"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tg"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tg"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gZ"),
					Langs: 18,
				},
				{
					Text:  []rune("dZ"),
					Langs: 40,
				},
				{
					Text:  []rune("x"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: 18,
				},
				{
					Text:  []rune("dZ"),
					Langs: 40,
				},
				{
					Text:  []rune("x"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("guy"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gue"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: 2,
				},
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ñ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ny"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("sç"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: 8,
				},
				{
					Text:  []rune("S"),
					Langs: 18,
				},
				{
					Text:  []rune("tS"),
					Langs: 32,
				},
				{
					Text:  []rune("dZ"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 32,
				},
				{
					Text:  []rune("dZ"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: 8,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cc"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: 8,
				},
				{
					Text:  []rune("ks"),
					Langs: 50,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: 8,
				},
				{
					Text:  []rune("s"),
					Langs: 50,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: 32,
				},
				{
					Text:  []rune("z"),
					Langs: 26,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("d"),
					[]rune("g"),
					[]rune("l"),
					[]rune("m"),
					[]rune("n"),
					[]rune("r"),
					[]rune("v"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 8,
				},
				{
					Text:  []rune("S"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("v"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("dz"),
					Langs: 8,
				},
				{
					Text:  []rune("Z"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("p"),
					[]rune("t"),
					[]rune("c"),
					[]rune("k"),
					[]rune("f"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 8,
				},
				{
					Text:  []rune("S"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("dz"),
					Langs: 8,
				},
				{
					Text:  []rune("ts"),
					Langs: 8,
				},
				{
					Text:  []rune("s"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("que"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: 2,
				},
				{
					Text:  []rune("ke"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ez"),
					Langs: 16,
				},
				{
					Text:  []rune("eS"),
					Langs: 16,
				},
				{
					Text:  []rune("eks"),
					Langs: -1,
				},
				{
					Text:  []rune("egz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("c"),
					[]rune("s"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: 16,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("l"),
					[]rune("n"),
					[]rune("r"),
					[]rune("s"),
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("f"),
					[]rune("p"),
					[]rune("v"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: 48,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("b"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
				{
					Text:  []rune("V"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("v"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("B"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("eau"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ouh"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("o"),
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: 2,
				},
				{
					Text:  []rune("uh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uh"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("o"),
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("uh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("o"),
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vo"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 2,
				},
			},
		},
		{
			Pattern: []rune("ão"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ãe"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ãi"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õe"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("où"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: 2,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("î"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ô"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ò"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
				{
					Text:  []rune("v"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: 32,
				},
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("b"),
					Langs: 32,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 48,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(French): {
		{
			Pattern: []rune("kh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("gn"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gue"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("que"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ouh"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("o"),
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("uh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vo"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eau"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("où"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oi"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("e"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Hebrew): {
		{
			Pattern: []rune("אי"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("עי"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("עו"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("VV"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("או"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("VV"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ג׳"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ד׳"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("א"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("L"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ב"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ג"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ד"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("וו"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("V"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("וי"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("WW"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ו"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("W"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ז"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ח"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("X"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ט"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("T"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("יי"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("י"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ך"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("X"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("כ"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("K"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("כ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ל"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ם"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("מ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ן"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("נ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ס"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ע"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("L"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ף"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("פ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ץ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("C"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("צ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("C"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ק"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("K"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ר"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ש"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ת"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("T"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Italian): {
		{
			Pattern: []rune("kh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gli"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("gli"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
				{
					Text:  []rune("gn"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gni"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ni"),
					Langs: -1,
				},
				{
					Text:  []rune("gni"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cc"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vo"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
				{
					Text:  []rune("dz"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Portuguese): {
		{
			Pattern: []rune("kh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sç"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("d"),
					[]rune("g"),
					[]rune("l"),
					[]rune("m"),
					[]rune("n"),
					[]rune("r"),
					[]rune("v"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("v"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("p"),
					[]rune("t"),
					[]rune("c"),
					[]rune("k"),
					[]rune("f"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vo"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ez"),
					Langs: -1,
				},
				{
					Text:  []rune("eS"),
					Langs: -1,
				},
				{
					Text:  []rune("eks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("c"),
					[]rune("s"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("f"),
					[]rune("g"),
					[]rune("l"),
					[]rune("n"),
					[]rune("p"),
					[]rune("r"),
					[]rune("s"),
					[]rune("t"),
					[]rune("v"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ão"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ãe"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ãi"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õe"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
				{
					Text:  []rune("an"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ô"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Spanish): {
		{
			Pattern: []rune("ñ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ny"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ig"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
				{
					Text:  []rune("ig"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ix"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tx"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tg"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tg"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("bh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("B"),
					Langs: -1,
				},
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
				{
					Text:  []rune("V"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
					[]rune("v"),
					[]rune("f"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vo"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ò"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
	},
}

var LangRules = []bmpm.LangRule{
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eau"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ou"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gni"),
			},
		},
		Langs:  10,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tx"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tj"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("guy"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sh"),
			},
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("lh"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("nh"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ny"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gue"),
			},
		},
		Langs:  34,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gui"),
			},
		},
		Langs:  34,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gia"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gie"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gio"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("giu"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ñ"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("â"),
			},
		},
		Langs:  18,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("á"),
			},
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("à"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ã"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ê"),
			},
		},
		Langs:  18,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("í"),
			},
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("î"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ô"),
			},
		},
		Langs:  18,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("õ"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ò"),
			},
		},
		Langs:  40,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ú"),
			},
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ù"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ü"),
			},
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("א"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ב"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ג"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ד"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ה"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ו"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ז"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ח"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ט"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("י"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("כ"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ל"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("מ"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("נ"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ס"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ע"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("פ"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("צ"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ק"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ר"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ש"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ת"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("a"),
			},
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("o"),
			},
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("e"),
			},
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("i"),
			},
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("y"),
			},
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("u"),
			},
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("kh"),
			},
		},
		Langs:  32,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gua"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("guo"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ç"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cha"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cho"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("chu"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("j"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dj"),
			},
		},
		Langs:  32,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sce"),
			},
		},
		Langs:  2,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sci"),
			},
		},
		Langs:  2,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ó"),
			},
		},
		Langs:  2,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("è"),
			},
		},
		Langs:  16,
		Accept: false,
	},
}

var FinalRules = bmpm.FinalRules{
	Approx: bmpm.FinalRule{
		First: bmpm.Rules{
			{
				Pattern: []rune("h"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("k"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("g"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("g"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("d"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("d"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("dZ"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("tS"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("nm"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("a"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("d"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("e"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("g"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("i"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("k"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("l"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("m"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("m"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("n"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("o"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("o"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("r"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("u"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("u"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("mbr"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("mr"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("mpr"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("mr"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bens"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("binz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("benS"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("binz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ben"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bin"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bar"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bar"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("abens"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("binz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("abenS"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("binz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aben"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bin"),
						Langs: -1,
					},
					{
						Text:  []rune("bun"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("abe"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bi"),
						Langs: -1,
					},
					{
						Text:  []rune("bu"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("abi"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bi"),
						Langs: -1,
					},
					{
						Text:  []rune("bu"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("abou"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bu"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 2,
					},
				},
			},
			{
				Pattern: []rune("abu"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bu"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bou"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bu"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 2,
					},
				},
			},
			{
				Pattern: []rune("bu"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("bu"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ibn"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("ibn"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("els"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("ilz"),
						Langs: -1,
					},
					{
						Text:  []rune("lz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("elS"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("ilz"),
						Langs: -1,
					},
					{
						Text:  []rune("lz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("el"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("il"),
						Langs: -1,
					},
					{
						Text:  []rune("l"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("als"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("alS"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lz"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("al"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("l"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dal"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("dal"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 8,
					},
				},
			},
			{
				Pattern: []rune("da"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("da"),
						Langs: -1,
					},
					{
						Text:  []rune("a"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("della"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("dila"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dela"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("dila"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("del"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("dil"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("des"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("dis"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("de"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("di"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("di"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("di"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 8,
					},
				},
			},
			{
				Pattern: []rune("do"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("du"),
						Langs: -1,
					},
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("du"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("du"),
						Langs: -1,
					},
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oa"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("va"),
						Langs: -1,
					},
					{
						Text:  []rune("a"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oe"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("vi"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ae"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("a"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("p"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ha"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("ha"),
						Langs: -1,
					},
					{
						Text:  []rune("a"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("h"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
					{
						Text:  []rune("h"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("x"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("h"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("h"),
						Langs: -1,
					},
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aja"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("Da"),
						Langs: -1,
					},
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aje"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("Di"),
						Langs: -1,
					},
					{
						Text:  []rune("Da"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aji"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("Di"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajo"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("Du"),
						Langs: -1,
					},
					{
						Text:  []rune("Da"),
						Langs: -1,
					},
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aju"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("Du"),
						Langs: -1,
					},
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aj"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ej"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oj"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uj"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("au"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eu"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ou"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ja"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("je"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jo"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ju"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("a"),
						Langs: -1,
					},
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("je"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("j"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("o"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("a"),
						Langs: -1,
					},
					{
						Text:  []rune("u"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("o"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("a"),
						Langs: -1,
					},
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("se"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
					{
						Text:  []rune("si"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Se"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
					{
						Text:  []rune("si"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dS"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("be"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
					{
						Text:  []rune("bi"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("pe"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
					{
						Text:  []rune("pi"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ve"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
					{
						Text:  []rune("vi"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("fe"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
					{
						Text:  []rune("fi"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ge"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
					{
						Text:  []rune("gi"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ke"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("g"),
						Langs: -1,
					},
					{
						Text:  []rune("ki"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("de"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
					{
						Text:  []rune("di"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("te"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("d"),
						Langs: -1,
					},
					{
						Text:  []rune("ti"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ze"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
					{
						Text:  []rune("zi"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("B"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("V"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[bmpm.Lang]bmpm.Rules{
			bmpm.Lang(Any):        {},
			bmpm.Lang(French):     {},
			bmpm.Lang(Hebrew):     {},
			bmpm.Lang(Italian):    {},
			bmpm.Lang(Portuguese): {},
			bmpm.Lang(Spanish):    {},
		},
	},
	Exact: bmpm.FinalRule{
		First: bmpm.Rules{
			{
				Pattern: []rune("h"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("k"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("g"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("g"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("s"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("d"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("d"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("dZ"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("tS"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("nm"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("a"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("d"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("e"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("f"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("g"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("i"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("k"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("l"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("l"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("m"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("m"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("n"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("o"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("o"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("r"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("u"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("u"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("v"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("z"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("h"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[^t]$"),
				},
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("d"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("s"),
						[]rune("t"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("g"),
						[]rune("z"),
						[]rune("d"),
					},
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("Z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("B"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("V"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[bmpm.Lang]bmpm.Rules{
			bmpm.Lang(Any):        {},
			bmpm.Lang(French):     {},
			bmpm.Lang(Hebrew):     {},
			bmpm.Lang(Italian):    {},
			bmpm.Lang(Portuguese): {},
			bmpm.Lang(Spanish):    {},
		},
	},
}

var Discards = []string{
	"abe",
	"aben",
	"abi",
	"abou",
	"abu",
	"al",
	"bar",
	"ben",
	"bou",
	"bu",
	"d",
	"da",
	"dal",
	"de",
	"del",
	"dela",
	"della",
	"des",
	"di",
	"el",
	"la",
	"le",
	"ibn",
	"ha",
}
