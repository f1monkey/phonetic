// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorsesep

import (
	"github.com/f1monkey/phonetic/beidermorse/common"
	"regexp"
)

type Lang common.Lang

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

var Rules = map[common.Lang]common.Rules{
	common.Lang(Any): {
		{
			Pattern: []rune("ph"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gli"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiou]$"),
			},
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiou]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tx"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tg"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix:           []rune("y"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gue"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ao]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ñ"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bdgv]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ptckf]"),
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiu]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ao]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[cs]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[cdglnrst]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bfpv]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ouh"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aioe]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aioe]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aioe]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aie]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuoóeéê]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuiíoóeéê]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiíou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("î"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ô"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õ"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(French): {
		{
			Pattern: []rune("kh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiy]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gue"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiy]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("que"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiouyéèê]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiouyéèê]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[bdgt]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ouh"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aioe]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeio]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeio]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eau"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("où"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[ou]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aoeu]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Hebrew): {
		{
			Pattern: []rune("אי"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("עי"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("עו"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("VV"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("או"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("VV"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ג׳"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ד׳"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("א"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("L"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ב"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ג"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ד"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("וו"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("V"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("וי"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("WW"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ו"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("W"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ז"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ח"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("X"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ט"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("T"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("יי"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("י"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ך"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("X"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("כ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("K"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("כ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ל"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ם"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("מ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ן"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("נ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ס"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ע"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("L"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ף"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("פ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ץ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("C"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("צ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("C"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ק"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("K"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ר"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ש"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ת"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("T"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Italian): {
		{
			Pattern: []rune("kh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gli"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[bdgt]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cc"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiou]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeou]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeou]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("�"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.RuleToken{
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
	common.Lang(Portuguese): {
		{
			Pattern: []rune("kh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sç"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bdgv]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ptckf]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiu]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ao]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[eiu]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ao]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[bdgt]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[cs]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuiíoóeéê]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiíou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bcdfglnprstv]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aáuoóeéê]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ô"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("õ"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Spanish): {
		{
			Pattern: []rune("ñ"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ig"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiou]$"),
			},
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[aeiou]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tx"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tg"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("[dgt]$"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[bpvf]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.RuleToken{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aei]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.RuleToken{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ò"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
	},
}

var LangRules = []common.LangRule{
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("eau"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ou"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gni"),
		},
		Langs:  10,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("tx"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("tj"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gy"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("guy"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("sh"),
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("lh"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("nh"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ny"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gue"),
		},
		Langs:  34,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gui"),
		},
		Langs:  34,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gia"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gie"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gio"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("giu"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ñ"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("â"),
		},
		Langs:  18,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("á"),
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("à"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ã"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ê"),
		},
		Langs:  18,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("í"),
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("î"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ô"),
		},
		Langs:  18,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("õ"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ò"),
		},
		Langs:  40,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ú"),
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ù"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ü"),
		},
		Langs:  48,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("א"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ב"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ג"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ד"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ה"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ו"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ז"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ח"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ט"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("י"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("כ"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ל"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("מ"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("נ"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ס"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ע"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("פ"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("צ"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ק"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ר"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ש"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ת"),
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("a"),
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("o"),
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("e"),
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("i"),
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("y"),
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("u"),
		},
		Langs:  4,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("kh"),
		},
		Langs:  32,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gua"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("guo"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ç"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("cha"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("cho"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("chu"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("j"),
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("dj"),
		},
		Langs:  32,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("sce"),
		},
		Langs:  2,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("sci"),
		},
		Langs:  2,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ó"),
		},
		Langs:  2,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("è"),
		},
		Langs:  16,
		Accept: false,
	},
}

var FinalRules = common.FinalRules{
	Approx: common.FinalRule{
		First: common.Rules{
			{
				Pattern: []rune("h"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[fktSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("p"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("b"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pktSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("f"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("v"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pftSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("k"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("g"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("g"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("t"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("d"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("d"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("dZ"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("tS"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("nm"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("a"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("b"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("d"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("e"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("f"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("g"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("i"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("k"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("l"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("l"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("m"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("m"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("n"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("o"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("o"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("p"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("r"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("r"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("t"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("u"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("u"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("v"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("z"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("mbr"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("mr"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("mpr"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("mr"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bens"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[bp]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ha"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("h"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uj"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("au"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eu"),
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ja"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("ia"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("je"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jo"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("iu"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ja"),
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jo"),
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("j"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("u"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Se"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[rmnl]$"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[rmnl]$"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dS"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("be"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[fktSs]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pktSs]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pftSs]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				Phonetic: []common.RuleToken{
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
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
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("V"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[common.Lang]common.Rules{
			common.Lang(Any):        {},
			common.Lang(French):     {},
			common.Lang(Hebrew):     {},
			common.Lang(Italian):    {},
			common.Lang(Portuguese): {},
			common.Lang(Spanish):    {},
		},
	},
	Exact: common.FinalRule{
		First: common.Rules{
			{
				Pattern: []rune("h"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[fktSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("p"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("p"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("b"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pktSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("f"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("f"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("v"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pftSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("k"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("k"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("g"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("g"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("t"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("t"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("d"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("d"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("dZ"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("tS"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("nm"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("m"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ji"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("i"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("a"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("b"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("b"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("d"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("d"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("e"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("f"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("f"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("g"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("g"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("i"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("i"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("k"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("k"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("l"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("l"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("m"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("m"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("n"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("o"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("o"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("p"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("p"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("r"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("r"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("t"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("t"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("u"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("u"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("v"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("v"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix:           []rune("z"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("h"),
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("[^t]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[bgZd]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[pfkst]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("S"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("S"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[bgzd]"),
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("Z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("z"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("B"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("b"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("V"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("v"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[common.Lang]common.Rules{
			common.Lang(Any):        {},
			common.Lang(French):     {},
			common.Lang(Hebrew):     {},
			common.Lang(Italian):    {},
			common.Lang(Portuguese): {},
			common.Lang(Spanish):    {},
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
