// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type sepLang languageID

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

var sepRules = map[sepLang]rules{
	sepany: rules{
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
					langs: -1,
				},
			},
		},
		{
			pattern: "gli",
			phoneticRules: []token{
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
			pattern: "gni",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "gh",
			phoneticRules: []token{
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
			pattern: "dh",
			phoneticRules: []token{
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
			pattern: "bh",
			phoneticRules: []token{
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
			pattern: "th",
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
			pattern: "lh",
			phoneticRules: []token{
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
			pattern: "nh",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "tx",
			phoneticRules: []token{
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
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tj",
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "tg",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "guy",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "gu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phoneticRules: []token{
				{
					text:  "gv",
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
					langs: -1,
				},
			},
		},
		{
			pattern: "ny",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
				{
					text:  "s",
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
			pattern: "ç",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "ch",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "eau",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
				pattern:          regexp.MustCompile("^[aioe]"),
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
				pattern:          regexp.MustCompile("^[aie]"),
			},
			phoneticRules: []token{
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
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aáuiíoóeéê]$"),
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
				pattern:          regexp.MustCompile("^[aeiíou]"),
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
					langs: 2,
				},
			},
		},
		{
			pattern: "ão",
			phoneticRules: []token{
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
			pattern: "ãe",
			phoneticRules: []token{
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
			pattern: "ãi",
			phoneticRules: []token{
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
			pattern: "õe",
			phoneticRules: []token{
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
			pattern: "â",
			phoneticRules: []token{
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
			pattern: "á",
			phoneticRules: []token{
				{
					text:  "a",
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
				{
					text:  "an",
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
			pattern: "ô",
			phoneticRules: []token{
				{
					text:  "o",
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
			pattern: "õ",
			phoneticRules: []token{
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
			pattern: "ò",
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
			pattern: "ü",
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
			pattern: "b",
			phoneticRules: []token{
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
					text:  "e",
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
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []token{
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
				{
					text:  "S",
					langs: 16,
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
				{
					text:  "b",
					langs: 32,
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
	sepfrench: rules{
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
			pattern: "ss",
			phoneticRules: []token{
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
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []token{
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
					text:  "e",
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
					text:  "i",
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
	sephebrew: rules{
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
					text:  "T",
					langs: -1,
				},
			},
		},
	},
	sepitalian: rules{
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
			pattern: "gli",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "gni",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "sc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeou]$"),
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
			pattern: "qu",
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
			pattern: "�",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "�",
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "�",
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "�",
			phoneticRules: []token{
				{
					text:  "o",
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
					text:  "e",
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
					text:  "i",
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
				{
					text:  "dz",
					langs: -1,
				},
			},
		},
	},
	sepportuguese: rules{
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
					text:  "S",
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
			pattern: "sc",
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
			pattern: "sç",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  "s",
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
			pattern: "s",
			leftContext: &ruleMatcher{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: "z",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
				{
					text:  "kv",
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
			phoneticRules: []token{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "lh",
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "nh",
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			phoneticRules: []token{
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
				pattern:          regexp.MustCompile("^[aeiíou]"),
			},
			phoneticRules: []token{
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
			pattern: "m",
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
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ão",
			phoneticRules: []token{
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
			pattern: "ãe",
			phoneticRules: []token{
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
			pattern: "ãi",
			phoneticRules: []token{
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
			pattern: "õe",
			phoneticRules: []token{
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
			pattern: "â",
			phoneticRules: []token{
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
			pattern: "á",
			phoneticRules: []token{
				{
					text:  "a",
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
			pattern: "í",
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ô",
			phoneticRules: []token{
				{
					text:  "o",
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
			pattern: "õ",
			phoneticRules: []token{
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
					text:  "u",
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
					text:  "i",
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
				{
					text:  "u",
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
	sepspanish: rules{
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
			pattern: "ny",
			phoneticRules: []token{
				{
					text:  "nj",
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
			pattern: "ig",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phoneticRules: []token{
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
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "tx",
			phoneticRules: []token{
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
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tj",
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "tg",
			phoneticRules: []token{
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
			pattern: "bh",
			phoneticRules: []token{
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
				{
					text:  "Z",
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
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
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
			phoneticRules: []token{
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
				{
					text:  "dZ",
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
			pattern: "à",
			phoneticRules: []token{
				{
					text:  "a",
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
			pattern: "ò",
			phoneticRules: []token{
				{
					text:  "o",
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
					text:  "e",
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
					text:  "i",
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
					pattern:          regexp.MustCompile("^[vgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgZz]"),
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
				pattern: "nm",
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				pattern: "e",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "e",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "i",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "o",
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
				pattern: "u",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "u",
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
				pattern: "mbr",
				phoneticRules: []token{
					{
						text:  "mr",
						langs: -1,
					},
				},
			},
			{
				pattern: "mpr",
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				pattern: "oa",
				phoneticRules: []token{
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
				pattern: "oe",
				phoneticRules: []token{
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
				pattern: "ae",
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
				pattern: "ha",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
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
				pattern: "h",
				phoneticRules: []token{
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
				pattern: "x",
				phoneticRules: []token{
					{
						text:  "h",
						langs: -1,
					},
				},
			},
			{
				pattern: "k",
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				pattern: "aj",
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
				pattern: "ej",
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
				pattern: "oj",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uj",
				phoneticRules: []token{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "au",
				phoneticRules: []token{
					{
						text:  "u",
						langs: -1,
					},
				},
			},
			{
				pattern: "eu",
				phoneticRules: []token{
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
				pattern: "ou",
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
					{
						text:  "iu",
						langs: -1,
					},
				},
			},
			{
				pattern: "ja",
				phoneticRules: []token{
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
				pattern: "je",
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: "ji",
				phoneticRules: []token{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: "jo",
				phoneticRules: []token{
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
				pattern: "ju",
				phoneticRules: []token{
					{
						text:  "u",
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
				pattern: "i",
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
				phoneticRules: []token{
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
				pattern: "o",
				phoneticRules: []token{
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
				pattern: "se",
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
			{
				pattern: "be",
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				pattern: "e",
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
				pattern: "B",
				phoneticRules: []token{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
			{
				pattern: "V",
				phoneticRules: []token{
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
				phoneticRules: []token{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
		},
		second: map[languageID]rules{
			languageID(sepany):        rules{},
			languageID(sepfrench):     rules{},
			languageID(sephebrew):     rules{},
			languageID(sepitalian):    rules{},
			languageID(sepportuguese): rules{},
			languageID(sepspanish):    rules{},
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
					pattern:          regexp.MustCompile("^[vgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgZz]"),
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
				pattern: "nm",
				phoneticRules: []token{
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
				phoneticRules: []token{
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
				pattern: "e",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "e",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "i",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "o",
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
				pattern: "u",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "u",
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
				pattern: "h",
				phoneticRules: []token{
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
				pattern: "B",
				phoneticRules: []token{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
			{
				pattern: "V",
				phoneticRules: []token{
					{
						text:  "v",
						langs: -1,
					},
				},
			},
		},
		second: map[languageID]rules{
			languageID(sepany):        rules{},
			languageID(sepfrench):     rules{},
			languageID(sephebrew):     rules{},
			languageID(sepitalian):    rules{},
			languageID(sepportuguese): rules{},
			languageID(sepspanish):    rules{},
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
