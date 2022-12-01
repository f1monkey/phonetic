// GENERATED CODE. DO NOT EDIT!
package bmpm

type sepLang int

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

var sepRules = map[sepLang][]rule{
	sepany: []rule{
		{
			patterns: []string{
				"ph",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"sh",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"gli",
				"",
				"",
				"(gli|l[8])",
			},
		},
		{
			patterns: []string{
				"gni",
				"",
				"",
				"(gni|ni[10])",
			},
		},
		{
			patterns: []string{
				"gn",
				"",
				"[aeou]",
				"(n[10]|nj[10]|gn)",
			},
		},
		{
			patterns: []string{
				"gh",
				"",
				"",
				"(g|gh)",
			},
		},
		{
			patterns: []string{
				"dh",
				"",
				"",
				"(d|dh)",
			},
		},
		{
			patterns: []string{
				"bh",
				"",
				"",
				"(b|bh)",
			},
		},
		{
			patterns: []string{
				"th",
				"",
				"",
				"(t|th)",
			},
		},
		{
			patterns: []string{
				"lh",
				"",
				"",
				"(l[16]|lh)",
			},
		},
		{
			patterns: []string{
				"nh",
				"",
				"",
				"(n[16]|nh)",
			},
		},
		{
			patterns: []string{
				"ig",
				"[aeiou]",
				"",
				"(ig|tS[32])",
			},
		},
		{
			patterns: []string{
				"ix",
				"[aeiou]",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"tx",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"tj",
				"",
				"$",
				"tS",
			},
		},
		{
			patterns: []string{
				"tj",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: []string{
				"tg",
				"",
				"",
				"(tg|dZ[32])",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"[aeou]",
				"dZ",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"y",
				"Z",
			},
		},
		{
			patterns: []string{
				"gg",
				"",
				"[ei]",
				"(gZ[18]|dZ[40]|x[32])",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[ei]",
				"(Z[18]|dZ[40]|x[32])",
			},
		},
		{
			patterns: []string{
				"guy",
				"",
				"",
				"gi",
			},
		},
		{
			patterns: []string{
				"gue",
				"",
				"$",
				"(k[2]|ge)",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[ei]",
				"(g|gv)",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[ao]",
				"gv",
			},
		},
		{
			patterns: []string{
				"ñ",
				"",
				"",
				"(n|nj)",
			},
		},
		{
			patterns: []string{
				"ny",
				"",
				"",
				"nj",
			},
		},
		{
			patterns: []string{
				"sc",
				"",
				"[ei]",
				"(s|S[8])",
			},
		},
		{
			patterns: []string{
				"sç",
				"",
				"[aeiou]",
				"s",
			},
		},
		{
			patterns: []string{
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"[ei]",
				"(k[8]|S[18]|tS[32]|dZ[32])",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"(S|tS[32]|dZ[32])",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"[aeou]",
				"(tS[8]|si)",
			},
		},
		{
			patterns: []string{
				"cc",
				"",
				"[eiyéèê]",
				"(tS[8]|ks[50])",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[eiyéèê]",
				"(tS[8]|s[50])",
			},
		},
		{
			patterns: []string{
				"s",
				"^",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"s",
				"[aáuiíoóeéêy]",
				"[aáuiíoóeéêy]",
				"(s[32]|z[26])",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"[dglmnrv]",
				"(z|s|Z[16])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"$",
				"(s|ts[8]|S[16])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"[bdgv]",
				"(z|dz[8]|Z[16])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"[ptckf]",
				"(s|ts[8]|S[16])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"",
				"(z|dz[8]|ts[8]|s[32])",
			},
		},
		{
			patterns: []string{
				"que",
				"",
				"$",
				"(k[2]|ke)",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"[eiu]",
				"k",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"[ao]",
				"(kv|k)",
			},
		},
		{
			patterns: []string{
				"ex",
				"",
				"[aáuiíoóeéêy]",
				"(ez[16]|eS[16]|eks|egz)",
			},
		},
		{
			patterns: []string{
				"ex",
				"",
				"[cs]",
				"(e[16]|ek)",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"[cdglnrst]",
				"(m|n[16])",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"[bfpv]",
				"(m|n[48])",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"$",
				"(m|n[16])",
			},
		},
		{
			patterns: []string{
				"b",
				"^",
				"",
				"(b|V[32])",
			},
		},
		{
			patterns: []string{
				"v",
				"^",
				"",
				"(v|B[32])",
			},
		},
		{
			patterns: []string{
				"eau",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ouh",
				"",
				"[aioe]",
				"(v[2]|uh)",
			},
		},
		{
			patterns: []string{
				"uh",
				"",
				"[aioe]",
				"(v|uh)",
			},
		},
		{
			patterns: []string{
				"ou",
				"",
				"[aioe]",
				"v",
			},
		},
		{
			patterns: []string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"[aie]",
				"v",
			},
		},
		{
			patterns: []string{
				"i",
				"[aáuoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"[aáuiíoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"[aeiíou]",
				"j",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"$",
				"(e|[2])",
			},
		},
		{
			patterns: []string{
				"ão",
				"",
				"",
				"(au|an)",
			},
		},
		{
			patterns: []string{
				"ãe",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: []string{
				"ãi",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: []string{
				"õe",
				"",
				"",
				"(oj|on)",
			},
		},
		{
			patterns: []string{
				"où",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ou",
				"",
				"",
				"(ou|u[2])",
			},
		},
		{
			patterns: []string{
				"â",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"á",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"ã",
				"",
				"",
				"(a|an)",
			},
		},
		{
			patterns: []string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"è",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"í",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"î",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ô",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ó",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"õ",
				"",
				"",
				"(o|on)",
			},
		},
		{
			patterns: []string{
				"ò",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ú",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ü",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"b",
				"",
				"",
				"(b|v[32])",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"(x[32]|Z)",
			},
		},
		{
			patterns: []string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"",
				"(s|S[16])",
			},
		},
		{
			patterns: []string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"v",
				"",
				"",
				"(v|b[32])",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"x",
				"",
				"",
				"(ks|gz|S[48])",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"",
				"z",
			},
		},
	},
	sepfrench: []rule{
		{
			patterns: []string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"ph",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"x",
				"",
				"",
				"ks",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[eiyéèê]",
				"s",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"gn",
				"",
				"",
				"(n|gn)",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[eiy]",
				"Z",
			},
		},
		{
			patterns: []string{
				"gue",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[eiy]",
				"g",
			},
		},
		{
			patterns: []string{
				"que",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"s",
				"[aeiouyéèê]",
				"[aeiouyéèê]",
				"z",
			},
		},
		{
			patterns: []string{
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"h",
				"[bdgt]",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"$",
				"",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"ouh",
				"",
				"[aioe]",
				"(v|uh)",
			},
		},
		{
			patterns: []string{
				"ou",
				"",
				"[aeio]",
				"v",
			},
		},
		{
			patterns: []string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"[aeio]",
				"v",
			},
		},
		{
			patterns: []string{
				"aue",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: []string{
				"eau",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ai",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ay",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"è",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"â",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"où",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ou",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"oi",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: []string{
				"y",
				"[ou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"$",
				"(e|)",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"[aoeu]",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"b",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"v",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"",
				"z",
			},
		},
	},
	sephebrew: []rule{
		{
			patterns: []string{
				"אי",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"עי",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"עו",
				"",
				"",
				"VV",
			},
		},
		{
			patterns: []string{
				"או",
				"",
				"",
				"VV",
			},
		},
		{
			patterns: []string{
				"ג׳",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"ד׳",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: []string{
				"א",
				"",
				"",
				"L",
			},
		},
		{
			patterns: []string{
				"ב",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"ג",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"ד",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"ה",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: []string{
				"ה",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: []string{
				"ה",
				"",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"וו",
				"",
				"",
				"V",
			},
		},
		{
			patterns: []string{
				"וי",
				"",
				"",
				"WW",
			},
		},
		{
			patterns: []string{
				"ו",
				"",
				"",
				"W",
			},
		},
		{
			patterns: []string{
				"ז",
				"",
				"",
				"z",
			},
		},
		{
			patterns: []string{
				"ח",
				"",
				"",
				"X",
			},
		},
		{
			patterns: []string{
				"ט",
				"",
				"",
				"T",
			},
		},
		{
			patterns: []string{
				"יי",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"י",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ך",
				"",
				"",
				"X",
			},
		},
		{
			patterns: []string{
				"כ",
				"^",
				"",
				"K",
			},
		},
		{
			patterns: []string{
				"כ",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"ל",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"ם",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"מ",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"ן",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"נ",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"ס",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ע",
				"",
				"",
				"L",
			},
		},
		{
			patterns: []string{
				"ף",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"פ",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"ץ",
				"",
				"",
				"C",
			},
		},
		{
			patterns: []string{
				"צ",
				"",
				"",
				"C",
			},
		},
		{
			patterns: []string{
				"ק",
				"",
				"",
				"K",
			},
		},
		{
			patterns: []string{
				"ר",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"ש",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ת",
				"",
				"",
				"T",
			},
		},
	},
	sepitalian: []rule{
		{
			patterns: []string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"gli",
				"",
				"",
				"(l|gli)",
			},
		},
		{
			patterns: []string{
				"gn",
				"",
				"[aeou]",
				"(n|nj|gn)",
			},
		},
		{
			patterns: []string{
				"gni",
				"",
				"",
				"(ni|gni)",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"[aeou]",
				"dZ",
			},
		},
		{
			patterns: []string{
				"gg",
				"",
				"[ei]",
				"dZ",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[ei]",
				"dZ",
			},
		},
		{
			patterns: []string{
				"h",
				"[bdgt]",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"[aeou]",
				"tS",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"[ei]",
				"k",
			},
		},
		{
			patterns: []string{
				"sc",
				"",
				"[ei]",
				"S",
			},
		},
		{
			patterns: []string{
				"cc",
				"",
				"[ei]",
				"tS",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[ei]",
				"tS",
			},
		},
		{
			patterns: []string{
				"s",
				"[aeiou]",
				"[aeiou]",
				"z",
			},
		},
		{
			patterns: []string{
				"i",
				"[aeou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"[aeou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"[aei]",
				"v",
			},
		},
		{
			patterns: []string{
				"�",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"�",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"�",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"�",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"b",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"(Z|dZ|j)",
			},
		},
		{
			patterns: []string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"v",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"x",
				"",
				"",
				"ks",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"",
				"(ts|dz)",
			},
		},
	},
	sepportuguese: []rule{
		{
			patterns: []string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"sc",
				"",
				"[ei]",
				"s",
			},
		},
		{
			patterns: []string{
				"sç",
				"",
				"[aou]",
				"s",
			},
		},
		{
			patterns: []string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[ei]",
				"s",
			},
		},
		{
			patterns: []string{
				"s",
				"^",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"s",
				"[aáuiíoóeéêy]",
				"[aáuiíoóeéêy]",
				"z",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"[dglmnrv]",
				"(Z|S)",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"$",
				"(Z|s|S)",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"[bdgv]",
				"(Z|z)",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"[ptckf]",
				"(s|S|z)",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[eiu]",
				"g",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[ao]",
				"gv",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[ei]",
				"Z",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"[eiu]",
				"k",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"[ao]",
				"kv",
			},
		},
		{
			patterns: []string{
				"uo",
				"",
				"",
				"(vo|o|u)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"[aei]",
				"v",
			},
		},
		{
			patterns: []string{
				"lh",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"nh",
				"",
				"",
				"nj",
			},
		},
		{
			patterns: []string{
				"h",
				"[bdgt]",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"ex",
				"",
				"[aáuiíoóeéêy]",
				"(ez|eS|eks)",
			},
		},
		{
			patterns: []string{
				"ex",
				"",
				"[cs]",
				"e",
			},
		},
		{
			patterns: []string{
				"y",
				"[aáuiíoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"[aeiíou]",
				"j",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"[bcdfglnprstv]",
				"(m|n)",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"$",
				"(m|n)",
			},
		},
		{
			patterns: []string{
				"ão",
				"",
				"",
				"(au|an|on)",
			},
		},
		{
			patterns: []string{
				"ãe",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: []string{
				"ãi",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: []string{
				"õe",
				"",
				"",
				"(oj|on)",
			},
		},
		{
			patterns: []string{
				"i",
				"[aáuoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: []string{
				"â",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"á",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"ã",
				"",
				"",
				"(a|an|on)",
			},
		},
		{
			patterns: []string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"í",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ô",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ó",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"õ",
				"",
				"",
				"(o|on)",
			},
		},
		{
			patterns: []string{
				"ú",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ü",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"aue",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"b",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"",
				"(e|i)",
			},
		},
		{
			patterns: []string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"",
				"(o|u)",
			},
		},
		{
			patterns: []string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"v",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"x",
				"",
				"",
				"(S|ks)",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"",
				"z",
			},
		},
	},
	sepspanish: []rule{
		{
			patterns: []string{
				"ñ",
				"",
				"",
				"(n|nj)",
			},
		},
		{
			patterns: []string{
				"ny",
				"",
				"",
				"nj",
			},
		},
		{
			patterns: []string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ig",
				"[aeiou]",
				"",
				"(tS|ig)",
			},
		},
		{
			patterns: []string{
				"ix",
				"[aeiou]",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"tx",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"tj",
				"",
				"$",
				"tS",
			},
		},
		{
			patterns: []string{
				"tj",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: []string{
				"tg",
				"",
				"",
				"(tg|dZ)",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"(tS|dZ)",
			},
		},
		{
			patterns: []string{
				"bh",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"h",
				"[dgt]",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"(x|Z)",
			},
		},
		{
			patterns: []string{
				"x",
				"",
				"",
				"(ks|gz|S)",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"v",
				"^",
				"",
				"(B|v)",
			},
		},
		{
			patterns: []string{
				"b",
				"^",
				"",
				"(b|V)",
			},
		},
		{
			patterns: []string{
				"v",
				"",
				"",
				"(b|v)",
			},
		},
		{
			patterns: []string{
				"b",
				"",
				"",
				"(b|v)",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"[bpvf]",
				"(m|n)",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[ei]",
				"s",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"",
				"(z|s)",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[ei]",
				"(g|gv)",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[ei]",
				"(x|g|dZ)",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"[aei]",
				"v",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"",
				"(i|j)",
			},
		},
		{
			patterns: []string{
				"ü",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"á",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"í",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ó",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ú",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"è",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ò",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"",
				"u",
			},
		},
	},
}

var sepLangRules = []langRule{
	{
		pattern: "eau",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ou",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "gni",
		langs:   10,
		accept:  true,
	},
	{
		pattern: "tx",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "tj",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "gy",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "guy",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "sh",
		langs:   48,
		accept:  true,
	},
	{
		pattern: "lh",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "nh",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ny",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "gue",
		langs:   34,
		accept:  true,
	},
	{
		pattern: "gui",
		langs:   34,
		accept:  true,
	},
	{
		pattern: "gia",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "gie",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "gio",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "giu",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ñ",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "â",
		langs:   18,
		accept:  true,
	},
	{
		pattern: "á",
		langs:   48,
		accept:  true,
	},
	{
		pattern: "à",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ã",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ê",
		langs:   18,
		accept:  true,
	},
	{
		pattern: "í",
		langs:   48,
		accept:  true,
	},
	{
		pattern: "î",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ô",
		langs:   18,
		accept:  true,
	},
	{
		pattern: "õ",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ò",
		langs:   40,
		accept:  true,
	},
	{
		pattern: "ú",
		langs:   48,
		accept:  true,
	},
	{
		pattern: "ù",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ü",
		langs:   48,
		accept:  true,
	},
	{
		pattern: "א",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ב",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ג",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ד",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ה",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ו",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ז",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ח",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ט",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "י",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "כ",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ל",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "מ",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "נ",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ס",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ע",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "פ",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "צ",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ק",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ר",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ש",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ת",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "a",
		langs:   4,
		accept:  false,
	},
	{
		pattern: "o",
		langs:   4,
		accept:  false,
	},
	{
		pattern: "e",
		langs:   4,
		accept:  false,
	},
	{
		pattern: "i",
		langs:   4,
		accept:  false,
	},
	{
		pattern: "y",
		langs:   4,
		accept:  false,
	},
	{
		pattern: "u",
		langs:   4,
		accept:  false,
	},
	{
		pattern: "kh",
		langs:   32,
		accept:  false,
	},
	{
		pattern: "gua",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "guo",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "ç",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "cha",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "cho",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "chu",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "j",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "dj",
		langs:   32,
		accept:  false,
	},
	{
		pattern: "sce",
		langs:   2,
		accept:  false,
	},
	{
		pattern: "sci",
		langs:   2,
		accept:  false,
	},
	{
		pattern: "ó",
		langs:   2,
		accept:  false,
	},
	{
		pattern: "è",
		langs:   16,
		accept:  false,
	},
}

var sepFinalRules = finalRules{
	approx: finalRule{
		first: []rule{
			{
				patterns: []string{
					"h",
					"",
					"$",
					"",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"[fktSs]",
					"p",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"p",
					"",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"$",
					"p",
				},
			},
			{
				patterns: []string{
					"p",
					"",
					"[vgdZz]",
					"b",
				},
			},
			{
				patterns: []string{
					"p",
					"",
					"b",
					"",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"[pktSs]",
					"f",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"f",
					"",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"$",
					"f",
				},
			},
			{
				patterns: []string{
					"f",
					"",
					"[vbgdZz]",
					"v",
				},
			},
			{
				patterns: []string{
					"f",
					"",
					"v",
					"",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"[pftSs]",
					"k",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"k",
					"",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"$",
					"k",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"[vbdZz]",
					"g",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"g",
					"",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"[pfkSs]",
					"t",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"t",
					"",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"$",
					"t",
				},
			},
			{
				patterns: []string{
					"t",
					"",
					"[vbgZz]",
					"d",
				},
			},
			{
				patterns: []string{
					"t",
					"",
					"d",
					"",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"dZ",
					"",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"tS",
					"",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"[pfkSt]",
					"s",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"Z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"nm",
					"",
					"",
					"m",
				},
			},
			{
				patterns: []string{
					"ji",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"a",
					"",
					"a",
					"",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"b",
					"",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"d",
					"",
				},
			},
			{
				patterns: []string{
					"e",
					"",
					"e",
					"",
				},
			},
			{
				patterns: []string{
					"f",
					"",
					"f",
					"",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"g",
					"",
				},
			},
			{
				patterns: []string{
					"i",
					"",
					"i",
					"",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"k",
					"",
				},
			},
			{
				patterns: []string{
					"l",
					"",
					"l",
					"",
				},
			},
			{
				patterns: []string{
					"m",
					"",
					"m",
					"",
				},
			},
			{
				patterns: []string{
					"n",
					"",
					"n",
					"",
				},
			},
			{
				patterns: []string{
					"o",
					"",
					"o",
					"",
				},
			},
			{
				patterns: []string{
					"p",
					"",
					"p",
					"",
				},
			},
			{
				patterns: []string{
					"r",
					"",
					"r",
					"",
				},
			},
			{
				patterns: []string{
					"t",
					"",
					"t",
					"",
				},
			},
			{
				patterns: []string{
					"u",
					"",
					"u",
					"",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"v",
					"",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"z",
					"",
				},
			},
			{
				patterns: []string{
					"mbr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: []string{
					"mpr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: []string{
					"bens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"benS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"ben",
					"^",
					"",
					"(bin|)",
				},
			},
			{
				patterns: []string{
					"bar",
					"^",
					"",
					"(bar|)",
				},
			},
			{
				patterns: []string{
					"abens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"abenS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"aben",
					"^",
					"",
					"(bin|bun|)",
				},
			},
			{
				patterns: []string{
					"abe",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: []string{
					"abi",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: []string{
					"abou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: []string{
					"abu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: []string{
					"bou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: []string{
					"bu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: []string{
					"ibn",
					"^",
					"",
					"(ibn|)",
				},
			},
			{
				patterns: []string{
					"els",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: []string{
					"elS",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: []string{
					"el",
					"^",
					"",
					"(il|l|)",
				},
			},
			{
				patterns: []string{
					"als",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: []string{
					"alS",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: []string{
					"al",
					"^",
					"",
					"(l|)",
				},
			},
			{
				patterns: []string{
					"dal",
					"^",
					"",
					"(dal|[8])",
				},
			},
			{
				patterns: []string{
					"da",
					"^",
					"",
					"(da|a|)",
				},
			},
			{
				patterns: []string{
					"della",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: []string{
					"dela",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: []string{
					"del",
					"^",
					"",
					"(dil|)",
				},
			},
			{
				patterns: []string{
					"des",
					"^",
					"",
					"(dis|)",
				},
			},
			{
				patterns: []string{
					"de",
					"^",
					"",
					"(di|i|)",
				},
			},
			{
				patterns: []string{
					"di",
					"^",
					"",
					"(di|i|[8])",
				},
			},
			{
				patterns: []string{
					"do",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: []string{
					"du",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: []string{
					"oa",
					"",
					"",
					"(va|a)",
				},
			},
			{
				patterns: []string{
					"oe",
					"",
					"",
					"(vi|i)",
				},
			},
			{
				patterns: []string{
					"ae",
					"",
					"",
					"(a|i)",
				},
			},
			{
				patterns: []string{
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: []string{
					"ha",
					"^",
					"",
					"(ha|a|)",
				},
			},
			{
				patterns: []string{
					"h",
					"",
					"",
					"(|h)",
				},
			},
			{
				patterns: []string{
					"x",
					"",
					"",
					"h",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"",
					"(h|k)",
				},
			},
			{
				patterns: []string{
					"aja",
					"^",
					"",
					"(Da|ia)",
				},
			},
			{
				patterns: []string{
					"aje",
					"^",
					"",
					"(Di|Da|i|ia)",
				},
			},
			{
				patterns: []string{
					"aji",
					"^",
					"",
					"(Di|i)",
				},
			},
			{
				patterns: []string{
					"ajo",
					"^",
					"",
					"(Du|Da|iu|ia)",
				},
			},
			{
				patterns: []string{
					"aju",
					"^",
					"",
					"(Du|iu)",
				},
			},
			{
				patterns: []string{
					"aj",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: []string{
					"ej",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: []string{
					"oj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"au",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"eu",
					"",
					"",
					"(iu|i|u)",
				},
			},
			{
				patterns: []string{
					"ou",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"a",
					"^",
					"",
					"",
				},
			},
			{
				patterns: []string{
					"ja",
					"^",
					"",
					"ia",
				},
			},
			{
				patterns: []string{
					"je",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"jo",
					"^",
					"",
					"(iu|ia)",
				},
			},
			{
				patterns: []string{
					"ju",
					"^",
					"",
					"iu",
				},
			},
			{
				patterns: []string{
					"ja",
					"",
					"",
					"(a|ia)",
				},
			},
			{
				patterns: []string{
					"je",
					"",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"ji",
					"",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"jo",
					"",
					"",
					"(u|iu)",
				},
			},
			{
				patterns: []string{
					"ju",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"j",
					"",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"i",
					"",
					"$",
					"(i|)",
				},
			},
			{
				patterns: []string{
					"o",
					"",
					"$",
					"(a|u|i)",
				},
			},
			{
				patterns: []string{
					"o",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"a",
					"",
					"$",
					"(a|i)",
				},
			},
			{
				patterns: []string{
					"se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: []string{
					"Se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: []string{
					"s",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"S",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"dS",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: []string{
					"dZ",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: []string{
					"Z",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"",
					"s",
				},
			},
			{
				patterns: []string{
					"dZ",
					"",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"Z",
					"",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"be",
					"",
					"[fktSs]",
					"(p|bi)",
				},
			},
			{
				patterns: []string{
					"pe",
					"",
					"[vgdZz]",
					"(b|pi)",
				},
			},
			{
				patterns: []string{
					"ve",
					"",
					"[pktSs]",
					"(f|vi)",
				},
			},
			{
				patterns: []string{
					"fe",
					"",
					"[vbgdZz]",
					"(v|fi)",
				},
			},
			{
				patterns: []string{
					"ge",
					"",
					"[pftSs]",
					"(k|gi)",
				},
			},
			{
				patterns: []string{
					"ke",
					"",
					"[vbdZz]",
					"(g|ki)",
				},
			},
			{
				patterns: []string{
					"de",
					"",
					"[pfkSs]",
					"(t|di)",
				},
			},
			{
				patterns: []string{
					"te",
					"",
					"[vbgZz]",
					"(d|ti)",
				},
			},
			{
				patterns: []string{
					"ze",
					"",
					"[pfkSt]",
					"(s|zi)",
				},
			},
			{
				patterns: []string{
					"e",
					"",
					"",
					"(i|)",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"",
					"b",
				},
			},
			{
				patterns: []string{
					"V",
					"",
					"",
					"v",
				},
			},
			{
				patterns: []string{
					"p",
					"^",
					"",
					"b",
				},
			},
		},
		second: []secondFinalRule{
			{
				langs: 0,
				rules: []rule{},
			},
			{
				langs: 1,
				rules: []rule{},
			},
			{
				langs: 2,
				rules: []rule{},
			},
			{
				langs: 3,
				rules: []rule{},
			},
			{
				langs: 4,
				rules: []rule{},
			},
			{
				langs: 5,
				rules: []rule{},
			},
		},
	},
	exact: finalRule{
		first: []rule{
			{
				patterns: []string{
					"h",
					"",
					"$",
					"",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"[fktSs]",
					"p",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"p",
					"",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"$",
					"p",
				},
			},
			{
				patterns: []string{
					"p",
					"",
					"[vgdZz]",
					"b",
				},
			},
			{
				patterns: []string{
					"p",
					"",
					"b",
					"",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"[pktSs]",
					"f",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"f",
					"",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"$",
					"f",
				},
			},
			{
				patterns: []string{
					"f",
					"",
					"[vbgdZz]",
					"v",
				},
			},
			{
				patterns: []string{
					"f",
					"",
					"v",
					"",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"[pftSs]",
					"k",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"k",
					"",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"$",
					"k",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"[vbdZz]",
					"g",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"g",
					"",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"[pfkSs]",
					"t",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"t",
					"",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"$",
					"t",
				},
			},
			{
				patterns: []string{
					"t",
					"",
					"[vbgZz]",
					"d",
				},
			},
			{
				patterns: []string{
					"t",
					"",
					"d",
					"",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"dZ",
					"",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"tS",
					"",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"[pfkSt]",
					"s",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"Z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: []string{
					"nm",
					"",
					"",
					"m",
				},
			},
			{
				patterns: []string{
					"ji",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"a",
					"",
					"a",
					"",
				},
			},
			{
				patterns: []string{
					"b",
					"",
					"b",
					"",
				},
			},
			{
				patterns: []string{
					"d",
					"",
					"d",
					"",
				},
			},
			{
				patterns: []string{
					"e",
					"",
					"e",
					"",
				},
			},
			{
				patterns: []string{
					"f",
					"",
					"f",
					"",
				},
			},
			{
				patterns: []string{
					"g",
					"",
					"g",
					"",
				},
			},
			{
				patterns: []string{
					"i",
					"",
					"i",
					"",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"k",
					"",
				},
			},
			{
				patterns: []string{
					"l",
					"",
					"l",
					"",
				},
			},
			{
				patterns: []string{
					"m",
					"",
					"m",
					"",
				},
			},
			{
				patterns: []string{
					"n",
					"",
					"n",
					"",
				},
			},
			{
				patterns: []string{
					"o",
					"",
					"o",
					"",
				},
			},
			{
				patterns: []string{
					"p",
					"",
					"p",
					"",
				},
			},
			{
				patterns: []string{
					"r",
					"",
					"r",
					"",
				},
			},
			{
				patterns: []string{
					"t",
					"",
					"t",
					"",
				},
			},
			{
				patterns: []string{
					"u",
					"",
					"u",
					"",
				},
			},
			{
				patterns: []string{
					"v",
					"",
					"v",
					"",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"z",
					"",
				},
			},
			{
				patterns: []string{
					"mbr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: []string{
					"mpr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: []string{
					"bens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"benS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"ben",
					"^",
					"",
					"(bin|)",
				},
			},
			{
				patterns: []string{
					"bar",
					"^",
					"",
					"(bar|)",
				},
			},
			{
				patterns: []string{
					"abens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"abenS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: []string{
					"aben",
					"^",
					"",
					"(bin|bun|)",
				},
			},
			{
				patterns: []string{
					"abe",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: []string{
					"abi",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: []string{
					"abou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: []string{
					"abu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: []string{
					"bou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: []string{
					"bu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: []string{
					"ibn",
					"^",
					"",
					"(ibn|)",
				},
			},
			{
				patterns: []string{
					"els",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: []string{
					"elS",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: []string{
					"el",
					"^",
					"",
					"(il|l|)",
				},
			},
			{
				patterns: []string{
					"als",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: []string{
					"alS",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: []string{
					"al",
					"^",
					"",
					"(l|)",
				},
			},
			{
				patterns: []string{
					"dal",
					"^",
					"",
					"(dal|[8])",
				},
			},
			{
				patterns: []string{
					"da",
					"^",
					"",
					"(da|a|)",
				},
			},
			{
				patterns: []string{
					"della",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: []string{
					"dela",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: []string{
					"del",
					"^",
					"",
					"(dil|)",
				},
			},
			{
				patterns: []string{
					"des",
					"^",
					"",
					"(dis|)",
				},
			},
			{
				patterns: []string{
					"de",
					"^",
					"",
					"(di|i|)",
				},
			},
			{
				patterns: []string{
					"di",
					"^",
					"",
					"(di|i|[8])",
				},
			},
			{
				patterns: []string{
					"do",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: []string{
					"du",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: []string{
					"oa",
					"",
					"",
					"(va|a)",
				},
			},
			{
				patterns: []string{
					"oe",
					"",
					"",
					"(vi|i)",
				},
			},
			{
				patterns: []string{
					"ae",
					"",
					"",
					"(a|i)",
				},
			},
			{
				patterns: []string{
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: []string{
					"ha",
					"^",
					"",
					"(ha|a|)",
				},
			},
			{
				patterns: []string{
					"h",
					"",
					"",
					"(|h)",
				},
			},
			{
				patterns: []string{
					"x",
					"",
					"",
					"h",
				},
			},
			{
				patterns: []string{
					"k",
					"",
					"",
					"(h|k)",
				},
			},
			{
				patterns: []string{
					"aja",
					"^",
					"",
					"(Da|ia)",
				},
			},
			{
				patterns: []string{
					"aje",
					"^",
					"",
					"(Di|Da|i|ia)",
				},
			},
			{
				patterns: []string{
					"aji",
					"^",
					"",
					"(Di|i)",
				},
			},
			{
				patterns: []string{
					"ajo",
					"^",
					"",
					"(Du|Da|iu|ia)",
				},
			},
			{
				patterns: []string{
					"aju",
					"^",
					"",
					"(Du|iu)",
				},
			},
			{
				patterns: []string{
					"aj",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: []string{
					"ej",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: []string{
					"oj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"au",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"eu",
					"",
					"",
					"(iu|i|u)",
				},
			},
			{
				patterns: []string{
					"ou",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"a",
					"^",
					"",
					"",
				},
			},
			{
				patterns: []string{
					"ja",
					"^",
					"",
					"ia",
				},
			},
			{
				patterns: []string{
					"je",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"jo",
					"^",
					"",
					"(iu|ia)",
				},
			},
			{
				patterns: []string{
					"ju",
					"^",
					"",
					"iu",
				},
			},
			{
				patterns: []string{
					"ja",
					"",
					"",
					"(a|ia)",
				},
			},
			{
				patterns: []string{
					"je",
					"",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"ji",
					"",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"jo",
					"",
					"",
					"(u|iu)",
				},
			},
			{
				patterns: []string{
					"ju",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"j",
					"",
					"",
					"i",
				},
			},
			{
				patterns: []string{
					"i",
					"",
					"$",
					"(i|)",
				},
			},
			{
				patterns: []string{
					"o",
					"",
					"$",
					"(a|u|i)",
				},
			},
			{
				patterns: []string{
					"o",
					"",
					"",
					"u",
				},
			},
			{
				patterns: []string{
					"a",
					"",
					"$",
					"(a|i)",
				},
			},
			{
				patterns: []string{
					"se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: []string{
					"s",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: []string{
					"Se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: []string{
					"s",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"S",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"dS",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: []string{
					"dZ",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: []string{
					"Z",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: []string{
					"z",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: []string{
					"S",
					"",
					"",
					"s",
				},
			},
			{
				patterns: []string{
					"dZ",
					"",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"Z",
					"",
					"",
					"z",
				},
			},
			{
				patterns: []string{
					"be",
					"",
					"[fktSs]",
					"(p|bi)",
				},
			},
			{
				patterns: []string{
					"pe",
					"",
					"[vgdZz]",
					"(b|pi)",
				},
			},
			{
				patterns: []string{
					"ve",
					"",
					"[pktSs]",
					"(f|vi)",
				},
			},
			{
				patterns: []string{
					"fe",
					"",
					"[vbgdZz]",
					"(v|fi)",
				},
			},
			{
				patterns: []string{
					"ge",
					"",
					"[pftSs]",
					"(k|gi)",
				},
			},
			{
				patterns: []string{
					"ke",
					"",
					"[vbdZz]",
					"(g|ki)",
				},
			},
			{
				patterns: []string{
					"de",
					"",
					"[pfkSs]",
					"(t|di)",
				},
			},
			{
				patterns: []string{
					"te",
					"",
					"[vbgZz]",
					"(d|ti)",
				},
			},
			{
				patterns: []string{
					"ze",
					"",
					"[pfkSt]",
					"(s|zi)",
				},
			},
			{
				patterns: []string{
					"e",
					"",
					"",
					"(i|)",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"",
					"b",
				},
			},
			{
				patterns: []string{
					"V",
					"",
					"",
					"v",
				},
			},
			{
				patterns: []string{
					"p",
					"^",
					"",
					"b",
				},
			},
		},
		second: []secondFinalRule{
			{
				langs: 0,
				rules: []rule{},
			},
			{
				langs: 1,
				rules: []rule{},
			},
			{
				langs: 2,
				rules: []rule{},
			},
			{
				langs: 3,
				rules: []rule{},
			},
			{
				langs: 4,
				rules: []rule{},
			},
			{
				langs: 5,
				rules: []rule{},
			},
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
