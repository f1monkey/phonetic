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
			pattern: []rune("ph"),
			phoneticRules: []token{
				{
					text:  []rune("f"),
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
					text:  []rune("gli"),
					langs: -1,
				},
				{
					text:  []rune("l"),
					langs: 8,
				},
			},
		},
		{
			pattern: []rune("gni"),
			phoneticRules: []token{
				{
					text:  []rune("gni"),
					langs: -1,
				},
				{
					text:  []rune("ni"),
					langs: 10,
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
					langs: 10,
				},
				{
					text:  []rune("nj"),
					langs: 10,
				},
				{
					text:  []rune("gn"),
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
					text:  []rune("gh"),
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
				{
					text:  []rune("dh"),
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
				{
					text:  []rune("bh"),
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
				{
					text:  []rune("th"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("lh"),
			phoneticRules: []token{
				{
					text:  []rune("l"),
					langs: 16,
				},
				{
					text:  []rune("lh"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("nh"),
			phoneticRules: []token{
				{
					text:  []rune("n"),
					langs: 16,
				},
				{
					text:  []rune("nh"),
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
					text:  []rune("ig"),
					langs: -1,
				},
				{
					text:  []rune("tS"),
					langs: 32,
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
					langs: 32,
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
			pattern: []rune("g"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				prefix:           []rune("y"),
			},
			phoneticRules: []token{
				{
					text:  []rune("Z"),
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
					text:  []rune("gZ"),
					langs: 18,
				},
				{
					text:  []rune("dZ"),
					langs: 40,
				},
				{
					text:  []rune("x"),
					langs: 32,
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
					langs: 18,
				},
				{
					text:  []rune("dZ"),
					langs: 40,
				},
				{
					text:  []rune("x"),
					langs: 32,
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
			pattern: []rune("gue"),
			rightContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("k"),
					langs: 2,
				},
				{
					text:  []rune("ge"),
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
				{
					text:  []rune("S"),
					langs: 8,
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
			pattern: []rune("ç"),
			phoneticRules: []token{
				{
					text:  []rune("s"),
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
					langs: 8,
				},
				{
					text:  []rune("S"),
					langs: 18,
				},
				{
					text:  []rune("tS"),
					langs: 32,
				},
				{
					text:  []rune("dZ"),
					langs: 32,
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
				{
					text:  []rune("tS"),
					langs: 32,
				},
				{
					text:  []rune("dZ"),
					langs: 32,
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
					langs: 8,
				},
				{
					text:  []rune("si"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("cc"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[eiyéèê]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("tS"),
					langs: 8,
				},
				{
					text:  []rune("ks"),
					langs: 50,
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
					text:  []rune("tS"),
					langs: 8,
				},
				{
					text:  []rune("s"),
					langs: 50,
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
					text:  []rune("s"),
					langs: 32,
				},
				{
					text:  []rune("z"),
					langs: 26,
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
					text:  []rune("z"),
					langs: -1,
				},
				{
					text:  []rune("s"),
					langs: -1,
				},
				{
					text:  []rune("Z"),
					langs: 16,
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
					langs: 8,
				},
				{
					text:  []rune("S"),
					langs: 16,
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
					langs: 8,
				},
				{
					text:  []rune("Z"),
					langs: 16,
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
					langs: 8,
				},
				{
					text:  []rune("S"),
					langs: 16,
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
					text:  []rune("dz"),
					langs: 8,
				},
				{
					text:  []rune("ts"),
					langs: 8,
				},
				{
					text:  []rune("s"),
					langs: 32,
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
					langs: 2,
				},
				{
					text:  []rune("ke"),
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
				{
					text:  []rune("k"),
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
					langs: 16,
				},
				{
					text:  []rune("eS"),
					langs: 16,
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
					langs: 16,
				},
				{
					text:  []rune("ek"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("m"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[cdglnrst]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("m"),
					langs: -1,
				},
				{
					text:  []rune("n"),
					langs: 16,
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
					langs: 48,
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
					langs: 16,
				},
			},
		},
		{
			pattern: []rune("b"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
				{
					text:  []rune("V"),
					langs: 32,
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
					text:  []rune("B"),
					langs: 32,
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
			pattern: []rune("ouh"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: 2,
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
			pattern: []rune("ou"),
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
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
				pattern:          regexp.MustCompile("^[aie]"),
			},
			phoneticRules: []token{
				{
					text:  []rune("v"),
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
					langs: 2,
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
					text:  []rune("ou"),
					langs: -1,
				},
				{
					text:  []rune("u"),
					langs: 2,
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
			pattern: []rune("ò"),
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
			pattern: []rune("ü"),
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
				{
					text:  []rune("v"),
					langs: 32,
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
					langs: 32,
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
				{
					text:  []rune("S"),
					langs: 16,
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
				{
					text:  []rune("b"),
					langs: 32,
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
					langs: 48,
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
	sepfrench: rules{
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
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
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
			pattern: []rune("q"),
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
			pattern: []rune("j"),
			phoneticRules: []token{
				{
					text:  []rune("Z"),
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
			pattern: []rune("ouh"),
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
			pattern: []rune("ey"),
			phoneticRules: []token{
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
			pattern: []rune("y"),
			phoneticRules: []token{
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
	sephebrew: rules{
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
					text:  []rune("T"),
					langs: -1,
				},
			},
		},
	},
	sepitalian: rules{
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
	sepportuguese: rules{
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
	sepspanish: rules{
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
			pattern: []rune("w"),
			phoneticRules: []token{
				{
					text:  []rune("v"),
					langs: -1,
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
					text:  []rune("B"),
					langs: -1,
				},
				{
					text:  []rune("v"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("b"),
			leftContext: &ruleMatcher{
				matchEmptyString: true,
			},
			phoneticRules: []token{
				{
					text:  []rune("b"),
					langs: -1,
				},
				{
					text:  []rune("V"),
					langs: -1,
				},
			},
		},
		{
			pattern: []rune("v"),
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
			pattern: []rune("c"),
			phoneticRules: []token{
				{
					text:  []rune("k"),
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
			pattern: []rune("q"),
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
	},
}

var sepLangRules = []langRule{
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("eau"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ou"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gni"),
		},
		langs:  10,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("tx"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("tj"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gy"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("guy"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sh"),
		},
		langs:  48,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("lh"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("nh"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ny"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gue"),
		},
		langs:  34,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gui"),
		},
		langs:  34,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gia"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gie"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gio"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("giu"),
		},
		langs:  8,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ñ"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("â"),
		},
		langs:  18,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("á"),
		},
		langs:  48,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("à"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ã"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ê"),
		},
		langs:  18,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("í"),
		},
		langs:  48,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("î"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ô"),
		},
		langs:  18,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("õ"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ò"),
		},
		langs:  40,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ú"),
		},
		langs:  48,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ù"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ü"),
		},
		langs:  48,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("א"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ב"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ג"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ד"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ה"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ו"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ז"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ח"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ט"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("י"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("כ"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ל"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("מ"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("נ"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ס"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ע"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("פ"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("צ"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ק"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ר"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ש"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ת"),
		},
		langs:  4,
		accept: true,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("a"),
		},
		langs:  4,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("o"),
		},
		langs:  4,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("e"),
		},
		langs:  4,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("i"),
		},
		langs:  4,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("y"),
		},
		langs:  4,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("u"),
		},
		langs:  4,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("kh"),
		},
		langs:  32,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("gua"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("guo"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ç"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("cha"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("cho"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("chu"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("j"),
		},
		langs:  8,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("dj"),
		},
		langs:  32,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sce"),
		},
		langs:  2,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("sci"),
		},
		langs:  2,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("ó"),
		},
		langs:  2,
		accept: false,
	},
	{
		match: &ruleMatcher{
			matchEmptyString: false,
			contains:         []rune("è"),
		},
		langs:  16,
		accept: false,
	},
}

var sepFinalRules = finalRules{
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
				pattern: []rune("nm"),
				phoneticRules: []token{
					{
						text:  []rune("m"),
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
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("a"),
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
				pattern: []rune("e"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("e"),
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
				pattern: []rune("i"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("i"),
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
				pattern: []rune("o"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("o"),
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
				pattern: []rune("u"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("u"),
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
				pattern: []rune("mbr"),
				phoneticRules: []token{
					{
						text:  []rune("mr"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("mpr"),
				phoneticRules: []token{
					{
						text:  []rune("mr"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bens"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("binz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("benS"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("binz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ben"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bin"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bar"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bar"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("abens"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("binz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("abenS"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("binz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aben"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bin"),
						langs: -1,
					},
					{
						text:  []rune("bun"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("abe"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bi"),
						langs: -1,
					},
					{
						text:  []rune("bu"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("abi"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bi"),
						langs: -1,
					},
					{
						text:  []rune("bu"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("abou"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bu"),
						langs: -1,
					},
					{
						text:  nil,
						langs: 2,
					},
				},
			},
			{
				pattern: []rune("abu"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bu"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("bou"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bu"),
						langs: -1,
					},
					{
						text:  nil,
						langs: 2,
					},
				},
			},
			{
				pattern: []rune("bu"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("bu"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ibn"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("ibn"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("els"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("ilz"),
						langs: -1,
					},
					{
						text:  []rune("lz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("elS"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("ilz"),
						langs: -1,
					},
					{
						text:  []rune("lz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("el"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("il"),
						langs: -1,
					},
					{
						text:  []rune("l"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("als"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("alS"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("lz"),
						langs: -1,
					},
					{
						text:  []rune("s"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("al"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("l"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dal"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("dal"),
						langs: -1,
					},
					{
						text:  nil,
						langs: 8,
					},
				},
			},
			{
				pattern: []rune("da"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("da"),
						langs: -1,
					},
					{
						text:  []rune("a"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("della"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("dila"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("dela"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("dila"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("del"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("dil"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("des"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("dis"),
						langs: -1,
					},
					{
						text:  nil,
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("de"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("di"),
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
				pattern: []rune("di"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("di"),
						langs: -1,
					},
					{
						text:  []rune("i"),
						langs: -1,
					},
					{
						text:  nil,
						langs: 8,
					},
				},
			},
			{
				pattern: []rune("do"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("du"),
						langs: -1,
					},
					{
						text:  []rune("u"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("du"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("du"),
						langs: -1,
					},
					{
						text:  []rune("u"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oa"),
				phoneticRules: []token{
					{
						text:  []rune("va"),
						langs: -1,
					},
					{
						text:  []rune("a"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("oe"),
				phoneticRules: []token{
					{
						text:  []rune("vi"),
						langs: -1,
					},
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ae"),
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
				pattern: []rune("ha"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("ha"),
						langs: -1,
					},
					{
						text:  []rune("a"),
						langs: -1,
					},
					{
						text:  nil,
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
					{
						text:  []rune("h"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("x"),
				phoneticRules: []token{
					{
						text:  []rune("h"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("k"),
				phoneticRules: []token{
					{
						text:  []rune("h"),
						langs: -1,
					},
					{
						text:  []rune("k"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aja"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("Da"),
						langs: -1,
					},
					{
						text:  []rune("ia"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aje"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("Di"),
						langs: -1,
					},
					{
						text:  []rune("Da"),
						langs: -1,
					},
					{
						text:  []rune("i"),
						langs: -1,
					},
					{
						text:  []rune("ia"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aji"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("Di"),
						langs: -1,
					},
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ajo"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("Du"),
						langs: -1,
					},
					{
						text:  []rune("Da"),
						langs: -1,
					},
					{
						text:  []rune("iu"),
						langs: -1,
					},
					{
						text:  []rune("ia"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aju"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("Du"),
						langs: -1,
					},
					{
						text:  []rune("iu"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("aj"),
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
				pattern: []rune("ej"),
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
				pattern: []rune("oj"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("uj"),
				phoneticRules: []token{
					{
						text:  []rune("D"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("au"),
				phoneticRules: []token{
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
						text:  []rune("iu"),
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
				pattern: []rune("ou"),
				phoneticRules: []token{
					{
						text:  []rune("u"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("a"),
				leftContext: &ruleMatcher{
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
				pattern: []rune("ja"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("ia"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("je"),
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
				pattern: []rune("jo"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("iu"),
						langs: -1,
					},
					{
						text:  []rune("ia"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ju"),
				leftContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("iu"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ja"),
				phoneticRules: []token{
					{
						text:  []rune("a"),
						langs: -1,
					},
					{
						text:  []rune("ia"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("je"),
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ji"),
				phoneticRules: []token{
					{
						text:  []rune("i"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("jo"),
				phoneticRules: []token{
					{
						text:  []rune("u"),
						langs: -1,
					},
					{
						text:  []rune("iu"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ju"),
				phoneticRules: []token{
					{
						text:  []rune("u"),
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
				pattern: []rune("i"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
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
				pattern: []rune("o"),
				rightContext: &ruleMatcher{
					matchEmptyString: true,
				},
				phoneticRules: []token{
					{
						text:  []rune("a"),
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
				pattern: []rune("o"),
				phoneticRules: []token{
					{
						text:  []rune("u"),
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
				pattern: []rune("se"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
					{
						text:  []rune("si"),
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
				pattern: []rune("Se"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("z"),
						langs: -1,
					},
					{
						text:  []rune("si"),
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
			{
				pattern: []rune("be"),
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
						text:  []rune("bi"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("pe"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("b"),
						langs: -1,
					},
					{
						text:  []rune("pi"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ve"),
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
						text:  []rune("vi"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("fe"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbgdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("v"),
						langs: -1,
					},
					{
						text:  []rune("fi"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ge"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("k"),
						langs: -1,
					},
					{
						text:  []rune("gi"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ke"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbdZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("g"),
						langs: -1,
					},
					{
						text:  []rune("ki"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("de"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("t"),
						langs: -1,
					},
					{
						text:  []rune("di"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("te"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[vbgZz]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("d"),
						langs: -1,
					},
					{
						text:  []rune("ti"),
						langs: -1,
					},
				},
			},
			{
				pattern: []rune("ze"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phoneticRules: []token{
					{
						text:  []rune("s"),
						langs: -1,
					},
					{
						text:  []rune("zi"),
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
						text:  nil,
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
			{
				pattern: []rune("p"),
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
				pattern: []rune("nm"),
				phoneticRules: []token{
					{
						text:  []rune("m"),
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
				pattern: []rune("a"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("a"),
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
				pattern: []rune("e"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("e"),
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
				pattern: []rune("i"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("i"),
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
				pattern: []rune("o"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("o"),
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
				pattern: []rune("u"),
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					prefix:           []rune("u"),
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
				pattern: []rune("h"),
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
