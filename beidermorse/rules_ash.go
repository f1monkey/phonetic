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
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "in",
					langs: 512,
				},
				{
					text:  "ina",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ina"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "in",
					langs: 512,
				},
				{
					text:  "ina",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lof",
					langs: 512,
				},
				{
					text:  "lef",
					langs: 512,
				},
				{
					text:  "lova",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lof",
					langs: 512,
				},
				{
					text:  "lef",
					langs: 512,
				},
				{
					text:  "lova",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "of",
					langs: 512,
				},
				{
					text:  "ova",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eva"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ef",
					langs: 512,
				},
				{
					text:  "eva",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aia"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("aja"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("aya"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("lowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lova",
					langs: -1,
				},
				{
					text:  "lof",
					langs: 128,
				},
				{
					text:  "l",
					langs: 128,
				},
				{
					text:  "el",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("kowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "kova",
					langs: -1,
				},
				{
					text:  "kof",
					langs: 128,
				},
				{
					text:  "k",
					langs: 128,
				},
				{
					text:  "ek",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("owa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ova",
					langs: -1,
				},
				{
					text:  "of",
					langs: 128,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lovna",
					langs: -1,
				},
				{
					text:  "levna",
					langs: -1,
				},
				{
					text:  "l",
					langs: 128,
				},
				{
					text:  "el",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("kowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "kovna",
					langs: -1,
				},
				{
					text:  "k",
					langs: 128,
				},
				{
					text:  "ek",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("owna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovna",
					langs: -1,
				},
				{
					text:  "",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("lówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "el",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("kówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "ek",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("ówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
				{
					text:  "i",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("rh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chsch"),
			phoneticRules: []token{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsch"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "sk",
					langs: 256,
				},
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "StS",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "StS",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ssh"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[äöü]"),
			},
			phoneticRules: []token{
				{
					text:  "sh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  "S",
					langs: 516,
				},
				{
					text:  "sh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: 516,
				},
				{
					text:  "kh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chs"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: 16,
				},
				{
					text:  "xs",
					langs: -1,
				},
				{
					text:  "tSs",
					langs: 516,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "k",
					langs: 256,
				},
				{
					text:  "tS",
					langs: 516,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 516,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "tsk",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("czy"),
			phoneticRules: []token{
				{
					text:  "tSi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cze"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: -1,
				},
				{
					text:  "tSF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciewicz"),
			phoneticRules: []token{
				{
					text:  "tsevitS",
					langs: -1,
				},
				{
					text:  "tSevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siewicz"),
			phoneticRules: []token{
				{
					text:  "sevitS",
					langs: -1,
				},
				{
					text:  "SevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziewicz"),
			phoneticRules: []token{
				{
					text:  "zevitS",
					langs: -1,
				},
				{
					text:  "ZevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("riewicz"),
			phoneticRules: []token{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("diewicz"),
			phoneticRules: []token{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tiewicz"),
			phoneticRules: []token{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iewicz"),
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewicz"),
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owicz"),
			phoneticRules: []token{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("icz"),
			phoneticRules: []token{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cz"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tSB",
					langs: 128,
				},
				{
					text:  "tsB",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			phoneticRules: []token{
				{
					text:  "tSa",
					langs: 128,
				},
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "tSom",
					langs: 128,
				},
				{
					text:  "tsom",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			phoneticRules: []token{
				{
					text:  "tSon",
					langs: 128,
				},
				{
					text:  "tson",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "tSem",
					langs: 128,
				},
				{
					text:  "tsem",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			phoneticRules: []token{
				{
					text:  "tSen",
					langs: 128,
				},
				{
					text:  "tsen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tSF",
					langs: 128,
				},
				{
					text:  "tsF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: 128,
				},
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cio"),
			phoneticRules: []token{
				{
					text:  "tSo",
					langs: 128,
				},
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciu"),
			phoneticRules: []token{
				{
					text:  "tSu",
					langs: 128,
				},
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "tsi",
					langs: 128,
				},
				{
					text:  "tSi",
					langs: 384,
				},
				{
					text:  "tS",
					langs: 256,
				},
				{
					text:  "si",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  "tsi",
					langs: 128,
				},
				{
					text:  "tSi",
					langs: 384,
				},
				{
					text:  "si",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ce"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tsF",
					langs: 128,
				},
				{
					text:  "tSe",
					langs: 384,
				},
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ce"),
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: 384,
				},
				{
					text:  "tse",
					langs: 128,
				},
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cy"),
			phoneticRules: []token{
				{
					text:  "si",
					langs: -1,
				},
				{
					text:  "tsi",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("ssz"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssp"),
			phoneticRules: []token{
				{
					text:  "Sp",
					langs: 16,
				},
				{
					text:  "sp",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			phoneticRules: []token{
				{
					text:  "Sp",
					langs: 16,
				},
				{
					text:  "sp",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sst"),
			phoneticRules: []token{
				{
					text:  "St",
					langs: 16,
				},
				{
					text:  "st",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			phoneticRules: []token{
				{
					text:  "St",
					langs: 16,
				},
				{
					text:  "st",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "SB",
					langs: 128,
				},
				{
					text:  "sB",
					langs: 128,
				},
				{
					text:  "sja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			phoneticRules: []token{
				{
					text:  "Sa",
					langs: 128,
				},
				{
					text:  "sja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Som",
					langs: 128,
				},
				{
					text:  "som",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			phoneticRules: []token{
				{
					text:  "Son",
					langs: 128,
				},
				{
					text:  "son",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Sem",
					langs: 128,
				},
				{
					text:  "sem",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			phoneticRules: []token{
				{
					text:  "Sen",
					langs: 128,
				},
				{
					text:  "sen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "SF",
					langs: 128,
				},
				{
					text:  "sF",
					langs: -1,
				},
				{
					text:  "zi",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
				{
					text:  "Se",
					langs: 128,
				},
				{
					text:  "zi",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  "So",
					langs: 128,
				},
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siu"),
			phoneticRules: []token{
				{
					text:  "Su",
					langs: 128,
				},
				{
					text:  "sju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("si"),
			phoneticRules: []token{
				{
					text:  "Si",
					langs: 128,
				},
				{
					text:  "si",
					langs: -1,
				},
				{
					text:  "zi",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouäöë]"),
			},
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "z",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("gue"),
			phoneticRules: []token{
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gui"),
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("guy"),
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: 256,
				},
				{
					text:  "gh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gauz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "haus",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gaus"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "haus",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'ts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("golts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goltz"),
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'ts"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("golts"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'tz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goltz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gendler"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hendler",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejmer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejm"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geymer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geym"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geimer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geim"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gof"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hof",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ger"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ger",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gen"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "gen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gin"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "gin",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gie"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ge",
					langs: -1,
				},
				{
					text:  "gi",
					langs: 16,
				},
				{
					text:  "ji",
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("gie"),
			phoneticRules: []token{
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ge"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  "gE",
					langs: -1,
				},
				{
					text:  "xe",
					langs: 1024,
				},
				{
					text:  "dZe",
					langs: 260,
				},
			},
		},
		{
			pattern: []rune("gi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "xi",
					langs: 1024,
				},
				{
					text:  "dZi",
					langs: 260,
				},
			},
		},
		{
			pattern: []rune("ge"),
			phoneticRules: []token{
				{
					text:  "gE",
					langs: -1,
				},
				{
					text:  "dZe",
					langs: 260,
				},
				{
					text:  "hE",
					langs: 512,
				},
				{
					text:  "xe",
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "dZi",
					langs: 260,
				},
				{
					text:  "hI",
					langs: 512,
				},
				{
					text:  "xi",
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("gy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "dj",
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("gy"),
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "d",
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[jyaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouyei]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouei]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "h",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "eZ",
					langs: 264,
				},
				{
					text:  "ex",
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ly"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "lj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("li"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "lj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "lj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("lyo"),
			phoneticRules: []token{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ll"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "J",
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("j"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aoeiuy]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: 4,
				},
				{
					text:  "x",
					langs: 1024,
				},
				{
					text:  "Z",
					langs: 264,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "x",
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  "pf",
					langs: -1,
				},
				{
					text:  "p",
					langs: -1,
				},
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  "ph",
					langs: -1,
				},
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "kv",
					langs: 16,
				},
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phoneticRules: []token{
				{
					text:  "Se",
					langs: 128,
				},
				{
					text:  "re",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			phoneticRules: []token{
				{
					text:  "rze",
					langs: -1,
				},
				{
					text:  "rtsE",
					langs: 16,
				},
				{
					text:  "Ze",
					langs: 128,
				},
				{
					text:  "re",
					langs: 128,
				},
				{
					text:  "rZe",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phoneticRules: []token{
				{
					text:  "Si",
					langs: 128,
				},
				{
					text:  "ri",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: 128,
				},
				{
					text:  "ri",
					langs: 128,
				},
				{
					text:  "rZi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phoneticRules: []token{
				{
					text:  "S",
					langs: 128,
				},
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			phoneticRules: []token{
				{
					text:  "rz",
					langs: -1,
				},
				{
					text:  "rts",
					langs: 16,
				},
				{
					text:  "Z",
					langs: 128,
				},
				{
					text:  "r",
					langs: 128,
				},
				{
					text:  "rZ",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 20,
				},
			},
		},
		{
			pattern: []rune("tz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 20,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: 532,
				},
				{
					text:  "tz",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zh"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
				{
					text:  "zh",
					langs: 128,
				},
				{
					text:  "tsh",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("zia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "ZB",
					langs: 128,
				},
				{
					text:  "zB",
					langs: 128,
				},
				{
					text:  "zja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  "Za",
					langs: 128,
				},
				{
					text:  "zja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Zom",
					langs: 128,
				},
				{
					text:  "zom",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			phoneticRules: []token{
				{
					text:  "Zon",
					langs: 128,
				},
				{
					text:  "zon",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Zem",
					langs: 128,
				},
				{
					text:  "zem",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			phoneticRules: []token{
				{
					text:  "Zen",
					langs: 128,
				},
				{
					text:  "zen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "ZF",
					langs: 128,
				},
				{
					text:  "zF",
					langs: 128,
				},
				{
					text:  "ze",
					langs: -1,
				},
				{
					text:  "tsi",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
				{
					text:  "Ze",
					langs: 128,
				},
				{
					text:  "tsi",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  "Zo",
					langs: 128,
				},
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  "Zu",
					langs: 128,
				},
				{
					text:  "zju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zi"),
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: 128,
				},
				{
					text:  "zi",
					langs: -1,
				},
				{
					text:  "tsi",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("thal"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "tal",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  "t",
					langs: 16,
				},
				{
					text:  "th",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vogel"),
			phoneticRules: []token{
				{
					text:  "vogel",
					langs: -1,
				},
				{
					text:  "fogel",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("v"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "f",
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
				{
					text:  "x",
					langs: 384,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
				{
					text:  "H",
					langs: 20,
				},
			},
		},
		{
			pattern: []rune("yi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile(" $"),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "in",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "",
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oue"),
			phoneticRules: []token{
				{
					text:  "oue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  "au",
					langs: -1,
				},
				{
					text:  "o",
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  "ou",
					langs: -1,
				},
				{
					text:  "u",
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "uje",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ae"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: 16,
				},
				{
					text:  "aje",
					langs: 512,
				},
				{
					text:  "ae",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: 16,
				},
				{
					text:  "oje",
					langs: 512,
				},
				{
					text:  "oe",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: 4,
				},
				{
					text:  "aje",
					langs: 512,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: 16,
				},
				{
					text:  "oj",
					langs: 16,
				},
				{
					text:  "eu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: 16,
				},
				{
					text:  "e",
					langs: 128,
				},
				{
					text:  "ije",
					langs: 512,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: 16,
				},
				{
					text:  "e",
					langs: 128,
				},
				{
					text:  "ije",
					langs: 512,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ye"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
				{
					text:  "ije",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  "jo",
					langs: -1,
				},
				{
					text:  "e",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("yo"),
			phoneticRules: []token{
				{
					text:  "jo",
					langs: -1,
				},
				{
					text:  "e",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  "ea",
					langs: -1,
				},
				{
					text:  "ja",
					langs: 256,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "je",
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: 4,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uu"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: 128,
				},
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ł"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ń"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "nj",
					langs: 1024,
				},
			},
		},
		{
			pattern: []rune("ś"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: 128,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ţ"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ż"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ź"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: 128,
				},
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("où"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "om",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			phoneticRules: []token{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ä"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ă"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: 256,
				},
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			phoneticRules: []token{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("î"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: 128,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ß"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("\""),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "A",
					langs: -1,
				},
				{
					text:  "B",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
				{
					text:  "F",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phoneticRules: []token{
				{
					text:  "O",
					langs: -1,
				},
				{
					text:  "P",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: 16,
				},
				{
					text:  "z",
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
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цю"),
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("циа"),
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цие"),
			phoneticRules: []token{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цио"),
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("циу"),
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("сие"),
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("сио"),
			phoneticRules: []token{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("зие"),
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("зио"),
			phoneticRules: []token{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гауз"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "haus",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гаус"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "haus",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гольц"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("геймер"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гейм"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гоф"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hof",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гер"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ger",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ген"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "gen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гин"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "gin",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(й|ё|я|ю|ы|а|е|о|и|у)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ля"),
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лю"),
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лё"),
			phoneticRules: []token{
				{
					text:  "le",
					langs: -1,
				},
				{
					text:  "lo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лио"),
			phoneticRules: []token{
				{
					text:  "le",
					langs: -1,
				},
				{
					text:  "lo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ле"),
			phoneticRules: []token{
				{
					text:  "lE",
					langs: -1,
				},
				{
					text:  "lo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ийе"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ие"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ыйе"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ые"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ий"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ый"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ий"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ый"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ё"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "jo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "jaj",
					langs: -1,
				},
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(а|е|о|у)$"),
			},
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("эй"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ауе"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ауэ"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("а"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("б"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("в"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("д"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ж"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("з"),
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("и"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("й"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("к"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("л"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("м"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("н"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("о"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("п"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("р"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("с"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "с",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("с"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("т"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("у"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ф"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("х"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ц"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ч"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ш"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("щ"),
			phoneticRules: []token{
				{
					text:  "StS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ъ"),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ы"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ь"),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("э"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ю"),
			phoneticRules: []token{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("я"),
			phoneticRules: []token{
				{
					text:  "ja",
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
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "c",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "f",
					langs: -1,
				},
				{
					text:  "w",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			phoneticRules: []token{
				{
					text:  "gn",
					langs: -1,
				},
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "sk",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("who"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "w",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kn"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mb"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ng"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "N",
					langs: -1,
				},
				{
					text:  "ng",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pn"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "pn",
					langs: -1,
				},
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ps"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ps",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "kw",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tia"),
			phoneticRules: []token{
				{
					text:  "So",
					langs: -1,
				},
				{
					text:  "Sa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tio"),
			phoneticRules: []token{
				{
					text:  "So",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wr"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "w",
					langs: -1,
				},
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile(" $"),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
				{
					text:  "oue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "o",
					langs: -1,
				},
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ear"),
			phoneticRules: []token{
				{
					text:  "ia",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oa"),
			phoneticRules: []token{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oy"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  "ju",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "r",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
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
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "gn",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gue"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("que"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouyéèê]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ouh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "uh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  "vo",
					langs: -1,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eau"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("où"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
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
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chsch"),
			phoneticRules: []token{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("evitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ovitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("witsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chs"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "Sp",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "St",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssp"),
			phoneticRules: []token{
				{
					text:  "Sp",
					langs: -1,
				},
				{
					text:  "sp",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			phoneticRules: []token{
				{
					text:  "Sp",
					langs: -1,
				},
				{
					text:  "sp",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sst"),
			phoneticRules: []token{
				{
					text:  "St",
					langs: -1,
				},
				{
					text:  "st",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			phoneticRules: []token{
				{
					text:  "St",
					langs: -1,
				},
				{
					text:  "st",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  "pf",
					langs: -1,
				},
				{
					text:  "p",
					langs: -1,
				},
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  "ph",
					langs: -1,
				},
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "kv",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "evits",
					langs: -1,
				},
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "evits",
					langs: -1,
				},
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("evitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "evits",
					langs: -1,
				},
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eviz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "evits",
					langs: -1,
				},
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovits",
					langs: -1,
				},
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovits",
					langs: -1,
				},
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ovitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovits",
					langs: -1,
				},
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oviz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovits",
					langs: -1,
				},
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("witz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "vits",
					langs: -1,
				},
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "vits",
					langs: -1,
				},
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "vits",
					langs: -1,
				},
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("viz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "vits",
					langs: -1,
				},
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("thal"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "tal",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "th",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[äöüaeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouyäöüj]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouyäöü]"),
			},
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ß"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ae"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ä"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ã"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "ts",
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
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("עי"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("עו"),
			phoneticRules: []token{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("או"),
			phoneticRules: []token{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ג׳"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ד׳"),
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("א"),
			phoneticRules: []token{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ב"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ג"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ד"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("וו"),
			phoneticRules: []token{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("וי"),
			phoneticRules: []token{
				{
					text:  "WW",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ו"),
			phoneticRules: []token{
				{
					text:  "W",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ז"),
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ח"),
			phoneticRules: []token{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ט"),
			phoneticRules: []token{
				{
					text:  "T",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("יי"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("י"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ך"),
			phoneticRules: []token{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("כ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("כ"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ל"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ם"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("מ"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ן"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("נ"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ס"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ע"),
			phoneticRules: []token{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ף"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("פ"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ץ"),
			phoneticRules: []token{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("צ"),
			phoneticRules: []token{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ק"),
			phoneticRules: []token{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ר"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ש"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ת"),
			phoneticRules: []token{
				{
					text:  "TB",
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
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zs"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cs"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aj"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[áo]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[áo]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ely"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "eli",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ly"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "li",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gy"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  "nj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "ni",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ty"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  "tj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ty"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "ti",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "ku",
					langs: -1,
				},
				{
					text:  "kv",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
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
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ski",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cka"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "tski",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lova",
					langs: -1,
				},
				{
					text:  "lof",
					langs: -1,
				},
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "el",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "kova",
					langs: -1,
				},
				{
					text:  "kof",
					langs: -1,
				},
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "ek",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ova",
					langs: -1,
				},
				{
					text:  "of",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lovna",
					langs: -1,
				},
				{
					text:  "levna",
					langs: -1,
				},
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "el",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "kovna",
					langs: -1,
				},
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "ek",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ovna",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "el",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "ek",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("czy"),
			phoneticRules: []token{
				{
					text:  "tSi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cze"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: -1,
				},
				{
					text:  "tSF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciewicz"),
			phoneticRules: []token{
				{
					text:  "tsevitS",
					langs: -1,
				},
				{
					text:  "tSevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siewicz"),
			phoneticRules: []token{
				{
					text:  "sevitS",
					langs: -1,
				},
				{
					text:  "SevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziewicz"),
			phoneticRules: []token{
				{
					text:  "zevitS",
					langs: -1,
				},
				{
					text:  "ZevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("riewicz"),
			phoneticRules: []token{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("diewicz"),
			phoneticRules: []token{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tiewicz"),
			phoneticRules: []token{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iewicz"),
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewicz"),
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owicz"),
			phoneticRules: []token{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("icz"),
			phoneticRules: []token{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cz"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tSB",
					langs: -1,
				},
				{
					text:  "tsB",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			phoneticRules: []token{
				{
					text:  "tSa",
					langs: -1,
				},
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "tSom",
					langs: -1,
				},
				{
					text:  "tsom",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			phoneticRules: []token{
				{
					text:  "tSon",
					langs: -1,
				},
				{
					text:  "tson",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "tSem",
					langs: -1,
				},
				{
					text:  "tsem",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			phoneticRules: []token{
				{
					text:  "tSen",
					langs: -1,
				},
				{
					text:  "tsen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "tSF",
					langs: -1,
				},
				{
					text:  "tsF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: -1,
				},
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cio"),
			phoneticRules: []token{
				{
					text:  "tSo",
					langs: -1,
				},
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciu"),
			phoneticRules: []token{
				{
					text:  "tSu",
					langs: -1,
				},
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  "tSi",
					langs: -1,
				},
				{
					text:  "tsI",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssz"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "SB",
					langs: -1,
				},
				{
					text:  "sB",
					langs: -1,
				},
				{
					text:  "sja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			phoneticRules: []token{
				{
					text:  "Sa",
					langs: -1,
				},
				{
					text:  "sja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Som",
					langs: -1,
				},
				{
					text:  "som",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			phoneticRules: []token{
				{
					text:  "Son",
					langs: -1,
				},
				{
					text:  "son",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Sem",
					langs: -1,
				},
				{
					text:  "sem",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			phoneticRules: []token{
				{
					text:  "Sen",
					langs: -1,
				},
				{
					text:  "sen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "SF",
					langs: -1,
				},
				{
					text:  "sF",
					langs: -1,
				},
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  "Se",
					langs: -1,
				},
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  "So",
					langs: -1,
				},
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siu"),
			phoneticRules: []token{
				{
					text:  "Su",
					langs: -1,
				},
				{
					text:  "sju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("si"),
			phoneticRules: []token{
				{
					text:  "Si",
					langs: -1,
				},
				{
					text:  "sI",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ś"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "ZB",
					langs: -1,
				},
				{
					text:  "zB",
					langs: -1,
				},
				{
					text:  "zja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  "Za",
					langs: -1,
				},
				{
					text:  "zja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Zom",
					langs: -1,
				},
				{
					text:  "zom",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			phoneticRules: []token{
				{
					text:  "Zon",
					langs: -1,
				},
				{
					text:  "zon",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "Zem",
					langs: -1,
				},
				{
					text:  "zem",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			phoneticRules: []token{
				{
					text:  "Zen",
					langs: -1,
				},
				{
					text:  "zen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "ZF",
					langs: -1,
				},
				{
					text:  "zF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  "Zo",
					langs: -1,
				},
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  "Zu",
					langs: -1,
				},
				{
					text:  "zju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zi"),
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: -1,
				},
				{
					text:  "zI",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
				{
					text:  "ZF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
				{
					text:  "ZF",
					langs: -1,
				},
				{
					text:  "ze",
					langs: -1,
				},
				{
					text:  "zF",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("źe"),
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ży"),
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("źi"),
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: -1,
				},
				{
					text:  "zi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ż"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ź"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phoneticRules: []token{
				{
					text:  "Se",
					langs: -1,
				},
				{
					text:  "re",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
				{
					text:  "re",
					langs: -1,
				},
				{
					text:  "rZe",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phoneticRules: []token{
				{
					text:  "Si",
					langs: -1,
				},
				{
					text:  "ri",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: -1,
				},
				{
					text:  "ri",
					langs: -1,
				},
				{
					text:  "rZi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
				{
					text:  "r",
					langs: -1,
				},
				{
					text:  "rZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ł"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ń"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "s",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "om",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			phoneticRules: []token{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			phoneticRules: []token{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ije"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yje"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iie"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yie"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iye"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yye"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rie"),
			phoneticRules: []token{
				{
					text:  "rje",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("die"),
			phoneticRules: []token{
				{
					text:  "dje",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tie"),
			phoneticRules: []token{
				{
					text:  "tje",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "F",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aj"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "B",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
				{
					text:  "F",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phoneticRules: []token{
				{
					text:  "P",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
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
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ce"),
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  "tSi",
					langs: -1,
				},
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  "dZi",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ţ"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("î"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ă"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
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
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "in",
					langs: -1,
				},
				{
					text:  "ina",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ina"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "in",
					langs: -1,
				},
				{
					text:  "ina",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lof",
					langs: -1,
				},
				{
					text:  "lef",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "lof",
					langs: -1,
				},
				{
					text:  "lef",
					langs: -1,
				},
				{
					text:  "lova",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "of",
					langs: -1,
				},
				{
					text:  "ova",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eva"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ef",
					langs: -1,
				},
				{
					text:  "ova",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aia"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aja"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aya"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsya"),
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsyu"),
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsia"),
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsie"),
			phoneticRules: []token{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsio"),
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsye"),
			phoneticRules: []token{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsyo"),
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsiu"),
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sye"),
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("syo"),
			phoneticRules: []token{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zye"),
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zyo"),
			phoneticRules: []token{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gauz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "haus",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gaus"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "haus",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'ts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("golts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gol'tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goltz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejmer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gejm"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geimer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geim"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geymer"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajmer",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("geym"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hajm",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gendler"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hendler",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gof"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hof",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gojf"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hojf",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goyf"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hojf",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("goif"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "hojf",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ger"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ger",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gen"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "gen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gin"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "gin",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kog"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  "kog",
					langs: -1,
				},
				{
					text:  "koh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kag"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  "kag",
					langs: -1,
				},
				{
					text:  "kah",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[jaeoiuy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  "StS",
					langs: -1,
				},
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssh"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zh"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
				{
					text:  "tz",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "kv",
					langs: -1,
				},
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "s",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lya"),
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lyu"),
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lia"),
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liu"),
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lja"),
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lju"),
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("le"),
			phoneticRules: []token{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "lE",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lyo"),
			phoneticRules: []token{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ije"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iye"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iie"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yje"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ye"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yye"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yie"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  "jo",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yo"),
			phoneticRules: []token{
				{
					text:  "jo",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aiou]$"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  "aje",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  "oo",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("\""),
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
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
					text:  "n",
					langs: -1,
				},
				{
					text:  "nj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ll"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bpvf]"),
			},
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "gv",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  "vo",
					langs: -1,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
	},
}

var ashLangRules = []langRule{
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "zh",
			prefix:           "",
			suffix:           "",
		},
		langs:  660,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eau",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[aoeiuäöü]h"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "vogel",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "vogel",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "witz",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "tz",
		},
		langs:  532,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "tz",
			suffix:           "",
		},
		langs:  516,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "güe",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "güi",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ghe",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ghi",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "vici",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "schi",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "chsch",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "tsch",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ssch",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "sch",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "sch",
			suffix:           "",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "rz",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "rz",
		},
		langs:  144,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[^aoeiuäöü]rz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("rz[^aoeiuäöü]"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cki",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ska",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cka",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ue",
			prefix:           "",
			suffix:           "",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ae",
			prefix:           "",
			suffix:           "",
		},
		langs:  532,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "oe",
			prefix:           "",
			suffix:           "",
		},
		langs:  540,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "th",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "th",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("th[^aoeiu]"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "mann",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cz",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cy",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "niew",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "stein",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "heim",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "heimer",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ii",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "iy",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yy",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yi",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yj",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ij",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gaus",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gauz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gauz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "goltz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gol'tz$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "golts",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gol'ts$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "goltz",
			suffix:           "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("^gol'tz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "golts",
			suffix:           "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("^gol'ts"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gendler",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gejmer",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gejm",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "geimer",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "geim",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "geymer",
			prefix:           "",
			suffix:           "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "geym",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gof",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "thal",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "zweig",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ck",
		},
		langs:  20,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "c",
		},
		langs:  448,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sz",
			prefix:           "",
			suffix:           "",
		},
		langs:  192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gue",
			prefix:           "",
			suffix:           "",
		},
		langs:  1032,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gui",
			prefix:           "",
			suffix:           "",
		},
		langs:  1032,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "guy",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cs",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "cs",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dzs",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "zs",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "zs",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "wl",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "wr",
			suffix:           "",
		},
		langs:  148,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gy",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gy[aeou]"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gy",
			prefix:           "",
			suffix:           "",
		},
		langs:  576,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ly",
			prefix:           "",
			suffix:           "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ny",
			prefix:           "",
			suffix:           "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ty",
			prefix:           "",
			suffix:           "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "â",
			prefix:           "",
			suffix:           "",
		},
		langs:  264,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ă",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "à",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ä",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "á",
			prefix:           "",
			suffix:           "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ą",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ć",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ç",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ę",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "é",
			prefix:           "",
			suffix:           "",
		},
		langs:  1096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "è",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ê",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "í",
			prefix:           "",
			suffix:           "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "î",
			prefix:           "",
			suffix:           "",
		},
		langs:  264,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ł",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ń",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ñ",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ó",
			prefix:           "",
			suffix:           "",
		},
		langs:  1216,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ö",
			prefix:           "",
			suffix:           "",
		},
		langs:  80,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "õ",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ş",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ś",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ţ",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ü",
			prefix:           "",
			suffix:           "",
		},
		langs:  80,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ù",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ű",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ú",
			prefix:           "",
			suffix:           "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ź",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ż",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ß",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "а",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ё",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "о",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "е",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "и",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "у",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ы",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "э",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ю",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "я",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "א",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ב",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ג",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ד",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ה",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ו",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ז",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ח",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ט",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "י",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "כ",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ל",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "מ",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "נ",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ס",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ע",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "פ",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "צ",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ק",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ר",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ש",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ת",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "a",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "o",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "e",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "i",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "y",
			prefix:           "",
			suffix:           "",
		},
		langs:  290,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "u",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("v[^aoeiuäüö]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("c[^aohk]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dzi",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ou",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "aj",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ej",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "oj",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "uj",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "k",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "v",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ky",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eu",
			prefix:           "",
			suffix:           "",
		},
		langs:  640,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "w",
			prefix:           "",
			suffix:           "",
		},
		langs:  1864,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "kie",
			prefix:           "",
			suffix:           "",
		},
		langs:  1032,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gie",
			prefix:           "",
			suffix:           "",
		},
		langs:  1288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "q",
			prefix:           "",
			suffix:           "",
		},
		langs:  960,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sch",
			prefix:           "",
			suffix:           "",
		},
		langs:  1224,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "h",
			suffix:           "",
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
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phoneticRules: []token{
					{
						text:  "p",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "p",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[gdZz]"),
				},
				phoneticRules: []token{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phoneticRules: []token{
					{
						text:  "f",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "f",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  "v",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  "k",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "k",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdZz]"),
				},
				phoneticRules: []token{
					{
						text:  "g",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  "t",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "t",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgZz]"),
				},
				phoneticRules: []token{
					{
						text:  "d",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "dZ",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "tS",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jnm"),
				phoneticRules: []token{
					{
						text:  "jm",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "I",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[aAB]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[AB]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("A"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "A",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "B",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("l"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "l",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("m"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "m",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "n",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("r"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "r",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "z",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bp]"),
				},
				phoneticRules: []token{
					{
						text:  "m",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("kAg"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  "kOg",
						langs: -1,
					},
					{
						text:  "kO",
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("kOg"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  "kAg",
						langs: -1,
					},
					{
						text:  "kA",
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("kog"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  "kog",
						langs: -1,
					},
					{
						text:  "ko",
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("kag"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phoneticRules: []token{
					{
						text:  "kag",
						langs: -1,
					},
					{
						text:  "ka",
						langs: 512,
					},
				},
			},
			{
				pattern: []rune("h"),
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("H"),
				phoneticRules: []token{
					{
						text:  "x",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				phoneticRules: []token{
					{
						text:  "e",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				phoneticRules: []token{
					{
						text:  "e",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				phoneticRules: []token{
					{
						text:  "a",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				phoneticRules: []token{
					{
						text:  "a",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("e"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("i"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("E"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("I"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Q"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Y"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("e"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "e",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("i"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("E"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "E",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("I"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "I",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("F"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "F",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Q"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "Q",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Y"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "Y",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEs"),
				phoneticRules: []token{
					{
						text:  "lEs",
						langs: -1,
					},
					{
						text:  "lz",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lE"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				phoneticRules: []token{
					{
						text:  "lE",
						langs: -1,
					},
					{
						text:  "l",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aue"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oue"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AvE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "AvE",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ave"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "Ave",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("avE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "avE",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ave"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "ave",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OvE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "OvE",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ove"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "Ove",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ovE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "ovE",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ove"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "ove",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ea"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "ea",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "EA",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ea"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "Ea",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
					{
						text:  "eA",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uje"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("iji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ije"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Iji"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjI"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ije"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjE"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ajo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ojo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ejo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uja"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ujo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ija"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ija"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjA"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ijo"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjO"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("j"),
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lYndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lander"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lAndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lAnder"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("landEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lender"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lendEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEnder"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "lYnder",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bUrk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "burk",
						langs: -1,
					},
					{
						text:  "berk",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("burk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "burk",
						langs: -1,
					},
					{
						text:  "berk",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bUrg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "burk",
						langs: -1,
					},
					{
						text:  "berk",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("burg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "burk",
						langs: -1,
					},
					{
						text:  "berk",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dS"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dZ"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				phoneticRules: []token{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dZ"),
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				phoneticRules: []token{
					{
						text:  "z",
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
							text:  "b",
							langs: -1,
						},
						{
							text:  "v",
							langs: 1024,
						},
					},
				},
				{
					pattern: []rune("J"),
					phoneticRules: []token{
						{
							text:  "z",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "om",
							langs: 128,
						},
						{
							text:  "im",
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "on",
							langs: 128,
						},
						{
							text:  "in",
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("UiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "im",
							langs: 128,
						},
						{
							text:  "om",
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "in",
							langs: 128,
						},
						{
							text:  "on",
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phoneticRules: []token{
						{
							text:  "Q",
							langs: 16,
						},
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "D",
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "Q",
							langs: 16,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "li",
							langs: -1,
						},
						{
							text:  "il",
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  "li",
							langs: -1,
						},
						{
							text:  "il",
							langs: 4,
						},
						{
							text:  "lY",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ei"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ei"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iA"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ia",
							langs: -1,
						},
						{
							text:  "io",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iA"),
					phoneticRules: []token{
						{
							text:  "ia",
							langs: -1,
						},
						{
							text:  "io",
							langs: -1,
						},
						{
							text:  "iY",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
						{
							text:  "D",
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
						{
							text:  "",
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
						{
							text:  "",
							langs: 4,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DoiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "uk",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "uk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sUts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "suts",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: 16,
						},
					},
				},
				{
					pattern: []rune("Uts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "Q",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("om"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "om",
							langs: -1,
						},
						{
							text:  "im",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("on"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "on",
							langs: -1,
						},
						{
							text:  "in",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "im",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("en"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "in",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "im",
							langs: -1,
						},
						{
							text:  "Ym",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("En"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "in",
							langs: -1,
						},
						{
							text:  "Yn",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("om"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "om",
							langs: -1,
						},
						{
							text:  "im",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("on"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "on",
							langs: -1,
						},
						{
							text:  "in",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "im",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("en"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "in",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "im",
							langs: -1,
						},
						{
							text:  "Ym",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("En"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "in",
							langs: -1,
						},
						{
							text:  "Yn",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					phoneticRules: []token{
						{
							text:  "Q",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "D",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  "il",
							langs: -1,
						},
						{
							text:  "li",
							langs: -1,
						},
						{
							text:  "lY",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiEuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQY]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiAEIOUouQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "Q",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AU"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aU"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OU"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oU"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoAOUiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoAOUeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "Y",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoeOUiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "Y",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "uk",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "uk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sUts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "suts",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "Q",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
						{
							text:  "im",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
						{
							text:  "in",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "im",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "in",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
						{
							text:  "im",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
						{
							text:  "in",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dm",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "Dn",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "im",
							langs: -1,
						},
						{
							text:  "om",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "in",
							langs: -1,
						},
						{
							text:  "on",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "ik",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []token{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Q",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "Y",
							langs: -1,
						},
						{
							text:  "i",
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
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phoneticRules: []token{
					{
						text:  "p",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "p",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[gdZz]"),
				},
				phoneticRules: []token{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phoneticRules: []token{
					{
						text:  "f",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "f",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  "v",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  "k",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "k",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdZz]"),
				},
				phoneticRules: []token{
					{
						text:  "g",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  "t",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "t",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgZz]"),
				},
				phoneticRules: []token{
					{
						text:  "d",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "dZ",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "tS",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jnm"),
				phoneticRules: []token{
					{
						text:  "jm",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "I",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[aAB]"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[AB]$"),
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("A"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "A",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("B"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "B",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("l"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "l",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("m"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "m",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "n",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("r"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "r",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "z",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("H"),
				phoneticRules: []token{
					{
						text:  "h",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[^t]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgZd]"),
				},
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkst]"),
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "S",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgzd]"),
				},
				phoneticRules: []token{
					{
						text:  "Z",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  "j",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  "j",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("je"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  "j",
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jE"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  "j",
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
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("J"),
					phoneticRules: []token{
						{
							text:  "l",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("J"),
					phoneticRules: []token{
						{
							text:  "l",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  "i",
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
