// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type genLang uint64

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

const genAll = genany +
	genarabic +
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

var genRules = map[genLang][]rule{
	genany: []rule{
		{
			pattern:      "yna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(in[131072]|ina)",
		},
		{
			pattern:      "ina",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(in[131072]|ina)",
		},
		{
			pattern:      "liova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lova|lof[131072]|lef[131072])",
		},
		{
			pattern:      "lova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lova|lof[131072]|lef[131072]|l[8]|el[8])",
		},
		{
			pattern:      "kova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kova|kof[131072]|k[8]|ek[8])",
		},
		{
			pattern:      "ova",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ova|of[131072]|[8])",
		},
		{
			pattern:      "ová",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ova|[8])",
		},
		{
			pattern:      "eva",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(eva|ef[131072])",
		},
		{
			pattern:      "aia",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i[131072])",
		},
		{
			pattern:      "aja",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i[131072])",
		},
		{
			pattern:      "aya",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(aja|i[131072])",
		},
		{
			pattern:      "lowa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lova|lof[16384]|l[16384]|el[16384])",
		},
		{
			pattern:      "kowa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kova|kof[16384]|k[16384]|ek[16384])",
		},
		{
			pattern:      "owa",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ova|of[16384]|)",
		},
		{
			pattern:      "lowna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lovna|levna|l[16384]|el[16384])",
		},
		{
			pattern:      "kowna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(kovna|k[16384]|ek[16384])",
		},
		{
			pattern:      "owna",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ovna|[16384])",
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
			pattern:      "á",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(a|i[8])",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(a|i[16392])",
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
		},
		{
			pattern:      "que",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k[64]|ke|kve)",
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
		},
		{
			pattern:      "m",
			rightContext: regexp.MustCompile("^[bfpv]"),
			phonetic:     "(m|n)",
		},
		{
			pattern:      "m",
			leftContext:  regexp.MustCompile("[aeiouy]$"),
			rightContext: regexp.MustCompile("^[aeiouy]"),
			phonetic:     "m",
		},
		{
			pattern:     "m",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(m|n[32832])",
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
			pattern:  "lio",
			phonetic: "(lo|le[131072])",
		},
		{
			pattern:  "lyo",
			phonetic: "(lo|le[131072])",
		},
		{
			pattern:      "lt",
			leftContext:  regexp.MustCompile("u$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lt|[64])",
		},
		{
			pattern:     "v",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(v|f[128]|b[262144])",
		},
		{
			pattern:      "ex",
			rightContext: regexp.MustCompile("^[aáuiíoóeéêy]"),
			phonetic:     "(ez[32768]|eS[32768]|eks|egz)",
		},
		{
			pattern:      "ex",
			rightContext: regexp.MustCompile("^[cs]"),
			phonetic:     "(e[32768]|ek)",
		},
		{
			pattern:      "x",
			leftContext:  regexp.MustCompile("u$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ks|[64])",
		},
		{
			pattern:  "ck",
			phonetic: "(k|tsk[16392])",
		},
		{
			pattern:  "cz",
			phonetic: "(tS|tsz[8])",
		},
		{
			pattern:     "rh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "r",
		},
		{
			pattern:     "dh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "d",
		},
		{
			pattern:     "bh",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "b",
		},
		{
			pattern:  "ph",
			phonetic: "(ph|f)",
		},
		{
			pattern:  "kh",
			phonetic: "(x[131104]|kh)",
		},
		{
			pattern:  "lh",
			phonetic: "(lh|l[32768])",
		},
		{
			pattern:  "nh",
			phonetic: "(nh|nj[32768])",
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
			leftContext:  regexp.MustCompile("[aeiouy]$"),
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(S|StS[131072]|sk[69632])",
		},
		{
			pattern:     "sch",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(S|StS[131072])",
		},
		{
			pattern:      "sch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(sk[69632]|S|StS[131072])",
		},
		{
			pattern:  "sch",
			phonetic: "(S|StS[131072])",
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
			phonetic:     "(S[131104]|sh)",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:  "zh",
			phonetic: "(Z[131104]|zh|tsh[128])",
		},
		{
			pattern:  "chs",
			phonetic: "(ks[128]|xs|tSs[131104])",
		},
		{
			pattern:      "ch",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(x|tS[393248]|k[69632]|S[32832])",
		},
		{
			pattern:  "ch",
			phonetic: "(x|tS[393248]|S[32832])",
		},
		{
			pattern:     "th",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "t",
		},
		{
			pattern:      "th",
			rightContext: regexp.MustCompile("^[äöüaeiou]"),
			phonetic:     "(t[672]|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:      "gh",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(g[70144]|gh)",
		},
		{
			pattern:      "ouh",
			rightContext: regexp.MustCompile("^[aioe]"),
			phonetic:     "(v[64]|uh)",
		},
		{
			pattern:      "uh",
			rightContext: regexp.MustCompile("^[aioe]"),
			phonetic:     "(v|uh)",
		},
		{
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("[aeiouyäöü]$"),
			phonetic:    "",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(h|x[66048]|H[381024])",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa[16384]|tsa)",
		},
		{
			pattern:      "cią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(tSom|tsom)",
		},
		{
			pattern:  "cią",
			phonetic: "(tSon[16384]|tson)",
		},
		{
			pattern:      "cię",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(tSem[16384]|tsem)",
		},
		{
			pattern:  "cię",
			phonetic: "(tSen[16384]|tsen)",
		},
		{
			pattern:  "cie",
			phonetic: "(tSe[16384]|tse)",
		},
		{
			pattern:  "cio",
			phonetic: "(tSo[16384]|tso)",
		},
		{
			pattern:  "ciu",
			phonetic: "(tSu[16384]|tsu)",
		},
		{
			pattern:      "sci",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(Si[4096]|stsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
		},
		{
			pattern:      "sc",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(S[4096]|sts[16392]|dZ[524288]|tS[81920]|s)",
		},
		{
			pattern:      "ci",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(tsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
		},
		{
			pattern:  "cy",
			phonetic: "(si|tsi[16384])",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(ts[16392]|dZ[524288]|tS[81920]|k[512]|s)",
		},
		{
			pattern:      "sç",
			rightContext: regexp.MustCompile("^[aeiou]"),
			phonetic:     "(s|stS[524288])",
		},
		{
			pattern:  "ssz",
			phonetic: "S",
		},
		{
			pattern:     "sz",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(S|s[2048])",
		},
		{
			pattern:      "sz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(S|s[2048])",
		},
		{
			pattern:  "sz",
			phonetic: "(S|s[2048]|sts[128])",
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp[128]|sp)",
		},
		{
			pattern:  "sp",
			phonetic: "(Sp[128]|sp)",
		},
		{
			pattern:  "sst",
			phonetic: "(St[128]|st)",
		},
		{
			pattern:  "st",
			phonetic: "(St[128]|st)",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern:     "sj",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "S",
		},
		{
			pattern:      "sj",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "S",
		},
		{
			pattern:  "sj",
			phonetic: "(sj|S[16]|sx[262144]|sZ[589824])",
		},
		{
			pattern:  "sia",
			phonetic: "(Sa[16384]|sa[16384]|sja)",
		},
		{
			pattern:      "sią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Som[16384]|som)",
		},
		{
			pattern:  "sią",
			phonetic: "(Son[16384]|son)",
		},
		{
			pattern:      "się",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Sem[16384]|sem)",
		},
		{
			pattern:  "się",
			phonetic: "(Sen[16384]|sen)",
		},
		{
			pattern:  "sie",
			phonetic: "(se|sje|Se[16384]|zi[128])",
		},
		{
			pattern:  "sio",
			phonetic: "(So[16384]|so)",
		},
		{
			pattern:  "siu",
			phonetic: "(Su[16384]|sju)",
		},
		{
			pattern:     "si",
			leftContext: regexp.MustCompile("[äöëaáuiíoóeéêy]$"),
			phonetic:    "(Si[16384]|si|zi[37056])",
		},
		{
			pattern:  "si",
			phonetic: "(Si[16384]|si|zi[128])",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aáuiíoóeéêy]$"),
			rightContext: regexp.MustCompile("^[aáuíoóeéêy]"),
			phonetic:     "(s|z[37056])",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[aeouäöë]"),
			phonetic:     "(s|z[128])",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aeiouy]$"),
			rightContext: regexp.MustCompile("^[dglmnrv]"),
			phonetic:     "(s|z|Z[32768]|[64])",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[dglmnrv]"),
			phonetic:     "(s|z|Z[32768])",
		},
		{
			pattern:      "gue",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k[64]|gve)",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[ei]"),
			phonetic:     "(g[64]|gv[294912])",
		},
		{
			pattern:      "gu",
			rightContext: regexp.MustCompile("^[ao]"),
			phonetic:     "gv",
		},
		{
			pattern:  "guy",
			phonetic: "gi",
		},
		{
			pattern:  "gli",
			phonetic: "(glI|l[4096])",
		},
		{
			pattern:  "gni",
			phonetic: "(gnI|ni[4160])",
		},
		{
			pattern:      "gn",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "(n[4160]|nj[4160]|gn)",
		},
		{
			pattern:  "ggie",
			phonetic: "(je[512]|dZe)",
		},
		{
			pattern:      "ggi",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "(j[512]|dZ)",
		},
		{
			pattern:      "ggi",
			leftContext:  regexp.MustCompile("[yaeiou]$"),
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "(gI|dZ[4096]|j[512])",
		},
		{
			pattern:     "gge",
			leftContext: regexp.MustCompile("[yaeiou]$"),
			phonetic:    "(gE|xe[262144]|gZe[32832]|dZe[331808]|je[512])",
		},
		{
			pattern:     "ggi",
			leftContext: regexp.MustCompile("[yaeiou]$"),
			phonetic:    "(gI|xi[262144]|gZi[32832]|dZi[331808]|i[512])",
		},
		{
			pattern:      "ggi",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "(gI|dZ[4096]|j[512])",
		},
		{
			pattern:      "gie",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ge|gi[128]|ji[64]|dZe[4096])",
		},
		{
			pattern:  "gie",
			phonetic: "(ge|gi[128]|dZe[4096]|je[512])",
		},
		{
			pattern:      "gi",
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "(i[512]|dZ)",
		},
		{
			pattern:     "ge",
			leftContext: regexp.MustCompile("[yaeiou]$"),
			phonetic:    "(gE|xe[262144]|Ze[32832]|dZe[331808])",
		},
		{
			pattern:     "gi",
			leftContext: regexp.MustCompile("[yaeiou]$"),
			phonetic:    "(gI|xi[262144]|Zi[32832]|dZi[331808])",
		},
		{
			pattern:  "ge",
			phonetic: "(gE|xe[262144]|hE[131072]|je[512]|Ze[32832]|dZe[331808])",
		},
		{
			pattern:  "gi",
			phonetic: "(gI|xi[262144]|hI[131072]|i[512]|Zi[32832]|dZi[331808])",
		},
		{
			pattern:      "gy",
			rightContext: regexp.MustCompile("^[aeouáéóúüöőű]"),
			phonetic:     "(gi|dj[2048])",
		},
		{
			pattern:  "gy",
			phonetic: "(gi|d[2048])",
		},
		{
			pattern:      "g",
			leftContext:  regexp.MustCompile("[yaeiou]$"),
			rightContext: regexp.MustCompile("^[aouyei]"),
			phonetic:     "g",
		},
		{
			pattern:      "g",
			rightContext: regexp.MustCompile("^[aouei]"),
			phonetic:     "(g|h[131072])",
		},
		{
			pattern:  "ij",
			phonetic: "(i|ej[16]|ix[262144]|iZ[622656])",
		},
		{
			pattern:      "j",
			rightContext: regexp.MustCompile("^[aoeiuy]"),
			phonetic:     "(j|dZ[32]|x[262144]|Z[622656])",
		},
		{
			pattern:     "rz",
			leftContext: regexp.MustCompile("t$"),
			phonetic:    "(S[16384]|r)",
		},
		{
			pattern:  "rz",
			phonetic: "(rz|rts[128]|Z[16384]|r[16384]|rZ[16384])",
		},
		{
			pattern:      "tz",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ts|tS[160])",
		},
		{
			pattern:     "tz",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(ts[131232]|tS[160])",
		},
		{
			pattern:  "tz",
			phonetic: "(ts[131232]|tz)",
		},
		{
			pattern:      "zia",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(Za[16384]|za[16384]|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za[16384]|zja)",
		},
		{
			pattern:      "zią",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Zom[16384]|zom)",
		},
		{
			pattern:  "zią",
			phonetic: "(Zon[16384]|zon)",
		},
		{
			pattern:      "zię",
			rightContext: regexp.MustCompile("^[bp]"),
			phonetic:     "(Zem[16384]|zem)",
		},
		{
			pattern:  "zię",
			phonetic: "(Zen[16384]|zen)",
		},
		{
			pattern:      "zie",
			rightContext: regexp.MustCompile("^[bcdgkpstwzż]"),
			phonetic:     "(Ze[16384]|ze[16384]|ze|tsi[128])",
		},
		{
			pattern:  "zie",
			phonetic: "(ze|Ze[16384]|tsi[128])",
		},
		{
			pattern:  "zio",
			phonetic: "(Zo[16384]|zo)",
		},
		{
			pattern:  "ziu",
			phonetic: "(Zu[16384]|zju)",
		},
		{
			pattern:  "zi",
			phonetic: "(Zi[16384]|zi|tsi[128]|dzi[4096]|tsi[4096]|si[262144])",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(s|ts[128]|ts[4096]|S[32768])",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^[bdgv]"),
			phonetic:     "(z|dz[4096]|Z[32768])",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^[ptckf]"),
			phonetic:     "(s|ts[4096]|S[32768])",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "oue",
			phonetic: "(oue|ve[64])",
		},
		{
			pattern:  "eau",
			phonetic: "o",
		},
		{
			pattern:  "ae",
			phonetic: "(Y[128]|aje[131072]|ae)",
		},
		{
			pattern:  "ai",
			phonetic: "aj",
		},
		{
			pattern:  "au",
			phonetic: "(au|o[64])",
		},
		{
			pattern:  "ay",
			phonetic: "aj",
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
			pattern:  "ea",
			phonetic: "(ea|ja[65536])",
		},
		{
			pattern:  "ee",
			phonetic: "(i[32]|aje[131072]|e)",
		},
		{
			pattern:  "ei",
			phonetic: "(aj|ej)",
		},
		{
			pattern:  "eu",
			phonetic: "(eu|Yj[128]|ej[128]|oj[128]|Y[16])",
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej)",
		},
		{
			pattern:  "ia",
			phonetic: "ja",
		},
		{
			pattern:  "ie",
			phonetic: "(i[128]|e[16384]|ije[131072]|Q[16]|je)",
		},
		{
			pattern:      "ii",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:  "io",
			phonetic: "(jo|e[131072])",
		},
		{
			pattern:  "iu",
			phonetic: "ju",
		},
		{
			pattern:      "iy",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:  "oe",
			phonetic: "(Y[128]|oje[131072]|u[16]|oe)",
		},
		{
			pattern:  "oi",
			phonetic: "oj",
		},
		{
			pattern:  "oo",
			phonetic: "(u[32]|o)",
		},
		{
			pattern:  "ou",
			phonetic: "(ou|u[576]|au[16])",
		},
		{
			pattern:  "où",
			phonetic: "u",
		},
		{
			pattern:  "oy",
			phonetic: "oj",
		},
		{
			pattern:  "õe",
			phonetic: "(oj|on)",
		},
		{
			pattern:  "ua",
			phonetic: "va",
		},
		{
			pattern:  "ue",
			phonetic: "(Q[128]|uje[131072]|ve)",
		},
		{
			pattern:  "ui",
			phonetic: "(uj|vi|Y[16])",
		},
		{
			pattern:  "uu",
			phonetic: "(u|Q[16])",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
		},
		{
			pattern:  "uy",
			phonetic: "uj",
		},
		{
			pattern:  "ya",
			phonetic: "ja",
		},
		{
			pattern:  "ye",
			phonetic: "(je|ije[131072])",
		},
		{
			pattern:     "yi",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "i",
		},
		{
			pattern:      "yi",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e[131072])",
		},
		{
			pattern:  "yu",
			phonetic: "ju",
		},
		{
			pattern:      "yy",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[áóéê]$"),
			phonetic:    "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[áóéê]$"),
			phonetic:    "j",
		},
		{
			pattern:     "e",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(e|je[131072])",
		},
		{
			pattern:      "e",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(e|EE[96])",
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
			pattern:  "à",
			phonetic: "a",
		},
		{
			pattern:  "â",
			phonetic: "a",
		},
		{
			pattern:  "ã",
			phonetic: "(a|an)",
		},
		{
			pattern:  "ă",
			phonetic: "(e[65536]|a)",
		},
		{
			pattern:  "ā",
			phonetic: "a",
		},
		{
			pattern:  "č",
			phonetic: "tS",
		},
		{
			pattern:  "ć",
			phonetic: "(tS[16384]|ts)",
		},
		{
			pattern:  "ç",
			phonetic: "(s|tS[524288])",
		},
		{
			pattern:  "ď",
			phonetic: "(d|dj[8])",
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
			pattern:  "ě",
			phonetic: "(e|je[8])",
		},
		{
			pattern:  "ē",
			phonetic: "e",
		},
		{
			pattern:  "ģ",
			phonetic: "(d|dj)",
		},
		{
			pattern:  "ğ",
			phonetic: "",
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
			pattern:  "ī",
			phonetic: "i",
		},
		{
			pattern:  "ı",
			phonetic: "(i|e[524288]|[524288])",
		},
		{
			pattern:  "ķ",
			phonetic: "(k|t[8192]|tj[8192])",
		},
		{
			pattern:  "ļ",
			phonetic: "l",
		},
		{
			pattern:  "ł",
			phonetic: "l",
		},
		{
			pattern:  "ń",
			phonetic: "(n|nj[16384])",
		},
		{
			pattern:  "ñ",
			phonetic: "(n|nj[262144])",
		},
		{
			pattern:  "ņ",
			phonetic: "(n|nj[8192])",
		},
		{
			pattern:  "ó",
			phonetic: "(u[16384]|o)",
		},
		{
			pattern:  "ô",
			phonetic: "o",
		},
		{
			pattern:  "õ",
			phonetic: "(o|on[32768]|Y[2048])",
		},
		{
			pattern:  "ò",
			phonetic: "o",
		},
		{
			pattern:  "ö",
			phonetic: "Y",
		},
		{
			pattern:  "ř",
			phonetic: "(r|rZ[8])",
		},
		{
			pattern:  "ś",
			phonetic: "(S[16384]|s)",
		},
		{
			pattern:  "ş",
			phonetic: "S",
		},
		{
			pattern:  "š",
			phonetic: "S",
		},
		{
			pattern:  "ţ",
			phonetic: "ts",
		},
		{
			pattern:  "ť",
			phonetic: "(t|tj[8])",
		},
		{
			pattern:  "ű",
			phonetic: "Q",
		},
		{
			pattern:  "ü",
			phonetic: "(Q|u[294912])",
		},
		{
			pattern:  "ū",
			phonetic: "u",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "ů",
			phonetic: "u",
		},
		{
			pattern:  "ù",
			phonetic: "u",
		},
		{
			pattern:  "ý",
			phonetic: "i",
		},
		{
			pattern:  "ż",
			phonetic: "Z",
		},
		{
			pattern:  "ź",
			phonetic: "(Z[16384]|z)",
		},
		{
			pattern:  "ž",
			phonetic: "Z",
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
			pattern:      "o",
			rightContext: regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			phonetic:     "(O|P[16384])",
		},
		{
			pattern:  "a",
			phonetic: "A",
		},
		{
			pattern:  "b",
			phonetic: "B",
		},
		{
			pattern:  "c",
			phonetic: "(k|ts[16392]|dZ[524288])",
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
			phonetic: "(h|x[65536]|H[299072])",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "(j|x[262144]|Z[622656])",
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
			phonetic: "(s|S[32768])",
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
			phonetic: "V",
		},
		{
			pattern:  "w",
			phonetic: "(v|w[48])",
		},
		{
			pattern:  "x",
			phonetic: "(ks|gz|S[294912])",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "(z|ts[128]|dz[4096]|ts[4096]|s[262144])",
		},
	},
	genarabic: []rule{
		{
			pattern:  "ا",
			phonetic: "a",
		},
		{
			pattern:      "ب",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "b",
		},
		{
			pattern:  "ب",
			phonetic: "b1",
		},
		{
			pattern:      "ت",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "t",
		},
		{
			pattern:  "ت",
			phonetic: "t1",
		},
		{
			pattern:      "ث",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "t",
		},
		{
			pattern:  "ث",
			phonetic: "t1",
		},
		{
			pattern:      "ج",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(dZ|Z)",
		},
		{
			pattern:  "ج",
			phonetic: "(dZ1|Z1)",
		},
		{
			pattern:     "ح",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "1",
		},
		{
			pattern:      "ح",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "1",
		},
		{
			pattern:  "ح",
			phonetic: "(h1|1)",
		},
		{
			pattern:      "خ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "x",
		},
		{
			pattern:  "خ",
			phonetic: "x1",
		},
		{
			pattern:      "د",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "d",
		},
		{
			pattern:  "د",
			phonetic: "d1",
		},
		{
			pattern:      "ذ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "d",
		},
		{
			pattern:  "ذ",
			phonetic: "d1",
		},
		{
			pattern:      "ر",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "r",
		},
		{
			pattern:  "ر",
			phonetic: "r1",
		},
		{
			pattern:      "ز",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "z",
		},
		{
			pattern:  "ز",
			phonetic: "z1",
		},
		{
			pattern:      "س",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "s",
		},
		{
			pattern:  "س",
			phonetic: "s1",
		},
		{
			pattern:      "ش",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "S",
		},
		{
			pattern:  "ش",
			phonetic: "S1",
		},
		{
			pattern:      "ص",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "s",
		},
		{
			pattern:  "ص",
			phonetic: "s1",
		},
		{
			pattern:      "ض",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "d",
		},
		{
			pattern:  "ض",
			phonetic: "d1",
		},
		{
			pattern:      "ط",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "t",
		},
		{
			pattern:  "ط",
			phonetic: "t1",
		},
		{
			pattern:      "ظ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "z",
		},
		{
			pattern:  "ظ",
			phonetic: "z1",
		},
		{
			pattern:     "ع",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "1",
		},
		{
			pattern:      "ع",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "1",
		},
		{
			pattern:  "ع",
			phonetic: "(h1|1)",
		},
		{
			pattern:      "غ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "g",
		},
		{
			pattern:  "غ",
			phonetic: "g1",
		},
		{
			pattern:      "ف",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "f",
		},
		{
			pattern:  "ف",
			phonetic: "f1",
		},
		{
			pattern:      "ق",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "k",
		},
		{
			pattern:  "ق",
			phonetic: "k1",
		},
		{
			pattern:      "ك",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "k",
		},
		{
			pattern:  "ك",
			phonetic: "k1",
		},
		{
			pattern:      "ل",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "l",
		},
		{
			pattern:  "ل",
			phonetic: "l1",
		},
		{
			pattern:      "م",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "m",
		},
		{
			pattern:  "م",
			phonetic: "m1",
		},
		{
			pattern:      "ن",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "n",
		},
		{
			pattern:  "ن",
			phonetic: "n1",
		},
		{
			pattern:     "ه",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "1",
		},
		{
			pattern:      "ه",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "1",
		},
		{
			pattern:  "ه",
			phonetic: "(h1|1)",
		},
		{
			pattern:      "و",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(u|v)",
		},
		{
			pattern:  "و",
			phonetic: "(u|v1)",
		},
		{
			pattern:      "ي\u200e",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(i|j)",
		},
		{
			pattern:  "ي\u200e",
			phonetic: "(i|j1)",
		},
	},
	gencyrillic: []rule{
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
			pattern:      "с",
			rightContext: regexp.MustCompile("^с"),
			phonetic:     "",
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
			phonetic:     "(hejmer|hajmer)",
		},
		{
			pattern:      "гейм",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(hejm|hajm)",
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
			pattern:     "ей",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "(jej|ej)",
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
			phonetic: "ej",
		},
		{
			pattern:  "ей",
			phonetic: "ej",
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
			pattern:  "ё",
			phonetic: "(e|jo)",
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
	genczech: []rule{
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:  "qu",
			phonetic: "(k|kv)",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "ei",
			phonetic: "(ej|aj)",
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
			pattern:  "č",
			phonetic: "tS",
		},
		{
			pattern:  "š",
			phonetic: "S",
		},
		{
			pattern:  "ž",
			phonetic: "Z",
		},
		{
			pattern:  "ň",
			phonetic: "n",
		},
		{
			pattern:  "ť",
			phonetic: "(t|tj)",
		},
		{
			pattern:  "ď",
			phonetic: "(d|dj)",
		},
		{
			pattern:  "ř",
			phonetic: "(r|rZ)",
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
			pattern:  "ý",
			phonetic: "i",
		},
		{
			pattern:  "ě",
			phonetic: "(e|je)",
		},
		{
			pattern:  "ů",
			phonetic: "u",
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
			phonetic: "(h|g)",
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
			phonetic: "(k|kv)",
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
			phonetic: "z",
		},
	},
	gendutch: []rule{
		{
			pattern:  "ssj",
			phonetic: "S",
		},
		{
			pattern:  "sj",
			phonetic: "S",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:      "c",
			rightContext: regexp.MustCompile("^[eiy]"),
			phonetic:     "ts",
		},
		{
			pattern:  "ck",
			phonetic: "k",
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
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern:     "h",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "ou",
			phonetic: "au",
		},
		{
			pattern:  "ie",
			phonetic: "(Q|i)",
		},
		{
			pattern:  "uu",
			phonetic: "(Q|u)",
		},
		{
			pattern:  "ee",
			phonetic: "e",
		},
		{
			pattern:  "eu",
			phonetic: "(Y|Yj)",
		},
		{
			pattern:  "aa",
			phonetic: "a",
		},
		{
			pattern:  "oo",
			phonetic: "o",
		},
		{
			pattern:  "oe",
			phonetic: "u",
		},
		{
			pattern:  "ij",
			phonetic: "ej",
		},
		{
			pattern:  "ui",
			phonetic: "(Y|uj)",
		},
		{
			pattern:  "ei",
			phonetic: "(ej|aj)",
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
			pattern:     "i",
			leftContext: regexp.MustCompile("[aou]$"),
			phonetic:    "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
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
			phonetic: "(g|x)",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "(i|Q)",
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
			phonetic: "(u|Q)",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
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
	genenglish: []rule{
		{
			pattern:  "�",
			phonetic: "",
		},
		{
			pattern:  "'",
			phonetic: "",
		},
		{
			pattern:     "mc",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "mak",
		},
		{
			pattern:  "tz",
			phonetic: "ts",
		},
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
			pattern:     "x",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "z",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "j",
		},
		{
			pattern:      "y",
			leftContext:  regexp.MustCompile("^$"),
			rightContext: regexp.MustCompile("^[aeiouy]"),
			phonetic:     "j",
		},
		{
			pattern:     "yi",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "i",
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
			phonetic: "(aj|ej|e)",
		},
		{
			pattern:  "ay",
			phonetic: "(aj|ej)",
		},
		{
			pattern:      "a",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "ej",
		},
		{
			pattern:  "ei",
			phonetic: "(ej|aj|i)",
		},
		{
			pattern:  "ey",
			phonetic: "(ej|aj|i)",
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
			pattern:  "ie",
			phonetic: "i",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[^aeiou]e"),
			phonetic:     "aj",
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
			pattern:  "a",
			phonetic: "(e|o|a)",
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
			phonetic: "dZ",
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
			phonetic: "(o|a)",
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
			phonetic: "(u|a)",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
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
	genfrench: []rule{
		{
			pattern:      "lt",
			leftContext:  regexp.MustCompile("u$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(lt|)",
		},
		{
			pattern:      "c",
			leftContext:  regexp.MustCompile("n$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k|)",
		},
		{
			pattern:      "d",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(t|)",
		},
		{
			pattern:      "g",
			leftContext:  regexp.MustCompile("n$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(k|)",
		},
		{
			pattern:      "p",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(p|)",
		},
		{
			pattern:      "r",
			leftContext:  regexp.MustCompile("e$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(r|)",
		},
		{
			pattern:      "t",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(t|)",
		},
		{
			pattern:      "z",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(s|)",
		},
		{
			pattern:      "ds",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ds|)",
		},
		{
			pattern:      "ps",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ps|)",
		},
		{
			pattern:      "rs",
			leftContext:  regexp.MustCompile("e$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(rs|)",
		},
		{
			pattern:      "ts",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ts|)",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(s|)",
		},
		{
			pattern:      "x",
			leftContext:  regexp.MustCompile("u$"),
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "(ks|)",
		},
		{
			pattern:      "s",
			leftContext:  regexp.MustCompile("[aeéèêiou]$"),
			rightContext: regexp.MustCompile("^[^aeéèêiou]"),
			phonetic:     "(s|)",
		},
		{
			pattern:      "t",
			leftContext:  regexp.MustCompile("[aeéèêiou]$"),
			rightContext: regexp.MustCompile("^[^aeéèêiou]"),
			phonetic:     "(t|)",
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
			pattern:      "aill",
			rightContext: regexp.MustCompile("^e"),
			phonetic:     "aj",
		},
		{
			pattern:      "ll",
			rightContext: regexp.MustCompile("^e"),
			phonetic:     "(l|j)",
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
			pattern:      "m",
			leftContext:  regexp.MustCompile("[aeiouy]$"),
			rightContext: regexp.MustCompile("^[aeiouy]"),
			phonetic:     "m",
		},
		{
			pattern:     "m",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(m|n)",
		},
		{
			pattern:      "ou",
			rightContext: regexp.MustCompile("^[aeio]"),
			phonetic:     "v",
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
			pattern:  "au",
			phonetic: "(o|au)",
		},
		{
			pattern:  "ai",
			phonetic: "(e|aj)",
		},
		{
			pattern:  "ay",
			phonetic: "(e|aj)",
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
			phonetic: "(oj|va)",
		},
		{
			pattern:  "ei",
			phonetic: "(aj|ej|e)",
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej|e)",
		},
		{
			pattern:  "eu",
			phonetic: "(ej|Y)",
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
			phonetic: "(u|Q)",
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
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	gengerman: []rule{
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
			pattern:  "ssch",
			phonetic: "S",
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
		},
		{
			pattern:  "sch",
			phonetic: "S",
		},
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
			phonetic: "(aj|ej)",
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej)",
		},
		{
			pattern:  "eu",
			phonetic: "(Yj|ej|aj|oj)",
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
	gengreek: []rule{
		{
			pattern:      "αυ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "af",
		},
		{
			pattern:      "αυ",
			rightContext: regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			phonetic:     "af",
		},
		{
			pattern:  "αυ",
			phonetic: "av",
		},
		{
			pattern:      "ευ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ef",
		},
		{
			pattern:      "ευ",
			rightContext: regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			phonetic:     "ef",
		},
		{
			pattern:  "ευ",
			phonetic: "ev",
		},
		{
			pattern:      "ηυ",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "if",
		},
		{
			pattern:      "ηυ",
			rightContext: regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			phonetic:     "if",
		},
		{
			pattern:  "ηυ",
			phonetic: "iv",
		},
		{
			pattern:  "ου",
			phonetic: "u",
		},
		{
			pattern:  "αι",
			phonetic: "aj",
		},
		{
			pattern:  "ει",
			phonetic: "ej",
		},
		{
			pattern:  "οι",
			phonetic: "oj",
		},
		{
			pattern:  "ωι",
			phonetic: "oj",
		},
		{
			pattern:  "ηι",
			phonetic: "ej",
		},
		{
			pattern:  "υι",
			phonetic: "i",
		},
		{
			pattern:      "γγ",
			leftContext:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			rightContext: regexp.MustCompile("^(ε|ι|η)"),
			phonetic:     "(nj|j)",
		},
		{
			pattern:      "γγ",
			rightContext: regexp.MustCompile("^(ε|ι|η)"),
			phonetic:     "j",
		},
		{
			pattern:     "γγ",
			leftContext: regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			phonetic:    "(ng|g)",
		},
		{
			pattern:  "γγ",
			phonetic: "g",
		},
		{
			pattern:     "γκ",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "g",
		},
		{
			pattern:      "γκ",
			leftContext:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			rightContext: regexp.MustCompile("^(ε|ι|η)"),
			phonetic:     "(nj|j)",
		},
		{
			pattern:      "γκ",
			rightContext: regexp.MustCompile("^(ε|ι|η)"),
			phonetic:     "j",
		},
		{
			pattern:     "γκ",
			leftContext: regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			phonetic:    "(ng|g)",
		},
		{
			pattern:  "γκ",
			phonetic: "g",
		},
		{
			pattern:      "γι",
			rightContext: regexp.MustCompile("^(α|ο|ω|υ)"),
			phonetic:     "j",
		},
		{
			pattern:  "γι",
			phonetic: "(gi|i)",
		},
		{
			pattern:      "γε",
			rightContext: regexp.MustCompile("^(α|ο|ω|υ)"),
			phonetic:     "j",
		},
		{
			pattern:  "γε",
			phonetic: "(ge|je)",
		},
		{
			pattern:  "κζ",
			phonetic: "gz",
		},
		{
			pattern:  "τζ",
			phonetic: "dz",
		},
		{
			pattern:      "σ",
			rightContext: regexp.MustCompile("^(β|γ|δ|μ|ν|ρ)"),
			phonetic:     "z",
		},
		{
			pattern:  "μβ",
			phonetic: "(mb|b)",
		},
		{
			pattern:     "μπ",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "b",
		},
		{
			pattern:     "μπ",
			leftContext: regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			phonetic:    "mb",
		},
		{
			pattern:  "μπ",
			phonetic: "b",
		},
		{
			pattern:     "ντ",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "d",
		},
		{
			pattern:     "ντ",
			leftContext: regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			phonetic:    "(nd|nt)",
		},
		{
			pattern:  "ντ",
			phonetic: "(nt|d)",
		},
		{
			pattern:  "ά",
			phonetic: "a",
		},
		{
			pattern:  "έ",
			phonetic: "e",
		},
		{
			pattern:  "ή",
			phonetic: "(i|e)",
		},
		{
			pattern:  "ί",
			phonetic: "i",
		},
		{
			pattern:  "ό",
			phonetic: "o",
		},
		{
			pattern:  "ύ",
			phonetic: "(Q|i|u)",
		},
		{
			pattern:  "ώ",
			phonetic: "o",
		},
		{
			pattern:  "ΰ",
			phonetic: "(Q|i|u)",
		},
		{
			pattern:  "ϋ",
			phonetic: "(Q|i|u)",
		},
		{
			pattern:  "ϊ",
			phonetic: "j",
		},
		{
			pattern:  "α",
			phonetic: "a",
		},
		{
			pattern:  "β",
			phonetic: "(v|b)",
		},
		{
			pattern:  "γ",
			phonetic: "g",
		},
		{
			pattern:  "δ",
			phonetic: "d",
		},
		{
			pattern:  "ε",
			phonetic: "e",
		},
		{
			pattern:  "ζ",
			phonetic: "z",
		},
		{
			pattern:  "η",
			phonetic: "(i|e)",
		},
		{
			pattern:  "ι",
			phonetic: "i",
		},
		{
			pattern:  "κ",
			phonetic: "k",
		},
		{
			pattern:  "λ",
			phonetic: "l",
		},
		{
			pattern:  "μ",
			phonetic: "m",
		},
		{
			pattern:  "ν",
			phonetic: "n",
		},
		{
			pattern:  "ξ",
			phonetic: "ks",
		},
		{
			pattern:  "ο",
			phonetic: "o",
		},
		{
			pattern:  "π",
			phonetic: "p",
		},
		{
			pattern:  "ρ",
			phonetic: "r",
		},
		{
			pattern:  "σ",
			phonetic: "s",
		},
		{
			pattern:  "ς",
			phonetic: "s",
		},
		{
			pattern:  "τ",
			phonetic: "t",
		},
		{
			pattern:  "υ",
			phonetic: "(Q|i|u)",
		},
		{
			pattern:  "φ",
			phonetic: "f",
		},
		{
			pattern:  "θ",
			phonetic: "t",
		},
		{
			pattern:  "χ",
			phonetic: "x",
		},
		{
			pattern:  "ψ",
			phonetic: "ps",
		},
		{
			pattern:  "ω",
			phonetic: "o",
		},
	},
	gengreeklatin: []rule{
		{
			pattern:      "au",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "af",
		},
		{
			pattern:      "au",
			rightContext: regexp.MustCompile("^[kpstfh]"),
			phonetic:     "af",
		},
		{
			pattern:  "au",
			phonetic: "av",
		},
		{
			pattern:      "eu",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "ef",
		},
		{
			pattern:      "eu",
			rightContext: regexp.MustCompile("^[kpstfh]"),
			phonetic:     "ef",
		},
		{
			pattern:  "eu",
			phonetic: "ev",
		},
		{
			pattern:  "ou",
			phonetic: "u",
		},
		{
			pattern:     "gge",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(nje|je)",
		},
		{
			pattern:      "ggi",
			leftContext:  regexp.MustCompile("[aeiouy]$"),
			rightContext: regexp.MustCompile("^[aou]"),
			phonetic:     "(nj|j)",
		},
		{
			pattern:     "ggi",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(ni|i)",
		},
		{
			pattern:  "gge",
			phonetic: "je",
		},
		{
			pattern:  "ggi",
			phonetic: "i",
		},
		{
			pattern:     "gg",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(ng|g)",
		},
		{
			pattern:  "gg",
			phonetic: "g",
		},
		{
			pattern:     "gk",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "g",
		},
		{
			pattern:     "gke",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(nje|je)",
		},
		{
			pattern:     "gki",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(ni|i)",
		},
		{
			pattern:  "gke",
			phonetic: "je",
		},
		{
			pattern:  "gki",
			phonetic: "i",
		},
		{
			pattern:     "gk",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(ng|g)",
		},
		{
			pattern:  "gk",
			phonetic: "g",
		},
		{
			pattern:      "nghi",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "Nj",
		},
		{
			pattern:  "nghi",
			phonetic: "(Ngi|Ni)",
		},
		{
			pattern:      "nghe",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "Nj",
		},
		{
			pattern:  "nghe",
			phonetic: "(Nje|Nge)",
		},
		{
			pattern:      "ghi",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "j",
		},
		{
			pattern:  "ghi",
			phonetic: "(gi|i)",
		},
		{
			pattern:      "ghe",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "j",
		},
		{
			pattern:  "ghe",
			phonetic: "(je|ge)",
		},
		{
			pattern:  "ngh",
			phonetic: "Ng",
		},
		{
			pattern:  "gh",
			phonetic: "g",
		},
		{
			pattern:      "ngi",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "Nj",
		},
		{
			pattern:  "ngi",
			phonetic: "(Ngi|Ni)",
		},
		{
			pattern:      "nge",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "Nj",
		},
		{
			pattern:  "nge",
			phonetic: "(Nje|Nge)",
		},
		{
			pattern:      "gi",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "j",
		},
		{
			pattern:  "gi",
			phonetic: "(gi|i)",
		},
		{
			pattern:      "ge",
			rightContext: regexp.MustCompile("^[aouy]"),
			phonetic:     "j",
		},
		{
			pattern:  "ge",
			phonetic: "(je|ge)",
		},
		{
			pattern:  "ng",
			phonetic: "Ng",
		},
		{
			pattern:      "i",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:     "i",
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "y",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "yi",
			rightContext: regexp.MustCompile("^[aeou]"),
			phonetic:     "j",
		},
		{
			pattern:  "yi",
			phonetic: "i",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "dh",
			phonetic: "d",
		},
		{
			pattern:  "dj",
			phonetic: "dZ",
		},
		{
			pattern:  "ph",
			phonetic: "f",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:  "kz",
			phonetic: "gz",
		},
		{
			pattern:  "tz",
			phonetic: "dz",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^[bgdmnr]"),
			phonetic:     "z",
		},
		{
			pattern:  "mb",
			phonetic: "(mb|b)",
		},
		{
			pattern:     "mp",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "b",
		},
		{
			pattern:     "mp",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "mp",
		},
		{
			pattern:  "mp",
			phonetic: "b",
		},
		{
			pattern:     "nt",
			leftContext: regexp.MustCompile("^$"),
			phonetic:    "d",
		},
		{
			pattern:     "nt",
			leftContext: regexp.MustCompile("[aeiouy]$"),
			phonetic:    "(nd|nt)",
		},
		{
			pattern:  "nt",
			phonetic: "(nt|d)",
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
			pattern:  "óu",
			phonetic: "u",
		},
		{
			pattern:  "ú",
			phonetic: "u",
		},
		{
			pattern:  "ý",
			phonetic: "(i|Q|u)",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "(b|v)",
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
			phonetic: "x",
		},
		{
			pattern:  "i",
			phonetic: "i",
		},
		{
			pattern:  "j",
			phonetic: "(j|Z)",
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
			pattern:  "ο",
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
			phonetic: "(i|Q|u)",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	genhebrew: []rule{
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
	genhungarian: []rule{
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
			phonetic: "(aj|ej)",
		},
		{
			pattern:  "ey",
			phonetic: "(aj|ej)",
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
			phonetic: "(ej|e)",
		},
		{
			pattern:  "ely",
			phonetic: "(ej|eli)",
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
			pattern:  "ú",
			phonetic: "u",
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
	genitalian: []rule{
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
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
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
	genlatvian: []rule{
		{
			pattern:  "č",
			phonetic: "tS",
		},
		{
			pattern:  "ģ",
			phonetic: "(d|dj)",
		},
		{
			pattern:  "ķ",
			phonetic: "(t|tj)",
		},
		{
			pattern:  "ļ",
			phonetic: "l",
		},
		{
			pattern:  "š",
			phonetic: "S",
		},
		{
			pattern:  "ņ",
			phonetic: "(n|nj)",
		},
		{
			pattern:  "ž",
			phonetic: "Z",
		},
		{
			pattern:  "ā",
			phonetic: "a",
		},
		{
			pattern:  "ē",
			phonetic: "e",
		},
		{
			pattern:  "ī",
			phonetic: "i",
		},
		{
			pattern:  "ū",
			phonetic: "u",
		},
		{
			pattern:  "ai",
			phonetic: "aj",
		},
		{
			pattern:  "ei",
			phonetic: "ej",
		},
		{
			pattern:  "io",
			phonetic: "jo",
		},
		{
			pattern:  "iu",
			phonetic: "ju",
		},
		{
			pattern:  "ie",
			phonetic: "je",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "ui",
			phonetic: "uj",
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
	genpolish: []rule{
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
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
		},
		{
			pattern:     "y",
			leftContext: regexp.MustCompile("[aeou]$"),
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
	genportuguese: []rule{
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
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
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
	genromanian: []rule{
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
			pattern:  "ţ",
			phonetic: "ts",
		},
		{
			pattern:  "ş",
			phonetic: "S",
		},
		{
			pattern:  "qu",
			phonetic: "k",
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
			phonetic: "(x|h)",
		},
		{
			pattern:  "i",
			phonetic: "I",
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
			phonetic: "z",
		},
	},
	genrussian: []rule{
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
			pattern:  "qu",
			phonetic: "(kv|k)",
		},
		{
			pattern:      "s",
			rightContext: regexp.MustCompile("^s"),
			phonetic:     "",
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
			leftContext: regexp.MustCompile("[aeou]$"),
			phonetic:    "j",
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
			leftContext: regexp.MustCompile("[aeiou]$"),
			phonetic:    "j",
		},
		{
			pattern:      "ii",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:      "iy",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:      "yy",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "i",
		},
		{
			pattern:      "yi",
			rightContext: regexp.MustCompile("^$"),
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
	genspanish: []rule{
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
			pattern:      "h",
			rightContext: regexp.MustCompile("^$"),
			phonetic:     "",
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
			pattern:  "uo",
			phonetic: "(vo|o)",
		},
		{
			pattern:      "u",
			rightContext: regexp.MustCompile("^[aei]"),
			phonetic:     "v",
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
			pattern:  "b",
			phonetic: "B",
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
			phonetic: "(x|Z)",
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
			phonetic: "V",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "(ks|gz|S)",
		},
		{
			pattern:  "y",
			phonetic: "(i|j)",
		},
		{
			pattern:  "z",
			phonetic: "(z|s)",
		},
	},
	genturkish: []rule{
		{
			pattern:  "ç",
			phonetic: "tS",
		},
		{
			pattern:  "ğ",
			phonetic: "",
		},
		{
			pattern:  "ş",
			phonetic: "S",
		},
		{
			pattern:  "ü",
			phonetic: "Q",
		},
		{
			pattern:  "ö",
			phonetic: "Y",
		},
		{
			pattern:  "ı",
			phonetic: "(e|i|)",
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
			phonetic: "dZ",
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
			phonetic: "j",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
}

var genLangRules = []langRule{
	{
		pattern: regexp.MustCompile("^o’"),
		langs:   32,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^o'"),
		langs:   32,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^mc"),
		langs:   32,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^fitz"),
		langs:   32,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ceau"),
		langs:   65600,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("eau"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ault$"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("oult$"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("eux$"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("eix$"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("glou$"),
		langs:   512,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("uu"),
		langs:   16,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("tx"),
		langs:   262144,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("witz"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("tz$"),
		langs:   131232,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^tz"),
		langs:   131104,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("poulos$"),
		langs:   512,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("pulos$"),
		langs:   512,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("iou"),
		langs:   512,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("sj$"),
		langs:   16,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^sj"),
		langs:   16,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("güe"),
		langs:   262144,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("güi"),
		langs:   262144,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ghe"),
		langs:   66048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ghi"),
		langs:   66048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("escu$"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("esco$"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("vici$"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("schi$"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ii$"),
		langs:   131072,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("iy$"),
		langs:   131072,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("yy$"),
		langs:   131072,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("yi$"),
		langs:   131072,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^rz"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("rz$"),
		langs:   16512,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("[bcdfgklmnpstwz]rz"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("rz[bcdfghklmnpstw]"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("cki$"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ska$"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("cka$"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ae"),
		langs:   131232,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("oe"),
		langs:   131312,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("th$"),
		langs:   160,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^th"),
		langs:   672,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("mann"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("cz"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("cy"),
		langs:   16896,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("niew"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("etti$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("eti$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ati$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ato$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("[aoei]no$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("[aoei]ni$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("esi$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("oli$"),
		langs:   4096,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("field$"),
		langs:   32,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("stein"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("heim$"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("heimer$"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("thal"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("zweig"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("[aeou]h"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("äh"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("öh"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("üh"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("[ln]h[ao]$"),
		langs:   32768,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("[ln]h[aou]"),
		langs:   819416,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("chsch"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("tsch"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("sch$"),
		langs:   131200,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^sch"),
		langs:   131200,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ck$"),
		langs:   160,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("c$"),
		langs:   608264,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("sz"),
		langs:   18432,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("cs$"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^cs"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("dzs"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("zs$"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^zs"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^wl"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("^wr"),
		langs:   16560,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("gy$"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("gy[aeou]"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("gy"),
		langs:   133696,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("guy"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("gu[ei]"),
		langs:   294976,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("gu[ao]"),
		langs:   294912,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("gi[aou]"),
		langs:   4608,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ly"),
		langs:   150016,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ny"),
		langs:   412160,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ty"),
		langs:   150016,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ā"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ć"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ç"),
		langs:   819264,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("č"),
		langs:   8200,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ď"),
		langs:   8,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ē"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ğ"),
		langs:   524288,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ģ"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ī"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ķ"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ļ"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ł"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ņ"),
		langs:   8192,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ń"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ñ"),
		langs:   262144,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ň"),
		langs:   8,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ř"),
		langs:   8,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ś"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ş"),
		langs:   589824,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("š"),
		langs:   8200,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ţ"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ť"),
		langs:   8,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ź"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ž"),
		langs:   8200,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ż"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ß"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ä"),
		langs:   128,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("á"),
		langs:   297480,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("â"),
		langs:   98368,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ă"),
		langs:   65536,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ą"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("à"),
		langs:   32768,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ã"),
		langs:   32768,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ę"),
		langs:   16384,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("é"),
		langs:   2632,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("è"),
		langs:   266304,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ê"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ě"),
		langs:   8,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ê"),
		langs:   32832,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("í"),
		langs:   297480,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("î"),
		langs:   65600,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ı"),
		langs:   524288,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ó"),
		langs:   317960,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ö"),
		langs:   526464,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ô"),
		langs:   32832,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("õ"),
		langs:   34816,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ò"),
		langs:   266240,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ű"),
		langs:   2048,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ú"),
		langs:   297480,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ü"),
		langs:   821376,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ù"),
		langs:   64,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ů"),
		langs:   8,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ý"),
		langs:   520,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("а"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ё"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("о"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("е"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("и"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("у"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ы"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("э"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ю"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("я"),
		langs:   4,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("α"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ε"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("η"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ι"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ο"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("υ"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ω"),
		langs:   256,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ا"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ب"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ت"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ث"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ج"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ح"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("خ'"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("د"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ذ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ر"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ز"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("س"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ش"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ص"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ض"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ط"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ظ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ع"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("غ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ف"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ق"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ك"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ل"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("م"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ن"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ه"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("و"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ي"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("آ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("إ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("أ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ؤ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ئ"),
		langs:   2,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("א"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ב"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ג"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ד"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ה"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ו"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ז"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ח"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ט"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("י"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("כ"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ל"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("מ"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("נ"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ס"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ע"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("פ"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("צ"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ק"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ר"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ש"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("ת"),
		langs:   1024,
		accept:  true,
	},
	{
		pattern: regexp.MustCompile("a"),
		langs:   1286,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("o"),
		langs:   1286,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("e"),
		langs:   1286,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("i"),
		langs:   1286,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("y"),
		langs:   75030,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("u"),
		langs:   1286,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("j"),
		langs:   4096,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("j[^aoeiuy]"),
		langs:   295488,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("g"),
		langs:   8,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("k"),
		langs:   364608,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("q"),
		langs:   748056,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("v"),
		langs:   16384,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("w"),
		langs:   993864,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("x"),
		langs:   534552,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("dj"),
		langs:   786432,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("v[^aoeiu]"),
		langs:   128,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("y[^aoeiu]"),
		langs:   128,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("c[^aohk]"),
		langs:   128,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("dzi"),
		langs:   524512,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ou"),
		langs:   128,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("a[eiou]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ö[eaiou]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ü[eaiou]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("e[aiou]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("i[aeou]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("o[aieu]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("u[aieo]"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("aj"),
		langs:   240,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ej"),
		langs:   240,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("oj"),
		langs:   240,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("uj"),
		langs:   240,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("eu"),
		langs:   147456,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ky"),
		langs:   16384,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("kie"),
		langs:   262720,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("gie"),
		langs:   360960,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ch[aou]"),
		langs:   4096,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("ch"),
		langs:   524288,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("son$"),
		langs:   128,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("sc[ei]"),
		langs:   64,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("sch"),
		langs:   280640,
		accept:  false,
	},
	{
		pattern: regexp.MustCompile("^h"),
		langs:   131072,
		accept:  false,
	},
}

var genFinalRules = finalRules{
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
				rightContext: regexp.MustCompile("^[aA]"),
				phonetic:     "",
			},
			{
				pattern:     "a",
				leftContext: regexp.MustCompile("A$"),
				phonetic:    "",
			},
			{
				pattern:      "A",
				rightContext: regexp.MustCompile("^A"),
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
				pattern:      "j",
				rightContext: regexp.MustCompile("^j"),
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
				pattern:     "vanden",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(vanden|)",
			},
			{
				pattern:     "vander",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(vander|)",
			},
			{
				pattern:      "van",
				leftContext:  regexp.MustCompile("^$"),
				rightContext: regexp.MustCompile("^[bp]"),
				phonetic:     "(vam|[16])",
			},
			{
				pattern:     "van",
				leftContext: regexp.MustCompile("^$"),
				phonetic:    "(van|[16])",
			},
			{
				pattern:      "n",
				rightContext: regexp.MustCompile("^[bp]"),
				phonetic:     "m",
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
				pattern:      "sen",
				leftContext:  regexp.MustCompile("[rmnl]$"),
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(zn|zon)",
			},
			{
				pattern:      "sen",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(sn|son)",
			},
			{
				pattern:      "sEn",
				leftContext:  regexp.MustCompile("[rmnl]$"),
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(zn|zon)",
			},
			{
				pattern:      "sEn",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(sn|son)",
			},
			{
				pattern:      "e",
				leftContext:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "i",
				leftContext:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "E",
				leftContext:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "I",
				leftContext:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "Q",
				leftContext:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				rightContext: regexp.MustCompile("^[ln]$"),
				phonetic:     "",
			},
			{
				pattern:      "Y",
				leftContext:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
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
				pattern:      "burk",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "bUrk",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "burg",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "bUrg",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "Burk",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "BUrk",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "Burg",
				rightContext: regexp.MustCompile("^$"),
				phonetic:     "(burk|berk)",
			},
			{
				pattern:      "BUrg",
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
			uint64(genany): []rule{
				{
					pattern:  "mb",
					phonetic: "(mb|b[512])",
				},
				{
					pattern:  "mp",
					phonetic: "(mp|b[512])",
				},
				{
					pattern:  "ng",
					phonetic: "(ng|g[512])",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^[fktSs]"),
					phonetic:     "(p|f[262144])",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^p"),
					phonetic:     "",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(p|f[262144])",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^[pktSs]"),
					phonetic:     "(f|p[262144])",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^f"),
					phonetic:     "",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(f|p[262144])",
				},
				{
					pattern:  "B",
					phonetic: "(b|v[262144])",
				},
				{
					pattern:  "V",
					phonetic: "(v|b[262144])",
				},
				{
					pattern:      "t",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(t|[64])",
				},
				{
					pattern:      "g",
					leftContext:  regexp.MustCompile("n$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(g|[64])",
				},
				{
					pattern:      "k",
					leftContext:  regexp.MustCompile("n$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(k|[64])",
				},
				{
					pattern:      "p",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(p|[64])",
				},
				{
					pattern:      "r",
					leftContext:  regexp.MustCompile("[Ee]$"),
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(r|[64])",
				},
				{
					pattern:      "s",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(s|[64])",
				},
				{
					pattern:      "t",
					leftContext:  regexp.MustCompile("[aeiouAEIOU]$"),
					rightContext: regexp.MustCompile("^[^aeiouAEIOU]"),
					phonetic:     "(t|[64])",
				},
				{
					pattern:      "s",
					leftContext:  regexp.MustCompile("[aeiouAEIOU]$"),
					rightContext: regexp.MustCompile("^[^aeiouAEIOU]"),
					phonetic:     "(s|[64])",
				},
				{
					pattern:     "I",
					leftContext: regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					phonetic:    "i",
				},
				{
					pattern:      "I",
					rightContext: regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					phonetic:     "(Q[128]|i|D[32])",
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
					phonetic:     "(ik|Qk[128])",
				},
				{
					pattern:      "Ik",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "ik",
				},
				{
					pattern:      "sIts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(sits|sQts[128])",
				},
				{
					pattern:      "Its",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "its",
				},
				{
					pattern:  "I",
					phonetic: "(Q[128]|i)",
				},
				{
					pattern:     "lEE",
					leftContext: regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					phonetic:    "(li|il[32])",
				},
				{
					pattern:     "rEE",
					leftContext: regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					phonetic:    "(ri|ir[32])",
				},
				{
					pattern:     "lE",
					leftContext: regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					phonetic:    "(li|il[32]|lY[128])",
				},
				{
					pattern:     "rE",
					leftContext: regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					phonetic:    "(ri|ir[32]|rY[128])",
				},
				{
					pattern:  "EE",
					phonetic: "(i|)",
				},
				{
					pattern:  "ea",
					phonetic: "(D|a|i)",
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
					pattern:  "eu",
					phonetic: "(D|e|u)",
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
					phonetic: "(ia|io|iY[128])",
				},
				{
					pattern:      "A",
					rightContext: regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					phonetic:     "(a|o|Y[128]|D[32])",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("i[^aeiouAEIOU]$"),
					phonetic:    "(i|Y[128]|[32])",
				},
				{
					pattern:     "E",
					leftContext: regexp.MustCompile("a[^aeiouAEIOU]$"),
					phonetic:    "(i|Y[128]|[32])",
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
					phonetic: "(i|Y[128])",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
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
					phonetic: "(o|Y[128])",
				},
				{
					pattern:  "O",
					phonetic: "o",
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
					phonetic: "(a|o|Y[128])",
				},
				{
					pattern:  "A",
					phonetic: "(a|o)",
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
					phonetic:     "(uk|Qk[128])",
				},
				{
					pattern:      "Uk",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "uk",
				},
				{
					pattern:      "sUts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "(suts|sQts[128])",
				},
				{
					pattern:      "Uts",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "uts",
				},
				{
					pattern:  "U",
					phonetic: "(u|Q[128])",
				},
				{
					pattern:  "U",
					phonetic: "u",
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
					phonetic: "(i|Y[128])",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
			},
			uint64(genarabic): []rule{
				{
					pattern:  "1a",
					phonetic: "(D|a)",
				},
				{
					pattern:  "1i",
					phonetic: "(D|i|e)",
				},
				{
					pattern:  "1u",
					phonetic: "(D|u|o)",
				},
				{
					pattern:  "j1",
					phonetic: "(ja|je|jo|ju|j)",
				},
				{
					pattern:  "1",
					phonetic: "(a|e|i|o|u|)",
				},
				{
					pattern:  "u",
					phonetic: "(o|u)",
				},
				{
					pattern:  "i",
					phonetic: "(i|e)",
				},
				{
					pattern:      "p",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "p",
				},
				{
					pattern:  "p",
					phonetic: "(p|b)",
				},
			},
			uint64(genrussian): []rule{
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
			uint64(gencyrillic): []rule{
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
			uint64(genfrench): []rule{
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
			},
			uint64(genczech): []rule{
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
			},
			uint64(gendutch): []rule{
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
			},
			uint64(genenglish): []rule{
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
			uint64(gengerman): []rule{
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
			uint64(gengreek): []rule{
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
			},
			uint64(gengreeklatin): []rule{
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
					pattern:  "N",
					phonetic: "",
				},
			},
			uint64(genhebrew): []rule{},
			uint64(genhungarian): []rule{
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
			},
			uint64(genitalian): []rule{
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
			},
			uint64(genlatvian): []rule{
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
			},
			uint64(genpolish): []rule{
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
			uint64(genportuguese): []rule{
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
			},
			uint64(genromanian): []rule{
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
			uint64(genspanish): []rule{
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
					pattern:  "B",
					phonetic: "(b|v)",
				},
				{
					pattern:  "V",
					phonetic: "(b|v)",
				},
			},
			uint64(genturkish): []rule{
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
				rightContext: regexp.MustCompile("^[aA]"),
				phonetic:     "",
			},
			{
				pattern:     "a",
				leftContext: regexp.MustCompile("A$"),
				phonetic:    "",
			},
			{
				pattern:      "A",
				rightContext: regexp.MustCompile("^A"),
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
				pattern:      "j",
				rightContext: regexp.MustCompile("^j"),
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
			uint64(genany): []rule{
				{
					pattern:      "EE",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "e",
				},
				{
					pattern:  "A",
					phonetic: "a",
				},
				{
					pattern:  "E",
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
					pattern:      "B",
					rightContext: regexp.MustCompile("^[fktSs]"),
					phonetic:     "p",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^p"),
					phonetic:     "",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "p",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^[pktSs]"),
					phonetic:     "f",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^f"),
					phonetic:     "",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "f",
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
			uint64(genarabic): []rule{
				{
					pattern:  "1",
					phonetic: "",
				},
			},
			uint64(genrussian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(gencyrillic): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(genczech): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(gendutch): []rule{},
			uint64(genenglish): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(genfrench): []rule{},
			uint64(gengerman): []rule{
				{
					pattern:      "EE",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "e",
				},
				{
					pattern:  "A",
					phonetic: "a",
				},
				{
					pattern:  "E",
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
					pattern:      "B",
					rightContext: regexp.MustCompile("^[fktSs]"),
					phonetic:     "p",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^p"),
					phonetic:     "",
				},
				{
					pattern:      "B",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "p",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^[pktSs]"),
					phonetic:     "f",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^f"),
					phonetic:     "",
				},
				{
					pattern:      "V",
					rightContext: regexp.MustCompile("^$"),
					phonetic:     "f",
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
			uint64(gengreek): []rule{},
			uint64(gengreeklatin): []rule{
				{
					pattern:  "N",
					phonetic: "n",
				},
			},
			uint64(genhebrew): []rule{},
			uint64(genhungarian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(genitalian): []rule{},
			uint64(genlatvian): []rule{},
			uint64(genpolish): []rule{
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
			uint64(genportuguese): []rule{},
			uint64(genromanian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			uint64(genspanish): []rule{
				{
					pattern:  "B",
					phonetic: "b",
				},
				{
					pattern:  "V",
					phonetic: "v",
				},
			},
			uint64(genturkish): []rule{},
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
