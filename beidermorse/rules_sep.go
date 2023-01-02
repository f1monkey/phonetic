// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type sepLang int64

const (
	sepany sepLang = 1 << iota
	sepfrench
	sephebrew
	sepitalian
	sepportuguese
	sepspanish
)

func (l sepLang) String() string {
	switch l {
	case sepany:
		return "any"
	case sepfrench:
		return "french"
	case sephebrew:
		return "hebrew"
	case sepitalian:
		return "italian"
	case sepportuguese:
		return "portuguese"
	case sepspanish:
		return "spanish"
	}
	return ""
}

const sepAll = sepfrench +
	sephebrew +
	sepitalian +
	sepportuguese +
	sepspanish

var sepRules = map[sepLang][]rule{
	sepany: []rule{
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
			phonetic: "(gli|l[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "gli",
					langs: -1,
				},
				{
					text:  "l",
					langs: 8,
				},
			},
		},
		{
			pattern:  "gni",
			phonetic: "(gni|ni[10])",
			phoneticRules: []phoneticRule{
				{
					text:  "gni",
					langs: -1,
				},
				{
					text:  "ni",
					langs: 10,
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
			phonetic: "(n[10]|nj[10]|gn)",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: 10,
				},
				{
					text:  "nj",
					langs: 10,
				},
				{
					text:  "gn",
					langs: -1,
				},
			},
		},
		{
			pattern:  "gh",
			phonetic: "(g|gh)",
			phoneticRules: []phoneticRule{
				{
					text:  "g",
					langs: -1,
				},
				{
					text:  "gh",
					langs: -1,
				},
			},
		},
		{
			pattern:  "dh",
			phonetic: "(d|dh)",
			phoneticRules: []phoneticRule{
				{
					text:  "d",
					langs: -1,
				},
				{
					text:  "dh",
					langs: -1,
				},
			},
		},
		{
			pattern:  "bh",
			phonetic: "(b|bh)",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
				{
					text:  "bh",
					langs: -1,
				},
			},
		},
		{
			pattern:  "th",
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
			pattern:  "lh",
			phonetic: "(l[16]|lh)",
			phoneticRules: []phoneticRule{
				{
					text:  "l",
					langs: 16,
				},
				{
					text:  "lh",
					langs: -1,
				},
			},
		},
		{
			pattern:  "nh",
			phonetic: "(n[16]|nh)",
			phoneticRules: []phoneticRule{
				{
					text:  "n",
					langs: 16,
				},
				{
					text:  "nh",
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
			phonetic: "(ig|tS[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "ig",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 32,
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
			phonetic: "(tg|dZ[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "tg",
					langs: -1,
				},
				{
					text:  "dZ",
					langs: 32,
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
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "y",
				suffix:           "",
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
			pattern: "gg",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(gZ[18]|dZ[40]|x[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "gZ",
					langs: 18,
				},
				{
					text:  "dZ",
					langs: 40,
				},
				{
					text:  "x",
					langs: 32,
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
			phonetic: "(Z[18]|dZ[40]|x[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "Z",
					langs: 18,
				},
				{
					text:  "dZ",
					langs: 40,
				},
				{
					text:  "x",
					langs: 32,
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
			pattern: "gue",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(k[2]|ge)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: 2,
				},
				{
					text:  "ge",
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
			pattern: "sc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(s|S[8])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "S",
					langs: 8,
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
			phonetic: "s",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
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
			pattern: "ch",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(k[8]|S[18]|tS[32]|dZ[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: 8,
				},
				{
					text:  "S",
					langs: 18,
				},
				{
					text:  "tS",
					langs: 32,
				},
				{
					text:  "dZ",
					langs: 32,
				},
			},
		},
		{
			pattern:  "ch",
			phonetic: "(S|tS[32]|dZ[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "S",
					langs: -1,
				},
				{
					text:  "tS",
					langs: 32,
				},
				{
					text:  "dZ",
					langs: 32,
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
			phonetic: "(tS[8]|si)",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: 8,
				},
				{
					text:  "si",
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
				pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			phonetic: "(tS[8]|ks[50])",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: 8,
				},
				{
					text:  "ks",
					langs: 50,
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
			phonetic: "(tS[8]|s[50])",
			phoneticRules: []phoneticRule{
				{
					text:  "tS",
					langs: 8,
				},
				{
					text:  "s",
					langs: 50,
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
			phonetic: "(s[32]|z[26])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: 32,
				},
				{
					text:  "z",
					langs: 26,
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
			phonetic: "(z|s|Z[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "Z",
					langs: 16,
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
			phonetic: "(s|ts[8]|S[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 8,
				},
				{
					text:  "S",
					langs: 16,
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
			phonetic: "(z|dz[8]|Z[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "dz",
					langs: 8,
				},
				{
					text:  "Z",
					langs: 16,
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
			phonetic: "(s|ts[8]|S[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "ts",
					langs: 8,
				},
				{
					text:  "S",
					langs: 16,
				},
			},
		},
		{
			pattern:  "z",
			phonetic: "(z|dz[8]|ts[8]|s[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "z",
					langs: -1,
				},
				{
					text:  "dz",
					langs: 8,
				},
				{
					text:  "ts",
					langs: 8,
				},
				{
					text:  "s",
					langs: 32,
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
			phonetic: "(k[2]|ke)",
			phoneticRules: []phoneticRule{
				{
					text:  "k",
					langs: 2,
				},
				{
					text:  "ke",
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
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "(ez[16]|eS[16]|eks|egz)",
			phoneticRules: []phoneticRule{
				{
					text:  "ez",
					langs: 16,
				},
				{
					text:  "eS",
					langs: 16,
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
			phonetic: "(e[16]|ek)",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: 16,
				},
				{
					text:  "ek",
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
				pattern:          regexp.MustCompile("^[cdglnrst]"),
			},
			phonetic: "(m|n[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
				{
					text:  "n",
					langs: 16,
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
			phonetic: "(m|n[48])",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
				{
					text:  "n",
					langs: 48,
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
			phonetic: "(m|n[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "m",
					langs: -1,
				},
				{
					text:  "n",
					langs: 16,
				},
			},
		},
		{
			pattern: "b",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(b|V[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
				{
					text:  "V",
					langs: 32,
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
			phonetic: "(v|B[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "B",
					langs: 32,
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
			pattern: "ouh",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phonetic: "(v[2]|uh)",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: 2,
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
			pattern: "ou",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aioe]"),
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
				pattern:          regexp.MustCompile("^[aie]"),
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
			pattern: "e",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(e|[2])",
			phoneticRules: []phoneticRule{
				{
					text:  "e",
					langs: -1,
				},
				{
					text:  "",
					langs: 2,
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
			phonetic: "(ou|u[2])",
			phoneticRules: []phoneticRule{
				{
					text:  "ou",
					langs: -1,
				},
				{
					text:  "u",
					langs: 2,
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
			phonetic: "(b|v[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
				{
					text:  "v",
					langs: 32,
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
			phonetic: "(x[32]|Z)",
			phoneticRules: []phoneticRule{
				{
					text:  "x",
					langs: 32,
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
			phonetic: "(s|S[16])",
			phoneticRules: []phoneticRule{
				{
					text:  "s",
					langs: -1,
				},
				{
					text:  "S",
					langs: 16,
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
			phonetic: "(v|b[32])",
			phoneticRules: []phoneticRule{
				{
					text:  "v",
					langs: -1,
				},
				{
					text:  "b",
					langs: 32,
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
			phonetic: "(ks|gz|S[48])",
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
					langs: 48,
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
	sepfrench: []rule{
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
			pattern: "ouh",
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
			phonetic: "oj",
			phoneticRules: []phoneticRule{
				{
					text:  "oj",
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
			pattern:  "ey",
			phonetic: "ej",
			phoneticRules: []phoneticRule{
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
	sephebrew: []rule{
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
			phonetic: "T",
			phoneticRules: []phoneticRule{
				{
					text:  "T",
					langs: -1,
				},
			},
		},
	},
	sepitalian: []rule{
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
	sepportuguese: []rule{
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
	sepspanish: []rule{
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
			pattern: "v",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(B|v)",
			phoneticRules: []phoneticRule{
				{
					text:  "B",
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
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(b|V)",
			phoneticRules: []phoneticRule{
				{
					text:  "b",
					langs: -1,
				},
				{
					text:  "V",
					langs: -1,
				},
			},
		},
		{
			pattern:  "v",
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
	},
}

var sepLangRules = []langRule{
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eau",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ou",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gni",
			prefix:           "",
			suffix:           "",
		},
		langs:  10,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "tx",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "tj",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gy",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "guy",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sh",
			prefix:           "",
			suffix:           "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "lh",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "nh",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ny",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gue",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gui",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gia",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gie",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gio",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "giu",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ñ",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "â",
			prefix:           "",
			suffix:           "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "á",
			prefix:           "",
			suffix:           "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "à",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ã",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ê",
			prefix:           "",
			suffix:           "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "í",
			prefix:           "",
			suffix:           "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "î",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ô",
			prefix:           "",
			suffix:           "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "õ",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ò",
			prefix:           "",
			suffix:           "",
		},
		langs:  40,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ú",
			prefix:           "",
			suffix:           "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ù",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ü",
			prefix:           "",
			suffix:           "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "א",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ב",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ג",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ד",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ה",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ו",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ז",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ח",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ט",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "י",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "כ",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ל",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "מ",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "נ",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ס",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ע",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "פ",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "צ",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ק",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ר",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ש",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ת",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "a",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "o",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "e",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "i",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "y",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "u",
			prefix:           "",
			suffix:           "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "kh",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gua",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "guo",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ç",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cha",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cho",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "chu",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "j",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dj",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sce",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sci",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ó",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "è",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: false,
	},
}

var sepFinalRules = finalRules{
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
				pattern:  "nm",
				phonetic: "m",
				phoneticRules: []phoneticRule{
					{
						text:  "m",
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
				pattern: "a",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "a",
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
				pattern: "e",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "e",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "i",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "o",
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
				pattern: "u",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "u",
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
				pattern:  "mbr",
				phonetic: "mr",
				phoneticRules: []phoneticRule{
					{
						text:  "mr",
						langs: -1,
					},
				},
			},
			{
				pattern:  "mpr",
				phonetic: "mr",
				phoneticRules: []phoneticRule{
					{
						text:  "mr",
						langs: -1,
					},
				},
			},
			{
				pattern: "bens",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(binz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "binz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "benS",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(binz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "binz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "ben",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bin|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bin",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "bar",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bar|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bar",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "abens",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(binz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "binz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "abenS",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(binz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "binz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "aben",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bin|bun|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bin",
						langs: -1,
					},
					{
						text:  "bun",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "abe",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bi|bu|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bi",
						langs: -1,
					},
					{
						text:  "bu",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "abi",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bi|bu|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bi",
						langs: -1,
					},
					{
						text:  "bu",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "abou",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bu|[2])",
				phoneticRules: []phoneticRule{
					{
						text:  "bu",
						langs: -1,
					},
					{
						text:  "",
						langs: 2,
					},
				},
			},
			{
				pattern: "abu",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bu|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bu",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "bou",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bu|[2])",
				phoneticRules: []phoneticRule{
					{
						text:  "bu",
						langs: -1,
					},
					{
						text:  "",
						langs: 2,
					},
				},
			},
			{
				pattern: "bu",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(bu|)",
				phoneticRules: []phoneticRule{
					{
						text:  "bu",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "ibn",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(ibn|)",
				phoneticRules: []phoneticRule{
					{
						text:  "ibn",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "els",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(ilz|lz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "ilz",
						langs: -1,
					},
					{
						text:  "lz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "elS",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(ilz|lz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "ilz",
						langs: -1,
					},
					{
						text:  "lz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "el",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(il|l|)",
				phoneticRules: []phoneticRule{
					{
						text:  "il",
						langs: -1,
					},
					{
						text:  "l",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "als",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(lz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "lz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "alS",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(lz|s)",
				phoneticRules: []phoneticRule{
					{
						text:  "lz",
						langs: -1,
					},
					{
						text:  "s",
						langs: -1,
					},
				},
			},
			{
				pattern: "al",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(l|)",
				phoneticRules: []phoneticRule{
					{
						text:  "l",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "dal",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(dal|[8])",
				phoneticRules: []phoneticRule{
					{
						text:  "dal",
						langs: -1,
					},
					{
						text:  "",
						langs: 8,
					},
				},
			},
			{
				pattern: "da",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(da|a|)",
				phoneticRules: []phoneticRule{
					{
						text:  "da",
						langs: -1,
					},
					{
						text:  "a",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "della",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(dila|)",
				phoneticRules: []phoneticRule{
					{
						text:  "dila",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "dela",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(dila|)",
				phoneticRules: []phoneticRule{
					{
						text:  "dila",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "del",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(dil|)",
				phoneticRules: []phoneticRule{
					{
						text:  "dil",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "des",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(dis|)",
				phoneticRules: []phoneticRule{
					{
						text:  "dis",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern: "de",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(di|i|)",
				phoneticRules: []phoneticRule{
					{
						text:  "di",
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
				pattern: "di",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(di|i|[8])",
				phoneticRules: []phoneticRule{
					{
						text:  "di",
						langs: -1,
					},
					{
						text:  "i",
						langs: -1,
					},
					{
						text:  "",
						langs: 8,
					},
				},
			},
			{
				pattern: "do",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(du|u)",
				phoneticRules: []phoneticRule{
					{
						text:  "du",
						langs: -1,
					},
					{
						text:  "u",
						langs: -1,
					},
				},
			},
			{
				pattern: "du",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(du|u)",
				phoneticRules: []phoneticRule{
					{
						text:  "du",
						langs: -1,
					},
					{
						text:  "u",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oa",
				phonetic: "(va|a)",
				phoneticRules: []phoneticRule{
					{
						text:  "va",
						langs: -1,
					},
					{
						text:  "a",
						langs: -1,
					},
				},
			},
			{
				pattern:  "oe",
				phonetic: "(vi|i)",
				phoneticRules: []phoneticRule{
					{
						text:  "vi",
						langs: -1,
					},
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ae",
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
				pattern: "ha",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(ha|a|)",
				phoneticRules: []phoneticRule{
					{
						text:  "ha",
						langs: -1,
					},
					{
						text:  "a",
						langs: -1,
					},
					{
						text:  "",
						langs: -1,
					},
				},
			},
			{
				pattern:  "h",
				phonetic: "(|h)",
				phoneticRules: []phoneticRule{
					{
						text:  "",
						langs: -1,
					},
					{
						text:  "h",
						langs: -1,
					},
				},
			},
			{
				pattern:  "x",
				phonetic: "h",
				phoneticRules: []phoneticRule{
					{
						text:  "h",
						langs: -1,
					},
				},
			},
			{
				pattern:  "k",
				phonetic: "(h|k)",
				phoneticRules: []phoneticRule{
					{
						text:  "h",
						langs: -1,
					},
					{
						text:  "k",
						langs: -1,
					},
				},
			},
			{
				pattern: "aja",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(Da|ia)",
				phoneticRules: []phoneticRule{
					{
						text:  "Da",
						langs: -1,
					},
					{
						text:  "ia",
						langs: -1,
					},
				},
			},
			{
				pattern: "aje",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(Di|Da|i|ia)",
				phoneticRules: []phoneticRule{
					{
						text:  "Di",
						langs: -1,
					},
					{
						text:  "Da",
						langs: -1,
					},
					{
						text:  "i",
						langs: -1,
					},
					{
						text:  "ia",
						langs: -1,
					},
				},
			},
			{
				pattern: "aji",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(Di|i)",
				phoneticRules: []phoneticRule{
					{
						text:  "Di",
						langs: -1,
					},
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: "ajo",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(Du|Da|iu|ia)",
				phoneticRules: []phoneticRule{
					{
						text:  "Du",
						langs: -1,
					},
					{
						text:  "Da",
						langs: -1,
					},
					{
						text:  "iu",
						langs: -1,
					},
					{
						text:  "ia",
						langs: -1,
					},
				},
			},
			{
				pattern: "aju",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(Du|iu)",
				phoneticRules: []phoneticRule{
					{
						text:  "Du",
						langs: -1,
					},
					{
						text:  "iu",
						langs: -1,
					},
				},
			},
			{
				pattern:  "aj",
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
				pattern:  "ej",
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
				pattern:  "oj",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "uj",
				phonetic: "D",
				phoneticRules: []phoneticRule{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern:  "au",
				phonetic: "u",
				phoneticRules: []phoneticRule{
					{
						text:  "u",
						langs: -1,
					},
				},
			},
			{
				pattern:  "eu",
				phonetic: "(iu|i|u)",
				phoneticRules: []phoneticRule{
					{
						text:  "iu",
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
				pattern: "a",
				leftContext: &ruleMatcher{
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
				pattern: "ja",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "ia",
				phoneticRules: []phoneticRule{
					{
						text:  "ia",
						langs: -1,
					},
				},
			},
			{
				pattern: "je",
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
				pattern: "jo",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(iu|ia)",
				phoneticRules: []phoneticRule{
					{
						text:  "iu",
						langs: -1,
					},
					{
						text:  "ia",
						langs: -1,
					},
				},
			},
			{
				pattern: "ju",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "iu",
				phoneticRules: []phoneticRule{
					{
						text:  "iu",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ja",
				phonetic: "(a|ia)",
				phoneticRules: []phoneticRule{
					{
						text:  "a",
						langs: -1,
					},
					{
						text:  "ia",
						langs: -1,
					},
				},
			},
			{
				pattern:  "je",
				phonetic: "i",
				phoneticRules: []phoneticRule{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ji",
				phonetic: "i",
				phoneticRules: []phoneticRule{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern:  "jo",
				phonetic: "(u|iu)",
				phoneticRules: []phoneticRule{
					{
						text:  "u",
						langs: -1,
					},
					{
						text:  "iu",
						langs: -1,
					},
				},
			},
			{
				pattern:  "ju",
				phonetic: "u",
				phoneticRules: []phoneticRule{
					{
						text:  "u",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(a|u|i)",
				phoneticRules: []phoneticRule{
					{
						text:  "a",
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
				pattern:  "o",
				phonetic: "u",
				phoneticRules: []phoneticRule{
					{
						text:  "u",
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
				pattern: "se",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phonetic: "(z|si)",
				phoneticRules: []phoneticRule{
					{
						text:  "z",
						langs: -1,
					},
					{
						text:  "si",
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
				pattern: "Se",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phonetic: "(z|si)",
				phoneticRules: []phoneticRule{
					{
						text:  "z",
						langs: -1,
					},
					{
						text:  "si",
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
			{
				pattern: "be",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phonetic: "(p|bi)",
				phoneticRules: []phoneticRule{
					{
						text:  "p",
						langs: -1,
					},
					{
						text:  "bi",
						langs: -1,
					},
				},
			},
			{
				pattern: "pe",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				phonetic: "(b|pi)",
				phoneticRules: []phoneticRule{
					{
						text:  "b",
						langs: -1,
					},
					{
						text:  "pi",
						langs: -1,
					},
				},
			},
			{
				pattern: "ve",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phonetic: "(f|vi)",
				phoneticRules: []phoneticRule{
					{
						text:  "f",
						langs: -1,
					},
					{
						text:  "vi",
						langs: -1,
					},
				},
			},
			{
				pattern: "fe",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				phonetic: "(v|fi)",
				phoneticRules: []phoneticRule{
					{
						text:  "v",
						langs: -1,
					},
					{
						text:  "fi",
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
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phonetic: "(k|gi)",
				phoneticRules: []phoneticRule{
					{
						text:  "k",
						langs: -1,
					},
					{
						text:  "gi",
						langs: -1,
					},
				},
			},
			{
				pattern: "ke",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				phonetic: "(g|ki)",
				phoneticRules: []phoneticRule{
					{
						text:  "g",
						langs: -1,
					},
					{
						text:  "ki",
						langs: -1,
					},
				},
			},
			{
				pattern: "de",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phonetic: "(t|di)",
				phoneticRules: []phoneticRule{
					{
						text:  "t",
						langs: -1,
					},
					{
						text:  "di",
						langs: -1,
					},
				},
			},
			{
				pattern: "te",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				phonetic: "(d|ti)",
				phoneticRules: []phoneticRule{
					{
						text:  "d",
						langs: -1,
					},
					{
						text:  "ti",
						langs: -1,
					},
				},
			},
			{
				pattern: "ze",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phonetic: "(s|zi)",
				phoneticRules: []phoneticRule{
					{
						text:  "s",
						langs: -1,
					},
					{
						text:  "zi",
						langs: -1,
					},
				},
			},
			{
				pattern:  "e",
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
			{
				pattern: "p",
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
		},
		second: map[int64][]rule{
			int64(sepany):        []rule{},
			int64(sepfrench):     []rule{},
			int64(sephebrew):     []rule{},
			int64(sepitalian):    []rule{},
			int64(sepportuguese): []rule{},
			int64(sepspanish):    []rule{},
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
				pattern:  "nm",
				phonetic: "m",
				phoneticRules: []phoneticRule{
					{
						text:  "m",
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
				pattern: "a",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "a",
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
				pattern: "e",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "e",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "i",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "o",
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
				pattern: "u",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "u",
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
		second: map[int64][]rule{
			int64(sepany):        []rule{},
			int64(sepfrench):     []rule{},
			int64(sephebrew):     []rule{},
			int64(sepitalian):    []rule{},
			int64(sepportuguese): []rule{},
			int64(sepspanish):    []rule{},
		},
	},
}

var sepDiscards = []string{
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
	"el",
	"la",
	"le",
	"ibn",
	"ha",
}
