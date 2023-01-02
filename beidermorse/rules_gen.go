// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type genLang int64

const (
	genany genLang = 1 << iota
	genarabic
	gencyrillic
	genczech
	gendutch
	genenglish
	genfrench
	gengerman
	gengreek
	gengreeklatin
	genhebrew
	genhungarian
	genitalian
	genlatvian
	genpolish
	genportuguese
	genromanian
	genrussian
	genspanish
	genturkish
)

func (l genLang) String() string {
	switch l {
	case genany:
		return "any"
	case genarabic:
		return "arabic"
	case gencyrillic:
		return "cyrillic"
	case genczech:
		return "czech"
	case gendutch:
		return "dutch"
	case genenglish:
		return "english"
	case genfrench:
		return "french"
	case gengerman:
		return "german"
	case gengreek:
		return "greek"
	case gengreeklatin:
		return "greeklatin"
	case genhebrew:
		return "hebrew"
	case genhungarian:
		return "hungarian"
	case genitalian:
		return "italian"
	case genlatvian:
		return "latvian"
	case genpolish:
		return "polish"
	case genportuguese:
		return "portuguese"
	case genromanian:
		return "romanian"
	case genrussian:
		return "russian"
	case genspanish:
		return "spanish"
	case genturkish:
		return "turkish"
	}
	return ""
}

const genAll = genarabic +
	gencyrillic +
	genczech +
	gendutch +
	genenglish +
	genfrench +
	gengerman +
	gengreek +
	gengreeklatin +
	genhebrew +
	genhungarian +
	genitalian +
	genlatvian +
	genpolish +
	genportuguese +
	genromanian +
	genrussian +
	genspanish +
	genturkish

var genRules = map[genLang][]rule{
	genany: []rule{
		{
			pattern: "yna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(in[131072]|ina)",
			phoneticRules: []phoneticRule{
				{
					text:  "in",
					langs: 131072,
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
			phonetic: "(in[131072]|ina)",
			phoneticRules: []phoneticRule{
				{
					text:  "in",
					langs: 131072,
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
			phonetic: "(lova|lof[131072]|lef[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "lova",
					langs: -1,
				},
				{
					text:  "lof",
					langs: 131072,
				},
				{
					text:  "lef",
					langs: 131072,
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
			phonetic: "(lova|lof[131072]|lef[131072]|l[8]|el[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "lova",
					langs: -1,
				},
				{
					text:  "lof",
					langs: 131072,
				},
				{
					text:  "lef",
					langs: 131072,
				},
				{
					text:  "l",
					langs: 8,
				},
				{
					text:  "el",
					langs: 8,
				},
			},
		},
		{
			pattern: "kova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(kova|kof[131072]|k[8]|ek[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "kova",
					langs: -1,
				},
				{
					text:  "kof",
					langs: 131072,
				},
				{
					text:  "k",
					langs: 8,
				},
				{
					text:  "ek",
					langs: 8,
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
			phonetic: "(ova|of[131072]|[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "ova",
					langs: -1,
				},
				{
					text:  "of",
					langs: 131072,
				},
				{
					text:  "",
					langs: 8,
				},
			},
		},
		{
			pattern: "ová",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ova|[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "ova",
					langs: -1,
				},
				{
					text:  "",
					langs: 8,
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
			phonetic: "(eva|ef[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "eva",
					langs: -1,
				},
				{
					text:  "ef",
					langs: 131072,
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
			phonetic: "(aja|i[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: 131072,
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
			phonetic: "(aja|i[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: 131072,
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
			phonetic: "(aja|i[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "aja",
					langs: -1,
				},
				{
					text:  "i",
					langs: 131072,
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
			phonetic: "(lova|lof[16384]|l[16384]|el[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "lova",
					langs: -1,
				},
				{
					text:  "lof",
					langs: 16384,
				},
				{
					text:  "l",
					langs: 16384,
				},
				{
					text:  "el",
					langs: 16384,
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
			phonetic: "(kova|kof[16384]|k[16384]|ek[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "kova",
					langs: -1,
				},
				{
					text:  "kof",
					langs: 16384,
				},
				{
					text:  "k",
					langs: 16384,
				},
				{
					text:  "ek",
					langs: 16384,
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
			phonetic: "(ova|of[16384]|)",
			phoneticRules: []phoneticRule{
				{
					text:  "ova",
					langs: -1,
				},
				{
					text:  "of",
					langs: 16384,
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
			phonetic: "(lovna|levna|l[16384]|el[16384])",
			phoneticRules: []phoneticRule{
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
					langs: 16384,
				},
				{
					text:  "el",
					langs: 16384,
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
			phonetic: "(kovna|k[16384]|ek[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "kovna",
					langs: -1,
				},
				{
					text:  "k",
					langs: 16384,
				},
				{
					text:  "ek",
					langs: 16384,
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
			phonetic: "(ovna|[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "ovna",
					langs: -1,
				},
				{
					text:  "",
					langs: 16384,
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "á",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(a|i[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
				{
					text:  "i",
					langs: 8,
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
			phonetic: "(a|i[16392])",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
				{
					text:  "i",
					langs: 16392,
				},
			},
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
			phoneticRules: []phoneticRule{
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
			pattern: "que",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(k[64]|ke|kve)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: 64,
				},
				{
					text:  "ke",
					langs: -1,
				},
				{
					text:  "kve",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
			phoneticRules: []phoneticRule{
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
			pattern: "m",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bfpv]"),
			},
			phonetic: "(m|n)",
			phoneticRules: []phoneticRule{
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
			pattern: "m",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(m|n[32832])",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
				{
					text:  "n",
					langs: 32832,
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(lo|le[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: 131072,
				},
			},
		},
		{
			pattern:  "lyo",
			phonetic: "(lo|le[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "lo",
					langs: -1,
				},
				{
					text:  "le",
					langs: 131072,
				},
			},
		},
		{
			pattern: "lt",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "u",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lt|[64])",
			phoneticRules: []phoneticRule{
				{
					text:  "lt",
					langs: -1,
				},
				{
					text:  "",
					langs: 64,
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
			phonetic: "(v|f[128]|b[262144])",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "f",
					langs: 128,
				},
				{
					text:  "b",
					langs: 262144,
				},
			},
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "(ez[32768]|eS[32768]|eks|egz)",
			phoneticRules: []phoneticRule{
				{
					text:  "ez",
					langs: 32768,
				},
				{
					text:  "eS",
					langs: 32768,
				},
				{
					text:  "eks",
					langs: -1,
				},
				{
					text:  "egz",
					langs: -1,
				},
			},
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[cs]"),
			},
			phonetic: "(e[32768]|ek)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: 32768,
				},
				{
					text:  "ek",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "u",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ks|[64])",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
				{
					text:  "",
					langs: 64,
				},
			},
		},
		{
			pattern:  "ck",
			phonetic: "(k|tsk[16392])",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "tsk",
					langs: 16392,
				},
			},
		},
		{
			pattern:  "cz",
			phonetic: "(tS|tsz[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
				{
					text:  "tsz",
					langs: 8,
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
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: "dh",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "bh",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ph",
			phonetic: "(ph|f)",
			phoneticRules: []phoneticRule{
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
			pattern:  "kh",
			phonetic: "(x[131104]|kh)",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: 131104,
				},
				{
					text:  "kh",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lh",
			phonetic: "(lh|l[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "lh",
					langs: -1,
				},
				{
					text:  "l",
					langs: 32768,
				},
			},
		},
		{
			pattern:  "nh",
			phonetic: "(nh|nj[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "nh",
					langs: -1,
				},
				{
					text:  "nj",
					langs: 32768,
				},
			},
		},
		{
			pattern:  "ssch",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
			phoneticRules: []phoneticRule{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsch",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "sch",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(S|StS[131072]|sk[69632])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "StS",
					langs: 131072,
				},
				{
					text:  "sk",
					langs: 69632,
				},
			},
		},
		{
			pattern: "sch",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(S|StS[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "StS",
					langs: 131072,
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
			phonetic: "(sk[69632]|S|StS[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "sk",
					langs: 69632,
				},
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "StS",
					langs: 131072,
				},
			},
		},
		{
			pattern:  "sch",
			phonetic: "(S|StS[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "StS",
					langs: 131072,
				},
			},
		},
		{
			pattern:  "ssh",
			phonetic: "S",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(S[131104]|sh)",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: 131104,
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
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zh",
			phonetic: "(Z[131104]|zh|tsh[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: 131104,
				},
				{
					text:  "zh",
					langs: -1,
				},
				{
					text:  "tsh",
					langs: 128,
				},
			},
		},
		{
			pattern:  "chs",
			phonetic: "(ks[128]|xs|tSs[131104])",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: 128,
				},
				{
					text:  "xs",
					langs: -1,
				},
				{
					text:  "tSs",
					langs: 131104,
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
			phonetic: "(x|tS[393248]|k[69632]|S[32832])",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 393248,
				},
				{
					text:  "k",
					langs: 69632,
				},
				{
					text:  "S",
					langs: 32832,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(x|tS[393248]|S[32832])",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 393248,
				},
				{
					text:  "S",
					langs: 32832,
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(t[672]|th)",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: 672,
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
			phoneticRules: []phoneticRule{
				{
					text:  "t",
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
			phonetic: "(g[70144]|gh)",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: 70144,
				},
				{
					text:  "gh",
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
			phonetic: "(v[64]|uh)",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: 64,
				},
				{
					text:  "uh",
					langs: -1,
				},
			},
		},
		{
			pattern: "uh",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phonetic: "(v|uh)",
			phoneticRules: []phoneticRule{
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
			pattern: "h",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(h|x[66048]|H[381024])",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
				{
					text:  "x",
					langs: 66048,
				},
				{
					text:  "H",
					langs: 381024,
				},
			},
		},
		{
			pattern:  "cia",
			phonetic: "(tSa[16384]|tsa)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSa",
					langs: 16384,
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(tSon[16384]|tson)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSon",
					langs: 16384,
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
			phonetic: "(tSem[16384]|tsem)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSem",
					langs: 16384,
				},
				{
					text:  "tsem",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cię",
			phonetic: "(tSen[16384]|tsen)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSen",
					langs: 16384,
				},
				{
					text:  "tsen",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cie",
			phonetic: "(tSe[16384]|tse)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSe",
					langs: 16384,
				},
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cio",
			phonetic: "(tSo[16384]|tso)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSo",
					langs: 16384,
				},
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ciu",
			phonetic: "(tSu[16384]|tsu)",
			phoneticRules: []phoneticRule{
				{
					text:  "tSu",
					langs: 16384,
				},
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern: "sci",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(Si[4096]|stsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
			phoneticRules: []phoneticRule{
				{
					text:  "Si",
					langs: 4096,
				},
				{
					text:  "stsi",
					langs: 16392,
				},
				{
					text:  "dZi",
					langs: 524288,
				},
				{
					text:  "tSi",
					langs: 81920,
				},
				{
					text:  "tS",
					langs: 65536,
				},
				{
					text:  "si",
					langs: -1,
				},
			},
		},
		{
			pattern: "sc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(S[4096]|sts[16392]|dZ[524288]|tS[81920]|s)",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: 4096,
				},
				{
					text:  "sts",
					langs: 16392,
				},
				{
					text:  "dZ",
					langs: 524288,
				},
				{
					text:  "tS",
					langs: 81920,
				},
				{
					text:  "s",
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
			phonetic: "(tsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
			phoneticRules: []phoneticRule{
				{
					text:  "tsi",
					langs: 16392,
				},
				{
					text:  "dZi",
					langs: 524288,
				},
				{
					text:  "tSi",
					langs: 81920,
				},
				{
					text:  "tS",
					langs: 65536,
				},
				{
					text:  "si",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cy",
			phonetic: "(si|tsi[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "si",
					langs: -1,
				},
				{
					text:  "tsi",
					langs: 16384,
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
			phonetic: "(ts[16392]|dZ[524288]|tS[81920]|k[512]|s)",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: 16392,
				},
				{
					text:  "dZ",
					langs: 524288,
				},
				{
					text:  "tS",
					langs: 81920,
				},
				{
					text:  "k",
					langs: 512,
				},
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "sç",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "(s|stS[524288])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "stS",
					langs: 524288,
				},
			},
		},
		{
			pattern:  "ssz",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sz",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(S|s[2048])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "s",
					langs: 2048,
				},
			},
		},
		{
			pattern: "sz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(S|s[2048])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "s",
					langs: 2048,
				},
			},
		},
		{
			pattern:  "sz",
			phonetic: "(S|s[2048]|sts[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "s",
					langs: 2048,
				},
				{
					text:  "sts",
					langs: 128,
				},
			},
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp[128]|sp)",
			phoneticRules: []phoneticRule{
				{
					text:  "Sp",
					langs: 128,
				},
				{
					text:  "sp",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sp",
			phonetic: "(Sp[128]|sp)",
			phoneticRules: []phoneticRule{
				{
					text:  "Sp",
					langs: 128,
				},
				{
					text:  "sp",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sst",
			phonetic: "(St[128]|st)",
			phoneticRules: []phoneticRule{
				{
					text:  "St",
					langs: 128,
				},
				{
					text:  "st",
					langs: -1,
				},
			},
		},
		{
			pattern:  "st",
			phonetic: "(St[128]|st)",
			phoneticRules: []phoneticRule{
				{
					text:  "St",
					langs: 128,
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
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "sj",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "sj",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sj",
			phonetic: "(sj|S[16]|sx[262144]|sZ[589824])",
			phoneticRules: []phoneticRule{
				{
					text:  "sj",
					langs: -1,
				},
				{
					text:  "S",
					langs: 16,
				},
				{
					text:  "sx",
					langs: 262144,
				},
				{
					text:  "sZ",
					langs: 589824,
				},
			},
		},
		{
			pattern:  "sia",
			phonetic: "(Sa[16384]|sa[16384]|sja)",
			phoneticRules: []phoneticRule{
				{
					text:  "Sa",
					langs: 16384,
				},
				{
					text:  "sa",
					langs: 16384,
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
			phonetic: "(Som[16384]|som)",
			phoneticRules: []phoneticRule{
				{
					text:  "Som",
					langs: 16384,
				},
				{
					text:  "som",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sią",
			phonetic: "(Son[16384]|son)",
			phoneticRules: []phoneticRule{
				{
					text:  "Son",
					langs: 16384,
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
			phonetic: "(Sem[16384]|sem)",
			phoneticRules: []phoneticRule{
				{
					text:  "Sem",
					langs: 16384,
				},
				{
					text:  "sem",
					langs: -1,
				},
			},
		},
		{
			pattern:  "się",
			phonetic: "(Sen[16384]|sen)",
			phoneticRules: []phoneticRule{
				{
					text:  "Sen",
					langs: 16384,
				},
				{
					text:  "sen",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sie",
			phonetic: "(se|sje|Se[16384]|zi[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "se",
					langs: -1,
				},
				{
					text:  "sje",
					langs: -1,
				},
				{
					text:  "Se",
					langs: 16384,
				},
				{
					text:  "zi",
					langs: 128,
				},
			},
		},
		{
			pattern:  "sio",
			phonetic: "(So[16384]|so)",
			phoneticRules: []phoneticRule{
				{
					text:  "So",
					langs: 16384,
				},
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "siu",
			phonetic: "(Su[16384]|sju)",
			phoneticRules: []phoneticRule{
				{
					text:  "Su",
					langs: 16384,
				},
				{
					text:  "sju",
					langs: -1,
				},
			},
		},
		{
			pattern: "si",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[äöëaáuiíoóeéêy]$"),
			},
			phonetic: "(Si[16384]|si|zi[37056])",
			phoneticRules: []phoneticRule{
				{
					text:  "Si",
					langs: 16384,
				},
				{
					text:  "si",
					langs: -1,
				},
				{
					text:  "zi",
					langs: 37056,
				},
			},
		},
		{
			pattern:  "si",
			phonetic: "(Si[16384]|si|zi[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "Si",
					langs: 16384,
				},
				{
					text:  "si",
					langs: -1,
				},
				{
					text:  "zi",
					langs: 128,
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
				pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuíoóeéêy]"),
			},
			phonetic: "(s|z[37056])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "z",
					langs: 37056,
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
				pattern:          regexp.MustCompile("^[aeouäöë]"),
			},
			phonetic: "(s|z[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "z",
					langs: 128,
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
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phonetic: "(s|z|Z[32768]|[64])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "Z",
					langs: 32768,
				},
				{
					text:  "",
					langs: 64,
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
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phonetic: "(s|z|Z[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "Z",
					langs: 32768,
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
			phonetic: "(k[64]|gve)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: 64,
				},
				{
					text:  "gve",
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
			phonetic: "(g[64]|gv[294912])",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: 64,
				},
				{
					text:  "gv",
					langs: 294912,
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
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phonetic: "gv",
			phoneticRules: []phoneticRule{
				{
					text:  "gv",
					langs: -1,
				},
			},
		},
		{
			pattern:  "guy",
			phonetic: "gi",
			phoneticRules: []phoneticRule{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gli",
			phonetic: "(glI|l[4096])",
			phoneticRules: []phoneticRule{
				{
					text:  "glI",
					langs: -1,
				},
				{
					text:  "l",
					langs: 4096,
				},
			},
		},
		{
			pattern:  "gni",
			phonetic: "(gnI|ni[4160])",
			phoneticRules: []phoneticRule{
				{
					text:  "gnI",
					langs: -1,
				},
				{
					text:  "ni",
					langs: 4160,
				},
			},
		},
		{
			pattern: "gn",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "(n[4160]|nj[4160]|gn)",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: 4160,
				},
				{
					text:  "nj",
					langs: 4160,
				},
				{
					text:  "gn",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ggie",
			phonetic: "(je[512]|dZe)",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: 512,
				},
				{
					text:  "dZe",
					langs: -1,
				},
			},
		},
		{
			pattern: "ggi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "(j[512]|dZ)",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: 512,
				},
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "(gI|dZ[4096]|j[512])",
			phoneticRules: []phoneticRule{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: 4096,
				},
				{
					text:  "j",
					langs: 512,
				},
			},
		},
		{
			pattern: "gge",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gE|xe[262144]|gZe[32832]|dZe[331808]|je[512])",
			phoneticRules: []phoneticRule{
				{
					text:  "gE",
					langs: -1,
				},
				{
					text:  "xe",
					langs: 262144,
				},
				{
					text:  "gZe",
					langs: 32832,
				},
				{
					text:  "dZe",
					langs: 331808,
				},
				{
					text:  "je",
					langs: 512,
				},
			},
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gI|xi[262144]|gZi[32832]|dZi[331808]|i[512])",
			phoneticRules: []phoneticRule{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "xi",
					langs: 262144,
				},
				{
					text:  "gZi",
					langs: 32832,
				},
				{
					text:  "dZi",
					langs: 331808,
				},
				{
					text:  "i",
					langs: 512,
				},
			},
		},
		{
			pattern: "ggi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "(gI|dZ[4096]|j[512])",
			phoneticRules: []phoneticRule{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: 4096,
				},
				{
					text:  "j",
					langs: 512,
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
			phonetic: "(ge|gi[128]|ji[64]|dZe[4096])",
			phoneticRules: []phoneticRule{
				{
					text:  "ge",
					langs: -1,
				},
				{
					text:  "gi",
					langs: 128,
				},
				{
					text:  "ji",
					langs: 64,
				},
				{
					text:  "dZe",
					langs: 4096,
				},
			},
		},
		{
			pattern:  "gie",
			phonetic: "(ge|gi[128]|dZe[4096]|je[512])",
			phoneticRules: []phoneticRule{
				{
					text:  "ge",
					langs: -1,
				},
				{
					text:  "gi",
					langs: 128,
				},
				{
					text:  "dZe",
					langs: 4096,
				},
				{
					text:  "je",
					langs: 512,
				},
			},
		},
		{
			pattern: "gi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "(i[512]|dZ)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: 512,
				},
				{
					text:  "dZ",
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
			phonetic: "(gE|xe[262144]|Ze[32832]|dZe[331808])",
			phoneticRules: []phoneticRule{
				{
					text:  "gE",
					langs: -1,
				},
				{
					text:  "xe",
					langs: 262144,
				},
				{
					text:  "Ze",
					langs: 32832,
				},
				{
					text:  "dZe",
					langs: 331808,
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
			phonetic: "(gI|xi[262144]|Zi[32832]|dZi[331808])",
			phoneticRules: []phoneticRule{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "xi",
					langs: 262144,
				},
				{
					text:  "Zi",
					langs: 32832,
				},
				{
					text:  "dZi",
					langs: 331808,
				},
			},
		},
		{
			pattern:  "ge",
			phonetic: "(gE|xe[262144]|hE[131072]|je[512]|Ze[32832]|dZe[331808])",
			phoneticRules: []phoneticRule{
				{
					text:  "gE",
					langs: -1,
				},
				{
					text:  "xe",
					langs: 262144,
				},
				{
					text:  "hE",
					langs: 131072,
				},
				{
					text:  "je",
					langs: 512,
				},
				{
					text:  "Ze",
					langs: 32832,
				},
				{
					text:  "dZe",
					langs: 331808,
				},
			},
		},
		{
			pattern:  "gi",
			phonetic: "(gI|xi[262144]|hI[131072]|i[512]|Zi[32832]|dZi[331808])",
			phoneticRules: []phoneticRule{
				{
					text:  "gI",
					langs: -1,
				},
				{
					text:  "xi",
					langs: 262144,
				},
				{
					text:  "hI",
					langs: 131072,
				},
				{
					text:  "i",
					langs: 512,
				},
				{
					text:  "Zi",
					langs: 32832,
				},
				{
					text:  "dZi",
					langs: 331808,
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
			phonetic: "(gi|dj[2048])",
			phoneticRules: []phoneticRule{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "dj",
					langs: 2048,
				},
			},
		},
		{
			pattern:  "gy",
			phonetic: "(gi|d[2048])",
			phoneticRules: []phoneticRule{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "d",
					langs: 2048,
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
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouyei]"),
			},
			phonetic: "g",
			phoneticRules: []phoneticRule{
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
			phonetic: "(g|h[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "h",
					langs: 131072,
				},
			},
		},
		{
			pattern:  "ij",
			phonetic: "(i|ej[16]|ix[262144]|iZ[622656])",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "ej",
					langs: 16,
				},
				{
					text:  "ix",
					langs: 262144,
				},
				{
					text:  "iZ",
					langs: 622656,
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
			phonetic: "(j|dZ[32]|x[262144]|Z[622656])",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: 32,
				},
				{
					text:  "x",
					langs: 262144,
				},
				{
					text:  "Z",
					langs: 622656,
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
			phonetic: "(S[16384]|r)",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: 16384,
				},
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "rz",
			phonetic: "(rz|rts[128]|Z[16384]|r[16384]|rZ[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "rz",
					langs: -1,
				},
				{
					text:  "rts",
					langs: 128,
				},
				{
					text:  "Z",
					langs: 16384,
				},
				{
					text:  "r",
					langs: 16384,
				},
				{
					text:  "rZ",
					langs: 16384,
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
			phonetic: "(ts|tS[160])",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 160,
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
			phonetic: "(ts[131232]|tS[160])",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: 131232,
				},
				{
					text:  "tS",
					langs: 160,
				},
			},
		},
		{
			pattern:  "tz",
			phonetic: "(ts[131232]|tz)",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: 131232,
				},
				{
					text:  "tz",
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
			phonetic: "(Za[16384]|za[16384]|zja)",
			phoneticRules: []phoneticRule{
				{
					text:  "Za",
					langs: 16384,
				},
				{
					text:  "za",
					langs: 16384,
				},
				{
					text:  "zja",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zia",
			phonetic: "(Za[16384]|zja)",
			phoneticRules: []phoneticRule{
				{
					text:  "Za",
					langs: 16384,
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
			phonetic: "(Zom[16384]|zom)",
			phoneticRules: []phoneticRule{
				{
					text:  "Zom",
					langs: 16384,
				},
				{
					text:  "zom",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zią",
			phonetic: "(Zon[16384]|zon)",
			phoneticRules: []phoneticRule{
				{
					text:  "Zon",
					langs: 16384,
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
			phonetic: "(Zem[16384]|zem)",
			phoneticRules: []phoneticRule{
				{
					text:  "Zem",
					langs: 16384,
				},
				{
					text:  "zem",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zię",
			phonetic: "(Zen[16384]|zen)",
			phoneticRules: []phoneticRule{
				{
					text:  "Zen",
					langs: 16384,
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
			phonetic: "(Ze[16384]|ze[16384]|ze|tsi[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "Ze",
					langs: 16384,
				},
				{
					text:  "ze",
					langs: 16384,
				},
				{
					text:  "ze",
					langs: -1,
				},
				{
					text:  "tsi",
					langs: 128,
				},
			},
		},
		{
			pattern:  "zie",
			phonetic: "(ze|Ze[16384]|tsi[128])",
			phoneticRules: []phoneticRule{
				{
					text:  "ze",
					langs: -1,
				},
				{
					text:  "Ze",
					langs: 16384,
				},
				{
					text:  "tsi",
					langs: 128,
				},
			},
		},
		{
			pattern:  "zio",
			phonetic: "(Zo[16384]|zo)",
			phoneticRules: []phoneticRule{
				{
					text:  "Zo",
					langs: 16384,
				},
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ziu",
			phonetic: "(Zu[16384]|zju)",
			phoneticRules: []phoneticRule{
				{
					text:  "Zu",
					langs: 16384,
				},
				{
					text:  "zju",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zi",
			phonetic: "(Zi[16384]|zi|tsi[128]|dzi[4096]|tsi[4096]|si[262144])",
			phoneticRules: []phoneticRule{
				{
					text:  "Zi",
					langs: 16384,
				},
				{
					text:  "zi",
					langs: -1,
				},
				{
					text:  "tsi",
					langs: 128,
				},
				{
					text:  "dzi",
					langs: 4096,
				},
				{
					text:  "tsi",
					langs: 4096,
				},
				{
					text:  "si",
					langs: 262144,
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
			phonetic: "(s|ts[128]|ts[4096]|S[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 128,
				},
				{
					text:  "ts",
					langs: 4096,
				},
				{
					text:  "S",
					langs: 32768,
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
				pattern:          regexp.MustCompile("^[bdgv]"),
			},
			phonetic: "(z|dz[4096]|Z[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "dz",
					langs: 4096,
				},
				{
					text:  "Z",
					langs: 32768,
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
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phonetic: "(s|ts[4096]|S[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 4096,
				},
				{
					text:  "S",
					langs: 32768,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oue",
			phonetic: "(oue|ve[64])",
			phoneticRules: []phoneticRule{
				{
					text:  "oue",
					langs: -1,
				},
				{
					text:  "ve",
					langs: 64,
				},
			},
		},
		{
			pattern:  "eau",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ae",
			phonetic: "(Y[128]|aje[131072]|ae)",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: 128,
				},
				{
					text:  "aje",
					langs: 131072,
				},
				{
					text:  "ae",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ai",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "au",
			phonetic: "(au|o[64])",
			phoneticRules: []phoneticRule{
				{
					text:  "au",
					langs: -1,
				},
				{
					text:  "o",
					langs: 64,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ão",
			phonetic: "(au|an)",
			phoneticRules: []phoneticRule{
				{
					text:  "au",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ãe",
			phonetic: "(aj|an)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ãi",
			phonetic: "(aj|an)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ea",
			phonetic: "(ea|ja[65536])",
			phoneticRules: []phoneticRule{
				{
					text:  "ea",
					langs: -1,
				},
				{
					text:  "ja",
					langs: 65536,
				},
			},
		},
		{
			pattern:  "ee",
			phonetic: "(i[32]|aje[131072]|e)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: 32,
				},
				{
					text:  "aje",
					langs: 131072,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "(eu|Yj[128]|ej[128]|oj[128]|Y[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "eu",
					langs: -1,
				},
				{
					text:  "Yj",
					langs: 128,
				},
				{
					text:  "ej",
					langs: 128,
				},
				{
					text:  "oj",
					langs: 128,
				},
				{
					text:  "Y",
					langs: 16,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ia",
			phonetic: "ja",
			phoneticRules: []phoneticRule{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "(i[128]|e[16384]|ije[131072]|Q[16]|je)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: 128,
				},
				{
					text:  "e",
					langs: 16384,
				},
				{
					text:  "ije",
					langs: 131072,
				},
				{
					text:  "Q",
					langs: 16,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "io",
			phonetic: "(jo|e[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "jo",
					langs: -1,
				},
				{
					text:  "e",
					langs: 131072,
				},
			},
		},
		{
			pattern:  "iu",
			phonetic: "ju",
			phoneticRules: []phoneticRule{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oe",
			phonetic: "(Y[128]|oje[131072]|u[16]|oe)",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: 128,
				},
				{
					text:  "oje",
					langs: 131072,
				},
				{
					text:  "u",
					langs: 16,
				},
				{
					text:  "oe",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oi",
			phonetic: "oj",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oo",
			phonetic: "(u[32]|o)",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: 32,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "(ou|u[576]|au[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "ou",
					langs: -1,
				},
				{
					text:  "u",
					langs: 576,
				},
				{
					text:  "au",
					langs: 16,
				},
			},
		},
		{
			pattern:  "où",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oy",
			phonetic: "oj",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "õe",
			phonetic: "(oj|on)",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ua",
			phonetic: "va",
			phoneticRules: []phoneticRule{
				{
					text:  "va",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ue",
			phonetic: "(Q[128]|uje[131072]|ve)",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: 128,
				},
				{
					text:  "uje",
					langs: 131072,
				},
				{
					text:  "ve",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ui",
			phonetic: "(uj|vi|Y[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "uj",
					langs: -1,
				},
				{
					text:  "vi",
					langs: -1,
				},
				{
					text:  "Y",
					langs: 16,
				},
			},
		},
		{
			pattern:  "uu",
			phonetic: "(u|Q[16])",
			phoneticRules: []phoneticRule{
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
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
			phoneticRules: []phoneticRule{
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
			pattern:  "uy",
			phonetic: "uj",
			phoneticRules: []phoneticRule{
				{
					text:  "uj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ya",
			phonetic: "ja",
			phoneticRules: []phoneticRule{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ye",
			phonetic: "(je|ije[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
				{
					text:  "ije",
					langs: 131072,
				},
			},
		},
		{
			pattern: "yi",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "jo",
					langs: -1,
				},
				{
					text:  "e",
					langs: 131072,
				},
			},
		},
		{
			pattern:  "yu",
			phonetic: "ju",
			phoneticRules: []phoneticRule{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
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
				pattern:          regexp.MustCompile("[áóéê]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[áóéê]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
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
			phonetic: "(e|je[131072])",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "je",
					langs: 131072,
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
			phonetic: "(e|EE[96])",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "EE",
					langs: 96,
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
			phoneticRules: []phoneticRule{
				{
					text:  "om",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ą",
			phonetic: "on",
			phoneticRules: []phoneticRule{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ä",
			phonetic: "(Y|e)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "à",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "â",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ã",
			phonetic: "(a|an)",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ă",
			phonetic: "(e[65536]|a)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: 65536,
				},
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ā",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "č",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ć",
			phonetic: "(tS[16384]|ts)",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: 16384,
				},
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "(s|tS[524288])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 524288,
				},
			},
		},
		{
			pattern:  "ď",
			phonetic: "(d|dj[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
				{
					text:  "dj",
					langs: 8,
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
			phoneticRules: []phoneticRule{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ę",
			phonetic: "en",
			phoneticRules: []phoneticRule{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "è",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ê",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ě",
			phonetic: "(e|je[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "je",
					langs: 8,
				},
			},
		},
		{
			pattern:  "ē",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ģ",
			phonetic: "(d|dj)",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ğ",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "î",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ī",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ı",
			phonetic: "(i|e[524288]|[524288])",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "e",
					langs: 524288,
				},
				{
					text:  "",
					langs: 524288,
				},
			},
		},
		{
			pattern:  "ķ",
			phonetic: "(k|t[8192]|tj[8192])",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "t",
					langs: 8192,
				},
				{
					text:  "tj",
					langs: 8192,
				},
			},
		},
		{
			pattern:  "ļ",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ł",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ń",
			phonetic: "(n|nj[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "nj",
					langs: 16384,
				},
			},
		},
		{
			pattern:  "ñ",
			phonetic: "(n|nj[262144])",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "nj",
					langs: 262144,
				},
			},
		},
		{
			pattern:  "ņ",
			phonetic: "(n|nj[8192])",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "nj",
					langs: 8192,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "(u[16384]|o)",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: 16384,
				},
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ô",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "õ",
			phonetic: "(o|on[32768]|Y[2048])",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
				{
					text:  "on",
					langs: 32768,
				},
				{
					text:  "Y",
					langs: 2048,
				},
			},
		},
		{
			pattern:  "ò",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ö",
			phonetic: "Y",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ř",
			phonetic: "(r|rZ[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
				{
					text:  "rZ",
					langs: 8,
				},
			},
		},
		{
			pattern:  "ś",
			phonetic: "(S[16384]|s)",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: 16384,
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
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "š",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ť",
			phonetic: "(t|tj[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "tj",
					langs: 8,
				},
			},
		},
		{
			pattern:  "ű",
			phonetic: "Q",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "(Q|u[294912])",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "u",
					langs: 294912,
				},
			},
		},
		{
			pattern:  "ū",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ů",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ù",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ý",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ż",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ź",
			phonetic: "(Z[16384]|z)",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: 16384,
				},
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ž",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ß",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "'",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "\"",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
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
			phonetic: "(O|P[16384])",
			phoneticRules: []phoneticRule{
				{
					text:  "O",
					langs: -1,
				},
				{
					text:  "P",
					langs: 16384,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "A",
			phoneticRules: []phoneticRule{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "B",
			phoneticRules: []phoneticRule{
				{
					text:  "B",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "(k|ts[16392]|dZ[524288])",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 16392,
				},
				{
					text:  "dZ",
					langs: 524288,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(h|x[65536]|H[299072])",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
				{
					text:  "x",
					langs: 65536,
				},
				{
					text:  "H",
					langs: 299072,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "(j|x[262144]|Z[622656])",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "x",
					langs: 262144,
				},
				{
					text:  "Z",
					langs: 622656,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "O",
			phoneticRules: []phoneticRule{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "(s|S[32768])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "S",
					langs: 32768,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "U",
			phoneticRules: []phoneticRule{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "V",
			phoneticRules: []phoneticRule{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "(v|w[48])",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "w",
					langs: 48,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "(ks|gz|S[294912])",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
				{
					text:  "gz",
					langs: -1,
				},
				{
					text:  "S",
					langs: 294912,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "(z|ts[128]|dz[4096]|ts[4096]|s[262144])",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 128,
				},
				{
					text:  "dz",
					langs: 4096,
				},
				{
					text:  "ts",
					langs: 4096,
				},
				{
					text:  "s",
					langs: 262144,
				},
			},
		},
	},
	genarabic: []rule{
		{
			pattern:  "ا",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: "ب",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ب",
			phonetic: "b1",
			phoneticRules: []phoneticRule{
				{
					text:  "b1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ت",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ت",
			phonetic: "t1",
			phoneticRules: []phoneticRule{
				{
					text:  "t1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ث",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ث",
			phonetic: "t1",
			phoneticRules: []phoneticRule{
				{
					text:  "t1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ج",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(dZ|Z)",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ج",
			phonetic: "(dZ1|Z1)",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ1",
					langs: -1,
				},
				{
					text:  "Z1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ح",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ح",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ح",
			phonetic: "(h1|1)",
			phoneticRules: []phoneticRule{
				{
					text:  "h1",
					langs: -1,
				},
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "خ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "خ",
			phonetic: "x1",
			phoneticRules: []phoneticRule{
				{
					text:  "x1",
					langs: -1,
				},
			},
		},
		{
			pattern: "د",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "د",
			phonetic: "d1",
			phoneticRules: []phoneticRule{
				{
					text:  "d1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ذ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ذ",
			phonetic: "d1",
			phoneticRules: []phoneticRule{
				{
					text:  "d1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ر",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ر",
			phonetic: "r1",
			phoneticRules: []phoneticRule{
				{
					text:  "r1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ز",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ز",
			phonetic: "z1",
			phoneticRules: []phoneticRule{
				{
					text:  "z1",
					langs: -1,
				},
			},
		},
		{
			pattern: "س",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "س",
			phonetic: "s1",
			phoneticRules: []phoneticRule{
				{
					text:  "s1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ش",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ش",
			phonetic: "S1",
			phoneticRules: []phoneticRule{
				{
					text:  "S1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ص",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ص",
			phonetic: "s1",
			phoneticRules: []phoneticRule{
				{
					text:  "s1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ض",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ض",
			phonetic: "d1",
			phoneticRules: []phoneticRule{
				{
					text:  "d1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ط",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ط",
			phonetic: "t1",
			phoneticRules: []phoneticRule{
				{
					text:  "t1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ظ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ظ",
			phonetic: "z1",
			phoneticRules: []phoneticRule{
				{
					text:  "z1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ع",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ع",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ع",
			phonetic: "(h1|1)",
			phoneticRules: []phoneticRule{
				{
					text:  "h1",
					langs: -1,
				},
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "غ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "غ",
			phonetic: "g1",
			phoneticRules: []phoneticRule{
				{
					text:  "g1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ف",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ف",
			phonetic: "f1",
			phoneticRules: []phoneticRule{
				{
					text:  "f1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ق",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ق",
			phonetic: "k1",
			phoneticRules: []phoneticRule{
				{
					text:  "k1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ك",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ك",
			phonetic: "k1",
			phoneticRules: []phoneticRule{
				{
					text:  "k1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ل",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ل",
			phonetic: "l1",
			phoneticRules: []phoneticRule{
				{
					text:  "l1",
					langs: -1,
				},
			},
		},
		{
			pattern: "م",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "م",
			phonetic: "m1",
			phoneticRules: []phoneticRule{
				{
					text:  "m1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ن",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ن",
			phonetic: "n1",
			phoneticRules: []phoneticRule{
				{
					text:  "n1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ه",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ه",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ه",
			phonetic: "(h1|1)",
			phoneticRules: []phoneticRule{
				{
					text:  "h1",
					langs: -1,
				},
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern: "و",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(u|v)",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "و",
			phonetic: "(u|v1)",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
				{
					text:  "v1",
					langs: -1,
				},
			},
		},
		{
			pattern: "ي\u200e",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(i|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ي\u200e",
			phonetic: "(i|j1)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "j1",
					langs: -1,
				},
			},
		},
	},
	gencyrillic: []rule{
		{
			pattern:  "ця",
			phonetic: "tsa",
			phoneticRules: []phoneticRule{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "цю",
			phonetic: "tsu",
			phoneticRules: []phoneticRule{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "циа",
			phonetic: "tsa",
			phoneticRules: []phoneticRule{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "цие",
			phonetic: "tse",
			phoneticRules: []phoneticRule{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "цио",
			phonetic: "tso",
			phoneticRules: []phoneticRule{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "циу",
			phonetic: "tsu",
			phoneticRules: []phoneticRule{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "сие",
			phonetic: "se",
			phoneticRules: []phoneticRule{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern:  "сио",
			phonetic: "so",
			phoneticRules: []phoneticRule{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "зие",
			phonetic: "ze",
			phoneticRules: []phoneticRule{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "зио",
			phonetic: "zo",
			phoneticRules: []phoneticRule{
				{
					text:  "zo",
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(hejmer|hajmer)",
			phoneticRules: []phoneticRule{
				{
					text:  "hejmer",
					langs: -1,
				},
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
			phonetic: "(hejm|hajm)",
			phoneticRules: []phoneticRule{
				{
					text:  "hejm",
					langs: -1,
				},
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "лю",
			phonetic: "lu",
			phoneticRules: []phoneticRule{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "лё",
			phonetic: "(le|lo)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ие",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ыйе",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ые",
			phonetic: "je",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "i",
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
			phonetic: "(jej|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "jej",
					langs: -1,
				},
				{
					text:  "ej",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "эй",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ей",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ауе",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ауэ",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "а",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "б",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "в",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "г",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "д",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "е",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ё",
			phonetic: "(e|jo)",
			phoneticRules: []phoneticRule{
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
			pattern:  "ж",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "з",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "и",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "й",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "к",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "л",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "м",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "н",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "о",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "п",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "р",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "с",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "т",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "у",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ф",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "х",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ц",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ч",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ш",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "щ",
			phonetic: "StS",
			phoneticRules: []phoneticRule{
				{
					text:  "StS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ъ",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ы",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ь",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "э",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ю",
			phonetic: "ju",
			phoneticRules: []phoneticRule{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern:  "я",
			phonetic: "ja",
			phoneticRules: []phoneticRule{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
	},
	genczech: []rule{
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "(k|kv)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "kv",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "(ej|aj)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "č",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "š",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ž",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ň",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ť",
			phonetic: "(t|tj)",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "tj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ď",
			phonetic: "(d|dj)",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ř",
			phonetic: "(r|rZ)",
			phoneticRules: []phoneticRule{
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
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ý",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ě",
			phonetic: "(e|je)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ů",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(h|g)",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "(k|kv)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "kv",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	gendutch: []rule{
		{
			pattern:  "ssj",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sj",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
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
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ck",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "kv",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ss",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
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
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "au",
			phoneticRules: []phoneticRule{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "(Q|i)",
			phoneticRules: []phoneticRule{
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
			pattern:  "uu",
			phonetic: "(Q|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ee",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "(Y|Yj)",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
				{
					text:  "Yj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aa",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oo",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oe",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ij",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ui",
			phonetic: "(Y|uj)",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
				{
					text:  "uj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "(ej|aj)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
				{
					text:  "aj",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "(g|x)",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "(i|Q)",
			phoneticRules: []phoneticRule{
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
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "(u|Q)",
			phoneticRules: []phoneticRule{
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
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
			phoneticRules: []phoneticRule{
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
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genenglish: []rule{
		{
			pattern:  "�",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "'",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "mc",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "mak",
			phoneticRules: []phoneticRule{
				{
					text:  "mak",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tz",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tch",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "s",
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
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gh",
			phonetic: "(g|f|w)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ph",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sch",
			phonetic: "(S|sk)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "H",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "kw",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tia",
			phonetic: "(So|Sa)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "r",
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
			phoneticRules: []phoneticRule{
				{
					text:  "z",
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
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "yi",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oue",
			phonetic: "(aue|oue)",
			phoneticRules: []phoneticRule{
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
			phonetic: "(aj|ej|e)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
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
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
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
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "(ej|aj|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
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
			phonetic: "(ej|aj|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
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
			phoneticRules: []phoneticRule{
				{
					text:  "ia",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ea",
			phonetic: "(i|e)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			pattern:  "ie",
			phonetic: "i",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oa",
			phonetic: "ou",
			phoneticRules: []phoneticRule{
				{
					text:  "ou",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oi",
			phonetic: "oj",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oo",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "(u|ou)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "ou",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			pattern:  "a",
			phonetic: "(e|o|a)",
			phoneticRules: []phoneticRule{
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
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "(o|a)",
			phoneticRules: []phoneticRule{
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
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "(u|a)",
			phoneticRules: []phoneticRule{
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
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
			phoneticRules: []phoneticRule{
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
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genfrench: []rule{
		{
			pattern: "lt",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "u",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lt|)",
			phoneticRules: []phoneticRule{
				{
					text:  "lt",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "c",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "n",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(k|)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
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
			phonetic: "(t|)",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "",
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
				suffix:           "n",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(k|)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "p",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(p|)",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "r",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "e",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(r|)",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(t|)",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "",
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
			phonetic: "(s|)",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ds",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ds|)",
			phoneticRules: []phoneticRule{
				{
					text:  "ds",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ps",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ps|)",
			phoneticRules: []phoneticRule{
				{
					text:  "ps",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "rs",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "e",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(rs|)",
			phoneticRules: []phoneticRule{
				{
					text:  "rs",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ts",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ts|)",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(s|)",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "x",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "u",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ks|)",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
				{
					text:  "",
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
				pattern:          regexp.MustCompile("[aeéèêiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			phonetic: "(s|)",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "t",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeéèêiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			phonetic: "(t|)",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ph",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "S",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gn",
			phonetic: "(n|gn)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "aill",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "e",
				suffix:           "",
			},
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: "ll",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "e",
				suffix:           "",
			},
			phonetic: "(l|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "j",
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
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(m|n)",
			phoneticRules: []phoneticRule{
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
			pattern: "ou",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
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
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eau",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "au",
			phonetic: "(o|au)",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ai",
			phonetic: "(e|aj)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "(e|aj)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ê",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "è",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "à",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "â",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "où",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oi",
			phonetic: "(oj|va)",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
				{
					text:  "va",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "(aj|ej|e)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej|e)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "(ej|Y)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
				{
					text:  "Y",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "(u|Q)",
			phoneticRules: []phoneticRule{
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
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	gengerman: []rule{
		{
			pattern: "ewitsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "evitS",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "vitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ssch",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
			phoneticRules: []phoneticRule{
				{
					text:  "xS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sch",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ziu",
			phonetic: "tsu",
			phoneticRules: []phoneticRule{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zia",
			phonetic: "tsa",
			phoneticRules: []phoneticRule{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zio",
			phonetic: "tso",
			phoneticRules: []phoneticRule{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "chs",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ck",
			phonetic: "k",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "St",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp|sp)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "H",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ss",
			phonetic: "s",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ß",
			phonetic: "s",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ue",
			phonetic: "Q",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ae",
			phonetic: "Y",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oe",
			phonetic: "Y",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "Q",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ä",
			phonetic: "(Y|e)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "(Yj|ej|aj|oj)",
			phoneticRules: []phoneticRule{
				{
					text:  "Yj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "I",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ñ",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ã",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ő",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ű",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "A",
			phoneticRules: []phoneticRule{
				{
					text:  "A",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "O",
			phoneticRules: []phoneticRule{
				{
					text:  "O",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "U",
			phoneticRules: []phoneticRule{
				{
					text:  "U",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "(f|v)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
	},
	gengreek: []rule{
		{
			pattern: "αυ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "af",
			phoneticRules: []phoneticRule{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern: "αυ",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phonetic: "af",
			phoneticRules: []phoneticRule{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern:  "αυ",
			phonetic: "av",
			phoneticRules: []phoneticRule{
				{
					text:  "av",
					langs: -1,
				},
			},
		},
		{
			pattern: "ευ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ef",
			phoneticRules: []phoneticRule{
				{
					text:  "ef",
					langs: -1,
				},
			},
		},
		{
			pattern: "ευ",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phonetic: "ef",
			phoneticRules: []phoneticRule{
				{
					text:  "ef",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ευ",
			phonetic: "ev",
			phoneticRules: []phoneticRule{
				{
					text:  "ev",
					langs: -1,
				},
			},
		},
		{
			pattern: "ηυ",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "if",
			phoneticRules: []phoneticRule{
				{
					text:  "if",
					langs: -1,
				},
			},
		},
		{
			pattern: "ηυ",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phonetic: "if",
			phoneticRules: []phoneticRule{
				{
					text:  "if",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ηυ",
			phonetic: "iv",
			phoneticRules: []phoneticRule{
				{
					text:  "iv",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ου",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "αι",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ει",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "οι",
			phonetic: "oj",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ωι",
			phonetic: "oj",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ηι",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "υι",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "γγ",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "(nj|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "nj",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "γγ",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "γγ",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "(ng|g)",
			phoneticRules: []phoneticRule{
				{
					text:  "ng",
					langs: -1,
				},
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "γγ",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "γκ",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "γκ",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "(nj|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "nj",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "γκ",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "γκ",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "(ng|g)",
			phoneticRules: []phoneticRule{
				{
					text:  "ng",
					langs: -1,
				},
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "γκ",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "γι",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "γι",
			phonetic: "(gi|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "γε",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "γε",
			phonetic: "(ge|je)",
			phoneticRules: []phoneticRule{
				{
					text:  "ge",
					langs: -1,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "κζ",
			phonetic: "gz",
			phoneticRules: []phoneticRule{
				{
					text:  "gz",
					langs: -1,
				},
			},
		},
		{
			pattern:  "τζ",
			phonetic: "dz",
			phoneticRules: []phoneticRule{
				{
					text:  "dz",
					langs: -1,
				},
			},
		},
		{
			pattern: "σ",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(β|γ|δ|μ|ν|ρ)"),
			},
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "μβ",
			phonetic: "(mb|b)",
			phoneticRules: []phoneticRule{
				{
					text:  "mb",
					langs: -1,
				},
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "μπ",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "μπ",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "mb",
			phoneticRules: []phoneticRule{
				{
					text:  "mb",
					langs: -1,
				},
			},
		},
		{
			pattern:  "μπ",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "ντ",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "ντ",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "(nd|nt)",
			phoneticRules: []phoneticRule{
				{
					text:  "nd",
					langs: -1,
				},
				{
					text:  "nt",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ντ",
			phonetic: "(nt|d)",
			phoneticRules: []phoneticRule{
				{
					text:  "nt",
					langs: -1,
				},
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ά",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "έ",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ή",
			phonetic: "(i|e)",
			phoneticRules: []phoneticRule{
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
			pattern:  "ί",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ό",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ύ",
			phonetic: "(Q|i|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ώ",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ΰ",
			phonetic: "(Q|i|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ϋ",
			phonetic: "(Q|i|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ϊ",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "α",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "β",
			phonetic: "(v|b)",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "γ",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "δ",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ε",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ζ",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "η",
			phonetic: "(i|e)",
			phoneticRules: []phoneticRule{
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
			pattern:  "ι",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "κ",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "λ",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "μ",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ν",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ξ",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ο",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "π",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ρ",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "σ",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ς",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "τ",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "υ",
			phonetic: "(Q|i|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "φ",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "θ",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "χ",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ψ",
			phonetic: "ps",
			phoneticRules: []phoneticRule{
				{
					text:  "ps",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ω",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
	},
	gengreeklatin: []rule{
		{
			pattern: "au",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "af",
			phoneticRules: []phoneticRule{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern: "au",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[kpstfh]"),
			},
			phonetic: "af",
			phoneticRules: []phoneticRule{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern:  "au",
			phonetic: "av",
			phoneticRules: []phoneticRule{
				{
					text:  "av",
					langs: -1,
				},
			},
		},
		{
			pattern: "eu",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ef",
			phoneticRules: []phoneticRule{
				{
					text:  "ef",
					langs: -1,
				},
			},
		},
		{
			pattern: "eu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[kpstfh]"),
			},
			phonetic: "ef",
			phoneticRules: []phoneticRule{
				{
					text:  "ef",
					langs: -1,
				},
			},
		},
		{
			pattern:  "eu",
			phonetic: "ev",
			phoneticRules: []phoneticRule{
				{
					text:  "ev",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ou",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: "gge",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(nje|je)",
			phoneticRules: []phoneticRule{
				{
					text:  "nje",
					langs: -1,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "(nj|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "nj",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ni|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "ni",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gge",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ggi",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "gg",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ng|g)",
			phoneticRules: []phoneticRule{
				{
					text:  "ng",
					langs: -1,
				},
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gg",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "gk",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "gke",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(nje|je)",
			phoneticRules: []phoneticRule{
				{
					text:  "nje",
					langs: -1,
				},
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: "gki",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ni|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "ni",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gke",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gki",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "gk",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ng|g)",
			phoneticRules: []phoneticRule{
				{
					text:  "ng",
					langs: -1,
				},
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gk",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "nghi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
			phoneticRules: []phoneticRule{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "nghi",
			phonetic: "(Ngi|Ni)",
			phoneticRules: []phoneticRule{
				{
					text:  "Ngi",
					langs: -1,
				},
				{
					text:  "Ni",
					langs: -1,
				},
			},
		},
		{
			pattern: "nghe",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
			phoneticRules: []phoneticRule{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "nghe",
			phonetic: "(Nje|Nge)",
			phoneticRules: []phoneticRule{
				{
					text:  "Nje",
					langs: -1,
				},
				{
					text:  "Nge",
					langs: -1,
				},
			},
		},
		{
			pattern: "ghi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ghi",
			phonetic: "(gi|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ghe",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ghe",
			phonetic: "(je|ge)",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ngh",
			phonetic: "Ng",
			phoneticRules: []phoneticRule{
				{
					text:  "Ng",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gh",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: "ngi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
			phoneticRules: []phoneticRule{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ngi",
			phonetic: "(Ngi|Ni)",
			phoneticRules: []phoneticRule{
				{
					text:  "Ngi",
					langs: -1,
				},
				{
					text:  "Ni",
					langs: -1,
				},
			},
		},
		{
			pattern: "nge",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
			phoneticRules: []phoneticRule{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "nge",
			phonetic: "(Nje|Nge)",
			phoneticRules: []phoneticRule{
				{
					text:  "Nje",
					langs: -1,
				},
				{
					text:  "Nge",
					langs: -1,
				},
			},
		},
		{
			pattern: "gi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gi",
			phonetic: "(gi|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "gi",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ge",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ge",
			phonetic: "(je|ge)",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
				{
					text:  "ge",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ng",
			phonetic: "Ng",
			phoneticRules: []phoneticRule{
				{
					text:  "Ng",
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
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yi",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "dh",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "dj",
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ph",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "th",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "kz",
			phonetic: "gz",
			phoneticRules: []phoneticRule{
				{
					text:  "gz",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tz",
			phonetic: "dz",
			phoneticRules: []phoneticRule{
				{
					text:  "dz",
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
				pattern:          regexp.MustCompile("^[bgdmnr]"),
			},
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "mb",
			phonetic: "(mb|b)",
			phoneticRules: []phoneticRule{
				{
					text:  "mb",
					langs: -1,
				},
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "mp",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "mp",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "mp",
			phoneticRules: []phoneticRule{
				{
					text:  "mp",
					langs: -1,
				},
			},
		},
		{
			pattern:  "mp",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: "nt",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: "nt",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(nd|nt)",
			phoneticRules: []phoneticRule{
				{
					text:  "nd",
					langs: -1,
				},
				{
					text:  "nt",
					langs: -1,
				},
			},
		},
		{
			pattern:  "nt",
			phonetic: "(nt|d)",
			phoneticRules: []phoneticRule{
				{
					text:  "nt",
					langs: -1,
				},
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "óu",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ý",
			phonetic: "(i|Q|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "(b|v)",
			phoneticRules: []phoneticRule{
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
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "(j|Z)",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ο",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "(i|Q|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "Q",
					langs: -1,
				},
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genhebrew: []rule{
		{
			pattern:  "אי",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "עי",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "עו",
			phonetic: "VV",
			phoneticRules: []phoneticRule{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern:  "או",
			phonetic: "VV",
			phoneticRules: []phoneticRule{
				{
					text:  "VV",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ג׳",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ד׳",
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "א",
			phonetic: "L",
			phoneticRules: []phoneticRule{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ב",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ג",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ד",
			phonetic: "d",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "1",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ה",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "וו",
			phonetic: "V",
			phoneticRules: []phoneticRule{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern:  "וי",
			phonetic: "WW",
			phoneticRules: []phoneticRule{
				{
					text:  "WW",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ו",
			phonetic: "W",
			phoneticRules: []phoneticRule{
				{
					text:  "W",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ז",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ח",
			phonetic: "X",
			phoneticRules: []phoneticRule{
				{
					text:  "X",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ט",
			phonetic: "T",
			phoneticRules: []phoneticRule{
				{
					text:  "T",
					langs: -1,
				},
			},
		},
		{
			pattern:  "יי",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "י",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ך",
			phonetic: "X",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern:  "כ",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ל",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ם",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "מ",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ן",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "נ",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ס",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ע",
			phonetic: "L",
			phoneticRules: []phoneticRule{
				{
					text:  "L",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ף",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "פ",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ץ",
			phonetic: "C",
			phoneticRules: []phoneticRule{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern:  "צ",
			phonetic: "C",
			phoneticRules: []phoneticRule{
				{
					text:  "C",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ק",
			phonetic: "K",
			phoneticRules: []phoneticRule{
				{
					text:  "K",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ר",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ש",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ת",
			phonetic: "TB",
			phoneticRules: []phoneticRule{
				{
					text:  "TB",
					langs: -1,
				},
			},
		},
	},
	genhungarian: []rule{
		{
			pattern:  "sz",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zs",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cs",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "(oj|aj)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "ej",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ee",
			phonetic: "(ej|e)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
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
			phonetic: "(ej|eli)",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gy",
			phonetic: "(d|gi)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ny",
			phonetic: "(n|ni)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "tj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ty",
			phonetic: "(t|ti)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ö",
			phonetic: "Y",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ő",
			phonetic: "Y",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "Q",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ű",
			phonetic: "Q",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "(S|s)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genitalian: []rule{
		{
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gli",
			phonetic: "(l|gli)",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
				{
					text:  "gli",
					langs: -1,
				},
			},
		},
		{
			pattern: "gn",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "(n|nj|gn)",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
				{
					text:  "nj",
					langs: -1,
				},
				{
					text:  "gn",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gni",
			phonetic: "(ni|gni)",
			phoneticRules: []phoneticRule{
				{
					text:  "ni",
					langs: -1,
				},
				{
					text:  "gni",
					langs: -1,
				},
			},
		},
		{
			pattern: "gi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "gg",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ci",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "tS",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "sc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
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
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
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
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
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
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "�",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "�",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "�",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "�",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "(Z|dZ|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "(ts|dz)",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
				{
					text:  "dz",
					langs: -1,
				},
			},
		},
	},
	genlatvian: []rule{
		{
			pattern:  "č",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ģ",
			phonetic: "(d|dj)",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
				{
					text:  "dj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ķ",
			phonetic: "(t|tj)",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
				{
					text:  "tj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ļ",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "š",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ņ",
			phonetic: "(n|nj)",
			phoneticRules: []phoneticRule{
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
			pattern:  "ž",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ā",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ē",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ī",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ū",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ai",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern:  "io",
			phonetic: "jo",
			phoneticRules: []phoneticRule{
				{
					text:  "jo",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iu",
			phonetic: "ju",
			phoneticRules: []phoneticRule{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ui",
			phonetic: "uj",
			phoneticRules: []phoneticRule{
				{
					text:  "uj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genpolish: []rule{
		{
			pattern: "ska",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ski",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "rjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
			phoneticRules: []phoneticRule{
				{
					text:  "djevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
			phoneticRules: []phoneticRule{
				{
					text:  "tjevitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
			phoneticRules: []phoneticRule{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
			phoneticRules: []phoneticRule{
				{
					text:  "evitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
			phoneticRules: []phoneticRule{
				{
					text:  "ovitS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "icz",
			phonetic: "itS",
			phoneticRules: []phoneticRule{
				{
					text:  "itS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "cz",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sz",
			phonetic: "S",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "Ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "źe",
			phonetic: "(Ze|ze)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "Zi",
					langs: -1,
				},
			},
		},
		{
			pattern:  "źi",
			phonetic: "(Zi|zi)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ź",
			phonetic: "(Z|z)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ń",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "(u|o)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "em",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ą",
			phonetic: "on",
			phoneticRules: []phoneticRule{
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ę",
			phonetic: "en",
			phoneticRules: []phoneticRule{
				{
					text:  "en",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ije",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yje",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iie",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yie",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iye",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yye",
			phonetic: "je",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "rie",
			phonetic: "rje",
			phoneticRules: []phoneticRule{
				{
					text:  "rje",
					langs: -1,
				},
			},
		},
		{
			pattern:  "die",
			phonetic: "dje",
			phoneticRules: []phoneticRule{
				{
					text:  "dje",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tie",
			phonetic: "tje",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "F",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "au",
			phonetic: "au",
			phoneticRules: []phoneticRule{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ei",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ey",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ej",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ai",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ay",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aj",
			phonetic: "aj",
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "P",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(h|x)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genportuguese: []rule{
		{
			pattern:  "kh",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ss",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "sc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "sç",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
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
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "s",
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
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
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phonetic: "(Z|S)",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
				{
					text:  "S",
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
			phonetic: "(Z|s|S)",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "S",
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
				pattern:          regexp.MustCompile("^[bdgv]"),
			},
			phonetic: "(Z|z)",
			phoneticRules: []phoneticRule{
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
			pattern: "z",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phonetic: "(s|S|z)",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "z",
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
				pattern:          regexp.MustCompile("^[eiu]"),
			},
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
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
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phonetic: "gv",
			phoneticRules: []phoneticRule{
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
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiu]"),
			},
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: "qu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phonetic: "kv",
			phoneticRules: []phoneticRule{
				{
					text:  "kv",
					langs: -1,
				},
			},
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o|u)",
			phoneticRules: []phoneticRule{
				{
					text:  "vo",
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
			pattern: "u",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lh",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "nh",
			phonetic: "nj",
			phoneticRules: []phoneticRule{
				{
					text:  "nj",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "(ez|eS|eks)",
			phoneticRules: []phoneticRule{
				{
					text:  "ez",
					langs: -1,
				},
				{
					text:  "eS",
					langs: -1,
				},
				{
					text:  "eks",
					langs: -1,
				},
			},
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[cs]"),
			},
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
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
				pattern:          regexp.MustCompile("[aáuiíoóeéê]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("^[aeiíou]"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
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
				pattern:          regexp.MustCompile("^[bcdfglnprstv]"),
			},
			phonetic: "(m|n)",
			phoneticRules: []phoneticRule{
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
			pattern: "m",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(m|n)",
			phoneticRules: []phoneticRule{
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
			pattern:  "ão",
			phonetic: "(au|an|on)",
			phoneticRules: []phoneticRule{
				{
					text:  "au",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ãe",
			phonetic: "(aj|an)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ãi",
			phonetic: "(aj|an)",
			phoneticRules: []phoneticRule{
				{
					text:  "aj",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
			},
		},
		{
			pattern:  "õe",
			phonetic: "(oj|on)",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
					langs: -1,
				},
				{
					text:  "on",
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
				pattern:          regexp.MustCompile("[aáuoóeéê]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "â",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "à",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ã",
			phonetic: "(a|an|on)",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
				{
					text:  "an",
					langs: -1,
				},
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ê",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ô",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "õ",
			phonetic: "(o|on)",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
				{
					text:  "on",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "(e|i)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "(o|u)",
			phoneticRules: []phoneticRule{
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
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "(S|ks)",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genromanian: []rule{
		{
			pattern:  "ce",
			phonetic: "tSe",
			phoneticRules: []phoneticRule{
				{
					text:  "tSe",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ci",
			phonetic: "(tSi|tS)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "x",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gi",
			phonetic: "(dZi|dZ)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gh",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ş",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "î",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ea",
			phonetic: "ja",
			phoneticRules: []phoneticRule{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ă",
			phonetic: "(e|a)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "(x|h)",
			phoneticRules: []phoneticRule{
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
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genrussian: []rule{
		{
			pattern: "yna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(in|ina)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsyu",
			phonetic: "tsu",
			phoneticRules: []phoneticRule{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsia",
			phonetic: "tsa",
			phoneticRules: []phoneticRule{
				{
					text:  "tsa",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsie",
			phonetic: "tse",
			phoneticRules: []phoneticRule{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsio",
			phonetic: "tso",
			phoneticRules: []phoneticRule{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsye",
			phonetic: "tse",
			phoneticRules: []phoneticRule{
				{
					text:  "tse",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsyo",
			phonetic: "tso",
			phoneticRules: []phoneticRule{
				{
					text:  "tso",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tsiu",
			phonetic: "tsu",
			phoneticRules: []phoneticRule{
				{
					text:  "tsu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sie",
			phonetic: "se",
			phoneticRules: []phoneticRule{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sio",
			phonetic: "so",
			phoneticRules: []phoneticRule{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zie",
			phonetic: "ze",
			phoneticRules: []phoneticRule{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zio",
			phonetic: "zo",
			phoneticRules: []phoneticRule{
				{
					text:  "zo",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sye",
			phonetic: "se",
			phoneticRules: []phoneticRule{
				{
					text:  "se",
					langs: -1,
				},
			},
		},
		{
			pattern:  "syo",
			phonetic: "so",
			phoneticRules: []phoneticRule{
				{
					text:  "so",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zye",
			phonetic: "ze",
			phoneticRules: []phoneticRule{
				{
					text:  "ze",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zyo",
			phonetic: "zo",
			phoneticRules: []phoneticRule{
				{
					text:  "zo",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "gin",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gg",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "sh",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "zh",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "ts",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tz",
			phonetic: "(ts|tz)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
			phoneticRules: []phoneticRule{
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
			pattern: "s",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "s",
				suffix:           "",
			},
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lya",
			phonetic: "la",
			phoneticRules: []phoneticRule{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lyu",
			phonetic: "lu",
			phoneticRules: []phoneticRule{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lia",
			phonetic: "la",
			phoneticRules: []phoneticRule{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "liu",
			phonetic: "lu",
			phoneticRules: []phoneticRule{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lja",
			phonetic: "la",
			phoneticRules: []phoneticRule{
				{
					text:  "la",
					langs: -1,
				},
			},
		},
		{
			pattern:  "lju",
			phonetic: "lu",
			phoneticRules: []phoneticRule{
				{
					text:  "lu",
					langs: -1,
				},
			},
		},
		{
			pattern:  "le",
			phonetic: "(lo|lE)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ie",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iye",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "iie",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yje",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ye",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yye",
			phonetic: "je",
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yie",
			phonetic: "je",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "io",
			phonetic: "(jo|e)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern:  "oo",
			phonetic: "(oo|u)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "\"",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "aue",
			phonetic: "aue",
			phoneticRules: []phoneticRule{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "E",
			phoneticRules: []phoneticRule{
				{
					text:  "E",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "I",
			phoneticRules: []phoneticRule{
				{
					text:  "I",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genspanish: []rule{
		{
			pattern:  "ñ",
			phonetic: "(n|nj)",
			phoneticRules: []phoneticRule{
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
			pattern:  "ny",
			phonetic: "nj",
			phoneticRules: []phoneticRule{
				{
					text:  "nj",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ç",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: "ig",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phonetic: "(tS|ig)",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
				{
					text:  "ig",
					langs: -1,
				},
			},
		},
		{
			pattern: "ix",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tx",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tj",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tj",
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "tg",
			phonetic: "(tg|dZ)",
			phoneticRules: []phoneticRule{
				{
					text:  "tg",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(tS|dZ)",
			phoneticRules: []phoneticRule{
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
			pattern:  "bh",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
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
				pattern:          regexp.MustCompile("[dgt]$"),
			},
			phonetic: "",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bpvf]"),
			},
			phonetic: "(m|n)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
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
			phonetic: "(x|g|dZ)",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
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
			pattern:  "qu",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
			phoneticRules: []phoneticRule{
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
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "á",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "é",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "í",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ó",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ú",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "à",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "è",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ò",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "B",
			phoneticRules: []phoneticRule{
				{
					text:  "B",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "(x|Z)",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: -1,
				},
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "V",
			phoneticRules: []phoneticRule{
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "(ks|gz|S)",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
				{
					text:  "gz",
					langs: -1,
				},
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "(i|j)",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "(z|s)",
			phoneticRules: []phoneticRule{
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
	},
	genturkish: []rule{
		{
			pattern:  "ç",
			phonetic: "tS",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ğ",
			phonetic: "",
			phoneticRules: []phoneticRule{
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ş",
			phonetic: "S",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ü",
			phonetic: "Q",
			phoneticRules: []phoneticRule{
				{
					text:  "Q",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ö",
			phonetic: "Y",
			phoneticRules: []phoneticRule{
				{
					text:  "Y",
					langs: -1,
				},
			},
		},
		{
			pattern:  "ı",
			phonetic: "(e|i|)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
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
			pattern:  "a",
			phonetic: "a",
			phoneticRules: []phoneticRule{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern:  "b",
			phonetic: "b",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern:  "c",
			phonetic: "dZ",
			phoneticRules: []phoneticRule{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern:  "d",
			phonetic: "d",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern:  "e",
			phonetic: "e",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern:  "f",
			phonetic: "f",
			phoneticRules: []phoneticRule{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern:  "g",
			phonetic: "g",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern:  "h",
			phonetic: "h",
			phoneticRules: []phoneticRule{
				{
					text:  "h",
					langs: -1,
				},
			},
		},
		{
			pattern:  "i",
			phonetic: "i",
			phoneticRules: []phoneticRule{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern:  "j",
			phonetic: "Z",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern:  "k",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "l",
			phonetic: "l",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern:  "m",
			phonetic: "m",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern:  "n",
			phonetic: "n",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern:  "o",
			phonetic: "o",
			phoneticRules: []phoneticRule{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern:  "p",
			phonetic: "p",
			phoneticRules: []phoneticRule{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern:  "q",
			phonetic: "k",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern:  "r",
			phonetic: "r",
			phoneticRules: []phoneticRule{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern:  "s",
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern:  "t",
			phonetic: "t",
			phoneticRules: []phoneticRule{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern:  "u",
			phonetic: "u",
			phoneticRules: []phoneticRule{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "w",
			phonetic: "v",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern:  "x",
			phonetic: "ks",
			phoneticRules: []phoneticRule{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern:  "y",
			phonetic: "j",
			phoneticRules: []phoneticRule{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "z",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
}

var genLangRules = []langRule{
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("^o’"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("^o'"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "mc",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "fitz",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ceau",
			prefix:           "",
			suffix:           "",
		},
		langs:  65600,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eau",
			prefix:           "",
			suffix:           "",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ault",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "oult",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "eux",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "eix",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "glou",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "uu",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "tx",
			prefix:           "",
			suffix:           "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "witz",
			prefix:           "",
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
			suffix:           "tz",
		},
		langs:  131232,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "tz",
			suffix:           "",
		},
		langs:  131104,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "poulos",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "pulos",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "iou",
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
			suffix:           "sj",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "sj",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "güe",
			prefix:           "",
			suffix:           "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "güi",
			prefix:           "",
			suffix:           "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ghe",
			prefix:           "",
			suffix:           "",
		},
		langs:  66048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ghi",
			prefix:           "",
			suffix:           "",
		},
		langs:  66048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "escu",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "esco",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "vici",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "schi",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ii",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "iy",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yy",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yi",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "rz",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "rz",
		},
		langs:  16512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[bcdfgklmnpstwz]rz"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("rz[bcdfghklmnpstw]"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cki",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ska",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cka",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ae",
			prefix:           "",
			suffix:           "",
		},
		langs:  131232,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "oe",
			prefix:           "",
			suffix:           "",
		},
		langs:  131312,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "th",
		},
		langs:  160,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "th",
			suffix:           "",
		},
		langs:  672,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "mann",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cz",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cy",
			prefix:           "",
			suffix:           "",
		},
		langs:  16896,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "niew",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "etti",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "eti",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ati",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ato",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[aoei]no$"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[aoei]ni$"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "esi",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "oli",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "field",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "stein",
			prefix:           "",
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
			suffix:           "heim",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "heimer",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "thal",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "zweig",
			prefix:           "",
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
			suffix:           "",
			pattern:          regexp.MustCompile("[aeou]h"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "äh",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "öh",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "üh",
			prefix:           "",
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
			suffix:           "",
			pattern:          regexp.MustCompile("[ln]h[ao]$"),
		},
		langs:  32768,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[ln]h[aou]"),
		},
		langs:  819416,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "chsch",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "tsch",
			prefix:           "",
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
			suffix:           "sch",
		},
		langs:  131200,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "sch",
			suffix:           "",
		},
		langs:  131200,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ck",
		},
		langs:  160,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "c",
		},
		langs:  608264,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sz",
			prefix:           "",
			suffix:           "",
		},
		langs:  18432,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cs",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "cs",
			suffix:           "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dzs",
			prefix:           "",
			suffix:           "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "zs",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "zs",
			suffix:           "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "wl",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "wr",
			suffix:           "",
		},
		langs:  16560,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gy",
		},
		langs:  2048,
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
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gy",
			prefix:           "",
			suffix:           "",
		},
		langs:  133696,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "guy",
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
			suffix:           "",
			pattern:          regexp.MustCompile("gu[ei]"),
		},
		langs:  294976,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gu[ao]"),
		},
		langs:  294912,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gi[aou]"),
		},
		langs:  4608,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ly",
			prefix:           "",
			suffix:           "",
		},
		langs:  150016,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ny",
			prefix:           "",
			suffix:           "",
		},
		langs:  412160,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ty",
			prefix:           "",
			suffix:           "",
		},
		langs:  150016,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ā",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ć",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ç",
			prefix:           "",
			suffix:           "",
		},
		langs:  819264,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "č",
			prefix:           "",
			suffix:           "",
		},
		langs:  8200,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ď",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ē",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ğ",
			prefix:           "",
			suffix:           "",
		},
		langs:  524288,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ģ",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ī",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ķ",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ļ",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ł",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ņ",
			prefix:           "",
			suffix:           "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ń",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ñ",
			prefix:           "",
			suffix:           "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ň",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ř",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ś",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ş",
			prefix:           "",
			suffix:           "",
		},
		langs:  589824,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "š",
			prefix:           "",
			suffix:           "",
		},
		langs:  8200,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ţ",
			prefix:           "",
			suffix:           "",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ť",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ź",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ž",
			prefix:           "",
			suffix:           "",
		},
		langs:  8200,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ż",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ß",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ä",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "á",
			prefix:           "",
			suffix:           "",
		},
		langs:  297480,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "â",
			prefix:           "",
			suffix:           "",
		},
		langs:  98368,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ă",
			prefix:           "",
			suffix:           "",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ą",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "à",
			prefix:           "",
			suffix:           "",
		},
		langs:  32768,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ã",
			prefix:           "",
			suffix:           "",
		},
		langs:  32768,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ę",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "é",
			prefix:           "",
			suffix:           "",
		},
		langs:  2632,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "è",
			prefix:           "",
			suffix:           "",
		},
		langs:  266304,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ê",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ě",
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
		langs:  32832,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "í",
			prefix:           "",
			suffix:           "",
		},
		langs:  297480,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "î",
			prefix:           "",
			suffix:           "",
		},
		langs:  65600,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ı",
			prefix:           "",
			suffix:           "",
		},
		langs:  524288,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ó",
			prefix:           "",
			suffix:           "",
		},
		langs:  317960,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ö",
			prefix:           "",
			suffix:           "",
		},
		langs:  526464,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ô",
			prefix:           "",
			suffix:           "",
		},
		langs:  32832,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "õ",
			prefix:           "",
			suffix:           "",
		},
		langs:  34816,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ò",
			prefix:           "",
			suffix:           "",
		},
		langs:  266240,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ű",
			prefix:           "",
			suffix:           "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ú",
			prefix:           "",
			suffix:           "",
		},
		langs:  297480,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ü",
			prefix:           "",
			suffix:           "",
		},
		langs:  821376,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ù",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ů",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ý",
			prefix:           "",
			suffix:           "",
		},
		langs:  520,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "а",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ё",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "о",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "е",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "и",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "у",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ы",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "э",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ю",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "я",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "α",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ε",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "η",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ι",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ο",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "υ",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ω",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ا",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ب",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ت",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ث",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ج",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ح",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("خ'"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "د",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ذ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ر",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ز",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "س",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ش",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ص",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ض",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ط",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ظ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ع",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "غ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ف",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ق",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ك",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ل",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "م",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ن",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ه",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "و",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ي",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "آ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "إ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "أ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ؤ",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ئ",
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
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ב",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ג",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ד",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ה",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ו",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ז",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ח",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ט",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "י",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "כ",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ל",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "מ",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "נ",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ס",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ע",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "פ",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "צ",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ק",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ר",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ש",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ת",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "a",
			prefix:           "",
			suffix:           "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "o",
			prefix:           "",
			suffix:           "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "e",
			prefix:           "",
			suffix:           "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "i",
			prefix:           "",
			suffix:           "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "y",
			prefix:           "",
			suffix:           "",
		},
		langs:  75030,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "u",
			prefix:           "",
			suffix:           "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "j",
			prefix:           "",
			suffix:           "",
		},
		langs:  4096,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("j[^aoeiuy]"),
		},
		langs:  295488,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "g",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "k",
			prefix:           "",
			suffix:           "",
		},
		langs:  364608,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "q",
			prefix:           "",
			suffix:           "",
		},
		langs:  748056,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "v",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "w",
			prefix:           "",
			suffix:           "",
		},
		langs:  993864,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "x",
			prefix:           "",
			suffix:           "",
		},
		langs:  534552,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dj",
			prefix:           "",
			suffix:           "",
		},
		langs:  786432,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("v[^aoeiu]"),
		},
		langs:  128,
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
		langs:  128,
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
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dzi",
			prefix:           "",
			suffix:           "",
		},
		langs:  524512,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ou",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("a[eiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("ö[eaiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("ü[eaiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("e[aiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("i[aeou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("o[aieu]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("u[aieo]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "aj",
			prefix:           "",
			suffix:           "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ej",
			prefix:           "",
			suffix:           "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "oj",
			prefix:           "",
			suffix:           "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "uj",
			prefix:           "",
			suffix:           "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eu",
			prefix:           "",
			suffix:           "",
		},
		langs:  147456,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ky",
			prefix:           "",
			suffix:           "",
		},
		langs:  16384,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "kie",
			prefix:           "",
			suffix:           "",
		},
		langs:  262720,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gie",
			prefix:           "",
			suffix:           "",
		},
		langs:  360960,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("ch[aou]"),
		},
		langs:  4096,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ch",
			prefix:           "",
			suffix:           "",
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "son",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("sc[ei]"),
		},
		langs:  64,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sch",
			prefix:           "",
			suffix:           "",
		},
		langs:  280640,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "h",
			suffix:           "",
		},
		langs:  131072,
		accept: false,
	},
}

var genFinalRules = finalRules{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				phonetic: "b",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				phonetic: "v",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				phonetic: "g",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				phonetic: "d",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[aA]"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					suffix:           "A",
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "j",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "j",
					suffix:           "",
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "vanden",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(vanden|)",
				phoneticRules: []phoneticRule{
					{
						text:  "vanden",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "vander",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(vander|)",
				phoneticRules: []phoneticRule{
					{
						text:  "vander",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "van",
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
					pattern:          regexp.MustCompile("^[bp]"),
				},
				phonetic: "(vam|[16])",
				phoneticRules: []phoneticRule{
					{
						text:  "vam",
						langs: -1,
					},
					{
						text:  "",
						langs: 16,
					},
				},
			},
			{
				pattern: "van",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(van|[16])",
				phoneticRules: []phoneticRule{
					{
						text:  "van",
						langs: -1,
					},
					{
						text:  "",
						langs: 16,
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
				phoneticRules: []phoneticRule{
					{
						text:  "m",
						langs: -1,
					},
				},
			},
			{
				pattern:  "h",
				phonetic: "",
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "H",
				phonetic: "(x|)",
				phoneticRules: []phoneticRule{
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
				pattern: "sen",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(zn|zon)",
				phoneticRules: []phoneticRule{
					{
						text:  "zn",
						langs: -1,
					},
					{
						text:  "zon",
						langs: -1,
					},
				},
			},
			{
				pattern: "sen",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(sn|son)",
				phoneticRules: []phoneticRule{
					{
						text:  "sn",
						langs: -1,
					},
					{
						text:  "son",
						langs: -1,
					},
				},
			},
			{
				pattern: "sEn",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(zn|zon)",
				phoneticRules: []phoneticRule{
					{
						text:  "zn",
						langs: -1,
					},
					{
						text:  "zon",
						langs: -1,
					},
				},
			},
			{
				pattern: "sEn",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(sn|son)",
				phoneticRules: []phoneticRule{
					{
						text:  "sn",
						langs: -1,
					},
					{
						text:  "son",
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oue",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AvE",
				phonetic: "(D|AvE)",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Oji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Oje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Eji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Eje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Uji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Uje",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "iji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ije",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Iji",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjI",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ije",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjE",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ajU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ajo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ojO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Oja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ojo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "OjO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ejO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Eja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ejo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "EjO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ujO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Uja",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ujo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "UjO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ija",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ijO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ija",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjA",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Ijo",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "IjO",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Aju",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "AjU",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "j",
				phonetic: "i",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "lYnder",
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
				phoneticRules: []phoneticRule{
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
				pattern: "bUrk",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				pattern: "Burk",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
				phoneticRules: []phoneticRule{
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
				pattern: "BUrk",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
				phoneticRules: []phoneticRule{
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
				pattern: "Burg",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
				phoneticRules: []phoneticRule{
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
				pattern: "BUrg",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern:  "dZ",
				phonetic: "z",
				phoneticRules: []phoneticRule{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
			{
				pattern:  "Z",
				phonetic: "z",
				phoneticRules: []phoneticRule{
					{
						text:  "z",
						langs: -1,
					},
				},
			},
		},
		second: map[int64][]rule{
			int64(genany): []rule{
				{
					pattern:  "mb",
					phonetic: "(mb|b[512])",
					phoneticRules: []phoneticRule{
						{
							text:  "mb",
							langs: -1,
						},
						{
							text:  "b",
							langs: 512,
						},
					},
				},
				{
					pattern:  "mp",
					phonetic: "(mp|b[512])",
					phoneticRules: []phoneticRule{
						{
							text:  "mp",
							langs: -1,
						},
						{
							text:  "b",
							langs: 512,
						},
					},
				},
				{
					pattern:  "ng",
					phonetic: "(ng|g[512])",
					phoneticRules: []phoneticRule{
						{
							text:  "ng",
							langs: -1,
						},
						{
							text:  "g",
							langs: 512,
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
						pattern:          regexp.MustCompile("^[fktSs]"),
					},
					phonetic: "(p|f[262144])",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
						{
							text:  "f",
							langs: 262144,
						},
					},
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "p",
						suffix:           "",
					},
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(p|f[262144])",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
						{
							text:  "f",
							langs: 262144,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[pktSs]"),
					},
					phonetic: "(f|p[262144])",
					phoneticRules: []phoneticRule{
						{
							text:  "f",
							langs: -1,
						},
						{
							text:  "p",
							langs: 262144,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "f",
						suffix:           "",
					},
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(f|p[262144])",
					phoneticRules: []phoneticRule{
						{
							text:  "f",
							langs: -1,
						},
						{
							text:  "p",
							langs: 262144,
						},
					},
				},
				{
					pattern:  "B",
					phonetic: "(b|v[262144])",
					phoneticRules: []phoneticRule{
						{
							text:  "b",
							langs: -1,
						},
						{
							text:  "v",
							langs: 262144,
						},
					},
				},
				{
					pattern:  "V",
					phonetic: "(v|b[262144])",
					phoneticRules: []phoneticRule{
						{
							text:  "v",
							langs: -1,
						},
						{
							text:  "b",
							langs: 262144,
						},
					},
				},
				{
					pattern: "t",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(t|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "t",
							langs: -1,
						},
						{
							text:  "",
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
						suffix:           "n",
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(g|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "g",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: "k",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "n",
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(k|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "k",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: "p",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(p|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: "r",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[Ee]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(r|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "r",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: "s",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(s|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "s",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: "t",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiouAEIOU]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					phonetic: "(t|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "t",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
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
						pattern:          regexp.MustCompile("[aeiouAEIOU]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					phonetic: "(s|[64])",
					phoneticRules: []phoneticRule{
						{
							text:  "s",
							langs: -1,
						},
						{
							text:  "",
							langs: 64,
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(Q[128]|i|D[32])",
					phoneticRules: []phoneticRule{
						{
							text:  "Q",
							langs: 128,
						},
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "D",
							langs: 32,
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(ik|Qk[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "ik",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: 128,
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(sits|sQts[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "sits",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: 128,
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
					phoneticRules: []phoneticRule{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(Q[128]|i)",
					phoneticRules: []phoneticRule{
						{
							text:  "Q",
							langs: 128,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: "lEE",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(li|il[32])",
					phoneticRules: []phoneticRule{
						{
							text:  "li",
							langs: -1,
						},
						{
							text:  "il",
							langs: 32,
						},
					},
				},
				{
					pattern: "rEE",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(ri|ir[32])",
					phoneticRules: []phoneticRule{
						{
							text:  "ri",
							langs: -1,
						},
						{
							text:  "ir",
							langs: 32,
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
					phonetic: "(li|il[32]|lY[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "li",
							langs: -1,
						},
						{
							text:  "il",
							langs: 32,
						},
						{
							text:  "lY",
							langs: 128,
						},
					},
				},
				{
					pattern: "rE",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(ri|ir[32]|rY[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "ri",
							langs: -1,
						},
						{
							text:  "ir",
							langs: 32,
						},
						{
							text:  "rY",
							langs: 128,
						},
					},
				},
				{
					pattern:  "EE",
					phonetic: "(i|)",
					phoneticRules: []phoneticRule{
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
					pattern:  "ea",
					phonetic: "(D|a|i)",
					phoneticRules: []phoneticRule{
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
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					pattern:  "eu",
					phonetic: "(D|e|u)",
					phoneticRules: []phoneticRule{
						{
							text:  "D",
							langs: -1,
						},
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
					pattern:  "ai",
					phonetic: "(D|a|i)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(ia|io|iY[128])",
					phoneticRules: []phoneticRule{
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
							langs: 128,
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
					phonetic: "(a|o|Y[128]|D[32])",
					phoneticRules: []phoneticRule{
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
							langs: 128,
						},
						{
							text:  "D",
							langs: 32,
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
					phonetic: "(i|Y[128]|[32])",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 128,
						},
						{
							text:  "",
							langs: 32,
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
					phonetic: "(i|Y[128]|[32])",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 128,
						},
						{
							text:  "",
							langs: 32,
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(i|Y[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 128,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
					phoneticRules: []phoneticRule{
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
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phonetic: "o",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "(o|Y[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 128,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "o",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(a|o|Y[128])",
					phoneticRules: []phoneticRule{
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
							langs: 128,
						},
					},
				},
				{
					pattern:  "A",
					phonetic: "(a|o)",
					phoneticRules: []phoneticRule{
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
					pattern: "U",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "u",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(uk|Qk[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "uk",
							langs: -1,
						},
						{
							text:  "Qk",
							langs: 128,
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
					phoneticRules: []phoneticRule{
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
					phonetic: "(suts|sQts[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "suts",
							langs: -1,
						},
						{
							text:  "sQts",
							langs: 128,
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
					phoneticRules: []phoneticRule{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "(u|Q[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "u",
							langs: -1,
						},
						{
							text:  "Q",
							langs: 128,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "u",
					phoneticRules: []phoneticRule{
						{
							text:  "u",
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
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phonetic: "i",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "e",
					phonetic: "(i|Y[128])",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
						{
							text:  "Y",
							langs: 128,
						},
					},
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
					phoneticRules: []phoneticRule{
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
			int64(genarabic): []rule{
				{
					pattern:  "1a",
					phonetic: "(D|a)",
					phoneticRules: []phoneticRule{
						{
							text:  "D",
							langs: -1,
						},
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "1i",
					phonetic: "(D|i|e)",
					phoneticRules: []phoneticRule{
						{
							text:  "D",
							langs: -1,
						},
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
					pattern:  "1u",
					phonetic: "(D|u|o)",
					phoneticRules: []phoneticRule{
						{
							text:  "D",
							langs: -1,
						},
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
					pattern:  "j1",
					phonetic: "(ja|je|jo|ju|j)",
					phoneticRules: []phoneticRule{
						{
							text:  "ja",
							langs: -1,
						},
						{
							text:  "je",
							langs: -1,
						},
						{
							text:  "jo",
							langs: -1,
						},
						{
							text:  "ju",
							langs: -1,
						},
						{
							text:  "j",
							langs: -1,
						},
					},
				},
				{
					pattern:  "1",
					phonetic: "(a|e|i|o|u|)",
					phoneticRules: []phoneticRule{
						{
							text:  "a",
							langs: -1,
						},
						{
							text:  "e",
							langs: -1,
						},
						{
							text:  "i",
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
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern:  "u",
					phonetic: "(o|u)",
					phoneticRules: []phoneticRule{
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
					pattern:  "i",
					phonetic: "(i|e)",
					phoneticRules: []phoneticRule{
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
					pattern: "p",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "p",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
					},
				},
				{
					pattern:  "p",
					phonetic: "(p|b)",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
						{
							text:  "b",
							langs: -1,
						},
					},
				},
			},
			int64(genrussian): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
					phoneticRules: []phoneticRule{
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
			int64(gencyrillic): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
					phoneticRules: []phoneticRule{
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
			int64(genfrench): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genczech): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(gendutch): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genenglish): []rule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
			int64(gengerman): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "its",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(Q|i)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "(o|Y)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "uts",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "(u|Q)",
					phoneticRules: []phoneticRule{
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
			int64(gengreek): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(gengreeklatin): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "N",
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
			},
			int64(genhebrew): []rule{},
			int64(genhungarian): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genitalian): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genlatvian): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genpolish): []rule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
					phoneticRules: []phoneticRule{
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
			int64(genportuguese): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genromanian): []rule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
					phoneticRules: []phoneticRule{
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
			int64(genspanish): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					pattern:  "B",
					phonetic: "(b|v)",
					phoneticRules: []phoneticRule{
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
					pattern:  "V",
					phonetic: "(b|v)",
					phoneticRules: []phoneticRule{
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
			},
			int64(genturkish): []rule{
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
					phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				phonetic: "b",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				phonetic: "v",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				phonetic: "g",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				phonetic: "d",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
					pattern:          regexp.MustCompile("^[aA]"),
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
					suffix:           "A",
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "j",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "j",
					suffix:           "",
				},
				phonetic: "",
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "H",
				phonetic: "",
				phoneticRules: []phoneticRule{
					{
						text:  "",
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
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
				phoneticRules: []phoneticRule{
					{
						text:  "j",
						langs: -1,
					},
				},
			},
		},
		second: map[int64][]rule{
			int64(genany): []rule{
				{
					pattern: "EE",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "A",
					phonetic: "a",
					phoneticRules: []phoneticRule{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "o",
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "o",
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "u",
					phoneticRules: []phoneticRule{
						{
							text:  "u",
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
						pattern:          regexp.MustCompile("^[fktSs]"),
					},
					phonetic: "p",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "p",
						suffix:           "",
					},
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "p",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[pktSs]"),
					},
					phonetic: "f",
					phoneticRules: []phoneticRule{
						{
							text:  "f",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "f",
						suffix:           "",
					},
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "f",
					phoneticRules: []phoneticRule{
						{
							text:  "f",
							langs: -1,
						},
					},
				},
				{
					pattern:  "B",
					phonetic: "b",
					phoneticRules: []phoneticRule{
						{
							text:  "b",
							langs: -1,
						},
					},
				},
				{
					pattern:  "V",
					phonetic: "v",
					phoneticRules: []phoneticRule{
						{
							text:  "v",
							langs: -1,
						},
					},
				},
			},
			int64(genarabic): []rule{
				{
					pattern:  "1",
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
			},
			int64(genrussian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(gencyrillic): []rule{
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genczech): []rule{
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(gendutch): []rule{},
			int64(genenglish): []rule{
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genfrench): []rule{},
			int64(gengerman): []rule{
				{
					pattern: "EE",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "A",
					phonetic: "a",
					phoneticRules: []phoneticRule{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern:  "O",
					phonetic: "o",
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "o",
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "U",
					phonetic: "u",
					phoneticRules: []phoneticRule{
						{
							text:  "u",
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
						pattern:          regexp.MustCompile("^[fktSs]"),
					},
					phonetic: "p",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "p",
						suffix:           "",
					},
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "p",
					phoneticRules: []phoneticRule{
						{
							text:  "p",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[pktSs]"),
					},
					phonetic: "f",
					phoneticRules: []phoneticRule{
						{
							text:  "f",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "f",
						suffix:           "",
					},
					phonetic: "",
					phoneticRules: []phoneticRule{
						{
							text:  "",
							langs: -1,
						},
					},
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "f",
					phoneticRules: []phoneticRule{
						{
							text:  "f",
							langs: -1,
						},
					},
				},
				{
					pattern:  "B",
					phonetic: "b",
					phoneticRules: []phoneticRule{
						{
							text:  "b",
							langs: -1,
						},
					},
				},
				{
					pattern:  "V",
					phonetic: "v",
					phoneticRules: []phoneticRule{
						{
							text:  "v",
							langs: -1,
						},
					},
				},
			},
			int64(gengreek): []rule{},
			int64(gengreeklatin): []rule{
				{
					pattern:  "N",
					phonetic: "n",
					phoneticRules: []phoneticRule{
						{
							text:  "n",
							langs: -1,
						},
					},
				},
			},
			int64(genhebrew): []rule{},
			int64(genhungarian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genitalian): []rule{},
			int64(genlatvian): []rule{},
			int64(genpolish): []rule{
				{
					pattern:  "B",
					phonetic: "a",
					phoneticRules: []phoneticRule{
						{
							text:  "a",
							langs: -1,
						},
					},
				},
				{
					pattern:  "F",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "P",
					phonetic: "o",
					phoneticRules: []phoneticRule{
						{
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genportuguese): []rule{},
			int64(genromanian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
					phoneticRules: []phoneticRule{
						{
							text:  "e",
							langs: -1,
						},
					},
				},
				{
					pattern:  "I",
					phonetic: "i",
					phoneticRules: []phoneticRule{
						{
							text:  "i",
							langs: -1,
						},
					},
				},
			},
			int64(genspanish): []rule{
				{
					pattern:  "B",
					phonetic: "b",
					phoneticRules: []phoneticRule{
						{
							text:  "b",
							langs: -1,
						},
					},
				},
				{
					pattern:  "V",
					phonetic: "v",
					phoneticRules: []phoneticRule{
						{
							text:  "v",
							langs: -1,
						},
					},
				},
			},
			int64(genturkish): []rule{},
		},
	},
}

var genDiscards = []string{
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
