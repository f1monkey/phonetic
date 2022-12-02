// GENERATED CODE. DO NOT EDIT!
package bmpm

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
			patterns: [4]string{
				"ph",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"sh",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"gli",
				"",
				"",
				"(gli|l[8])",
			},
		},
		{
			patterns: [4]string{
				"gni",
				"",
				"",
				"(gni|ni[10])",
			},
		},
		{
			patterns: [4]string{
				"gn",
				"",
				"[aeou]",
				"(n[10]|nj[10]|gn)",
			},
		},
		{
			patterns: [4]string{
				"gh",
				"",
				"",
				"(g|gh)",
			},
		},
		{
			patterns: [4]string{
				"dh",
				"",
				"",
				"(d|dh)",
			},
		},
		{
			patterns: [4]string{
				"bh",
				"",
				"",
				"(b|bh)",
			},
		},
		{
			patterns: [4]string{
				"th",
				"",
				"",
				"(t|th)",
			},
		},
		{
			patterns: [4]string{
				"lh",
				"",
				"",
				"(l[16]|lh)",
			},
		},
		{
			patterns: [4]string{
				"nh",
				"",
				"",
				"(n[16]|nh)",
			},
		},
		{
			patterns: [4]string{
				"ig",
				"[aeiou]",
				"",
				"(ig|tS[32])",
			},
		},
		{
			patterns: [4]string{
				"ix",
				"[aeiou]",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"tx",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"tj",
				"",
				"$",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"tj",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"tg",
				"",
				"",
				"(tg|dZ[32])",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"[aeou]",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"y",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"gg",
				"",
				"[ei]",
				"(gZ[18]|dZ[40]|x[32])",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[ei]",
				"(Z[18]|dZ[40]|x[32])",
			},
		},
		{
			patterns: [4]string{
				"guy",
				"",
				"",
				"gi",
			},
		},
		{
			patterns: [4]string{
				"gue",
				"",
				"$",
				"(k[2]|ge)",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[ei]",
				"(g|gv)",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[ao]",
				"gv",
			},
		},
		{
			patterns: [4]string{
				"ñ",
				"",
				"",
				"(n|nj)",
			},
		},
		{
			patterns: [4]string{
				"ny",
				"",
				"",
				"nj",
			},
		},
		{
			patterns: [4]string{
				"sc",
				"",
				"[ei]",
				"(s|S[8])",
			},
		},
		{
			patterns: [4]string{
				"sç",
				"",
				"[aeiou]",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"[ei]",
				"(k[8]|S[18]|tS[32]|dZ[32])",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"(S|tS[32]|dZ[32])",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"[aeou]",
				"(tS[8]|si)",
			},
		},
		{
			patterns: [4]string{
				"cc",
				"",
				"[eiyéèê]",
				"(tS[8]|ks[50])",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[eiyéèê]",
				"(tS[8]|s[50])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"^",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aáuiíoóeéêy]",
				"[aáuiíoóeéêy]",
				"(s[32]|z[26])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"[dglmnrv]",
				"(z|s|Z[16])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"$",
				"(s|ts[8]|S[16])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"[bdgv]",
				"(z|dz[8]|Z[16])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"[ptckf]",
				"(s|ts[8]|S[16])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"",
				"(z|dz[8]|ts[8]|s[32])",
			},
		},
		{
			patterns: [4]string{
				"que",
				"",
				"$",
				"(k[2]|ke)",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"[eiu]",
				"k",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"[ao]",
				"(kv|k)",
			},
		},
		{
			patterns: [4]string{
				"ex",
				"",
				"[aáuiíoóeéêy]",
				"(ez[16]|eS[16]|eks|egz)",
			},
		},
		{
			patterns: [4]string{
				"ex",
				"",
				"[cs]",
				"(e[16]|ek)",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"[cdglnrst]",
				"(m|n[16])",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"[bfpv]",
				"(m|n[48])",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"$",
				"(m|n[16])",
			},
		},
		{
			patterns: [4]string{
				"b",
				"^",
				"",
				"(b|V[32])",
			},
		},
		{
			patterns: [4]string{
				"v",
				"^",
				"",
				"(v|B[32])",
			},
		},
		{
			patterns: [4]string{
				"eau",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ouh",
				"",
				"[aioe]",
				"(v[2]|uh)",
			},
		},
		{
			patterns: [4]string{
				"uh",
				"",
				"[aioe]",
				"(v|uh)",
			},
		},
		{
			patterns: [4]string{
				"ou",
				"",
				"[aioe]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"[aie]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aáuoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[aáuiíoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"[aeiíou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"$",
				"(e|[2])",
			},
		},
		{
			patterns: [4]string{
				"ão",
				"",
				"",
				"(au|an)",
			},
		},
		{
			patterns: [4]string{
				"ãe",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: [4]string{
				"ãi",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: [4]string{
				"õe",
				"",
				"",
				"(oj|on)",
			},
		},
		{
			patterns: [4]string{
				"où",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ou",
				"",
				"",
				"(ou|u[2])",
			},
		},
		{
			patterns: [4]string{
				"â",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"á",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"ã",
				"",
				"",
				"(a|an)",
			},
		},
		{
			patterns: [4]string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"è",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"í",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"î",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ô",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ó",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"õ",
				"",
				"",
				"(o|on)",
			},
		},
		{
			patterns: [4]string{
				"ò",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ú",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ü",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"b",
				"",
				"",
				"(b|v[32])",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"(x[32]|Z)",
			},
		},
		{
			patterns: [4]string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"",
				"(s|S[16])",
			},
		},
		{
			patterns: [4]string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"v",
				"",
				"",
				"(v|b[32])",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"x",
				"",
				"",
				"(ks|gz|S[48])",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"",
				"z",
			},
		},
	},
	sepfrench: []rule{
		{
			patterns: [4]string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"ph",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"x",
				"",
				"",
				"ks",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[eiyéèê]",
				"s",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"gn",
				"",
				"",
				"(n|gn)",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[eiy]",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"gue",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[eiy]",
				"g",
			},
		},
		{
			patterns: [4]string{
				"que",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aeiouyéèê]",
				"[aeiouyéèê]",
				"z",
			},
		},
		{
			patterns: [4]string{
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[bdgt]",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"$",
				"",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"ouh",
				"",
				"[aioe]",
				"(v|uh)",
			},
		},
		{
			patterns: [4]string{
				"ou",
				"",
				"[aeio]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"[aeio]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"aue",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: [4]string{
				"eau",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ai",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ay",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"è",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"â",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"où",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ou",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"oi",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[ou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"$",
				"(e|)",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"[aoeu]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"b",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"v",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"",
				"z",
			},
		},
	},
	sephebrew: []rule{
		{
			patterns: [4]string{
				"אי",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"עי",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"עו",
				"",
				"",
				"VV",
			},
		},
		{
			patterns: [4]string{
				"או",
				"",
				"",
				"VV",
			},
		},
		{
			patterns: [4]string{
				"ג׳",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"ד׳",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"א",
				"",
				"",
				"L",
			},
		},
		{
			patterns: [4]string{
				"ב",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"ג",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"ד",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"ה",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ה",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ה",
				"",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"וו",
				"",
				"",
				"V",
			},
		},
		{
			patterns: [4]string{
				"וי",
				"",
				"",
				"WW",
			},
		},
		{
			patterns: [4]string{
				"ו",
				"",
				"",
				"W",
			},
		},
		{
			patterns: [4]string{
				"ז",
				"",
				"",
				"z",
			},
		},
		{
			patterns: [4]string{
				"ח",
				"",
				"",
				"X",
			},
		},
		{
			patterns: [4]string{
				"ט",
				"",
				"",
				"T",
			},
		},
		{
			patterns: [4]string{
				"יי",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"י",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ך",
				"",
				"",
				"X",
			},
		},
		{
			patterns: [4]string{
				"כ",
				"^",
				"",
				"K",
			},
		},
		{
			patterns: [4]string{
				"כ",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"ל",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"ם",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"מ",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"ן",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"נ",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"ס",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ע",
				"",
				"",
				"L",
			},
		},
		{
			patterns: [4]string{
				"ף",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"פ",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"ץ",
				"",
				"",
				"C",
			},
		},
		{
			patterns: [4]string{
				"צ",
				"",
				"",
				"C",
			},
		},
		{
			patterns: [4]string{
				"ק",
				"",
				"",
				"K",
			},
		},
		{
			patterns: [4]string{
				"ר",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"ש",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ת",
				"",
				"",
				"T",
			},
		},
	},
	sepitalian: []rule{
		{
			patterns: [4]string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"gli",
				"",
				"",
				"(l|gli)",
			},
		},
		{
			patterns: [4]string{
				"gn",
				"",
				"[aeou]",
				"(n|nj|gn)",
			},
		},
		{
			patterns: [4]string{
				"gni",
				"",
				"",
				"(ni|gni)",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"[aeou]",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"gg",
				"",
				"[ei]",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[ei]",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[bdgt]",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"[aeou]",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"[ei]",
				"k",
			},
		},
		{
			patterns: [4]string{
				"sc",
				"",
				"[ei]",
				"S",
			},
		},
		{
			patterns: [4]string{
				"cc",
				"",
				"[ei]",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[ei]",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aeiou]",
				"[aeiou]",
				"z",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aeou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[aeou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"[aei]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"�",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"�",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"�",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"�",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"b",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"(Z|dZ|j)",
			},
		},
		{
			patterns: [4]string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"v",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"x",
				"",
				"",
				"ks",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"",
				"(ts|dz)",
			},
		},
	},
	sepportuguese: []rule{
		{
			patterns: [4]string{
				"kh",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"sc",
				"",
				"[ei]",
				"s",
			},
		},
		{
			patterns: [4]string{
				"sç",
				"",
				"[aou]",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[ei]",
				"s",
			},
		},
		{
			patterns: [4]string{
				"s",
				"^",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aáuiíoóeéêy]",
				"[aáuiíoóeéêy]",
				"z",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"[dglmnrv]",
				"(Z|S)",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"$",
				"(Z|s|S)",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"[bdgv]",
				"(Z|z)",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"[ptckf]",
				"(s|S|z)",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[eiu]",
				"g",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[ao]",
				"gv",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[ei]",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"[eiu]",
				"k",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"[ao]",
				"kv",
			},
		},
		{
			patterns: [4]string{
				"uo",
				"",
				"",
				"(vo|o|u)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"[aei]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"lh",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"nh",
				"",
				"",
				"nj",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[bdgt]",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"ex",
				"",
				"[aáuiíoóeéêy]",
				"(ez|eS|eks)",
			},
		},
		{
			patterns: [4]string{
				"ex",
				"",
				"[cs]",
				"e",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[aáuiíoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"[aeiíou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"[bcdfglnprstv]",
				"(m|n)",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"$",
				"(m|n)",
			},
		},
		{
			patterns: [4]string{
				"ão",
				"",
				"",
				"(au|an|on)",
			},
		},
		{
			patterns: [4]string{
				"ãe",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: [4]string{
				"ãi",
				"",
				"",
				"(aj|an)",
			},
		},
		{
			patterns: [4]string{
				"õe",
				"",
				"",
				"(oj|on)",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aáuoóeéê]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[aeou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"â",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"á",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"ã",
				"",
				"",
				"(a|an|on)",
			},
		},
		{
			patterns: [4]string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"í",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ô",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ó",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"õ",
				"",
				"",
				"(o|on)",
			},
		},
		{
			patterns: [4]string{
				"ú",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ü",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"aue",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"b",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"",
				"(e|i)",
			},
		},
		{
			patterns: [4]string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"",
				"(o|u)",
			},
		},
		{
			patterns: [4]string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"v",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"x",
				"",
				"",
				"(S|ks)",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"",
				"z",
			},
		},
	},
	sepspanish: []rule{
		{
			patterns: [4]string{
				"ñ",
				"",
				"",
				"(n|nj)",
			},
		},
		{
			patterns: [4]string{
				"ny",
				"",
				"",
				"nj",
			},
		},
		{
			patterns: [4]string{
				"ç",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ig",
				"[aeiou]",
				"",
				"(tS|ig)",
			},
		},
		{
			patterns: [4]string{
				"ix",
				"[aeiou]",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"tx",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"tj",
				"",
				"$",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"tj",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"tg",
				"",
				"",
				"(tg|dZ)",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"(tS|dZ)",
			},
		},
		{
			patterns: [4]string{
				"bh",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[dgt]",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"(x|Z)",
			},
		},
		{
			patterns: [4]string{
				"x",
				"",
				"",
				"(ks|gz|S)",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"v",
				"^",
				"",
				"(B|v)",
			},
		},
		{
			patterns: [4]string{
				"b",
				"^",
				"",
				"(b|V)",
			},
		},
		{
			patterns: [4]string{
				"v",
				"",
				"",
				"(b|v)",
			},
		},
		{
			patterns: [4]string{
				"b",
				"",
				"",
				"(b|v)",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"[bpvf]",
				"(m|n)",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[ei]",
				"s",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"",
				"(z|s)",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[ei]",
				"(g|gv)",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[ei]",
				"(x|g|dZ)",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"q",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"uo",
				"",
				"",
				"(vo|o)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"[aei]",
				"v",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"",
				"(i|j)",
			},
		},
		{
			patterns: [4]string{
				"ü",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"á",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"é",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"í",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ó",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ú",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"à",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"è",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ò",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"d",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"f",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"h",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"k",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"l",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"m",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"n",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"p",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"r",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"t",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
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
				patterns: [4]string{
					"h",
					"",
					"$",
					"",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"[fktSs]",
					"p",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"p",
					"",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"$",
					"p",
				},
			},
			{
				patterns: [4]string{
					"p",
					"",
					"[vgdZz]",
					"b",
				},
			},
			{
				patterns: [4]string{
					"p",
					"",
					"b",
					"",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"[pktSs]",
					"f",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"f",
					"",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"$",
					"f",
				},
			},
			{
				patterns: [4]string{
					"f",
					"",
					"[vbgdZz]",
					"v",
				},
			},
			{
				patterns: [4]string{
					"f",
					"",
					"v",
					"",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"[pftSs]",
					"k",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"k",
					"",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"$",
					"k",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"[vbdZz]",
					"g",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"g",
					"",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"[pfkSs]",
					"t",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"t",
					"",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"$",
					"t",
				},
			},
			{
				patterns: [4]string{
					"t",
					"",
					"[vbgZz]",
					"d",
				},
			},
			{
				patterns: [4]string{
					"t",
					"",
					"d",
					"",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"dZ",
					"",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"tS",
					"",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"[pfkSt]",
					"s",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"Z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"nm",
					"",
					"",
					"m",
				},
			},
			{
				patterns: [4]string{
					"ji",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"a",
					"",
					"a",
					"",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"b",
					"",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"d",
					"",
				},
			},
			{
				patterns: [4]string{
					"e",
					"",
					"e",
					"",
				},
			},
			{
				patterns: [4]string{
					"f",
					"",
					"f",
					"",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"g",
					"",
				},
			},
			{
				patterns: [4]string{
					"i",
					"",
					"i",
					"",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"k",
					"",
				},
			},
			{
				patterns: [4]string{
					"l",
					"",
					"l",
					"",
				},
			},
			{
				patterns: [4]string{
					"m",
					"",
					"m",
					"",
				},
			},
			{
				patterns: [4]string{
					"n",
					"",
					"n",
					"",
				},
			},
			{
				patterns: [4]string{
					"o",
					"",
					"o",
					"",
				},
			},
			{
				patterns: [4]string{
					"p",
					"",
					"p",
					"",
				},
			},
			{
				patterns: [4]string{
					"r",
					"",
					"r",
					"",
				},
			},
			{
				patterns: [4]string{
					"t",
					"",
					"t",
					"",
				},
			},
			{
				patterns: [4]string{
					"u",
					"",
					"u",
					"",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"v",
					"",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"z",
					"",
				},
			},
			{
				patterns: [4]string{
					"mbr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: [4]string{
					"mpr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: [4]string{
					"bens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"benS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"ben",
					"^",
					"",
					"(bin|)",
				},
			},
			{
				patterns: [4]string{
					"bar",
					"^",
					"",
					"(bar|)",
				},
			},
			{
				patterns: [4]string{
					"abens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"abenS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"aben",
					"^",
					"",
					"(bin|bun|)",
				},
			},
			{
				patterns: [4]string{
					"abe",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: [4]string{
					"abi",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: [4]string{
					"abou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: [4]string{
					"abu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: [4]string{
					"bou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: [4]string{
					"bu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: [4]string{
					"ibn",
					"^",
					"",
					"(ibn|)",
				},
			},
			{
				patterns: [4]string{
					"els",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: [4]string{
					"elS",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: [4]string{
					"el",
					"^",
					"",
					"(il|l|)",
				},
			},
			{
				patterns: [4]string{
					"als",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: [4]string{
					"alS",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: [4]string{
					"al",
					"^",
					"",
					"(l|)",
				},
			},
			{
				patterns: [4]string{
					"dal",
					"^",
					"",
					"(dal|[8])",
				},
			},
			{
				patterns: [4]string{
					"da",
					"^",
					"",
					"(da|a|)",
				},
			},
			{
				patterns: [4]string{
					"della",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: [4]string{
					"dela",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: [4]string{
					"del",
					"^",
					"",
					"(dil|)",
				},
			},
			{
				patterns: [4]string{
					"des",
					"^",
					"",
					"(dis|)",
				},
			},
			{
				patterns: [4]string{
					"de",
					"^",
					"",
					"(di|i|)",
				},
			},
			{
				patterns: [4]string{
					"di",
					"^",
					"",
					"(di|i|[8])",
				},
			},
			{
				patterns: [4]string{
					"do",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: [4]string{
					"du",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: [4]string{
					"oa",
					"",
					"",
					"(va|a)",
				},
			},
			{
				patterns: [4]string{
					"oe",
					"",
					"",
					"(vi|i)",
				},
			},
			{
				patterns: [4]string{
					"ae",
					"",
					"",
					"(a|i)",
				},
			},
			{
				patterns: [4]string{
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: [4]string{
					"ha",
					"^",
					"",
					"(ha|a|)",
				},
			},
			{
				patterns: [4]string{
					"h",
					"",
					"",
					"(|h)",
				},
			},
			{
				patterns: [4]string{
					"x",
					"",
					"",
					"h",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"",
					"(h|k)",
				},
			},
			{
				patterns: [4]string{
					"aja",
					"^",
					"",
					"(Da|ia)",
				},
			},
			{
				patterns: [4]string{
					"aje",
					"^",
					"",
					"(Di|Da|i|ia)",
				},
			},
			{
				patterns: [4]string{
					"aji",
					"^",
					"",
					"(Di|i)",
				},
			},
			{
				patterns: [4]string{
					"ajo",
					"^",
					"",
					"(Du|Da|iu|ia)",
				},
			},
			{
				patterns: [4]string{
					"aju",
					"^",
					"",
					"(Du|iu)",
				},
			},
			{
				patterns: [4]string{
					"aj",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: [4]string{
					"ej",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: [4]string{
					"oj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"au",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"eu",
					"",
					"",
					"(iu|i|u)",
				},
			},
			{
				patterns: [4]string{
					"ou",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"a",
					"^",
					"",
					"",
				},
			},
			{
				patterns: [4]string{
					"ja",
					"^",
					"",
					"ia",
				},
			},
			{
				patterns: [4]string{
					"je",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"jo",
					"^",
					"",
					"(iu|ia)",
				},
			},
			{
				patterns: [4]string{
					"ju",
					"^",
					"",
					"iu",
				},
			},
			{
				patterns: [4]string{
					"ja",
					"",
					"",
					"(a|ia)",
				},
			},
			{
				patterns: [4]string{
					"je",
					"",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"ji",
					"",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"jo",
					"",
					"",
					"(u|iu)",
				},
			},
			{
				patterns: [4]string{
					"ju",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"j",
					"",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"i",
					"",
					"$",
					"(i|)",
				},
			},
			{
				patterns: [4]string{
					"o",
					"",
					"$",
					"(a|u|i)",
				},
			},
			{
				patterns: [4]string{
					"o",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"a",
					"",
					"$",
					"(a|i)",
				},
			},
			{
				patterns: [4]string{
					"se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: [4]string{
					"Se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: [4]string{
					"s",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"S",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"dS",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: [4]string{
					"dZ",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: [4]string{
					"Z",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"",
					"s",
				},
			},
			{
				patterns: [4]string{
					"dZ",
					"",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"Z",
					"",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"be",
					"",
					"[fktSs]",
					"(p|bi)",
				},
			},
			{
				patterns: [4]string{
					"pe",
					"",
					"[vgdZz]",
					"(b|pi)",
				},
			},
			{
				patterns: [4]string{
					"ve",
					"",
					"[pktSs]",
					"(f|vi)",
				},
			},
			{
				patterns: [4]string{
					"fe",
					"",
					"[vbgdZz]",
					"(v|fi)",
				},
			},
			{
				patterns: [4]string{
					"ge",
					"",
					"[pftSs]",
					"(k|gi)",
				},
			},
			{
				patterns: [4]string{
					"ke",
					"",
					"[vbdZz]",
					"(g|ki)",
				},
			},
			{
				patterns: [4]string{
					"de",
					"",
					"[pfkSs]",
					"(t|di)",
				},
			},
			{
				patterns: [4]string{
					"te",
					"",
					"[vbgZz]",
					"(d|ti)",
				},
			},
			{
				patterns: [4]string{
					"ze",
					"",
					"[pfkSt]",
					"(s|zi)",
				},
			},
			{
				patterns: [4]string{
					"e",
					"",
					"",
					"(i|)",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"",
					"b",
				},
			},
			{
				patterns: [4]string{
					"V",
					"",
					"",
					"v",
				},
			},
			{
				patterns: [4]string{
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
				patterns: [4]string{
					"h",
					"",
					"$",
					"",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"[fktSs]",
					"p",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"p",
					"",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"$",
					"p",
				},
			},
			{
				patterns: [4]string{
					"p",
					"",
					"[vgdZz]",
					"b",
				},
			},
			{
				patterns: [4]string{
					"p",
					"",
					"b",
					"",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"[pktSs]",
					"f",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"f",
					"",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"$",
					"f",
				},
			},
			{
				patterns: [4]string{
					"f",
					"",
					"[vbgdZz]",
					"v",
				},
			},
			{
				patterns: [4]string{
					"f",
					"",
					"v",
					"",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"[pftSs]",
					"k",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"k",
					"",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"$",
					"k",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"[vbdZz]",
					"g",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"g",
					"",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"[pfkSs]",
					"t",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"t",
					"",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"$",
					"t",
				},
			},
			{
				patterns: [4]string{
					"t",
					"",
					"[vbgZz]",
					"d",
				},
			},
			{
				patterns: [4]string{
					"t",
					"",
					"d",
					"",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"dZ",
					"",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"tS",
					"",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"[pfkSt]",
					"s",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"Z",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"[sSzZ]",
					"",
				},
			},
			{
				patterns: [4]string{
					"nm",
					"",
					"",
					"m",
				},
			},
			{
				patterns: [4]string{
					"ji",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"a",
					"",
					"a",
					"",
				},
			},
			{
				patterns: [4]string{
					"b",
					"",
					"b",
					"",
				},
			},
			{
				patterns: [4]string{
					"d",
					"",
					"d",
					"",
				},
			},
			{
				patterns: [4]string{
					"e",
					"",
					"e",
					"",
				},
			},
			{
				patterns: [4]string{
					"f",
					"",
					"f",
					"",
				},
			},
			{
				patterns: [4]string{
					"g",
					"",
					"g",
					"",
				},
			},
			{
				patterns: [4]string{
					"i",
					"",
					"i",
					"",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"k",
					"",
				},
			},
			{
				patterns: [4]string{
					"l",
					"",
					"l",
					"",
				},
			},
			{
				patterns: [4]string{
					"m",
					"",
					"m",
					"",
				},
			},
			{
				patterns: [4]string{
					"n",
					"",
					"n",
					"",
				},
			},
			{
				patterns: [4]string{
					"o",
					"",
					"o",
					"",
				},
			},
			{
				patterns: [4]string{
					"p",
					"",
					"p",
					"",
				},
			},
			{
				patterns: [4]string{
					"r",
					"",
					"r",
					"",
				},
			},
			{
				patterns: [4]string{
					"t",
					"",
					"t",
					"",
				},
			},
			{
				patterns: [4]string{
					"u",
					"",
					"u",
					"",
				},
			},
			{
				patterns: [4]string{
					"v",
					"",
					"v",
					"",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"z",
					"",
				},
			},
			{
				patterns: [4]string{
					"mbr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: [4]string{
					"mpr",
					"",
					"",
					"mr",
				},
			},
			{
				patterns: [4]string{
					"bens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"benS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"ben",
					"^",
					"",
					"(bin|)",
				},
			},
			{
				patterns: [4]string{
					"bar",
					"^",
					"",
					"(bar|)",
				},
			},
			{
				patterns: [4]string{
					"abens",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"abenS",
					"^",
					"",
					"(binz|s)",
				},
			},
			{
				patterns: [4]string{
					"aben",
					"^",
					"",
					"(bin|bun|)",
				},
			},
			{
				patterns: [4]string{
					"abe",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: [4]string{
					"abi",
					"^",
					"",
					"(bi|bu|)",
				},
			},
			{
				patterns: [4]string{
					"abou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: [4]string{
					"abu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: [4]string{
					"bou",
					"^",
					"",
					"(bu|[2])",
				},
			},
			{
				patterns: [4]string{
					"bu",
					"^",
					"",
					"(bu|)",
				},
			},
			{
				patterns: [4]string{
					"ibn",
					"^",
					"",
					"(ibn|)",
				},
			},
			{
				patterns: [4]string{
					"els",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: [4]string{
					"elS",
					"^",
					"",
					"(ilz|lz|s)",
				},
			},
			{
				patterns: [4]string{
					"el",
					"^",
					"",
					"(il|l|)",
				},
			},
			{
				patterns: [4]string{
					"als",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: [4]string{
					"alS",
					"^",
					"",
					"(lz|s)",
				},
			},
			{
				patterns: [4]string{
					"al",
					"^",
					"",
					"(l|)",
				},
			},
			{
				patterns: [4]string{
					"dal",
					"^",
					"",
					"(dal|[8])",
				},
			},
			{
				patterns: [4]string{
					"da",
					"^",
					"",
					"(da|a|)",
				},
			},
			{
				patterns: [4]string{
					"della",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: [4]string{
					"dela",
					"^",
					"",
					"(dila|)",
				},
			},
			{
				patterns: [4]string{
					"del",
					"^",
					"",
					"(dil|)",
				},
			},
			{
				patterns: [4]string{
					"des",
					"^",
					"",
					"(dis|)",
				},
			},
			{
				patterns: [4]string{
					"de",
					"^",
					"",
					"(di|i|)",
				},
			},
			{
				patterns: [4]string{
					"di",
					"^",
					"",
					"(di|i|[8])",
				},
			},
			{
				patterns: [4]string{
					"do",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: [4]string{
					"du",
					"^",
					"",
					"(du|u)",
				},
			},
			{
				patterns: [4]string{
					"oa",
					"",
					"",
					"(va|a)",
				},
			},
			{
				patterns: [4]string{
					"oe",
					"",
					"",
					"(vi|i)",
				},
			},
			{
				patterns: [4]string{
					"ae",
					"",
					"",
					"(a|i)",
				},
			},
			{
				patterns: [4]string{
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: [4]string{
					"ha",
					"^",
					"",
					"(ha|a|)",
				},
			},
			{
				patterns: [4]string{
					"h",
					"",
					"",
					"(|h)",
				},
			},
			{
				patterns: [4]string{
					"x",
					"",
					"",
					"h",
				},
			},
			{
				patterns: [4]string{
					"k",
					"",
					"",
					"(h|k)",
				},
			},
			{
				patterns: [4]string{
					"aja",
					"^",
					"",
					"(Da|ia)",
				},
			},
			{
				patterns: [4]string{
					"aje",
					"^",
					"",
					"(Di|Da|i|ia)",
				},
			},
			{
				patterns: [4]string{
					"aji",
					"^",
					"",
					"(Di|i)",
				},
			},
			{
				patterns: [4]string{
					"ajo",
					"^",
					"",
					"(Du|Da|iu|ia)",
				},
			},
			{
				patterns: [4]string{
					"aju",
					"^",
					"",
					"(Du|iu)",
				},
			},
			{
				patterns: [4]string{
					"aj",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: [4]string{
					"ej",
					"",
					"",
					"(D|i)",
				},
			},
			{
				patterns: [4]string{
					"oj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uj",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"au",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"eu",
					"",
					"",
					"(iu|i|u)",
				},
			},
			{
				patterns: [4]string{
					"ou",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"a",
					"^",
					"",
					"",
				},
			},
			{
				patterns: [4]string{
					"ja",
					"^",
					"",
					"ia",
				},
			},
			{
				patterns: [4]string{
					"je",
					"^",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"jo",
					"^",
					"",
					"(iu|ia)",
				},
			},
			{
				patterns: [4]string{
					"ju",
					"^",
					"",
					"iu",
				},
			},
			{
				patterns: [4]string{
					"ja",
					"",
					"",
					"(a|ia)",
				},
			},
			{
				patterns: [4]string{
					"je",
					"",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"ji",
					"",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"jo",
					"",
					"",
					"(u|iu)",
				},
			},
			{
				patterns: [4]string{
					"ju",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"j",
					"",
					"",
					"i",
				},
			},
			{
				patterns: [4]string{
					"i",
					"",
					"$",
					"(i|)",
				},
			},
			{
				patterns: [4]string{
					"o",
					"",
					"$",
					"(a|u|i)",
				},
			},
			{
				patterns: [4]string{
					"o",
					"",
					"",
					"u",
				},
			},
			{
				patterns: [4]string{
					"a",
					"",
					"$",
					"(a|i)",
				},
			},
			{
				patterns: [4]string{
					"se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: [4]string{
					"s",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: [4]string{
					"Se",
					"",
					"[rmnl]",
					"(z|si)",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"[rmnl]",
					"z",
				},
			},
			{
				patterns: [4]string{
					"s",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"S",
					"[rmnl]",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"dS",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: [4]string{
					"dZ",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: [4]string{
					"Z",
					"",
					"$",
					"S",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: [4]string{
					"z",
					"",
					"$",
					"(S|s)",
				},
			},
			{
				patterns: [4]string{
					"S",
					"",
					"",
					"s",
				},
			},
			{
				patterns: [4]string{
					"dZ",
					"",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"Z",
					"",
					"",
					"z",
				},
			},
			{
				patterns: [4]string{
					"be",
					"",
					"[fktSs]",
					"(p|bi)",
				},
			},
			{
				patterns: [4]string{
					"pe",
					"",
					"[vgdZz]",
					"(b|pi)",
				},
			},
			{
				patterns: [4]string{
					"ve",
					"",
					"[pktSs]",
					"(f|vi)",
				},
			},
			{
				patterns: [4]string{
					"fe",
					"",
					"[vbgdZz]",
					"(v|fi)",
				},
			},
			{
				patterns: [4]string{
					"ge",
					"",
					"[pftSs]",
					"(k|gi)",
				},
			},
			{
				patterns: [4]string{
					"ke",
					"",
					"[vbdZz]",
					"(g|ki)",
				},
			},
			{
				patterns: [4]string{
					"de",
					"",
					"[pfkSs]",
					"(t|di)",
				},
			},
			{
				patterns: [4]string{
					"te",
					"",
					"[vbgZz]",
					"(d|ti)",
				},
			},
			{
				patterns: [4]string{
					"ze",
					"",
					"[pfkSt]",
					"(s|zi)",
				},
			},
			{
				patterns: [4]string{
					"e",
					"",
					"",
					"(i|)",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"",
					"b",
				},
			},
			{
				patterns: [4]string{
					"V",
					"",
					"",
					"v",
				},
			},
			{
				patterns: [4]string{
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
