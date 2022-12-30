// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type ashLang uint64

const (
	ashany ashLang = 1 << iota
	ashcyrillic
	ashenglish
	ashfrench
	ashgerman
	ashhebrew
	ashhungarian
	ashpolish
	ashromanian
	ashrussian
	ashspanish
)

func (l ashLang) String() string {
	switch l {
	case ashany:
		return "any"
	case ashcyrillic:
		return "cyrillic"
	case ashenglish:
		return "english"
	case ashfrench:
		return "french"
	case ashgerman:
		return "german"
	case ashhebrew:
		return "hebrew"
	case ashhungarian:
		return "hungarian"
	case ashpolish:
		return "polish"
	case ashromanian:
		return "romanian"
	case ashrussian:
		return "russian"
	case ashspanish:
		return "spanish"
	}
	return ""
}

const ashAll = ashany +
	ashcyrillic +
	ashenglish +
	ashfrench +
	ashgerman +
	ashhebrew +
	ashhungarian +
	ashpolish +
	ashromanian +
	ashrussian +
	ashspanish

var ashRules = map[ashLang][]rule{
	ashany: []rule{
		{
			pattern:      "yna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(in[512]|ina)",
		},
		{
			pattern:      "ina",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(in[512]|ina)",
		},
		{
			pattern:      "liova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lof[512]|lef[512]|lova)",
		},
		{
			pattern:      "lova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lof[512]|lef[512]|lova)",
		},
		{
			pattern:      "ova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(of[512]|ova)",
		},
		{
			pattern:      "eva",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ef[512]|eva)",
		},
		{
			pattern:      "aia",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i[512])",
		},
		{
			pattern:      "aja",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i[512])",
		},
		{
			pattern:      "aya",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i[512])",
		},
		{
			pattern:      "lowa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lova|lof[128]|l[128]|el[128])",
		},
		{
			pattern:      "kowa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kova|kof[128]|k[128]|ek[128])",
		},
		{
			pattern:      "owa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ova|of[128]|)",
		},
		{
			pattern:      "lowna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lovna|levna|l[128]|el[128])",
		},
		{
			pattern:      "kowna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kovna|k[128]|ek[128])",
		},
		{
			pattern:      "owna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovna|[128])",
		},
		{
			pattern:      "lówna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(l|el[128])",
		},
		{
			pattern:      "kówna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k|ek[128])",
		},
		{
			pattern:      "ówna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(a|i[128])",
		},
		{
			pattern:     "rh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "r",
		},
		{
			pattern:  "ssch",
			phonetic: "S",
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
		},
		{
			pattern:  "tsch",
			phonetic: "tS",
		},
		{
			pattern:      "sch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(sk[256]|S|StS[512])",
		},
		{
			pattern:  "sch",
			phonetic: "(S|StS[512])",
		},
		{
			pattern:  "ssh",
			phonetic: "S",
		},
		{
			pattern:      "sh",
			rightContext: regexp.MustCompile("^[äöü]"),
			phonetic:     "sh",
		},
		{
			pattern:      "sh",
			rightContext: regexp.MustCompile("^[aeiou]"),
			phonetic:     "(S[516]|sh)",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:  "kh",
			phonetic: "(x[516]|kh)",
		},
		{
			pattern:  "chs",
			phonetic: "(ks[16]|xs|tSs[516])",
		},
		{
			pattern:      "ch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(x|k[256]|tS[516])",
		},
		{
			pattern:  "ch",
			phonetic: "(x|tS[516])",
		},
		{
			pattern:  "ck",
			phonetic: "(k|tsk[128])",
		},
		{
			pattern:  "czy",
			phonetic: "tSi",
		},
		{
			pattern:      "cze",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tSe|tSF)",
		},
		{
			pattern:  "ciewicz",
			phonetic: "(tsevitS|tSevitS)",
		},
		{
			pattern:  "siewicz",
			phonetic: "(sevitS|SevitS)",
		},
		{
			pattern:  "ziewicz",
			phonetic: "(zevitS|ZevitS)",
		},
		{
			pattern:  "riewicz",
			phonetic: "rjevitS",
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
		},
		{
			pattern:  "icz",
			phonetic: "itS",
		},
		{
			pattern:  "cz",
			phonetic: "tS",
		},
		{
			pattern:      "cia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tSB[128]|tsB)",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa[128]|tsa)",
		},
		{
			pattern:      "cią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(tSom[128]|tsom)",
		},
		{
			pattern:  "cią",
			phonetic: "(tSon[128]|tson)",
		},
		{
			pattern:      "cię",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(tSem[128]|tsem)",
		},
		{
			pattern:  "cię",
			phonetic: "(tSen[128]|tsen)",
		},
		{
			pattern:      "cie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tSF[128]|tsF)",
		},
		{
			pattern:  "cie",
			phonetic: "(tSe[128]|tse)",
		},
		{
			pattern:  "cio",
			phonetic: "(tSo[128]|tso)",
		},
		{
			pattern:  "ciu",
			phonetic: "(tSu[128]|tsu)",
		},
		{
			pattern:      "ci",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(tsi[128]|tSi[384]|tS[256]|si)",
		},
		{
			pattern:  "ci",
			phonetic: "(tsi[128]|tSi[384]|si)",
		},
		{
			pattern:      "ce",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tsF[128]|tSe[384]|se)",
		},
		{
			pattern:  "ce",
			phonetic: "(tSe[384]|tse[128]|se)",
		},
		{
			pattern:  "cy",
			phonetic: "(si|tsi[128])",
		},
		{
			pattern:  "ssz",
			phonetic: "S",
		},
		{
			pattern:  "sz",
			phonetic: "S",
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp[16]|sp)",
		},
		{
			pattern:  "sp",
			phonetic: "(Sp[16]|sp)",
		},
		{
			pattern:  "sst",
			phonetic: "(St[16]|st)",
		},
		{
			pattern:  "st",
			phonetic: "(St[16]|st)",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern:      "sia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(SB[128]|sB[128]|sja)",
		},
		{
			pattern:  "sia",
			phonetic: "(Sa[128]|sja)",
		},
		{
			pattern:      "sią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Som[128]|som)",
		},
		{
			pattern:  "sią",
			phonetic: "(Son[128]|son)",
		},
		{
			pattern:      "się",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Sem[128]|sem)",
		},
		{
			pattern:  "się",
			phonetic: "(Sen[128]|sen)",
		},
		{
			pattern:      "sie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(SF[128]|sF|zi[16])",
		},
		{
			pattern:  "sie",
			phonetic: "(se|Se[128]|zi[16])",
		},
		{
			pattern:  "sio",
			phonetic: "(So[128]|so)",
		},
		{
			pattern:  "siu",
			phonetic: "(Su[128]|sju)",
		},
		{
			pattern:  "si",
			phonetic: "(Si[128]|si|zi[16])",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[aeiouäöë]"),
			phonetic:     "(s|z[16])",
		},
		{
			pattern:  "gue",
			phonetic: "ge",
		},
		{
			pattern:  "gui",
			phonetic: "gi",
		},
		{
			pattern:  "guy",
			phonetic: "gi",
		},
		{
			pattern:      "gh",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(g[256]|gh)",
		},
		{
			pattern:      "gauz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "haus",
		},
		{
			pattern:      "gaus",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "haus",
		},
		{
			pattern:      "gol'ts",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "golts",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "gol'tz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:  "goltz",
			phonetic: "holts",
		},
		{
			pattern:     "gol'ts",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "holts",
		},
		{
			pattern:     "golts",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "holts",
		},
		{
			pattern:     "gol'tz",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "holts",
		},
		{
			pattern:     "goltz",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "holts",
		},
		{
			pattern:      "gendler",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hendler",
		},
		{
			pattern:      "gejmer",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "gejm",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "geymer",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "geym",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "geimer",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "geim",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "gof",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hof",
		},
		{
			pattern:      "ger",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ger",
		},
		{
			pattern:      "gen",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "gen",
		},
		{
			pattern:      "gin",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "gin",
		},
		{
			pattern:      "gie",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ge|gi[16]|ji[8])",
		},
		{
			pattern:  "gie",
			phonetic: "ge",
		},
		{
			pattern:     "ge",
			leftContext: regexp.MustCompile("[yaeiou]$"),
			phonetic:    "(gE|xe[1024]|dZe[260])",
		},
		{
			pattern:     "gi",
			leftContext: regexp.MustCompile("[yaeiou]$"),
			phonetic:    "(gI|xi[1024]|dZi[260])",
		},
		{
			pattern:  "ge",
			phonetic: "(gE|dZe[260]|hE[512]|xe[1024])",
		},
		{
			pattern:  "gi",
			phonetic: "(gI|dZi[260]|hI[512]|xi[1024])",
		},
		{
			pattern:      "gy",
			rightContext: regexp.MustCompile("^[aeouáéóúüöőű]"),
			phonetic:     "(gi|dj[64])",
		},
		{
			pattern:  "gy",
			phonetic: "(gi|d[64])",
		},
		{
			pattern:      "g",
			leftContext:  regexp.MustCompile("[jyaeiou]$"),
			rightContext: regexp.MustCompile("^[aouyei]"),
			phonetic:     "g",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[aouei]"),
			phonetic:     "(g|h[512])",
		},
		{
			pattern:  "ej",
			phonetic: "(aj|eZ[264]|ex[1024])",
		},
		{
			pattern:  "ej",
			phonetic: "aj",
		},
		{
			pattern:      "ly",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "(l|lj)",
		},
		{
			pattern:      "li",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "(l|lj)",
		},
		{
			pattern:      "lj",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "(l|lj)",
		},
		{
			pattern:  "lio",
			phonetic: "(lo|le[512])",
		},
		{
			pattern:  "lyo",
			phonetic: "(lo|le[512])",
		},
		{
			pattern:  "ll",
			phonetic: "(l|J[1024])",
		},
		{
			pattern:      "j",
			rightContext: regexp.MustCompile("^[aoeiuy]"),
			phonetic:     "(j|dZ[4]|x[1024]|Z[264])",
		},
		{
			pattern:  "j",
			phonetic: "(j|x[1024])",
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
		},
		{
			pattern:  "ph",
			phonetic: "(ph|f)",
		},
		{
			pattern:  "qu",
			phonetic: "(kv[16]|k)",
		},
		{
			pattern:     "rze",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(Se[128]|re)",
		},
		{
			pattern:  "rze",
			phonetic: "(rze|rtsE[16]|Ze[128]|re[128]|rZe[128])",
		},
		{
			pattern:     "rzy",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(Si[128]|ri)",
		},
		{
			pattern:  "rzy",
			phonetic: "(Zi[128]|ri[128]|rZi)",
		},
		{
			pattern:     "rz",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(S[128]|r)",
		},
		{
			pattern:  "rz",
			phonetic: "(rz|rts[16]|Z[128]|r[128]|rZ[128])",
		},
		{
			pattern:      "tz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ts|tS[20])",
		},
		{
			pattern:     "tz",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(ts|tS[20])",
		},
		{
			pattern:  "tz",
			phonetic: "(ts[532]|tz)",
		},
		{
			pattern:  "zh",
			phonetic: "(Z|zh[128]|tsh[16])",
		},
		{
			pattern:      "zia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(ZB[128]|zB[128]|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za[128]|zja)",
		},
		{
			pattern:      "zią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Zom[128]|zom)",
		},
		{
			pattern:  "zią",
			phonetic: "(Zon[128]|zon)",
		},
		{
			pattern:      "zię",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Zem[128]|zem)",
		},
		{
			pattern:  "zię",
			phonetic: "(Zen[128]|zen)",
		},
		{
			pattern:      "zie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(ZF[128]|zF[128]|ze|tsi[16])",
		},
		{
			pattern:  "zie",
			phonetic: "(ze|Ze[128]|tsi[16])",
		},
		{
			pattern:  "zio",
			phonetic: "(Zo[128]|zo)",
		},
		{
			pattern:  "ziu",
			phonetic: "(Zu[128]|zju)",
		},
		{
			pattern:  "zi",
			phonetic: "(Zi[128]|zi|tsi[16])",
		},
		{
			pattern:      "thal",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "tal",
		},
		{
			pattern:     "th",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "t",
		},
		{
			pattern:      "th",
			rightContext: regexp.MustCompile("^[aeiou]"),
			phonetic:     "(t[16]|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:  "vogel",
			phonetic: "(vogel|fogel[16])",
		},
		{
			pattern:     "v",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(v|f[16])",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("[aeiouyäöü]$"),
			phonetic:    "",
		},
		{
			pattern:  "h",
			phonetic: "(h|x[384])",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(h|H[20])",
		},
		{
			pattern:     "yi",
			leftContext: regexp.MustCompile(" $"),
			phonetic:    "i",
		},
		{
			pattern:      "ii",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "iy",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "yy",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "e",
			leftContext:  regexp.MustCompile("in$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(e|[8])",
		},
		{
			pattern:      "yj",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:      "ij",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "oue",
			phonetic: "oue",
		},
		{
			pattern:  "au",
			phonetic: "(au|o[8])",
		},
		{
			pattern:  "ou",
			phonetic: "(ou|u[8])",
		},
		{
			pattern:  "ue",
			phonetic: "(Q|uje[512])",
		},
		{
			pattern:  "ae",
			phonetic: "(Y[16]|aje[512]|ae)",
		},
		{
			pattern:  "oe",
			phonetic: "(Y[16]|oje[512]|oe)",
		},
		{
			pattern:  "ee",
			phonetic: "(i[4]|aje[512]|e)",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "eu",
			phonetic: "(aj[16]|oj[16]|eu)",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "ie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(i[16]|e[128]|ije[512]|je)",
		},
		{
			pattern:  "ie",
			phonetic: "(i[16]|e[128]|ije[512]|je)",
		},
		{
			pattern:  "ye",
			phonetic: "(je|ije[512])",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "j",
		},
		{
			pattern:  "io",
			phonetic: "(jo|e[512])",
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e[512])",
		},
		{
			pattern:  "ea",
			phonetic: "(ea|ja[256])",
		},
		{
			pattern:     "e",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(e|je[512])",
		},
		{
			pattern:  "oo",
			phonetic: "(u[4]|o)",
		},
		{
			pattern:  "uu",
			phonetic: "u",
		},
		{
			pattern:  "ć",
			phonetic: "(tS[128]|ts)",
		},
		{
			pattern:  "ł",
			phonetic: "l",
		},
		{
			pattern:  "ń",
			phonetic: "n",
		},
		{
			pattern:  "ñ",
			phonetic: "(n|nj[1024])",
		},
		{
			pattern:  "ś",
			phonetic: "(S[128]|s)",
		},
		{
			pattern:  "ş",
			phonetic: "S",
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
		},
		{
			pattern:  "ż",
			phonetic: "Z",
		},
		{
			pattern:  "ź",
			phonetic: "(Z[128]|z)",
		},
		{
			pattern:  "où",
			phonetic: "u",
		},
		{
			pattern:      "ą",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "om",
		},
		{
			pattern:  "ą",
			phonetic: "on",
		},
		{
			pattern:  "ä",
			phonetic: "(Y|e)",
		},
		{
			pattern:  "á",
			phonetic: "a",
		},
		{
			pattern:  "ă",
			phonetic: "(e[256]|a)",
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
			pattern:  "é",
			phonetic: "e",
		},
		{
			pattern:  "è",
			phonetic: "e",
		},
		{
			pattern:  "ê",
			phonetic: "e",
		},
		{
			pattern:      "ę",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "em",
		},
		{
			pattern:  "ę",
			phonetic: "en",
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
			pattern:  "ö",
			phonetic: "Y",
		},
		{
			pattern:  "ő",
			phonetic: "Y",
		},
		{
			pattern:  "ó",
			phonetic: "(u[128]|o)",
		},
		{
			pattern:  "ű",
			phonetic: "Q",
		},
		{
			pattern:  "ü",
			phonetic: "Q",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "ű",
			phonetic: "Q",
		},
		{
			pattern:  "ß",
			phonetic: "s",
		},
		{
			pattern:  "'",
			phonetic: "",
		},
		{
			pattern:  "\"",
			phonetic: "",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(A|B[128])",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(E|F[128])",
		},
		{
			pattern:      "o",
			rightContext: regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			phonetic:     "(O|P[128])",
		},
		{
			pattern:  "a",
			phonetic: "A",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "(k|ts[128])",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
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
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
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
			phonetic: "O",
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
			phonetic: "U",
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
			phonetic: "(ts[16]|z)",
		},
	},
	ashcyrillic: []rule{
		{
			pattern:  "ця",
			phonetic: "tsa",
		},
		{
			pattern:  "цю",
			phonetic: "tsu",
		},
		{
			pattern:  "циа",
			phonetic: "tsa",
		},
		{
			pattern:  "цие",
			phonetic: "tse",
		},
		{
			pattern:  "цио",
			phonetic: "tso",
		},
		{
			pattern:  "циу",
			phonetic: "tsu",
		},
		{
			pattern:  "сие",
			phonetic: "se",
		},
		{
			pattern:  "сио",
			phonetic: "so",
		},
		{
			pattern:  "зие",
			phonetic: "ze",
		},
		{
			pattern:  "зио",
			phonetic: "zo",
		},
		{
			pattern:      "гауз",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "haus",
		},
		{
			pattern:      "гаус",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "haus",
		},
		{
			pattern:      "гольц",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "геймер",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "гейм",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "гоф",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hof",
		},
		{
			pattern:      "гер",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ger",
		},
		{
			pattern:      "ген",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "gen",
		},
		{
			pattern:      "гин",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "gin",
		},
		{
			pattern:      "г",
			leftContext:  regexp.MustCompile("(й|ё|я|ю|ы|а|е|о|и|у)$"),
			rightContext: regexp.MustCompile("^(а|е|о|и|у)"),
			phonetic:     "g",
		},
		{
			pattern:      "г",
			rightContext: regexp.MustCompile("^(а|е|о|и|у)"),
			phonetic:     "(g|h)",
		},
		{
			pattern:  "ля",
			phonetic: "la",
		},
		{
			pattern:  "лю",
			phonetic: "lu",
		},
		{
			pattern:  "лё",
			phonetic: "(le|lo)",
		},
		{
			pattern:  "лио",
			phonetic: "(le|lo)",
		},
		{
			pattern:  "ле",
			phonetic: "(lE|lo)",
		},
		{
			pattern:  "ийе",
			phonetic: "je",
		},
		{
			pattern:  "ие",
			phonetic: "je",
		},
		{
			pattern:  "ыйе",
			phonetic: "je",
		},
		{
			pattern:  "ые",
			phonetic: "je",
		},
		{
			pattern:      "ий",
			rightContext: regexp.MustCompile("^(а|о|у)"),
			phonetic:     "j",
		},
		{
			pattern:      "ый",
			rightContext: regexp.MustCompile("^(а|о|у)"),
			phonetic:     "j",
		},
		{
			pattern:      "ий",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:      "ый",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:  "ё",
			phonetic: "(e|jo)",
		},
		{
			pattern:     "ей",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(jaj|aj)",
		},
		{
			pattern:     "е",
			leftContext: regexp.MustCompile("(а|е|о|у)$"),
			phonetic:    "je",
		},
		{
			pattern:     "е",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "je",
		},
		{
			pattern:  "эй",
			phonetic: "aj",
		},
		{
			pattern:  "ей",
			phonetic: "aj",
		},
		{
			pattern:  "ауе",
			phonetic: "aue",
		},
		{
			pattern:  "ауэ",
			phonetic: "aue",
		},
		{
			pattern:  "а",
			phonetic: "a",
		},
		{
			pattern:  "б",
			phonetic: "b",
		},
		{
			pattern:  "в",
			phonetic: "v",
		},
		{
			pattern:  "г",
			phonetic: "g",
		},
		{
			pattern:  "д",
			phonetic: "d",
		},
		{
			pattern:  "е",
			phonetic: "E",
		},
		{
			pattern:  "ж",
			phonetic: "Z",
		},
		{
			pattern:  "з",
			phonetic: "z",
		},
		{
			pattern:  "и",
			phonetic: "I",
		},
		{
			pattern:  "й",
			phonetic: "j",
		},
		{
			pattern:  "к",
			phonetic: "k",
		},
		{
			pattern:  "л",
			phonetic: "l",
		},
		{
			pattern:  "м",
			phonetic: "m",
		},
		{
			pattern:  "н",
			phonetic: "n",
		},
		{
			pattern:  "о",
			phonetic: "o",
		},
		{
			pattern:  "п",
			phonetic: "p",
		},
		{
			pattern:  "р",
			phonetic: "r",
		},
		{
			pattern:      "с",
			rightContext: regexp.MustCompile("^с"),
			phonetic:     "",
		},
		{
			pattern:  "с",
			phonetic: "s",
		},
		{
			pattern:  "т",
			phonetic: "t",
		},
		{
			pattern:  "у",
			phonetic: "u",
		},
		{
			pattern:  "ф",
			phonetic: "f",
		},
		{
			pattern:  "х",
			phonetic: "x",
		},
		{
			pattern:  "ц",
			phonetic: "ts",
		},
		{
			pattern:  "ч",
			phonetic: "tS",
		},
		{
			pattern:  "ш",
			phonetic: "S",
		},
		{
			pattern:  "щ",
			phonetic: "StS",
		},
		{
			pattern:  "ъ",
			phonetic: "",
		},
		{
			pattern:  "ы",
			phonetic: "I",
		},
		{
			pattern:  "ь",
			phonetic: "",
		},
		{
			pattern:  "э",
			phonetic: "E",
		},
		{
			pattern:  "ю",
			phonetic: "ju",
		},
		{
			pattern:  "я",
			phonetic: "ja",
		},
	},
	ashenglish: []rule{
		{
			pattern:  "tch",
			phonetic: "tS",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
		},
		{
			pattern:  "ck",
			phonetic: "k",
		},
		{
			pattern:      "cc",
			rightContext: regexp.MustCompile("^[iey]"),
			phonetic:     "ks",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^c"),
			phonetic:     "",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[iey]"),
			phonetic:     "s",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:     "gh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "g",
		},
		{
			pattern:  "gh",
			phonetic: "(g|f|w)",
		},
		{
			pattern:  "gn",
			phonetic: "(gn|n)",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[iey]"),
			phonetic:     "(g|dZ)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ph",
			phonetic: "f",
		},
		{
			pattern:  "sch",
			phonetic: "(S|sk)",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:     "who",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "hu",
		},
		{
			pattern:     "wh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "w",
		},
		{
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
		},
		{
			pattern:      "h",
			rightContext: regexp.MustCompile("^[^aeiou]"),
			phonetic:     "",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "H",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "j",
			phonetic: "dZ",
		},
		{
			pattern:     "kn",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "n",
		},
		{
			pattern:      "mb",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "m",
		},
		{
			pattern:      "ng",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(N|ng)",
		},
		{
			pattern:     "pn",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(pn|n)",
		},
		{
			pattern:     "ps",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(ps|s)",
		},
		{
			pattern:  "qu",
			phonetic: "kw",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "tia",
			phonetic: "(So|Sa)",
		},
		{
			pattern:  "tio",
			phonetic: "So",
		},
		{
			pattern:     "wr",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "r",
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
		},
		{
			pattern:     "x",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "z",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:     "yi",
			leftContext: regexp.MustCompile(" $"),
			phonetic:    "i",
		},
		{
			pattern:      "y",
			leftContext:  regexp.MustCompile("^$"),
			rightContext: regexp.MustCompile("^[aeiouy]"),
			phonetic:     "j",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "oue",
			phonetic: "(aue|oue)",
		},
		{
			pattern:  "ai",
			phonetic: "(aj|e)",
		},
		{
			pattern:  "ay",
			phonetic: "aj",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "aj",
		},
		{
			pattern:  "a",
			phonetic: "(e|o|a)",
		},
		{
			pattern:  "ei",
			phonetic: "(aj|i)",
		},
		{
			pattern:  "ey",
			phonetic: "(aj|i)",
		},
		{
			pattern:  "ear",
			phonetic: "ia",
		},
		{
			pattern:  "ea",
			phonetic: "(i|e)",
		},
		{
			pattern:  "ee",
			phonetic: "i",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "i",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(|E)",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "ie",
			phonetic: "i",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "aj",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "oa",
			phonetic: "ou",
		},
		{
			pattern:  "oi",
			phonetic: "oj",
		},
		{
			pattern:  "oo",
			phonetic: "u",
		},
		{
			pattern:  "ou",
			phonetic: "(u|ou)",
		},
		{
			pattern:  "oy",
			phonetic: "oj",
		},
		{
			pattern:      "o",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "ou",
		},
		{
			pattern:  "o",
			phonetic: "(o|a)",
		},
		{
			pattern:      "u",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "(ju|u)",
		},
		{
			pattern:      "u",
			rightContext: regexp.MustCompile("^r"),
			phonetic:     "(e|u)",
		},
		{
			pattern:  "u",
			phonetic: "(u|a)",
		},
		{
			pattern:  "y",
			phonetic: "i",
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
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
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
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashfrench: []rule{
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
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
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
			pattern:  "yi",
			phonetic: "i",
		},
		{
			pattern:  "ii",
			phonetic: "i",
		},
		{
			pattern:  "yy",
			phonetic: "i",
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
			phonetic: "E",
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
			phonetic: "I",
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
	ashgerman: []rule{
		{
			pattern:  "ziu",
			phonetic: "tsu",
		},
		{
			pattern:  "zia",
			phonetic: "tsa",
		},
		{
			pattern:  "zio",
			phonetic: "tso",
		},
		{
			pattern:  "ssch",
			phonetic: "S",
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
		},
		{
			pattern:      "ewitsch",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "evitS",
		},
		{
			pattern:      "owitsch",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ovitS",
		},
		{
			pattern:      "evitsch",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "evitS",
		},
		{
			pattern:      "ovitsch",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ovitS",
		},
		{
			pattern:      "witsch",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "vitS",
		},
		{
			pattern:      "vitsch",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "vitS",
		},
		{
			pattern:  "sch",
			phonetic: "S",
		},
		{
			pattern:  "chs",
			phonetic: "ks",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:  "ck",
			phonetic: "k",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[eiy]"),
			phonetic:     "ts",
		},
		{
			pattern:     "sp",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "Sp",
		},
		{
			pattern:     "st",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "St",
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp|sp)",
		},
		{
			pattern:  "sp",
			phonetic: "(Sp|sp)",
		},
		{
			pattern:  "sst",
			phonetic: "(St|st)",
		},
		{
			pattern:  "st",
			phonetic: "(St|st)",
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
		},
		{
			pattern:  "ph",
			phonetic: "(ph|f)",
		},
		{
			pattern:  "qu",
			phonetic: "kv",
		},
		{
			pattern:      "ewitz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(evits|evitS)",
		},
		{
			pattern:      "ewiz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(evits|evitS)",
		},
		{
			pattern:      "evitz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(evits|evitS)",
		},
		{
			pattern:      "eviz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(evits|evitS)",
		},
		{
			pattern:      "owitz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovits|ovitS)",
		},
		{
			pattern:      "owiz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovits|ovitS)",
		},
		{
			pattern:      "ovitz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovits|ovitS)",
		},
		{
			pattern:      "oviz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovits|ovitS)",
		},
		{
			pattern:      "witz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(vits|vitS)",
		},
		{
			pattern:      "wiz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(vits|vitS)",
		},
		{
			pattern:      "vitz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(vits|vitS)",
		},
		{
			pattern:      "viz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(vits|vitS)",
		},
		{
			pattern:  "tz",
			phonetic: "ts",
		},
		{
			pattern:      "thal",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "tal",
		},
		{
			pattern:     "th",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "t",
		},
		{
			pattern:      "th",
			rightContext: regexp.MustCompile("^[äöüaeiou]"),
			phonetic:     "(t|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:     "rh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "r",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("[aeiouyäöü]$"),
			phonetic:    "",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "H",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[äöüaeiouy]"),
			phonetic:     "(z|s)",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aeiouyäöüj]$"),
			rightContext: regexp.MustCompile("^[aeiouyäöü]"),
			phonetic:     "z",
		},
		{
			pattern:  "ß",
			phonetic: "s",
		},
		{
			pattern:      "ij",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "ue",
			phonetic: "Q",
		},
		{
			pattern:  "ae",
			phonetic: "Y",
		},
		{
			pattern:  "oe",
			phonetic: "Y",
		},
		{
			pattern:  "ü",
			phonetic: "Q",
		},
		{
			pattern:  "ä",
			phonetic: "(Y|e)",
		},
		{
			pattern:  "ö",
			phonetic: "Y",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "eu",
			phonetic: "(aj|oj)",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:  "ie",
			phonetic: "I",
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
			pattern:  "ñ",
			phonetic: "n",
		},
		{
			pattern:  "ã",
			phonetic: "a",
		},
		{
			pattern:  "ő",
			phonetic: "o",
		},
		{
			pattern:  "ű",
			phonetic: "u",
		},
		{
			pattern:  "ç",
			phonetic: "s",
		},
		{
			pattern:  "a",
			phonetic: "A",
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
			phonetic: "E",
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
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
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
			phonetic: "O",
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
			phonetic: "U",
		},
		{
			pattern:  "v",
			phonetic: "(f|v)",
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
			phonetic: "ts",
		},
	},
	ashhebrew: []rule{
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
			phonetic: "TB",
		},
	},
	ashhungarian: []rule{
		{
			pattern:  "sz",
			phonetic: "s",
		},
		{
			pattern:  "zs",
			phonetic: "Z",
		},
		{
			pattern:  "cs",
			phonetic: "tS",
		},
		{
			pattern:  "ay",
			phonetic: "(oj|aj)",
		},
		{
			pattern:  "ai",
			phonetic: "(oj|aj)",
		},
		{
			pattern:  "aj",
			phonetic: "(oj|aj)",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[áo]$"),
			phonetic:    "j",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[áo]$"),
			phonetic:    "j",
		},
		{
			pattern:  "ee",
			phonetic: "(aj|e)",
		},
		{
			pattern:  "ely",
			phonetic: "(aj|eli)",
		},
		{
			pattern:  "ly",
			phonetic: "(j|li)",
		},
		{
			pattern:      "gy",
			rightContext: regexp.MustCompile("^[aeouáéóúüöőű]"),
			phonetic:     "dj",
		},
		{
			pattern:  "gy",
			phonetic: "(d|gi)",
		},
		{
			pattern:      "ny",
			rightContext: regexp.MustCompile("^[aeouáéóúüöőű]"),
			phonetic:     "nj",
		},
		{
			pattern:  "ny",
			phonetic: "(n|ni)",
		},
		{
			pattern:      "ty",
			rightContext: regexp.MustCompile("^[aeouáéóúüöőű]"),
			phonetic:     "tj",
		},
		{
			pattern:  "ty",
			phonetic: "(t|ti)",
		},
		{
			pattern:  "qu",
			phonetic: "(ku|kv)",
		},
		{
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
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
			pattern:  "ö",
			phonetic: "Y",
		},
		{
			pattern:  "ő",
			phonetic: "Y",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "ü",
			phonetic: "Q",
		},
		{
			pattern:  "ű",
			phonetic: "Q",
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
			phonetic: "ts",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
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
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
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
			phonetic: "(S|s)",
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
			phonetic: "z",
		},
	},
	ashpolish: []rule{
		{
			pattern:      "ska",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ski",
		},
		{
			pattern:      "cka",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "tski",
		},
		{
			pattern:      "lowa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lova|lof|l|el)",
		},
		{
			pattern:      "kowa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kova|kof|k|ek)",
		},
		{
			pattern:      "owa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ova|of|)",
		},
		{
			pattern:      "lowna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lovna|levna|l|el)",
		},
		{
			pattern:      "kowna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kovna|k|ek)",
		},
		{
			pattern:      "owna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovna|)",
		},
		{
			pattern:      "lówna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(l|el)",
		},
		{
			pattern:      "kówna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k|ek)",
		},
		{
			pattern:      "ówna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(a|i)",
		},
		{
			pattern:  "czy",
			phonetic: "tSi",
		},
		{
			pattern:      "cze",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tSe|tSF)",
		},
		{
			pattern:  "ciewicz",
			phonetic: "(tsevitS|tSevitS)",
		},
		{
			pattern:  "siewicz",
			phonetic: "(sevitS|SevitS)",
		},
		{
			pattern:  "ziewicz",
			phonetic: "(zevitS|ZevitS)",
		},
		{
			pattern:  "riewicz",
			phonetic: "rjevitS",
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
		},
		{
			pattern:  "icz",
			phonetic: "itS",
		},
		{
			pattern:  "cz",
			phonetic: "tS",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:      "cia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tSB|tsB)",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa|tsa)",
		},
		{
			pattern:      "cią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(tSom|tsom)",
		},
		{
			pattern:  "cią",
			phonetic: "(tSon|tson)",
		},
		{
			pattern:      "cię",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(tSem|tsem)",
		},
		{
			pattern:  "cię",
			phonetic: "(tSen|tsen)",
		},
		{
			pattern:      "cie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(tSF|tsF)",
		},
		{
			pattern:  "cie",
			phonetic: "(tSe|tse)",
		},
		{
			pattern:  "cio",
			phonetic: "(tSo|tso)",
		},
		{
			pattern:  "ciu",
			phonetic: "(tSu|tsu)",
		},
		{
			pattern:  "ci",
			phonetic: "(tSi|tsI)",
		},
		{
			pattern:  "ć",
			phonetic: "(tS|ts)",
		},
		{
			pattern:  "ssz",
			phonetic: "S",
		},
		{
			pattern:  "sz",
			phonetic: "S",
		},
		{
			pattern:      "sia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(SB|sB|sja)",
		},
		{
			pattern:  "sia",
			phonetic: "(Sa|sja)",
		},
		{
			pattern:      "sią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Som|som)",
		},
		{
			pattern:  "sią",
			phonetic: "(Son|son)",
		},
		{
			pattern:      "się",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Sem|sem)",
		},
		{
			pattern:  "się",
			phonetic: "(Sen|sen)",
		},
		{
			pattern:      "sie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(SF|sF|se)",
		},
		{
			pattern:  "sie",
			phonetic: "(Se|se)",
		},
		{
			pattern:  "sio",
			phonetic: "(So|so)",
		},
		{
			pattern:  "siu",
			phonetic: "(Su|sju)",
		},
		{
			pattern:  "si",
			phonetic: "(Si|sI)",
		},
		{
			pattern:  "ś",
			phonetic: "(S|s)",
		},
		{
			pattern:      "zia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(ZB|zB|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za|zja)",
		},
		{
			pattern:      "zią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Zom|zom)",
		},
		{
			pattern:  "zią",
			phonetic: "(Zon|zon)",
		},
		{
			pattern:      "zię",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Zem|zem)",
		},
		{
			pattern:  "zię",
			phonetic: "(Zen|zen)",
		},
		{
			pattern:      "zie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(ZF|zF)",
		},
		{
			pattern:  "zie",
			phonetic: "(Ze|ze)",
		},
		{
			pattern:  "zio",
			phonetic: "(Zo|zo)",
		},
		{
			pattern:  "ziu",
			phonetic: "(Zu|zju)",
		},
		{
			pattern:  "zi",
			phonetic: "(Zi|zI)",
		},
		{
			pattern:      "że",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(Ze|ZF)",
		},
		{
			pattern:      "że",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(Ze|ZF|ze|zF)",
		},
		{
			pattern:  "że",
			phonetic: "Ze",
		},
		{
			pattern:  "źe",
			phonetic: "(Ze|ze)",
		},
		{
			pattern:  "ży",
			phonetic: "Zi",
		},
		{
			pattern:  "źi",
			phonetic: "(Zi|zi)",
		},
		{
			pattern:  "ż",
			phonetic: "Z",
		},
		{
			pattern:  "ź",
			phonetic: "(Z|z)",
		},
		{
			pattern:     "rze",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(Se|re)",
		},
		{
			pattern:  "rze",
			phonetic: "(Ze|re|rZe)",
		},
		{
			pattern:     "rzy",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(Si|ri)",
		},
		{
			pattern:  "rzy",
			phonetic: "(Zi|ri|rZi)",
		},
		{
			pattern:     "rz",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(S|r)",
		},
		{
			pattern:  "rz",
			phonetic: "(Z|r|rZ)",
		},
		{
			pattern:  "lio",
			phonetic: "(lo|le)",
		},
		{
			pattern:  "ł",
			phonetic: "l",
		},
		{
			pattern:  "ń",
			phonetic: "n",
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^s"),
			phonetic:     "",
		},
		{
			pattern:  "ó",
			phonetic: "(u|o)",
		},
		{
			pattern:      "ą",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "om",
		},
		{
			pattern:      "ę",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "em",
		},
		{
			pattern:  "ą",
			phonetic: "on",
		},
		{
			pattern:  "ę",
			phonetic: "en",
		},
		{
			pattern:  "ije",
			phonetic: "je",
		},
		{
			pattern:  "yje",
			phonetic: "je",
		},
		{
			pattern:  "iie",
			phonetic: "je",
		},
		{
			pattern:  "yie",
			phonetic: "je",
		},
		{
			pattern:  "iye",
			phonetic: "je",
		},
		{
			pattern:  "yye",
			phonetic: "je",
		},
		{
			pattern:      "ij",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "yj",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "ii",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "yi",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "iy",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "yy",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:  "rie",
			phonetic: "rje",
		},
		{
			pattern:  "die",
			phonetic: "dje",
		},
		{
			pattern:  "tie",
			phonetic: "tje",
		},
		{
			pattern:      "ie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "F",
		},
		{
			pattern:  "ie",
			phonetic: "e",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "au",
			phonetic: "au",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "ej",
			phonetic: "aj",
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
			pattern:  "aj",
			phonetic: "aj",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[ou]$"),
			phonetic:    "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[ou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "B",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(E|F)",
		},
		{
			pattern:      "o",
			rightContext: regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			phonetic:     "P",
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
			phonetic: "ts",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
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
			phonetic: "(h|x)",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
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
			phonetic: "I",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashromanian: []rule{
		{
			pattern:  "j",
			phonetic: "Z",
		},
		{
			pattern:  "ce",
			phonetic: "tSe",
		},
		{
			pattern:  "ci",
			phonetic: "(tSi|tS)",
		},
		{
			pattern:      "ch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "k",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "gi",
			phonetic: "(dZi|dZ)",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "dZ",
		},
		{
			pattern:  "gh",
			phonetic: "g",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
		},
		{
			pattern:  "ş",
			phonetic: "S",
		},
		{
			pattern:  "h",
			phonetic: "(x|h)",
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
			pattern:  "î",
			phonetic: "i",
		},
		{
			pattern:  "ea",
			phonetic: "ja",
		},
		{
			pattern:  "ă",
			phonetic: "(e|a)",
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
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
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
			pattern:  "i",
			phonetic: "I",
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
	ashrussian: []rule{
		{
			pattern:      "yna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(in|ina)",
		},
		{
			pattern:      "ina",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(in|ina)",
		},
		{
			pattern:      "liova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lof|lef)",
		},
		{
			pattern:      "lova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lof|lef|lova)",
		},
		{
			pattern:      "ova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(of|ova)",
		},
		{
			pattern:      "eva",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ef|ova)",
		},
		{
			pattern:      "aia",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i)",
		},
		{
			pattern:      "aja",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i)",
		},
		{
			pattern:      "aya",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i)",
		},
		{
			pattern:  "tsya",
			phonetic: "tsa",
		},
		{
			pattern:  "tsyu",
			phonetic: "tsu",
		},
		{
			pattern:  "tsia",
			phonetic: "tsa",
		},
		{
			pattern:  "tsie",
			phonetic: "tse",
		},
		{
			pattern:  "tsio",
			phonetic: "tso",
		},
		{
			pattern:  "tsye",
			phonetic: "tse",
		},
		{
			pattern:  "tsyo",
			phonetic: "tso",
		},
		{
			pattern:  "tsiu",
			phonetic: "tsu",
		},
		{
			pattern:  "sie",
			phonetic: "se",
		},
		{
			pattern:  "sio",
			phonetic: "so",
		},
		{
			pattern:  "zie",
			phonetic: "ze",
		},
		{
			pattern:  "zio",
			phonetic: "zo",
		},
		{
			pattern:  "sye",
			phonetic: "se",
		},
		{
			pattern:  "syo",
			phonetic: "so",
		},
		{
			pattern:  "zye",
			phonetic: "ze",
		},
		{
			pattern:  "zyo",
			phonetic: "zo",
		},
		{
			pattern:      "gauz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "haus",
		},
		{
			pattern:      "gaus",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "haus",
		},
		{
			pattern:      "gol'ts",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "golts",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "gol'tz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "goltz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "holts",
		},
		{
			pattern:      "gejmer",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "gejm",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "geimer",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "geim",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "geymer",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajmer",
		},
		{
			pattern:      "geym",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hajm",
		},
		{
			pattern:      "gendler",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hendler",
		},
		{
			pattern:      "gof",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hof",
		},
		{
			pattern:      "gojf",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hojf",
		},
		{
			pattern:      "goyf",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hojf",
		},
		{
			pattern:      "goif",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "hojf",
		},
		{
			pattern:      "ger",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ger",
		},
		{
			pattern:      "gen",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "gen",
		},
		{
			pattern:      "gin",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "gin",
		},
		{
			pattern:  "gg",
			phonetic: "g",
		},
		{
			pattern:      "kog",
			leftContext:  regexp.MustCompile("^$"),
			rightContext: regexp.MustCompile("^[aeoiu]"),
			phonetic:     "(kog|koh)",
		},
		{
			pattern:      "kag",
			leftContext:  regexp.MustCompile("^$"),
			rightContext: regexp.MustCompile("^[aeoiu]"),
			phonetic:     "(kag|kah)",
		},
		{
			pattern:      "g",
			leftContext:  regexp.MustCompile("[jaeoiuy]$"),
			rightContext: regexp.MustCompile("^[aeoiu]"),
			phonetic:     "g",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[aeoiu]"),
			phonetic:     "(g|h)",
		},
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
		},
		{
			pattern:  "sch",
			phonetic: "(StS|S)",
		},
		{
			pattern:  "ssh",
			phonetic: "S",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:  "zh",
			phonetic: "Z",
		},
		{
			pattern:      "tz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ts",
		},
		{
			pattern:  "tz",
			phonetic: "(ts|tz)",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[iey]"),
			phonetic:     "s",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^s"),
			phonetic:     "",
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
			pattern:  "lya",
			phonetic: "la",
		},
		{
			pattern:  "lyu",
			phonetic: "lu",
		},
		{
			pattern:  "lia",
			phonetic: "la",
		},
		{
			pattern:  "liu",
			phonetic: "lu",
		},
		{
			pattern:  "lja",
			phonetic: "la",
		},
		{
			pattern:  "lju",
			phonetic: "lu",
		},
		{
			pattern:  "le",
			phonetic: "(lo|lE)",
		},
		{
			pattern:  "lyo",
			phonetic: "(lo|le)",
		},
		{
			pattern:  "lio",
			phonetic: "(lo|le)",
		},
		{
			pattern:  "ije",
			phonetic: "je",
		},
		{
			pattern:  "ie",
			phonetic: "je",
		},
		{
			pattern:  "iye",
			phonetic: "je",
		},
		{
			pattern:  "iie",
			phonetic: "je",
		},
		{
			pattern:  "yje",
			phonetic: "je",
		},
		{
			pattern:  "ye",
			phonetic: "je",
		},
		{
			pattern:  "yye",
			phonetic: "je",
		},
		{
			pattern:  "yie",
			phonetic: "je",
		},
		{
			pattern:      "ij",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "iy",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "ii",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "yj",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "yy",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:      "yi",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "j",
		},
		{
			pattern:  "io",
			phonetic: "(jo|e)",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "j",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "ej",
			phonetic: "aj",
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e)",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[au]"),
			phonetic:     "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aiou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "ii",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "iy",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "yy",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "yi",
			rightContext: regexp.MustCompile("^ "),
			phonetic:     "i",
		},
		{
			pattern:      "yj",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:      "ij",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:     "e",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(je|E)",
		},
		{
			pattern:  "ee",
			phonetic: "(aje|i)",
		},
		{
			pattern:     "e",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "je",
		},
		{
			pattern:  "y",
			phonetic: "I",
		},
		{
			pattern:  "oo",
			phonetic: "(oo|u)",
		},
		{
			pattern:  "'",
			phonetic: "",
		},
		{
			pattern:  "\"",
			phonetic: "",
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
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
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
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
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
	ashspanish: []rule{
		{
			pattern:  "ñ",
			phonetic: "(n|nj)",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|dZ)",
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
			phonetic: "x",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "ll",
			phonetic: "(l|Z)",
		},
		{
			pattern:  "w",
			phonetic: "v",
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
			phonetic:     "(x|g)",
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
			phonetic: "(i|j|S|Z)",
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
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
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
			phonetic: "I",
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

var ashLangRules = []langRule{
	{
		match: ruleMatch{
			contains: "zh",
			prefix:   "",
			suffix:   "",
		},
		langs:  660,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "eau",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[aoeiuäöü]h"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "vogel",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "vogel",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "witz",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "tz",
		},
		langs:  532,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "tz",
			suffix:   "",
		},
		langs:  516,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "güe",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "güi",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ghe",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ghi",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "vici",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "schi",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "chsch",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "tsch",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ssch",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "sch",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "sch",
			suffix:   "",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "rz",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "rz",
		},
		langs:  144,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[^aoeiuäöü]rz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("rz[^aoeiuäöü]"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "cki",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "ska",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "cka",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ue",
			prefix:   "",
			suffix:   "",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ae",
			prefix:   "",
			suffix:   "",
		},
		langs:  532,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "oe",
			prefix:   "",
			suffix:   "",
		},
		langs:  540,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "th",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "th",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("th[^aoeiu]"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "mann",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "cz",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "cy",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "niew",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "stein",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "heim",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "heimer",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "ii",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "iy",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "yy",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "yi",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "yj",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "ij",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gaus",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gauz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gauz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "goltz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gol'tz$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "golts",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gol'ts$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "goltz",
			suffix:   "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("^gol'tz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "golts",
			suffix:   "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("^gol'ts"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gendler",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gejmer",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gejm",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "geimer",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "geim",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "geymer",
			prefix:   "",
			suffix:   "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "geym",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gof",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "thal",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "zweig",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "ck",
		},
		langs:  20,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "c",
		},
		langs:  448,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "sz",
			prefix:   "",
			suffix:   "",
		},
		langs:  192,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gue",
			prefix:   "",
			suffix:   "",
		},
		langs:  1032,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gui",
			prefix:   "",
			suffix:   "",
		},
		langs:  1032,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "guy",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "cs",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "cs",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "dzs",
			prefix:   "",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "zs",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "zs",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "wl",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "wr",
			suffix:   "",
		},
		langs:  148,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "gy",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gy[aeou]"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "gy",
			prefix:   "",
			suffix:   "",
		},
		langs:  576,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ly",
			prefix:   "",
			suffix:   "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ny",
			prefix:   "",
			suffix:   "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ty",
			prefix:   "",
			suffix:   "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "â",
			prefix:   "",
			suffix:   "",
		},
		langs:  264,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ă",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "à",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ä",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "á",
			prefix:   "",
			suffix:   "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ą",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ć",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ç",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ę",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "é",
			prefix:   "",
			suffix:   "",
		},
		langs:  1096,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "è",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ê",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "í",
			prefix:   "",
			suffix:   "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "î",
			prefix:   "",
			suffix:   "",
		},
		langs:  264,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ł",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ń",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ñ",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ó",
			prefix:   "",
			suffix:   "",
		},
		langs:  1216,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ö",
			prefix:   "",
			suffix:   "",
		},
		langs:  80,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "õ",
			prefix:   "",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ş",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ś",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ţ",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ü",
			prefix:   "",
			suffix:   "",
		},
		langs:  80,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ù",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ű",
			prefix:   "",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ú",
			prefix:   "",
			suffix:   "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ź",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ż",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ß",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "а",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ё",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "о",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "е",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "и",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "у",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ы",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "э",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ю",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "я",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "א",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ב",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ג",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ד",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ה",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ו",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ז",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ח",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ט",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "י",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "כ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ל",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "מ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "נ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ס",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ע",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "פ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "צ",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ק",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ר",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ש",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "ת",
			prefix:   "",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatch{
			contains: "a",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "o",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "e",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "i",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "y",
			prefix:   "",
			suffix:   "",
		},
		langs:  290,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "u",
			prefix:   "",
			suffix:   "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("v[^aoeiuäüö]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("y[^aoeiu]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("c[^aohk]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "dzi",
			prefix:   "",
			suffix:   "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "ou",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "aj",
			prefix:   "",
			suffix:   "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "ej",
			prefix:   "",
			suffix:   "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "oj",
			prefix:   "",
			suffix:   "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "uj",
			prefix:   "",
			suffix:   "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "k",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "v",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "ky",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "eu",
			prefix:   "",
			suffix:   "",
		},
		langs:  640,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "w",
			prefix:   "",
			suffix:   "",
		},
		langs:  1864,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "kie",
			prefix:   "",
			suffix:   "",
		},
		langs:  1032,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "gie",
			prefix:   "",
			suffix:   "",
		},
		langs:  1288,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "q",
			prefix:   "",
			suffix:   "",
		},
		langs:  960,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "sch",
			prefix:   "",
			suffix:   "",
		},
		langs:  1224,
		accept: false,
	},
	{
		match: ruleMatch{
			contains: "",
			prefix:   "h",
			suffix:   "",
		},
		langs:  512,
		accept: false,
	},
}

var ashFinalRules = finalRules{
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
				rightContext: regexp.MustCompile("^[gdZz]"),
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
				rightContext: regexp.MustCompile("^[bgdZz]"),
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
				rightContext: regexp.MustCompile("^[bdZz]"),
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
				rightContext: regexp.MustCompile("^[bgZz]"),
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
				pattern:  "jnm",
				phonetic: "jm",
			},
			{
				pattern:     "ji",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "i",
			},
			{
				pattern:     "jI",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "I",
			},
			{
				pattern:      "a",
				rightContext: regexp.MustCompile("^[aAB]"),
				phonetic:     "",
			},
			{
				pattern:     "a",
				leftContext: regexp.MustCompile("[AB]$"),
				phonetic:    "",
			},
			{
				pattern:      "A",
				rightContext: regexp.MustCompile("^A"),
				phonetic:     "",
			},
			{
				pattern:      "B",
				rightContext: regexp.MustCompile("^B"),
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
				pattern:      "n",
				rightContext: regexp.MustCompile("^[bp]"),
				phonetic:     "m",
			},
			{
				pattern:      "kAg",
				leftContext:  regexp.MustCompile("^$"),
				rightContext: regexp.MustCompile("^[AEOIUaeoiu]"),
				phonetic:     "(kOg|kO[512])",
			},
			{
				pattern:      "kOg",
				leftContext:  regexp.MustCompile("^$"),
				rightContext: regexp.MustCompile("^[AEOIUaeoiu]"),
				phonetic:     "(kAg|kA[512])",
			},
			{
				pattern:      "kog",
				leftContext:  regexp.MustCompile("^$"),
				rightContext: regexp.MustCompile("^[AEOIUaeoiu]"),
				phonetic:     "(kog|ko[512])",
			},
			{
				pattern:      "kag",
				leftContext:  regexp.MustCompile("^$"),
				rightContext: regexp.MustCompile("^[AEOIUaeoiu]"),
				phonetic:     "(kag|ka[512])",
			},
			{
				pattern:  "h",
				phonetic: "",
			},
			{
				pattern:  "H",
				phonetic: "(x|)",
			},
			{
				pattern:      "F",
				rightContext: regexp.MustCompile("^[bdgkpstvzZ]h"),
				phonetic:     "e",
			},
			{
				pattern:      "F",
				rightContext: regexp.MustCompile("^[bdgkpstvzZ]x"),
				phonetic:     "e",
			},
			{
				pattern:      "B",
				rightContext: regexp.MustCompile("^[bdgkpstvzZ]h"),
				phonetic:     "a",
			},
			{
				pattern:      "B",
				rightContext: regexp.MustCompile("^[bdgkpstvzZ]x"),
				phonetic:     "a",
			},
			{
				pattern:      "e",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "i",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "E",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "I",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "F",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "Q",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "Y",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "e",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(e|)",
			},
			{
				pattern:      "i",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(i|)",
			},
			{
				pattern:      "E",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(E|)",
			},
			{
				pattern:      "I",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(I|)",
			},
			{
				pattern:      "F",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(F|)",
			},
			{
				pattern:      "Q",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(Q|)",
			},
			{
				pattern:      "Y",
				leftContext:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				phonetic:     "(Y|)",
			},
			{
				pattern:  "lEs",
				phonetic: "(lEs|lz)",
			},
			{
				pattern:     "lE",
				leftContext: regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				phonetic:    "(lE|l)",
			},
			{
				pattern:  "aue",
				phonetic: "D",
			},
			{
				pattern:  "oue",
				phonetic: "D",
			},
			{
				pattern:  "AvE",
				phonetic: "(D|AvE)",
			},
			{
				pattern:  "Ave",
				phonetic: "(D|Ave)",
			},
			{
				pattern:  "avE",
				phonetic: "(D|avE)",
			},
			{
				pattern:  "ave",
				phonetic: "(D|ave)",
			},
			{
				pattern:  "OvE",
				phonetic: "(D|OvE)",
			},
			{
				pattern:  "Ove",
				phonetic: "(D|Ove)",
			},
			{
				pattern:  "ovE",
				phonetic: "(D|ovE)",
			},
			{
				pattern:  "ove",
				phonetic: "(D|ove)",
			},
			{
				pattern:  "ea",
				phonetic: "(D|ea)",
			},
			{
				pattern:  "EA",
				phonetic: "(D|EA)",
			},
			{
				pattern:  "Ea",
				phonetic: "(D|Ea)",
			},
			{
				pattern:  "eA",
				phonetic: "(D|eA)",
			},
			{
				pattern:  "aji",
				phonetic: "D",
			},
			{
				pattern:  "ajI",
				phonetic: "D",
			},
			{
				pattern:  "aje",
				phonetic: "D",
			},
			{
				pattern:  "ajE",
				phonetic: "D",
			},
			{
				pattern:  "Aji",
				phonetic: "D",
			},
			{
				pattern:  "AjI",
				phonetic: "D",
			},
			{
				pattern:  "Aje",
				phonetic: "D",
			},
			{
				pattern:  "AjE",
				phonetic: "D",
			},
			{
				pattern:  "oji",
				phonetic: "D",
			},
			{
				pattern:  "ojI",
				phonetic: "D",
			},
			{
				pattern:  "oje",
				phonetic: "D",
			},
			{
				pattern:  "ojE",
				phonetic: "D",
			},
			{
				pattern:  "Oji",
				phonetic: "D",
			},
			{
				pattern:  "OjI",
				phonetic: "D",
			},
			{
				pattern:  "Oje",
				phonetic: "D",
			},
			{
				pattern:  "OjE",
				phonetic: "D",
			},
			{
				pattern:  "eji",
				phonetic: "D",
			},
			{
				pattern:  "ejI",
				phonetic: "D",
			},
			{
				pattern:  "eje",
				phonetic: "D",
			},
			{
				pattern:  "ejE",
				phonetic: "D",
			},
			{
				pattern:  "Eji",
				phonetic: "D",
			},
			{
				pattern:  "EjI",
				phonetic: "D",
			},
			{
				pattern:  "Eje",
				phonetic: "D",
			},
			{
				pattern:  "EjE",
				phonetic: "D",
			},
			{
				pattern:  "uji",
				phonetic: "D",
			},
			{
				pattern:  "ujI",
				phonetic: "D",
			},
			{
				pattern:  "uje",
				phonetic: "D",
			},
			{
				pattern:  "ujE",
				phonetic: "D",
			},
			{
				pattern:  "Uji",
				phonetic: "D",
			},
			{
				pattern:  "UjI",
				phonetic: "D",
			},
			{
				pattern:  "Uje",
				phonetic: "D",
			},
			{
				pattern:  "UjE",
				phonetic: "D",
			},
			{
				pattern:  "iji",
				phonetic: "D",
			},
			{
				pattern:  "ijI",
				phonetic: "D",
			},
			{
				pattern:  "ije",
				phonetic: "D",
			},
			{
				pattern:  "ijE",
				phonetic: "D",
			},
			{
				pattern:  "Iji",
				phonetic: "D",
			},
			{
				pattern:  "IjI",
				phonetic: "D",
			},
			{
				pattern:  "Ije",
				phonetic: "D",
			},
			{
				pattern:  "IjE",
				phonetic: "D",
			},
			{
				pattern:  "aja",
				phonetic: "D",
			},
			{
				pattern:  "ajA",
				phonetic: "D",
			},
			{
				pattern:  "ajo",
				phonetic: "D",
			},
			{
				pattern:  "ajO",
				phonetic: "D",
			},
			{
				pattern:  "aju",
				phonetic: "D",
			},
			{
				pattern:  "ajU",
				phonetic: "D",
			},
			{
				pattern:  "Aja",
				phonetic: "D",
			},
			{
				pattern:  "AjA",
				phonetic: "D",
			},
			{
				pattern:  "Ajo",
				phonetic: "D",
			},
			{
				pattern:  "AjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "oja",
				phonetic: "D",
			},
			{
				pattern:  "ojA",
				phonetic: "D",
			},
			{
				pattern:  "ojo",
				phonetic: "D",
			},
			{
				pattern:  "ojO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Oja",
				phonetic: "D",
			},
			{
				pattern:  "OjA",
				phonetic: "D",
			},
			{
				pattern:  "Ojo",
				phonetic: "D",
			},
			{
				pattern:  "OjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "eja",
				phonetic: "D",
			},
			{
				pattern:  "ejA",
				phonetic: "D",
			},
			{
				pattern:  "ejo",
				phonetic: "D",
			},
			{
				pattern:  "ejO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Eja",
				phonetic: "D",
			},
			{
				pattern:  "EjA",
				phonetic: "D",
			},
			{
				pattern:  "Ejo",
				phonetic: "D",
			},
			{
				pattern:  "EjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "uja",
				phonetic: "D",
			},
			{
				pattern:  "ujA",
				phonetic: "D",
			},
			{
				pattern:  "ujo",
				phonetic: "D",
			},
			{
				pattern:  "ujO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Uja",
				phonetic: "D",
			},
			{
				pattern:  "UjA",
				phonetic: "D",
			},
			{
				pattern:  "Ujo",
				phonetic: "D",
			},
			{
				pattern:  "UjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "ija",
				phonetic: "D",
			},
			{
				pattern:  "ijA",
				phonetic: "D",
			},
			{
				pattern:  "ijo",
				phonetic: "D",
			},
			{
				pattern:  "ijO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Ija",
				phonetic: "D",
			},
			{
				pattern:  "IjA",
				phonetic: "D",
			},
			{
				pattern:  "Ijo",
				phonetic: "D",
			},
			{
				pattern:  "IjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "j",
				phonetic: "i",
			},
			{
				pattern:      "lYndEr",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lander",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lAndEr",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lAnder",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "landEr",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lender",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lEndEr",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lendEr",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "lEnder",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "lYnder",
			},
			{
				pattern:      "bUrk",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "burk",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "bUrg",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "burg",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "s",
				rightContext: regexp.MustCompile("^[rmnl]"),
				phonetic:     "z",
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
		},
		second: map[uint64][]rule{
			uint64(ashany): []rule{
				{
					pattern:  "b",
					phonetic: "(b|v[1024])",
				},
				{
					pattern:  "J",
					phonetic: "z",
				},
				{
					pattern:      "aiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "AiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "oiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "OiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "uiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "UiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "eiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "EiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "iiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "IiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "aiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "AiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "oiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "OiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "uiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "UiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "eiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "EiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "iiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "IiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(o|om[128]|im[128])",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(a|o|on[128]|in[128])",
				},
				{
					pattern:  "B",
					phonetic: "(a|o)",
				},
				{
					pattern:      "aiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "AiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "oiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "OiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "uiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "UiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "eiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "EiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "iiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "IiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "aiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "AiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "oiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "OiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "uiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "UiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "eiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "EiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "iiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "IiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "F",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(i|im[128]|om[128])",
				},
				{
					pattern:      "F",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(i|in[128]|on[128])",
				},
				{
					pattern:  "F",
					phonetic: "i",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					phonetic:     "(Q[16]|i|D[4])",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk[16])",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts[16])",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(Q[16]|i)",
				},
				{
					pattern:      "lE",
					leftContext:  regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(li|il[4])",
				},
				{
					pattern:     "lE",
					leftContext: regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					phonetic:    "(li|il[4]|lY[16])",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "Ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "Oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "Ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "ei",
					phonetic: "(D|i)",
				},
				{
					pattern:  "Ei",
					phonetic: "(D|i)",
				},
				{
					pattern:      "iA",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ia|io)",
				},
				{
					pattern:  "iA",
					phonetic: "(ia|io|iY[16])",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					phonetic:     "(a|o|Y[16]|D[4])",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("i[^aeiouAEIOU]$"),
					phonetic:    "(i|Y[16]|[4])",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("a[^aeiouAEIOU]$"),
					phonetic:    "(i|Y[16]|[4])",
				},
				{
					pattern:      "e",
					rightContext: regexp.MustCompile("^[fklmnprstv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "e",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:      "e",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "e",
					leftContext: regexp.MustCompile("[DaoiuAOIUQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "e",
					rightContext: regexp.MustCompile("^[aoAOQY]"),
					phonetic:     "i",
				},
				{
					pattern:  "e",
					phonetic: "(i|Y[16])",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoiuAOIUQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoAOQY]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(i|Y[16])",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:      "O",
					rightContext: regexp.MustCompile("^[fklmnprstv]$"),
					phonetic:     "o",
				},
				{
					pattern:      "O",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "o",
				},
				{
					pattern:      "O",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "o",
				},
				{
					pattern:     "O",
					leftContext: regexp.MustCompile("[oeiuQY]$"),
					phonetic:    "o",
				},
				{
					pattern:  "O",
					phonetic: "(o|Y[16])",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "(a|o)",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "(a|o)",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(a|o)",
				},
				{
					pattern:     "A",
					leftContext: regexp.MustCompile("[oeiuQY]$"),
					phonetic:    "(a|o)",
				},
				{
					pattern:  "A",
					phonetic: "(a|o|Y[16])",
				},
				{
					pattern:      "U",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "u",
				},
				{
					pattern:     "U",
					leftContext: regexp.MustCompile("[DoiuQY]$"),
					phonetic:    "u",
				},
				{
					pattern:      "U",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "u",
				},
				{
					pattern:      "Uk",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(uk|Qk[16])",
				},
				{
					pattern:      "Uk",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "uk",
				},
				{
					pattern:      "sUts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(suts|sQts[16])",
				},
				{
					pattern:      "Uts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "uts",
				},
				{
					pattern:  "U",
					phonetic: "(u|Q[16])",
				},
			},
			uint64(ashrussian): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiEIou]$"),
					phonetic:    "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:      "om",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(om|im)",
				},
				{
					pattern:      "on",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(on|in)",
				},
				{
					pattern:      "em",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(im|om)",
				},
				{
					pattern:      "en",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(in|on)",
				},
				{
					pattern:      "Em",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(im|Ym|om)",
				},
				{
					pattern:      "En",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(in|Yn|on)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprsStv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(ashcyrillic): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiEIou]$"),
					phonetic:    "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:      "om",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(om|im)",
				},
				{
					pattern:      "on",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(on|in)",
				},
				{
					pattern:      "em",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(im|om)",
				},
				{
					pattern:      "en",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(in|on)",
				},
				{
					pattern:      "Em",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(im|Ym|om)",
				},
				{
					pattern:      "En",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(in|Yn|on)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprsStv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(ashenglish): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^aEIeiou]e"),
					phonetic:     "(Q|i|D)",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aEIeiou]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:     "lE",
					leftContext: regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					phonetic:    "(il|li|lY)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("D[^aeiEIou]$"),
					phonetic:    "(i|)",
				},
				{
					pattern:     "e",
					leftContext: regexp.MustCompile("D[^aeiEIou]$"),
					phonetic:    "(i|)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprsStv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoiEuQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQY]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
			},
			uint64(ashfrench): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aEIeiou]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprsStv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[aoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(ashgerman): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiAEIOUouQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(Q|i)",
				},
				{
					pattern:  "AU",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "aU",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "Au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "OU",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "oU",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "Ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "Ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "Oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "Ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoAOUiuQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoAOQY]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
				{
					pattern:      "O",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "o",
				},
				{
					pattern:      "O",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "o",
				},
				{
					pattern:      "O",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "o",
				},
				{
					pattern:     "O",
					leftContext: regexp.MustCompile("[aoAOUeiuQY]$"),
					phonetic:    "o",
				},
				{
					pattern:  "O",
					phonetic: "(o|Y)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(a|o)",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "(a|o)",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "(a|o)",
				},
				{
					pattern:     "A",
					leftContext: regexp.MustCompile("[aoeOUiuQY]$"),
					phonetic:    "(a|o)",
				},
				{
					pattern:  "A",
					phonetic: "(a|o|Y)",
				},
				{
					pattern:      "U",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "u",
				},
				{
					pattern:     "U",
					leftContext: regexp.MustCompile("[DaoiuUQY]$"),
					phonetic:    "u",
				},
				{
					pattern:      "U",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "u",
				},
				{
					pattern:      "Uk",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(uk|Qk)",
				},
				{
					pattern:      "Uk",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "uk",
				},
				{
					pattern:      "sUts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(suts|sQts)",
				},
				{
					pattern:      "Uts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "uts",
				},
				{
					pattern:  "U",
					phonetic: "(u|Q)",
				},
			},
			uint64(ashhebrew): []rule{},
			uint64(ashhungarian): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aEIeiou]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprsStv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[aoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(ashpolish): []rule{
				{
					pattern:      "aiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "oiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "uiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "eiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "EiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "iiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "IiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "aiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "oiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "uiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "eiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "EiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "iiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "IiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(o|om|im)",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(o|on|in)",
				},
				{
					pattern:  "B",
					phonetic: "o",
				},
				{
					pattern:      "aiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "oiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "uiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "eiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "EiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "iiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "IiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "aiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "oiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "uiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "eiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "EiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "iiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "IiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "F",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(i|im|om)",
				},
				{
					pattern:      "F",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(i|in|on)",
				},
				{
					pattern:  "F",
					phonetic: "i",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiAEBFIou]$"),
					phonetic:    "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(ashromanian): []rule{
				{
					pattern:      "aiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "oiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "uiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "eiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "EiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "iiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "IiB",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "aiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "oiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "uiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "eiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "EiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "iiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "IiB",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(o|om|im)",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(o|on|in)",
				},
				{
					pattern:  "B",
					phonetic: "o",
				},
				{
					pattern:      "aiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "oiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "uiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "eiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "EiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "iiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "IiF",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(D|Dm)",
				},
				{
					pattern:      "aiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "oiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "uiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "eiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "EiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "iiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "IiF",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(D|Dn)",
				},
				{
					pattern:      "F",
					rightContext: regexp.MustCompile("^[bp]"),
					phonetic:     "(i|im|om)",
				},
				{
					pattern:      "F",
					rightContext: regexp.MustCompile("^[dgkstvz]"),
					phonetic:     "(i|in|on)",
				},
				{
					pattern:  "F",
					phonetic: "i",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiAEBFIou]$"),
					phonetic:    "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprst]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[DaoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(ashspanish): []rule{
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "i",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aEIeiou]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^k]$"),
					phonetic:     "i",
				},
				{
					pattern:      "Ik",
					leftContext:  regexp.MustCompile("[lr]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(ik|Qk)",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts)",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[fklmnprsStv]$"),
					phonetic:     "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^ts$"),
					phonetic:     "i",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("[aoiuQ]$"),
					phonetic:    "i",
				},
				{
					pattern:      "E",
					rightContext: regexp.MustCompile("^[aoQ]"),
					phonetic:     "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
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
				rightContext: regexp.MustCompile("^[gdZz]"),
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
				rightContext: regexp.MustCompile("^[bgdZz]"),
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
				rightContext: regexp.MustCompile("^[bdZz]"),
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
				rightContext: regexp.MustCompile("^[bgZz]"),
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
				pattern:  "jnm",
				phonetic: "jm",
			},
			{
				pattern:     "ji",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "i",
			},
			{
				pattern:     "jI",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "I",
			},
			{
				pattern:      "a",
				rightContext: regexp.MustCompile("^[aAB]"),
				phonetic:     "",
			},
			{
				pattern:     "a",
				leftContext: regexp.MustCompile("[AB]$"),
				phonetic:    "",
			},
			{
				pattern:      "A",
				rightContext: regexp.MustCompile("^A"),
				phonetic:     "",
			},
			{
				pattern:      "B",
				rightContext: regexp.MustCompile("^B"),
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
				pattern:  "H",
				phonetic: "h",
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
				pattern:     "ji",
				leftContext: regexp.MustCompile("[aAoOeEiIuU]$"),
				phonetic:    "j",
			},
			{
				pattern:     "jI",
				leftContext: regexp.MustCompile("[aAoOeEiIuU]$"),
				phonetic:    "j",
			},
			{
				pattern:     "je",
				leftContext: regexp.MustCompile("[aAoOeEiIuU]$"),
				phonetic:    "j",
			},
			{
				pattern:     "jE",
				leftContext: regexp.MustCompile("[aAoOeEiIuU]$"),
				phonetic:    "j",
			},
		},
		second: map[uint64][]rule{
			uint64(ashany): []rule{
				{
					pattern:  "A",
					phonetic: "a",
				},
				{
					pattern:  "B",
					phonetic: "a",
				},
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "F",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
				{
					pattern:  "O",
					phonetic: "o",
				},
				{
					pattern:  "P",
					phonetic: "o",
				},
				{
					pattern:  "U",
					phonetic: "u",
				},
				{
					pattern:  "J",
					phonetic: "l",
				},
			},
			uint64(ashrussian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashcyrillic): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashenglish): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashfrench): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashgerman): []rule{
				{
					pattern:  "A",
					phonetic: "a",
				},
				{
					pattern:  "B",
					phonetic: "a",
				},
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "F",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
				{
					pattern:  "O",
					phonetic: "o",
				},
				{
					pattern:  "P",
					phonetic: "o",
				},
				{
					pattern:  "U",
					phonetic: "u",
				},
				{
					pattern:  "J",
					phonetic: "l",
				},
			},
			uint64(ashhebrew): []rule{},
			uint64(ashhungarian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashpolish): []rule{
				{
					pattern:  "B",
					phonetic: "a",
				},
				{
					pattern:  "F",
					phonetic: "e",
				},
				{
					pattern:  "P",
					phonetic: "o",
				},
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashromanian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(ashspanish): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
		},
	},
}

var ashDiscards = []string{
	"ben",
	"bar",
	"ha",
}
