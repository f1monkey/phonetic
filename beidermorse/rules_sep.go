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

var sepRules = map[sepLang][]rule{
	sepany: []rule{
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
					langs: -1,
				},
			},
		},
		{
			pattern: "gli",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "lh",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "tx",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tj",
			phoneticRules: []phonetic{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "tg",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "gu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phoneticRules: []phonetic{
				{
					text:  "gv",
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
					langs: -1,
				},
			},
		},
		{
			pattern: "ny",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "s",
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
			pattern: "ç",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "ex",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
				pattern:          regexp.MustCompile("^[aioe]"),
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
				pattern:          regexp.MustCompile("^[aie]"),
			},
			phoneticRules: []phonetic{
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
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aáuiíoóeéê]$"),
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
				pattern:          regexp.MustCompile("^[aeiíou]"),
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
					langs: 2,
				},
			},
		},
		{
			pattern: "ão",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "á",
			phoneticRules: []phonetic{
				{
					text:  "a",
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
				{
					text:  "an",
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
			pattern: "ô",
			phoneticRules: []phonetic{
				{
					text:  "o",
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
			pattern: "õ",
			phoneticRules: []phonetic{
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
			pattern: "ü",
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
			pattern: "b",
			phoneticRules: []phonetic{
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
					text:  "e",
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
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "j",
			phoneticRules: []phonetic{
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
				{
					text:  "S",
					langs: 16,
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
				{
					text:  "b",
					langs: 32,
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
	sepfrench: []rule{
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
			pattern: "ss",
			phoneticRules: []phonetic{
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
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: "ey",
			phoneticRules: []phonetic{
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
					text:  "e",
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
					text:  "i",
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
	sephebrew: []rule{
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
					text:  "T",
					langs: -1,
				},
			},
		},
	},
	sepitalian: []rule{
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
			pattern: "gli",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "sc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeou]$"),
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
			pattern: "qu",
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
			pattern: "�",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "�",
			phoneticRules: []phonetic{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: "�",
			phoneticRules: []phonetic{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: "�",
			phoneticRules: []phonetic{
				{
					text:  "o",
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
					text:  "e",
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
					text:  "i",
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
				{
					text:  "dz",
					langs: -1,
				},
			},
		},
	},
	sepportuguese: []rule{
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
					text:  "S",
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
			pattern: "sc",
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
			pattern: "sç",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []phonetic{
				{
					text:  "s",
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
			pattern: "s",
			leftContext: &ruleMatcher{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "z",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "kv",
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
			phoneticRules: []phonetic{
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: "lh",
			phoneticRules: []phonetic{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: "nh",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
				pattern:          regexp.MustCompile("^[aeiíou]"),
			},
			phoneticRules: []phonetic{
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
			pattern: "m",
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
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: "ão",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			pattern: "â",
			phoneticRules: []phonetic{
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
			pattern: "á",
			phoneticRules: []phonetic{
				{
					text:  "a",
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
			pattern: "í",
			phoneticRules: []phonetic{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: "ô",
			phoneticRules: []phonetic{
				{
					text:  "o",
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
			pattern: "õ",
			phoneticRules: []phonetic{
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
					text:  "u",
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
					text:  "i",
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
				{
					text:  "u",
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
	sepspanish: []rule{
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
			pattern: "ny",
			phoneticRules: []phonetic{
				{
					text:  "nj",
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
			pattern: "ig",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: "tx",
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: "tj",
			phoneticRules: []phonetic{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: "tg",
			phoneticRules: []phonetic{
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
			pattern: "bh",
			phoneticRules: []phonetic{
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
				{
					text:  "Z",
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
			phoneticRules: []phonetic{
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
				{
					text:  "dZ",
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
			pattern: "à",
			phoneticRules: []phonetic{
				{
					text:  "a",
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
			pattern: "ò",
			phoneticRules: []phonetic{
				{
					text:  "o",
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
					text:  "e",
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
					text:  "i",
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
					pattern:          regexp.MustCompile("^[vgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgZz]"),
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
				pattern: "nm",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				pattern: "e",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "e",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "i",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "o",
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
				pattern: "u",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "u",
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
				pattern: "mbr",
				phoneticRules: []phonetic{
					{
						text:  "mr",
						langs: -1,
					},
				},
			},
			{
				pattern: "mpr",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				pattern: "ha",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "h",
						langs: -1,
					},
				},
			},
			{
				pattern: "k",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				pattern: "ej",
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
				pattern: "oj",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "uj",
				phoneticRules: []phonetic{
					{
						text:  "D",
						langs: -1,
					},
				},
			},
			{
				pattern: "au",
				phoneticRules: []phonetic{
					{
						text:  "u",
						langs: -1,
					},
				},
			},
			{
				pattern: "eu",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "iu",
						langs: -1,
					},
				},
			},
			{
				pattern: "ja",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: "ji",
				phoneticRules: []phonetic{
					{
						text:  "i",
						langs: -1,
					},
				},
			},
			{
				pattern: "jo",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "u",
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
				pattern: "i",
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				pattern: "se",
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
			{
				pattern: "be",
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				pattern: "B",
				phoneticRules: []phonetic{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
			{
				pattern: "V",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
		},
		second: map[languageID][]rule{
			languageID(sepany):        []rule{},
			languageID(sepfrench):     []rule{},
			languageID(sephebrew):     []rule{},
			languageID(sepitalian):    []rule{},
			languageID(sepportuguese): []rule{},
			languageID(sepspanish):    []rule{},
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
					pattern:          regexp.MustCompile("^[vgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgdZz]"),
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
					pattern:          regexp.MustCompile("^[vbdZz]"),
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
					pattern:          regexp.MustCompile("^[vbgZz]"),
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
				pattern: "nm",
				phoneticRules: []phonetic{
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
				phoneticRules: []phonetic{
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
				pattern: "e",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "e",
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
				pattern: "i",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "i",
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
				pattern: "o",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "o",
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
				pattern: "u",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "u",
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
				pattern: "h",
				phoneticRules: []phonetic{
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
				pattern: "B",
				phoneticRules: []phonetic{
					{
						text:  "b",
						langs: -1,
					},
				},
			},
			{
				pattern: "V",
				phoneticRules: []phonetic{
					{
						text:  "v",
						langs: -1,
					},
				},
			},
		},
		second: map[languageID][]rule{
			languageID(sepany):        []rule{},
			languageID(sepfrench):     []rule{},
			languageID(sephebrew):     []rule{},
			languageID(sepitalian):    []rule{},
			languageID(sepportuguese): []rule{},
			languageID(sepspanish):    []rule{},
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
