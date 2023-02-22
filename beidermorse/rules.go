// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorse

import (
	"github.com/f1monkey/phonetic/internal/bmpm"
	"regexp"
)

type Lang bmpm.Lang

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

var Rules = map[bmpm.Lang]bmpm.Rules{
	bmpm.Lang(Any): {
		{
			Pattern: []rune("yna"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("á"),
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
					Langs: 8,
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
					Langs: 16392,
				},
			},
		},
		{
			Pattern: []rune("pf"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("dh"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("bh"),
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
		{
			Pattern: []rune("ph"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("xS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			LeftContext: &bmpm.Matcher{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("sh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zh"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			Pattern: []rune("th"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gh"),
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
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
					[]rune("i"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
				{
					Text:  []rune("stS"),
					Langs: 524288,
				},
			},
		},
		{
			Pattern: []rune("ssz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
					Langs: 2048,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sj"),
			LeftContext: &bmpm.Matcher{
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
			Pattern: []rune("sj"),
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
			Pattern: []rune("sj"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("guy"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gli"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
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
					Text:  []rune("h"),
					Langs: 131072,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oue"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ae"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
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
			Pattern: []rune("ea"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("io"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iy"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("oe"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oy"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
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
			Pattern: []rune("ua"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("va"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ue"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("uy"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("uj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ya"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ye"),
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("yi"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("yo"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("i"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("ó"),
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("ó"),
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
			Pattern: []rune("e"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("om"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("â"),
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
			Pattern: []rune("ă"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("č"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ć"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("en"),
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
			Pattern: []rune("è"),
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
			Pattern: []rune("ě"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ģ"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
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
			Pattern: []rune("ī"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ı"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ł"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ř"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ţ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ť"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("ů"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ù"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ý"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ż"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ß"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("A"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("B"),
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
					Text:  []rune("E"),
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("O"),
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
					Langs: 32768,
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
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("V"),
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
				{
					Text:  []rune("w"),
					Langs: 48,
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
					Langs: 294912,
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
	bmpm.Lang(Arabic): {
		{
			Pattern: []rune("ا"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ب"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ب"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ت"),
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
			Pattern: []rune("ت"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ث"),
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
			Pattern: []rune("ث"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ج"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("ح"),
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
			Pattern: []rune("ح"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("خ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("د"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("د"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ذ"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ذ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ر"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ر"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ز"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ز"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("س"),
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
			Pattern: []rune("س"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ش"),
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
			Pattern: []rune("ش"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ص"),
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
			Pattern: []rune("ص"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ض"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ض"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ط"),
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
			Pattern: []rune("ط"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ظ"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ظ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ع"),
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
			Pattern: []rune("ع"),
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
			Pattern: []rune("ع"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("غ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ف"),
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
			Pattern: []rune("ف"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ق"),
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
			Pattern: []rune("ق"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ك"),
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
			Pattern: []rune("ك"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ل"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ل"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("م"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("م"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ن"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ن"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ه"),
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
			Pattern: []rune("ه"),
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
			Pattern: []rune("ه"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
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
			Pattern: []rune("ي\u200e"),
			Phonetic: []bmpm.RuleToken{
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
	bmpm.Lang(Cyrillic): {
		{
			Pattern: []rune("ця"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цю"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("циа"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цие"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цио"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("циу"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("сие"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("сио"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("зие"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("зио"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("с"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("с"),
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
			Pattern: []rune("гауз"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("haus"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гаус"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("haus"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гольц"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("holts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("геймер"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("hof"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гер"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ger"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ген"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("гин"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gin"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("г"),
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("и"),
					[]rune("у"),
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
			Pattern: []rune("г"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("и"),
					[]rune("у"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лю"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лё"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ие"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ыйе"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ые"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ий"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("о"),
					[]rune("у"),
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
			Pattern: []rune("ый"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("а"),
					[]rune("о"),
					[]rune("у"),
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
			Pattern: []rune("ий"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("ый"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("ей"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("а"),
					[]rune("е"),
					[]rune("о"),
					[]rune("у"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("е"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("эй"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ей"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ауе"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ауэ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("а"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("б"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("в"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("г"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("д"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("е"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ё"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("з"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("и"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("й"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("к"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("л"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("м"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("н"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("о"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("п"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("р"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("с"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("т"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("у"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ф"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("х"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ц"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ч"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ш"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("щ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("StS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ъ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ы"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ь"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("э"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ю"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("я"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Czech): {
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
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
				{
					Text:  []rune("kv"),
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
			Pattern: []rune("ei"),
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
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
			Pattern: []rune("č"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ž"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ň"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ť"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("ý"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ě"),
			Phonetic: []bmpm.RuleToken{
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
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
				{
					Text:  []rune("kv"),
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Dutch): {
		{
			Pattern: []rune("ssj"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sj"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
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
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("pf"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			LeftContext: &bmpm.Matcher{
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
			Pattern: []rune("th"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("th"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
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
					[]rune("a"),
					[]rune("e"),
					[]rune("i"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
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
			Pattern: []rune("aue"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ui"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
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
				{
					Text:  []rune("x"),
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
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
				{
					Text:  []rune("Q"),
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(English): {
		{
			Pattern: []rune("�"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mc"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("mak"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
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
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cc"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("c"),
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
			Pattern: []rune("c"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
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
			Pattern: []rune("gh"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
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
			Pattern: []rune("ph"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("who"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("hu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("wh"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("w"),
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
			Pattern: []rune("h"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]"),
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("H"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kn"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mb"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ng"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("kw"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tia"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("So"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("wr"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
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
				MatchEmptyString: true,
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
			Pattern: []rune("yi"),
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
			Pattern: []rune("aue"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oue"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ia"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oa"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ou"),
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
			Pattern: []rune("oo"),
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
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oy"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("r"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
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
					Text:  []rune("a"),
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
				{
					Text:  []rune("a"),
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(French): {
		{
			Pattern: []rune("lt"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("n"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("n"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("e"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("e"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("u"),
				},
			},
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("aill"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ll"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("e"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("m"),
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			LeftContext: &bmpm.Matcher{
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
			Pattern: []rune("au"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
				{
					Text:  []rune("va"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
				{
					Text:  []rune("Q"),
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
	bmpm.Lang(German): {
		{
			Pattern: []rune("ewitsch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owitsch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("evitsch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ovitsch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("witsch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("vitsch"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("xS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ziu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chs"),
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
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
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
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sp"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("st"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("St"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssp"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("kv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewitz"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("thal"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tal"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("th"),
			LeftContext: &bmpm.Matcher{
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
			Pattern: []rune("th"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("th"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rh"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("H"),
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
			Pattern: []rune("s"),
			RightContext: &bmpm.Matcher{
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
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
					[]rune("j"),
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
					[]rune("ä"),
					[]rune("ö"),
					[]rune("ü"),
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
			Pattern: []rune("ß"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("aue"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ue"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ae"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
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
			Pattern: []rune("ie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
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
			Pattern: []rune("ñ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
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
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("a"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("A"),
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("O"),
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
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []bmpm.RuleToken{
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
			},
		},
	},
	bmpm.Lang(Greek): {
		{
			Pattern: []rune("αυ"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("αυ"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("αυ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("av"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ευ"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ευ"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ευ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ev"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηυ"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("if"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηυ"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("if"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηυ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("iv"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ου"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("αι"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ει"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("οι"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ωι"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ηι"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("υι"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γγ"),
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
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
			Pattern: []rune("γγ"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γκ"),
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("ε"),
					[]rune("ι"),
					[]rune("η"),
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
			Pattern: []rune("γκ"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("γι"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
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
			Pattern: []rune("γι"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("α"),
					[]rune("ο"),
					[]rune("ω"),
					[]rune("υ"),
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
			Pattern: []rune("γε"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("τζ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("σ"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μβ"),
			Phonetic: []bmpm.RuleToken{
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
		{
			Pattern: []rune("μπ"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("mb"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μπ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ντ"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ντ"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("έ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ή"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ό"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ύ"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ΰ"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("α"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("β"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("δ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ε"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ζ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("η"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("κ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("λ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("μ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ν"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ξ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ο"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("π"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ρ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("σ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ς"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("τ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("υ"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("θ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("χ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ψ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ps"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ω"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Greeklatin): {
		{
			Pattern: []rune("au"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("af"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("av"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ef"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ev"),
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
			Pattern: []rune("gge"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ggi"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gk"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gke"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gki"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gk"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghi"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghi"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nghe"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
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
			Pattern: []rune("ghi"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
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
			Pattern: []rune("ghe"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Ng"),
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
			},
		},
		{
			Pattern: []rune("ngi"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ngi"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nge"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
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
			Pattern: []rune("gi"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
					[]rune("y"),
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
			Pattern: []rune("ge"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Ng"),
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
			Pattern: []rune("yi"),
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
			Pattern: []rune("yi"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
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
			Pattern: []rune("dh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("dj"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dZ"),
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
			Pattern: []rune("th"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dz"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mb"),
			Phonetic: []bmpm.RuleToken{
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
		{
			Pattern: []rune("mp"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("mp"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("mp"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nt"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("nt"),
			LeftContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("óu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("ý"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("x"),
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
			Pattern: []rune("ο"),
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
					Text:  []rune("TB"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Hungarian): {
		{
			Pattern: []rune("sz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zs"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cs"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("o"),
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("á"),
					[]rune("o"),
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
			Pattern: []rune("ee"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ty"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("ö"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Q"),
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
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("z"),
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
	bmpm.Lang(Latvian): {
		{
			Pattern: []rune("č"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ģ"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("š"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ņ"),
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
			Pattern: []rune("ž"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ā"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ē"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ī"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ū"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("u"),
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
			Pattern: []rune("ei"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ej"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("io"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
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
			Pattern: []rune("ui"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("uj"),
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
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
	bmpm.Lang(Polish): {
		{
			Pattern: []rune("ska"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ski"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cka"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tski"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lowa"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("czy"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tSi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cze"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("rjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("diewicz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("djevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tiewicz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iewicz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewicz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owicz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("icz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("itS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cia"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sia"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("zia"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("źe"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Zi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("źi"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
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
			Pattern: []rune("rze"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("t"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("n"),
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
			Pattern: []rune("s"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("s"),
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
			Pattern: []rune("ó"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("om"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("b"),
					[]rune("p"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("en"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ije"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yje"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
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
			Pattern: []rune("yj"),
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
			Pattern: []rune("ii"),
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
			Pattern: []rune("yi"),
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
			Pattern: []rune("iy"),
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
			Pattern: []rune("yy"),
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
			Pattern: []rune("rie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("rje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("die"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("dje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("F"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("e"),
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
			Pattern: []rune("au"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ej"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
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
			Pattern: []rune("aj"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("aj"),
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
			Pattern: []rune("a"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("B"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("P"),
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
					Text:  []rune("ts"),
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
					Text:  []rune("E"),
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
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("I"),
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
	bmpm.Lang(Romanian): {
		{
			Pattern: []rune("ce"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("ch"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			Phonetic: []bmpm.RuleToken{
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
			Pattern: []rune("gh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
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
			Pattern: []rune("ţ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
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
			Pattern: []rune("î"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ă"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("E"),
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("I"),
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
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
	},
	bmpm.Lang(Russian): {
		{
			Pattern: []rune("yna"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsyu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsia"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsio"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsyo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsiu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("syo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zyo"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ger"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ger"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gen"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gen"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gin"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("gin"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			LeftContext: &bmpm.Matcher{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
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
			Pattern: []rune("g"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
					[]rune("e"),
					[]rune("o"),
					[]rune("i"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
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
			Pattern: []rune("zh"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("i"),
					[]rune("e"),
					[]rune("y"),
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
			Pattern: []rune("qu"),
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
			Pattern: []rune("s"),
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("s"),
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
			Pattern: []rune("lya"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lyu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lia"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("liu"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lja"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lju"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("le"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yje"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yye"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yie"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ij"),
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
			Pattern: []rune("iy"),
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
			Pattern: []rune("ii"),
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
			Pattern: []rune("yj"),
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
			Pattern: []rune("yy"),
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
			Pattern: []rune("yi"),
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
			Pattern: []rune("io"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
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
			Pattern: []rune("yo"),
			Phonetic: []bmpm.RuleToken{
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
			RightContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Prefix: [][]rune{
					[]rune("a"),
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
					[]rune("i"),
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
			Pattern: []rune("ii"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("iy"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("yy"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("yi"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("yj"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("ij"),
			RightContext: &bmpm.Matcher{
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
			Pattern: []rune("e"),
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
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
			LeftContext: &bmpm.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("a"),
					[]rune("o"),
					[]rune("u"),
				},
			},
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []bmpm.RuleToken{
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
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
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
					Text:  []rune("E"),
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
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("I"),
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
			Pattern: []rune("b"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("B"),
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
					Text:  []rune("V"),
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
	},
	bmpm.Lang(Turkish): {
		{
			Pattern: []rune("ç"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ğ"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []bmpm.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ı"),
			Phonetic: []bmpm.RuleToken{
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
					Text:  []rune("dZ"),
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
					Text:  []rune("j"),
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
}

var LangRules = []bmpm.LangRule{
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("o’"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("o'"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("mc"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("fitz"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ceau"),
			},
		},
		Langs:  65600,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eau"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ault"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("oult"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("eux"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("eix"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("glou"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("uu"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tx"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("witz"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("tz"),
			},
		},
		Langs:  131232,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("tz"),
			},
		},
		Langs:  131104,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("poulos"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("pulos"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("iou"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("sj"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("sj"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("güe"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("güi"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ghe"),
			},
		},
		Langs:  66048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ghi"),
			},
		},
		Langs:  66048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("escu"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("esco"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("vici"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("schi"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ii"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("iy"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yy"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yi"),
			},
		},
		Langs:  131072,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("rz"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("rz"),
			},
		},
		Langs:  16512,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cki"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ska"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cka"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ae"),
			},
		},
		Langs:  131232,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oe"),
			},
		},
		Langs:  131312,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("th"),
			},
		},
		Langs:  160,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("th"),
			},
		},
		Langs:  672,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("mann"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cz"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cy"),
			},
		},
		Langs:  16896,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("niew"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("etti"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("eti"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ati"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ato"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("esi"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("oli"),
			},
		},
		Langs:  4096,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("field"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("stein"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("heim"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("heimer"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("thal"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("zweig"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("äh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("öh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("üh"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("chsch"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tsch"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  131200,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  131200,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ck"),
			},
		},
		Langs:  160,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("c"),
			},
		},
		Langs:  608264,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sz"),
			},
		},
		Langs:  18432,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("cs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dzs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("zs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("zs"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("wl"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("wr"),
			},
		},
		Langs:  16560,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  133696,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("guy"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ly"),
			},
		},
		Langs:  150016,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ny"),
			},
		},
		Langs:  412160,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ty"),
			},
		},
		Langs:  150016,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ā"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ć"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ç"),
			},
		},
		Langs:  819264,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("č"),
			},
		},
		Langs:  8200,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ď"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ē"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ğ"),
			},
		},
		Langs:  524288,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ģ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ī"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ķ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ļ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ł"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ņ"),
			},
		},
		Langs:  8192,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ń"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ñ"),
			},
		},
		Langs:  262144,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ň"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ř"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ś"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ş"),
			},
		},
		Langs:  589824,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("š"),
			},
		},
		Langs:  8200,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ţ"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ť"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ź"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ž"),
			},
		},
		Langs:  8200,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ż"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ß"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ä"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("á"),
			},
		},
		Langs:  297480,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("â"),
			},
		},
		Langs:  98368,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ă"),
			},
		},
		Langs:  65536,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ą"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("à"),
			},
		},
		Langs:  32768,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ã"),
			},
		},
		Langs:  32768,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ę"),
			},
		},
		Langs:  16384,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("é"),
			},
		},
		Langs:  2632,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("è"),
			},
		},
		Langs:  266304,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ê"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ě"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ê"),
			},
		},
		Langs:  32832,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("í"),
			},
		},
		Langs:  297480,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("î"),
			},
		},
		Langs:  65600,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ı"),
			},
		},
		Langs:  524288,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ó"),
			},
		},
		Langs:  317960,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ö"),
			},
		},
		Langs:  526464,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ô"),
			},
		},
		Langs:  32832,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("õ"),
			},
		},
		Langs:  34816,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ò"),
			},
		},
		Langs:  266240,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ű"),
			},
		},
		Langs:  2048,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ú"),
			},
		},
		Langs:  297480,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ü"),
			},
		},
		Langs:  821376,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ù"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ů"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ý"),
			},
		},
		Langs:  520,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("а"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ё"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("о"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("е"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("и"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("у"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ы"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("э"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ю"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("я"),
			},
		},
		Langs:  4,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("α"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ε"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("η"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ι"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ο"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("υ"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ω"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ا"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ب"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ت"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ث"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ج"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ح"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("خ'"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("د"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ذ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ر"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ز"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("س"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ش"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ص"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ض"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ط"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ظ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ع"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("غ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ف"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ق"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ك"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ل"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("م"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ن"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ه"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("و"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ي"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("آ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("إ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("أ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ؤ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ئ"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("א"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ב"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ג"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ד"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ה"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ו"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ז"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ח"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ט"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("י"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("כ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ל"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("מ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("נ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ס"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ע"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("פ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("צ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ק"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ר"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ש"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ת"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("a"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("o"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("e"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("i"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("y"),
			},
		},
		Langs:  75030,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("u"),
			},
		},
		Langs:  1286,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("j"),
			},
		},
		Langs:  4096,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("j[^aoeiuy]"),
		},
		Langs:  295488,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("g"),
			},
		},
		Langs:  8,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("k"),
			},
		},
		Langs:  364608,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("q"),
			},
		},
		Langs:  748056,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("v"),
			},
		},
		Langs:  16384,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("w"),
			},
		},
		Langs:  993864,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("x"),
			},
		},
		Langs:  534552,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dj"),
			},
		},
		Langs:  786432,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("v[^aoeiu]"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("c[^aohk]"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dzi"),
			},
		},
		Langs:  524512,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ou"),
			},
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("aj"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ej"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oj"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("uj"),
			},
		},
		Langs:  240,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eu"),
			},
		},
		Langs:  147456,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ky"),
			},
		},
		Langs:  16384,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("kie"),
			},
		},
		Langs:  262720,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gie"),
			},
		},
		Langs:  360960,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ch"),
			},
		},
		Langs:  524288,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("son"),
			},
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
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
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  280640,
		Accept: false,
	},
	{
		Matcher: &bmpm.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("h"),
			},
		},
		Langs:  131072,
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
				Pattern: []rune("jnm"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("jm"),
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
				Pattern: []rune("jI"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("I"),
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
						[]rune("A"),
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
				Pattern: []rune("a"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("A"),
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
				Pattern: []rune("A"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
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
				Pattern: []rune("j"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("j"),
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
				Pattern: []rune("vanden"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("b"),
						[]rune("p"),
					},
				},
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				Pattern: []rune("h"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("H"),
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("r"),
						[]rune("m"),
						[]rune("n"),
						[]rune("l"),
					},
				},
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
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
				Pattern: []rune("i"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
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
				Pattern: []rune("E"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
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
				Pattern: []rune("I"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
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
				Pattern: []rune("Q"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
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
				Pattern: []rune("Y"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Exact: [][]rune{
						[]rune("l"),
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
				Pattern: []rune("e"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
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
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
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
				Pattern: []rune("E"),
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
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
				RightContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				LeftContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oue"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AvE"),
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
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
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uje"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("iji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ije"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Iji"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjI"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ije"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjE"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ajo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ojo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ejo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uja"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ujo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ija"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ija"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjA"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ijo"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjO"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("D"),
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
				Pattern: []rune("lYndEr"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lander"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lAndEr"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lAnder"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("landEr"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lender"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEndEr"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lendEr"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("lEnder"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("lYnder"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("burk"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
				RightContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
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
		},
		Second: map[bmpm.Lang]bmpm.Rules{
			bmpm.Lang(Any): {
				{
					Pattern: []rune("mb"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
							Text:  []rune("f"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("B"),
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
					Pattern: []rune("B"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
							Text:  []rune("p"),
							Langs: 262144,
						},
					},
				},
				{
					Pattern: []rune("V"),
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
					Pattern: []rune("V"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("n"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("n"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("E"),
							[]rune("e"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Pattern: []rune("ea"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Pattern: []rune("Ei"),
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
					Pattern: []rune("iA"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
					Pattern: []rune("O"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sUts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
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
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
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
							Text:  []rune("Y"),
							Langs: 128,
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
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Arabic): {
				{
					Pattern: []rune("1a"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Pattern: []rune("i"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(Russian): {
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(Cyrillic): {
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(French): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Czech): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Dutch): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(English): {
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("D[^aeiEIou]$"),
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
					Pattern: []rune("e"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("D[^aeiEIou]$"),
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
					Pattern: []rune("e"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
							[]rune("Y"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(German): {
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Uk"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("uk"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sUts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(Greek): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Greeklatin): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("N"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Hebrew): {},
			bmpm.Lang(Hungarian): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Italian): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Latvian): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Polish): {
				{
					Pattern: []rune("aiB"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiF"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(Portuguese): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Romanian): {
				{
					Pattern: []rune("aiB"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("aiF"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("b"),
							[]rune("p"),
						},
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("I"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Pattern:          regexp.MustCompile("^[^k]$"),
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("Ik"),
					LeftContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Suffix: [][]rune{
							[]rune("l"),
							[]rune("r"),
						},
					},
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("ik"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("sIts"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
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
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Exact: [][]rune{
							[]rune("ts"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
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
					Pattern: []rune("E"),
					LeftContext: &bmpm.Matcher{
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
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: false,
						Prefix: [][]rune{
							[]rune("a"),
							[]rune("o"),
							[]rune("Q"),
						},
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
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
			bmpm.Lang(Spanish): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Pattern: []rune("V"),
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
			},
			bmpm.Lang(Turkish): {
				{
					Pattern: []rune("au"),
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
					Phonetic: []bmpm.RuleToken{
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
				Pattern: []rune("jnm"),
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("jm"),
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
				Pattern: []rune("jI"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: true,
				},
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("I"),
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
						[]rune("A"),
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
				Pattern: []rune("a"),
				LeftContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Suffix: [][]rune{
						[]rune("A"),
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
				Pattern: []rune("A"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
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
				Pattern: []rune("j"),
				RightContext: &bmpm.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("j"),
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
				Pattern: []rune("H"),
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
				Pattern: []rune("ji"),
				LeftContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jI"),
				LeftContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("je"),
				LeftContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jE"),
				LeftContext: &bmpm.Matcher{
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
				Phonetic: []bmpm.RuleToken{
					{
						Text:  []rune("j"),
						Langs: -1,
					},
				},
			},
		},
		Second: map[bmpm.Lang]bmpm.Rules{
			bmpm.Lang(Any): {
				{
					Pattern: []rune("EE"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
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
					Pattern: []rune("B"),
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
					Pattern: []rune("B"),
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
					Pattern: []rune("V"),
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
					Pattern: []rune("V"),
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
					Pattern: []rune("V"),
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
			bmpm.Lang(Arabic): {
				{
					Pattern: []rune("1"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  nil,
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Russian): {
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Cyrillic): {
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Czech): {
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Dutch): {},
			bmpm.Lang(English): {
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(French): {},
			bmpm.Lang(German): {
				{
					Pattern: []rune("EE"),
					RightContext: &bmpm.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("A"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
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
					Pattern: []rune("B"),
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
					Pattern: []rune("B"),
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
					Pattern: []rune("V"),
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
					Pattern: []rune("V"),
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
					Pattern: []rune("V"),
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
			bmpm.Lang(Greek): {},
			bmpm.Lang(Greeklatin): {
				{
					Pattern: []rune("N"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("n"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Hebrew): {},
			bmpm.Lang(Hungarian): {
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Italian): {},
			bmpm.Lang(Latvian): {},
			bmpm.Lang(Polish): {
				{
					Pattern: []rune("B"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Portuguese): {},
			bmpm.Lang(Romanian): {
				{
					Pattern: []rune("E"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []bmpm.RuleToken{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
			},
			bmpm.Lang(Spanish): {
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
			bmpm.Lang(Turkish): {},
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
