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
			pattern:      "gn",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "(n[10]|nj[10]|gn)",
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
			pattern:     "ig",
			leftContext: regexp.MustCompile("[aeiou]$"),
			phonetic:    "(ig|tS[32])",
		},
		{
			pattern:     "ix",
			leftContext: regexp.MustCompile("[aeiou]$"),
			phonetic:    "S",
		},
		{
			pattern:  "tx",
			phonetic: "tS",
		},
		{
			pattern:      "tj",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "tS",
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
			pattern:      "gi",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "dZ",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^y"),
			phonetic:     "Z",
		},
		{
			pattern:      "gg",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(gZ[18]|dZ[40]|x[32])",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(Z[18]|dZ[40]|x[32])",
		},
		{
			pattern:  "guy",
			phonetic: "gi",
		},
		{
			pattern:      "gue",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k[2]|ge)",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(g|gv)",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[ao]"),
			phonetic:     "gv",
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
			pattern:      "sc",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(s|S[8])",
		},
		{
			pattern:      "sç",
			rightContext: regexp.MustCompile("^[aeiou]"),
			phonetic:     "s",
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
			pattern:      "ch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(k[8]|S[18]|tS[32]|dZ[32])",
		},
		{
			pattern:  "ch",
			phonetic: "(S|tS[32]|dZ[32])",
		},
		{
			pattern:      "ci",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "(tS[8]|si)",
		},
		{
			pattern:      "cc",
			rightContext: regexp.MustCompile("^[eiyéèê]"),
			phonetic:     "(tS[8]|ks[50])",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[eiyéèê]"),
			phonetic:     "(tS[8]|s[50])",
		},
		{
			pattern:     "s",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "s",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aáuiíoóeéêy]$"),
			rightContext: regexp.MustCompile("^[aáuiíoóeéêy]"),
			phonetic:     "(s[32]|z[26])",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[dglmnrv]"),
			phonetic:     "(z|s|Z[16])",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(s|ts[8]|S[16])",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^[bdgv]"),
			phonetic:     "(z|dz[8]|Z[16])",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^[ptckf]"),
			phonetic:     "(s|ts[8]|S[16])",
		},
		{
			pattern:  "z",
			phonetic: "(z|dz[8]|ts[8]|s[32])",
		},
		{
			pattern:      "que",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k[2]|ke)",
		},
		{
			pattern:      "qu",
			rightContext: regexp.MustCompile("^[eiu]"),
			phonetic:     "k",
		},
		{
			pattern:      "qu",
			rightContext: regexp.MustCompile("^[ao]"),
			phonetic:     "(kv|k)",
		},
		{
			pattern:      "ex",
			rightContext: regexp.MustCompile("^[aáuiíoóeéêy]"),
			phonetic:     "(ez[16]|eS[16]|eks|egz)",
		},
		{
			pattern:      "ex",
			rightContext: regexp.MustCompile("^[cs]"),
			phonetic:     "(e[16]|ek)",
		},
		{
			pattern:      "m",
			rightContext: regexp.MustCompile("^[cdglnrst]"),
			phonetic:     "(m|n[16])",
		},
		{
			pattern:      "m",
			rightContext: regexp.MustCompile("^[bfpv]"),
			phonetic:     "(m|n[48])",
		},
		{
			pattern:      "m",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(m|n[16])",
		},
		{
			pattern:     "b",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(b|V[32])",
		},
		{
			pattern:     "v",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(v|B[32])",
		},
		{
			pattern:  "eau",
			phonetic: "o",
		},
		{
			pattern:      "ouh",
			rightContext: regexp.MustCompile("^[aioe]"),
			phonetic:     "(v[2]|uh)",
		},
		{
			pattern:      "uh",
			rightContext: regexp.MustCompile("^[aioe]"),
			phonetic:     "(v|uh)",
		},
		{
			pattern:      "ou",
			rightContext: regexp.MustCompile("^[aioe]"),
			phonetic:     "v",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
		},
		{
			pattern:      "u",
			rightContext: regexp.MustCompile("^[aie]"),
			phonetic:     "v",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aáuoóeéê]$"),
			phonetic:    "j",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aáuiíoóeéê]$"),
			phonetic:    "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[aeiíou]"),
			phonetic:     "j",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(e|[2])",
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
			pattern:      "c",
			rightContext: regexp.MustCompile("^[eiyéèê]"),
			phonetic:     "s",
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
			pattern:      "g",
			rightContext: regexp.MustCompile("^[eiy]"),
			phonetic:     "Z",
		},
		{
			pattern:      "gue",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "k",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[eiy]"),
			phonetic:     "g",
		},
		{
			pattern:      "que",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "k",
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
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aeiouyéèê]$"),
			rightContext: regexp.MustCompile("^[aeiouyéèê]"),
			phonetic:     "z",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("[bdgt]$"),
			phonetic:    "",
		},
		{
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
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
			pattern:      "ouh",
			rightContext: regexp.MustCompile("^[aioe]"),
			phonetic:     "(v|uh)",
		},
		{
			pattern:      "ou",
			rightContext: regexp.MustCompile("^[aeio]"),
			phonetic:     "v",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
		},
		{
			pattern:      "u",
			rightContext: regexp.MustCompile("^[aeio]"),
			phonetic:     "v",
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
			pattern:     "y",
			leftContext: regexp.MustCompile("[ou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(e|)",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[aoeu]"),
			phonetic:     "j",
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
			pattern:     "ה",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "1",
		},
		{
			pattern:      "ה",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "1",
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
			pattern:     "כ",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "K",
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
			pattern:      "gn",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "(n|nj|gn)",
		},
		{
			pattern:  "gni",
			phonetic: "(ni|gni)",
		},
		{
			pattern:      "gi",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "dZ",
		},
		{
			pattern:      "gg",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "dZ",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "dZ",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("[bdgt]$"),
			phonetic:    "g",
		},
		{
			pattern:      "ci",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "tS",
		},
		{
			pattern:      "ch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "k",
		},
		{
			pattern:      "sc",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "S",
		},
		{
			pattern:      "cc",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "tS",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "tS",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aeiou]$"),
			rightContext: regexp.MustCompile("^[aeiou]"),
			phonetic:     "z",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
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
			pattern:      "u",
			rightContext: regexp.MustCompile("^[aei]"),
			phonetic:     "v",
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
			pattern:      "sc",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "s",
		},
		{
			pattern:      "sç",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "s",
		},
		{
			pattern:  "ç",
			phonetic: "s",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "s",
		},
		{
			pattern:     "s",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "s",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aáuiíoóeéêy]$"),
			rightContext: regexp.MustCompile("^[aáuiíoóeéêy]"),
			phonetic:     "z",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[dglmnrv]"),
			phonetic:     "(Z|S)",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(Z|s|S)",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^[bdgv]"),
			phonetic:     "(Z|z)",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^[ptckf]"),
			phonetic:     "(s|S|z)",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[eiu]"),
			phonetic:     "g",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[ao]"),
			phonetic:     "gv",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "Z",
		},
		{
			pattern:      "qu",
			rightContext: regexp.MustCompile("^[eiu]"),
			phonetic:     "k",
		},
		{
			pattern:      "qu",
			rightContext: regexp.MustCompile("^[ao]"),
			phonetic:     "kv",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o|u)",
		},
		{
			pattern:      "u",
			rightContext: regexp.MustCompile("^[aei]"),
			phonetic:     "v",
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
			pattern:     "h",
			leftContext: regexp.MustCompile("[bdgt]$"),
			phonetic:    "",
		},
		{
			pattern:      "ex",
			rightContext: regexp.MustCompile("^[aáuiíoóeéêy]"),
			phonetic:     "(ez|eS|eks)",
		},
		{
			pattern:      "ex",
			rightContext: regexp.MustCompile("^[cs]"),
			phonetic:     "e",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aáuiíoóeéê]$"),
			phonetic:    "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[aeiíou]"),
			phonetic:     "j",
		},
		{
			pattern:      "m",
			rightContext: regexp.MustCompile("^[bcdfglnprstv]"),
			phonetic:     "(m|n)",
		},
		{
			pattern:      "m",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(m|n)",
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
			pattern:     "i",
			leftContext: regexp.MustCompile("[aáuoóeéê]$"),
			phonetic:    "j",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
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
			pattern:     "ig",
			leftContext: regexp.MustCompile("[aeiou]$"),
			phonetic:    "(tS|ig)",
		},
		{
			pattern:     "ix",
			leftContext: regexp.MustCompile("[aeiou]$"),
			phonetic:    "S",
		},
		{
			pattern:  "tx",
			phonetic: "tS",
		},
		{
			pattern:      "tj",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "tS",
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
			pattern:     "h",
			leftContext: regexp.MustCompile("[dgt]$"),
			phonetic:    "",
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
			pattern:     "v",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(B|v)",
		},
		{
			pattern:     "b",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(b|V)",
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
			pattern:      "m",
			rightContext: regexp.MustCompile("^[bpvf]"),
			phonetic:     "(m|n)",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "s",
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
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(g|gv)",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(x|g|dZ)",
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
			pattern:      "u",
			rightContext: regexp.MustCompile("^[aei]"),
			phonetic:     "v",
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
		match: ruleMatch{
			contains: "eau",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ou",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gni",
			prefix:   "",
			suffix:   "",
		},
		langs:  10,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "tx",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "tj",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gy",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "guy",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "sh",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "lh",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "nh",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ny",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gue",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gui",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gia",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gie",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gio",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "giu",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ñ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "â",
			prefix:   "",
			suffix:   "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "á",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "à",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ã",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ê",
			prefix:   "",
			suffix:   "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "í",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "î",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ô",
			prefix:   "",
			suffix:   "",
		},
		langs:  18,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "õ",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ò",
			prefix:   "",
			suffix:   "",
		},
		langs:  40,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ú",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ù",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ü",
			prefix:   "",
			suffix:   "",
		},
		langs:  48,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "א",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ב",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ג",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ד",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ה",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ו",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ז",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ח",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ט",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "י",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "כ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ל",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "מ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "נ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ס",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ע",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "פ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "צ",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ק",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ר",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ש",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ת",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "a",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "o",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "e",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "i",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "y",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "u",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "kh",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "gua",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "guo",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "ç",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "cha",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "cho",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "chu",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "j",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "dj",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "sce",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "sci",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "ó",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: false,
	},
	{
		match: ruleMatch{
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
				pattern:      "h",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^[fktSs]"),
				phonetic:     "p",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^p"),
				phonetic:     "",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "p",
			},
			{
				pattern:      "p",
				rightContext: regexp.MustCompile("^[vgdZz]"),
				phonetic:     "b",
			},
			{
				pattern:      "p",
				rightContext: regexp.MustCompile("^b"),
				phonetic:     "",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^[pktSs]"),
				phonetic:     "f",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^f"),
				phonetic:     "",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "f",
			},
			{
				pattern:      "f",
				rightContext: regexp.MustCompile("^[vbgdZz]"),
				phonetic:     "v",
			},
			{
				pattern:      "f",
				rightContext: regexp.MustCompile("^v"),
				phonetic:     "",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^[pftSs]"),
				phonetic:     "k",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^k"),
				phonetic:     "",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "k",
			},
			{
				pattern:      "k",
				rightContext: regexp.MustCompile("^[vbdZz]"),
				phonetic:     "g",
			},
			{
				pattern:      "k",
				rightContext: regexp.MustCompile("^g"),
				phonetic:     "",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^[pfkSs]"),
				phonetic:     "t",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^t"),
				phonetic:     "",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "t",
			},
			{
				pattern:      "t",
				rightContext: regexp.MustCompile("^[vbgZz]"),
				phonetic:     "d",
			},
			{
				pattern:      "t",
				rightContext: regexp.MustCompile("^d"),
				phonetic:     "",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^dZ"),
				phonetic:     "",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^tS"),
				phonetic:     "",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^[pfkSt]"),
				phonetic:     "s",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:      "Z",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:      "S",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:  "nm",
				phonetic: "m",
			},
			{
				pattern:     "ji",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "i",
			},
			{
				pattern:      "a",
				rightContext: regexp.MustCompile("^a"),
				phonetic:     "",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^b"),
				phonetic:     "",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^d"),
				phonetic:     "",
			},
			{
				pattern:      "e",
				rightContext: regexp.MustCompile("^e"),
				phonetic:     "",
			},
			{
				pattern:      "f",
				rightContext: regexp.MustCompile("^f"),
				phonetic:     "",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^g"),
				phonetic:     "",
			},
			{
				pattern:      "i",
				rightContext: regexp.MustCompile("^i"),
				phonetic:     "",
			},
			{
				pattern:      "k",
				rightContext: regexp.MustCompile("^k"),
				phonetic:     "",
			},
			{
				pattern:      "l",
				rightContext: regexp.MustCompile("^l"),
				phonetic:     "",
			},
			{
				pattern:      "m",
				rightContext: regexp.MustCompile("^m"),
				phonetic:     "",
			},
			{
				pattern:      "n",
				rightContext: regexp.MustCompile("^n"),
				phonetic:     "",
			},
			{
				pattern:      "o",
				rightContext: regexp.MustCompile("^o"),
				phonetic:     "",
			},
			{
				pattern:      "p",
				rightContext: regexp.MustCompile("^p"),
				phonetic:     "",
			},
			{
				pattern:      "r",
				rightContext: regexp.MustCompile("^r"),
				phonetic:     "",
			},
			{
				pattern:      "t",
				rightContext: regexp.MustCompile("^t"),
				phonetic:     "",
			},
			{
				pattern:      "u",
				rightContext: regexp.MustCompile("^u"),
				phonetic:     "",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^v"),
				phonetic:     "",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^z"),
				phonetic:     "",
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
				pattern:     "bens",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(binz|s)",
			},
			{
				pattern:     "benS",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(binz|s)",
			},
			{
				pattern:     "ben",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bin|)",
			},
			{
				pattern:     "bar",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bar|)",
			},
			{
				pattern:     "abens",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(binz|s)",
			},
			{
				pattern:     "abenS",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(binz|s)",
			},
			{
				pattern:     "aben",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bin|bun|)",
			},
			{
				pattern:     "abe",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bi|bu|)",
			},
			{
				pattern:     "abi",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bi|bu|)",
			},
			{
				pattern:     "abou",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bu|[2])",
			},
			{
				pattern:     "abu",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bu|)",
			},
			{
				pattern:     "bou",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bu|[2])",
			},
			{
				pattern:     "bu",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(bu|)",
			},
			{
				pattern:     "ibn",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(ibn|)",
			},
			{
				pattern:     "els",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(ilz|lz|s)",
			},
			{
				pattern:     "elS",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(ilz|lz|s)",
			},
			{
				pattern:     "el",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(il|l|)",
			},
			{
				pattern:     "als",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(lz|s)",
			},
			{
				pattern:     "alS",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(lz|s)",
			},
			{
				pattern:     "al",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(l|)",
			},
			{
				pattern:     "dal",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(dal|[8])",
			},
			{
				pattern:     "da",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(da|a|)",
			},
			{
				pattern:     "della",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(dila|)",
			},
			{
				pattern:     "dela",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(dila|)",
			},
			{
				pattern:     "del",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(dil|)",
			},
			{
				pattern:     "des",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(dis|)",
			},
			{
				pattern:     "de",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(di|i|)",
			},
			{
				pattern:     "di",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(di|i|[8])",
			},
			{
				pattern:     "do",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(du|u)",
			},
			{
				pattern:     "du",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(du|u)",
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
				pattern:      "n",
				rightContext: regexp.MustCompile("^[bp]"),
				phonetic:     "m",
			},
			{
				pattern:     "ha",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(ha|a|)",
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
				pattern:     "aja",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(Da|ia)",
			},
			{
				pattern:     "aje",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(Di|Da|i|ia)",
			},
			{
				pattern:     "aji",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(Di|i)",
			},
			{
				pattern:     "ajo",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(Du|Da|iu|ia)",
			},
			{
				pattern:     "aju",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(Du|iu)",
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
				pattern:     "a",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "",
			},
			{
				pattern:     "ja",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "ia",
			},
			{
				pattern:     "je",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "i",
			},
			{
				pattern:     "jo",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(iu|ia)",
			},
			{
				pattern:     "ju",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "iu",
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
				pattern:      "i",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(i|)",
			},
			{
				pattern:      "o",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(a|u|i)",
			},
			{
				pattern:  "o",
				phonetic: "u",
			},
			{
				pattern:      "a",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(a|i)",
			},
			{
				pattern:      "se",
				rightContext: regexp.MustCompile("^[rmnl]"),
				phonetic:     "(z|si)",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^[rmnl]"),
				phonetic:     "z",
			},
			{
				pattern:      "Se",
				rightContext: regexp.MustCompile("^[rmnl]"),
				phonetic:     "(z|si)",
			},
			{
				pattern:      "S",
				rightContext: regexp.MustCompile("^[rmnl]"),
				phonetic:     "z",
			},
			{
				pattern:     "s",
				leftContext: regexp.MustCompile("[rmnl]$"),
				phonetic:    "z",
			},
			{
				pattern:     "S",
				leftContext: regexp.MustCompile("[rmnl]$"),
				phonetic:    "z",
			},
			{
				pattern:      "dS",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "S",
			},
			{
				pattern:      "dZ",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "S",
			},
			{
				pattern:      "Z",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "S",
			},
			{
				pattern:      "S",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(S|s)",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(S|s)",
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
				pattern:      "be",
				rightContext: regexp.MustCompile("^[fktSs]"),
				phonetic:     "(p|bi)",
			},
			{
				pattern:      "pe",
				rightContext: regexp.MustCompile("^[vgdZz]"),
				phonetic:     "(b|pi)",
			},
			{
				pattern:      "ve",
				rightContext: regexp.MustCompile("^[pktSs]"),
				phonetic:     "(f|vi)",
			},
			{
				pattern:      "fe",
				rightContext: regexp.MustCompile("^[vbgdZz]"),
				phonetic:     "(v|fi)",
			},
			{
				pattern:      "ge",
				rightContext: regexp.MustCompile("^[pftSs]"),
				phonetic:     "(k|gi)",
			},
			{
				pattern:      "ke",
				rightContext: regexp.MustCompile("^[vbdZz]"),
				phonetic:     "(g|ki)",
			},
			{
				pattern:      "de",
				rightContext: regexp.MustCompile("^[pfkSs]"),
				phonetic:     "(t|di)",
			},
			{
				pattern:      "te",
				rightContext: regexp.MustCompile("^[vbgZz]"),
				phonetic:     "(d|ti)",
			},
			{
				pattern:      "ze",
				rightContext: regexp.MustCompile("^[pfkSt]"),
				phonetic:     "(s|zi)",
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
				pattern:     "p",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "b",
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
				pattern:      "h",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^[fktSs]"),
				phonetic:     "p",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^p"),
				phonetic:     "",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "p",
			},
			{
				pattern:      "p",
				rightContext: regexp.MustCompile("^[vgdZz]"),
				phonetic:     "b",
			},
			{
				pattern:      "p",
				rightContext: regexp.MustCompile("^b"),
				phonetic:     "",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^[pktSs]"),
				phonetic:     "f",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^f"),
				phonetic:     "",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "f",
			},
			{
				pattern:      "f",
				rightContext: regexp.MustCompile("^[vbgdZz]"),
				phonetic:     "v",
			},
			{
				pattern:      "f",
				rightContext: regexp.MustCompile("^v"),
				phonetic:     "",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^[pftSs]"),
				phonetic:     "k",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^k"),
				phonetic:     "",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "k",
			},
			{
				pattern:      "k",
				rightContext: regexp.MustCompile("^[vbdZz]"),
				phonetic:     "g",
			},
			{
				pattern:      "k",
				rightContext: regexp.MustCompile("^g"),
				phonetic:     "",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^[pfkSs]"),
				phonetic:     "t",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^t"),
				phonetic:     "",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "t",
			},
			{
				pattern:      "t",
				rightContext: regexp.MustCompile("^[vbgZz]"),
				phonetic:     "d",
			},
			{
				pattern:      "t",
				rightContext: regexp.MustCompile("^d"),
				phonetic:     "",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^dZ"),
				phonetic:     "",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^tS"),
				phonetic:     "",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^[pfkSt]"),
				phonetic:     "s",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:      "Z",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:      "S",
				rightContext: regexp.MustCompile("^[sSzZ]"),
				phonetic:     "",
			},
			{
				pattern:  "nm",
				phonetic: "m",
			},
			{
				pattern:     "ji",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "i",
			},
			{
				pattern:      "a",
				rightContext: regexp.MustCompile("^a"),
				phonetic:     "",
			},
			{
				pattern:      "b",
				rightContext: regexp.MustCompile("^b"),
				phonetic:     "",
			},
			{
				pattern:      "d",
				rightContext: regexp.MustCompile("^d"),
				phonetic:     "",
			},
			{
				pattern:      "e",
				rightContext: regexp.MustCompile("^e"),
				phonetic:     "",
			},
			{
				pattern:      "f",
				rightContext: regexp.MustCompile("^f"),
				phonetic:     "",
			},
			{
				pattern:      "g",
				rightContext: regexp.MustCompile("^g"),
				phonetic:     "",
			},
			{
				pattern:      "i",
				rightContext: regexp.MustCompile("^i"),
				phonetic:     "",
			},
			{
				pattern:      "k",
				rightContext: regexp.MustCompile("^k"),
				phonetic:     "",
			},
			{
				pattern:      "l",
				rightContext: regexp.MustCompile("^l"),
				phonetic:     "",
			},
			{
				pattern:      "m",
				rightContext: regexp.MustCompile("^m"),
				phonetic:     "",
			},
			{
				pattern:      "n",
				rightContext: regexp.MustCompile("^n"),
				phonetic:     "",
			},
			{
				pattern:      "o",
				rightContext: regexp.MustCompile("^o"),
				phonetic:     "",
			},
			{
				pattern:      "p",
				rightContext: regexp.MustCompile("^p"),
				phonetic:     "",
			},
			{
				pattern:      "r",
				rightContext: regexp.MustCompile("^r"),
				phonetic:     "",
			},
			{
				pattern:      "t",
				rightContext: regexp.MustCompile("^t"),
				phonetic:     "",
			},
			{
				pattern:      "u",
				rightContext: regexp.MustCompile("^u"),
				phonetic:     "",
			},
			{
				pattern:      "v",
				rightContext: regexp.MustCompile("^v"),
				phonetic:     "",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^z"),
				phonetic:     "",
			},
			{
				pattern:  "h",
				phonetic: "",
			},
			{
				pattern:      "s",
				leftContext:  regexp.MustCompile("[^t]$"),
				rightContext: regexp.MustCompile("^[bgZd]"),
				phonetic:     "z",
			},
			{
				pattern:      "Z",
				rightContext: regexp.MustCompile("^[pfkst]"),
				phonetic:     "S",
			},
			{
				pattern:      "Z",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "S",
			},
			{
				pattern:      "S",
				rightContext: regexp.MustCompile("^[bgzd]"),
				phonetic:     "Z",
			},
			{
				pattern:      "z",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "s",
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
