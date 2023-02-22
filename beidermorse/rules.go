// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorse

import (
	"github.com/f1monkey/phonetic/beidermorse/common"
	"regexp"
)

type Lang common.Lang

const (
	Any Lang = 1 << iota
	Arabic
	Cyrillic
	Czech
	Dutch
	English
	French
	German
	Greek
	Greeklatin
	Hebrew
	Hungarian
	Italian
	Latvian
	Polish
	Portuguese
	Romanian
	Russian
	Spanish
	Turkish
)

func (l Lang) String() string {
	switch l {
	case Any:
		return "any"
	case Arabic:
		return "arabic"
	case Cyrillic:
		return "cyrillic"
	case Czech:
		return "czech"
	case Dutch:
		return "dutch"
	case English:
		return "english"
	case French:
		return "french"
	case German:
		return "german"
	case Greek:
		return "greek"
	case Greeklatin:
		return "greeklatin"
	case Hebrew:
		return "hebrew"
	case Hungarian:
		return "hungarian"
	case Italian:
		return "italian"
	case Latvian:
		return "latvian"
	case Polish:
		return "polish"
	case Portuguese:
		return "portuguese"
	case Romanian:
		return "romanian"
	case Russian:
		return "russian"
	case Spanish:
		return "spanish"
	case Turkish:
		return "turkish"
	}
	return ""
}

const All = Arabic +
	Cyrillic +
	Czech +
	Dutch +
	English +
	French +
	German +
	Greek +
	Greeklatin +
	Hebrew +
	Hungarian +
	Italian +
	Latvian +
	Polish +
	Portuguese +
	Romanian +
	Russian +
	Spanish +
	Turkish

var Rules = map[common.Lang]common.Rules{
	common.Lang(Any): {
		{
			Pattern: []rune("yna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("in"),
					Langs: 131072,
				},
				{
					Text:  []rune("ina"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ina"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("in"),
					Langs: 131072,
				},
				{
					Text:  []rune("ina"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("liova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: 131072,
				},
				{
					Text:  []rune("lef"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: 131072,
				},
				{
					Text:  []rune("lef"),
					Langs: 131072,
				},
				{
					Text:  []rune("l"),
					Langs: 8,
				},
				{
					Text:  []rune("el"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("kova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kova"),
					Langs: -1,
				},
				{
					Text:  []rune("kof"),
					Langs: 131072,
				},
				{
					Text:  []rune("k"),
					Langs: 8,
				},
				{
					Text:  []rune("ek"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  []rune("of"),
					Langs: 131072,
				},
				{
					Text:  nil,
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ová"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("eva"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("eva"),
					Langs: -1,
				},
				{
					Text:  []rune("ef"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("aia"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("aja"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("aya"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lowa"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: 16384,
				},
				{
					Text:  []rune("l"),
					Langs: 16384,
				},
				{
					Text:  []rune("el"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("kowa"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kova"),
					Langs: -1,
				},
				{
					Text:  []rune("kof"),
					Langs: 16384,
				},
				{
					Text:  []rune("k"),
					Langs: 16384,
				},
				{
					Text:  []rune("ek"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("owa"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  []rune("of"),
					Langs: 16384,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lowna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lovna"),
					Langs: -1,
				},
				{
					Text:  []rune("levna"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 16384,
				},
				{
					Text:  []rune("el"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("kowna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kovna"),
					Langs: -1,
				},
				{
					Text:  []rune("k"),
					Langs: 16384,
				},
				{
					Text:  []rune("ek"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("owna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovna"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("lówna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("el"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kówna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ówna"),
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
			Pattern: []rune("á"),
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
					Langs: 8,
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
					Langs: 16392,
				},
			},
		},
		{
			Pattern: []rune("pf"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("pf"),
					Langs: -1,
				},
				{
					Text:  []rune("p"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
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
					Langs: 64,
				},
				{
					Text:  []rune("ke"),
					Langs: -1,
				},
				{
					Text:  []rune("kve"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
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
			Pattern: []rune("m"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("f"),
					[]rune("p"),
					[]rune("v"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: 32832,
				},
			},
		},
		{
			Pattern: []rune("ly"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("lj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("li"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("lj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lyo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("lt"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lt"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 64,
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
					Text:  []rune("f"),
					Langs: 128,
				},
				{
					Text:  []rune("b"),
					Langs: 262144,
				},
			},
		},
		{
			Pattern: []rune("ex"),
			RightContext: &common.Matcher{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ez"),
					Langs: 32768,
				},
				{
					Text:  []rune("eS"),
					Langs: 32768,
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
				Prefix: [][]rune{
					[]rune("c"),
					[]rune("s"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: 32768,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("tsk"),
					Langs: 16392,
				},
			},
		},
		{
			Pattern: []rune("cz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
				{
					Text:  []rune("tsz"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("rh"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("dh"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("bh"),
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
		{
			Pattern: []rune("ph"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ph"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: 131104,
				},
				{
					Text:  []rune("kh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lh"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("nh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nh"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("ssch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("xS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
				{
					Text:  []rune("sk"),
					Langs: 69632,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("sk"),
					Langs: 69632,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("ssh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("sh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: 131104,
				},
				{
					Text:  []rune("sh"),
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
			Pattern: []rune("zh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: 131104,
				},
				{
					Text:  []rune("zh"),
					Langs: -1,
				},
				{
					Text:  []rune("tsh"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("chs"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: 128,
				},
				{
					Text:  []rune("xs"),
					Langs: -1,
				},
				{
					Text:  []rune("tSs"),
					Langs: 131104,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 393248,
				},
				{
					Text:  []rune("k"),
					Langs: 69632,
				},
				{
					Text:  []rune("S"),
					Langs: 32832,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 393248,
				},
				{
					Text:  []rune("S"),
					Langs: 32832,
				},
			},
		},
		{
			Pattern: []rune("th"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("th"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: 672,
				},
				{
					Text:  []rune("th"),
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
			},
		},
		{
			Pattern: []rune("gh"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: 70144,
				},
				{
					Text:  []rune("gh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ouh"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("o"),
					[]rune("e"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: 64,
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("i"),
					[]rune("o"),
					[]rune("e"),
				},
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
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: 66048,
				},
				{
					Text:  []rune("H"),
					Langs: 381024,
				},
			},
		},
		{
			Pattern: []rune("cia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSa"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cią"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSom"),
					Langs: -1,
				},
				{
					Text:  []rune("tsom"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cią"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSon"),
					Langs: 16384,
				},
				{
					Text:  []rune("tson"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cię"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSem"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cię"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSen"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: 16384,
				},
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSo"),
					Langs: 16384,
				},
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ciu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSu"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sci"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Si"),
					Langs: 4096,
				},
				{
					Text:  []rune("stsi"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZi"),
					Langs: 524288,
				},
				{
					Text:  []rune("tSi"),
					Langs: 81920,
				},
				{
					Text:  []rune("tS"),
					Langs: 65536,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sc"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: 4096,
				},
				{
					Text:  []rune("sts"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZ"),
					Langs: 524288,
				},
				{
					Text:  []rune("tS"),
					Langs: 81920,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsi"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZi"),
					Langs: 524288,
				},
				{
					Text:  []rune("tSi"),
					Langs: 81920,
				},
				{
					Text:  []rune("tS"),
					Langs: 65536,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZ"),
					Langs: 524288,
				},
				{
					Text:  []rune("tS"),
					Langs: 81920,
				},
				{
					Text:  []rune("k"),
					Langs: 512,
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("stS"),
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ssz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("sz"),
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
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: 2048,
				},
				{
					Text:  []rune("sts"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("ssp"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: 128,
				},
				{
					Text:  []rune("sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sp"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: 128,
				},
				{
					Text:  []rune("sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sst"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("St"),
					Langs: 128,
				},
				{
					Text:  []rune("st"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("st"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("St"),
					Langs: 128,
				},
				{
					Text:  []rune("st"),
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
			Pattern: []rune("sj"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("sj"),
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
			Pattern: []rune("sj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("sj"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: 16,
				},
				{
					Text:  []rune("sx"),
					Langs: 262144,
				},
				{
					Text:  []rune("sZ"),
					Langs: 589824,
				},
			},
		},
		{
			Pattern: []rune("sia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sa"),
					Langs: 16384,
				},
				{
					Text:  []rune("sa"),
					Langs: 16384,
				},
				{
					Text:  []rune("sja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sią"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Som"),
					Langs: 16384,
				},
				{
					Text:  []rune("som"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sią"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Son"),
					Langs: 16384,
				},
				{
					Text:  []rune("son"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("się"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sem"),
					Langs: 16384,
				},
				{
					Text:  []rune("sem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("się"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sen"),
					Langs: 16384,
				},
				{
					Text:  []rune("sen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
				{
					Text:  []rune("sje"),
					Langs: -1,
				},
				{
					Text:  []rune("Se"),
					Langs: 16384,
				},
				{
					Text:  []rune("zi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("So"),
					Langs: 16384,
				},
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("siu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Su"),
					Langs: 16384,
				},
				{
					Text:  []rune("sju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("si"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ë"),
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Si"),
					Langs: 16384,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: 37056,
				},
			},
		},
		{
			Pattern: []rune("si"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Si"),
					Langs: 16384,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
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
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("á"),
					[]rune("u"),
					[]rune("í"),
					[]rune("o"),
					[]rune("ó"),
					[]rune("e"),
					[]rune("é"),
					[]rune("ê"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: 37056,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ë"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			RightContext: &common.Matcher{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: 32768,
				},
				{
					Text:  nil,
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: 32768,
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
					Langs: 64,
				},
				{
					Text:  []rune("gve"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: 64,
				},
				{
					Text:  []rune("gv"),
					Langs: 294912,
				},
			},
		},
		{
			Pattern: []rune("gu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gv"),
					Langs: -1,
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
			Pattern: []rune("gli"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("glI"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: 4096,
				},
			},
		},
		{
			Pattern: []rune("gni"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gnI"),
					Langs: -1,
				},
				{
					Text:  []rune("ni"),
					Langs: 4160,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: 4160,
				},
				{
					Text:  []rune("nj"),
					Langs: 4160,
				},
				{
					Text:  []rune("gn"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: 512,
				},
				{
					Text:  []rune("dZe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: 512,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("y"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 4096,
				},
				{
					Text:  []rune("j"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("gge"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("y"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gE"),
					Langs: -1,
				},
				{
					Text:  []rune("xe"),
					Langs: 262144,
				},
				{
					Text:  []rune("gZe"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZe"),
					Langs: 331808,
				},
				{
					Text:  []rune("je"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("y"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("xi"),
					Langs: 262144,
				},
				{
					Text:  []rune("gZi"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZi"),
					Langs: 331808,
				},
				{
					Text:  []rune("i"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 4096,
				},
				{
					Text:  []rune("j"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("gie"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
				{
					Text:  []rune("gi"),
					Langs: 128,
				},
				{
					Text:  []rune("ji"),
					Langs: 64,
				},
				{
					Text:  []rune("dZe"),
					Langs: 4096,
				},
			},
		},
		{
			Pattern: []rune("gie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
				{
					Text:  []rune("gi"),
					Langs: 128,
				},
				{
					Text:  []rune("dZe"),
					Langs: 4096,
				},
				{
					Text:  []rune("je"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: 512,
				},
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ge"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("y"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gE"),
					Langs: -1,
				},
				{
					Text:  []rune("xe"),
					Langs: 262144,
				},
				{
					Text:  []rune("Ze"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZe"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("y"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("xi"),
					Langs: 262144,
				},
				{
					Text:  []rune("Zi"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZi"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("ge"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gE"),
					Langs: -1,
				},
				{
					Text:  []rune("xe"),
					Langs: 262144,
				},
				{
					Text:  []rune("hE"),
					Langs: 131072,
				},
				{
					Text:  []rune("je"),
					Langs: 512,
				},
				{
					Text:  []rune("Ze"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZe"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gI"),
					Langs: -1,
				},
				{
					Text:  []rune("xi"),
					Langs: 262144,
				},
				{
					Text:  []rune("hI"),
					Langs: 131072,
				},
				{
					Text:  []rune("i"),
					Langs: 512,
				},
				{
					Text:  []rune("Zi"),
					Langs: 32832,
				},
				{
					Text:  []rune("dZi"),
					Langs: 331808,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
					[]rune("á"),
					[]rune("é"),
					[]rune("ó"),
					[]rune("ú"),
					[]rune("ü"),
					[]rune("ö"),
					[]rune("ő"),
					[]rune("ű"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("d"),
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("g"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("y"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("h"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: 16,
				},
				{
					Text:  []rune("ix"),
					Langs: 262144,
				},
				{
					Text:  []rune("iZ"),
					Langs: 622656,
				},
			},
		},
		{
			Pattern: []rune("j"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("e"),
					[]rune("i"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
				{
					Text:  []rune("dZ"),
					Langs: 32,
				},
				{
					Text:  []rune("x"),
					Langs: 262144,
				},
				{
					Text:  []rune("Z"),
					Langs: 622656,
				},
			},
		},
		{
			Pattern: []rune("rz"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: 16384,
				},
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("rz"),
					Langs: -1,
				},
				{
					Text:  []rune("rts"),
					Langs: 128,
				},
				{
					Text:  []rune("Z"),
					Langs: 16384,
				},
				{
					Text:  []rune("r"),
					Langs: 16384,
				},
				{
					Text:  []rune("rZ"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 160,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 131232,
				},
				{
					Text:  []rune("tS"),
					Langs: 160,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 131232,
				},
				{
					Text:  []rune("tz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Za"),
					Langs: 16384,
				},
				{
					Text:  []rune("za"),
					Langs: 16384,
				},
				{
					Text:  []rune("zja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Za"),
					Langs: 16384,
				},
				{
					Text:  []rune("zja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zią"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zom"),
					Langs: 16384,
				},
				{
					Text:  []rune("zom"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zią"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zon"),
					Langs: 16384,
				},
				{
					Text:  []rune("zon"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zię"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zem"),
					Langs: 16384,
				},
				{
					Text:  []rune("zem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zię"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zen"),
					Langs: 16384,
				},
				{
					Text:  []rune("zen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: 16384,
				},
				{
					Text:  []rune("ze"),
					Langs: 16384,
				},
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
				{
					Text:  []rune("Ze"),
					Langs: 16384,
				},
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zo"),
					Langs: 16384,
				},
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ziu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zu"),
					Langs: 16384,
				},
				{
					Text:  []rune("zju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: 16384,
				},
				{
					Text:  []rune("zi"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
				{
					Text:  []rune("dzi"),
					Langs: 4096,
				},
				{
					Text:  []rune("tsi"),
					Langs: 4096,
				},
				{
					Text:  []rune("si"),
					Langs: 262144,
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
					Langs: 128,
				},
				{
					Text:  []rune("ts"),
					Langs: 4096,
				},
				{
					Text:  []rune("S"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("v"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
				{
					Text:  []rune("dz"),
					Langs: 4096,
				},
				{
					Text:  []rune("Z"),
					Langs: 32768,
				},
			},
		},
		{
			Pattern: []rune("z"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("p"),
					[]rune("t"),
					[]rune("c"),
					[]rune("k"),
					[]rune("f"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: 4096,
				},
				{
					Text:  []rune("S"),
					Langs: 32768,
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
			Pattern: []rune("oue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oue"),
					Langs: -1,
				},
				{
					Text:  []rune("ve"),
					Langs: 64,
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
			Pattern: []rune("ae"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: 128,
				},
				{
					Text:  []rune("aje"),
					Langs: 131072,
				},
				{
					Text:  []rune("ae"),
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
			Pattern: []rune("au"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: 64,
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
			Pattern: []rune("ea"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ea"),
					Langs: -1,
				},
				{
					Text:  []rune("ja"),
					Langs: 65536,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: 32,
				},
				{
					Text:  []rune("aje"),
					Langs: 131072,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("eu"),
					Langs: -1,
				},
				{
					Text:  []rune("Yj"),
					Langs: 128,
				},
				{
					Text:  []rune("ej"),
					Langs: 128,
				},
				{
					Text:  []rune("oj"),
					Langs: 128,
				},
				{
					Text:  []rune("Y"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: 128,
				},
				{
					Text:  []rune("e"),
					Langs: 16384,
				},
				{
					Text:  []rune("ije"),
					Langs: 131072,
				},
				{
					Text:  []rune("Q"),
					Langs: 16,
				},
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("io"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("iu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iy"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("oe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: 128,
				},
				{
					Text:  []rune("oje"),
					Langs: 131072,
				},
				{
					Text:  []rune("u"),
					Langs: 16,
				},
				{
					Text:  []rune("oe"),
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
			Pattern: []rune("oo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: 32,
				},
				{
					Text:  []rune("o"),
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
					Langs: 576,
				},
				{
					Text:  []rune("au"),
					Langs: 16,
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
			Pattern: []rune("oy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
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
			Pattern: []rune("ua"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("va"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: 128,
				},
				{
					Text:  []rune("uje"),
					Langs: 131072,
				},
				{
					Text:  []rune("ve"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ui"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("uj"),
					Langs: -1,
				},
				{
					Text:  []rune("vi"),
					Langs: -1,
				},
				{
					Text:  []rune("Y"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("uu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
				{
					Text:  []rune("Q"),
					Langs: 16,
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
			Pattern: []rune("uy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("uj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ya"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
				{
					Text:  []rune("ije"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("yi"),
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
			Pattern: []rune("yi"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("yo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("yu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("ó"),
					[]rune("é"),
					[]rune("ê"),
				},
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
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("ó"),
					[]rune("é"),
					[]rune("ê"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: 131072,
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
					Text:  []rune("EE"),
					Langs: 96,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("om"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
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
			Pattern: []rune("ă"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: 65536,
				},
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ā"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("č"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ć"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: 16384,
				},
				{
					Text:  []rune("ts"),
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
				{
					Text:  []rune("tS"),
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ď"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("en"),
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
			Pattern: []rune("è"),
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
			Pattern: []rune("ě"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ē"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ģ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ğ"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
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
			Pattern: []rune("ī"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ı"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 524288,
				},
				{
					Text:  nil,
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ķ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("t"),
					Langs: 8192,
				},
				{
					Text:  []rune("tj"),
					Langs: 8192,
				},
			},
		},
		{
			Pattern: []rune("ļ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ł"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 16384,
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
					Langs: 262144,
				},
			},
		},
		{
			Pattern: []rune("ņ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 8192,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: 16384,
				},
				{
					Text:  []rune("o"),
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
			Pattern: []rune("õ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("on"),
					Langs: 32768,
				},
				{
					Text:  []rune("Y"),
					Langs: 2048,
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
			Pattern: []rune("ö"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ř"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
				{
					Text:  []rune("rZ"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ś"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: 16384,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ţ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ť"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  []rune("tj"),
					Langs: 8,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: 294912,
				},
			},
		},
		{
			Pattern: []rune("ū"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("ů"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ù"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ý"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ż"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: 16384,
				},
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ž"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ß"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
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
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("ć"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("l"),
					[]rune("ł"),
					[]rune("m"),
					[]rune("n"),
					[]rune("ń"),
					[]rune("r"),
					[]rune("s"),
					[]rune("ś"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ź"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("O"),
					Langs: -1,
				},
				{
					Text:  []rune("P"),
					Langs: 16384,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("A"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("B"),
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
				{
					Text:  []rune("ts"),
					Langs: 16392,
				},
				{
					Text:  []rune("dZ"),
					Langs: 524288,
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
					Text:  []rune("E"),
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
				{
					Text:  []rune("x"),
					Langs: 65536,
				},
				{
					Text:  []rune("H"),
					Langs: 299072,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: 262144,
				},
				{
					Text:  []rune("Z"),
					Langs: 622656,
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
					Text:  []rune("O"),
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
					Langs: 32768,
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
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("V"),
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
				{
					Text:  []rune("w"),
					Langs: 48,
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
					Langs: 294912,
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
				{
					Text:  []rune("ts"),
					Langs: 128,
				},
				{
					Text:  []rune("dz"),
					Langs: 4096,
				},
				{
					Text:  []rune("ts"),
					Langs: 4096,
				},
				{
					Text:  []rune("s"),
					Langs: 262144,
				},
			},
		},
	},
	common.Lang(Arabic): {
		{
			Pattern: []rune("ا"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ب"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ب"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ت"),
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
			Pattern: []rune("ت"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ث"),
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
			Pattern: []rune("ث"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ج"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ج"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ1"),
					Langs: -1,
				},
				{
					Text:  []rune("Z1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ح"),
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
			Pattern: []rune("ح"),
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
			Pattern: []rune("ح"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h1"),
					Langs: -1,
				},
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("خ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("خ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("د"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("د"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ذ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ذ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ر"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ر"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ز"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ز"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("س"),
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
			Pattern: []rune("س"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ش"),
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
			Pattern: []rune("ش"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ص"),
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
			Pattern: []rune("ص"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ض"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ض"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ط"),
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
			Pattern: []rune("ط"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ظ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ظ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ع"),
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
			Pattern: []rune("ع"),
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
			Pattern: []rune("ع"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h1"),
					Langs: -1,
				},
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("غ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("غ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ف"),
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
			Pattern: []rune("ف"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ق"),
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
			Pattern: []rune("ق"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ك"),
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
			Pattern: []rune("ك"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ل"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ل"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("م"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("م"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ن"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ن"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ه"),
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
			Pattern: []rune("ه"),
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
			Pattern: []rune("ه"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h1"),
					Langs: -1,
				},
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("و"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("و"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
				{
					Text:  []rune("v1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ي\u200e"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
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
			Pattern: []rune("ي\u200e"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("j1"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Cyrillic): {
		{
			Pattern: []rune("ця"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цю"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("циа"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цие"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цио"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("циу"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("сие"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("сио"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("зие"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("зио"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("с"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("с"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гауз"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("haus"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гаус"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("haus"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гольц"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("holts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("геймер"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hejmer"),
					Langs: -1,
				},
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гейм"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hejm"),
					Langs: -1,
				},
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гоф"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hof"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гер"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ger"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ген"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гин"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gin"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("г"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("й"),
					[]rune("ё"),
					[]rune("я"),
					[]rune("ю"),
					[]rune("ы"),
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("и"),
					[]rune("у"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("и"),
					[]rune("у"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("г"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("и"),
					[]rune("у"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ля"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лю"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лё"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("le"),
					Langs: -1,
				},
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лио"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("le"),
					Langs: -1,
				},
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ле"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lE"),
					Langs: -1,
				},
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ийе"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ие"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ыйе"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ые"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ий"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("о"),
					[]rune("у"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ый"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("о"),
					[]rune("у"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ий"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("ый"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("ей"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jej"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("е"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("у"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("е"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("эй"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ей"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ауе"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ауэ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("а"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("б"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("в"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("г"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("д"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("е"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ё"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ж"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("з"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("и"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("й"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("к"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("л"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("м"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("н"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("о"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("п"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("р"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("с"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("т"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("у"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ф"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("х"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ц"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ч"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ш"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("щ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("StS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ъ"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ы"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ь"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("э"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ю"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("я"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Czech): {
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
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
				{
					Text:  []rune("kv"),
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
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("č"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ž"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ň"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ť"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  []rune("tj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ď"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ř"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
				{
					Text:  []rune("rZ"),
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
			Pattern: []rune("ý"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ě"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ů"),
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
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
				{
					Text:  []rune("kv"),
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Dutch): {
		{
			Pattern: []rune("ssj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("pf"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("pf"),
					Langs: -1,
				},
				{
					Text:  []rune("p"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ph"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("th"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
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
			Pattern: []rune("th"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
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
			Pattern: []rune("ou"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
				{
					Text:  []rune("Yj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aa"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ui"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
				{
					Text:  []rune("uj"),
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
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
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
				{
					Text:  []rune("x"),
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
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
				{
					Text:  []rune("Q"),
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
					Text:  []rune("w"),
					Langs: -1,
				},
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(English): {
		{
			Pattern: []rune("�"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mc"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("mak"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
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
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cc"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("c"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gh"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
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
					Text:  []rune("f"),
					Langs: -1,
				},
				{
					Text:  []rune("w"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gn"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
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
			Pattern: []rune("th"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
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
			Pattern: []rune("ph"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("sk"),
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
			Pattern: []rune("who"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("wh"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("w"),
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
			Pattern: []rune("h"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]"),
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("H"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kn"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mb"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ng"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("N"),
					Langs: -1,
				},
				{
					Text:  []rune("ng"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("pn"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("pn"),
					Langs: -1,
				},
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ps"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ps"),
					Langs: -1,
				},
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kw"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("So"),
					Langs: -1,
				},
				{
					Text:  []rune("Sa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("So"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("wr"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
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
				MatchEmptyString: true,
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
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
			Pattern: []rune("aue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
				{
					Text:  []rune("oue"),
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
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
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
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
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
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
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
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ear"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ia"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
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
					Text:  nil,
					Langs: -1,
				},
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oa"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ou"),
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
			Pattern: []rune("oo"),
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
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ju"),
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
				Prefix: [][]rune{
					[]rune("r"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
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
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
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
					Text:  []rune("a"),
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
				{
					Text:  []rune("a"),
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
					Text:  []rune("w"),
					Langs: -1,
				},
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(French): {
		{
			Pattern: []rune("lt"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lt"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("n"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
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
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("n"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("e"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  nil,
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
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ds"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ds"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ps"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ps"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rs"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("e"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("rs"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ts"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  nil,
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
					[]rune("é"),
					[]rune("è"),
					[]rune("ê"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aill"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ll"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("j"),
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
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
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
			RightContext: &common.Matcher{
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
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
			Pattern: []rune("ou"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
				},
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
			Pattern: []rune("au"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
				{
					Text:  []rune("au"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
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
					Text:  []rune("e"),
					Langs: -1,
				},
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
				{
					Text:  []rune("va"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("e"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
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
				{
					Text:  []rune("Q"),
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
	common.Lang(German): {
		{
			Pattern: []rune("ewitsch"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owitsch"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("evitsch"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ovitsch"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("witsch"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("vitsch"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("xS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ziu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chs"),
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
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sp"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("st"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("St"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssp"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: -1,
				},
				{
					Text:  []rune("sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sp"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: -1,
				},
				{
					Text:  []rune("sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sst"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("St"),
					Langs: -1,
				},
				{
					Text:  []rune("st"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("st"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("St"),
					Langs: -1,
				},
				{
					Text:  []rune("st"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("pf"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("pf"),
					Langs: -1,
				},
				{
					Text:  []rune("p"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ph"),
					Langs: -1,
				},
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewitz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evits"),
					Langs: -1,
				},
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewiz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evits"),
					Langs: -1,
				},
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("evitz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evits"),
					Langs: -1,
				},
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eviz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evits"),
					Langs: -1,
				},
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owitz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovits"),
					Langs: -1,
				},
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owiz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovits"),
					Langs: -1,
				},
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ovitz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovits"),
					Langs: -1,
				},
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oviz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovits"),
					Langs: -1,
				},
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("witz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vits"),
					Langs: -1,
				},
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("wiz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vits"),
					Langs: -1,
				},
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("vitz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vits"),
					Langs: -1,
				},
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("viz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vits"),
					Langs: -1,
				},
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("thal"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tal"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("th"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
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
			Pattern: []rune("th"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rh"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("H"),
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
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
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
			},
		},
		{
			Pattern: []rune("s"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
					[]rune("j"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ß"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("aue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ae"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
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
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Yj"),
					Langs: -1,
				},
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("e"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
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
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("a"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("A"),
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
					Text:  []rune("O"),
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
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
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
			},
		},
	},
	common.Lang(Greek): {
		{
			Pattern: []rune("αυ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("αυ"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("κ"),
					[]rune("π"),
					[]rune("σ"),
					[]rune("τ"),
					[]rune("φ"),
					[]rune("θ"),
					[]rune("χ"),
					[]rune("ψ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("αυ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("av"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ευ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ευ"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("κ"),
					[]rune("π"),
					[]rune("σ"),
					[]rune("τ"),
					[]rune("φ"),
					[]rune("θ"),
					[]rune("χ"),
					[]rune("ψ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ευ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ev"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηυ"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("if"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηυ"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("κ"),
					[]rune("π"),
					[]rune("σ"),
					[]rune("τ"),
					[]rune("φ"),
					[]rune("θ"),
					[]rune("χ"),
					[]rune("ψ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("if"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηυ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("iv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ου"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("αι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ει"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("οι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ωι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("υι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γγ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γγ"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γγ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ng"),
					Langs: -1,
				},
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γγ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ng"),
					Langs: -1,
				},
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γι"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γε"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γε"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("κζ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("τζ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("σ"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("β"),
					[]rune("γ"),
					[]rune("δ"),
					[]rune("μ"),
					[]rune("ν"),
					[]rune("ρ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μβ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("mb"),
					Langs: -1,
				},
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μπ"),
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
		{
			Pattern: []rune("μπ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("mb"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μπ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ντ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ντ"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nd"),
					Langs: -1,
				},
				{
					Text:  []rune("nt"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ντ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nt"),
					Langs: -1,
				},
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ά"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("έ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ή"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ί"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ό"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ύ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
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
			Pattern: []rune("ώ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ΰ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
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
			Pattern: []rune("ϋ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
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
			Pattern: []rune("ϊ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("α"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("β"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("δ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ε"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ζ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("η"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ι"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("κ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("λ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ν"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ξ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ο"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("π"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ρ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("σ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ς"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("τ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("υ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
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
			Pattern: []rune("φ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("θ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("χ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ψ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ps"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ω"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Greeklatin): {
		{
			Pattern: []rune("au"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("f"),
					[]rune("h"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("av"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("f"),
					[]rune("h"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ev"),
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
			Pattern: []rune("gge"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nje"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ni"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gge"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ng"),
					Langs: -1,
				},
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gk"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gke"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nje"),
					Langs: -1,
				},
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gki"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ni"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gke"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gki"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gk"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ng"),
					Langs: -1,
				},
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gk"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ngi"),
					Langs: -1,
				},
				{
					Text:  []rune("Ni"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghe"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Nje"),
					Langs: -1,
				},
				{
					Text:  []rune("Nge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ghi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ghi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ghe"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ghe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ngh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ng"),
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
			},
		},
		{
			Pattern: []rune("ngi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ngi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ngi"),
					Langs: -1,
				},
				{
					Text:  []rune("Ni"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nge"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nge"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Nje"),
					Langs: -1,
				},
				{
					Text:  []rune("Nge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ge"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ge"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ng"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ng"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
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
			Pattern: []rune("dh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("dj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
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
			Pattern: []rune("th"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("g"),
					[]rune("d"),
					[]rune("m"),
					[]rune("n"),
					[]rune("r"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mb"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("mb"),
					Langs: -1,
				},
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mp"),
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
		{
			Pattern: []rune("mp"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("mp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mp"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nt"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nt"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nd"),
					Langs: -1,
				},
				{
					Text:  []rune("nt"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nt"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("nt"),
					Langs: -1,
				},
				{
					Text:  []rune("d"),
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
			Pattern: []rune("óu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("ý"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
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
					Text:  []rune("x"),
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
					Text:  []rune("j"),
					Langs: -1,
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
			Pattern: []rune("ο"),
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
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
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
					Text:  []rune("TB"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Hungarian): {
		{
			Pattern: []rune("sz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zs"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cs"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
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
					Text:  []rune("aj"),
					Langs: -1,
				},
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
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("o"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("o"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ely"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
				{
					Text:  []rune("eli"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ly"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
				{
					Text:  []rune("li"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
					[]rune("á"),
					[]rune("é"),
					[]rune("ó"),
					[]rune("ú"),
					[]rune("ü"),
					[]rune("ö"),
					[]rune("ő"),
					[]rune("ű"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ny"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
					[]rune("á"),
					[]rune("é"),
					[]rune("ó"),
					[]rune("ú"),
					[]rune("ü"),
					[]rune("ö"),
					[]rune("ő"),
					[]rune("ű"),
				},
			},
			Phonetic: []common.RuleToken{
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
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("ni"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ty"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
					[]rune("á"),
					[]rune("é"),
					[]rune("ó"),
					[]rune("ú"),
					[]rune("ü"),
					[]rune("ö"),
					[]rune("ő"),
					[]rune("ű"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ty"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  []rune("ti"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ku"),
					Langs: -1,
				},
				{
					Text:  []rune("kv"),
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
			Pattern: []rune("ö"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
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
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
					Text:  []rune("z"),
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Suffix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
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
			Pattern: []rune("ci"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
				},
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
	common.Lang(Latvian): {
		{
			Pattern: []rune("č"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ģ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ķ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
				{
					Text:  []rune("tj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ļ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ņ"),
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
			Pattern: []rune("ž"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ā"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ē"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ī"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ū"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("io"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
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
			Pattern: []rune("ui"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("uj"),
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
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
	common.Lang(Polish): {
		{
			Pattern: []rune("ska"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ski"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cka"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tski"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lowa"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
				{
					Text:  []rune("lof"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("el"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kowa"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kova"),
					Langs: -1,
				},
				{
					Text:  []rune("kof"),
					Langs: -1,
				},
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owa"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
				{
					Text:  []rune("of"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lowna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lovna"),
					Langs: -1,
				},
				{
					Text:  []rune("levna"),
					Langs: -1,
				},
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("el"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kowna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("kovna"),
					Langs: -1,
				},
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovna"),
					Langs: -1,
				},
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lówna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("el"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kówna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
				{
					Text:  []rune("ek"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ówna"),
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
			Pattern: []rune("czy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cze"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: -1,
				},
				{
					Text:  []rune("tSF"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ciewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsevitS"),
					Langs: -1,
				},
				{
					Text:  []rune("tSevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("siewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("sevitS"),
					Langs: -1,
				},
				{
					Text:  []rune("SevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ziewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("zevitS"),
					Langs: -1,
				},
				{
					Text:  []rune("ZevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("riewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("rjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("diewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("djevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tiewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owicz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("icz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("itS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cia"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSB"),
					Langs: -1,
				},
				{
					Text:  []rune("tsB"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSa"),
					Langs: -1,
				},
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cią"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSom"),
					Langs: -1,
				},
				{
					Text:  []rune("tsom"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cią"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSon"),
					Langs: -1,
				},
				{
					Text:  []rune("tson"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cię"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSem"),
					Langs: -1,
				},
				{
					Text:  []rune("tsem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cię"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSen"),
					Langs: -1,
				},
				{
					Text:  []rune("tsen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cie"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSF"),
					Langs: -1,
				},
				{
					Text:  []rune("tsF"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: -1,
				},
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSo"),
					Langs: -1,
				},
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ciu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSu"),
					Langs: -1,
				},
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSi"),
					Langs: -1,
				},
				{
					Text:  []rune("tsI"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ć"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sia"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("SB"),
					Langs: -1,
				},
				{
					Text:  []rune("sB"),
					Langs: -1,
				},
				{
					Text:  []rune("sja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sa"),
					Langs: -1,
				},
				{
					Text:  []rune("sja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sią"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Som"),
					Langs: -1,
				},
				{
					Text:  []rune("som"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sią"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Son"),
					Langs: -1,
				},
				{
					Text:  []rune("son"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("się"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sem"),
					Langs: -1,
				},
				{
					Text:  []rune("sem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("się"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sen"),
					Langs: -1,
				},
				{
					Text:  []rune("sen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("SF"),
					Langs: -1,
				},
				{
					Text:  []rune("sF"),
					Langs: -1,
				},
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Se"),
					Langs: -1,
				},
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("So"),
					Langs: -1,
				},
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("siu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Su"),
					Langs: -1,
				},
				{
					Text:  []rune("sju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("si"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Si"),
					Langs: -1,
				},
				{
					Text:  []rune("sI"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ś"),
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
			Pattern: []rune("zia"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ZB"),
					Langs: -1,
				},
				{
					Text:  []rune("zB"),
					Langs: -1,
				},
				{
					Text:  []rune("zja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Za"),
					Langs: -1,
				},
				{
					Text:  []rune("zja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zią"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zom"),
					Langs: -1,
				},
				{
					Text:  []rune("zom"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zią"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zon"),
					Langs: -1,
				},
				{
					Text:  []rune("zon"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zię"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zem"),
					Langs: -1,
				},
				{
					Text:  []rune("zem"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zię"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zen"),
					Langs: -1,
				},
				{
					Text:  []rune("zen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ZF"),
					Langs: -1,
				},
				{
					Text:  []rune("zF"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zo"),
					Langs: -1,
				},
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ziu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zu"),
					Langs: -1,
				},
				{
					Text:  []rune("zju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: -1,
				},
				{
					Text:  []rune("zI"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("że"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
				{
					Text:  []rune("ZF"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("że"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
				{
					Text:  []rune("ZF"),
					Langs: -1,
				},
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
				{
					Text:  []rune("zF"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("że"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("źe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ży"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("źi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ż"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
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
			Pattern: []rune("rze"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Se"),
					Langs: -1,
				},
				{
					Text:  []rune("re"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rze"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
				{
					Text:  []rune("re"),
					Langs: -1,
				},
				{
					Text:  []rune("rZe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rzy"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Si"),
					Langs: -1,
				},
				{
					Text:  []rune("ri"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rzy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: -1,
				},
				{
					Text:  []rune("ri"),
					Langs: -1,
				},
				{
					Text:  []rune("rZi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rz"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
				{
					Text:  []rune("r"),
					Langs: -1,
				},
				{
					Text:  []rune("rZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ł"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
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
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("s"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("om"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("en"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ije"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yje"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yj"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iy"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("rje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("die"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("F"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
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
			Pattern: []rune("au"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ej"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
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
			Pattern: []rune("aj"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("B"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("p"),
					[]rune("s"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
				{
					Text:  []rune("F"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("c"),
					[]rune("ć"),
					[]rune("d"),
					[]rune("g"),
					[]rune("k"),
					[]rune("l"),
					[]rune("ł"),
					[]rune("m"),
					[]rune("n"),
					[]rune("ń"),
					[]rune("r"),
					[]rune("s"),
					[]rune("ś"),
					[]rune("t"),
					[]rune("w"),
					[]rune("z"),
					[]rune("ź"),
					[]rune("ż"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("P"),
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
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
					Text:  []rune("I"),
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
			RightContext: &common.Matcher{
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
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("v"),
				},
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
				Prefix: [][]rune{
					[]rune("p"),
					[]rune("t"),
					[]rune("c"),
					[]rune("k"),
					[]rune("f"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
				},
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
				Suffix: [][]rune{
					[]rune("b"),
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
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
			Pattern: []rune("ex"),
			RightContext: &common.Matcher{
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
				Prefix: [][]rune{
					[]rune("c"),
					[]rune("s"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("í"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
	common.Lang(Romanian): {
		{
			Pattern: []rune("ce"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSi"),
					Langs: -1,
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZi"),
					Langs: -1,
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("dZ"),
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
			},
		},
		{
			Pattern: []rune("i"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ţ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
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
			Pattern: []rune("î"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ă"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
				{
					Text:  []rune("a"),
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
					Text:  []rune("E"),
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
					Text:  []rune("x"),
					Langs: -1,
				},
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
					Text:  []rune("I"),
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	common.Lang(Russian): {
		{
			Pattern: []rune("yna"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("in"),
					Langs: -1,
				},
				{
					Text:  []rune("ina"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ina"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("in"),
					Langs: -1,
				},
				{
					Text:  []rune("ina"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("liova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lof"),
					Langs: -1,
				},
				{
					Text:  []rune("lef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lof"),
					Langs: -1,
				},
				{
					Text:  []rune("lef"),
					Langs: -1,
				},
				{
					Text:  []rune("lova"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ova"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("of"),
					Langs: -1,
				},
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eva"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
				{
					Text:  []rune("ova"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aia"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aja"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aya"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aja"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsya"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsyu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsyo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsiu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("syo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zyo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ger"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ger"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gen"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gin"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gin"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("j"),
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("i"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("i"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("i"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
				{
					Text:  []rune("h"),
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
			Pattern: []rune("ch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("StS"),
					Langs: -1,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
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
			Pattern: []rune("zh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
				{
					Text:  []rune("tz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
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
			Pattern: []rune("s"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("s"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lya"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lyu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lia"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("liu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lja"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lju"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("le"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("lE"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lyo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("lo"),
					Langs: -1,
				},
				{
					Text:  []rune("le"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ije"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yje"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yye"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yie"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iy"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yj"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("io"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("u"),
				},
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
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("iy"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("yy"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("yi"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("yj"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("ij"),
			RightContext: &common.Matcher{
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
			Pattern: []rune("e"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aje"),
					Langs: -1,
				},
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("oo"),
					Langs: -1,
				},
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
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
					Text:  []rune("I"),
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
				},
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
				Suffix: [][]rune{
					[]rune("d"),
					[]rune("g"),
					[]rune("t"),
				},
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
			Pattern: []rune("m"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
					[]rune("v"),
					[]rune("f"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
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
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("v"),
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
			Pattern: []rune("b"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("B"),
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
					Text:  []rune("V"),
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
	},
	common.Lang(Turkish): {
		{
			Pattern: []rune("ç"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ğ"),
			Phonetic: []common.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ı"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
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
					Text:  []rune("dZ"),
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
					Text:  []rune("j"),
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
}

var LangRules = []common.LangRule{
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("o’"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("o'"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("mc"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("fitz"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ceau"),
			},
		},
		Langs:  65600,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eau"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ault"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("oult"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("eux"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("eix"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("glou"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("uu"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tx"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("witz"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("tz"),
			},
		},
		Langs:  131232,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("tz"),
			},
		},
		Langs:  131104,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("poulos"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("pulos"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("iou"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("sj"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("sj"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("güe"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("güi"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ghe"),
			},
		},
		Langs:  66048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ghi"),
			},
		},
		Langs:  66048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("escu"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("esco"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("vici"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("schi"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ii"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("iy"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yy"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yi"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("rz"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("rz"),
			},
		},
		Langs:  16512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("brz"),
				[]rune("crz"),
				[]rune("drz"),
				[]rune("frz"),
				[]rune("grz"),
				[]rune("krz"),
				[]rune("lrz"),
				[]rune("mrz"),
				[]rune("nrz"),
				[]rune("prz"),
				[]rune("srz"),
				[]rune("trz"),
				[]rune("wrz"),
				[]rune("zrz"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("rzb"),
				[]rune("rzc"),
				[]rune("rzd"),
				[]rune("rzf"),
				[]rune("rzg"),
				[]rune("rzh"),
				[]rune("rzk"),
				[]rune("rzl"),
				[]rune("rzm"),
				[]rune("rzn"),
				[]rune("rzp"),
				[]rune("rzs"),
				[]rune("rzt"),
				[]rune("rzw"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cki"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ska"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cka"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ae"),
			},
		},
		Langs:  131232,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oe"),
			},
		},
		Langs:  131312,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("th"),
			},
		},
		Langs:  160,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("th"),
			},
		},
		Langs:  672,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("mann"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cz"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cy"),
			},
		},
		Langs:  16896,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("niew"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("etti"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("eti"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ati"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ato"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ano"),
				[]rune("ono"),
				[]rune("eno"),
				[]rune("ino"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ani"),
				[]rune("oni"),
				[]rune("eni"),
				[]rune("ini"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("esi"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("oli"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("field"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("stein"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("heim"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("heimer"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("thal"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("zweig"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ah"),
				[]rune("eh"),
				[]rune("oh"),
				[]rune("uh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("äh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("öh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("üh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("lha"),
				[]rune("lho"),
				[]rune("nha"),
				[]rune("nho"),
			},
		},
		Langs:  32768,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("lha"),
				[]rune("lho"),
				[]rune("lhu"),
				[]rune("nha"),
				[]rune("nho"),
				[]rune("nhu"),
			},
		},
		Langs:  819416,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("chsch"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tsch"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  131200,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  131200,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ck"),
			},
		},
		Langs:  160,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("c"),
			},
		},
		Langs:  608264,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sz"),
			},
		},
		Langs:  18432,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("cs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dzs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("zs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("zs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("wl"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("wr"),
			},
		},
		Langs:  16560,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gya"),
				[]rune("gye"),
				[]rune("gyo"),
				[]rune("gyu"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  133696,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("guy"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gue"),
				[]rune("gui"),
			},
		},
		Langs:  294976,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gua"),
				[]rune("guo"),
			},
		},
		Langs:  294912,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gia"),
				[]rune("gio"),
				[]rune("giu"),
			},
		},
		Langs:  4608,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ly"),
			},
		},
		Langs:  150016,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ny"),
			},
		},
		Langs:  412160,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ty"),
			},
		},
		Langs:  150016,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ā"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ć"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ç"),
			},
		},
		Langs:  819264,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("č"),
			},
		},
		Langs:  8200,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ď"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ē"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ğ"),
			},
		},
		Langs:  524288,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ģ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ī"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ķ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ļ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ł"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ņ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ń"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ñ"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ň"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ř"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ś"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ş"),
			},
		},
		Langs:  589824,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("š"),
			},
		},
		Langs:  8200,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ţ"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ť"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ź"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ž"),
			},
		},
		Langs:  8200,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ż"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ß"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ä"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("á"),
			},
		},
		Langs:  297480,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("â"),
			},
		},
		Langs:  98368,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ă"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ą"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("à"),
			},
		},
		Langs:  32768,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ã"),
			},
		},
		Langs:  32768,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ę"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("é"),
			},
		},
		Langs:  2632,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("è"),
			},
		},
		Langs:  266304,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ê"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ě"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ê"),
			},
		},
		Langs:  32832,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("í"),
			},
		},
		Langs:  297480,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("î"),
			},
		},
		Langs:  65600,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ı"),
			},
		},
		Langs:  524288,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ó"),
			},
		},
		Langs:  317960,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ö"),
			},
		},
		Langs:  526464,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ô"),
			},
		},
		Langs:  32832,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("õ"),
			},
		},
		Langs:  34816,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ò"),
			},
		},
		Langs:  266240,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ű"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ú"),
			},
		},
		Langs:  297480,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ü"),
			},
		},
		Langs:  821376,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ù"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ů"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ý"),
			},
		},
		Langs:  520,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("а"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ё"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("о"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("е"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("и"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("у"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ы"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("э"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ю"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("я"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("α"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ε"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("η"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ι"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ο"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("υ"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ω"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ا"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ب"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ت"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ث"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ج"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ح"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("خ'"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("د"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ذ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ر"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ز"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("س"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ش"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ص"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ض"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ط"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ظ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ع"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("غ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ف"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ق"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ك"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ل"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("م"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ن"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ه"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("و"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ي"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("آ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("إ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("أ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ؤ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ئ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("א"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ב"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ג"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ד"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ה"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ו"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ז"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ח"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ט"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("י"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("כ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ל"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("מ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("נ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ס"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ע"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("פ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("צ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ק"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ר"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ש"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ת"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("a"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("o"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("e"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("i"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("y"),
			},
		},
		Langs:  75030,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("u"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("j"),
			},
		},
		Langs:  4096,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("j[^aoeiuy]"),
		},
		Langs:  295488,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("g"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("k"),
			},
		},
		Langs:  364608,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("q"),
			},
		},
		Langs:  748056,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("v"),
			},
		},
		Langs:  16384,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("w"),
			},
		},
		Langs:  993864,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("x"),
			},
		},
		Langs:  534552,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dj"),
			},
		},
		Langs:  786432,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("v[^aoeiu]"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("c[^aohk]"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dzi"),
			},
		},
		Langs:  524512,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ou"),
			},
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ae"),
				[]rune("ai"),
				[]rune("ao"),
				[]rune("au"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("öe"),
				[]rune("öa"),
				[]rune("öi"),
				[]rune("öo"),
				[]rune("öu"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("üe"),
				[]rune("üa"),
				[]rune("üi"),
				[]rune("üo"),
				[]rune("üu"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ea"),
				[]rune("ei"),
				[]rune("eo"),
				[]rune("eu"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ia"),
				[]rune("ie"),
				[]rune("io"),
				[]rune("iu"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oa"),
				[]rune("oi"),
				[]rune("oe"),
				[]rune("ou"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ua"),
				[]rune("ui"),
				[]rune("ue"),
				[]rune("uo"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("aj"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ej"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oj"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("uj"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eu"),
			},
		},
		Langs:  147456,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ky"),
			},
		},
		Langs:  16384,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("kie"),
			},
		},
		Langs:  262720,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gie"),
			},
		},
		Langs:  360960,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cha"),
				[]rune("cho"),
				[]rune("chu"),
			},
		},
		Langs:  4096,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ch"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("son"),
			},
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sce"),
				[]rune("sci"),
			},
		},
		Langs:  64,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  280640,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("h"),
			},
		},
		Langs:  131072,
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
					Prefix: [][]rune{
						[]rune("f"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("b"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("f"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("k"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("g"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("d"),
					},
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
					Prefix: [][]rune{
						[]rune("dZ"),
					},
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
					Prefix: [][]rune{
						[]rune("tS"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jnm"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("jm"),
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
				Pattern: []rune("jI"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("I"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("a"),
						[]rune("A"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("A"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("A"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
					},
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
					Prefix: [][]rune{
						[]rune("b"),
					},
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
					Prefix: [][]rune{
						[]rune("d"),
					},
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
					Prefix: [][]rune{
						[]rune("f"),
					},
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
					Prefix: [][]rune{
						[]rune("g"),
					},
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("j"),
					},
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
					Prefix: [][]rune{
						[]rune("k"),
					},
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
					Prefix: [][]rune{
						[]rune("l"),
					},
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
					Prefix: [][]rune{
						[]rune("m"),
					},
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
					Prefix: [][]rune{
						[]rune("n"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
					},
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
					Prefix: [][]rune{
						[]rune("r"),
					},
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
					Prefix: [][]rune{
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
					},
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
					Prefix: [][]rune{
						[]rune("z"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("vanden"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("vanden"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("vander"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("vander"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("van"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("p"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("vam"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 16,
					},
				},
			},
			{
				Pattern: []rune("van"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("van"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: 16,
					},
				},
			},
			{
				Pattern: []rune("n"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("p"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("m"),
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
				Pattern: []rune("H"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("x"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sen"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("zn"),
						Langs: -1,
					},
					{
						Text:  []rune("zon"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sen"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("sn"),
						Langs: -1,
					},
					{
						Text:  []rune("son"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sEn"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("zn"),
						Langs: -1,
					},
					{
						Text:  []rune("zon"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("sEn"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("sn"),
						Langs: -1,
					},
					{
						Text:  []rune("son"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("e"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("B"),
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
						[]rune("n"),
					},
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
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("B"),
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
						[]rune("n"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("E"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("B"),
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
						[]rune("n"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("I"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("B"),
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
						[]rune("n"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Q"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("B"),
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
						[]rune("n"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Y"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("B"),
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
						[]rune("n"),
					},
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
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("lb"),
						[]rune("ld"),
						[]rune("lf"),
						[]rune("lg"),
						[]rune("lk"),
						[]rune("ll"),
						[]rune("lm"),
						[]rune("ln"),
						[]rune("lp"),
						[]rune("lr"),
						[]rune("ls"),
						[]rune("lS"),
						[]rune("lt"),
						[]rune("lv"),
						[]rune("lz"),
						[]rune("lZ"),
						[]rune("nb"),
						[]rune("nd"),
						[]rune("nf"),
						[]rune("ng"),
						[]rune("nk"),
						[]rune("nl"),
						[]rune("nm"),
						[]rune("nn"),
						[]rune("np"),
						[]rune("nr"),
						[]rune("ns"),
						[]rune("nS"),
						[]rune("nt"),
						[]rune("nv"),
						[]rune("nz"),
						[]rune("nZ"),
					},
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
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("lb"),
						[]rune("ld"),
						[]rune("lf"),
						[]rune("lg"),
						[]rune("lk"),
						[]rune("ll"),
						[]rune("lm"),
						[]rune("ln"),
						[]rune("lp"),
						[]rune("lr"),
						[]rune("ls"),
						[]rune("lS"),
						[]rune("lt"),
						[]rune("lv"),
						[]rune("lz"),
						[]rune("lZ"),
						[]rune("nb"),
						[]rune("nd"),
						[]rune("nf"),
						[]rune("ng"),
						[]rune("nk"),
						[]rune("nl"),
						[]rune("nm"),
						[]rune("nn"),
						[]rune("np"),
						[]rune("nr"),
						[]rune("ns"),
						[]rune("nS"),
						[]rune("nt"),
						[]rune("nv"),
						[]rune("nz"),
						[]rune("nZ"),
					},
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
				Pattern: []rune("E"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("lb"),
						[]rune("ld"),
						[]rune("lf"),
						[]rune("lg"),
						[]rune("lk"),
						[]rune("ll"),
						[]rune("lm"),
						[]rune("ln"),
						[]rune("lp"),
						[]rune("lr"),
						[]rune("ls"),
						[]rune("lS"),
						[]rune("lt"),
						[]rune("lv"),
						[]rune("lz"),
						[]rune("lZ"),
						[]rune("nb"),
						[]rune("nd"),
						[]rune("nf"),
						[]rune("ng"),
						[]rune("nk"),
						[]rune("nl"),
						[]rune("nm"),
						[]rune("nn"),
						[]rune("np"),
						[]rune("nr"),
						[]rune("ns"),
						[]rune("nS"),
						[]rune("nt"),
						[]rune("nv"),
						[]rune("nz"),
						[]rune("nZ"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("E"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("I"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("lb"),
						[]rune("ld"),
						[]rune("lf"),
						[]rune("lg"),
						[]rune("lk"),
						[]rune("ll"),
						[]rune("lm"),
						[]rune("ln"),
						[]rune("lp"),
						[]rune("lr"),
						[]rune("ls"),
						[]rune("lS"),
						[]rune("lt"),
						[]rune("lv"),
						[]rune("lz"),
						[]rune("lZ"),
						[]rune("nb"),
						[]rune("nd"),
						[]rune("nf"),
						[]rune("ng"),
						[]rune("nk"),
						[]rune("nl"),
						[]rune("nm"),
						[]rune("nn"),
						[]rune("np"),
						[]rune("nr"),
						[]rune("ns"),
						[]rune("nS"),
						[]rune("nt"),
						[]rune("nv"),
						[]rune("nz"),
						[]rune("nZ"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("I"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Q"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("lb"),
						[]rune("ld"),
						[]rune("lf"),
						[]rune("lg"),
						[]rune("lk"),
						[]rune("ll"),
						[]rune("lm"),
						[]rune("ln"),
						[]rune("lp"),
						[]rune("lr"),
						[]rune("ls"),
						[]rune("lS"),
						[]rune("lt"),
						[]rune("lv"),
						[]rune("lz"),
						[]rune("lZ"),
						[]rune("nb"),
						[]rune("nd"),
						[]rune("nf"),
						[]rune("ng"),
						[]rune("nk"),
						[]rune("nl"),
						[]rune("nm"),
						[]rune("nn"),
						[]rune("np"),
						[]rune("nr"),
						[]rune("ns"),
						[]rune("nS"),
						[]rune("nt"),
						[]rune("nv"),
						[]rune("nz"),
						[]rune("nZ"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("Q"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Y"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("l"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("s"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("lb"),
						[]rune("ld"),
						[]rune("lf"),
						[]rune("lg"),
						[]rune("lk"),
						[]rune("ll"),
						[]rune("lm"),
						[]rune("ln"),
						[]rune("lp"),
						[]rune("lr"),
						[]rune("ls"),
						[]rune("lS"),
						[]rune("lt"),
						[]rune("lv"),
						[]rune("lz"),
						[]rune("lZ"),
						[]rune("nb"),
						[]rune("nd"),
						[]rune("nf"),
						[]rune("ng"),
						[]rune("nk"),
						[]rune("nl"),
						[]rune("nm"),
						[]rune("nn"),
						[]rune("np"),
						[]rune("nr"),
						[]rune("ns"),
						[]rune("nS"),
						[]rune("nt"),
						[]rune("nv"),
						[]rune("nz"),
						[]rune("nZ"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("Y"),
						Langs: -1,
					},
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEs"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lEs"),
						Langs: -1,
					},
					{
						Text:  []rune("lz"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lE"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("b"),
						[]rune("d"),
						[]rune("f"),
						[]rune("g"),
						[]rune("k"),
						[]rune("m"),
						[]rune("n"),
						[]rune("p"),
						[]rune("r"),
						[]rune("S"),
						[]rune("t"),
						[]rune("v"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lE"),
						Langs: -1,
					},
					{
						Text:  []rune("l"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aue"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oue"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AvE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("AvE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ave"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("Ave"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("avE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("avE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ave"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ave"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OvE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("OvE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ove"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("Ove"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ovE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ovE"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ove"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ove"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ea"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("ea"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("EA"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ea"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("Ea"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
					{
						Text:  []rune("eA"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uje"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("iji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ije"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Iji"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjI"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ije"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjE"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ajo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ojo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ejo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uja"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ujo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ija"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ija"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjA"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ijo"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjO"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("D"),
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
				Pattern: []rune("lYndEr"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lander"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lAndEr"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lAnder"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("landEr"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lender"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEndEr"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lendEr"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEnder"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("burk"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bUrk"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("burg"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("bUrg"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Burk"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("BUrk"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Burg"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("BUrg"),
				RightContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("burk"),
						Langs: -1,
					},
					{
						Text:  []rune("berk"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("s"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
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
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
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
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
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
		},
		Second: map[common.Lang]common.Rules{
			common.Lang(Any): {
				{
					Pattern: []rune("mb"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("mb"),
							Langs: -1,
						},
						{
							Text:  []rune("b"),
							Langs: 512,
						},
					},
				},
				{
					Pattern: []rune("mp"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("mp"),
							Langs: -1,
						},
						{
							Text:  []rune("b"),
							Langs: 512,
						},
					},
				},
				{
					Pattern: []rune("ng"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ng"),
							Langs: -1,
						},
						{
							Text:  []rune("g"),
							Langs: 512,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("t"),
							[]rune("S"),
							[]rune("s"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  []rune("f"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  []rune("f"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("p"),
							[]rune("k"),
							[]rune("t"),
							[]rune("S"),
							[]rune("s"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("f"),
							Langs: -1,
						},
						{
							Text:  []rune("p"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("f"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("f"),
							Langs: -1,
						},
						{
							Text:  []rune("p"),
							Langs: 262144,
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
						{
							Text:  []rune("v"),
							Langs: 262144,
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
						{
							Text:  []rune("b"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("t"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("t"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("g"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("n"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("g"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("k"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("n"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("k"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("p"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("r"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("E"),
							[]rune("e"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("r"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("s"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("s"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("t"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("o"),
							[]rune("u"),
							[]rune("A"),
							[]rune("E"),
							[]rune("I"),
							[]rune("O"),
							[]rune("U"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("t"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("s"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("o"),
							[]rune("u"),
							[]rune("A"),
							[]rune("E"),
							[]rune("I"),
							[]rune("O"),
							[]rune("U"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("s"),
							Langs: -1,
						},
						{
							Text:  nil,
							Langs: 64,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("o"),
							[]rune("u"),
							[]rune("A"),
							[]rune("E"),
							[]rune("I"),
							[]rune("B"),
							[]rune("F"),
							[]rune("O"),
							[]rune("U"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Q"),
							Langs: 128,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("D"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Q"),
							Langs: 128,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("lEE"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("b"),
							[]rune("d"),
							[]rune("f"),
							[]rune("g"),
							[]rune("k"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
							[]rune("Z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("li"),
							Langs: -1,
						},
						{
							Text:  []rune("il"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("rEE"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("b"),
							[]rune("d"),
							[]rune("f"),
							[]rune("g"),
							[]rune("k"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
							[]rune("Z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ri"),
							Langs: -1,
						},
						{
							Text:  []rune("ir"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("lE"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("b"),
							[]rune("d"),
							[]rune("f"),
							[]rune("g"),
							[]rune("k"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
							[]rune("Z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("li"),
							Langs: -1,
						},
						{
							Text:  []rune("il"),
							Langs: 32,
						},
						{
							Text:  []rune("lY"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("rE"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("b"),
							[]rune("d"),
							[]rune("f"),
							[]rune("g"),
							[]rune("k"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
							[]rune("Z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ri"),
							Langs: -1,
						},
						{
							Text:  []rune("ir"),
							Langs: 32,
						},
						{
							Text:  []rune("rY"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("EE"),
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
					Pattern: []rune("ea"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("eu"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("e"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("Ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("Ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("ei"),
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
					Pattern: []rune("Ei"),
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
					Pattern: []rune("iA"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ia"),
							Langs: -1,
						},
						{
							Text:  []rune("io"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iA"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ia"),
							Langs: -1,
						},
						{
							Text:  []rune("io"),
							Langs: -1,
						},
						{
							Text:  []rune("iY"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
						{
							Text:  []rune("D"),
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
						{
							Text:  nil,
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
						{
							Text:  nil,
							Langs: 32,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("A"),
							[]rune("O"),
							[]rune("I"),
							[]rune("U"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("A"),
							[]rune("O"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("P"),
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
					Pattern: []rune("O"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("o"),
							[]rune("e"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("o"),
							[]rune("e"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sUts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("suts"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("Uts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
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
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("A"),
							[]rune("O"),
							[]rune("I"),
							[]rune("U"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("A"),
							[]rune("O"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
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
							Text:  []rune("Y"),
							Langs: 128,
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
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Arabic): {
				{
					Pattern: []rune("1a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("1i"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("1u"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("j1"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ja"),
							Langs: -1,
						},
						{
							Text:  []rune("je"),
							Langs: -1,
						},
						{
							Text:  []rune("jo"),
							Langs: -1,
						},
						{
							Text:  []rune("ju"),
							Langs: -1,
						},
						{
							Text:  []rune("j"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("1"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("e"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
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
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("u"),
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
					Pattern: []rune("i"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("p"),
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
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
						{
							Text:  []rune("b"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Russian): {
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("E"),
							[]rune("I"),
							[]rune("o"),
							[]rune("u"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("om"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("om"),
							Langs: -1,
						},
						{
							Text:  []rune("im"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("on"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("on"),
							Langs: -1,
						},
						{
							Text:  []rune("in"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("em"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("im"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("en"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("in"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Em"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("im"),
							Langs: -1,
						},
						{
							Text:  []rune("Ym"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("En"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("in"),
							Langs: -1,
						},
						{
							Text:  []rune("Yn"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
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
						{
							Text:  []rune("o"),
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
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Cyrillic): {
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("E"),
							[]rune("I"),
							[]rune("o"),
							[]rune("u"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("om"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("om"),
							Langs: -1,
						},
						{
							Text:  []rune("im"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("on"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("on"),
							Langs: -1,
						},
						{
							Text:  []rune("in"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("em"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("im"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("en"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("in"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Em"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("im"),
							Langs: -1,
						},
						{
							Text:  []rune("Ym"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("En"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("in"),
							Langs: -1,
						},
						{
							Text:  []rune("Yn"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
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
						{
							Text:  []rune("o"),
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
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(French): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Czech): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Dutch): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(English): {
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("D"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("E"),
							[]rune("I"),
							[]rune("e"),
							[]rune("i"),
							[]rune("o"),
							[]rune("u"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("lE"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("b"),
							[]rune("d"),
							[]rune("f"),
							[]rune("g"),
							[]rune("k"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
							[]rune("Z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("il"),
							Langs: -1,
						},
						{
							Text:  []rune("li"),
							Langs: -1,
						},
						{
							Text:  []rune("lY"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("D[^aeiEIou]$"),
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
					Pattern: []rune("e"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("D[^aeiEIou]$"),
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
					Pattern: []rune("e"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("S"),
							[]rune("t"),
							[]rune("v"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("E"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
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
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(German): {
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("A"),
							[]rune("E"),
							[]rune("I"),
							[]rune("O"),
							[]rune("U"),
							[]rune("o"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("AU"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aU"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("OU"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("oU"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("Ou"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("Ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("Ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("e"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("A"),
							[]rune("O"),
							[]rune("U"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("A"),
							[]rune("O"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("A"),
							[]rune("O"),
							[]rune("U"),
							[]rune("e"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
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
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("e"),
							[]rune("O"),
							[]rune("U"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("U"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sUts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("suts"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Greek): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Greeklatin): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
				{
					Pattern: []rune("N"),
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Hebrew): {},
			common.Lang(Hungarian): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Italian): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Latvian): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Polish): {
				{
					Pattern: []rune("aiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
						{
							Text:  []rune("im"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: -1,
						},
						{
							Text:  []rune("in"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("im"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("in"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("A"),
							[]rune("E"),
							[]rune("B"),
							[]rune("F"),
							[]rune("I"),
							[]rune("o"),
							[]rune("u"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Portuguese): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
			},
			common.Lang(Romanian): {
				{
					Pattern: []rune("aiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiB"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
						{
							Text:  []rune("im"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: -1,
						},
						{
							Text:  []rune("in"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dm"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("oiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("uiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("eiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("EiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("iiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("IiF"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("Dn"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("im"),
							Langs: -1,
						},
						{
							Text:  []rune("om"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("d"),
							[]rune("g"),
							[]rune("k"),
							[]rune("s"),
							[]rune("t"),
							[]rune("v"),
							[]rune("z"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("in"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
						{
							Text:  []rune("Qk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("sits"),
							Langs: -1,
						},
						{
							Text:  []rune("sQts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Its"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("a"),
							[]rune("e"),
							[]rune("i"),
							[]rune("A"),
							[]rune("E"),
							[]rune("B"),
							[]rune("F"),
							[]rune("I"),
							[]rune("o"),
							[]rune("u"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("Q"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("l"),
							[]rune("m"),
							[]rune("n"),
							[]rune("p"),
							[]rune("r"),
							[]rune("s"),
							[]rune("t"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &common.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("D"),
							[]rune("a"),
							[]rune("o"),
							[]rune("i"),
							[]rune("u"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("Y"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Spanish): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
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
						{
							Text:  []rune("v"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
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
			},
			common.Lang(Turkish): {
				{
					Pattern: []rune("au"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("a"),
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
							Text:  []rune("D"),
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
					Pattern: []rune("ai"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
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
					Pattern: []rune("oi"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("ui"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("D"),
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
					Pattern: []rune("a"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
			},
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
					Prefix: [][]rune{
						[]rune("f"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("b"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("k"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("f"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("t"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("k"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("d"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("g"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("s"),
					},
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
					Prefix: [][]rune{
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("z"),
					},
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
					Prefix: [][]rune{
						[]rune("d"),
					},
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
					Prefix: [][]rune{
						[]rune("dZ"),
					},
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
					Prefix: [][]rune{
						[]rune("tS"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("S"),
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
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
					Prefix: [][]rune{
						[]rune("s"),
						[]rune("S"),
						[]rune("z"),
						[]rune("Z"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jnm"),
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("jm"),
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
				Pattern: []rune("jI"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("I"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("a"),
						[]rune("A"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("a"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("A"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("A"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
					},
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
					Prefix: [][]rune{
						[]rune("b"),
					},
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
					Prefix: [][]rune{
						[]rune("d"),
					},
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
					Prefix: [][]rune{
						[]rune("f"),
					},
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
					Prefix: [][]rune{
						[]rune("g"),
					},
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
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("j"),
					},
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
					Prefix: [][]rune{
						[]rune("k"),
					},
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
					Prefix: [][]rune{
						[]rune("l"),
					},
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
					Prefix: [][]rune{
						[]rune("m"),
					},
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
					Prefix: [][]rune{
						[]rune("n"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
					},
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
					Prefix: [][]rune{
						[]rune("r"),
					},
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
					Prefix: [][]rune{
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("v"),
					},
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
					Prefix: [][]rune{
						[]rune("z"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("H"),
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
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("g"),
						[]rune("Z"),
						[]rune("d"),
					},
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
					Prefix: [][]rune{
						[]rune("p"),
						[]rune("f"),
						[]rune("k"),
						[]rune("s"),
						[]rune("t"),
					},
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
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("g"),
						[]rune("z"),
						[]rune("d"),
					},
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
				Pattern: []rune("ji"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("a"),
						[]rune("A"),
						[]rune("o"),
						[]rune("O"),
						[]rune("e"),
						[]rune("E"),
						[]rune("i"),
						[]rune("I"),
						[]rune("u"),
						[]rune("U"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jI"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("a"),
						[]rune("A"),
						[]rune("o"),
						[]rune("O"),
						[]rune("e"),
						[]rune("E"),
						[]rune("i"),
						[]rune("I"),
						[]rune("u"),
						[]rune("U"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("je"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("a"),
						[]rune("A"),
						[]rune("o"),
						[]rune("O"),
						[]rune("e"),
						[]rune("E"),
						[]rune("i"),
						[]rune("I"),
						[]rune("u"),
						[]rune("U"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jE"),
				LeftContext: &common.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("a"),
						[]rune("A"),
						[]rune("o"),
						[]rune("O"),
						[]rune("e"),
						[]rune("E"),
						[]rune("i"),
						[]rune("I"),
						[]rune("u"),
						[]rune("U"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[common.Lang]common.Rules{
			common.Lang(Any): {
				{
					Pattern: []rune("EE"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("t"),
							[]rune("S"),
							[]rune("s"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
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
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("p"),
							[]rune("k"),
							[]rune("t"),
							[]rune("S"),
							[]rune("s"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("f"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("f"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
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
			common.Lang(Arabic): {
				{
					Pattern: []rune("1"),
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Russian): {
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Cyrillic): {
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Czech): {
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Dutch): {},
			common.Lang(English): {
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(French): {},
			common.Lang(German): {
				{
					Pattern: []rune("EE"),
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("f"),
							[]rune("k"),
							[]rune("t"),
							[]rune("S"),
							[]rune("s"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("p"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("p"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
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
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("p"),
							[]rune("k"),
							[]rune("t"),
							[]rune("S"),
							[]rune("s"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("f"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
					RightContext: &common.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("f"),
						},
					},
					Phonetic: []common.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("V"),
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
			common.Lang(Greek): {},
			common.Lang(Greeklatin): {
				{
					Pattern: []rune("N"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("n"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Hebrew): {},
			common.Lang(Hungarian): {
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Italian): {},
			common.Lang(Latvian): {},
			common.Lang(Polish): {
				{
					Pattern: []rune("B"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Portuguese): {},
			common.Lang(Romanian): {
				{
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			common.Lang(Spanish): {
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
			common.Lang(Turkish): {},
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
	"dos",
	"du",
	"el",
	"la",
	"le",
	"ibn",
	"van",
	"von",
	"ha",
	"vanden",
	"vander",
}
