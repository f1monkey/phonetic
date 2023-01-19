// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type ashLang languageID

const (
	ashany ashLang = 1 << iota
	ashcyrillic
	ashenglish
	ashfrench
	ashgerman
	ashhebrew
	ashhungarian
	ashpolish
	ashromanian
	ashrussian
	ashspanish
)

func (l ashLang) String() string {
	switch l {
	case ashany:
		return "any"
	case ashcyrillic:
		return "cyrillic"
	case ashenglish:
		return "english"
	case ashfrench:
		return "french"
	case ashgerman:
		return "german"
	case ashhebrew:
		return "hebrew"
	case ashhungarian:
		return "hungarian"
	case ashpolish:
		return "polish"
	case ashromanian:
		return "romanian"
	case ashrussian:
		return "russian"
	case ashspanish:
		return "spanish"
	}
	return ""
}

const ashAll = ashcyrillic +
	ashenglish +
	ashfrench +
	ashgerman +
	ashhebrew +
	ashhungarian +
	ashpolish +
	ashromanian +
	ashrussian +
	ashspanish

var ashRules = map[ashLang]rules{
	ashany: rules{
		{
			pattern: []rune("yna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: 512,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ina"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: 512,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lof"),
					langs: 512,
				},
				{
					text:  []rune("lef"),
					langs: 512,
				},
				{
					text:  []rune("lova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lof"),
					langs: 512,
				},
				{
					text:  []rune("lef"),
					langs: 512,
				},
				{
					text:  []rune("lova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("of"),
					langs: 512,
				},
				{
					text:  []rune("ova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eva"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: 512,
				},
				{
					text:  []rune("eva"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aia"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("aja"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("aya"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("lowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lova"),
					langs: -1,
				},
				{
					text:  []rune("lof"),
					langs: 128,
				},
				{
					text:  []rune("l"),
					langs: 128,
				},
				{
					text:  []rune("el"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("kowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kova"),
					langs: -1,
				},
				{
					text:  []rune("kof"),
					langs: 128,
				},
				{
					text:  []rune("k"),
					langs: 128,
				},
				{
					text:  []rune("ek"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("owa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ova"),
					langs: -1,
				},
				{
					text:  []rune("of"),
					langs: 128,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lovna"),
					langs: -1,
				},
				{
					text:  []rune("levna"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: 128,
				},
				{
					text:  []rune("el"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("kowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kovna"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: 128,
				},
				{
					text:  []rune("ek"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("owna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovna"),
					langs: -1,
				},
				{
					text:  nil,
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("lówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("kówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("ówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("rh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chsch"),
			phoneticRules: []token{
				{
					text:  []rune("xS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("sk"),
					langs: 256,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("StS"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("StS"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ssh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöü]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("sh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 516,
				},
				{
					text:  []rune("sh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: 516,
				},
				{
					text:  []rune("kh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chs"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: 16,
				},
				{
					text:  []rune("xs"),
					langs: -1,
				},
				{
					text:  []rune("tSs"),
					langs: 516,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: 256,
				},
				{
					text:  []rune("tS"),
					langs: 516,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 516,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("tsk"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("czy"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cze"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: -1,
				},
				{
					text:  []rune("tSF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciewicz"),
			phoneticRules: []token{
				{
					text:  []rune("tsevitS"),
					langs: -1,
				},
				{
					text:  []rune("tSevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siewicz"),
			phoneticRules: []token{
				{
					text:  []rune("sevitS"),
					langs: -1,
				},
				{
					text:  []rune("SevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziewicz"),
			phoneticRules: []token{
				{
					text:  []rune("zevitS"),
					langs: -1,
				},
				{
					text:  []rune("ZevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("riewicz"),
			phoneticRules: []token{
				{
					text:  []rune("rjevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("diewicz"),
			phoneticRules: []token{
				{
					text:  []rune("djevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tiewicz"),
			phoneticRules: []token{
				{
					text:  []rune("tjevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iewicz"),
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewicz"),
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owicz"),
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("icz"),
			phoneticRules: []token{
				{
					text:  []rune("itS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cz"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSB"),
					langs: 128,
				},
				{
					text:  []rune("tsB"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			phoneticRules: []token{
				{
					text:  []rune("tSa"),
					langs: 128,
				},
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSom"),
					langs: 128,
				},
				{
					text:  []rune("tsom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			phoneticRules: []token{
				{
					text:  []rune("tSon"),
					langs: 128,
				},
				{
					text:  []rune("tson"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSem"),
					langs: 128,
				},
				{
					text:  []rune("tsem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			phoneticRules: []token{
				{
					text:  []rune("tSen"),
					langs: 128,
				},
				{
					text:  []rune("tsen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSF"),
					langs: 128,
				},
				{
					text:  []rune("tsF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: 128,
				},
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cio"),
			phoneticRules: []token{
				{
					text:  []rune("tSo"),
					langs: 128,
				},
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciu"),
			phoneticRules: []token{
				{
					text:  []rune("tSu"),
					langs: 128,
				},
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tsi"),
					langs: 128,
				},
				{
					text:  []rune("tSi"),
					langs: 384,
				},
				{
					text:  []rune("tS"),
					langs: 256,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  []rune("tsi"),
					langs: 128,
				},
				{
					text:  []rune("tSi"),
					langs: 384,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ce"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tsF"),
					langs: 128,
				},
				{
					text:  []rune("tSe"),
					langs: 384,
				},
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ce"),
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: 384,
				},
				{
					text:  []rune("tse"),
					langs: 128,
				},
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cy"),
			phoneticRules: []token{
				{
					text:  []rune("si"),
					langs: -1,
				},
				{
					text:  []rune("tsi"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("ssz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: 16,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: 16,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sst"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: 16,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: 16,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("SB"),
					langs: 128,
				},
				{
					text:  []rune("sB"),
					langs: 128,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			phoneticRules: []token{
				{
					text:  []rune("Sa"),
					langs: 128,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Som"),
					langs: 128,
				},
				{
					text:  []rune("som"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			phoneticRules: []token{
				{
					text:  []rune("Son"),
					langs: 128,
				},
				{
					text:  []rune("son"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Sem"),
					langs: 128,
				},
				{
					text:  []rune("sem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			phoneticRules: []token{
				{
					text:  []rune("Sen"),
					langs: 128,
				},
				{
					text:  []rune("sen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("SF"),
					langs: 128,
				},
				{
					text:  []rune("sF"),
					langs: -1,
				},
				{
					text:  []rune("zi"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
				{
					text:  []rune("Se"),
					langs: 128,
				},
				{
					text:  []rune("zi"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: 128,
				},
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siu"),
			phoneticRules: []token{
				{
					text:  []rune("Su"),
					langs: 128,
				},
				{
					text:  []rune("sju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("si"),
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: 128,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
				{
					text:  []rune("zi"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouäöë]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("gue"),
			phoneticRules: []token{
				{
					text:  []rune("ge"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gui"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("guy"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: 256,
				},
				{
					text:  []rune("gh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gauz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gaus"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'ts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("golts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goltz"),
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'ts"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("golts"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'tz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goltz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gendler"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hendler"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejmer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejm"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geymer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geym"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geimer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geim"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gof"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hof"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ger"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ger"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gen"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gin"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gin"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gie"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ge"),
					langs: -1,
				},
				{
					text:  []rune("gi"),
					langs: 16,
				},
				{
					text:  []rune("ji"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("gie"),
			phoneticRules: []token{
				{
					text:  []rune("ge"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ge"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gE"),
					langs: -1,
				},
				{
					text:  []rune("xe"),
					langs: 1024,
				},
				{
					text:  []rune("dZe"),
					langs: 260,
				},
			},
		},
		{
			pattern: []rune("gi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("xi"),
					langs: 1024,
				},
				{
					text:  []rune("dZi"),
					langs: 260,
				},
			},
		},
		{
			pattern: []rune("ge"),
			phoneticRules: []token{
				{
					text:  []rune("gE"),
					langs: -1,
				},
				{
					text:  []rune("dZe"),
					langs: 260,
				},
				{
					text:  []rune("hE"),
					langs: 512,
				},
				{
					text:  []rune("xe"),
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("dZi"),
					langs: 260,
				},
				{
					text:  []rune("hI"),
					langs: 512,
				},
				{
					text:  []rune("xi"),
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("gy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("dj"),
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("gy"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("d"),
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[jyaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouyei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("eZ"),
					langs: 264,
				},
				{
					text:  []rune("ex"),
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ly"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("lj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("li"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("lj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("lj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("lyo"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ll"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("J"),
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("j"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aoeiuy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: 4,
				},
				{
					text:  []rune("x"),
					langs: 1024,
				},
				{
					text:  []rune("Z"),
					langs: 264,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  []rune("pf"),
					langs: -1,
				},
				{
					text:  []rune("p"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("ph"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: 16,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Se"),
					langs: 128,
				},
				{
					text:  []rune("re"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			phoneticRules: []token{
				{
					text:  []rune("rze"),
					langs: -1,
				},
				{
					text:  []rune("rtsE"),
					langs: 16,
				},
				{
					text:  []rune("Ze"),
					langs: 128,
				},
				{
					text:  []rune("re"),
					langs: 128,
				},
				{
					text:  []rune("rZe"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: 128,
				},
				{
					text:  []rune("ri"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: 128,
				},
				{
					text:  []rune("ri"),
					langs: 128,
				},
				{
					text:  []rune("rZi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 128,
				},
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			phoneticRules: []token{
				{
					text:  []rune("rz"),
					langs: -1,
				},
				{
					text:  []rune("rts"),
					langs: 16,
				},
				{
					text:  []rune("Z"),
					langs: 128,
				},
				{
					text:  []rune("r"),
					langs: 128,
				},
				{
					text:  []rune("rZ"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 20,
				},
			},
		},
		{
			pattern: []rune("tz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 20,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: 532,
				},
				{
					text:  []rune("tz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zh"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("zh"),
					langs: 128,
				},
				{
					text:  []rune("tsh"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("zia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ZB"),
					langs: 128,
				},
				{
					text:  []rune("zB"),
					langs: 128,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  []rune("Za"),
					langs: 128,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zom"),
					langs: 128,
				},
				{
					text:  []rune("zom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			phoneticRules: []token{
				{
					text:  []rune("Zon"),
					langs: 128,
				},
				{
					text:  []rune("zon"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zem"),
					langs: 128,
				},
				{
					text:  []rune("zem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			phoneticRules: []token{
				{
					text:  []rune("Zen"),
					langs: 128,
				},
				{
					text:  []rune("zen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ZF"),
					langs: 128,
				},
				{
					text:  []rune("zF"),
					langs: 128,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
				{
					text:  []rune("tsi"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
				{
					text:  []rune("Ze"),
					langs: 128,
				},
				{
					text:  []rune("tsi"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("Zo"),
					langs: 128,
				},
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  []rune("Zu"),
					langs: 128,
				},
				{
					text:  []rune("zju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zi"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: 128,
				},
				{
					text:  []rune("zi"),
					langs: -1,
				},
				{
					text:  []rune("tsi"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("thal"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tal"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: 16,
				},
				{
					text:  []rune("th"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vogel"),
			phoneticRules: []token{
				{
					text:  []rune("vogel"),
					langs: -1,
				},
				{
					text:  []rune("fogel"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("v"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: 384,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("H"),
					langs: 20,
				},
			},
		},
		{
			pattern: []rune("yi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile(" $"),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("in"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  nil,
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oue"),
			phoneticRules: []token{
				{
					text:  []rune("oue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("ou"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("uje"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ae"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: 16,
				},
				{
					text:  []rune("aje"),
					langs: 512,
				},
				{
					text:  []rune("ae"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: 16,
				},
				{
					text:  []rune("oje"),
					langs: 512,
				},
				{
					text:  []rune("oe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: 4,
				},
				{
					text:  []rune("aje"),
					langs: 512,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: 16,
				},
				{
					text:  []rune("oj"),
					langs: 16,
				},
				{
					text:  []rune("eu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: 16,
				},
				{
					text:  []rune("e"),
					langs: 128,
				},
				{
					text:  []rune("ije"),
					langs: 512,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: 16,
				},
				{
					text:  []rune("e"),
					langs: 128,
				},
				{
					text:  []rune("ije"),
					langs: 512,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
				{
					text:  []rune("ije"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("yo"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  []rune("ea"),
					langs: -1,
				},
				{
					text:  []rune("ja"),
					langs: 256,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: 4,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uu"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: 128,
				},
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ł"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ń"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("ś"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 128,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ţ"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ż"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ź"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: 128,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("où"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("om"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			phoneticRules: []token{
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ä"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ă"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: 256,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("em"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			phoneticRules: []token{
				{
					text:  []rune("en"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("î"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: 128,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ß"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("\""),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("A"),
					langs: -1,
				},
				{
					text:  []rune("B"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
				{
					text:  []rune("F"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("O"),
					langs: -1,
				},
				{
					text:  []rune("P"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("A"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("O"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("U"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: 16,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashcyrillic: rules{
		{
			pattern: []rune("ця"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цю"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("циа"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цие"),
			phoneticRules: []token{
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цио"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("циу"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("сие"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("сио"),
			phoneticRules: []token{
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("зие"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("зио"),
			phoneticRules: []token{
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гауз"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гаус"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гольц"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("геймер"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гейм"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гоф"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hof"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гер"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ger"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ген"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гин"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gin"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(й|ё|я|ю|ы|а|е|о|и|у)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ля"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лю"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лё"),
			phoneticRules: []token{
				{
					text:  []rune("le"),
					langs: -1,
				},
				{
					text:  []rune("lo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лио"),
			phoneticRules: []token{
				{
					text:  []rune("le"),
					langs: -1,
				},
				{
					text:  []rune("lo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ле"),
			phoneticRules: []token{
				{
					text:  []rune("lE"),
					langs: -1,
				},
				{
					text:  []rune("lo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ийе"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ие"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ыйе"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ые"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ий"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ый"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ий"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ый"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ё"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("jo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("jaj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(а|е|о|у)$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("эй"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ауе"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ауэ"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("а"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("б"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("в"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("д"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ж"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("з"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("и"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("й"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("к"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("л"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("м"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("н"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("о"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("п"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("р"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("с"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("с"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("с"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("т"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("у"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ф"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("х"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ц"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ч"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ш"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("щ"),
			phoneticRules: []token{
				{
					text:  []rune("StS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ъ"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ы"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ь"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("э"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ю"),
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("я"),
			phoneticRules: []token{
				{
					text:  []rune("ja"),
					langs: -1,
				},
			},
		},
	},
	ashenglish: rules{
		{
			pattern: []rune("tch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("c"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
				{
					text:  []rune("w"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			phoneticRules: []token{
				{
					text:  []rune("gn"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("sk"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("who"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("w"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("H"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kn"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mb"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ng"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("N"),
					langs: -1,
				},
				{
					text:  []rune("ng"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pn"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("pn"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ps"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ps"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kw"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tia"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: -1,
				},
				{
					text:  []rune("Sa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tio"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wr"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("w"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile(" $"),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
				{
					text:  []rune("oue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ear"),
			phoneticRules: []token{
				{
					text:  []rune("ia"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oa"),
			phoneticRules: []token{
				{
					text:  []rune("ou"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("ou"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oy"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ou"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("r"),
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashfrench: rules{
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("gn"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gue"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("que"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyéèê]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ouh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
				{
					text:  []rune("uh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  []rune("vo"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eau"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("où"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashgerman: rules{
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chsch"),
			phoneticRules: []token{
				{
					text:  []rune("xS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("evitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ovitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("witsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chs"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: -1,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: -1,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sst"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: -1,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: -1,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  []rune("pf"),
					langs: -1,
				},
				{
					text:  []rune("p"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("ph"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("evitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eviz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ovitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oviz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("witz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("viz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("thal"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tal"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("th"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("H"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöüaeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyäöüj]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouyäöü]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ß"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ae"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ä"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ã"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("A"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("O"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("U"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
	},
	ashhebrew: rules{
		{
			pattern: []rune("אי"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("עי"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("עו"),
			phoneticRules: []token{
				{
					text:  []rune("VV"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("או"),
			phoneticRules: []token{
				{
					text:  []rune("VV"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ג׳"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ד׳"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("א"),
			phoneticRules: []token{
				{
					text:  []rune("L"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ב"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ג"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ד"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("וו"),
			phoneticRules: []token{
				{
					text:  []rune("V"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("וי"),
			phoneticRules: []token{
				{
					text:  []rune("WW"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ו"),
			phoneticRules: []token{
				{
					text:  []rune("W"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ז"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ח"),
			phoneticRules: []token{
				{
					text:  []rune("X"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ט"),
			phoneticRules: []token{
				{
					text:  []rune("T"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("יי"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("י"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ך"),
			phoneticRules: []token{
				{
					text:  []rune("X"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("כ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("K"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("כ"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ל"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ם"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("מ"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ן"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("נ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ס"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ע"),
			phoneticRules: []token{
				{
					text:  []rune("L"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ף"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("פ"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ץ"),
			phoneticRules: []token{
				{
					text:  []rune("C"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("צ"),
			phoneticRules: []token{
				{
					text:  []rune("C"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ק"),
			phoneticRules: []token{
				{
					text:  []rune("K"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ר"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ש"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ת"),
			phoneticRules: []token{
				{
					text:  []rune("TB"),
					langs: -1,
				},
			},
		},
	},
	ashhungarian: rules{
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zs"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cs"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aj"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[áo]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[áo]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ely"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("eli"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ly"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("li"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gy"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
				{
					text:  []rune("gi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("ni"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ty"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ty"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("ti"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("ku"),
					langs: -1,
				},
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashpolish: rules{
		{
			pattern: []rune("ska"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ski"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cka"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tski"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lova"),
					langs: -1,
				},
				{
					text:  []rune("lof"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kova"),
					langs: -1,
				},
				{
					text:  []rune("kof"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ova"),
					langs: -1,
				},
				{
					text:  []rune("of"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lovna"),
					langs: -1,
				},
				{
					text:  []rune("levna"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kovna"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovna"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("czy"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cze"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: -1,
				},
				{
					text:  []rune("tSF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciewicz"),
			phoneticRules: []token{
				{
					text:  []rune("tsevitS"),
					langs: -1,
				},
				{
					text:  []rune("tSevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siewicz"),
			phoneticRules: []token{
				{
					text:  []rune("sevitS"),
					langs: -1,
				},
				{
					text:  []rune("SevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziewicz"),
			phoneticRules: []token{
				{
					text:  []rune("zevitS"),
					langs: -1,
				},
				{
					text:  []rune("ZevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("riewicz"),
			phoneticRules: []token{
				{
					text:  []rune("rjevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("diewicz"),
			phoneticRules: []token{
				{
					text:  []rune("djevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tiewicz"),
			phoneticRules: []token{
				{
					text:  []rune("tjevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iewicz"),
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewicz"),
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owicz"),
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("icz"),
			phoneticRules: []token{
				{
					text:  []rune("itS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cz"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSB"),
					langs: -1,
				},
				{
					text:  []rune("tsB"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			phoneticRules: []token{
				{
					text:  []rune("tSa"),
					langs: -1,
				},
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSom"),
					langs: -1,
				},
				{
					text:  []rune("tsom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			phoneticRules: []token{
				{
					text:  []rune("tSon"),
					langs: -1,
				},
				{
					text:  []rune("tson"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSem"),
					langs: -1,
				},
				{
					text:  []rune("tsem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			phoneticRules: []token{
				{
					text:  []rune("tSen"),
					langs: -1,
				},
				{
					text:  []rune("tsen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSF"),
					langs: -1,
				},
				{
					text:  []rune("tsF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: -1,
				},
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cio"),
			phoneticRules: []token{
				{
					text:  []rune("tSo"),
					langs: -1,
				},
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciu"),
			phoneticRules: []token{
				{
					text:  []rune("tSu"),
					langs: -1,
				},
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
				{
					text:  []rune("tsI"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("SB"),
					langs: -1,
				},
				{
					text:  []rune("sB"),
					langs: -1,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			phoneticRules: []token{
				{
					text:  []rune("Sa"),
					langs: -1,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Som"),
					langs: -1,
				},
				{
					text:  []rune("som"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			phoneticRules: []token{
				{
					text:  []rune("Son"),
					langs: -1,
				},
				{
					text:  []rune("son"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Sem"),
					langs: -1,
				},
				{
					text:  []rune("sem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			phoneticRules: []token{
				{
					text:  []rune("Sen"),
					langs: -1,
				},
				{
					text:  []rune("sen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("SF"),
					langs: -1,
				},
				{
					text:  []rune("sF"),
					langs: -1,
				},
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  []rune("Se"),
					langs: -1,
				},
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: -1,
				},
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siu"),
			phoneticRules: []token{
				{
					text:  []rune("Su"),
					langs: -1,
				},
				{
					text:  []rune("sju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("si"),
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: -1,
				},
				{
					text:  []rune("sI"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ś"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ZB"),
					langs: -1,
				},
				{
					text:  []rune("zB"),
					langs: -1,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  []rune("Za"),
					langs: -1,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zom"),
					langs: -1,
				},
				{
					text:  []rune("zom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			phoneticRules: []token{
				{
					text:  []rune("Zon"),
					langs: -1,
				},
				{
					text:  []rune("zon"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zem"),
					langs: -1,
				},
				{
					text:  []rune("zem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			phoneticRules: []token{
				{
					text:  []rune("Zen"),
					langs: -1,
				},
				{
					text:  []rune("zen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ZF"),
					langs: -1,
				},
				{
					text:  []rune("zF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("Zo"),
					langs: -1,
				},
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  []rune("Zu"),
					langs: -1,
				},
				{
					text:  []rune("zju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zi"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
				{
					text:  []rune("zI"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ZF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ZF"),
					langs: -1,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
				{
					text:  []rune("zF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("źe"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ży"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("źi"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
				{
					text:  []rune("zi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ż"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ź"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Se"),
					langs: -1,
				},
				{
					text:  []rune("re"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("re"),
					langs: -1,
				},
				{
					text:  []rune("rZe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: -1,
				},
				{
					text:  []rune("ri"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
				{
					text:  []rune("ri"),
					langs: -1,
				},
				{
					text:  []rune("rZi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("r"),
					langs: -1,
				},
				{
					text:  []rune("rZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ł"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ń"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("s"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("om"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("em"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			phoneticRules: []token{
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			phoneticRules: []token{
				{
					text:  []rune("en"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ije"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yje"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rie"),
			phoneticRules: []token{
				{
					text:  []rune("rje"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("die"),
			phoneticRules: []token{
				{
					text:  []rune("dje"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tie"),
			phoneticRules: []token{
				{
					text:  []rune("tje"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("F"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aj"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("B"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
				{
					text:  []rune("F"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("P"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashromanian: rules{
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ce"),
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  []rune("dZi"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ţ"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("î"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  []rune("ja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ă"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashrussian: rules{
		{
			pattern: []rune("yna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: -1,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ina"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: -1,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lof"),
					langs: -1,
				},
				{
					text:  []rune("lef"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lof"),
					langs: -1,
				},
				{
					text:  []rune("lef"),
					langs: -1,
				},
				{
					text:  []rune("lova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("of"),
					langs: -1,
				},
				{
					text:  []rune("ova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eva"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: -1,
				},
				{
					text:  []rune("ova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aia"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aja"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aya"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsya"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsyu"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsia"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsie"),
			phoneticRules: []token{
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsio"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsye"),
			phoneticRules: []token{
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsyo"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsiu"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sye"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("syo"),
			phoneticRules: []token{
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zye"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zyo"),
			phoneticRules: []token{
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gauz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gaus"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'ts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("golts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goltz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejmer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejm"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geimer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geim"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geymer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geym"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gendler"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hendler"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gof"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hof"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gojf"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hojf"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goyf"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hojf"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goif"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hojf"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ger"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ger"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gen"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gin"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gin"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kog"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("kog"),
					langs: -1,
				},
				{
					text:  []rune("koh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kag"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("kag"),
					langs: -1,
				},
				{
					text:  []rune("kah"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[jaeoiuy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("StS"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zh"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  []rune("tz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("s"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lya"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lyu"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lia"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liu"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lja"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lju"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("le"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("lE"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lyo"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ije"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yje"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yo"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("aje"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("oo"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("\""),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
	ashspanish: rules{
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ll"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bpvf]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("gv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  []rune("vo"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
	},
}

var ashLangRules = []langRule{
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("zh"),
		},
		langs:  660,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("eau"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[aoeiuäöü]h"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("vogel"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("vogel"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("witz"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("tz"),
		},
		langs:  532,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("tz"),
		},
		langs:  516,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("güe"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("güi"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ghe"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ghi"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("vici"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("schi"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("chsch"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("tsch"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ssch"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("sch"),
		},
		langs:  528,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("sch"),
		},
		langs:  528,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("rz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("rz"),
		},
		langs:  144,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[^aoeiuäöü]rz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("rz[^aoeiuäöü]"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("cki"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ska"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("cka"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ue"),
		},
		langs:  528,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ae"),
		},
		langs:  532,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("oe"),
		},
		langs:  540,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("th"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("th"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("th[^aoeiu]"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("mann"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("cz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("cy"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("niew"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("stein"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("heim"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("heimer"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ii"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("iy"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("yy"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("yi"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("yj"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ij"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gaus"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gauz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gauz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("goltz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gol'tz$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("golts"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gol'ts$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("goltz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("^gol'tz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("golts"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("^gol'ts"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gendler"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gejmer"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gejm"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("geimer"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("geim"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("geymer"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("geym"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gof"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("thal"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("zweig"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ck"),
		},
		langs:  20,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("c"),
		},
		langs:  448,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sz"),
		},
		langs:  192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gue"),
		},
		langs:  1032,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gui"),
		},
		langs:  1032,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("guy"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("cs"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("cs"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("dzs"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("zs"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("zs"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("wl"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("wr"),
		},
		langs:  148,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gy"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gy[aeou]"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gy"),
		},
		langs:  576,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ly"),
		},
		langs:  704,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ny"),
		},
		langs:  704,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ty"),
		},
		langs:  704,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("â"),
		},
		langs:  264,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ă"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("à"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ä"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("á"),
		},
		langs:  1088,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ą"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ć"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ç"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ę"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("é"),
		},
		langs:  1096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("è"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ê"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("í"),
		},
		langs:  1088,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("î"),
		},
		langs:  264,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ł"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ń"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ñ"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ó"),
		},
		langs:  1216,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ö"),
		},
		langs:  80,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("õ"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ş"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ś"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ţ"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ü"),
		},
		langs:  80,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ù"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ű"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ú"),
		},
		langs:  1088,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ź"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ż"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ß"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("а"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ё"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("о"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("е"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("и"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("у"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ы"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("э"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ю"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("я"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("א"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ב"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ג"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ד"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ה"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ו"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ז"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ח"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ט"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("י"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("כ"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ל"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("מ"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("נ"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ס"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ע"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("פ"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("צ"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ק"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ר"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ש"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ת"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("a"),
		},
		langs:  34,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("o"),
		},
		langs:  34,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("e"),
		},
		langs:  34,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("i"),
		},
		langs:  34,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("y"),
		},
		langs:  290,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("u"),
		},
		langs:  34,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("v[^aoeiuäüö]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("c[^aohk]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("dzi"),
		},
		langs:  28,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ou"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("aj"),
		},
		langs:  28,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ej"),
		},
		langs:  28,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("oj"),
		},
		langs:  28,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("uj"),
		},
		langs:  28,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("k"),
		},
		langs:  256,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("v"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ky"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("eu"),
		},
		langs:  640,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("w"),
		},
		langs:  1864,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("kie"),
		},
		langs:  1032,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gie"),
		},
		langs:  1288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("q"),
		},
		langs:  960,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sch"),
		},
		langs:  1224,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("h"),
		},
		langs:  512,
		accept: false,
	},
}

var ashFinalRules = finalRules{
	approx: finalRule{
		first: rules{
			{
				pattern: []rune("h"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[gdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("b"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("v"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("g"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("d"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("dZ"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("tS"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jnm"),
				phoneticRules: []token{
					{
						text:  []rune("jm"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("I"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[aAB]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[AB]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("A"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("A"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("B"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("l"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("l"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("m"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("m"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("n"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("r"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("r"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("z"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bp]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("m"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("kAg"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("kOg"),
						langs: -1,
					},
					{
						text:  []rune("kO"),
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("kOg"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("kAg"),
						langs: -1,
					},
					{
						text:  []rune("kA"),
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("kog"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("kog"),
						langs: -1,
					},
					{
						text:  []rune("ko"),
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("kag"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("kag"),
						langs: -1,
					},
					{
						text:  []rune("ka"),
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("h"),
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("H"),
				phoneticRules: []token{
					{
						text:  []rune("x"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				phoneticRules: []token{
					{
						text:  []rune("e"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				phoneticRules: []token{
					{
						text:  []rune("e"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				phoneticRules: []token{
					{
						text:  []rune("a"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				phoneticRules: []token{
					{
						text:  []rune("a"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("e"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("i"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("E"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("I"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Q"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Y"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("e"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("e"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("i"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("E"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("E"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("I"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("I"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("F"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Q"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("Q"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Y"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("Y"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEs"),
				phoneticRules: []token{
					{
						text:  []rune("lEs"),
						langs: -1,
					},
					{
						text:  []rune("lz"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lE"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("lE"),
						langs: -1,
					},
					{
						text:  []rune("l"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aue"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oue"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AvE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("AvE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ave"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("Ave"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("avE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("avE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ave"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ave"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OvE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("OvE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ove"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("Ove"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ovE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ovE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ove"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ove"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ea"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ea"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("EA"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ea"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("Ea"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("eA"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("iji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ije"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Iji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ije"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ajo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ojo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ejo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ujo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ija"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ija"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ijo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("j"),
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lYndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lander"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lAndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lAnder"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("landEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lender"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lendEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEnder"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bUrk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("burk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bUrg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("burg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dS"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dZ"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dZ"),
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
		},
		second: map[languageID]rules{
			languageID(ashany): rules{
				{
					pattern: []rune("b"),
					phoneticRules: []token{
						{
							text:  []rune("b"),
							langs: -1,
						},
						{
							text:  []rune("v"),
							langs: 1024,
						},
					},
				},
				{
					pattern: []rune("J"),
					phoneticRules: []token{
						{
							text:  []rune("z"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: 128,
						},
						{
							text:  []rune("im"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: 128,
						},
						{
							text:  []rune("in"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: 128,
						},
						{
							text:  []rune("om"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: 128,
						},
						{
							text:  []rune("on"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: 16,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("D"),
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: 16,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("li"),
							langs: -1,
						},
						{
							text:  []rune("il"),
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("li"),
							langs: -1,
						},
						{
							text:  []rune("il"),
							langs: 4,
						},
						{
							text:  []rune("lY"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ei"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ei"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iA"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ia"),
							langs: -1,
						},
						{
							text:  []rune("io"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iA"),
					phoneticRules: []token{
						{
							text:  []rune("ia"),
							langs: -1,
						},
						{
							text:  []rune("io"),
							langs: -1,
						},
						{
							text:  []rune("iY"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
						{
							text:  []rune("D"),
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
						{
							text:  nil,
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
						{
							text:  nil,
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DoiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sUts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("suts"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Uts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: 16,
						},
					},
				},
			},
			languageID(ashrussian): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("om"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("on"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("en"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("Ym"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("En"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("Yn"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashcyrillic): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("om"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("on"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("en"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("Ym"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("En"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("Yn"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashenglish): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("D"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("il"),
							langs: -1,
						},
						{
							text:  []rune("li"),
							langs: -1,
						},
						{
							text:  []rune("lY"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiEuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashfrench): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashgerman): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiAEIOUouQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoAOUiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoAOUeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoeOUiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sUts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("suts"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashhebrew): rules{},
			languageID(ashhungarian): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashpolish): rules{
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashromanian): rules{
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashspanish): rules{
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
		},
	},
	exact: finalRule{
		first: rules{
			{
				pattern: []rune("h"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[gdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("b"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("v"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("g"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("d"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("dZ"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("tS"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jnm"),
				phoneticRules: []token{
					{
						text:  []rune("jm"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("I"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[aAB]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[AB]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("A"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("A"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("B"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("l"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("l"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("m"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("m"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("n"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("r"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("r"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("z"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("H"),
				phoneticRules: []token{
					{
						text:  []rune("h"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[^t]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgZd]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkst]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgzd]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("Z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("je"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jE"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
		},
		second: map[languageID]rules{
			languageID(ashany): rules{
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("J"),
					phoneticRules: []token{
						{
							text:  []rune("l"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashrussian): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashcyrillic): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashenglish): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashfrench): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashgerman): rules{
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("J"),
					phoneticRules: []token{
						{
							text:  []rune("l"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashhebrew): rules{},
			languageID(ashhungarian): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashpolish): rules{
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashromanian): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
			languageID(ashspanish): rules{
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
			},
		},
	},
}

var ashDiscards = []string{
	"ben",
	"bar",
	"ha",
}
