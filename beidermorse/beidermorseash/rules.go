// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorseash

import (
	"github.com/f1monkey/phonetic/beidermorse/common"
	"regexp"
)

type Lang common.Lang

const (
	Any Lang = 1 << iota
	Cyrillic
	English
	French
	German
	Hebrew
	Hungarian
	Polish
	Romanian
	Russian
	Spanish
)

func (l Lang) String() string {
	switch l {
	case Any:
		return "any"
	case Cyrillic:
		return "cyrillic"
	case English:
		return "english"
	case French:
		return "french"
	case German:
		return "german"
	case Hebrew:
		return "hebrew"
	case Hungarian:
		return "hungarian"
	case Polish:
		return "polish"
	case Romanian:
		return "romanian"
	case Russian:
		return "russian"
	case Spanish:
		return "spanish"
	}
	return ""
}

const All = Cyrillic +
	English +
	French +
	German +
	Hebrew +
	Hungarian +
	Polish +
	Romanian +
	Russian +
	Spanish

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
					Langs: 512,
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
					Langs: 512,
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
					Langs: 512,
				},
				{
					Text:  []rune("lef"),
					Langs: 512,
				},
				{
					Text:  []rune("lova"),
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
					Langs: 512,
				},
				{
					Text:  []rune("lef"),
					Langs: 512,
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
					Langs: 512,
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
					Langs: 512,
				},
				{
					Text:  []rune("eva"),
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
					Langs: 512,
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
					Langs: 512,
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
					Langs: 512,
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
					Langs: 128,
				},
				{
					Text:  []rune("l"),
					Langs: 128,
				},
				{
					Text:  []rune("el"),
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("k"),
					Langs: 128,
				},
				{
					Text:  []rune("ek"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("el"),
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("ek"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 256,
				},
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("StS"),
					Langs: 512,
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
					Langs: 512,
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
					Langs: 516,
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
			Pattern: []rune("kh"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
					Langs: 516,
				},
				{
					Text:  []rune("kh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chs"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: 16,
				},
				{
					Text:  []rune("xs"),
					Langs: -1,
				},
				{
					Text:  []rune("tSs"),
					Langs: 516,
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
					Text:  []rune("k"),
					Langs: 256,
				},
				{
					Text:  []rune("tS"),
					Langs: 516,
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
					Langs: 516,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("tsu"),
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
					Langs: 128,
				},
				{
					Text:  []rune("tSi"),
					Langs: 384,
				},
				{
					Text:  []rune("tS"),
					Langs: 256,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tsi"),
					Langs: 128,
				},
				{
					Text:  []rune("tSi"),
					Langs: 384,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ce"),
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
					Text:  []rune("tsF"),
					Langs: 128,
				},
				{
					Text:  []rune("tSe"),
					Langs: 384,
				},
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ce"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tSe"),
					Langs: 384,
				},
				{
					Text:  []rune("tse"),
					Langs: 128,
				},
				{
					Text:  []rune("se"),
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
					Langs: 128,
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
			Pattern: []rune("ssp"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Sp"),
					Langs: 16,
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
					Langs: 16,
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
					Langs: 16,
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
					Langs: 16,
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
					Langs: 128,
				},
				{
					Text:  []rune("sB"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("sF"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: 16,
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
					Text:  []rune("Se"),
					Langs: 128,
				},
				{
					Text:  []rune("zi"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("So"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("si"),
					Langs: -1,
				},
				{
					Text:  []rune("zi"),
					Langs: 16,
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
					[]rune("i"),
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
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("gue"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gui"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("gi"),
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
					Langs: 256,
				},
				{
					Text:  []rune("gh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gauz"),
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
			Pattern: []rune("gaus"),
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
			Pattern: []rune("gol'ts"),
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
			Pattern: []rune("golts"),
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
			Pattern: []rune("gol'tz"),
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
			Pattern: []rune("goltz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("holts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gol'ts"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("golts"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("gol'tz"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("goltz"),
			LeftContext: &common.Matcher{
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
			Pattern: []rune("gendler"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hendler"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gejmer"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gejm"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geymer"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geym"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geimer"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geim"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gof"),
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
					Langs: 16,
				},
				{
					Text:  []rune("ji"),
					Langs: 8,
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
					Langs: 1024,
				},
				{
					Text:  []rune("dZe"),
					Langs: 260,
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
					Langs: 1024,
				},
				{
					Text:  []rune("dZi"),
					Langs: 260,
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
					Text:  []rune("dZe"),
					Langs: 260,
				},
				{
					Text:  []rune("hE"),
					Langs: 512,
				},
				{
					Text:  []rune("xe"),
					Langs: 1024,
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
					Text:  []rune("dZi"),
					Langs: 260,
				},
				{
					Text:  []rune("hI"),
					Langs: 512,
				},
				{
					Text:  []rune("xi"),
					Langs: 1024,
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
					Langs: 64,
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
					Langs: 64,
				},
			},
		},
		{
			Pattern: []rune("g"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Suffix: [][]rune{
					[]rune("j"),
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
					Langs: 512,
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
				{
					Text:  []rune("eZ"),
					Langs: 264,
				},
				{
					Text:  []rune("ex"),
					Langs: 1024,
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
			Pattern: []rune("lj"),
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
					Langs: 512,
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
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("ll"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
				{
					Text:  []rune("J"),
					Langs: 1024,
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
					Langs: 4,
				},
				{
					Text:  []rune("x"),
					Langs: 1024,
				},
				{
					Text:  []rune("Z"),
					Langs: 264,
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
					Langs: 1024,
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
					Langs: 16,
				},
				{
					Text:  []rune("k"),
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
					Langs: 128,
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
					Text:  []rune("rze"),
					Langs: -1,
				},
				{
					Text:  []rune("rtsE"),
					Langs: 16,
				},
				{
					Text:  []rune("Ze"),
					Langs: 128,
				},
				{
					Text:  []rune("re"),
					Langs: 128,
				},
				{
					Text:  []rune("rZe"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("ri"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 16,
				},
				{
					Text:  []rune("Z"),
					Langs: 128,
				},
				{
					Text:  []rune("r"),
					Langs: 128,
				},
				{
					Text:  []rune("rZ"),
					Langs: 128,
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
					Langs: 20,
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
					Langs: -1,
				},
				{
					Text:  []rune("tS"),
					Langs: 20,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
					Langs: 532,
				},
				{
					Text:  []rune("tz"),
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
				{
					Text:  []rune("zh"),
					Langs: 128,
				},
				{
					Text:  []rune("tsh"),
					Langs: 16,
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
					Langs: 128,
				},
				{
					Text:  []rune("zB"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("zF"),
					Langs: 128,
				},
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 16,
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
					Langs: 128,
				},
				{
					Text:  []rune("tsi"),
					Langs: 16,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Zo"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
				},
				{
					Text:  []rune("zi"),
					Langs: -1,
				},
				{
					Text:  []rune("tsi"),
					Langs: 16,
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
					Langs: 16,
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
			Pattern: []rune("vogel"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("vogel"),
					Langs: -1,
				},
				{
					Text:  []rune("fogel"),
					Langs: 16,
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
					Langs: 16,
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
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
				{
					Text:  []rune("x"),
					Langs: 384,
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
					Text:  []rune("H"),
					Langs: 20,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile(" $"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
					[]rune("in"),
				},
			},
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
					Langs: 8,
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
					Langs: 8,
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
					Langs: 8,
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
				{
					Text:  []rune("uje"),
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("ae"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: 16,
				},
				{
					Text:  []rune("aje"),
					Langs: 512,
				},
				{
					Text:  []rune("ae"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Y"),
					Langs: 16,
				},
				{
					Text:  []rune("oje"),
					Langs: 512,
				},
				{
					Text:  []rune("oe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: 4,
				},
				{
					Text:  []rune("aje"),
					Langs: 512,
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
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
					Langs: 16,
				},
				{
					Text:  []rune("oj"),
					Langs: 16,
				},
				{
					Text:  []rune("eu"),
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
					Text:  []rune("i"),
					Langs: 16,
				},
				{
					Text:  []rune("e"),
					Langs: 128,
				},
				{
					Text:  []rune("ije"),
					Langs: 512,
				},
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
					Text:  []rune("i"),
					Langs: 16,
				},
				{
					Text:  []rune("e"),
					Langs: 128,
				},
				{
					Text:  []rune("ije"),
					Langs: 512,
				},
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
				{
					Text:  []rune("ije"),
					Langs: 512,
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
			Pattern: []rune("io"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jo"),
					Langs: -1,
				},
				{
					Text:  []rune("e"),
					Langs: 512,
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
					Langs: 512,
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
					Langs: 256,
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
					Langs: 512,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: 4,
				},
				{
					Text:  []rune("o"),
					Langs: -1,
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
			},
		},
		{
			Pattern: []rune("ć"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("tS"),
					Langs: 128,
				},
				{
					Text:  []rune("ts"),
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
			Pattern: []rune("ñ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
				{
					Text:  []rune("nj"),
					Langs: 1024,
				},
			},
		},
		{
			Pattern: []rune("ś"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
					Langs: 128,
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
			Pattern: []rune("ţ"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ts"),
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
					Langs: 128,
				},
				{
					Text:  []rune("z"),
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
			Pattern: []rune("ă"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("e"),
					Langs: 256,
				},
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
			Pattern: []rune("ó"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("u"),
					Langs: 128,
				},
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
			Pattern: []rune("ű"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("Q"),
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
					Text:  []rune("A"),
					Langs: -1,
				},
				{
					Text:  []rune("B"),
					Langs: 128,
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
					Langs: 128,
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
					Langs: 128,
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
				{
					Text:  []rune("ts"),
					Langs: 128,
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
					Langs: 16,
				},
				{
					Text:  []rune("z"),
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
			Pattern: []rune("ей"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("jaj"),
					Langs: -1,
				},
				{
					Text:  []rune("aj"),
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
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ей"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("aj"),
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
	common.Lang(English): {
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
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
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
			Pattern: []rune("h"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("h"),
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
			Pattern: []rune("q"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
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
			Pattern: []rune("x"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			LeftContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile(" $"),
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
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
					Text:  []rune("aj"),
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
			Pattern: []rune("ei"),
			Phonetic: []common.RuleToken{
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
			Pattern: []rune("e"),
			Phonetic: []common.RuleToken{
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
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
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
			Pattern: []rune("y"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
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
			Pattern: []rune("yi"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("i"),
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
	common.Lang(German): {
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
			Pattern: []rune("sch"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("S"),
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
			Pattern: []rune("eu"),
			Phonetic: []common.RuleToken{
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
					Text:  []rune("aj"),
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
					Text:  []rune("aj"),
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
	common.Lang(Romanian): {
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
			Pattern: []rune("c"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("k"),
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
			Pattern: []rune("ei"),
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
			Pattern: []rune("i"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
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
			Pattern: []rune("gauz"),
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
			Pattern: []rune("gaus"),
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
			Pattern: []rune("gol'ts"),
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
			Pattern: []rune("golts"),
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
			Pattern: []rune("gol'tz"),
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
			Pattern: []rune("goltz"),
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
			Pattern: []rune("gejmer"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gejm"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geimer"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geim"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geymer"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajmer"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("geym"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hajm"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gendler"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hendler"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gof"),
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
			Pattern: []rune("gojf"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hojf"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("goyf"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hojf"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("goif"),
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("hojf"),
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
			Pattern: []rune("kog"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
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
					Text:  []rune("kog"),
					Langs: -1,
				},
				{
					Text:  []rune("koh"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kag"),
			LeftContext: &common.Matcher{
				MatchEmptyString: true,
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
					Text:  []rune("kag"),
					Langs: -1,
				},
				{
					Text:  []rune("kah"),
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
			Pattern: []rune("c"),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^ "),
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
			Pattern: []rune("y"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("I"),
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
			Pattern: []rune("j"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("x"),
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
			Pattern: []rune("ll"),
			Phonetic: []common.RuleToken{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
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
				{
					Text:  []rune("S"),
					Langs: -1,
				},
				{
					Text:  []rune("Z"),
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
			Contains: [][]rune{
				[]rune("zh"),
			},
		},
		Langs:  660,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eau"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ah"),
				[]rune("oh"),
				[]rune("eh"),
				[]rune("ih"),
				[]rune("uh"),
				[]rune("äh"),
				[]rune("öh"),
				[]rune("üh"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("vogel"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("vogel"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("witz"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("tz"),
			},
		},
		Langs:  532,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("tz"),
			},
		},
		Langs:  516,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("güe"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("güi"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ghe"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ghi"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("vici"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("schi"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("chsch"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("tsch"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ssch"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  528,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  528,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("rz"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("rz"),
			},
		},
		Langs:  144,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("[^aoeiuäöü]rz"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("rz[^aoeiuäöü]"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cki"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ska"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cka"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ue"),
			},
		},
		Langs:  528,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ae"),
			},
		},
		Langs:  532,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oe"),
			},
		},
		Langs:  540,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("th"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("th"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("th[^aoeiu]"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("mann"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cz"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("cy"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("niew"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("stein"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("heim"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("heimer"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ii"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("iy"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yy"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yi"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("yj"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ij"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gaus"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gauz"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gauz"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("goltz"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("gol'tz$"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("golts"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("gol'ts$"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("goltz"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("^gol'tz"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("golts"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("^gol'ts"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gendler"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gejmer"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gejm"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("geimer"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("geim"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("geymer"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("geym"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gof"),
			},
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("thal"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("zweig"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("ck"),
			},
		},
		Langs:  20,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("c"),
			},
		},
		Langs:  448,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sz"),
			},
		},
		Langs:  192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gue"),
			},
		},
		Langs:  1032,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gui"),
			},
		},
		Langs:  1032,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("guy"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("cs"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("cs"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dzs"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("zs"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("zs"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("wl"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("wr"),
			},
		},
		Langs:  148,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  64,
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
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gy"),
			},
		},
		Langs:  576,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ly"),
			},
		},
		Langs:  704,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ny"),
			},
		},
		Langs:  704,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ty"),
			},
		},
		Langs:  704,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("â"),
			},
		},
		Langs:  264,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ă"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("à"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ä"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("á"),
			},
		},
		Langs:  1088,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ą"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ć"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ç"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ę"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("é"),
			},
		},
		Langs:  1096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("è"),
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
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("í"),
			},
		},
		Langs:  1088,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("î"),
			},
		},
		Langs:  264,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ł"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ń"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ñ"),
			},
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ó"),
			},
		},
		Langs:  1216,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ö"),
			},
		},
		Langs:  80,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("õ"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ş"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ś"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ţ"),
			},
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ü"),
			},
		},
		Langs:  80,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ù"),
			},
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ű"),
			},
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ú"),
			},
		},
		Langs:  1088,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ź"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ż"),
			},
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ß"),
			},
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("а"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ё"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("о"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("е"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("и"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("у"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ы"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("э"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ю"),
			},
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("я"),
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
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ב"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ג"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ד"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ה"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ו"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ז"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ח"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ט"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("י"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("כ"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ל"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("מ"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("נ"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ס"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ע"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("פ"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("צ"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ק"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ר"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ש"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ת"),
			},
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("a"),
			},
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("o"),
			},
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("e"),
			},
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("i"),
			},
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("y"),
			},
		},
		Langs:  290,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("u"),
			},
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("v[^aoeiuäüö]"),
		},
		Langs:  16,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		Langs:  16,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("c[^aohk]"),
		},
		Langs:  16,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("dzi"),
			},
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ou"),
			},
		},
		Langs:  16,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("aj"),
			},
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ej"),
			},
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("oj"),
			},
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("uj"),
			},
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("k"),
			},
		},
		Langs:  256,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("v"),
			},
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("ky"),
			},
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("eu"),
			},
		},
		Langs:  640,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("w"),
			},
		},
		Langs:  1864,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("kie"),
			},
		},
		Langs:  1032,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("gie"),
			},
		},
		Langs:  1288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("q"),
			},
		},
		Langs:  960,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains: [][]rune{
				[]rune("sch"),
			},
		},
		Langs:  1224,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix: [][]rune{
				[]rune("h"),
			},
		},
		Langs:  512,
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
						[]rune("B"),
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
						[]rune("B"),
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
				Pattern: []rune("B"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("B"),
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
				Pattern: []rune("kAg"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
						[]rune("E"),
						[]rune("O"),
						[]rune("I"),
						[]rune("U"),
						[]rune("a"),
						[]rune("e"),
						[]rune("o"),
						[]rune("i"),
						[]rune("u"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("kOg"),
						Langs: -1,
					},
					{
						Text:  []rune("kO"),
						Langs: 512,
					},
				},
			},
			{
				Pattern: []rune("kOg"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
						[]rune("E"),
						[]rune("O"),
						[]rune("I"),
						[]rune("U"),
						[]rune("a"),
						[]rune("e"),
						[]rune("o"),
						[]rune("i"),
						[]rune("u"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("kAg"),
						Langs: -1,
					},
					{
						Text:  []rune("kA"),
						Langs: 512,
					},
				},
			},
			{
				Pattern: []rune("kog"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
						[]rune("E"),
						[]rune("O"),
						[]rune("I"),
						[]rune("U"),
						[]rune("a"),
						[]rune("e"),
						[]rune("o"),
						[]rune("i"),
						[]rune("u"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("kog"),
						Langs: -1,
					},
					{
						Text:  []rune("ko"),
						Langs: 512,
					},
				},
			},
			{
				Pattern: []rune("kag"),
				LeftContext: &common.Matcher{
					MatchEmptyString: true,
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("A"),
						[]rune("E"),
						[]rune("O"),
						[]rune("I"),
						[]rune("U"),
						[]rune("a"),
						[]rune("e"),
						[]rune("o"),
						[]rune("i"),
						[]rune("u"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("kag"),
						Langs: -1,
					},
					{
						Text:  []rune("ka"),
						Langs: 512,
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
				Pattern: []rune("F"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("bh"),
						[]rune("dh"),
						[]rune("gh"),
						[]rune("kh"),
						[]rune("ph"),
						[]rune("sh"),
						[]rune("th"),
						[]rune("vh"),
						[]rune("zh"),
						[]rune("Zh"),
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
				Pattern: []rune("F"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("bx"),
						[]rune("dx"),
						[]rune("gx"),
						[]rune("kx"),
						[]rune("px"),
						[]rune("sx"),
						[]rune("tx"),
						[]rune("vx"),
						[]rune("zx"),
						[]rune("Zx"),
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
				Pattern: []rune("B"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("bh"),
						[]rune("dh"),
						[]rune("gh"),
						[]rune("kh"),
						[]rune("ph"),
						[]rune("sh"),
						[]rune("th"),
						[]rune("vh"),
						[]rune("zh"),
						[]rune("Zh"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("a"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("B"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("bx"),
						[]rune("dx"),
						[]rune("gx"),
						[]rune("kx"),
						[]rune("px"),
						[]rune("sx"),
						[]rune("tx"),
						[]rune("vx"),
						[]rune("zx"),
						[]rune("Zx"),
					},
				},
				Phonetic: []common.RuleToken{
					{
						Text:  []rune("a"),
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
				Pattern: []rune("F"),
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
				Pattern: []rune("F"),
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
						Text:  []rune("F"),
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
					Pattern: []rune("b"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("b"),
							Langs: -1,
						},
						{
							Text:  []rune("v"),
							Langs: 1024,
						},
					},
				},
				{
					Pattern: []rune("J"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("z"),
							Langs: -1,
						},
					},
				},
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
					Pattern: []rune("AiB"),
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
					Pattern: []rune("OiB"),
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
					Pattern: []rune("UiB"),
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
					Pattern: []rune("AiB"),
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
					Pattern: []rune("OiB"),
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
					Pattern: []rune("UiB"),
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
							Langs: 128,
						},
						{
							Text:  []rune("im"),
							Langs: 128,
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
							Text:  []rune("a"),
							Langs: -1,
						},
						{
							Text:  []rune("o"),
							Langs: -1,
						},
						{
							Text:  []rune("on"),
							Langs: 128,
						},
						{
							Text:  []rune("in"),
							Langs: 128,
						},
					},
				},
				{
					Pattern: []rune("B"),
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
					Pattern: []rune("AiF"),
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
					Pattern: []rune("OiF"),
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
					Pattern: []rune("UiF"),
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
					Pattern: []rune("AiF"),
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
					Pattern: []rune("OiF"),
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
					Pattern: []rune("UiF"),
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
							Langs: 128,
						},
						{
							Text:  []rune("om"),
							Langs: 128,
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
							Langs: 128,
						},
						{
							Text:  []rune("on"),
							Langs: 128,
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
							Langs: 16,
						},
						{
							Text:  []rune("i"),
							Langs: -1,
						},
						{
							Text:  []rune("D"),
							Langs: 4,
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
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
						},
						{
							Text:  []rune("i"),
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
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("li"),
							Langs: -1,
						},
						{
							Text:  []rune("il"),
							Langs: 4,
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
							Langs: 4,
						},
						{
							Text:  []rune("lY"),
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
						},
						{
							Text:  []rune("D"),
							Langs: 4,
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
							Langs: 16,
						},
						{
							Text:  nil,
							Langs: 4,
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
							Langs: 16,
						},
						{
							Text:  nil,
							Langs: 4,
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
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
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
							Langs: 16,
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
			common.Lang(French): {
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
			common.Lang(Hebrew): {},
			common.Lang(Hungarian): {
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
						[]rune("B"),
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
						[]rune("B"),
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
				Pattern: []rune("B"),
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Prefix: [][]rune{
						[]rune("B"),
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
						Text:  []rune("h"),
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
					Pattern: []rune("A"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
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
					Pattern: []rune("E"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("e"),
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
					Pattern: []rune("J"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("l"),
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
			common.Lang(French): {
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
			common.Lang(German): {
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
					Pattern: []rune("B"),
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
					Pattern: []rune("F"),
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
					Pattern: []rune("J"),
					Phonetic: []common.RuleToken{
						{
							Text:  []rune("l"),
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
		},
	},
}

var Discards = []string{
	"ben",
	"bar",
	"ha",
}
