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
			phonetic: "(in[512]|ina)",
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
			phonetic: "(in[512]|ina)",
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
			phonetic: "(lof[512]|lef[512]|lova)",
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
			phonetic: "(lof[512]|lef[512]|lova)",
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
			phonetic: "(of[512]|ova)",
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
			phonetic: "(ef[512]|eva)",
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
			phonetic: "(aja|i[512])",
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
			phonetic: "(aja|i[512])",
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
			phonetic: "(aja|i[512])",
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
			phonetic: "(lova|lof[128]|l[128]|el[128])",
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
			phonetic: "(kova|kof[128]|k[128]|ek[128])",
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
			phonetic: "(ova|of[128]|)",
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
			phonetic: "(lovna|levna|l[128]|el[128])",
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
			phonetic: "(kovna|k[128]|ek[128])",
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
			phonetic: "(ovna|[128])",
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
			phonetic: "(l|el[128])",
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
			phonetic: "(k|ek[128])",
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
			phonetic: "",
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
			phonetic: "(a|i[128])",
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
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ssch",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
			phoneticRules: []phonetic{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsch",
			phonetic: "tS",
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
			phonetic: "(sk[256]|S|StS[512])",
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
			pattern:  "sch",
			phonetic: "(S|StS[512])",
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
			pattern:  "ssh",
			phonetic: "S",
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
			phonetic: "sh",
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
			phonetic: "(S[516]|sh)",
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
			pattern:  "sh",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "kh",
			phonetic: "(x[516]|kh)",
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
			pattern:  "chs",
			phonetic: "(ks[16]|xs|tSs[516])",
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
			phonetic: "(x|k[256]|tS[516])",
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
			pattern:  "ch",
			phonetic: "(x|tS[516])",
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
			pattern:  "ck",
			phonetic: "(k|tsk[128])",
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
			pattern:  "czy",
			phonetic: "tSi",
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
			phonetic: "(tSe|tSF)",
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
			pattern:  "ciewicz",
			phonetic: "(tsevitS|tSevitS)",
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
			pattern:  "siewicz",
			phonetic: "(sevitS|SevitS)",
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
			pattern:  "ziewicz",
			phonetic: "(zevitS|ZevitS)",
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
			pattern:  "riewicz",
			phonetic: "rjevitS",
			phoneticRules: []phonetic{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
			phoneticRules: []phonetic{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
			phoneticRules: []phonetic{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
			phoneticRules: []phonetic{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "icz",
			phonetic: "itS",
			phoneticRules: []phonetic{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cz",
			phonetic: "tS",
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
			phonetic: "(tSB[128]|tsB)",
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
			pattern:  "cia",
			phonetic: "(tSa[128]|tsa)",
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
			phonetic: "(tSom[128]|tsom)",
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
			pattern:  "cią",
			phonetic: "(tSon[128]|tson)",
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
			phonetic: "(tSem[128]|tsem)",
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
			pattern:  "cię",
			phonetic: "(tSen[128]|tsen)",
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
			phonetic: "(tSF[128]|tsF)",
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
			pattern:  "cie",
			phonetic: "(tSe[128]|tse)",
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
			pattern:  "cio",
			phonetic: "(tSo[128]|tso)",
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
			pattern:  "ciu",
			phonetic: "(tSu[128]|tsu)",
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
			phonetic: "(tsi[128]|tSi[384]|tS[256]|si)",
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
			pattern:  "ci",
			phonetic: "(tsi[128]|tSi[384]|si)",
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
			phonetic: "(tsF[128]|tSe[384]|se)",
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
			pattern:  "ce",
			phonetic: "(tSe[384]|tse[128]|se)",
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
			pattern:  "cy",
			phonetic: "(si|tsi[128])",
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
			pattern:  "ssz",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sz",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp[16]|sp)",
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
			pattern:  "sp",
			phonetic: "(Sp[16]|sp)",
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
			pattern:  "sst",
			phonetic: "(St[16]|st)",
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
			pattern:  "st",
			phonetic: "(St[16]|st)",
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
			pattern:  "ss",
			phonetic: "s",
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
			phonetic: "(SB[128]|sB[128]|sja)",
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
			pattern:  "sia",
			phonetic: "(Sa[128]|sja)",
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
			phonetic: "(Som[128]|som)",
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
			pattern:  "sią",
			phonetic: "(Son[128]|son)",
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
			phonetic: "(Sem[128]|sem)",
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
			pattern:  "się",
			phonetic: "(Sen[128]|sen)",
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
			phonetic: "(SF[128]|sF|zi[16])",
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
			pattern:  "sie",
			phonetic: "(se|Se[128]|zi[16])",
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
			pattern:  "sio",
			phonetic: "(So[128]|so)",
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
			pattern:  "siu",
			phonetic: "(Su[128]|sju)",
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
			pattern:  "si",
			phonetic: "(Si[128]|si|zi[16])",
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
			phonetic: "(s|z[16])",
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
			pattern:  "gue",
			phonetic: "ge",
			phoneticRules: []phonetic{
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gui",
			phonetic: "gi",
			phoneticRules: []phonetic{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern:  "guy",
			phonetic: "gi",
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
			phonetic: "(g[256]|gh)",
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
			phonetic: "haus",
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
			phonetic: "haus",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "holts",
			phoneticRules: []phonetic{
				{
					text:  "holts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "goltz",
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "hendler",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hof",
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
			phonetic: "ger",
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
			phonetic: "gen",
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
			phonetic: "gin",
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
			phonetic: "(ge|gi[16]|ji[8])",
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
			pattern:  "gie",
			phonetic: "ge",
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
			phonetic: "(gE|xe[1024]|dZe[260])",
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
			phonetic: "(gI|xi[1024]|dZi[260])",
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
			pattern:  "ge",
			phonetic: "(gE|dZe[260]|hE[512]|xe[1024])",
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
			pattern:  "gi",
			phonetic: "(gI|dZi[260]|hI[512]|xi[1024])",
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
			phonetic: "(gi|dj[64])",
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
			pattern:  "gy",
			phonetic: "(gi|d[64])",
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
			phonetic: "g",
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
			phonetic: "(g|h[512])",
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
			pattern:  "ej",
			phonetic: "(aj|eZ[264]|ex[1024])",
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
			pattern:  "ej",
			phonetic: "aj",
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
			phonetic: "(l|lj)",
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
			phonetic: "(l|lj)",
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
			phonetic: "(l|lj)",
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
			pattern:  "lio",
			phonetic: "(lo|le[512])",
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
			pattern:  "lyo",
			phonetic: "(lo|le[512])",
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
			pattern:  "ll",
			phonetic: "(l|J[1024])",
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
			phonetic: "(j|dZ[4]|x[1024]|Z[264])",
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
			pattern:  "j",
			phonetic: "(j|x[1024])",
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
			pattern:  "pf",
			phonetic: "(pf|p|f)",
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
			pattern:  "ph",
			phonetic: "(ph|f)",
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
			pattern:  "qu",
			phonetic: "(kv[16]|k)",
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
			phonetic: "(Se[128]|re)",
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
			pattern:  "rze",
			phonetic: "(rze|rtsE[16]|Ze[128]|re[128]|rZe[128])",
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
			phonetic: "(Si[128]|ri)",
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
			pattern:  "rzy",
			phonetic: "(Zi[128]|ri[128]|rZi)",
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
			phonetic: "(S[128]|r)",
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
			pattern:  "rz",
			phonetic: "(rz|rts[16]|Z[128]|r[128]|rZ[128])",
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
			phonetic: "(ts|tS[20])",
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
			phonetic: "(ts|tS[20])",
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
			pattern:  "tz",
			phonetic: "(ts[532]|tz)",
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
			pattern:  "zh",
			phonetic: "(Z|zh[128]|tsh[16])",
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
			phonetic: "(ZB[128]|zB[128]|zja)",
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
			pattern:  "zia",
			phonetic: "(Za[128]|zja)",
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
			phonetic: "(Zom[128]|zom)",
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
			pattern:  "zią",
			phonetic: "(Zon[128]|zon)",
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
			phonetic: "(Zem[128]|zem)",
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
			pattern:  "zię",
			phonetic: "(Zen[128]|zen)",
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
			phonetic: "(ZF[128]|zF[128]|ze|tsi[16])",
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
			pattern:  "zie",
			phonetic: "(ze|Ze[128]|tsi[16])",
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
			pattern:  "zio",
			phonetic: "(Zo[128]|zo)",
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
			pattern:  "ziu",
			phonetic: "(Zu[128]|zju)",
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
			pattern:  "zi",
			phonetic: "(Zi[128]|zi|tsi[16])",
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
			phonetic: "tal",
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
			phonetic: "t",
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
			phonetic: "(t[16]|th)",
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
			pattern:  "th",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "vogel",
			phonetic: "(vogel|fogel[16])",
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
			phonetic: "(v|f[16])",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(h|x[384])",
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
			phonetic: "(h|H[20])",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "(e|[8])",
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
			phonetic: "i",
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
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oue",
			phonetic: "oue",
			phoneticRules: []phonetic{
				{
					text:  "oue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "au",
			phonetic: "(au|o[8])",
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
			pattern:  "ou",
			phonetic: "(ou|u[8])",
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
			pattern:  "ue",
			phonetic: "(Q|uje[512])",
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
			pattern:  "ae",
			phonetic: "(Y[16]|aje[512]|ae)",
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
			pattern:  "oe",
			phonetic: "(Y[16]|oje[512]|oe)",
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
			pattern:  "ee",
			phonetic: "(i[4]|aje[512]|e)",
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
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "(aj[16]|oj[16]|eu)",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "(i[16]|e[128]|ije[512]|je)",
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
			pattern:  "ie",
			phonetic: "(i[16]|e[128]|ije[512]|je)",
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
			pattern:  "ye",
			phonetic: "(je|ije[512])",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "io",
			phonetic: "(jo|e[512])",
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
			pattern:  "yo",
			phonetic: "(jo|e[512])",
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
			pattern:  "ea",
			phonetic: "(ea|ja[256])",
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
			phonetic: "(e|je[512])",
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
			pattern:  "oo",
			phonetic: "(u[4]|o)",
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
			pattern:  "uu",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ć",
			phonetic: "(tS[128]|ts)",
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
			pattern:  "ł",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ń",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ñ",
			phonetic: "(n|nj[1024])",
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
			pattern:  "ś",
			phonetic: "(S[128]|s)",
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
			pattern:  "ş",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ż",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ź",
			phonetic: "(Z[128]|z)",
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
			pattern:  "où",
			phonetic: "u",
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
			phonetic: "om",
			phoneticRules: []phonetic{
				{
					text:  "om",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ą",
			phonetic: "on",
			phoneticRules: []phonetic{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ä",
			phonetic: "(Y|e)",
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
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ă",
			phonetic: "(e[256]|a)",
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
			pattern:  "à",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "â",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "è",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ê",
			phonetic: "e",
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
			phonetic: "em",
			phoneticRules: []phonetic{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ę",
			phonetic: "en",
			phoneticRules: []phonetic{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "î",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ö",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ő",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "(u[128]|o)",
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
			pattern:  "ű",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ű",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ß",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "'",
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "\"",
			phonetic: "",
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
			phonetic: "(A|B[128])",
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
			phonetic: "(E|F[128])",
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
			phonetic: "(O|P[128])",
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
			pattern:  "a",
			phonetic: "A",
			phoneticRules: []phonetic{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "(k|ts[128])",
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
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "O",
			phoneticRules: []phonetic{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "U",
			phoneticRules: []phonetic{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "(ts[16]|z)",
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
			pattern:  "ця",
			phonetic: "tsa",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "цю",
			phonetic: "tsu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "циа",
			phonetic: "tsa",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "цие",
			phonetic: "tse",
			phoneticRules: []phonetic{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "цио",
			phonetic: "tso",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "циу",
			phonetic: "tsu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "сие",
			phonetic: "se",
			phoneticRules: []phonetic{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern:  "сио",
			phonetic: "so",
			phoneticRules: []phonetic{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "зие",
			phonetic: "ze",
			phoneticRules: []phonetic{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "зио",
			phonetic: "zo",
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
			phonetic: "haus",
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
			phonetic: "haus",
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
			phonetic: "holts",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hof",
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
			phonetic: "ger",
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
			phonetic: "gen",
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
			phonetic: "gin",
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
			phonetic: "g",
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
			phonetic: "(g|h)",
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
			pattern:  "ля",
			phonetic: "la",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "лю",
			phonetic: "lu",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "лё",
			phonetic: "(le|lo)",
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
			pattern:  "лио",
			phonetic: "(le|lo)",
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
			pattern:  "ле",
			phonetic: "(lE|lo)",
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
			pattern:  "ийе",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ие",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ыйе",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ые",
			phonetic: "je",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "i",
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
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ё",
			phonetic: "(e|jo)",
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
			phonetic: "(jaj|aj)",
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
			phonetic: "je",
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
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "эй",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ей",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ауе",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ауэ",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "а",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "б",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "в",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "г",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "д",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "е",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ж",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "з",
			phonetic: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "и",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "й",
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "к",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "л",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "м",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "н",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "о",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "п",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "р",
			phonetic: "r",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "с",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "т",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "у",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ф",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "х",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ц",
			phonetic: "ts",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ч",
			phonetic: "tS",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ш",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "щ",
			phonetic: "StS",
			phoneticRules: []phonetic{
				{
					text:  "StS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ъ",
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ы",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ь",
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "э",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ю",
			phonetic: "ju",
			phoneticRules: []phonetic{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern:  "я",
			phonetic: "ja",
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
			pattern:  "tch",
			phonetic: "tS",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
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
			pattern:  "ck",
			phonetic: "k",
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
			phonetic: "ks",
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
			phonetic: "",
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
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
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
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gh",
			phonetic: "(g|f|w)",
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
			pattern:  "gn",
			phonetic: "(gn|n)",
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
			phonetic: "(g|dZ)",
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
			pattern:  "th",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ph",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sch",
			phonetic: "(S|sk)",
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
			pattern:  "sh",
			phonetic: "S",
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
			phonetic: "hu",
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
			phonetic: "w",
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
			phonetic: "",
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
			phonetic: "",
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
			phonetic: "H",
			phoneticRules: []phonetic{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "dZ",
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
			phonetic: "n",
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
			phonetic: "m",
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
			phonetic: "(N|ng)",
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
			phonetic: "(pn|n)",
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
			phonetic: "(ps|s)",
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
			pattern:  "qu",
			phonetic: "kw",
			phoneticRules: []phonetic{
				{
					text:  "kw",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tia",
			phonetic: "(So|Sa)",
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
			pattern:  "tio",
			phonetic: "So",
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
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
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
			phonetic: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
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
			phonetic: "i",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oue",
			phonetic: "(aue|oue)",
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
			pattern:  "ai",
			phonetic: "(aj|e)",
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
			pattern:  "ay",
			phonetic: "aj",
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
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "(e|o|a)",
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
			pattern:  "ei",
			phonetic: "(aj|i)",
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
			pattern:  "ey",
			phonetic: "(aj|i)",
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
			pattern:  "ear",
			phonetic: "ia",
			phoneticRules: []phonetic{
				{
					text:  "ia",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ea",
			phonetic: "(i|e)",
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
			pattern:  "ee",
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "(|E)",
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
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "i",
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
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oa",
			phonetic: "ou",
			phoneticRules: []phonetic{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oi",
			phonetic: "oj",
			phoneticRules: []phonetic{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oo",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "(u|ou)",
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
			pattern:  "oy",
			phonetic: "oj",
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
			phonetic: "ou",
			phoneticRules: []phonetic{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "(o|a)",
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
			phonetic: "(ju|u)",
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
			phonetic: "(e|u)",
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
			pattern:  "u",
			phonetic: "(u|a)",
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
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
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
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ph",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "S",
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
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gn",
			phonetic: "(n|gn)",
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
			phonetic: "Z",
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
			phonetic: "k",
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
			phonetic: "g",
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
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
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
			phonetic: "z",
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
			phonetic: "",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
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
			phonetic: "(v|uh)",
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
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
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
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eau",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ai",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ê",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "è",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "à",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "â",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "où",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oi",
			phonetic: "oj",
			phoneticRules: []phonetic{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
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
			phonetic: "j",
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
			phonetic: "(e|)",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yi",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ii",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yy",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
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
			pattern:  "ziu",
			phonetic: "tsu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zia",
			phonetic: "tsa",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zio",
			phonetic: "tso",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ssch",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
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
			phonetic: "evitS",
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
			phonetic: "ovitS",
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
			phonetic: "evitS",
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
			phonetic: "ovitS",
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
			phonetic: "vitS",
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
			phonetic: "vitS",
			phoneticRules: []phonetic{
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sch",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "chs",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ck",
			phonetic: "k",
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
			phonetic: "ts",
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
			phonetic: "Sp",
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
			phonetic: "St",
			phoneticRules: []phonetic{
				{
					text:  "St",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp|sp)",
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
			pattern:  "sp",
			phonetic: "(Sp|sp)",
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
			pattern:  "sst",
			phonetic: "(St|st)",
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
			pattern:  "st",
			phonetic: "(St|st)",
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
			pattern:  "pf",
			phonetic: "(pf|p|f)",
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
			pattern:  "ph",
			phonetic: "(ph|f)",
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
			pattern:  "qu",
			phonetic: "kv",
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
			phonetic: "(evits|evitS)",
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
			phonetic: "(evits|evitS)",
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
			phonetic: "(evits|evitS)",
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
			phonetic: "(evits|evitS)",
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
			phonetic: "(ovits|ovitS)",
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
			phonetic: "(ovits|ovitS)",
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
			phonetic: "(ovits|ovitS)",
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
			phonetic: "(ovits|ovitS)",
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
			phonetic: "(vits|vitS)",
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
			phonetic: "(vits|vitS)",
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
			phonetic: "(vits|vitS)",
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
			phonetic: "(vits|vitS)",
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
			pattern:  "tz",
			phonetic: "ts",
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
			phonetic: "tal",
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
			phonetic: "t",
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
			phonetic: "(t|th)",
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
			pattern:  "th",
			phonetic: "t",
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
			phonetic: "r",
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
			phonetic: "",
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
			phonetic: "H",
			phoneticRules: []phonetic{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ss",
			phonetic: "s",
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
			phonetic: "(z|s)",
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
			phonetic: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ß",
			phonetic: "s",
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
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ue",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ae",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oe",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ä",
			phonetic: "(Y|e)",
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
			pattern:  "ö",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "(aj|oj)",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "I",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ñ",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ã",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ő",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ű",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "A",
			phoneticRules: []phonetic{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "O",
			phoneticRules: []phonetic{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "U",
			phoneticRules: []phonetic{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "(f|v)",
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
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "ts",
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
			pattern:  "אי",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "עי",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "עו",
			phonetic: "VV",
			phoneticRules: []phonetic{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern:  "או",
			phonetic: "VV",
			phoneticRules: []phonetic{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ג׳",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ד׳",
			phonetic: "dZ",
			phoneticRules: []phonetic{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "א",
			phonetic: "L",
			phoneticRules: []phonetic{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ב",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ג",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ד",
			phonetic: "d",
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
			phonetic: "1",
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
			phonetic: "1",
			phoneticRules: []phonetic{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ה",
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "וו",
			phonetic: "V",
			phoneticRules: []phonetic{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern:  "וי",
			phonetic: "WW",
			phoneticRules: []phonetic{
				{
					text:  "WW",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ו",
			phonetic: "W",
			phoneticRules: []phonetic{
				{
					text:  "W",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ז",
			phonetic: "z",
			phoneticRules: []phonetic{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ח",
			phonetic: "X",
			phoneticRules: []phonetic{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ט",
			phonetic: "T",
			phoneticRules: []phonetic{
				{
					text:  "T",
					langs: -1,
				},
			},
		},
		{
			pattern:  "יי",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "י",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ך",
			phonetic: "X",
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
			phonetic: "K",
			phoneticRules: []phonetic{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern:  "כ",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ל",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ם",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "מ",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ן",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "נ",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ס",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ע",
			phonetic: "L",
			phoneticRules: []phonetic{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ף",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "פ",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ץ",
			phonetic: "C",
			phoneticRules: []phonetic{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern:  "צ",
			phonetic: "C",
			phoneticRules: []phonetic{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ק",
			phonetic: "K",
			phoneticRules: []phonetic{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ר",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ש",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ת",
			phonetic: "TB",
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
			pattern:  "sz",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zs",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cs",
			phonetic: "tS",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "(oj|aj)",
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
			pattern:  "ai",
			phonetic: "(oj|aj)",
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
			pattern:  "aj",
			phonetic: "(oj|aj)",
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
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ee",
			phonetic: "(aj|e)",
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
			pattern:  "ely",
			phonetic: "(aj|eli)",
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
			pattern:  "ly",
			phonetic: "(j|li)",
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
			phonetic: "dj",
			phoneticRules: []phonetic{
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gy",
			phonetic: "(d|gi)",
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
			phonetic: "nj",
			phoneticRules: []phonetic{
				{
					text:  "nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ny",
			phonetic: "(n|ni)",
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
			phonetic: "tj",
			phoneticRules: []phonetic{
				{
					text:  "tj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ty",
			phonetic: "(t|ti)",
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
			pattern:  "qu",
			phonetic: "(ku|kv)",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ö",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ő",
			phonetic: "Y",
			phoneticRules: []phonetic{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ű",
			phonetic: "Q",
			phoneticRules: []phonetic{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "ts",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "(S|s)",
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
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
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
			phonetic: "ski",
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
			phonetic: "tski",
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
			phonetic: "(lova|lof|l|el)",
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
			phonetic: "(kova|kof|k|ek)",
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
			phonetic: "(ova|of|)",
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
			phonetic: "(lovna|levna|l|el)",
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
			phonetic: "(kovna|k|ek)",
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
			phonetic: "(ovna|)",
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
			phonetic: "(l|el)",
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
			phonetic: "(k|ek)",
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
			phonetic: "",
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
			phonetic: "(a|i)",
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
			pattern:  "czy",
			phonetic: "tSi",
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
			phonetic: "(tSe|tSF)",
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
			pattern:  "ciewicz",
			phonetic: "(tsevitS|tSevitS)",
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
			pattern:  "siewicz",
			phonetic: "(sevitS|SevitS)",
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
			pattern:  "ziewicz",
			phonetic: "(zevitS|ZevitS)",
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
			pattern:  "riewicz",
			phonetic: "rjevitS",
			phoneticRules: []phonetic{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
			phoneticRules: []phonetic{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
			phoneticRules: []phonetic{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
			phoneticRules: []phonetic{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
			phoneticRules: []phonetic{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "icz",
			phonetic: "itS",
			phoneticRules: []phonetic{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cz",
			phonetic: "tS",
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
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
			phonetic: "(tSB|tsB)",
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
			pattern:  "cia",
			phonetic: "(tSa|tsa)",
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
			phonetic: "(tSom|tsom)",
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
			pattern:  "cią",
			phonetic: "(tSon|tson)",
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
			phonetic: "(tSem|tsem)",
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
			pattern:  "cię",
			phonetic: "(tSen|tsen)",
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
			phonetic: "(tSF|tsF)",
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
			pattern:  "cie",
			phonetic: "(tSe|tse)",
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
			pattern:  "cio",
			phonetic: "(tSo|tso)",
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
			pattern:  "ciu",
			phonetic: "(tSu|tsu)",
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
			pattern:  "ci",
			phonetic: "(tSi|tsI)",
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
			pattern:  "ć",
			phonetic: "(tS|ts)",
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
			pattern:  "ssz",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sz",
			phonetic: "S",
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
			phonetic: "(SB|sB|sja)",
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
			pattern:  "sia",
			phonetic: "(Sa|sja)",
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
			phonetic: "(Som|som)",
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
			pattern:  "sią",
			phonetic: "(Son|son)",
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
			phonetic: "(Sem|sem)",
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
			pattern:  "się",
			phonetic: "(Sen|sen)",
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
			phonetic: "(SF|sF|se)",
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
			pattern:  "sie",
			phonetic: "(Se|se)",
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
			pattern:  "sio",
			phonetic: "(So|so)",
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
			pattern:  "siu",
			phonetic: "(Su|sju)",
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
			pattern:  "si",
			phonetic: "(Si|sI)",
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
			pattern:  "ś",
			phonetic: "(S|s)",
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
			phonetic: "(ZB|zB|zja)",
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
			pattern:  "zia",
			phonetic: "(Za|zja)",
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
			phonetic: "(Zom|zom)",
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
			pattern:  "zią",
			phonetic: "(Zon|zon)",
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
			phonetic: "(Zem|zem)",
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
			pattern:  "zię",
			phonetic: "(Zen|zen)",
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
			phonetic: "(ZF|zF)",
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
			pattern:  "zie",
			phonetic: "(Ze|ze)",
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
			pattern:  "zio",
			phonetic: "(Zo|zo)",
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
			pattern:  "ziu",
			phonetic: "(Zu|zju)",
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
			pattern:  "zi",
			phonetic: "(Zi|zI)",
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
			phonetic: "(Ze|ZF)",
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
			phonetic: "(Ze|ZF|ze|zF)",
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
			pattern:  "że",
			phonetic: "Ze",
			phoneticRules: []phonetic{
				{
					text:  "Ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "źe",
			phonetic: "(Ze|ze)",
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
			pattern:  "ży",
			phonetic: "Zi",
			phoneticRules: []phonetic{
				{
					text:  "Zi",
					langs: -1,
				},
			},
		},
		{
			pattern:  "źi",
			phonetic: "(Zi|zi)",
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
			pattern:  "ż",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ź",
			phonetic: "(Z|z)",
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
			phonetic: "(Se|re)",
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
			pattern:  "rze",
			phonetic: "(Ze|re|rZe)",
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
			phonetic: "(Si|ri)",
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
			pattern:  "rzy",
			phonetic: "(Zi|ri|rZi)",
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
			phonetic: "(S|r)",
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
			pattern:  "rz",
			phonetic: "(Z|r|rZ)",
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
			pattern:  "lio",
			phonetic: "(lo|le)",
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
			pattern:  "ł",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ń",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "k",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "(u|o)",
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
			phonetic: "om",
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
			phonetic: "em",
			phoneticRules: []phonetic{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ą",
			phonetic: "on",
			phoneticRules: []phonetic{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ę",
			phonetic: "en",
			phoneticRules: []phonetic{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ije",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yje",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iie",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yie",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iye",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yye",
			phonetic: "je",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "rie",
			phonetic: "rje",
			phoneticRules: []phonetic{
				{
					text:  "rje",
					langs: -1,
				},
			},
		},
		{
			pattern:  "die",
			phonetic: "dje",
			phoneticRules: []phonetic{
				{
					text:  "dje",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tie",
			phonetic: "tje",
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
			phonetic: "F",
			phoneticRules: []phonetic{
				{
					text:  "F",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "au",
			phonetic: "au",
			phoneticRules: []phonetic{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ej",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ai",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aj",
			phonetic: "aj",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "B",
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
			phonetic: "(E|F)",
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
			phonetic: "P",
			phoneticRules: []phonetic{
				{
					text:  "P",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "ts",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(h|x)",
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
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
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
			pattern:  "j",
			phonetic: "Z",
			phoneticRules: []phonetic{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ce",
			phonetic: "tSe",
			phoneticRules: []phonetic{
				{
					text:  "tSe",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ci",
			phonetic: "(tSi|tS)",
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
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gi",
			phonetic: "(dZi|dZ)",
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
			phonetic: "dZ",
			phoneticRules: []phonetic{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gh",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "aj",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ş",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(x|h)",
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
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "î",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ea",
			phonetic: "ja",
			phoneticRules: []phonetic{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ă",
			phonetic: "(e|a)",
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
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
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
			phonetic: "(in|ina)",
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
			phonetic: "(in|ina)",
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
			phonetic: "(lof|lef)",
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
			phonetic: "(lof|lef|lova)",
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
			phonetic: "(of|ova)",
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
			phonetic: "(ef|ova)",
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
			phonetic: "(aja|i)",
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
			phonetic: "(aja|i)",
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
			phonetic: "(aja|i)",
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
			pattern:  "tsya",
			phonetic: "tsa",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsyu",
			phonetic: "tsu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsia",
			phonetic: "tsa",
			phoneticRules: []phonetic{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsie",
			phonetic: "tse",
			phoneticRules: []phonetic{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsio",
			phonetic: "tso",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsye",
			phonetic: "tse",
			phoneticRules: []phonetic{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsyo",
			phonetic: "tso",
			phoneticRules: []phonetic{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsiu",
			phonetic: "tsu",
			phoneticRules: []phonetic{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sie",
			phonetic: "se",
			phoneticRules: []phonetic{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sio",
			phonetic: "so",
			phoneticRules: []phonetic{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zie",
			phonetic: "ze",
			phoneticRules: []phonetic{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zio",
			phonetic: "zo",
			phoneticRules: []phonetic{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sye",
			phonetic: "se",
			phoneticRules: []phonetic{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern:  "syo",
			phonetic: "so",
			phoneticRules: []phonetic{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zye",
			phonetic: "ze",
			phoneticRules: []phonetic{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zyo",
			phonetic: "zo",
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
			phonetic: "haus",
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
			phonetic: "haus",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "holts",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hajmer",
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
			phonetic: "hajm",
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
			phonetic: "hendler",
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
			phonetic: "hof",
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
			phonetic: "hojf",
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
			phonetic: "hojf",
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
			phonetic: "hojf",
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
			phonetic: "ger",
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
			phonetic: "gen",
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
			phonetic: "gin",
			phoneticRules: []phonetic{
				{
					text:  "gin",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gg",
			phonetic: "g",
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
			phonetic: "(kog|koh)",
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
			phonetic: "(kag|kah)",
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
			phonetic: "g",
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
			phonetic: "(g|h)",
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
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
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
			pattern:  "sch",
			phonetic: "(StS|S)",
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
			pattern:  "ssh",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sh",
			phonetic: "S",
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zh",
			phonetic: "Z",
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
			phonetic: "ts",
			phoneticRules: []phonetic{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tz",
			phonetic: "(ts|tz)",
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
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
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
			pattern:  "q",
			phonetic: "k",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lya",
			phonetic: "la",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lyu",
			phonetic: "lu",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lia",
			phonetic: "la",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "liu",
			phonetic: "lu",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lja",
			phonetic: "la",
			phoneticRules: []phonetic{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lju",
			phonetic: "lu",
			phoneticRules: []phonetic{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "le",
			phonetic: "(lo|lE)",
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
			pattern:  "lyo",
			phonetic: "(lo|le)",
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
			pattern:  "lio",
			phonetic: "(lo|le)",
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
			pattern:  "ije",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iye",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iie",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yje",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ye",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yye",
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yie",
			phonetic: "je",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "io",
			phonetic: "(jo|e)",
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
			phonetic: "j",
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
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ej",
			phonetic: "aj",
			phoneticRules: []phonetic{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e)",
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
			phonetic: "j",
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
			phonetic: "j",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "i",
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
			phonetic: "(je|E)",
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
			pattern:  "ee",
			phonetic: "(aje|i)",
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
			phonetic: "je",
			phoneticRules: []phonetic{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oo",
			phonetic: "(oo|u)",
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
			pattern:  "'",
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "\"",
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phonetic{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phonetic{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phonetic{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
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
			pattern:  "ñ",
			phonetic: "(n|nj)",
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
			pattern:  "ch",
			phonetic: "(tS|dZ)",
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
			phonetic: "",
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
			phonetic: "",
			phoneticRules: []phonetic{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "x",
			phoneticRules: []phonetic{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phonetic{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ll",
			phonetic: "(l|Z)",
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
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "(b|v)",
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
			pattern:  "b",
			phonetic: "(b|v)",
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
			phonetic: "(m|n)",
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
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "(z|s)",
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
			phonetic: "(g|gv)",
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
			phonetic: "(x|g)",
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
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
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
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "(i|j|S|Z)",
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
			pattern:  "ü",
			phonetic: "v",
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phonetic{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phonetic{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phonetic{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phonetic{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phonetic{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phonetic{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phonetic{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phonetic{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phonetic{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phonetic{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phonetic{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phonetic{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phonetic{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phonetic{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phonetic{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
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
				phonetic: "",
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
				phonetic: "p",
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
				phonetic: "",
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
				phonetic: "p",
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
				phonetic: "b",
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
				phonetic: "",
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
				phonetic: "f",
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
				phonetic: "",
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
				phonetic: "f",
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
				phonetic: "v",
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
				phonetic: "",
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
				phonetic: "k",
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
				phonetic: "",
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
				phonetic: "k",
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
				phonetic: "g",
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
				phonetic: "",
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
				phonetic: "t",
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
				phonetic: "",
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
				phonetic: "t",
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
				phonetic: "d",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "s",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
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
				phonetic: "i",
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
				phonetic: "I",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "m",
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
				phonetic: "(kOg|kO[512])",
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
				phonetic: "(kAg|kA[512])",
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
				phonetic: "(kog|ko[512])",
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
				phonetic: "(kag|ka[512])",
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
				pattern:  "h",
				phonetic: "",
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "H",
				phonetic: "(x|)",
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
				phonetic: "e",
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
				phonetic: "e",
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
				phonetic: "a",
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
				phonetic: "a",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "(e|)",
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
				phonetic: "(i|)",
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
				phonetic: "(E|)",
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
				phonetic: "(I|)",
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
				phonetic: "(F|)",
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
				phonetic: "(Q|)",
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
				phonetic: "(Y|)",
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
				pattern:  "lEs",
				phonetic: "(lEs|lz)",
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
				phonetic: "(lE|l)",
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
				pattern:  "aue",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oue",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AvE",
				phonetic: "(D|AvE)",
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
				pattern:  "Ave",
				phonetic: "(D|Ave)",
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
				pattern:  "avE",
				phonetic: "(D|avE)",
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
				pattern:  "ave",
				phonetic: "(D|ave)",
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
				pattern:  "OvE",
				phonetic: "(D|OvE)",
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
				pattern:  "Ove",
				phonetic: "(D|Ove)",
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
				pattern:  "ovE",
				phonetic: "(D|ovE)",
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
				pattern:  "ove",
				phonetic: "(D|ove)",
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
				pattern:  "ea",
				phonetic: "(D|ea)",
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
				pattern:  "EA",
				phonetic: "(D|EA)",
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
				pattern:  "Ea",
				phonetic: "(D|Ea)",
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
				pattern:  "eA",
				phonetic: "(D|eA)",
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
				pattern:  "aji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Oji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Oje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Eji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Eje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Uji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Uje",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "iji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ije",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Iji",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjI",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ije",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjE",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ajo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Oja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ojo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Eja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ejo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Uja",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ujo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ija",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ija",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjA",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ijo",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjO",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "j",
				phonetic: "i",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "lYnder",
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
				phonetic: "(burk|berk)",
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
				phonetic: "(burk|berk)",
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
				phonetic: "(burk|berk)",
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
				phonetic: "(burk|berk)",
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
				phonetic: "z",
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
				phonetic: "z",
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
				phonetic: "z",
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
				phonetic: "z",
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
				phonetic: "S",
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
				phonetic: "S",
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
				phonetic: "S",
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
				phonetic: "(S|s)",
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
				phonetic: "(S|s)",
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
				pattern:  "S",
				phonetic: "s",
				phoneticRules: []phonetic{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern:  "dZ",
				phonetic: "z",
				phoneticRules: []phonetic{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Z",
				phonetic: "z",
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
					pattern:  "b",
					phonetic: "(b|v[1024])",
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
					pattern:  "J",
					phonetic: "z",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(o|om[128]|im[128])",
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
					phonetic: "(a|o|on[128]|in[128])",
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
					pattern:  "B",
					phonetic: "(a|o)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(i|im[128]|om[128])",
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
					phonetic: "(i|in[128]|on[128])",
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
					pattern:  "F",
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
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
					phonetic: "i",
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
					phonetic: "(Q[16]|i|D[4])",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk[16])",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts[16])",
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
					phonetic: "its",
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(Q[16]|i)",
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
					phonetic: "(li|il[4])",
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
					phonetic: "(li|il[4]|lY[16])",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "Ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "Oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "Ui",
					phonetic: "(D|u|i)",
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
					pattern:  "ei",
					phonetic: "(D|i)",
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
					pattern:  "Ei",
					phonetic: "(D|i)",
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
					phonetic: "(ia|io)",
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
					pattern:  "iA",
					phonetic: "(ia|io|iY[16])",
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
					phonetic: "(a|o|Y[16]|D[4])",
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
					phonetic: "(i|Y[16]|[4])",
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
					phonetic: "(i|Y[16]|[4])",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "e",
					phonetic: "(i|Y[16])",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(i|Y[16])",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					phonetic: "o",
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
					phonetic: "o",
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
					phonetic: "o",
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
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "(o|Y[16])",
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
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					pattern:  "A",
					phonetic: "(a|o|Y[16])",
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
					phonetic: "u",
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
					phonetic: "u",
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
					phonetic: "u",
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
					phonetic: "(uk|Qk[16])",
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
					phonetic: "uk",
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
					phonetic: "(suts|sQts[16])",
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
					phonetic: "uts",
					phoneticRules: []phonetic{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "(u|Q[16])",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					phonetic: "(om|im)",
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
					phonetic: "(on|in)",
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
					phonetic: "(im|om)",
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
					phonetic: "(in|on)",
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
					phonetic: "(im|Ym|om)",
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
					phonetic: "(in|Yn|on)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					phonetic: "(om|im)",
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
					phonetic: "(on|in)",
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
					phonetic: "(im|om)",
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
					phonetic: "(in|on)",
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
					phonetic: "(im|Ym|om)",
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
					phonetic: "(in|Yn|on)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "(Q|i|D)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					phonetic: "(il|li|lY)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					phonetic: "(i|)",
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
					phonetic: "(i|)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(Q|i)",
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
					pattern:  "AU",
					phonetic: "(D|a|u)",
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
					pattern:  "aU",
					phonetic: "(D|a|u)",
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
					pattern:  "Au",
					phonetic: "(D|a|u)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "OU",
					phonetic: "(D|o|u)",
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
					pattern:  "oU",
					phonetic: "(D|o|u)",
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
					pattern:  "Ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "Ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "Oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "Ui",
					phonetic: "(D|u|i)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "o",
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
					phonetic: "o",
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
					phonetic: "o",
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
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "(o|Y)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					phonetic: "(a|o)",
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
					pattern:  "A",
					phonetic: "(a|o|Y)",
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
					phonetic: "u",
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
					phonetic: "u",
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
					phonetic: "u",
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
					phonetic: "(uk|Qk)",
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
					phonetic: "uk",
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
					phonetic: "(suts|sQts)",
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
					phonetic: "uts",
					phoneticRules: []phonetic{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "(u|Q)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(o|om|im)",
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
					phonetic: "(o|on|in)",
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
					pattern:  "B",
					phonetic: "o",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(i|im|om)",
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
					phonetic: "(i|in|on)",
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
					pattern:  "F",
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(o|om|im)",
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
					phonetic: "(o|on|in)",
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
					pattern:  "B",
					phonetic: "o",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dm)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(D|Dn)",
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
					phonetic: "(i|im|om)",
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
					phonetic: "(i|in|on)",
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
					pattern:  "F",
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "(ik|Qk)",
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
					phonetic: "ik",
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
					phonetic: "(sits|sQts)",
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
					phonetic: "its",
					phoneticRules: []phonetic{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
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
					pattern:  "au",
					phonetic: "(D|a|u)",
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
					pattern:  "ou",
					phonetic: "(D|o|u)",
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
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
					pattern:  "oi",
					phonetic: "(D|o|i)",
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
					pattern:  "ui",
					phonetic: "(D|u|i)",
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
					pattern:  "a",
					phonetic: "(a|o)",
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
					pattern:  "e",
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
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
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
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
				phonetic: "",
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
				phonetic: "p",
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
				phonetic: "",
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
				phonetic: "p",
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
				phonetic: "b",
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
				phonetic: "",
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
				phonetic: "f",
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
				phonetic: "",
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
				phonetic: "f",
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
				phonetic: "v",
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
				phonetic: "",
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
				phonetic: "k",
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
				phonetic: "",
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
				phonetic: "k",
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
				phonetic: "g",
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
				phonetic: "",
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
				phonetic: "t",
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
				phonetic: "",
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
				phonetic: "t",
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
				phonetic: "d",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "s",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
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
				phonetic: "i",
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
				phonetic: "I",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
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
				phonetic: "",
				phoneticRules: []phonetic{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "H",
				phonetic: "h",
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
				phonetic: "z",
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
				phonetic: "S",
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
				phonetic: "S",
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
				phonetic: "Z",
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
				phonetic: "s",
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
				phonetic: "j",
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
				phonetic: "j",
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
				phonetic: "j",
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
				phonetic: "j",
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
					pattern:  "A",
					phonetic: "a",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "B",
					phonetic: "a",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "F",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "u",
					phoneticRules: []phonetic{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern:  "J",
					phonetic: "l",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "A",
					phonetic: "a",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "B",
					phonetic: "a",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "F",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phonetic{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "u",
					phoneticRules: []phonetic{
						{
							text:  "u",
							langs: -1,
						},
					},
				},
				{
					pattern:  "J",
					phonetic: "l",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "B",
					phonetic: "a",
					phoneticRules: []phonetic{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "F",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "o",
					phoneticRules: []phonetic{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phonetic{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
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
