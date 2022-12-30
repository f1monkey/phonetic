// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type sepLang uint64

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

const sepAll = sepany +
	sepfrench +
	sephebrew +
	sepitalian +
	sepportuguese +
	sepspanish

var sepRules = map[sepLang][]rule{
	sepany: []rule{
		{
			pattern:  "ph",
			phonetic: "f",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "gli",
			phonetic: "(gli|l[8])",
		},
		{
			pattern:  "gni",
			phonetic: "(gni|ni[10])",
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
		},
		{
			pattern:  "gh",
			phonetic: "(g|gh)",
		},
		{
			pattern:  "dh",
			phonetic: "(d|dh)",
		},
		{
			pattern:  "bh",
			phonetic: "(b|bh)",
		},
		{
			pattern:  "th",
			phonetic: "(t|th)",
		},
		{
			pattern:  "lh",
			phonetic: "(l[16]|lh)",
		},
		{
			pattern:  "nh",
			phonetic: "(n[16]|nh)",
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
		},
		{
			pattern:  "tx",
			phonetic: "tS",
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
		},
		{
			pattern:  "tj",
			phonetic: "dZ",
		},
		{
			pattern:  "tg",
			phonetic: "(tg|dZ[32])",
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
		},
		{
			pattern:  "guy",
			phonetic: "gi",
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
		},
		{
			pattern:  "ñ",
			phonetic: "(n|nj)",
		},
		{
			pattern:  "ny",
			phonetic: "nj",
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
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern:  "ç",
			phonetic: "s",
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
		},
		{
			pattern:  "ch",
			phonetic: "(S|tS[32]|dZ[32])",
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
		},
		{
			pattern:  "z",
			phonetic: "(z|dz[8]|ts[8]|s[32])",
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
		},
		{
			pattern:  "eau",
			phonetic: "o",
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
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
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
		},
		{
			pattern:  "ão",
			phonetic: "(au|an)",
		},
		{
			pattern:  "ãe",
			phonetic: "(aj|an)",
		},
		{
			pattern:  "ãi",
			phonetic: "(aj|an)",
		},
		{
			pattern:  "õe",
			phonetic: "(oj|on)",
		},
		{
			pattern:  "où",
			phonetic: "u",
		},
		{
			pattern:  "ou",
			phonetic: "(ou|u[2])",
		},
		{
			pattern:  "â",
			phonetic: "a",
		},
		{
			pattern:  "à",
			phonetic: "a",
		},
		{
			pattern:  "á",
			phonetic: "a",
		},
		{
			pattern:  "ã",
			phonetic: "(a|an)",
		},
		{
			pattern:  "é",
			phonetic: "e",
		},
		{
			pattern:  "ê",
			phonetic: "e",
		},
		{
			pattern:  "è",
			phonetic: "e",
		},
		{
			pattern:  "í",
			phonetic: "i",
		},
		{
			pattern:  "î",
			phonetic: "i",
		},
		{
			pattern:  "ô",
			phonetic: "o",
		},
		{
			pattern:  "ó",
			phonetic: "o",
		},
		{
			pattern:  "õ",
			phonetic: "(o|on)",
		},
		{
			pattern:  "ò",
			phonetic: "o",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "ü",
			phonetic: "u",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "(b|v[32])",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "e",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "i",
		},
		{
			pattern:  "j",
			phonetic: "(x[32]|Z)",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "(s|S[16])",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "(v|b[32])",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "(ks|gz|S[48])",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	sepfrench: []rule{
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ph",
			phonetic: "f",
		},
		{
			pattern:  "ç",
			phonetic: "s",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "ch",
			phonetic: "S",
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
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "gn",
			phonetic: "(n|gn)",
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
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:  "q",
			phonetic: "k",
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
		},
		{
			pattern:  "ss",
			phonetic: "s",
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
		},
		{
			pattern:  "j",
			phonetic: "Z",
		},
		{
			pattern:  "w",
			phonetic: "v",
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
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
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
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "eau",
			phonetic: "o",
		},
		{
			pattern:  "ai",
			phonetic: "aj",
		},
		{
			pattern:  "ay",
			phonetic: "aj",
		},
		{
			pattern:  "é",
			phonetic: "e",
		},
		{
			pattern:  "ê",
			phonetic: "e",
		},
		{
			pattern:  "è",
			phonetic: "e",
		},
		{
			pattern:  "à",
			phonetic: "a",
		},
		{
			pattern:  "â",
			phonetic: "a",
		},
		{
			pattern:  "où",
			phonetic: "u",
		},
		{
			pattern:  "ou",
			phonetic: "u",
		},
		{
			pattern:  "oi",
			phonetic: "oj",
		},
		{
			pattern:  "ei",
			phonetic: "ej",
		},
		{
			pattern:  "ey",
			phonetic: "ej",
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
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "e",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "i",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	sephebrew: []rule{
		{
			pattern:  "אי",
			phonetic: "i",
		},
		{
			pattern:  "עי",
			phonetic: "i",
		},
		{
			pattern:  "עו",
			phonetic: "VV",
		},
		{
			pattern:  "או",
			phonetic: "VV",
		},
		{
			pattern:  "ג׳",
			phonetic: "Z",
		},
		{
			pattern:  "ד׳",
			phonetic: "dZ",
		},
		{
			pattern:  "א",
			phonetic: "L",
		},
		{
			pattern:  "ב",
			phonetic: "b",
		},
		{
			pattern:  "ג",
			phonetic: "g",
		},
		{
			pattern:  "ד",
			phonetic: "d",
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
		},
		{
			pattern:  "ה",
			phonetic: "",
		},
		{
			pattern:  "וו",
			phonetic: "V",
		},
		{
			pattern:  "וי",
			phonetic: "WW",
		},
		{
			pattern:  "ו",
			phonetic: "W",
		},
		{
			pattern:  "ז",
			phonetic: "z",
		},
		{
			pattern:  "ח",
			phonetic: "X",
		},
		{
			pattern:  "ט",
			phonetic: "T",
		},
		{
			pattern:  "יי",
			phonetic: "i",
		},
		{
			pattern:  "י",
			phonetic: "i",
		},
		{
			pattern:  "ך",
			phonetic: "X",
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
		},
		{
			pattern:  "כ",
			phonetic: "k",
		},
		{
			pattern:  "ל",
			phonetic: "l",
		},
		{
			pattern:  "ם",
			phonetic: "m",
		},
		{
			pattern:  "מ",
			phonetic: "m",
		},
		{
			pattern:  "ן",
			phonetic: "n",
		},
		{
			pattern:  "נ",
			phonetic: "n",
		},
		{
			pattern:  "ס",
			phonetic: "s",
		},
		{
			pattern:  "ע",
			phonetic: "L",
		},
		{
			pattern:  "ף",
			phonetic: "f",
		},
		{
			pattern:  "פ",
			phonetic: "f",
		},
		{
			pattern:  "ץ",
			phonetic: "C",
		},
		{
			pattern:  "צ",
			phonetic: "C",
		},
		{
			pattern:  "ק",
			phonetic: "K",
		},
		{
			pattern:  "ר",
			phonetic: "r",
		},
		{
			pattern:  "ש",
			phonetic: "s",
		},
		{
			pattern:  "ת",
			phonetic: "T",
		},
	},
	sepitalian: []rule{
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "gli",
			phonetic: "(l|gli)",
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
		},
		{
			pattern:  "gni",
			phonetic: "(ni|gni)",
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
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
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
		},
		{
			pattern:  "�",
			phonetic: "e",
		},
		{
			pattern:  "�",
			phonetic: "e",
		},
		{
			pattern:  "�",
			phonetic: "o",
		},
		{
			pattern:  "�",
			phonetic: "o",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "e",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "i",
		},
		{
			pattern:  "j",
			phonetic: "(Z|dZ|j)",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "(ts|dz)",
		},
	},
	sepportuguese: []rule{
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ch",
			phonetic: "S",
		},
		{
			pattern:  "ss",
			phonetic: "s",
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
		},
		{
			pattern:  "ç",
			phonetic: "s",
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
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o|u)",
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
		},
		{
			pattern:  "lh",
			phonetic: "l",
		},
		{
			pattern:  "nh",
			phonetic: "nj",
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
		},
		{
			pattern:  "ão",
			phonetic: "(au|an|on)",
		},
		{
			pattern:  "ãe",
			phonetic: "(aj|an)",
		},
		{
			pattern:  "ãi",
			phonetic: "(aj|an)",
		},
		{
			pattern:  "õe",
			phonetic: "(oj|on)",
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
		},
		{
			pattern:  "â",
			phonetic: "a",
		},
		{
			pattern:  "à",
			phonetic: "a",
		},
		{
			pattern:  "á",
			phonetic: "a",
		},
		{
			pattern:  "ã",
			phonetic: "(a|an|on)",
		},
		{
			pattern:  "é",
			phonetic: "e",
		},
		{
			pattern:  "ê",
			phonetic: "e",
		},
		{
			pattern:  "í",
			phonetic: "i",
		},
		{
			pattern:  "ô",
			phonetic: "o",
		},
		{
			pattern:  "ó",
			phonetic: "o",
		},
		{
			pattern:  "õ",
			phonetic: "(o|on)",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "ü",
			phonetic: "u",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "(e|i)",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "i",
		},
		{
			pattern:  "j",
			phonetic: "Z",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "(o|u)",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "S",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "(S|ks)",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	sepspanish: []rule{
		{
			pattern:  "ñ",
			phonetic: "(n|nj)",
		},
		{
			pattern:  "ny",
			phonetic: "nj",
		},
		{
			pattern:  "ç",
			phonetic: "s",
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
		},
		{
			pattern:  "tx",
			phonetic: "tS",
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
		},
		{
			pattern:  "tj",
			phonetic: "dZ",
		},
		{
			pattern:  "tg",
			phonetic: "(tg|dZ)",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|dZ)",
		},
		{
			pattern:  "bh",
			phonetic: "b",
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
		},
		{
			pattern:  "j",
			phonetic: "(x|Z)",
		},
		{
			pattern:  "x",
			phonetic: "(ks|gz|S)",
		},
		{
			pattern:  "w",
			phonetic: "v",
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
		},
		{
			pattern:  "v",
			phonetic: "(b|v)",
		},
		{
			pattern:  "b",
			phonetic: "(b|v)",
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
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "z",
			phonetic: "(z|s)",
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
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
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
		},
		{
			pattern:  "y",
			phonetic: "(i|j)",
		},
		{
			pattern:  "ü",
			phonetic: "v",
		},
		{
			pattern:  "á",
			phonetic: "a",
		},
		{
			pattern:  "é",
			phonetic: "e",
		},
		{
			pattern:  "í",
			phonetic: "i",
		},
		{
			pattern:  "ó",
			phonetic: "o",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "à",
			phonetic: "a",
		},
		{
			pattern:  "è",
			phonetic: "e",
		},
		{
			pattern:  "ò",
			phonetic: "o",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "e",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "i",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
	},
}

var sepLangRules = []langRule{
	{
		match: ruleMatcher{
			contains: "eau",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ou",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gni",
			prefix:   "",
			suffix:   "",
		},
		langs:  10,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "tx",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "tj",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gy",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "guy",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "sh",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "lh",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "nh",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ny",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gue",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gui",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gia",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gie",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gio",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "giu",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ñ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "â",
			prefix:   "",
			suffix:   "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "á",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "à",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ã",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ê",
			prefix:   "",
			suffix:   "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "í",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "î",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ô",
			prefix:   "",
			suffix:   "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "õ",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ò",
			prefix:   "",
			suffix:   "",
		},
		langs:  40,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ú",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ù",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ü",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "א",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ב",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ג",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ד",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ה",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ו",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ז",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ח",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ט",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "י",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "כ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ל",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "מ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "נ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ס",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ע",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "פ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "צ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ק",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ר",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ש",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ת",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "a",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "o",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "e",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "i",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "y",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "u",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "kh",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "gua",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "guo",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "ç",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "cha",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "cho",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "chu",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "j",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "dj",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "sce",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "sci",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "ó",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "è",
			prefix:   "",
			suffix:   "",
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
			},
			{
				pattern:  "nm",
				phonetic: "m",
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
			},
			{
				pattern:  "mbr",
				phonetic: "mr",
			},
			{
				pattern:  "mpr",
				phonetic: "mr",
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
			},
			{
				pattern:  "oa",
				phonetic: "(va|a)",
			},
			{
				pattern:  "oe",
				phonetic: "(vi|i)",
			},
			{
				pattern:  "ae",
				phonetic: "(a|i)",
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
			},
			{
				pattern:  "h",
				phonetic: "(|h)",
			},
			{
				pattern:  "x",
				phonetic: "h",
			},
			{
				pattern:  "k",
				phonetic: "(h|k)",
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
			},
			{
				pattern:  "aj",
				phonetic: "(D|i)",
			},
			{
				pattern:  "ej",
				phonetic: "(D|i)",
			},
			{
				pattern:  "oj",
				phonetic: "D",
			},
			{
				pattern:  "uj",
				phonetic: "D",
			},
			{
				pattern:  "au",
				phonetic: "u",
			},
			{
				pattern:  "eu",
				phonetic: "(iu|i|u)",
			},
			{
				pattern:  "ou",
				phonetic: "u",
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
			},
			{
				pattern:  "ja",
				phonetic: "(a|ia)",
			},
			{
				pattern:  "je",
				phonetic: "i",
			},
			{
				pattern:  "ji",
				phonetic: "i",
			},
			{
				pattern:  "jo",
				phonetic: "(u|iu)",
			},
			{
				pattern:  "ju",
				phonetic: "u",
			},
			{
				pattern:  "j",
				phonetic: "i",
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
			},
			{
				pattern:  "o",
				phonetic: "u",
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
			},
			{
				pattern:  "S",
				phonetic: "s",
			},
			{
				pattern:  "dZ",
				phonetic: "z",
			},
			{
				pattern:  "Z",
				phonetic: "z",
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
			},
			{
				pattern:  "e",
				phonetic: "(i|)",
			},
			{
				pattern:  "B",
				phonetic: "b",
			},
			{
				pattern:  "V",
				phonetic: "v",
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
			},
		},
		second: map[uint64][]rule{
			uint64(sepany):        []rule{},
			uint64(sepfrench):     []rule{},
			uint64(sephebrew):     []rule{},
			uint64(sepitalian):    []rule{},
			uint64(sepportuguese): []rule{},
			uint64(sepspanish):    []rule{},
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
			},
			{
				pattern:  "nm",
				phonetic: "m",
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
			},
			{
				pattern:  "h",
				phonetic: "",
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
			},
			{
				pattern:  "B",
				phonetic: "b",
			},
			{
				pattern:  "V",
				phonetic: "v",
			},
		},
		second: map[uint64][]rule{
			uint64(sepany):        []rule{},
			uint64(sepfrench):     []rule{},
			uint64(sephebrew):     []rule{},
			uint64(sepitalian):    []rule{},
			uint64(sepportuguese): []rule{},
			uint64(sepspanish):    []rule{},
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
