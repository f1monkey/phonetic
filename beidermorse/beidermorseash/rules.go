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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssch"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []common.Token{
				{
					Text:  []rune("xS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsch"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[äöü]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeiou]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kh"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("rjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("diewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("djevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tiewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("icz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("itS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cz"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssp"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeiouäöë]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("ge"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gui"),
			Phonetic: []common.Token{
				{
					Text:  []rune("gi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("guy"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("holts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("goltz"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[jyaeiou]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aouyei]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aouei]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aoeiuy]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Suffix:           []rune("t"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Suffix:           []rune("t"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Suffix:           []rune("t"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeiou]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("vogel"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Suffix:           []rune("in"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: true,
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("oue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("io"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ć"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ñ"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ţ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ż"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("om"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []common.Token{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ă"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []common.Token{
				{
					Text:  []rune("en"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("î"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ß"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("'"),
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("A"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("O"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цю"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("циа"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цие"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("цио"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("циу"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("сие"),
			Phonetic: []common.Token{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("сио"),
			Phonetic: []common.Token{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("зие"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("зио"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("(й|ё|я|ю|ы|а|е|о|и|у)$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лю"),
			Phonetic: []common.Token{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("лё"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ие"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ыйе"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ые"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ё"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("(а|е|о|у)$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("эй"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ей"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ауе"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ауэ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("а"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("б"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("в"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("г"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("д"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("е"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ж"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("з"),
			Phonetic: []common.Token{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("и"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("й"),
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("к"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("л"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("м"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("н"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("о"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("п"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("р"),
			Phonetic: []common.Token{
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
				Prefix:           []rune("с"),
			},
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("с"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("т"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("у"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ф"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("х"),
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ц"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ч"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ш"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("щ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("StS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ъ"),
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ы"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ь"),
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("э"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ю"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ju"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("я"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[iey]"),
			},
			Phonetic: []common.Token{
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
				Prefix:           []rune("c"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[iey]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gh"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[iey]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("kh"),
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("H"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("kw"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tia"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oue"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("ia"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oa"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oi"),
			Phonetic: []common.Token{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("ou"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Prefix:           []rune("r"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ph"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gn"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[bdgt]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eau"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ê"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("è"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("à"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("â"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("où"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ou"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oi"),
			Phonetic: []common.Token{
				{
					Text:  []rune("oj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[ou]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yi"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ii"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yy"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zia"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssch"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chsch"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("vitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sch"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("chs"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ck"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[eiy]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("St"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ssp"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("H"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ss"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[äöüaeiouy]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aeiouyäöüj]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeiouyäöü]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ß"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ae"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oe"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ä"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("eu"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ñ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ã"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ç"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("A"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("O"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("U"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("עי"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("עו"),
			Phonetic: []common.Token{
				{
					Text:  []rune("VV"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("או"),
			Phonetic: []common.Token{
				{
					Text:  []rune("VV"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ג׳"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ד׳"),
			Phonetic: []common.Token{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("א"),
			Phonetic: []common.Token{
				{
					Text:  []rune("L"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ב"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ג"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ד"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("1"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ה"),
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("וו"),
			Phonetic: []common.Token{
				{
					Text:  []rune("V"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("וי"),
			Phonetic: []common.Token{
				{
					Text:  []rune("WW"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ו"),
			Phonetic: []common.Token{
				{
					Text:  []rune("W"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ז"),
			Phonetic: []common.Token{
				{
					Text:  []rune("z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ח"),
			Phonetic: []common.Token{
				{
					Text:  []rune("X"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ט"),
			Phonetic: []common.Token{
				{
					Text:  []rune("T"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("יי"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("י"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ך"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("K"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("כ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ל"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ם"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("מ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ן"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("נ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ס"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ע"),
			Phonetic: []common.Token{
				{
					Text:  []rune("L"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ף"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("פ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ץ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("C"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("צ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("C"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ק"),
			Phonetic: []common.Token{
				{
					Text:  []rune("K"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ר"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ש"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ת"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zs"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cs"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[áo]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[áo]$"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ee"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("dj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gy"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("nj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ny"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("tj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ty"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ö"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ő"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Y"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ü"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ű"),
			Phonetic: []common.Token{
				{
					Text:  []rune("Q"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("rjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("diewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("djevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tiewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tjevitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ewicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("evitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("owicz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ovitS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("icz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("itS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("cz"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tS"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sz"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("Ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("źe"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("Zi"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("źi"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ź"),
			Phonetic: []common.Token{
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
				Suffix:           []rune("t"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Suffix:           []rune("t"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Suffix:           []rune("t"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ń"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.Token{
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
				Prefix:           []rune("s"),
			},
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bp]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("em"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ą"),
			Phonetic: []common.Token{
				{
					Text:  []rune("on"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ę"),
			Phonetic: []common.Token{
				{
					Text:  []rune("en"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ije"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yje"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yye"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("rie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("rje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("die"),
			Phonetic: []common.Token{
				{
					Text:  []rune("dje"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tie"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("F"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("au"),
			Phonetic: []common.Token{
				{
					Text:  []rune("au"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ej"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ai"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ay"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aj"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[ou]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[ou]$"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("P"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("Z"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ce"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tSe"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ci"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gi"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[ei]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("dZ"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gh"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ţ"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ş"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("î"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ea"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ja"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ă"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsyu"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsia"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsa"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsio"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tse"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsyo"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tso"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tsiu"),
			Phonetic: []common.Token{
				{
					Text:  []rune("tsu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sio"),
			Phonetic: []common.Token{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zio"),
			Phonetic: []common.Token{
				{
					Text:  []rune("zo"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("se"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("syo"),
			Phonetic: []common.Token{
				{
					Text:  []rune("so"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ze"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zyo"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("gin"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("gg"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[jaeoiuy]$"),
			},
			RightContext: &common.Matcher{
				MatchEmptyString: false,
				Pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ch"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("sh"),
			Phonetic: []common.Token{
				{
					Text:  []rune("S"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("zh"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("ts"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("tz"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[iey]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("qu"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Prefix:           []rune("s"),
			},
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("w"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lya"),
			Phonetic: []common.Token{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lyu"),
			Phonetic: []common.Token{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lia"),
			Phonetic: []common.Token{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("liu"),
			Phonetic: []common.Token{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lja"),
			Phonetic: []common.Token{
				{
					Text:  []rune("la"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("lju"),
			Phonetic: []common.Token{
				{
					Text:  []rune("lu"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("le"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("iie"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yje"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yye"),
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yie"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[aou]"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("io"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ei"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ey"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ej"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aj"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("yo"),
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("^[au]"),
			},
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aiou]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[aou]$"),
			},
			Phonetic: []common.Token{
				{
					Text:  []rune("je"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("oo"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("\""),
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("aue"),
			Phonetic: []common.Token{
				{
					Text:  []rune("aue"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("b"),
			Phonetic: []common.Token{
				{
					Text:  []rune("b"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("j"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
				Pattern:          regexp.MustCompile("[bdgt]$"),
			},
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  nil,
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("j"),
			Phonetic: []common.Token{
				{
					Text:  []rune("x"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("x"),
			Phonetic: []common.Token{
				{
					Text:  []rune("ks"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ll"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("v"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("c"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("z"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("q"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("uo"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("y"),
			Phonetic: []common.Token{
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
			Phonetic: []common.Token{
				{
					Text:  []rune("v"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("á"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("é"),
			Phonetic: []common.Token{
				{
					Text:  []rune("e"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("í"),
			Phonetic: []common.Token{
				{
					Text:  []rune("i"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ó"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("ú"),
			Phonetic: []common.Token{
				{
					Text:  []rune("u"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("a"),
			Phonetic: []common.Token{
				{
					Text:  []rune("a"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("d"),
			Phonetic: []common.Token{
				{
					Text:  []rune("d"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("e"),
			Phonetic: []common.Token{
				{
					Text:  []rune("E"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("f"),
			Phonetic: []common.Token{
				{
					Text:  []rune("f"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("g"),
			Phonetic: []common.Token{
				{
					Text:  []rune("g"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("h"),
			Phonetic: []common.Token{
				{
					Text:  []rune("h"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("i"),
			Phonetic: []common.Token{
				{
					Text:  []rune("I"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("k"),
			Phonetic: []common.Token{
				{
					Text:  []rune("k"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("l"),
			Phonetic: []common.Token{
				{
					Text:  []rune("l"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("m"),
			Phonetic: []common.Token{
				{
					Text:  []rune("m"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("n"),
			Phonetic: []common.Token{
				{
					Text:  []rune("n"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("o"),
			Phonetic: []common.Token{
				{
					Text:  []rune("o"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("p"),
			Phonetic: []common.Token{
				{
					Text:  []rune("p"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("r"),
			Phonetic: []common.Token{
				{
					Text:  []rune("r"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("s"),
			Phonetic: []common.Token{
				{
					Text:  []rune("s"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("t"),
			Phonetic: []common.Token{
				{
					Text:  []rune("t"),
					Langs: -1,
				},
			},
		},
		{
			Pattern: []rune("u"),
			Phonetic: []common.Token{
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
			Contains:         []rune("zh"),
		},
		Langs:  660,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("eau"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("[aoeiuäöü]h"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("vogel"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("vogel"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("witz"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("tz"),
		},
		Langs:  532,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("tz"),
		},
		Langs:  516,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("güe"),
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("güi"),
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ghe"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ghi"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("vici"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("schi"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("chsch"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("tsch"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ssch"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("sch"),
		},
		Langs:  528,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("sch"),
		},
		Langs:  528,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("rz"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("rz"),
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
			Suffix:           []rune("cki"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("ska"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("cka"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ue"),
		},
		Langs:  528,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ae"),
		},
		Langs:  532,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("oe"),
		},
		Langs:  540,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("th"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("th"),
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
			Contains:         []rune("mann"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("cz"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("cy"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("niew"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("stein"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("heim"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("heimer"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("ii"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("iy"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("yy"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("yi"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("yj"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("ij"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gaus"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gauz"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gauz"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("goltz"),
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
			Suffix:           []rune("golts"),
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
			Prefix:           []rune("goltz"),
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
			Prefix:           []rune("golts"),
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
			Suffix:           []rune("gendler"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gejmer"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gejm"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("geimer"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("geim"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("geymer"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("geym"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gof"),
		},
		Langs:  512,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("thal"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("zweig"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("ck"),
		},
		Langs:  20,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("c"),
		},
		Langs:  448,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("sz"),
		},
		Langs:  192,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gue"),
		},
		Langs:  1032,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gui"),
		},
		Langs:  1032,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("guy"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("cs"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("cs"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("dzs"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("zs"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("zs"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("wl"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("wr"),
		},
		Langs:  148,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Suffix:           []rune("gy"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Pattern:          regexp.MustCompile("gy[aeou]"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gy"),
		},
		Langs:  576,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ly"),
		},
		Langs:  704,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ny"),
		},
		Langs:  704,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ty"),
		},
		Langs:  704,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("â"),
		},
		Langs:  264,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ă"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("à"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ä"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("á"),
		},
		Langs:  1088,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ą"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ć"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ç"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ę"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("é"),
		},
		Langs:  1096,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("è"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ê"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("í"),
		},
		Langs:  1088,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("î"),
		},
		Langs:  264,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ł"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ń"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ñ"),
		},
		Langs:  1024,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ó"),
		},
		Langs:  1216,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ö"),
		},
		Langs:  80,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("õ"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ş"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ś"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ţ"),
		},
		Langs:  256,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ü"),
		},
		Langs:  80,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ù"),
		},
		Langs:  8,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ű"),
		},
		Langs:  64,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ú"),
		},
		Langs:  1088,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ź"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ż"),
		},
		Langs:  128,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ß"),
		},
		Langs:  16,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("а"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ё"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("о"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("е"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("и"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("у"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ы"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("э"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ю"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("я"),
		},
		Langs:  2,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("א"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ב"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ג"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ד"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ה"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ו"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ז"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ח"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ט"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("י"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("כ"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ל"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("מ"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("נ"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ס"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ע"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("פ"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("צ"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ק"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ר"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ש"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ת"),
		},
		Langs:  32,
		Accept: true,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("a"),
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("o"),
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("e"),
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("i"),
		},
		Langs:  34,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("y"),
		},
		Langs:  290,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("u"),
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
			Contains:         []rune("dzi"),
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ou"),
		},
		Langs:  16,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("aj"),
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ej"),
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("oj"),
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("uj"),
		},
		Langs:  28,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("k"),
		},
		Langs:  256,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("v"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("ky"),
		},
		Langs:  128,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("eu"),
		},
		Langs:  640,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("w"),
		},
		Langs:  1864,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("kie"),
		},
		Langs:  1032,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("gie"),
		},
		Langs:  1288,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("q"),
		},
		Langs:  960,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Contains:         []rune("sch"),
		},
		Langs:  1224,
		Accept: false,
	},
	{
		Matcher: &common.Matcher{
			MatchEmptyString: false,
			Prefix:           []rune("h"),
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[gdZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bdZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bgZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jnm"),
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[aAB]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[AB]$"),
				},
				Phonetic: []common.Token{
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
					Prefix:           []rune("A"),
				},
				Phonetic: []common.Token{
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
					Prefix:           []rune("B"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bp]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("H"),
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				RightContext: &common.Matcher{
					MatchEmptyString: false,
					Pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oue"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AvE"),
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uje"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("iji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ije"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Iji"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjI"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ije"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjE"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ajU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ajo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("oja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ojO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Oja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ojo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("OjO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("eja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ejO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Eja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ejo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("EjO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("uja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ujO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Uja"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ujo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("UjO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ija"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("ijO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ija"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjA"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Ijo"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("IjO"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Aju"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("AjU"),
				Phonetic: []common.Token{
					{
						Text:  []rune("D"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("j"),
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[rmnl]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  []rune("s"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("dZ"),
				Phonetic: []common.Token{
					{
						Text:  []rune("z"),
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("Z"),
				Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("e"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DoiuQY]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiEuQY]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQY]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aeiAEIOUouQY]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoAOUiuQY]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aoAOUeiuQY]$"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aoeOUiuQY]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuUQY]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("uts"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[bp]"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[lr]$"),
					},
					RightContext: &common.Matcher{
						MatchEmptyString: true,
					},
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("its"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^ts$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					Phonetic: []common.Token{
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
						Pattern:          regexp.MustCompile("^[aoQ]"),
					},
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[gdZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bdZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bgZz]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("jnm"),
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[aAB]"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[AB]$"),
				},
				Phonetic: []common.Token{
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
					Prefix:           []rune("A"),
				},
				Phonetic: []common.Token{
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
					Prefix:           []rune("B"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
					{
						Text:  nil,
						Langs: -1,
					},
				},
			},
			{
				Pattern: []rune("H"),
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("^[bgZd]"),
				},
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				Phonetic: []common.Token{
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
					Pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					Phonetic: []common.Token{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.Token{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("J"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("B"),
					Phonetic: []common.Token{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
						{
							Text:  []rune("i"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("O"),
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("U"),
					Phonetic: []common.Token{
						{
							Text:  []rune("u"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("J"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("a"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("F"),
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("P"),
					Phonetic: []common.Token{
						{
							Text:  []rune("o"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("E"),
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
					Phonetic: []common.Token{
						{
							Text:  []rune("e"),
							Langs: -1,
						},
					},
				},
				{
					Pattern: []rune("I"),
					Phonetic: []common.Token{
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
