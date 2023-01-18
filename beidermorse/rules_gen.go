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
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: 131072,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ina"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: 131072,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lova"),
					langs: -1,
				},
				{
					text:  []rune("lof"),
					langs: 131072,
				},
				{
					text:  []rune("lef"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("lova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lova"),
					langs: -1,
				},
				{
					text:  []rune("lof"),
					langs: 131072,
				},
				{
					text:  []rune("lef"),
					langs: 131072,
				},
				{
					text:  []rune("l"),
					langs: 8,
				},
				{
					text:  []rune("el"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("kova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kova"),
					langs: -1,
				},
				{
					text:  []rune("kof"),
					langs: 131072,
				},
				{
					text:  []rune("k"),
					langs: 8,
				},
				{
					text:  []rune("ek"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ova"),
					langs: -1,
				},
				{
					text:  []rune("of"),
					langs: 131072,
				},
				{
					text:  nil,
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ová"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ova"),
					langs: -1,
				},
				{
					text:  nil,
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("eva"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("eva"),
					langs: -1,
				},
				{
					text:  []rune("ef"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("aia"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("aja"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("aya"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("lowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lova"),
					langs: -1,
				},
				{
					text:  []rune("lof"),
					langs: 16384,
				},
				{
					text:  []rune("l"),
					langs: 16384,
				},
				{
					text:  []rune("el"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("kowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kova"),
					langs: -1,
				},
				{
					text:  []rune("kof"),
					langs: 16384,
				},
				{
					text:  []rune("k"),
					langs: 16384,
				},
				{
					text:  []rune("ek"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("owa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ova"),
					langs: -1,
				},
				{
					text:  []rune("of"),
					langs: 16384,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lovna"),
					langs: -1,
				},
				{
					text:  []rune("levna"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: 16384,
				},
				{
					text:  []rune("el"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("kowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kovna"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: 16384,
				},
				{
					text:  []rune("ek"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("owna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovna"),
					langs: -1,
				},
				{
					text:  nil,
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("lówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: 16392,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  []rune("pf"),
					langs: -1,
				},
				{
					text:  []rune("p"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("que"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: 64,
				},
				{
					text:  []rune("ke"),
					langs: -1,
				},
				{
					text:  []rune("kve"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bfpv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: 32832,
				},
			},
		},
		{
			pattern: []rune("ly"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("lj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("li"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("lj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("lyo"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("lt"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("u"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lt"),
					langs: -1,
				},
				{
					text:  nil,
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("v"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: 128,
				},
				{
					text:  []rune("b"),
					langs: 262144,
				},
			},
		},
		{
			pattern: []rune("ex"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ez"),
					langs: 32768,
				},
				{
					text:  []rune("eS"),
					langs: 32768,
				},
				{
					text:  []rune("eks"),
					langs: -1,
				},
				{
					text:  []rune("egz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ex"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[cs]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: 32768,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("u"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
				{
					text:  nil,
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("tsk"),
					langs: 16392,
				},
			},
		},
		{
			pattern: []rune("cz"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("tsz"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("rh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("dh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("bh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("ph"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: 131104,
				},
				{
					text:  []rune("kh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lh"),
			phoneticRules: []token{
				{
					text:  []rune("lh"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("nh"),
			phoneticRules: []token{
				{
					text:  []rune("nh"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("ssch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chsch"),
			phoneticRules: []token{
				{
					text:  []rune("xS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("StS"),
					langs: 131072,
				},
				{
					text:  []rune("sk"),
					langs: 69632,
				},
			},
		},
		{
			pattern: []rune("sch"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("StS"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("sch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("sk"),
					langs: 69632,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("StS"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("StS"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("ssh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöü]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("sh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 131104,
				},
				{
					text:  []rune("sh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zh"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: 131104,
				},
				{
					text:  []rune("zh"),
					langs: -1,
				},
				{
					text:  []rune("tsh"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("chs"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: 128,
				},
				{
					text:  []rune("xs"),
					langs: -1,
				},
				{
					text:  []rune("tSs"),
					langs: 131104,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 393248,
				},
				{
					text:  []rune("k"),
					langs: 69632,
				},
				{
					text:  []rune("S"),
					langs: 32832,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 393248,
				},
				{
					text:  []rune("S"),
					langs: 32832,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: 672,
				},
				{
					text:  []rune("th"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: 70144,
				},
				{
					text:  []rune("gh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ouh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: 64,
				},
				{
					text:  []rune("uh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
				{
					text:  []rune("uh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: 66048,
				},
				{
					text:  []rune("H"),
					langs: 381024,
				},
			},
		},
		{
			pattern: []rune("cia"),
			phoneticRules: []token{
				{
					text:  []rune("tSa"),
					langs: 16384,
				},
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSom"),
					langs: -1,
				},
				{
					text:  []rune("tsom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			phoneticRules: []token{
				{
					text:  []rune("tSon"),
					langs: 16384,
				},
				{
					text:  []rune("tson"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSem"),
					langs: 16384,
				},
				{
					text:  []rune("tsem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			phoneticRules: []token{
				{
					text:  []rune("tSen"),
					langs: 16384,
				},
				{
					text:  []rune("tsen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: 16384,
				},
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cio"),
			phoneticRules: []token{
				{
					text:  []rune("tSo"),
					langs: 16384,
				},
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciu"),
			phoneticRules: []token{
				{
					text:  []rune("tSu"),
					langs: 16384,
				},
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sci"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: 4096,
				},
				{
					text:  []rune("stsi"),
					langs: 16392,
				},
				{
					text:  []rune("dZi"),
					langs: 524288,
				},
				{
					text:  []rune("tSi"),
					langs: 81920,
				},
				{
					text:  []rune("tS"),
					langs: 65536,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 4096,
				},
				{
					text:  []rune("sts"),
					langs: 16392,
				},
				{
					text:  []rune("dZ"),
					langs: 524288,
				},
				{
					text:  []rune("tS"),
					langs: 81920,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tsi"),
					langs: 16392,
				},
				{
					text:  []rune("dZi"),
					langs: 524288,
				},
				{
					text:  []rune("tSi"),
					langs: 81920,
				},
				{
					text:  []rune("tS"),
					langs: 65536,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cy"),
			phoneticRules: []token{
				{
					text:  []rune("si"),
					langs: -1,
				},
				{
					text:  []rune("tsi"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: 16392,
				},
				{
					text:  []rune("dZ"),
					langs: 524288,
				},
				{
					text:  []rune("tS"),
					langs: 81920,
				},
				{
					text:  []rune("k"),
					langs: 512,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sç"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("stS"),
					langs: 524288,
				},
			},
		},
		{
			pattern: []rune("ssz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: 2048,
				},
			},
		},
		{
			pattern: []rune("sz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: 2048,
				},
			},
		},
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: 2048,
				},
				{
					text:  []rune("sts"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("ssp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: 128,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: 128,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sst"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: 128,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: 128,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sj"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sj"),
			phoneticRules: []token{
				{
					text:  []rune("sj"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: 16,
				},
				{
					text:  []rune("sx"),
					langs: 262144,
				},
				{
					text:  []rune("sZ"),
					langs: 589824,
				},
			},
		},
		{
			pattern: []rune("sia"),
			phoneticRules: []token{
				{
					text:  []rune("Sa"),
					langs: 16384,
				},
				{
					text:  []rune("sa"),
					langs: 16384,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Som"),
					langs: 16384,
				},
				{
					text:  []rune("som"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			phoneticRules: []token{
				{
					text:  []rune("Son"),
					langs: 16384,
				},
				{
					text:  []rune("son"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Sem"),
					langs: 16384,
				},
				{
					text:  []rune("sem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			phoneticRules: []token{
				{
					text:  []rune("Sen"),
					langs: 16384,
				},
				{
					text:  []rune("sen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
				{
					text:  []rune("sje"),
					langs: -1,
				},
				{
					text:  []rune("Se"),
					langs: 16384,
				},
				{
					text:  []rune("zi"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: 16384,
				},
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siu"),
			phoneticRules: []token{
				{
					text:  []rune("Su"),
					langs: 16384,
				},
				{
					text:  []rune("sju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("si"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[äöëaáuiíoóeéêy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: 16384,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
				{
					text:  []rune("zi"),
					langs: 37056,
				},
			},
		},
		{
			pattern: []rune("si"),
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: 16384,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
				{
					text:  []rune("zi"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aáuíoóeéêy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: 37056,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouäöë]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: 32768,
				},
				{
					text:  nil,
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("gue"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: 64,
				},
				{
					text:  []rune("gve"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: 64,
				},
				{
					text:  []rune("gv"),
					langs: 294912,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("guy"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gli"),
			phoneticRules: []token{
				{
					text:  []rune("glI"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: 4096,
				},
			},
		},
		{
			pattern: []rune("gni"),
			phoneticRules: []token{
				{
					text:  []rune("gnI"),
					langs: -1,
				},
				{
					text:  []rune("ni"),
					langs: 4160,
				},
			},
		},
		{
			pattern: []rune("gn"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: 4160,
				},
				{
					text:  []rune("nj"),
					langs: 4160,
				},
				{
					text:  []rune("gn"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: 512,
				},
				{
					text:  []rune("dZe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: 512,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: 4096,
				},
				{
					text:  []rune("j"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("gge"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gE"),
					langs: -1,
				},
				{
					text:  []rune("xe"),
					langs: 262144,
				},
				{
					text:  []rune("gZe"),
					langs: 32832,
				},
				{
					text:  []rune("dZe"),
					langs: 331808,
				},
				{
					text:  []rune("je"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("xi"),
					langs: 262144,
				},
				{
					text:  []rune("gZi"),
					langs: 32832,
				},
				{
					text:  []rune("dZi"),
					langs: 331808,
				},
				{
					text:  []rune("i"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: 4096,
				},
				{
					text:  []rune("j"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("gie"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ge"),
					langs: -1,
				},
				{
					text:  []rune("gi"),
					langs: 128,
				},
				{
					text:  []rune("ji"),
					langs: 64,
				},
				{
					text:  []rune("dZe"),
					langs: 4096,
				},
			},
		},
		{
			pattern: []rune("gie"),
			phoneticRules: []token{
				{
					text:  []rune("ge"),
					langs: -1,
				},
				{
					text:  []rune("gi"),
					langs: 128,
				},
				{
					text:  []rune("dZe"),
					langs: 4096,
				},
				{
					text:  []rune("je"),
					langs: 512,
				},
			},
		},
		{
			pattern: []rune("gi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: 512,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ge"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gE"),
					langs: -1,
				},
				{
					text:  []rune("xe"),
					langs: 262144,
				},
				{
					text:  []rune("Ze"),
					langs: 32832,
				},
				{
					text:  []rune("dZe"),
					langs: 331808,
				},
			},
		},
		{
			pattern: []rune("gi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("xi"),
					langs: 262144,
				},
				{
					text:  []rune("Zi"),
					langs: 32832,
				},
				{
					text:  []rune("dZi"),
					langs: 331808,
				},
			},
		},
		{
			pattern: []rune("ge"),
			phoneticRules: []token{
				{
					text:  []rune("gE"),
					langs: -1,
				},
				{
					text:  []rune("xe"),
					langs: 262144,
				},
				{
					text:  []rune("hE"),
					langs: 131072,
				},
				{
					text:  []rune("je"),
					langs: 512,
				},
				{
					text:  []rune("Ze"),
					langs: 32832,
				},
				{
					text:  []rune("dZe"),
					langs: 331808,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  []rune("gI"),
					langs: -1,
				},
				{
					text:  []rune("xi"),
					langs: 262144,
				},
				{
					text:  []rune("hI"),
					langs: 131072,
				},
				{
					text:  []rune("i"),
					langs: 512,
				},
				{
					text:  []rune("Zi"),
					langs: 32832,
				},
				{
					text:  []rune("dZi"),
					langs: 331808,
				},
			},
		},
		{
			pattern: []rune("gy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("dj"),
					langs: 2048,
				},
			},
		},
		{
			pattern: []rune("gy"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("d"),
					langs: 2048,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouyei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("ij"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: 16,
				},
				{
					text:  []rune("ix"),
					langs: 262144,
				},
				{
					text:  []rune("iZ"),
					langs: 622656,
				},
			},
		},
		{
			pattern: []rune("j"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aoeiuy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: 32,
				},
				{
					text:  []rune("x"),
					langs: 262144,
				},
				{
					text:  []rune("Z"),
					langs: 622656,
				},
			},
		},
		{
			pattern: []rune("rz"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 16384,
				},
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			phoneticRules: []token{
				{
					text:  []rune("rz"),
					langs: -1,
				},
				{
					text:  []rune("rts"),
					langs: 128,
				},
				{
					text:  []rune("Z"),
					langs: 16384,
				},
				{
					text:  []rune("r"),
					langs: 16384,
				},
				{
					text:  []rune("rZ"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 160,
				},
			},
		},
		{
			pattern: []rune("tz"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: 131232,
				},
				{
					text:  []rune("tS"),
					langs: 160,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: 131232,
				},
				{
					text:  []rune("tz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Za"),
					langs: 16384,
				},
				{
					text:  []rune("za"),
					langs: 16384,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  []rune("Za"),
					langs: 16384,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zom"),
					langs: 16384,
				},
				{
					text:  []rune("zom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			phoneticRules: []token{
				{
					text:  []rune("Zon"),
					langs: 16384,
				},
				{
					text:  []rune("zon"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zem"),
					langs: 16384,
				},
				{
					text:  []rune("zem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			phoneticRules: []token{
				{
					text:  []rune("Zen"),
					langs: 16384,
				},
				{
					text:  []rune("zen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: 16384,
				},
				{
					text:  []rune("ze"),
					langs: 16384,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
				{
					text:  []rune("tsi"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
				{
					text:  []rune("Ze"),
					langs: 16384,
				},
				{
					text:  []rune("tsi"),
					langs: 128,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("Zo"),
					langs: 16384,
				},
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  []rune("Zu"),
					langs: 16384,
				},
				{
					text:  []rune("zju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zi"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: 16384,
				},
				{
					text:  []rune("zi"),
					langs: -1,
				},
				{
					text:  []rune("tsi"),
					langs: 128,
				},
				{
					text:  []rune("dzi"),
					langs: 4096,
				},
				{
					text:  []rune("tsi"),
					langs: 4096,
				},
				{
					text:  []rune("si"),
					langs: 262144,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: 128,
				},
				{
					text:  []rune("ts"),
					langs: 4096,
				},
				{
					text:  []rune("S"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bdgv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("dz"),
					langs: 4096,
				},
				{
					text:  []rune("Z"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: 4096,
				},
				{
					text:  []rune("S"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oue"),
			phoneticRules: []token{
				{
					text:  []rune("oue"),
					langs: -1,
				},
				{
					text:  []rune("ve"),
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("eau"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ae"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: 128,
				},
				{
					text:  []rune("aje"),
					langs: 131072,
				},
				{
					text:  []rune("ae"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: 64,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ão"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ãe"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ãi"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  []rune("ea"),
					langs: -1,
				},
				{
					text:  []rune("ja"),
					langs: 65536,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: 32,
				},
				{
					text:  []rune("aje"),
					langs: 131072,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("eu"),
					langs: -1,
				},
				{
					text:  []rune("Yj"),
					langs: 128,
				},
				{
					text:  []rune("ej"),
					langs: 128,
				},
				{
					text:  []rune("oj"),
					langs: 128,
				},
				{
					text:  []rune("Y"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ia"),
			phoneticRules: []token{
				{
					text:  []rune("ja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: 128,
				},
				{
					text:  []rune("e"),
					langs: 16384,
				},
				{
					text:  []rune("ije"),
					langs: 131072,
				},
				{
					text:  []rune("Q"),
					langs: 16,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("iu"),
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: 128,
				},
				{
					text:  []rune("oje"),
					langs: 131072,
				},
				{
					text:  []rune("u"),
					langs: 16,
				},
				{
					text:  []rune("oe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: 32,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("ou"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: 576,
				},
				{
					text:  []rune("au"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("où"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oy"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("õe"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ua"),
			phoneticRules: []token{
				{
					text:  []rune("va"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: 128,
				},
				{
					text:  []rune("uje"),
					langs: 131072,
				},
				{
					text:  []rune("ve"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ui"),
			phoneticRules: []token{
				{
					text:  []rune("uj"),
					langs: -1,
				},
				{
					text:  []rune("vi"),
					langs: -1,
				},
				{
					text:  []rune("Y"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("uu"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("Q"),
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  []rune("vo"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uy"),
			phoneticRules: []token{
				{
					text:  []rune("uj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ya"),
			phoneticRules: []token{
				{
					text:  []rune("ja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
				{
					text:  []rune("ije"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("yi"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yo"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("yu"),
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[áóéê]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[áóéê]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: 131072,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("EE"),
					langs: 96,
				},
			},
		},
		{
			pattern: []rune("ą"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("om"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			phoneticRules: []token{
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ä"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ã"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ă"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: 65536,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ā"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("č"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: 16384,
				},
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 524288,
				},
			},
		},
		{
			pattern: []rune("ď"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
				{
					text:  []rune("dj"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ę"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("em"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			phoneticRules: []token{
				{
					text:  []rune("en"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ě"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ē"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ģ"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
				{
					text:  []rune("dj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ğ"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("î"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ī"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ı"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: 524288,
				},
				{
					text:  nil,
					langs: 524288,
				},
			},
		},
		{
			pattern: []rune("ķ"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("t"),
					langs: 8192,
				},
				{
					text:  []rune("tj"),
					langs: 8192,
				},
			},
		},
		{
			pattern: []rune("ļ"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ł"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ń"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: 262144,
				},
			},
		},
		{
			pattern: []rune("ņ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: 8192,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: 16384,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ô"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("õ"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("on"),
					langs: 32768,
				},
				{
					text:  []rune("Y"),
					langs: 2048,
				},
			},
		},
		{
			pattern: []rune("ò"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ř"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
				{
					text:  []rune("rZ"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ś"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: 16384,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("š"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ţ"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ť"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("tj"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: 294912,
				},
			},
		},
		{
			pattern: []rune("ū"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ů"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ù"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ý"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ż"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ź"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: 16384,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ž"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ß"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("\""),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("O"),
					langs: -1,
				},
				{
					text:  []rune("P"),
					langs: 16384,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("A"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("B"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: 16392,
				},
				{
					text:  []rune("dZ"),
					langs: 524288,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: 65536,
				},
				{
					text:  []rune("H"),
					langs: 299072,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: 262144,
				},
				{
					text:  []rune("Z"),
					langs: 622656,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("O"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: 32768,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("U"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("V"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
				{
					text:  []rune("w"),
					langs: 48,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
				{
					text:  []rune("gz"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: 294912,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: 128,
				},
				{
					text:  []rune("dz"),
					langs: 4096,
				},
				{
					text:  []rune("ts"),
					langs: 4096,
				},
				{
					text:  []rune("s"),
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
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ب"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ب"),
			phoneticRules: []token{
				{
					text:  []rune("b1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ت"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ت"),
			phoneticRules: []token{
				{
					text:  []rune("t1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ث"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ث"),
			phoneticRules: []token{
				{
					text:  []rune("t1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ج"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ج"),
			phoneticRules: []token{
				{
					text:  []rune("dZ1"),
					langs: -1,
				},
				{
					text:  []rune("Z1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ح"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ح"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ح"),
			phoneticRules: []token{
				{
					text:  []rune("h1"),
					langs: -1,
				},
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("خ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("خ"),
			phoneticRules: []token{
				{
					text:  []rune("x1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("د"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("د"),
			phoneticRules: []token{
				{
					text:  []rune("d1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ذ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ذ"),
			phoneticRules: []token{
				{
					text:  []rune("d1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ر"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ر"),
			phoneticRules: []token{
				{
					text:  []rune("r1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ز"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ز"),
			phoneticRules: []token{
				{
					text:  []rune("z1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("س"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("س"),
			phoneticRules: []token{
				{
					text:  []rune("s1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ش"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ش"),
			phoneticRules: []token{
				{
					text:  []rune("S1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ص"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ص"),
			phoneticRules: []token{
				{
					text:  []rune("s1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ض"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ض"),
			phoneticRules: []token{
				{
					text:  []rune("d1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ط"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ط"),
			phoneticRules: []token{
				{
					text:  []rune("t1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ظ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ظ"),
			phoneticRules: []token{
				{
					text:  []rune("z1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ع"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ع"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ع"),
			phoneticRules: []token{
				{
					text:  []rune("h1"),
					langs: -1,
				},
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("غ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("غ"),
			phoneticRules: []token{
				{
					text:  []rune("g1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ف"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ف"),
			phoneticRules: []token{
				{
					text:  []rune("f1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ق"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ق"),
			phoneticRules: []token{
				{
					text:  []rune("k1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ك"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ك"),
			phoneticRules: []token{
				{
					text:  []rune("k1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ل"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ل"),
			phoneticRules: []token{
				{
					text:  []rune("l1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("م"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("م"),
			phoneticRules: []token{
				{
					text:  []rune("m1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ن"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ن"),
			phoneticRules: []token{
				{
					text:  []rune("n1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ه"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ه"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ه"),
			phoneticRules: []token{
				{
					text:  []rune("h1"),
					langs: -1,
				},
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("و"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("و"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("v1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ي\u200e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ي\u200e"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("j1"),
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
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цю"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("циа"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цие"),
			phoneticRules: []token{
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("цио"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("циу"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("сие"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("сио"),
			phoneticRules: []token{
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("зие"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("зио"),
			phoneticRules: []token{
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("с"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("с"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гауз"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гаус"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("haus"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гольц"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("holts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("геймер"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hejmer"),
					langs: -1,
				},
				{
					text:  []rune("hajmer"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гейм"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hejm"),
					langs: -1,
				},
				{
					text:  []rune("hajm"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гоф"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hof"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гер"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ger"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ген"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("гин"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gin"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(й|ё|я|ю|ы|а|е|о|и|у)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ля"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лю"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лё"),
			phoneticRules: []token{
				{
					text:  []rune("le"),
					langs: -1,
				},
				{
					text:  []rune("lo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("лио"),
			phoneticRules: []token{
				{
					text:  []rune("le"),
					langs: -1,
				},
				{
					text:  []rune("lo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ле"),
			phoneticRules: []token{
				{
					text:  []rune("lE"),
					langs: -1,
				},
				{
					text:  []rune("lo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ийе"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ие"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ыйе"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ые"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ий"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ый"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(а|о|у)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ий"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ый"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("jej"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(а|е|о|у)$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("эй"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ей"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ауе"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ауэ"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("а"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("б"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("в"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("г"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("д"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("е"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ё"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("jo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ж"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("з"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("и"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("й"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("к"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("л"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("м"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("н"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("о"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("п"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("р"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("с"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("т"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("у"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ф"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("х"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ц"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ч"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ш"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("щ"),
			phoneticRules: []token{
				{
					text:  []rune("StS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ъ"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ы"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ь"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("э"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ю"),
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("я"),
			phoneticRules: []token{
				{
					text:  []rune("ja"),
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
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("č"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("š"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ž"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ň"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ť"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("tj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ď"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
				{
					text:  []rune("dj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ř"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
				{
					text:  []rune("rZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ý"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ě"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ů"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sj"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  []rune("pf"),
					langs: -1,
				},
				{
					text:  []rune("p"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("ph"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("th"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uu"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
				{
					text:  []rune("Yj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aa"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ui"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
				{
					text:  []rune("uj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("w"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mc"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("mak"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("c"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
				{
					text:  []rune("w"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			phoneticRules: []token{
				{
					text:  []rune("gn"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("sk"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("who"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("hu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("w"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("H"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kn"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mb"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ng"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("N"),
					langs: -1,
				},
				{
					text:  []rune("ng"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pn"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("pn"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ps"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ps"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kw"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tia"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: -1,
				},
				{
					text:  []rune("Sa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tio"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wr"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
				{
					text:  []rune("oue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ear"),
			phoneticRules: []token{
				{
					text:  []rune("ia"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oa"),
			phoneticRules: []token{
				{
					text:  []rune("ou"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("ou"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oy"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ou"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("r"),
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("w"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
				suffix:           []rune("u"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lt"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("n"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("n"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("e"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ds"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ds"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ps"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ps"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rs"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("e"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("rs"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ts"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("u"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeéèêiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeéèêiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[^aeéèêiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("gn"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gue"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aill"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ll"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("e"),
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("que"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyéèê]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eau"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("au"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("où"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oi"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("va"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
			},
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("evitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ovitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("witsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vitsch"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chsch"),
			phoneticRules: []token{
				{
					text:  []rune("xS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("chs"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ck"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: -1,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sp"),
			phoneticRules: []token{
				{
					text:  []rune("Sp"),
					langs: -1,
				},
				{
					text:  []rune("sp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sst"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: -1,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("st"),
			phoneticRules: []token{
				{
					text:  []rune("St"),
					langs: -1,
				},
				{
					text:  []rune("st"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("pf"),
			phoneticRules: []token{
				{
					text:  []rune("pf"),
					langs: -1,
				},
				{
					text:  []rune("p"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("ph"),
					langs: -1,
				},
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("evitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eviz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("evits"),
					langs: -1,
				},
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ovitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oviz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovits"),
					langs: -1,
				},
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("witz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("wiz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("vitz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("viz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("vits"),
					langs: -1,
				},
				{
					text:  []rune("vitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("thal"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tal"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöüaeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("th"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rh"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyäöü]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("H"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[äöüaeiouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouyäöüj]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiouyäöü]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ß"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ue"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ae"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oe"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ä"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("Yj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ñ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ã"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("A"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("O"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("U"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
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
			},
			phoneticRules: []token{
				{
					text:  []rune("af"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("αυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("af"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("αυ"),
			phoneticRules: []token{
				{
					text:  []rune("av"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ευ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ευ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ευ"),
			phoneticRules: []token{
				{
					text:  []rune("ev"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("if"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηυ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("if"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηυ"),
			phoneticRules: []token{
				{
					text:  []rune("iv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ου"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("αι"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ει"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("οι"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ωι"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ηι"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("υι"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γγ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γγ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γγ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ng"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γγ"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(ε|ι|η)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ng"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γκ"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γι"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γι"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γε"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γε"),
			phoneticRules: []token{
				{
					text:  []rune("ge"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("κζ"),
			phoneticRules: []token{
				{
					text:  []rune("gz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("τζ"),
			phoneticRules: []token{
				{
					text:  []rune("dz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("σ"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^(β|γ|δ|μ|ν|ρ)"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μβ"),
			phoneticRules: []token{
				{
					text:  []rune("mb"),
					langs: -1,
				},
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μπ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μπ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("mb"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μπ"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ντ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ντ"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nd"),
					langs: -1,
				},
				{
					text:  []rune("nt"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ντ"),
			phoneticRules: []token{
				{
					text:  []rune("nt"),
					langs: -1,
				},
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ά"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("έ"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ή"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ί"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ό"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ύ"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ώ"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ΰ"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ϋ"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ϊ"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("α"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("β"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("γ"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("δ"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ε"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ζ"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("η"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ι"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("κ"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("λ"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("μ"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ν"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ξ"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ο"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("π"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ρ"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("σ"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ς"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("τ"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("υ"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("φ"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("θ"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("χ"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ψ"),
			phoneticRules: []token{
				{
					text:  []rune("ps"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ω"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
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
			},
			phoneticRules: []token{
				{
					text:  []rune("af"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[kpstfh]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("af"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  []rune("av"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[kpstfh]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eu"),
			phoneticRules: []token{
				{
					text:  []rune("ev"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ou"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gge"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nje"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ni"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gge"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ggi"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ng"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gk"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gke"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nje"),
					langs: -1,
				},
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gki"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ni"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gke"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gki"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gk"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ng"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gk"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghi"),
			phoneticRules: []token{
				{
					text:  []rune("Ngi"),
					langs: -1,
				},
				{
					text:  []rune("Ni"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghe"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nghe"),
			phoneticRules: []token{
				{
					text:  []rune("Nje"),
					langs: -1,
				},
				{
					text:  []rune("Nge"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ghi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ghi"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ghe"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ghe"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
				{
					text:  []rune("ge"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ngh"),
			phoneticRules: []token{
				{
					text:  []rune("Ng"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ngi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ngi"),
			phoneticRules: []token{
				{
					text:  []rune("Ngi"),
					langs: -1,
				},
				{
					text:  []rune("Ni"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nge"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nge"),
			phoneticRules: []token{
				{
					text:  []rune("Nje"),
					langs: -1,
				},
				{
					text:  []rune("Nge"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  []rune("gi"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ge"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aouy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ge"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
				{
					text:  []rune("ge"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ng"),
			phoneticRules: []token{
				{
					text:  []rune("Ng"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("dh"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("dj"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("th"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kz"),
			phoneticRules: []token{
				{
					text:  []rune("gz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("dz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bgdmnr]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mb"),
			phoneticRules: []token{
				{
					text:  []rune("mb"),
					langs: -1,
				},
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mp"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mp"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("mp"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("mp"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nt"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nt"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiouy]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nd"),
					langs: -1,
				},
				{
					text:  []rune("nt"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nt"),
			phoneticRules: []token{
				{
					text:  []rune("nt"),
					langs: -1,
				},
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("óu"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ý"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ο"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("Q"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("עי"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("עו"),
			phoneticRules: []token{
				{
					text:  []rune("VV"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("או"),
			phoneticRules: []token{
				{
					text:  []rune("VV"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ג׳"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ד׳"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("א"),
			phoneticRules: []token{
				{
					text:  []rune("L"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ב"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ג"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ד"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("1"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ה"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("וו"),
			phoneticRules: []token{
				{
					text:  []rune("V"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("וי"),
			phoneticRules: []token{
				{
					text:  []rune("WW"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ו"),
			phoneticRules: []token{
				{
					text:  []rune("W"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ז"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ח"),
			phoneticRules: []token{
				{
					text:  []rune("X"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ט"),
			phoneticRules: []token{
				{
					text:  []rune("T"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("יי"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("י"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ך"),
			phoneticRules: []token{
				{
					text:  []rune("X"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("כ"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("K"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("כ"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ל"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ם"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("מ"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ן"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("נ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ס"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ע"),
			phoneticRules: []token{
				{
					text:  []rune("L"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ף"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("פ"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ץ"),
			phoneticRules: []token{
				{
					text:  []rune("C"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("צ"),
			phoneticRules: []token{
				{
					text:  []rune("C"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ק"),
			phoneticRules: []token{
				{
					text:  []rune("K"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ר"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ש"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ת"),
			phoneticRules: []token{
				{
					text:  []rune("TB"),
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
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zs"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cs"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aj"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[áo]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[áo]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ely"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
				{
					text:  []rune("eli"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ly"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
				{
					text:  []rune("li"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gy"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
				{
					text:  []rune("gi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("ni"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ty"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ty"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("ti"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("ku"),
					langs: -1,
				},
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ő"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ű"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gli"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("gli"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gn"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: -1,
				},
				{
					text:  []rune("gn"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gni"),
			phoneticRules: []token{
				{
					text:  []rune("ni"),
					langs: -1,
				},
				{
					text:  []rune("gni"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  []rune("vo"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("�"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  []rune("dz"),
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
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ģ"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
				{
					text:  []rune("dj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ķ"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
				{
					text:  []rune("tj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ļ"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("š"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ņ"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ž"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ā"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ē"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ī"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ū"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("ej"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iu"),
			phoneticRules: []token{
				{
					text:  []rune("ju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ui"),
			phoneticRules: []token{
				{
					text:  []rune("uj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
			},
			phoneticRules: []token{
				{
					text:  []rune("ski"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cka"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tski"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lova"),
					langs: -1,
				},
				{
					text:  []rune("lof"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kowa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kova"),
					langs: -1,
				},
				{
					text:  []rune("kof"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owa"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ova"),
					langs: -1,
				},
				{
					text:  []rune("of"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lovna"),
					langs: -1,
				},
				{
					text:  []rune("levna"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kowna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("kovna"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ovna"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
				{
					text:  []rune("el"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ówna"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("czy"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cze"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: -1,
				},
				{
					text:  []rune("tSF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciewicz"),
			phoneticRules: []token{
				{
					text:  []rune("tsevitS"),
					langs: -1,
				},
				{
					text:  []rune("tSevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siewicz"),
			phoneticRules: []token{
				{
					text:  []rune("sevitS"),
					langs: -1,
				},
				{
					text:  []rune("SevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziewicz"),
			phoneticRules: []token{
				{
					text:  []rune("zevitS"),
					langs: -1,
				},
				{
					text:  []rune("ZevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("riewicz"),
			phoneticRules: []token{
				{
					text:  []rune("rjevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("diewicz"),
			phoneticRules: []token{
				{
					text:  []rune("djevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tiewicz"),
			phoneticRules: []token{
				{
					text:  []rune("tjevitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iewicz"),
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ewicz"),
			phoneticRules: []token{
				{
					text:  []rune("evitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("owicz"),
			phoneticRules: []token{
				{
					text:  []rune("ovitS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("icz"),
			phoneticRules: []token{
				{
					text:  []rune("itS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cz"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSB"),
					langs: -1,
				},
				{
					text:  []rune("tsB"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cia"),
			phoneticRules: []token{
				{
					text:  []rune("tSa"),
					langs: -1,
				},
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSom"),
					langs: -1,
				},
				{
					text:  []rune("tsom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cią"),
			phoneticRules: []token{
				{
					text:  []rune("tSon"),
					langs: -1,
				},
				{
					text:  []rune("tson"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSem"),
					langs: -1,
				},
				{
					text:  []rune("tsem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cię"),
			phoneticRules: []token{
				{
					text:  []rune("tSen"),
					langs: -1,
				},
				{
					text:  []rune("tsen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tSF"),
					langs: -1,
				},
				{
					text:  []rune("tsF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cie"),
			phoneticRules: []token{
				{
					text:  []rune("tSe"),
					langs: -1,
				},
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cio"),
			phoneticRules: []token{
				{
					text:  []rune("tSo"),
					langs: -1,
				},
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ciu"),
			phoneticRules: []token{
				{
					text:  []rune("tSu"),
					langs: -1,
				},
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
				{
					text:  []rune("tsI"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ć"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sz"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("SB"),
					langs: -1,
				},
				{
					text:  []rune("sB"),
					langs: -1,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sia"),
			phoneticRules: []token{
				{
					text:  []rune("Sa"),
					langs: -1,
				},
				{
					text:  []rune("sja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Som"),
					langs: -1,
				},
				{
					text:  []rune("som"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sią"),
			phoneticRules: []token{
				{
					text:  []rune("Son"),
					langs: -1,
				},
				{
					text:  []rune("son"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Sem"),
					langs: -1,
				},
				{
					text:  []rune("sem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("się"),
			phoneticRules: []token{
				{
					text:  []rune("Sen"),
					langs: -1,
				},
				{
					text:  []rune("sen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("SF"),
					langs: -1,
				},
				{
					text:  []rune("sF"),
					langs: -1,
				},
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  []rune("Se"),
					langs: -1,
				},
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  []rune("So"),
					langs: -1,
				},
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("siu"),
			phoneticRules: []token{
				{
					text:  []rune("Su"),
					langs: -1,
				},
				{
					text:  []rune("sju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("si"),
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: -1,
				},
				{
					text:  []rune("sI"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ś"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ZB"),
					langs: -1,
				},
				{
					text:  []rune("zB"),
					langs: -1,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zia"),
			phoneticRules: []token{
				{
					text:  []rune("Za"),
					langs: -1,
				},
				{
					text:  []rune("zja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zom"),
					langs: -1,
				},
				{
					text:  []rune("zom"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zią"),
			phoneticRules: []token{
				{
					text:  []rune("Zon"),
					langs: -1,
				},
				{
					text:  []rune("zon"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Zem"),
					langs: -1,
				},
				{
					text:  []rune("zem"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zię"),
			phoneticRules: []token{
				{
					text:  []rune("Zen"),
					langs: -1,
				},
				{
					text:  []rune("zen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ZF"),
					langs: -1,
				},
				{
					text:  []rune("zF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("Zo"),
					langs: -1,
				},
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ziu"),
			phoneticRules: []token{
				{
					text:  []rune("Zu"),
					langs: -1,
				},
				{
					text:  []rune("zju"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zi"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
				{
					text:  []rune("zI"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ZF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ZF"),
					langs: -1,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
				{
					text:  []rune("zF"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("że"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("źe"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ży"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("źi"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
				{
					text:  []rune("zi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ż"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ź"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Se"),
					langs: -1,
				},
				{
					text:  []rune("re"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rze"),
			phoneticRules: []token{
				{
					text:  []rune("Ze"),
					langs: -1,
				},
				{
					text:  []rune("re"),
					langs: -1,
				},
				{
					text:  []rune("rZe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Si"),
					langs: -1,
				},
				{
					text:  []rune("ri"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rzy"),
			phoneticRules: []token{
				{
					text:  []rune("Zi"),
					langs: -1,
				},
				{
					text:  []rune("ri"),
					langs: -1,
				},
				{
					text:  []rune("rZi"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				suffix:           []rune("t"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rz"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("r"),
					langs: -1,
				},
				{
					text:  []rune("rZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ł"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ń"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("s"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("om"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("em"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ą"),
			phoneticRules: []token{
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ę"),
			phoneticRules: []token{
				{
					text:  []rune("en"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ije"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yje"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("rie"),
			phoneticRules: []token{
				{
					text:  []rune("rje"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("die"),
			phoneticRules: []token{
				{
					text:  []rune("dje"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tie"),
			phoneticRules: []token{
				{
					text:  []rune("tje"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("F"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("au"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ei"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ey"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ej"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ai"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ay"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aj"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("B"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
				{
					text:  []rune("F"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("P"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ss"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sç"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[dglmnrv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bdgv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ptckf]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("gv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ao]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  []rune("vo"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lh"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nh"),
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ex"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("ez"),
					langs: -1,
				},
				{
					text:  []rune("eS"),
					langs: -1,
				},
				{
					text:  []rune("eks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ex"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[cs]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aáuiíoóeéê]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeiíou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bcdfglnprstv]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ão"),
			phoneticRules: []token{
				{
					text:  []rune("au"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ãe"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ãi"),
			phoneticRules: []token{
				{
					text:  []rune("aj"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("õe"),
			phoneticRules: []token{
				{
					text:  []rune("oj"),
					langs: -1,
				},
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aáuoóeéê]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("â"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ã"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
				{
					text:  []rune("an"),
					langs: -1,
				},
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ê"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ô"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("õ"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("on"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  []rune("tSe"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ci"),
			phoneticRules: []token{
				{
					text:  []rune("tSi"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gi"),
			phoneticRules: []token{
				{
					text:  []rune("dZi"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gh"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ţ"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("î"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ea"),
			phoneticRules: []token{
				{
					text:  []rune("ja"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ă"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: -1,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ina"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("in"),
					langs: -1,
				},
				{
					text:  []rune("ina"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lof"),
					langs: -1,
				},
				{
					text:  []rune("lef"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("lof"),
					langs: -1,
				},
				{
					text:  []rune("lef"),
					langs: -1,
				},
				{
					text:  []rune("lova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ova"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("of"),
					langs: -1,
				},
				{
					text:  []rune("ova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("eva"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ef"),
					langs: -1,
				},
				{
					text:  []rune("ova"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aia"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aja"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aya"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("aja"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsya"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsyu"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsia"),
			phoneticRules: []token{
				{
					text:  []rune("tsa"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsie"),
			phoneticRules: []token{
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsio"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsye"),
			phoneticRules: []token{
				{
					text:  []rune("tse"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsyo"),
			phoneticRules: []token{
				{
					text:  []rune("tso"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tsiu"),
			phoneticRules: []token{
				{
					text:  []rune("tsu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sie"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sio"),
			phoneticRules: []token{
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zie"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zio"),
			phoneticRules: []token{
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sye"),
			phoneticRules: []token{
				{
					text:  []rune("se"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("syo"),
			phoneticRules: []token{
				{
					text:  []rune("so"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zye"),
			phoneticRules: []token{
				{
					text:  []rune("ze"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zyo"),
			phoneticRules: []token{
				{
					text:  []rune("zo"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ger"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ger"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gen"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gen"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gin"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("gin"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gg"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[jaeoiuy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("kh"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("x"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sch"),
			phoneticRules: []token{
				{
					text:  []rune("StS"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ssh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("sh"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("zh"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tz"),
			phoneticRules: []token{
				{
					text:  []rune("ts"),
					langs: -1,
				},
				{
					text:  []rune("tz"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("kv"),
					langs: -1,
				},
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("s"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lya"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lyu"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lia"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("liu"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lja"),
			phoneticRules: []token{
				{
					text:  []rune("la"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lju"),
			phoneticRules: []token{
				{
					text:  []rune("lu"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("le"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("lE"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lyo"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lio"),
			phoneticRules: []token{
				{
					text:  []rune("lo"),
					langs: -1,
				},
				{
					text:  []rune("le"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ije"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yje"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yye"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yie"),
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("io"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yo"),
			phoneticRules: []token{
				{
					text:  []rune("jo"),
					langs: -1,
				},
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[au]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ii"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("iy"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yy"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yi"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("yj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ij"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ee"),
			phoneticRules: []token{
				{
					text:  []rune("aje"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("je"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("oo"),
			phoneticRules: []token{
				{
					text:  []rune("oo"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("'"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("\""),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("aue"),
			phoneticRules: []token{
				{
					text:  []rune("aue"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("E"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("I"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
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
					text:  []rune("n"),
					langs: -1,
				},
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ny"),
			phoneticRules: []token{
				{
					text:  []rune("nj"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ig"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("ig"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ix"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[aeiou]$"),
			},
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tx"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tj"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tj"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("tg"),
			phoneticRules: []token{
				{
					text:  []rune("tg"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ch"),
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("bh"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("[dgt]$"),
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[bpvf]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("gu"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("gv"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("g"),
					langs: -1,
				},
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("qu"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("uo"),
			phoneticRules: []token{
				{
					text:  []rune("vo"),
					langs: -1,
				},
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("á"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("é"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("í"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ó"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ú"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("à"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("è"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ò"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("B"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("x"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("V"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
				{
					text:  []rune("gz"),
					langs: -1,
				},
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("s"),
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
					text:  []rune("tS"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ğ"),
			phoneticRules: []token{
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ş"),
			phoneticRules: []token{
				{
					text:  []rune("S"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ü"),
			phoneticRules: []token{
				{
					text:  []rune("Q"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ö"),
			phoneticRules: []token{
				{
					text:  []rune("Y"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("ı"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
				{
					text:  []rune("i"),
					langs: -1,
				},
				{
					text:  nil,
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("a"),
			phoneticRules: []token{
				{
					text:  []rune("a"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("dZ"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("d"),
			phoneticRules: []token{
				{
					text:  []rune("d"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("e"),
			phoneticRules: []token{
				{
					text:  []rune("e"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("f"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("g"),
			phoneticRules: []token{
				{
					text:  []rune("g"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("h"),
			phoneticRules: []token{
				{
					text:  []rune("h"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("i"),
			phoneticRules: []token{
				{
					text:  []rune("i"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("k"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("l"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("n"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("o"),
			phoneticRules: []token{
				{
					text:  []rune("o"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("p"),
			phoneticRules: []token{
				{
					text:  []rune("p"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("q"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("r"),
			phoneticRules: []token{
				{
					text:  []rune("r"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("s"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("t"),
			phoneticRules: []token{
				{
					text:  []rune("t"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("u"),
			phoneticRules: []token{
				{
					text:  []rune("u"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("x"),
			phoneticRules: []token{
				{
					text:  []rune("ks"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("y"),
			phoneticRules: []token{
				{
					text:  []rune("j"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("z"),
			phoneticRules: []token{
				{
					text:  []rune("z"),
					langs: -1,
				},
			},
		},
	},
}

var genLangRules = []langRule{
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("^o’"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("^o'"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("mc"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("fitz"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ceau"),
		},
		langs:  65600,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("eau"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ault"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("oult"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("eux"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("eix"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("glou"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("uu"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("tx"),
		},
		langs:  262144,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("witz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("tz"),
		},
		langs:  131232,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("tz"),
		},
		langs:  131104,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("poulos"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("pulos"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("iou"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("sj"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("sj"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("güe"),
		},
		langs:  262144,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("güi"),
		},
		langs:  262144,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ghe"),
		},
		langs:  66048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ghi"),
		},
		langs:  66048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("escu"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("esco"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("vici"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("schi"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ii"),
		},
		langs:  131072,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("iy"),
		},
		langs:  131072,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("yy"),
		},
		langs:  131072,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("yi"),
		},
		langs:  131072,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("rz"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("rz"),
		},
		langs:  16512,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[bcdfgklmnpstwz]rz"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("rz[bcdfghklmnpstw]"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("cki"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ska"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("cka"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ae"),
		},
		langs:  131232,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("oe"),
		},
		langs:  131312,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("th"),
		},
		langs:  160,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("th"),
		},
		langs:  672,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("mann"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("cz"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("cy"),
		},
		langs:  16896,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("niew"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("etti"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("eti"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ati"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ato"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[aoei]no$"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[aoei]ni$"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("esi"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("oli"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("field"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("stein"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("heim"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("heimer"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("thal"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("zweig"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[aeou]h"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("äh"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("öh"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("üh"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[ln]h[ao]$"),
		},
		langs:  32768,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("[ln]h[aou]"),
		},
		langs:  819416,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("chsch"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("tsch"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("sch"),
		},
		langs:  131200,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("sch"),
		},
		langs:  131200,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("ck"),
		},
		langs:  160,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("c"),
		},
		langs:  608264,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sz"),
		},
		langs:  18432,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("cs"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("cs"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("dzs"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("zs"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("zs"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("wl"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("wr"),
		},
		langs:  16560,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("gy"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gy[aeou]"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gy"),
		},
		langs:  133696,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("guy"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gu[ei]"),
		},
		langs:  294976,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gu[ao]"),
		},
		langs:  294912,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("gi[aou]"),
		},
		langs:  4608,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ly"),
		},
		langs:  150016,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ny"),
		},
		langs:  412160,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ty"),
		},
		langs:  150016,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ā"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ć"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ç"),
		},
		langs:  819264,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("č"),
		},
		langs:  8200,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ď"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ē"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ğ"),
		},
		langs:  524288,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ģ"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ī"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ķ"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ļ"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ł"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ņ"),
		},
		langs:  8192,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ń"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ñ"),
		},
		langs:  262144,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ň"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ř"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ś"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ş"),
		},
		langs:  589824,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("š"),
		},
		langs:  8200,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ţ"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ť"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ź"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ž"),
		},
		langs:  8200,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ż"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ß"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ä"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("á"),
		},
		langs:  297480,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("â"),
		},
		langs:  98368,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ă"),
		},
		langs:  65536,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ą"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("à"),
		},
		langs:  32768,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ã"),
		},
		langs:  32768,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ę"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("é"),
		},
		langs:  2632,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("è"),
		},
		langs:  266304,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ê"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ě"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ê"),
		},
		langs:  32832,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("í"),
		},
		langs:  297480,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("î"),
		},
		langs:  65600,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ı"),
		},
		langs:  524288,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ó"),
		},
		langs:  317960,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ö"),
		},
		langs:  526464,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ô"),
		},
		langs:  32832,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("õ"),
		},
		langs:  34816,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ò"),
		},
		langs:  266240,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ű"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ú"),
		},
		langs:  297480,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ü"),
		},
		langs:  821376,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ù"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ů"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ý"),
		},
		langs:  520,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("а"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ё"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("о"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("е"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("и"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("у"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ы"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("э"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ю"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("я"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("α"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ε"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("η"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ι"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ο"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("υ"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ω"),
		},
		langs:  256,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ا"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ب"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ت"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ث"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ج"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ح"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("خ'"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("د"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ذ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ر"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ز"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("س"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ش"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ص"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ض"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ط"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ظ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ع"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("غ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ف"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ق"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ك"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ل"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("م"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ن"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ه"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("و"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ي"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("آ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("إ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("أ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ؤ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ئ"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("א"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ב"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ג"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ד"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ה"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ו"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ז"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ח"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ט"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("י"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("כ"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ל"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("מ"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("נ"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ס"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ע"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("פ"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("צ"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ק"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ר"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ש"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ת"),
		},
		langs:  1024,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("a"),
		},
		langs:  1286,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("o"),
		},
		langs:  1286,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("e"),
		},
		langs:  1286,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("i"),
		},
		langs:  1286,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("y"),
		},
		langs:  75030,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("u"),
		},
		langs:  1286,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("j"),
		},
		langs:  4096,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("j[^aoeiuy]"),
		},
		langs:  295488,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("g"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("k"),
		},
		langs:  364608,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("q"),
		},
		langs:  748056,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("v"),
		},
		langs:  16384,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("w"),
		},
		langs:  993864,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("x"),
		},
		langs:  534552,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("dj"),
		},
		langs:  786432,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("v[^aoeiu]"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("c[^aohk]"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("dzi"),
		},
		langs:  524512,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ou"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("a[eiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("ö[eaiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("ü[eaiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("e[aiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("i[aeou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("o[aieu]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("u[aieo]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("aj"),
		},
		langs:  240,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ej"),
		},
		langs:  240,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("oj"),
		},
		langs:  240,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("uj"),
		},
		langs:  240,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("eu"),
		},
		langs:  147456,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ky"),
		},
		langs:  16384,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("kie"),
		},
		langs:  262720,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gie"),
		},
		langs:  360960,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("ch[aou]"),
		},
		langs:  4096,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ch"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			suffix:           []rune("son"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			pattern:          regexp.MustCompile("sc[ei]"),
		},
		langs:  64,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sch"),
		},
		langs:  280640,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			prefix:           []rune("h"),
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
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("b"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("v"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("g"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("d"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("dZ"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("tS"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jnm"),
				phoneticRules: []token{
					{
						text:  []rune("jm"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("I"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[aA]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					suffix:           []rune("A"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("A"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("A"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("j"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("j"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("l"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("l"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("m"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("m"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("n"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("r"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("r"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("z"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("vanden"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("vanden"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("vander"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("vander"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("van"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bp]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("vam"),
						langs: -1,
					},
					{
						text:  nil,
						langs: 16,
					},
				},
			},
			{
				pattern: []rune("van"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("van"),
						langs: -1,
					},
					{
						text:  nil,
						langs: 16,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bp]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("m"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("h"),
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("H"),
				phoneticRules: []token{
					{
						text:  []rune("x"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("sen"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("zn"),
						langs: -1,
					},
					{
						text:  []rune("zon"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("sen"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("sn"),
						langs: -1,
					},
					{
						text:  []rune("son"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("sEn"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("zn"),
						langs: -1,
					},
					{
						text:  []rune("zon"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("sEn"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("sn"),
						langs: -1,
					},
					{
						text:  []rune("son"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("e"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("i"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("E"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("I"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Q"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Y"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("e"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("e"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("i"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("E"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("E"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("I"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("I"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Q"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("Q"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Y"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("Y"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEs"),
				phoneticRules: []token{
					{
						text:  []rune("lEs"),
						langs: -1,
					},
					{
						text:  []rune("lz"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lE"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("lE"),
						langs: -1,
					},
					{
						text:  []rune("l"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aue"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oue"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AvE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("AvE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ave"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("Ave"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("avE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("avE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ave"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ave"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OvE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("OvE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ove"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("Ove"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ovE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ovE"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ove"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ove"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ea"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("ea"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("EA"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ea"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("Ea"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
					{
						text:  []rune("eA"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uje"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("iji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ije"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Iji"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjI"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ije"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjE"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ajo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ojO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Oja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ojo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("OjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("eja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ejO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Eja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ejo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("EjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ujO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Uja"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ujo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("UjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ija"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ijO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ija"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjA"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Ijo"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("IjO"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Aju"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("AjU"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("j"),
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lYndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lander"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lAndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lAnder"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("landEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lender"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEndEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lendEr"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("lEnder"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lYnder"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("burk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bUrk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("burg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bUrg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Burk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("BUrk"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Burg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("BUrg"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("burk"),
						langs: -1,
					},
					{
						text:  []rune("berk"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dS"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dZ"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dZ"),
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				phoneticRules: []token{
					{
						text:  []rune("z"),
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
							text:  []rune("mb"),
							langs: -1,
						},
						{
							text:  []rune("b"),
							langs: 512,
						},
					},
				},
				{
					pattern: []rune("mp"),
					phoneticRules: []token{
						{
							text:  []rune("mp"),
							langs: -1,
						},
						{
							text:  []rune("b"),
							langs: 512,
						},
					},
				},
				{
					pattern: []rune("ng"),
					phoneticRules: []token{
						{
							text:  []rune("ng"),
							langs: -1,
						},
						{
							text:  []rune("g"),
							langs: 512,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fktSs]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
						{
							text:  []rune("f"),
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						prefix:           []rune("p"),
					},
					phoneticRules: []token{
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
						{
							text:  []rune("f"),
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[pktSs]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("f"),
							langs: -1,
						},
						{
							text:  []rune("p"),
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						prefix:           []rune("f"),
					},
					phoneticRules: []token{
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("f"),
							langs: -1,
						},
						{
							text:  []rune("p"),
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("b"),
							langs: -1,
						},
						{
							text:  []rune("v"),
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  []rune("v"),
							langs: -1,
						},
						{
							text:  []rune("b"),
							langs: 262144,
						},
					},
				},
				{
					pattern: []rune("t"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("t"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("g"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						suffix:           []rune("n"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("g"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("k"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						suffix:           []rune("n"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("k"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("p"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("r"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[Ee]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("r"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("s"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("s"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("t"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiouAEIOU]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("t"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("s"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiouAEIOU]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("s"),
							langs: -1,
						},
						{
							text:  nil,
							langs: 64,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: 128,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("D"),
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: 128,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lEE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("li"),
							langs: -1,
						},
						{
							text:  []rune("il"),
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("rEE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("ri"),
							langs: -1,
						},
						{
							text:  []rune("ir"),
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("li"),
							langs: -1,
						},
						{
							text:  []rune("il"),
							langs: 32,
						},
						{
							text:  []rune("lY"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("rE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("ri"),
							langs: -1,
						},
						{
							text:  []rune("ir"),
							langs: 32,
						},
						{
							text:  []rune("rY"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("EE"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ea"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eu"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("e"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ei"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ei"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iA"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ia"),
							langs: -1,
						},
						{
							text:  []rune("io"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iA"),
					phoneticRules: []token{
						{
							text:  []rune("ia"),
							langs: -1,
						},
						{
							text:  []rune("io"),
							langs: -1,
						},
						{
							text:  []rune("iY"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
						{
							text:  []rune("D"),
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
						{
							text:  nil,
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
						{
							text:  nil,
							langs: 32,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DoiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sUts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("suts"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("Uts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: 128,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("1i"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("1u"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("j1"),
					phoneticRules: []token{
						{
							text:  []rune("ja"),
							langs: -1,
						},
						{
							text:  []rune("je"),
							langs: -1,
						},
						{
							text:  []rune("jo"),
							langs: -1,
						},
						{
							text:  []rune("ju"),
							langs: -1,
						},
						{
							text:  []rune("j"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("1"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("e"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("u"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("i"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("p"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("p"),
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
						{
							text:  []rune("b"),
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
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("om"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("on"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("en"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("Ym"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("En"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("Yn"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
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
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("om"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("on"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("en"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Em"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("Ym"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("En"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("Yn"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
						pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("D"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("lE"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("il"),
							langs: -1,
						},
						{
							text:  []rune("li"),
							langs: -1,
						},
						{
							text:  []rune("lY"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiEuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
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
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiAEIOUouQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("Q"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("AU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("OU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oU"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoAOUiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoAOUeiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aoeOUiuQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("Y"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuUQY]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uk"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sUts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("suts"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Uts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("uts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("N"),
					phoneticRules: []token{
						{
							text:  nil,
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiB"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dm"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("aiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("uiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("eiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("EiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("iiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("IiF"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("Dn"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("im"),
							langs: -1,
						},
						{
							text:  []rune("om"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("in"),
							langs: -1,
						},
						{
							text:  []rune("on"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
						{
							text:  []rune("Qk"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Ik"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("ik"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("sIts"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("sits"),
							langs: -1,
						},
						{
							text:  []rune("sQts"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("Its"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("its"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
						{
							text:  []rune("Q"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("au"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("e"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^ts$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("Y"),
							langs: -1,
						},
						{
							text:  []rune("i"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("b"),
							langs: -1,
						},
						{
							text:  []rune("v"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  []rune("b"),
							langs: -1,
						},
						{
							text:  []rune("v"),
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
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ou"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ai"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("oi"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("o"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("ui"),
					phoneticRules: []token{
						{
							text:  []rune("D"),
							langs: -1,
						},
						{
							text:  []rune("u"),
							langs: -1,
						},
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("a"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
						{
							text:  []rune("o"),
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
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("p"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("b"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("f"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("v"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("g"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("d"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("dZ"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("tS"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jnm"),
				phoneticRules: []token{
					{
						text:  []rune("jm"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("I"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[aA]"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					suffix:           []rune("A"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("A"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("A"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("b"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("b"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("d"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("d"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("f"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("f"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("g"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("g"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("j"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("j"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("k"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("l"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("l"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("m"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("m"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("n"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("n"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("p"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("p"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("r"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("r"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("t"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("t"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("v"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("v"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("z"),
				},
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("H"),
				phoneticRules: []token{
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("s"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[^t]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgZd]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkst]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("Z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("S"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("S"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[bgzd]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("Z"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("z"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jI"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("je"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jE"),
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phoneticRules: []token{
					{
						text:  []rune("j"),
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
					},
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fktSs]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						prefix:           []rune("p"),
					},
					phoneticRules: []token{
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[pktSs]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("f"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						prefix:           []rune("f"),
					},
					phoneticRules: []token{
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("f"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("b"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  []rune("v"),
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
							text:  nil,
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
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
					},
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("A"),
					phoneticRules: []token{
						{
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("O"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("U"),
					phoneticRules: []token{
						{
							text:  []rune("u"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[fktSs]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						prefix:           []rune("p"),
					},
					phoneticRules: []token{
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("p"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						pattern:          regexp.MustCompile("^[pktSs]"),
					},
					phoneticRules: []token{
						{
							text:  []rune("f"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						prefix:           []rune("f"),
					},
					phoneticRules: []token{
						{
							text:  nil,
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					rightContext: &ruleMatcher{
						matchEmptyString: true,
					},
					phoneticRules: []token{
						{
							text:  []rune("f"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("B"),
					phoneticRules: []token{
						{
							text:  []rune("b"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  []rune("v"),
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
							text:  []rune("n"),
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
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("a"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("F"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("P"),
					phoneticRules: []token{
						{
							text:  []rune("o"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("E"),
					phoneticRules: []token{
						{
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("e"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("I"),
					phoneticRules: []token{
						{
							text:  []rune("i"),
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
							text:  []rune("b"),
							langs: -1,
						},
					},
				},
				{
					pattern: []rune("V"),
					phoneticRules: []token{
						{
							text:  []rune("v"),
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
