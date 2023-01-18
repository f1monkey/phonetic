// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type genLang languageID

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

var genRules = map[genLang]rules{
	genany: rules{
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
					langs: 131072,
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
					langs: 131072,
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
			pattern: []rune("lova"),
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
			pattern: []rune("kova"),
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
			pattern: []rune("ova"),
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
					langs: 131072,
				},
				{
					text:  "",
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ová"),
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
					text:  "",
					langs: 8,
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
					langs: 131072,
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
					langs: 131072,
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
					langs: 131072,
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
					langs: 16384,
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
					langs: 16384,
				},
				{
					text:  "el",
					langs: 16384,
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
					langs: 16384,
				},
				{
					text:  "ek",
					langs: 16384,
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
					langs: 16384,
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
			pattern: []rune("á"),
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
					langs: 8,
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
					langs: 16392,
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
			pattern: []rune("m"),
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
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
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
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("lio"),
			phoneticRules: []token{
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
			pattern: []rune("lyo"),
			phoneticRules: []token{
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
			pattern: []rune("lt"),
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
			phoneticRules: []token{
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
					langs: 128,
				},
				{
					text:  "b",
					langs: 262144,
				},
			},
		},
		{
			pattern: []rune("ex"),
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
			pattern: []rune("ex"),
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
					langs: 32768,
				},
				{
					text:  "ek",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
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
			phoneticRules: []token{
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
			pattern: []rune("ck"),
			phoneticRules: []token{
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
			pattern: []rune("cz"),
			phoneticRules: []token{
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
			pattern: []rune("dh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("bh"),
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
			pattern: []rune("kh"),
			phoneticRules: []token{
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
			pattern: []rune("lh"),
			phoneticRules: []token{
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
			pattern: []rune("nh"),
			phoneticRules: []token{
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
			phoneticRules: []token{
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
			pattern: []rune("sch"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("sch"),
			phoneticRules: []token{
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
					langs: 131104,
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
			pattern: []rune("zh"),
			phoneticRules: []token{
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
			pattern: []rune("chs"),
			phoneticRules: []token{
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
			pattern: []rune("ch"),
			phoneticRules: []token{
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
					langs: 672,
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
					langs: 70144,
				},
				{
					text:  "gh",
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
					langs: 64,
				},
				{
					text:  "uh",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uh"),
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
			pattern: []rune("cia"),
			phoneticRules: []token{
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
					langs: 16384,
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
					langs: 16384,
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
					langs: 16384,
				},
				{
					text:  "tsen",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
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
			pattern: []rune("cio"),
			phoneticRules: []token{
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
			pattern: []rune("ciu"),
			phoneticRules: []token{
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
			pattern: []rune("sci"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
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
			pattern: []rune("sc"),
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
			pattern: []rune("cy"),
			phoneticRules: []token{
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
			pattern: []rune("sç"),
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
				{
					text:  "stS",
					langs: 524288,
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
			leftContext: &ruleMatcher{
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
					langs: 2048,
				},
			},
		},
		{
			pattern: []rune("sz"),
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
					langs: 2048,
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
			pattern: []rune("ssp"),
			phoneticRules: []token{
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
			pattern: []rune("sp"),
			phoneticRules: []token{
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
			pattern: []rune("sst"),
			phoneticRules: []token{
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
			pattern: []rune("st"),
			phoneticRules: []token{
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
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sj"),
			leftContext: &ruleMatcher{
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
			pattern: []rune("sj"),
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
			pattern: []rune("sj"),
			phoneticRules: []token{
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
			pattern: []rune("sia"),
			phoneticRules: []token{
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
					langs: 16384,
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
					langs: 16384,
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
					langs: 16384,
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
					langs: 16384,
				},
				{
					text:  "sen",
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
			pattern: []rune("sio"),
			phoneticRules: []token{
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
			pattern: []rune("siu"),
			phoneticRules: []token{
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
			pattern: []rune("si"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[äöëaáuiíoóeéêy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("si"),
			phoneticRules: []token{
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
			pattern: []rune("s"),
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
			phoneticRules: []token{
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
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeouäöë]"),
			},
			phoneticRules: []token{
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
			pattern: []rune("s"),
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
			phoneticRules: []token{
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
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phoneticRules: []token{
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
					langs: 64,
				},
				{
					text:  "gve",
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
					langs: 64,
				},
				{
					text:  "gv",
					langs: 294912,
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
			pattern: []rune("guy"),
			phoneticRules: []token{
				{
					text:  "gi",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gli"),
			phoneticRules: []token{
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
			pattern: []rune("gni"),
			phoneticRules: []token{
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
			pattern: []rune("gn"),
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
			pattern: []rune("ggie"),
			phoneticRules: []token{
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
			pattern: []rune("ggi"),
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
					langs: 512,
				},
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
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
			phoneticRules: []token{
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
			pattern: []rune("gge"),
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
			pattern: []rune("ggi"),
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
			pattern: []rune("ggi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
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
			pattern: []rune("gie"),
			phoneticRules: []token{
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
			pattern: []rune("gi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
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
			pattern: []rune("ge"),
			phoneticRules: []token{
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
			pattern: []rune("gi"),
			phoneticRules: []token{
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
					langs: 2048,
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
					langs: 2048,
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
				pattern:          regexp.MustCompile("[yaeiou]$"),
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
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("ij"),
			phoneticRules: []token{
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
					langs: 16384,
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
					langs: 160,
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
					langs: 131232,
				},
				{
					text:  "tS",
					langs: 160,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
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
			pattern: []rune("zia"),
			phoneticRules: []token{
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
					langs: 16384,
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
					langs: 16384,
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
					langs: 16384,
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
					langs: 16384,
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
			pattern: []rune("zie"),
			phoneticRules: []token{
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
			pattern: []rune("zio"),
			phoneticRules: []token{
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
			pattern: []rune("ziu"),
			phoneticRules: []token{
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
			pattern: []rune("zi"),
			phoneticRules: []token{
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
			pattern: []rune("z"),
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
					langs: 4096,
				},
				{
					text:  "Z",
					langs: 32768,
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
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phoneticRules: []token{
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
				{
					text:  "ve",
					langs: 64,
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
			pattern: []rune("ae"),
			phoneticRules: []token{
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
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  "aj",
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
					langs: 64,
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
			pattern: []rune("ão"),
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
			pattern: []rune("ãe"),
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
			pattern: []rune("ãi"),
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
			pattern: []rune("ea"),
			phoneticRules: []token{
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
			pattern: []rune("ee"),
			phoneticRules: []token{
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
			pattern: []rune("ei"),
			phoneticRules: []token{
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
			pattern: []rune("eu"),
			phoneticRules: []token{
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
			pattern: []rune("ey"),
			phoneticRules: []token{
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
			pattern: []rune("ia"),
			phoneticRules: []token{
				{
					text:  "ja",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
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
			pattern: []rune("ii"),
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
			pattern: []rune("io"),
			phoneticRules: []token{
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
			pattern: []rune("iu"),
			phoneticRules: []token{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
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
			pattern: []rune("oe"),
			phoneticRules: []token{
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
					langs: 32,
				},
				{
					text:  "o",
					langs: -1,
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
					langs: 576,
				},
				{
					text:  "au",
					langs: 16,
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
			pattern: []rune("oy"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("õe"),
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
			pattern: []rune("ua"),
			phoneticRules: []token{
				{
					text:  "va",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
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
			pattern: []rune("ui"),
			phoneticRules: []token{
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
			pattern: []rune("uu"),
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
			pattern: []rune("uy"),
			phoneticRules: []token{
				{
					text:  "uj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ya"),
			phoneticRules: []token{
				{
					text:  "ja",
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
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("yi"),
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
			pattern: []rune("yi"),
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
			pattern: []rune("yo"),
			phoneticRules: []token{
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
			pattern: []rune("yu"),
			phoneticRules: []token{
				{
					text:  "ju",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
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
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[áóéê]$"),
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
				pattern:          regexp.MustCompile("[áóéê]$"),
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
					langs: 131072,
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
					text:  "EE",
					langs: 96,
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
			pattern: []rune("ã"),
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
			pattern: []rune("ă"),
			phoneticRules: []token{
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
			pattern: []rune("ā"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("č"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
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
			pattern: []rune("ç"),
			phoneticRules: []token{
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
			pattern: []rune("ď"),
			phoneticRules: []token{
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
			pattern: []rune("ě"),
			phoneticRules: []token{
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
			pattern: []rune("ē"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ģ"),
			phoneticRules: []token{
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
			pattern: []rune("ğ"),
			phoneticRules: []token{
				{
					text:  "",
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
			pattern: []rune("ī"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ı"),
			phoneticRules: []token{
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
			pattern: []rune("ķ"),
			phoneticRules: []token{
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
			pattern: []rune("ļ"),
			phoneticRules: []token{
				{
					text:  "l",
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
				{
					text:  "nj",
					langs: 16384,
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
					langs: 262144,
				},
			},
		},
		{
			pattern: []rune("ņ"),
			phoneticRules: []token{
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
			pattern: []rune("ó"),
			phoneticRules: []token{
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
			pattern: []rune("ô"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("õ"),
			phoneticRules: []token{
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
			pattern: []rune("ò"),
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
			pattern: []rune("ř"),
			phoneticRules: []token{
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
			pattern: []rune("ś"),
			phoneticRules: []token{
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
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("š"),
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
			pattern: []rune("ť"),
			phoneticRules: []token{
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
				{
					text:  "u",
					langs: 294912,
				},
			},
		},
		{
			pattern: []rune("ū"),
			phoneticRules: []token{
				{
					text:  "u",
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
			pattern: []rune("ů"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ù"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ý"),
			phoneticRules: []token{
				{
					text:  "i",
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
					langs: 16384,
				},
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ž"),
			phoneticRules: []token{
				{
					text:  "Z",
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
					langs: 16384,
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
					text:  "B",
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
					langs: 16392,
				},
				{
					text:  "dZ",
					langs: 524288,
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
					langs: 65536,
				},
				{
					text:  "H",
					langs: 299072,
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
				{
					text:  "S",
					langs: 32768,
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
					text:  "V",
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
				{
					text:  "w",
					langs: 48,
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
	genarabic: rules{
		{
			pattern: []rune("ا"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ب"),
			rightContext: &ruleMatcher{
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
		{
			pattern: []rune("ب"),
			phoneticRules: []token{
				{
					text:  "b1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ت"),
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
			pattern: []rune("ت"),
			phoneticRules: []token{
				{
					text:  "t1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ث"),
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
			pattern: []rune("ث"),
			phoneticRules: []token{
				{
					text:  "t1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ج"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
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
			pattern: []rune("ج"),
			phoneticRules: []token{
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
			pattern: []rune("ح"),
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
			pattern: []rune("ح"),
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
			pattern: []rune("ح"),
			phoneticRules: []token{
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
			pattern: []rune("خ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("خ"),
			phoneticRules: []token{
				{
					text:  "x1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("د"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("د"),
			phoneticRules: []token{
				{
					text:  "d1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ذ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ذ"),
			phoneticRules: []token{
				{
					text:  "d1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ر"),
			rightContext: &ruleMatcher{
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
			pattern: []rune("ر"),
			phoneticRules: []token{
				{
					text:  "r1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ز"),
			rightContext: &ruleMatcher{
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
			pattern: []rune("ز"),
			phoneticRules: []token{
				{
					text:  "z1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("س"),
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
			pattern: []rune("س"),
			phoneticRules: []token{
				{
					text:  "s1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ش"),
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
			pattern: []rune("ش"),
			phoneticRules: []token{
				{
					text:  "S1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ص"),
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
			pattern: []rune("ص"),
			phoneticRules: []token{
				{
					text:  "s1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ض"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ض"),
			phoneticRules: []token{
				{
					text:  "d1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ط"),
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
			pattern: []rune("ط"),
			phoneticRules: []token{
				{
					text:  "t1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ظ"),
			rightContext: &ruleMatcher{
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
			pattern: []rune("ظ"),
			phoneticRules: []token{
				{
					text:  "z1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ع"),
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
			pattern: []rune("ع"),
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
			pattern: []rune("ع"),
			phoneticRules: []token{
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
			pattern: []rune("غ"),
			rightContext: &ruleMatcher{
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
			pattern: []rune("غ"),
			phoneticRules: []token{
				{
					text:  "g1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ف"),
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
			pattern: []rune("ف"),
			phoneticRules: []token{
				{
					text:  "f1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ق"),
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
			pattern: []rune("ق"),
			phoneticRules: []token{
				{
					text:  "k1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ك"),
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
			pattern: []rune("ك"),
			phoneticRules: []token{
				{
					text:  "k1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ل"),
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
			},
		},
		{
			pattern: []rune("ل"),
			phoneticRules: []token{
				{
					text:  "l1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("م"),
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
			pattern: []rune("م"),
			phoneticRules: []token{
				{
					text:  "m1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ن"),
			rightContext: &ruleMatcher{
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
			pattern: []rune("ن"),
			phoneticRules: []token{
				{
					text:  "n1",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ه"),
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
			pattern: []rune("ه"),
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
			pattern: []rune("ه"),
			phoneticRules: []token{
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
			pattern: []rune("و"),
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
				{
					text:  "v",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("و"),
			phoneticRules: []token{
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
			pattern: []rune("ي\u200e"),
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
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ي\u200e"),
			phoneticRules: []token{
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
	gencyrillic: rules{
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
			pattern: []rune("гейм"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
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
			pattern: []rune("ей"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
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
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			phoneticRules: []token{
				{
					text:  "ej",
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
	genczech: rules{
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
			pattern: []rune("qu"),
			phoneticRules: []token{
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
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  "aue",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
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
			pattern: []rune("č"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("š"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ž"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ň"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ť"),
			phoneticRules: []token{
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
			pattern: []rune("ď"),
			phoneticRules: []token{
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
			pattern: []rune("ř"),
			phoneticRules: []token{
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
			pattern: []rune("ý"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ě"),
			phoneticRules: []token{
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
			pattern: []rune("ů"),
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
				{
					text:  "kv",
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
	gendutch: rules{
		{
			pattern: []rune("ssj"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sj"),
			phoneticRules: []token{
				{
					text:  "S",
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
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  "k",
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
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  "s",
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
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
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
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  "au",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
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
			pattern: []rune("uu"),
			phoneticRules: []token{
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
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
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
			pattern: []rune("aa"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			phoneticRules: []token{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ui"),
			phoneticRules: []token{
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
			pattern: []rune("ei"),
			phoneticRules: []token{
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
					text:  "e",
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
				{
					text:  "x",
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
				{
					text:  "Q",
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
	genenglish: rules{
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  "",
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
			pattern: []rune("mc"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "mak",
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
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
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
			pattern: []rune("yi"),
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
			pattern: []rune("ay"),
			phoneticRules: []token{
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
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
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
			pattern: []rune("ey"),
			phoneticRules: []token{
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
					text:  "dZ",
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
				{
					text:  "a",
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
				{
					text:  "a",
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
	genfrench: rules{
		{
			pattern: []rune("lt"),
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
			phoneticRules: []token{
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
			pattern: []rune("c"),
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
			phoneticRules: []token{
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
				{
					text:  "",
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
				suffix:           "n",
			},
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
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
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
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
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
			phoneticRules: []token{
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
			pattern: []rune("t"),
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
				{
					text:  "",
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
				{
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ds"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
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
			pattern: []rune("ps"),
			rightContext: &ruleMatcher{
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
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rs"),
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
			phoneticRules: []token{
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
			pattern: []rune("ts"),
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
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
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
					text:  "",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
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
			phoneticRules: []token{
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
			pattern: []rune("s"),
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
			phoneticRules: []token{
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
			pattern: []rune("t"),
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
			phoneticRules: []token{
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
			pattern: []rune("aill"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "e",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ll"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "e",
				suffix:           "",
			},
			phoneticRules: []token{
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
			pattern: []rune("m"),
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
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
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
			pattern: []rune("au"),
			phoneticRules: []token{
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
			pattern: []rune("ai"),
			phoneticRules: []token{
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
			pattern: []rune("ay"),
			phoneticRules: []token{
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
				{
					text:  "va",
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
			pattern: []rune("ey"),
			phoneticRules: []token{
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
			pattern: []rune("eu"),
			phoneticRules: []token{
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
					text:  "e",
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
					text:  "i",
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
				{
					text:  "Q",
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
	gengerman: rules{
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
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
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
				{
					text:  "ej",
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
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
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
	gengreek: rules{
		{
			pattern: []rune("αυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("αυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phoneticRules: []token{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("αυ"),
			phoneticRules: []token{
				{
					text:  "av",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ευ"),
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
			},
		},
		{
			pattern: []rune("ευ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phoneticRules: []token{
				{
					text:  "ef",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ευ"),
			phoneticRules: []token{
				{
					text:  "ev",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "if",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phoneticRules: []token{
				{
					text:  "if",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηυ"),
			phoneticRules: []token{
				{
					text:  "iv",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ου"),
			phoneticRules: []token{
				{
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("αι"),
			phoneticRules: []token{
				{
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ει"),
			phoneticRules: []token{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("οι"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ωι"),
			phoneticRules: []token{
				{
					text:  "oj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηι"),
			phoneticRules: []token{
				{
					text:  "ej",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("υι"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γγ"),
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
			phoneticRules: []token{
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
			pattern: []rune("γγ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γγ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("γγ"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
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
			pattern: []rune("γκ"),
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
			phoneticRules: []token{
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
			pattern: []rune("γκ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("γκ"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γι"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γι"),
			phoneticRules: []token{
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
			pattern: []rune("γε"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γε"),
			phoneticRules: []token{
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
			pattern: []rune("κζ"),
			phoneticRules: []token{
				{
					text:  "gz",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("τζ"),
			phoneticRules: []token{
				{
					text:  "dz",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("σ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(β|γ|δ|μ|ν|ρ)"),
			},
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μβ"),
			phoneticRules: []token{
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
			pattern: []rune("μπ"),
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
		{
			pattern: []rune("μπ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
				{
					text:  "mb",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μπ"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ντ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ντ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("ντ"),
			phoneticRules: []token{
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
			pattern: []rune("ά"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("έ"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ή"),
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
			pattern: []rune("ί"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ό"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ύ"),
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
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ώ"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ΰ"),
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
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ϋ"),
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
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ϊ"),
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("α"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("β"),
			phoneticRules: []token{
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
			pattern: []rune("γ"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("δ"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ε"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ζ"),
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("η"),
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
			pattern: []rune("ι"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("κ"),
			phoneticRules: []token{
				{
					text:  "k",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("λ"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μ"),
			phoneticRules: []token{
				{
					text:  "m",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ν"),
			phoneticRules: []token{
				{
					text:  "n",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ξ"),
			phoneticRules: []token{
				{
					text:  "ks",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ο"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("π"),
			phoneticRules: []token{
				{
					text:  "p",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ρ"),
			phoneticRules: []token{
				{
					text:  "r",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("σ"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ς"),
			phoneticRules: []token{
				{
					text:  "s",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("τ"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("υ"),
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
					text:  "u",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("φ"),
			phoneticRules: []token{
				{
					text:  "f",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("θ"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("χ"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ψ"),
			phoneticRules: []token{
				{
					text:  "ps",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ω"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
	},
	gengreeklatin: rules{
		{
			pattern: []rune("au"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[kpstfh]"),
			},
			phoneticRules: []token{
				{
					text:  "af",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  "av",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
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
			},
		},
		{
			pattern: []rune("eu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[kpstfh]"),
			},
			phoneticRules: []token{
				{
					text:  "ef",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  "ev",
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
			pattern: []rune("gge"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("ggi"),
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
			phoneticRules: []token{
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
			pattern: []rune("ggi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("gge"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("gg"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gk"),
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
			pattern: []rune("gke"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("gki"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("gke"),
			phoneticRules: []token{
				{
					text:  "je",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gki"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gk"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("gk"),
			phoneticRules: []token{
				{
					text:  "g",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghi"),
			phoneticRules: []token{
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
			pattern: []rune("nghe"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghe"),
			phoneticRules: []token{
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
			pattern: []rune("ghi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ghi"),
			phoneticRules: []token{
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
			pattern: []rune("ghe"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ghe"),
			phoneticRules: []token{
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
			pattern: []rune("ngh"),
			phoneticRules: []token{
				{
					text:  "Ng",
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
			pattern: []rune("ngi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ngi"),
			phoneticRules: []token{
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
			pattern: []rune("nge"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "Nj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nge"),
			phoneticRules: []token{
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
			pattern: []rune("gi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
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
			pattern: []rune("ge"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  "j",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ge"),
			phoneticRules: []token{
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
			pattern: []rune("ng"),
			phoneticRules: []token{
				{
					text:  "Ng",
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
			pattern: []rune("i"),
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
			pattern: []rune("y"),
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
			pattern: []rune("yi"),
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
			pattern: []rune("yi"),
			phoneticRules: []token{
				{
					text:  "i",
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
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("dh"),
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("dj"),
			phoneticRules: []token{
				{
					text:  "dZ",
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
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  "t",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kz"),
			phoneticRules: []token{
				{
					text:  "gz",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  "dz",
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
				pattern:          regexp.MustCompile("^[bgdmnr]"),
			},
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mb"),
			phoneticRules: []token{
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
			pattern: []rune("mp"),
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
		{
			pattern: []rune("mp"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  "mp",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mp"),
			phoneticRules: []token{
				{
					text:  "b",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nt"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phoneticRules: []token{
				{
					text:  "d",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nt"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
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
			pattern: []rune("nt"),
			phoneticRules: []token{
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
			pattern: []rune("óu"),
			phoneticRules: []token{
				{
					text:  "u",
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
			pattern: []rune("ý"),
			phoneticRules: []token{
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
				{
					text:  "v",
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
					text:  "e",
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
					text:  "x",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  "i",
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
				{
					text:  "Z",
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
			pattern: []rune("ο"),
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
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  "z",
					langs: -1,
				},
			},
		},
	},
	genhebrew: rules{
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
	genhungarian: rules{
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
				{
					text:  "ej",
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
					text:  "ej",
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
			pattern: []rune("ely"),
			phoneticRules: []token{
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
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  "u",
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
	genitalian: rules{
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
			pattern: []rune("gli"),
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
			pattern: []rune("gn"),
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
			pattern: []rune("gni"),
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
			pattern: []rune("gi"),
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
			pattern: []rune("gg"),
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
					text:  "g",
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
			pattern: []rune("ci"),
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
			pattern: []rune("sc"),
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
			pattern: []rune("cc"),
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
					text:  "tS",
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
			pattern: []rune("i"),
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
			pattern: []rune("y"),
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
			pattern: []rune("qu"),
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
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  "o",
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
					text:  "e",
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
					text:  "i",
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
				{
					text:  "dz",
					langs: -1,
				},
			},
		},
	},
	genlatvian: rules{
		{
			pattern: []rune("č"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ģ"),
			phoneticRules: []token{
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
			pattern: []rune("ķ"),
			phoneticRules: []token{
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
			pattern: []rune("ļ"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("š"),
			phoneticRules: []token{
				{
					text:  "S",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ņ"),
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
			pattern: []rune("ž"),
			phoneticRules: []token{
				{
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ā"),
			phoneticRules: []token{
				{
					text:  "a",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ē"),
			phoneticRules: []token{
				{
					text:  "e",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ī"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ū"),
			phoneticRules: []token{
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
					text:  "aj",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  "ej",
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
			},
		},
		{
			pattern: []rune("iu"),
			phoneticRules: []token{
				{
					text:  "ju",
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
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  "o",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ui"),
			phoneticRules: []token{
				{
					text:  "uj",
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
	genpolish: rules{
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
			pattern: []rune("y"),
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
	genportuguese: rules{
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
					text:  "S",
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
			pattern: []rune("sc"),
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
			pattern: []rune("sç"),
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
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  "s",
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
			pattern: []rune("s"),
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
			pattern: []rune("s"),
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
			pattern: []rune("s"),
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
			pattern: []rune("z"),
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
			pattern: []rune("z"),
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
			pattern: []rune("z"),
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
			pattern: []rune("gu"),
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
			pattern: []rune("gu"),
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
					text:  "Z",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
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
			pattern: []rune("qu"),
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
			pattern: []rune("lh"),
			phoneticRules: []token{
				{
					text:  "l",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nh"),
			phoneticRules: []token{
				{
					text:  "nj",
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
			pattern: []rune("ex"),
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
			pattern: []rune("ex"),
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
			pattern: []rune("y"),
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
			pattern: []rune("y"),
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
			pattern: []rune("m"),
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
			pattern: []rune("m"),
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
			pattern: []rune("ão"),
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
			pattern: []rune("ãe"),
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
			pattern: []rune("ãi"),
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
			pattern: []rune("õe"),
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
			pattern: []rune("i"),
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
			pattern: []rune("â"),
			phoneticRules: []token{
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
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  "a",
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
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  "i",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ô"),
			phoneticRules: []token{
				{
					text:  "o",
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
			pattern: []rune("õ"),
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
					text:  "u",
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
					text:  "i",
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
				{
					text:  "u",
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
	genromanian: rules{
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
			pattern: []rune("i"),
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
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  "k",
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
					text:  "Z",
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
	genrussian: rules{
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
				pattern:          regexp.MustCompile("[aeiou]$"),
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
			pattern: []rune("iy"),
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
			pattern: []rune("yy"),
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
			pattern: []rune("yi"),
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
	genspanish: rules{
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
			pattern: []rune("ny"),
			phoneticRules: []token{
				{
					text:  "nj",
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
			pattern: []rune("ig"),
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
			pattern: []rune("ix"),
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
			pattern: []rune("tx"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tj"),
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
			pattern: []rune("tj"),
			phoneticRules: []token{
				{
					text:  "dZ",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tg"),
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
			pattern: []rune("bh"),
			phoneticRules: []token{
				{
					text:  "b",
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
				{
					text:  "dZ",
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
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  "a",
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
			pattern: []rune("ò"),
			phoneticRules: []token{
				{
					text:  "o",
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
					text:  "B",
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
					text:  "e",
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
					text:  "i",
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
				{
					text:  "Z",
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
					text:  "V",
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
	},
	genturkish: rules{
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  "tS",
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ğ"),
			phoneticRules: []token{
				{
					text:  "",
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
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  "Q",
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
			pattern: []rune("ı"),
			phoneticRules: []token{
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
					text:  "dZ",
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
					text:  "e",
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
					text:  "i",
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
					text:  "j",
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
					pattern:          regexp.MustCompile("^[aA]"),
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
					suffix:           "A",
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
				pattern: []rune("j"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "j",
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
				pattern: []rune("vanden"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
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
				pattern: []rune("vander"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
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
				pattern: []rune("van"),
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
				phoneticRules: []token{
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
				pattern: []rune("van"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
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
				pattern: []rune("sen"),
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
				phoneticRules: []token{
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
				pattern: []rune("sen"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
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
				pattern: []rune("sEn"),
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
				phoneticRules: []token{
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
				pattern: []rune("sEn"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phoneticRules: []token{
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
				pattern: []rune("e"),
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
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
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
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
				pattern: []rune("Burk"),
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
				pattern: []rune("BUrk"),
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
				pattern: []rune("Burg"),
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
				pattern: []rune("BUrg"),
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
			languageID(genany): rules{
				{
					pattern: []rune("mb"),
					phoneticRules: []token{
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
					pattern: []rune("mp"),
					phoneticRules: []token{
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
					pattern: []rune("ng"),
					phoneticRules: []token{
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
					pattern: []rune("B"),
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
							text:  "f",
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("B"),
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
					pattern: []rune("B"),
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
						{
							text:  "f",
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("V"),
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
							text:  "p",
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("V"),
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
					pattern: []rune("V"),
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
						{
							text:  "p",
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
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
					pattern: []rune("V"),
					phoneticRules: []token{
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
					pattern: []rune("t"),
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
						{
							text:  "",
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
						suffix:           "n",
					},
					rightContext: &ruleMatcher{
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
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("k"),
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
					phoneticRules: []token{
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
					pattern: []rune("p"),
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
						{
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("r"),
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
					phoneticRules: []token{
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
					pattern: []rune("s"),
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
							text:  "",
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("t"),
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
					phoneticRules: []token{
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
					pattern: []rune("s"),
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
					phoneticRules: []token{
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
							langs: 128,
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
							langs: 128,
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
							langs: 128,
						},
						{
							text:  "i",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lEE"),
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
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("rEE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
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
							langs: 32,
						},
						{
							text:  "lY",
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("rE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
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
					pattern: []rune("EE"),
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
					pattern: []rune("ea"),
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
					pattern: []rune("eu"),
					phoneticRules: []token{
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
							langs: 128,
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
							langs: 128,
						},
						{
							text:  "D",
							langs: 32,
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
							langs: 128,
						},
						{
							text:  "",
							langs: 32,
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
							langs: 128,
						},
						{
							text:  "",
							langs: 32,
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
							langs: 128,
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
							langs: 128,
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
							langs: 128,
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
							langs: 128,
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
							langs: 128,
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
							langs: 128,
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
							langs: 128,
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
			languageID(genarabic): rules{
				{
					pattern: []rune("1a"),
					phoneticRules: []token{
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
					pattern: []rune("1i"),
					phoneticRules: []token{
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
					pattern: []rune("1u"),
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
							text:  "o",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("j1"),
					phoneticRules: []token{
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
					pattern: []rune("1"),
					phoneticRules: []token{
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
					pattern: []rune("u"),
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
					pattern: []rune("i"),
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
					pattern: []rune("p"),
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
					phoneticRules: []token{
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
			languageID(genrussian): rules{
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
			languageID(gencyrillic): rules{
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
			languageID(genfrench): rules{
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
			},
			languageID(genczech): rules{
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
			},
			languageID(gendutch): rules{
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
			},
			languageID(genenglish): rules{
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
			languageID(gengerman): rules{
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
			languageID(gengreek): rules{
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
			},
			languageID(gengreeklatin): rules{
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
					pattern: []rune("N"),
					phoneticRules: []token{
						{
							text:  "",
							langs: -1,
						},
					},
				},
			},
			languageID(genhebrew): rules{},
			languageID(genhungarian): rules{
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
			},
			languageID(genitalian): rules{
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
			},
			languageID(genlatvian): rules{
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
			},
			languageID(genpolish): rules{
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
			languageID(genportuguese): rules{
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
			},
			languageID(genromanian): rules{
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
			languageID(genspanish): rules{
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
					pattern: []rune("B"),
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
					pattern: []rune("V"),
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
			},
			languageID(genturkish): rules{
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
					pattern:          regexp.MustCompile("^[aA]"),
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
					suffix:           "A",
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
				pattern: []rune("j"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "j",
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
						text:  "",
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
			languageID(genany): rules{
				{
					pattern: []rune("EE"),
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
					},
				},
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
					pattern: []rune("B"),
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
					pattern: []rune("B"),
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
					pattern: []rune("B"),
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
					pattern: []rune("V"),
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
					pattern: []rune("V"),
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
					pattern: []rune("V"),
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
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "b",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  "v",
							langs: -1,
						},
					},
				},
			},
			languageID(genarabic): rules{
				{
					pattern: []rune("1"),
					phoneticRules: []token{
						{
							text:  "",
							langs: -1,
						},
					},
				},
			},
			languageID(genrussian): rules{
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
			languageID(gencyrillic): rules{
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
			languageID(genczech): rules{
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
			languageID(gendutch): rules{},
			languageID(genenglish): rules{
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
			languageID(genfrench): rules{},
			languageID(gengerman): rules{
				{
					pattern: []rune("EE"),
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
					},
				},
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
					pattern: []rune("B"),
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
					pattern: []rune("B"),
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
					pattern: []rune("B"),
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
					pattern: []rune("V"),
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
					pattern: []rune("V"),
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
					pattern: []rune("V"),
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
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "b",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  "v",
							langs: -1,
						},
					},
				},
			},
			languageID(gengreek): rules{},
			languageID(gengreeklatin): rules{
				{
					pattern: []rune("N"),
					phoneticRules: []token{
						{
							text:  "n",
							langs: -1,
						},
					},
				},
			},
			languageID(genhebrew): rules{},
			languageID(genhungarian): rules{
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
			languageID(genitalian): rules{},
			languageID(genlatvian): rules{},
			languageID(genpolish): rules{
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
			languageID(genportuguese): rules{},
			languageID(genromanian): rules{
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
			languageID(genspanish): rules{
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  "b",
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  "v",
							langs: -1,
						},
					},
				},
			},
			languageID(genturkish): rules{},
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
