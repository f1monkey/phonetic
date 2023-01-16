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

var ashRules = map[ashLang][]rule{
	ashany: []rule{
		{
			pattern: "yna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "ssch",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "chsch",
			phoneticRules: []phonetic{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "kh",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "diewicz",
			phoneticRules: []phonetic{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tiewicz",
			phoneticRules: []phonetic{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "iewicz",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ewicz",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "owicz",
			phoneticRules: []phonetic{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "icz",
			phoneticRules: []phonetic{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern: "cz",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sz",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "ssp",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern: "gui",
			phoneticRules: []phonetic{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: "guy",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern: "goltz",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "vogel",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "oue",
			phoneticRules: []phonetic{
				{
					text:  "oue",
					langs: -1,
				},
			},
		},
		{
			pattern: "au",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "eu",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "io",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ć",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "ń",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ñ",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "ţ",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "ż",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ź",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "om",
					langs: -1,
				},
			},
		},
		{
			pattern: "ą",
			phoneticRules: []phonetic{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern: "ä",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "ă",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "â",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "è",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "ê",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern: "ę",
			phoneticRules: []phonetic{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern: "í",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "î",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ö",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ő",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ü",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ú",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ű",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ß",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "'",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "\"",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
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
	ashcyrillic: []rule{
		{
			pattern: "ця",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "цю",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "циа",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "цие",
			phoneticRules: []phonetic{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: "цио",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "циу",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "сие",
			phoneticRules: []phonetic{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: "сио",
			phoneticRules: []phonetic{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: "зие",
			phoneticRules: []phonetic{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "зио",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "лю",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "лё",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ие",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ыйе",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ые",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ё",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "эй",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ей",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ауе",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "ауэ",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "а",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "б",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "в",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "г",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "д",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "е",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "ж",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "з",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: "и",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "й",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "к",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "л",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "м",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "н",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "о",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "п",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "р",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "с",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "т",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "у",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ф",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "х",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ц",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "ч",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ш",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "щ",
			phoneticRules: []phonetic{
				{
					text:  "StS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ъ",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ы",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "ь",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "э",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "ю",
			phoneticRules: []phonetic{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: "я",
			phoneticRules: []phonetic{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
	},
	ashenglish: []rule{
		{
			pattern: "tch",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "gh",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "kh",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ph",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "sch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "kw",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "tia",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "oue",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "ia",
					langs: -1,
				},
			},
		},
		{
			pattern: "ea",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "oa",
			phoneticRules: []phonetic{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern: "oi",
			phoneticRules: []phonetic{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: "oo",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ou",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	ashfrench: []rule{
		{
			pattern: "kh",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ph",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "ç",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "gn",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "uo",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "eau",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ai",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ay",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "ê",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "è",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "à",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "â",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "où",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ou",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "oi",
			phoneticRules: []phonetic{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "yi",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ii",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "yy",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	ashgerman: []rule{
		{
			pattern: "ziu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "zia",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "zio",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "ssch",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "chsch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "sch",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "chs",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ck",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "St",
					langs: -1,
				},
			},
		},
		{
			pattern: "ssp",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern: "ss",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ß",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "ue",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ae",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "oe",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ü",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ä",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "eu",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ñ",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ã",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "ő",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ű",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ç",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
	},
	ashhebrew: []rule{
		{
			pattern: "אי",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "עי",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "עו",
			phoneticRules: []phonetic{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern: "או",
			phoneticRules: []phonetic{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern: "ג׳",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ד׳",
			phoneticRules: []phonetic{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "א",
			phoneticRules: []phonetic{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern: "ב",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "ג",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "ד",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ה",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "וו",
			phoneticRules: []phonetic{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern: "וי",
			phoneticRules: []phonetic{
				{
					text:  "WW",
					langs: -1,
				},
			},
		},
		{
			pattern: "ו",
			phoneticRules: []phonetic{
				{
					text:  "W",
					langs: -1,
				},
			},
		},
		{
			pattern: "ז",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ח",
			phoneticRules: []phonetic{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern: "ט",
			phoneticRules: []phonetic{
				{
					text:  "T",
					langs: -1,
				},
			},
		},
		{
			pattern: "יי",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "י",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ך",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern: "כ",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "ל",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "ם",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "מ",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "ן",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "נ",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ס",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "ע",
			phoneticRules: []phonetic{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern: "ף",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "פ",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "ץ",
			phoneticRules: []phonetic{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern: "צ",
			phoneticRules: []phonetic{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern: "ק",
			phoneticRules: []phonetic{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern: "ר",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "ש",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "ת",
			phoneticRules: []phonetic{
				{
					text:  "TB",
					langs: -1,
				},
			},
		},
	},
	ashhungarian: []rule{
		{
			pattern: "sz",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "zs",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "cs",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ay",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ee",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern: "gy",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "nj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ny",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "tj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ty",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "á",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "í",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ö",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ő",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern: "ú",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "ü",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "ű",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	ashpolish: []rule{
		{
			pattern: "ska",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "diewicz",
			phoneticRules: []phonetic{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tiewicz",
			phoneticRules: []phonetic{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "iewicz",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ewicz",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "owicz",
			phoneticRules: []phonetic{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern: "icz",
			phoneticRules: []phonetic{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern: "cz",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sz",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "Ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "źe",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "Zi",
					langs: -1,
				},
			},
		},
		{
			pattern: "źi",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ź",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "ń",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern: "ą",
			phoneticRules: []phonetic{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern: "ę",
			phoneticRules: []phonetic{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern: "ije",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yje",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iie",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yie",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iye",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yye",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "rie",
			phoneticRules: []phonetic{
				{
					text:  "rje",
					langs: -1,
				},
			},
		},
		{
			pattern: "die",
			phoneticRules: []phonetic{
				{
					text:  "dje",
					langs: -1,
				},
			},
		},
		{
			pattern: "tie",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "F",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "au",
			phoneticRules: []phonetic{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ej",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ai",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ay",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "aj",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "P",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	ashromanian: []rule{
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "ce",
			phoneticRules: []phonetic{
				{
					text:  "tSe",
					langs: -1,
				},
			},
		},
		{
			pattern: "ci",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "gi",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "gh",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ţ",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "ş",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "î",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ea",
			phoneticRules: []phonetic{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern: "ă",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	ashrussian: []rule{
		{
			pattern: "yna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsyu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsia",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsie",
			phoneticRules: []phonetic{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsio",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsye",
			phoneticRules: []phonetic{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsyo",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern: "tsiu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "sie",
			phoneticRules: []phonetic{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: "sio",
			phoneticRules: []phonetic{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: "zie",
			phoneticRules: []phonetic{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "zio",
			phoneticRules: []phonetic{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern: "sye",
			phoneticRules: []phonetic{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern: "syo",
			phoneticRules: []phonetic{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern: "zye",
			phoneticRules: []phonetic{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern: "zyo",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "gin",
					langs: -1,
				},
			},
		},
		{
			pattern: "gg",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "ch",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sh",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "zh",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern: "tz",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "w",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "lya",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "lyu",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "lia",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "liu",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "lja",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern: "lju",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern: "le",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ie",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iye",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "iie",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yje",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ye",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yye",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "yie",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "io",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ei",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ej",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "yo",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "oo",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "\"",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	ashspanish: []rule{
		{
			pattern: "ñ",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ll",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "v",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "z",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "q",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "uo",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "y",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "á",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "é",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "í",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ó",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "ú",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "e",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern: "i",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: "u",
			phoneticRules: []phonetic{
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
		first: []rule{
			{
				pattern: "h",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "jnm",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "H",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oue",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AvE",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "aje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Oji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Oje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "eji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "eje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Eji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Eje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Uji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Uje",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "iji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ije",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Iji",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjI",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ije",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjE",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "aja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ajo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "oja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ojO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Oja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ojo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "OjO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "eja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ejO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Eja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ejo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "EjO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ujO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Uja",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ujo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "UjO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ija",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "ijO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ija",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjA",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Ijo",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "IjO",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "Aju",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "AjU",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "j",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "dZ",
				phoneticRules: []phonetic{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern: "Z",
				phoneticRules: []phonetic{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
		},
		second: map[languageID][]rule{
			languageID(ashany): []rule{
				{
					pattern: "b",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "e",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "O",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern: "U",
					phoneticRules: []phonetic{
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
			languageID(ashrussian): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
			languageID(ashcyrillic): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
			languageID(ashenglish): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
			languageID(ashfrench): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
			languageID(ashgerman): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "O",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern: "U",
					phoneticRules: []phonetic{
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
			languageID(ashhebrew): []rule{},
			languageID(ashhungarian): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
			languageID(ashpolish): []rule{
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
			languageID(ashromanian): []rule{
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
			languageID(ashspanish): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
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
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
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
		first: []rule{
			{
				pattern: "h",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "jnm",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "H",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "j",
						langs: -1,
					},
				},
			},
		},
		second: map[languageID][]rule{
			languageID(ashany): []rule{
				{
					pattern: "A",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "F",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "O",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "U",
					phoneticRules: []phonetic{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: "J",
					phoneticRules: []phonetic{
						{
							text:  "l",
							langs: -1,
						},
					},
				},
			},
			languageID(ashrussian): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashcyrillic): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashenglish): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashfrench): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashgerman): []rule{
				{
					pattern: "A",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "F",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "O",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "U",
					phoneticRules: []phonetic{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern: "J",
					phoneticRules: []phonetic{
						{
							text:  "l",
							langs: -1,
						},
					},
				},
			},
			languageID(ashhebrew): []rule{},
			languageID(ashhungarian): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashpolish): []rule{
				{
					pattern: "B",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern: "F",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "P",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashromanian): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			languageID(ashspanish): []rule{
				{
					pattern: "E",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern: "I",
					phoneticRules: []phonetic{
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
