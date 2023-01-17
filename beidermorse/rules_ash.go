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
			pattern: "yna",
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
			pattern: "ina",
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
			pattern: "liova",
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
			pattern: "lova",
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
			pattern: "ova",
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
			pattern: "eva",
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
			pattern: "aia",
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
			pattern: "aja",
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
			pattern: "aya",
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
			pattern: "lowa",
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
			pattern: "kowa",
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
			pattern: "owa",
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
			pattern: "lowna",
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
			pattern: "kowna",
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
			pattern: "owna",
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
			pattern: "lówna",
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
			pattern: "kówna",
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
			pattern: "ówna",
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
			pattern: "a",
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
			pattern: "rh",
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
			pattern: "ssch",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "chsch",
			phoneticRules: []token{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsch",
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "sch",
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
			pattern: "sch",
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
			pattern: "ssh",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sh",
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
			pattern: "sh",
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
			pattern: "sh",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "kh",
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
			pattern: "chs",
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
			pattern: "ch",
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
			pattern: "ch",
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
			pattern: "ck",
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
			pattern: "czy",
			phoneticRules: []token{
				{
					text:  "tSi",
					langs: -1,
				},
			},
		},
		{
			pattern: "cze",
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
			pattern: "ciewicz",
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
			pattern: "siewicz",
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
			pattern: "ziewicz",
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
			pattern: "riewicz",
			phoneticRules: []token{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "diewicz",
			phoneticRules: []token{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tiewicz",
			phoneticRules: []token{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "iewicz",
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ewicz",
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "owicz",
			phoneticRules: []token{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "icz",
			phoneticRules: []token{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern: "cz",
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "cia",
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
			pattern: "cia",
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
			pattern: "cią",
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
			pattern: "cią",
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
			pattern: "cię",
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
			pattern: "cię",
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
			pattern: "cie",
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
			pattern: "cie",
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
			pattern: "cio",
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
			pattern: "ciu",
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
			pattern: "ci",
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
			pattern: "ci",
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
			pattern: "ce",
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
			pattern: "ce",
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
			pattern: "cy",
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
			pattern: "ssz",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sz",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "ssp",
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
			pattern: "sp",
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
			pattern: "sst",
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
			pattern: "st",
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
			pattern: "ss",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "sia",
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
			pattern: "sia",
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
			pattern: "sią",
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
			pattern: "sią",
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
			pattern: "się",
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
			pattern: "się",
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
			pattern: "sie",
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
			pattern: "sie",
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
			pattern: "sio",
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
			pattern: "siu",
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
			pattern: "si",
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
			pattern: "s",
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
			pattern: "gue",
			phoneticRules: []token{
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern: "gui",
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: "guy",
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: "gh",
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
			pattern: "gauz",
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
			pattern: "gaus",
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
			pattern: "gol'ts",
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
			pattern: "golts",
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
			pattern: "gol'tz",
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
			pattern: "goltz",
			phoneticRules: []token{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: "gol'ts",
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
			pattern: "golts",
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
			pattern: "gol'tz",
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
			pattern: "goltz",
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
			pattern: "gendler",
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
			pattern: "gejmer",
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
			pattern: "gejm",
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
			pattern: "geymer",
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
			pattern: "geym",
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
			pattern: "geimer",
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
			pattern: "geim",
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
			pattern: "gof",
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
			pattern: "ger",
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
			pattern: "gen",
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
			pattern: "gin",
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
			pattern: "gie",
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
			pattern: "gie",
			phoneticRules: []token{
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern: "ge",
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
			pattern: "gi",
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
			pattern: "ge",
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
			pattern: "gi",
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
			pattern: "gy",
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
			pattern: "gy",
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
			pattern: "g",
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
			pattern: "g",
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
			pattern: "ej",
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
			pattern: "ej",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ly",
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
			pattern: "li",
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
			pattern: "lj",
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
			pattern: "lio",
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
			pattern: "lyo",
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
			pattern: "ll",
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
			pattern: "j",
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
			pattern: "j",
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
			pattern: "pf",
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
			pattern: "ph",
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
			pattern: "qu",
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
			pattern: "rze",
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
			pattern: "rze",
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
			pattern: "rzy",
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
			pattern: "rzy",
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
			pattern: "rz",
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
			pattern: "rz",
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
			pattern: "tz",
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
			pattern: "tz",
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
			pattern: "tz",
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
			pattern: "zh",
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
			pattern: "zia",
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
			pattern: "zia",
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
			pattern: "zią",
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
			pattern: "zią",
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
			pattern: "zię",
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
			pattern: "zię",
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
			pattern: "zie",
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
			pattern: "zie",
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
			pattern: "zio",
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
			pattern: "ziu",
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
			pattern: "zi",
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
			pattern: "thal",
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
			pattern: "th",
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
			pattern: "th",
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
			pattern: "th",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "vogel",
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
			pattern: "v",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "yi",
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
			pattern: "ii",
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
			pattern: "iy",
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
			pattern: "yy",
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
			pattern: "e",
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
			pattern: "yj",
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
			pattern: "ij",
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
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "oue",
			phoneticRules: []token{
				{
					text:  "oue",
					langs: -1,
				},
			},
		},
		{
			pattern: "au",
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
			pattern: "ou",
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
			pattern: "ue",
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
			pattern: "ae",
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
			pattern: "oe",
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
			pattern: "ee",
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
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "eu",
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
			pattern: "i",
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
			pattern: "y",
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
			pattern: "ie",
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
			pattern: "ie",
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
			pattern: "ye",
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
			pattern: "i",
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
			pattern: "y",
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
			pattern: "io",
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
			pattern: "yo",
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
			pattern: "ea",
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
			pattern: "e",
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
			pattern: "oo",
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
			pattern: "uu",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ć",
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
			pattern: "ł",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "ń",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ñ",
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
			pattern: "ś",
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
			pattern: "ş",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "ţ",
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "ż",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ź",
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
			pattern: "où",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ą",
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
			pattern: "ą",
			phoneticRules: []token{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern: "ä",
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
			pattern: "á",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "ă",
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
			pattern: "à",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "â",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "è",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "ê",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "ę",
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
			pattern: "ę",
			phoneticRules: []token{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern: "í",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "î",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ö",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ő",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
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
			pattern: "ű",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ü",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ú",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ű",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ß",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "'",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "\"",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
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
			pattern: "e",
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
			pattern: "o",
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
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
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
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "ця",
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "цю",
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "циа",
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "цие",
			phoneticRules: []token{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: "цио",
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "циу",
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "сие",
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: "сио",
			phoneticRules: []token{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: "зие",
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "зио",
			phoneticRules: []token{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: "гауз",
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
			pattern: "гаус",
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
			pattern: "гольц",
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
			pattern: "геймер",
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
			pattern: "гейм",
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
			pattern: "гоф",
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
			pattern: "гер",
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
			pattern: "ген",
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
			pattern: "гин",
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
			pattern: "г",
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
			pattern: "г",
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
			pattern: "ля",
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "лю",
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "лё",
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
			pattern: "лио",
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
			pattern: "ле",
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
			pattern: "ийе",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ие",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ыйе",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ые",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ий",
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
			pattern: "ый",
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
			pattern: "ий",
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
			pattern: "ый",
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
			pattern: "ё",
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
			pattern: "ей",
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
			pattern: "е",
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
			pattern: "е",
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
			pattern: "эй",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ей",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ауе",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "ауэ",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "а",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "б",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "в",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "г",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "д",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "е",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "ж",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "з",
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: "и",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "й",
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "к",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "л",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "м",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "н",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "о",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "п",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "р",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "с",
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
			pattern: "с",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "т",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "у",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ф",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "х",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ц",
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "ч",
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ш",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "щ",
			phoneticRules: []token{
				{
					text:  "StS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ъ",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ы",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "ь",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "э",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "ю",
			phoneticRules: []token{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: "я",
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
			pattern: "tch",
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
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
			pattern: "ck",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "cc",
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
			pattern: "c",
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
			pattern: "c",
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
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "gh",
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
			pattern: "gh",
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
			pattern: "gn",
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
			pattern: "g",
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
			pattern: "th",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "kh",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ph",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "sch",
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
			pattern: "sh",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "who",
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
			pattern: "wh",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "kn",
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
			pattern: "mb",
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
			pattern: "ng",
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
			pattern: "pn",
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
			pattern: "ps",
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
			pattern: "qu",
			phoneticRules: []token{
				{
					text:  "kw",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "tia",
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
			pattern: "tio",
			phoneticRules: []token{
				{
					text:  "So",
					langs: -1,
				},
			},
		},
		{
			pattern: "wr",
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
			pattern: "w",
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
			pattern: "x",
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
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "yi",
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
			pattern: "y",
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
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "oue",
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
			pattern: "ai",
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
			pattern: "ay",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
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
			pattern: "a",
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
			pattern: "ei",
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
			pattern: "ey",
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
			pattern: "ear",
			phoneticRules: []token{
				{
					text:  "ia",
					langs: -1,
				},
			},
		},
		{
			pattern: "ea",
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
			pattern: "ee",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
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
			pattern: "e",
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
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
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
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "oa",
			phoneticRules: []token{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern: "oi",
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: "oo",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ou",
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
			pattern: "oy",
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
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
			pattern: "o",
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
			pattern: "u",
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
			pattern: "u",
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
			pattern: "u",
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
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "kh",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ph",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "ç",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
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
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "gn",
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
			pattern: "g",
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
			pattern: "gue",
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
			pattern: "gu",
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
			pattern: "que",
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
			pattern: "qu",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "ouh",
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
			pattern: "ou",
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
			pattern: "uo",
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
			pattern: "u",
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
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "eau",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ai",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ay",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "ê",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "è",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "à",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "â",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "où",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ou",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "oi",
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
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
			pattern: "e",
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
			pattern: "i",
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
			pattern: "y",
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
			pattern: "yi",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ii",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "yy",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "ziu",
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "zia",
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "zio",
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "ssch",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "chsch",
			phoneticRules: []token{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ewitsch",
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
			pattern: "owitsch",
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
			pattern: "evitsch",
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
			pattern: "ovitsch",
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
			pattern: "witsch",
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
			pattern: "vitsch",
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
			pattern: "sch",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "chs",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ck",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
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
			pattern: "sp",
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
			pattern: "st",
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
			pattern: "ssp",
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
			pattern: "sp",
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
			pattern: "sst",
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
			pattern: "st",
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
			pattern: "pf",
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
			pattern: "ph",
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
			pattern: "qu",
			phoneticRules: []token{
				{
					text:  "kv",
					langs: -1,
				},
			},
		},
		{
			pattern: "ewitz",
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
			pattern: "ewiz",
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
			pattern: "evitz",
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
			pattern: "eviz",
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
			pattern: "owitz",
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
			pattern: "owiz",
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
			pattern: "ovitz",
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
			pattern: "oviz",
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
			pattern: "witz",
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
			pattern: "wiz",
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
			pattern: "vitz",
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
			pattern: "viz",
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
			pattern: "tz",
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "thal",
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
			pattern: "th",
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
			pattern: "th",
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
			pattern: "th",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "rh",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "ss",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
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
			pattern: "s",
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
			pattern: "ß",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "ij",
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
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "ue",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ae",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "oe",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ü",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ä",
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
			pattern: "ö",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "eu",
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
			pattern: "i",
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
			pattern: "y",
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
			pattern: "ie",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
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
			pattern: "y",
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
			pattern: "ñ",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ã",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "ő",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ű",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ç",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
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
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "אי",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "עי",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "עו",
			phoneticRules: []token{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern: "או",
			phoneticRules: []token{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern: "ג׳",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ד׳",
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "א",
			phoneticRules: []token{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern: "ב",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "ג",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "ד",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "ה",
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
			pattern: "ה",
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
			pattern: "ה",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "וו",
			phoneticRules: []token{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern: "וי",
			phoneticRules: []token{
				{
					text:  "WW",
					langs: -1,
				},
			},
		},
		{
			pattern: "ו",
			phoneticRules: []token{
				{
					text:  "W",
					langs: -1,
				},
			},
		},
		{
			pattern: "ז",
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ח",
			phoneticRules: []token{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern: "ט",
			phoneticRules: []token{
				{
					text:  "T",
					langs: -1,
				},
			},
		},
		{
			pattern: "יי",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "י",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ך",
			phoneticRules: []token{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern: "כ",
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
			pattern: "כ",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "ל",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "ם",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "מ",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "ן",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "נ",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ס",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "ע",
			phoneticRules: []token{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern: "ף",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "פ",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "ץ",
			phoneticRules: []token{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern: "צ",
			phoneticRules: []token{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern: "ק",
			phoneticRules: []token{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern: "ר",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "ש",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "ת",
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
			pattern: "sz",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "zs",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "cs",
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ay",
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
			pattern: "ai",
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
			pattern: "aj",
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
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
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
			pattern: "i",
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
			pattern: "ee",
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
			pattern: "ely",
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
			pattern: "ly",
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
			pattern: "gy",
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
			pattern: "gy",
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
			pattern: "ny",
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
			pattern: "ny",
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
			pattern: "ty",
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
			pattern: "ty",
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
			pattern: "qu",
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
			pattern: "h",
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
			pattern: "á",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "í",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ö",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ő",
			phoneticRules: []token{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ú",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ü",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ű",
			phoneticRules: []token{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
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
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "ska",
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
			pattern: "cka",
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
			pattern: "lowa",
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
			pattern: "kowa",
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
			pattern: "owa",
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
			pattern: "lowna",
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
			pattern: "kowna",
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
			pattern: "owna",
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
			pattern: "lówna",
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
			pattern: "kówna",
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
			pattern: "ówna",
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
			pattern: "a",
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
			pattern: "czy",
			phoneticRules: []token{
				{
					text:  "tSi",
					langs: -1,
				},
			},
		},
		{
			pattern: "cze",
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
			pattern: "ciewicz",
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
			pattern: "siewicz",
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
			pattern: "ziewicz",
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
			pattern: "riewicz",
			phoneticRules: []token{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "diewicz",
			phoneticRules: []token{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tiewicz",
			phoneticRules: []token{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "iewicz",
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ewicz",
			phoneticRules: []token{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "owicz",
			phoneticRules: []token{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "icz",
			phoneticRules: []token{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern: "cz",
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "cia",
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
			pattern: "cia",
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
			pattern: "cią",
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
			pattern: "cią",
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
			pattern: "cię",
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
			pattern: "cię",
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
			pattern: "cie",
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
			pattern: "cie",
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
			pattern: "cio",
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
			pattern: "ciu",
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
			pattern: "ci",
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
			pattern: "ć",
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
			pattern: "ssz",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sz",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sia",
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
			pattern: "sia",
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
			pattern: "sią",
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
			pattern: "sią",
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
			pattern: "się",
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
			pattern: "się",
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
			pattern: "sie",
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
			pattern: "sie",
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
			pattern: "sio",
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
			pattern: "siu",
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
			pattern: "si",
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
			pattern: "ś",
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
			pattern: "zia",
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
			pattern: "zia",
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
			pattern: "zią",
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
			pattern: "zią",
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
			pattern: "zię",
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
			pattern: "zię",
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
			pattern: "zie",
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
			pattern: "zie",
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
			pattern: "zio",
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
			pattern: "ziu",
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
			pattern: "zi",
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
			pattern: "że",
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
			pattern: "że",
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
			pattern: "że",
			phoneticRules: []token{
				{
					text:  "Ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "źe",
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
			pattern: "ży",
			phoneticRules: []token{
				{
					text:  "Zi",
					langs: -1,
				},
			},
		},
		{
			pattern: "źi",
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
			pattern: "ż",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ź",
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
			pattern: "rze",
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
			pattern: "rze",
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
			pattern: "rzy",
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
			pattern: "rzy",
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
			pattern: "rz",
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
			pattern: "rz",
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
			pattern: "lio",
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
			pattern: "ł",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "ń",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
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
			pattern: "ó",
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
			pattern: "ą",
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
			pattern: "ę",
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
			pattern: "ą",
			phoneticRules: []token{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern: "ę",
			phoneticRules: []token{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern: "ije",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yje",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iie",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yie",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iye",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yye",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ij",
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
			pattern: "yj",
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
			pattern: "ii",
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
			pattern: "yi",
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
			pattern: "iy",
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
			pattern: "yy",
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
			pattern: "rie",
			phoneticRules: []token{
				{
					text:  "rje",
					langs: -1,
				},
			},
		},
		{
			pattern: "die",
			phoneticRules: []token{
				{
					text:  "dje",
					langs: -1,
				},
			},
		},
		{
			pattern: "tie",
			phoneticRules: []token{
				{
					text:  "tje",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
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
			pattern: "ie",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "au",
			phoneticRules: []token{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ej",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ai",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ay",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "aj",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
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
			pattern: "y",
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
			pattern: "i",
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
			pattern: "y",
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
			pattern: "a",
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
			pattern: "e",
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
			pattern: "o",
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
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
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
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ce",
			phoneticRules: []token{
				{
					text:  "tSe",
					langs: -1,
				},
			},
		},
		{
			pattern: "ci",
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
			pattern: "ch",
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
			pattern: "ch",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "gi",
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
			pattern: "g",
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
			pattern: "gh",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
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
			pattern: "i",
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
			pattern: "ţ",
			phoneticRules: []token{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "ş",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
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
			pattern: "qu",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "î",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ea",
			phoneticRules: []token{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern: "ă",
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
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "yna",
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
			pattern: "ina",
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
			pattern: "liova",
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
			pattern: "lova",
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
			pattern: "ova",
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
			pattern: "eva",
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
			pattern: "aia",
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
			pattern: "aja",
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
			pattern: "aya",
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
			pattern: "tsya",
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsyu",
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsia",
			phoneticRules: []token{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsie",
			phoneticRules: []token{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsio",
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsye",
			phoneticRules: []token{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsyo",
			phoneticRules: []token{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsiu",
			phoneticRules: []token{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "sie",
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: "sio",
			phoneticRules: []token{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: "zie",
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "zio",
			phoneticRules: []token{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: "sye",
			phoneticRules: []token{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: "syo",
			phoneticRules: []token{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: "zye",
			phoneticRules: []token{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "zyo",
			phoneticRules: []token{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: "gauz",
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
			pattern: "gaus",
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
			pattern: "gol'ts",
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
			pattern: "golts",
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
			pattern: "gol'tz",
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
			pattern: "goltz",
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
			pattern: "gejmer",
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
			pattern: "gejm",
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
			pattern: "geimer",
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
			pattern: "geim",
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
			pattern: "geymer",
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
			pattern: "geym",
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
			pattern: "gendler",
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
			pattern: "gof",
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
			pattern: "gojf",
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
			pattern: "goyf",
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
			pattern: "goif",
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
			pattern: "ger",
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
			pattern: "gen",
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
			pattern: "gin",
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
			pattern: "gg",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "kog",
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
			pattern: "kag",
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
			pattern: "g",
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
			pattern: "g",
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
			pattern: "kh",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
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
			pattern: "sch",
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
			pattern: "ssh",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sh",
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "zh",
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "tz",
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
			pattern: "tz",
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
			pattern: "c",
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
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
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
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
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
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "lya",
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "lyu",
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "lia",
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "liu",
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "lja",
			phoneticRules: []token{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "lju",
			phoneticRules: []token{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "le",
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
			pattern: "lyo",
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
			pattern: "lio",
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
			pattern: "ije",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iye",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iie",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yje",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ye",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yye",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yie",
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ij",
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
			pattern: "iy",
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
			pattern: "ii",
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
			pattern: "yj",
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
			pattern: "yy",
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
			pattern: "yi",
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
			pattern: "io",
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
			pattern: "i",
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
			pattern: "i",
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
			pattern: "ei",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ej",
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "yo",
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
			pattern: "y",
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
			pattern: "y",
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
			pattern: "ii",
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
			pattern: "iy",
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
			pattern: "yy",
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
			pattern: "yi",
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
			pattern: "yj",
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
			pattern: "ij",
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
			pattern: "e",
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
			pattern: "ee",
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
			pattern: "e",
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
			pattern: "y",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "oo",
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
			pattern: "'",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "\"",
			phoneticRules: []token{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "ñ",
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
			pattern: "ch",
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
			pattern: "h",
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
			pattern: "h",
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
			pattern: "j",
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ll",
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
			pattern: "w",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
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
			pattern: "b",
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
			pattern: "m",
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
			pattern: "c",
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
			pattern: "c",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
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
			pattern: "gu",
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
			pattern: "g",
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
			pattern: "qu",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "uo",
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
			pattern: "u",
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
			pattern: "y",
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
			pattern: "ü",
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "á",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "í",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ú",
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []token{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []token{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []token{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
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
				pattern: "h",
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
				pattern: "b",
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
				pattern: "b",
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
				pattern: "b",
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
				pattern: "p",
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
				pattern: "p",
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
				pattern: "v",
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
				pattern: "v",
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
				pattern: "v",
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
				pattern: "f",
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
				pattern: "f",
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
				pattern: "g",
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
				pattern: "g",
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
				pattern: "g",
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
				pattern: "k",
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
				pattern: "k",
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
				pattern: "d",
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
				pattern: "d",
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
				pattern: "d",
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
				pattern: "t",
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
				pattern: "t",
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
				pattern: "s",
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
				pattern: "s",
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
				pattern: "z",
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
				pattern: "z",
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
				pattern: "s",
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
				pattern: "Z",
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
				pattern: "S",
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
				pattern: "jnm",
				phoneticRules: []token{
					{
						text:  "jm",
						langs: -1,
					},
				},
			},
			{
				pattern: "ji",
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
				pattern: "jI",
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
				pattern: "a",
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
				pattern: "a",
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
				pattern: "A",
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
				pattern: "B",
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
				pattern: "b",
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
				pattern: "d",
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
				pattern: "f",
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
				pattern: "g",
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
				pattern: "k",
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
				pattern: "l",
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
				pattern: "m",
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
				pattern: "n",
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
				pattern: "p",
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
				pattern: "r",
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
				pattern: "t",
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
				pattern: "v",
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
				pattern: "z",
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
				pattern: "n",
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
				pattern: "kAg",
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
				pattern: "kOg",
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
				pattern: "kog",
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
				pattern: "kag",
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
				pattern: "h",
				phoneticRules: []token{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "H",
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
				pattern: "F",
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
				pattern: "F",
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
				pattern: "B",
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
				pattern: "B",
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
				pattern: "e",
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
				pattern: "i",
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
				pattern: "E",
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
				pattern: "I",
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
				pattern: "F",
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
				pattern: "Q",
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
				pattern: "Y",
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
				pattern: "e",
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
				pattern: "i",
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
				pattern: "E",
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
				pattern: "I",
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
				pattern: "F",
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
				pattern: "Q",
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
				pattern: "Y",
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
				pattern: "lEs",
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
				pattern: "lE",
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
				pattern: "aue",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oue",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AvE",
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
				pattern: "Ave",
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
				pattern: "avE",
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
				pattern: "ave",
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
				pattern: "OvE",
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
				pattern: "Ove",
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
				pattern: "ovE",
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
				pattern: "ove",
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
				pattern: "ea",
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
				pattern: "EA",
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
				pattern: "Ea",
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
				pattern: "eA",
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
				pattern: "aji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "aje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Oji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Oje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "eji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "eje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Eji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Eje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Uji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Uje",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "iji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ije",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Iji",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjI",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ije",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjE",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "aja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ajo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Oja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ojo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "eja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Eja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ejo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Uja",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ujo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ija",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ija",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjA",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ijo",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjO",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "j",
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: "lYndEr",
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
				pattern: "lander",
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
				pattern: "lAndEr",
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
				pattern: "lAnder",
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
				pattern: "landEr",
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
				pattern: "lender",
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
				pattern: "lEndEr",
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
				pattern: "lendEr",
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
				pattern: "lEnder",
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
				pattern: "bUrk",
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
				pattern: "burk",
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
				pattern: "bUrg",
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
				pattern: "burg",
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
				pattern: "s",
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
				pattern: "S",
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
				pattern: "s",
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
				pattern: "S",
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
				pattern: "dS",
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
				pattern: "dZ",
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
				pattern: "Z",
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
				pattern: "S",
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
				pattern: "z",
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
				pattern: "S",
				phoneticRules: []token{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "dZ",
				phoneticRules: []token{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: "Z",
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
					pattern: "b",
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
					pattern: "J",
					phoneticRules: []token{
						{
							text:  "z",
							langs: -1,
						},
					},
				},
				{
					pattern: "aiB",
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
					pattern: "AiB",
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
					pattern: "oiB",
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
					pattern: "OiB",
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
					pattern: "uiB",
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
					pattern: "UiB",
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
					pattern: "eiB",
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
					pattern: "EiB",
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
					pattern: "iiB",
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
					pattern: "IiB",
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
					pattern: "aiB",
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
					pattern: "AiB",
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
					pattern: "oiB",
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
					pattern: "OiB",
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
					pattern: "uiB",
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
					pattern: "UiB",
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
					pattern: "eiB",
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
					pattern: "EiB",
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
					pattern: "iiB",
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
					pattern: "IiB",
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
					pattern: "B",
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
					pattern: "B",
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
					pattern: "B",
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
					pattern: "aiF",
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
					pattern: "AiF",
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
					pattern: "oiF",
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
					pattern: "OiF",
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
					pattern: "uiF",
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
					pattern: "UiF",
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
					pattern: "eiF",
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
					pattern: "EiF",
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
					pattern: "iiF",
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
					pattern: "IiF",
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
					pattern: "aiF",
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
					pattern: "AiF",
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
					pattern: "oiF",
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
					pattern: "OiF",
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
					pattern: "uiF",
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
					pattern: "UiF",
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
					pattern: "eiF",
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
					pattern: "EiF",
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
					pattern: "iiF",
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
					pattern: "IiF",
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
					pattern: "F",
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
					pattern: "F",
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
					pattern: "F",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "lE",
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
					pattern: "lE",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "Ai",
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
					pattern: "oi",
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
					pattern: "Oi",
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
					pattern: "ui",
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
					pattern: "Ui",
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
					pattern: "ei",
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
					pattern: "Ei",
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
					pattern: "iA",
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
					pattern: "iA",
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
					pattern: "A",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "e",
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
					pattern: "e",
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
					pattern: "e",
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
					pattern: "e",
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
					pattern: "e",
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
					pattern: "e",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "a",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "U",
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
					pattern: "U",
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
					pattern: "U",
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
					pattern: "Uk",
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
					pattern: "Uk",
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
					pattern: "sUts",
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
					pattern: "Uts",
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
					pattern: "U",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "om",
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
					pattern: "on",
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
					pattern: "em",
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
					pattern: "en",
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
					pattern: "Em",
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
					pattern: "En",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "om",
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
					pattern: "on",
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
					pattern: "em",
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
					pattern: "en",
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
					pattern: "Em",
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
					pattern: "En",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "lE",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "E",
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
					pattern: "e",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "a",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "AU",
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
					pattern: "aU",
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
					pattern: "Au",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "OU",
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
					pattern: "oU",
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
					pattern: "Ou",
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
					pattern: "ai",
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
					pattern: "Ai",
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
					pattern: "oi",
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
					pattern: "Oi",
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
					pattern: "ui",
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
					pattern: "Ui",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "O",
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
					pattern: "a",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "A",
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
					pattern: "U",
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
					pattern: "U",
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
					pattern: "U",
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
					pattern: "Uk",
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
					pattern: "Uk",
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
					pattern: "sUts",
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
					pattern: "Uts",
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
					pattern: "U",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "aiB",
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
					pattern: "oiB",
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
					pattern: "uiB",
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
					pattern: "eiB",
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
					pattern: "EiB",
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
					pattern: "iiB",
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
					pattern: "IiB",
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
					pattern: "aiB",
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
					pattern: "oiB",
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
					pattern: "uiB",
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
					pattern: "eiB",
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
					pattern: "EiB",
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
					pattern: "iiB",
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
					pattern: "IiB",
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
					pattern: "B",
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
					pattern: "B",
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
					pattern: "B",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "aiF",
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
					pattern: "oiF",
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
					pattern: "uiF",
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
					pattern: "eiF",
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
					pattern: "EiF",
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
					pattern: "iiF",
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
					pattern: "IiF",
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
					pattern: "aiF",
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
					pattern: "oiF",
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
					pattern: "uiF",
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
					pattern: "eiF",
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
					pattern: "EiF",
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
					pattern: "iiF",
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
					pattern: "IiF",
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
					pattern: "F",
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
					pattern: "F",
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
					pattern: "F",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "aiB",
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
					pattern: "oiB",
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
					pattern: "uiB",
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
					pattern: "eiB",
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
					pattern: "EiB",
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
					pattern: "iiB",
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
					pattern: "IiB",
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
					pattern: "aiB",
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
					pattern: "oiB",
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
					pattern: "uiB",
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
					pattern: "eiB",
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
					pattern: "EiB",
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
					pattern: "iiB",
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
					pattern: "IiB",
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
					pattern: "B",
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
					pattern: "B",
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
					pattern: "B",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "aiF",
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
					pattern: "oiF",
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
					pattern: "uiF",
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
					pattern: "eiF",
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
					pattern: "EiF",
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
					pattern: "iiF",
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
					pattern: "IiF",
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
					pattern: "aiF",
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
					pattern: "oiF",
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
					pattern: "uiF",
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
					pattern: "eiF",
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
					pattern: "EiF",
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
					pattern: "iiF",
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
					pattern: "IiF",
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
					pattern: "F",
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
					pattern: "F",
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
					pattern: "F",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "I",
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
					pattern: "Ik",
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
					pattern: "Ik",
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
					pattern: "sIts",
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
					pattern: "Its",
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
					pattern: "I",
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
					pattern: "au",
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
					pattern: "ou",
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
					pattern: "ai",
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
					pattern: "oi",
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
					pattern: "ui",
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
					pattern: "a",
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
					pattern: "e",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
					pattern: "E",
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
				pattern: "h",
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
				pattern: "b",
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
				pattern: "b",
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
				pattern: "b",
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
				pattern: "p",
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
				pattern: "p",
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
				pattern: "v",
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
				pattern: "v",
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
				pattern: "v",
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
				pattern: "f",
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
				pattern: "f",
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
				pattern: "g",
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
				pattern: "g",
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
				pattern: "g",
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
				pattern: "k",
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
				pattern: "k",
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
				pattern: "d",
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
				pattern: "d",
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
				pattern: "d",
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
				pattern: "t",
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
				pattern: "t",
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
				pattern: "s",
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
				pattern: "s",
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
				pattern: "z",
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
				pattern: "z",
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
				pattern: "s",
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
				pattern: "Z",
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
				pattern: "S",
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
				pattern: "jnm",
				phoneticRules: []token{
					{
						text:  "jm",
						langs: -1,
					},
				},
			},
			{
				pattern: "ji",
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
				pattern: "jI",
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
				pattern: "a",
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
				pattern: "a",
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
				pattern: "A",
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
				pattern: "B",
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
				pattern: "b",
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
				pattern: "d",
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
				pattern: "f",
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
				pattern: "g",
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
				pattern: "k",
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
				pattern: "l",
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
				pattern: "m",
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
				pattern: "n",
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
				pattern: "p",
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
				pattern: "r",
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
				pattern: "t",
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
				pattern: "v",
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
				pattern: "z",
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
				pattern: "H",
				phoneticRules: []token{
					{
						text:  "h",
						langs: -1,
					},
				},
			},
			{
				pattern: "s",
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
				pattern: "Z",
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
				pattern: "Z",
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
				pattern: "S",
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
				pattern: "z",
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
				pattern: "ji",
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
				pattern: "jI",
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
				pattern: "je",
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
				pattern: "jE",
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
					pattern: "A",
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "F",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "O",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "U",
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: "J",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "A",
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "F",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []token{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "O",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "U",
					phoneticRules: []token{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: "J",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "B",
					phoneticRules: []token{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "F",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []token{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
					pattern: "E",
					phoneticRules: []token{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
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
