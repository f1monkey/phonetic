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
			pattern: "yna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(in[131072]|ina)",
		},
		{
			pattern: "ina",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(in[131072]|ina)",
		},
		{
			pattern: "liova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lova|lof[131072]|lef[131072])",
		},
		{
			pattern: "lova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lova|lof[131072]|lef[131072]|l[8]|el[8])",
		},
		{
			pattern: "kova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(kova|kof[131072]|k[8]|ek[8])",
		},
		{
			pattern: "ova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ova|of[131072]|[8])",
		},
		{
			pattern: "ová",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ova|[8])",
		},
		{
			pattern: "eva",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(eva|ef[131072])",
		},
		{
			pattern: "aia",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(aja|i[131072])",
		},
		{
			pattern: "aja",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(aja|i[131072])",
		},
		{
			pattern: "aya",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(aja|i[131072])",
		},
		{
			pattern: "lowa",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lova|lof[16384]|l[16384]|el[16384])",
		},
		{
			pattern: "kowa",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(kova|kof[16384]|k[16384]|ek[16384])",
		},
		{
			pattern: "owa",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ova|of[16384]|)",
		},
		{
			pattern: "lowna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lovna|levna|l[16384]|el[16384])",
		},
		{
			pattern: "kowna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(kovna|k[16384]|ek[16384])",
		},
		{
			pattern: "owna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ovna|[16384])",
		},
		{
			pattern: "lówna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(l|el)",
		},
		{
			pattern: "kówna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(k|ek)",
		},
		{
			pattern: "ówna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "á",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(a|i[8])",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(a|i[16392])",
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
		},
		{
			pattern: "que",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(k[64]|ke|kve)",
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
		},
		{
			pattern: "m",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bfpv]"),
			},
			phonetic: "(m|n)",
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiouy]"),
			},
			phonetic: "m",
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(m|n[32832])",
		},
		{
			pattern: "ly",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[au]"),
			},
			phonetic: "(l|lj)",
		},
		{
			pattern: "li",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[au]"),
			},
			phonetic: "(l|lj)",
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
			pattern: "lt",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "u",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lt|[64])",
		},
		{
			pattern: "v",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(v|f[128]|b[262144])",
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "(ez[32768]|eS[32768]|eks|egz)",
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[cs]"),
			},
			phonetic: "(e[32768]|ek)",
		},
		{
			pattern: "x",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "u",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ks|[64])",
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
			pattern: "rh",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "r",
		},
		{
			pattern: "dh",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "d",
		},
		{
			pattern: "bh",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "b",
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
			pattern: "sch",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(S|StS[131072]|sk[69632])",
		},
		{
			pattern: "sch",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(S|StS[131072])",
		},
		{
			pattern: "sch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(sk[69632]|S|StS[131072])",
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
			pattern: "sh",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[äöü]"),
			},
			phonetic: "sh",
		},
		{
			pattern: "sh",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "(S[131104]|sh)",
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
			pattern: "ch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(x|tS[393248]|k[69632]|S[32832])",
		},
		{
			pattern:  "ch",
			phonetic: "(x|tS[393248]|S[32832])",
		},
		{
			pattern: "th",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "t",
		},
		{
			pattern: "th",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[äöüaeiou]"),
			},
			phonetic: "(t[672]|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern: "gh",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(g[70144]|gh)",
		},
		{
			pattern: "ouh",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aioe]"),
			},
			phonetic: "(v[64]|uh)",
		},
		{
			pattern: "uh",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aioe]"),
			},
			phonetic: "(v|uh)",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouyäöü]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(h|x[66048]|H[381024])",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa[16384]|tsa)",
		},
		{
			pattern: "cią",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSom|tsom)",
		},
		{
			pattern:  "cią",
			phonetic: "(tSon[16384]|tson)",
		},
		{
			pattern: "cię",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSem[16384]|tsem)",
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
			pattern: "sci",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(Si[4096]|stsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
		},
		{
			pattern: "sc",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(S[4096]|sts[16392]|dZ[524288]|tS[81920]|s)",
		},
		{
			pattern: "ci",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(tsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
		},
		{
			pattern:  "cy",
			phonetic: "(si|tsi[16384])",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(ts[16392]|dZ[524288]|tS[81920]|k[512]|s)",
		},
		{
			pattern: "sç",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "(s|stS[524288])",
		},
		{
			pattern:  "ssz",
			phonetic: "S",
		},
		{
			pattern: "sz",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(S|s[2048])",
		},
		{
			pattern: "sz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(S|s[2048])",
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
			pattern: "sj",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "S",
		},
		{
			pattern: "sj",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "S",
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
			pattern: "sią",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Som[16384]|som)",
		},
		{
			pattern:  "sią",
			phonetic: "(Son[16384]|son)",
		},
		{
			pattern: "się",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Sem[16384]|sem)",
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
			pattern: "si",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[äöëaáuiíoóeéêy]$"),
			},
			phonetic: "(Si[16384]|si|zi[37056])",
		},
		{
			pattern:  "si",
			phonetic: "(Si[16384]|si|zi[128])",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aáuíoóeéêy]"),
			},
			phonetic: "(s|z[37056])",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeouäöë]"),
			},
			phonetic: "(s|z[128])",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[dglmnrv]"),
			},
			phonetic: "(s|z|Z[32768]|[64])",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[dglmnrv]"),
			},
			phonetic: "(s|z|Z[32768])",
		},
		{
			pattern: "gue",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(k[64]|gve)",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(g[64]|gv[294912])",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ao]"),
			},
			phonetic: "gv",
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
			pattern: "gn",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "(n[4160]|nj[4160]|gn)",
		},
		{
			pattern:  "ggie",
			phonetic: "(je[512]|dZe)",
		},
		{
			pattern: "ggi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "(j[512]|dZ)",
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[yaeiou]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "(gI|dZ[4096]|j[512])",
		},
		{
			pattern: "gge",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gE|xe[262144]|gZe[32832]|dZe[331808]|je[512])",
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gI|xi[262144]|gZi[32832]|dZi[331808]|i[512])",
		},
		{
			pattern: "ggi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "(gI|dZ[4096]|j[512])",
		},
		{
			pattern: "gie",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ge|gi[128]|ji[64]|dZe[4096])",
		},
		{
			pattern:  "gie",
			phonetic: "(ge|gi[128]|dZe[4096]|je[512])",
		},
		{
			pattern: "gi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "(i[512]|dZ)",
		},
		{
			pattern: "ge",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gE|xe[262144]|Ze[32832]|dZe[331808])",
		},
		{
			pattern: "gi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gI|xi[262144]|Zi[32832]|dZi[331808])",
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
			pattern: "gy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phonetic: "(gi|dj[2048])",
		},
		{
			pattern:  "gy",
			phonetic: "(gi|d[2048])",
		},
		{
			pattern: "g",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[yaeiou]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouyei]"),
			},
			phonetic: "g",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouei]"),
			},
			phonetic: "(g|h[131072])",
		},
		{
			pattern:  "ij",
			phonetic: "(i|ej[16]|ix[262144]|iZ[622656])",
		},
		{
			pattern: "j",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aoeiuy]"),
			},
			phonetic: "(j|dZ[32]|x[262144]|Z[622656])",
		},
		{
			pattern: "rz",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "t",
			},
			phonetic: "(S[16384]|r)",
		},
		{
			pattern:  "rz",
			phonetic: "(rz|rts[128]|Z[16384]|r[16384]|rZ[16384])",
		},
		{
			pattern: "tz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ts|tS[160])",
		},
		{
			pattern: "tz",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ts[131232]|tS[160])",
		},
		{
			pattern:  "tz",
			phonetic: "(ts[131232]|tz)",
		},
		{
			pattern: "zia",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(Za[16384]|za[16384]|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za[16384]|zja)",
		},
		{
			pattern: "zią",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zom[16384]|zom)",
		},
		{
			pattern:  "zią",
			phonetic: "(Zon[16384]|zon)",
		},
		{
			pattern: "zię",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zem[16384]|zem)",
		},
		{
			pattern:  "zię",
			phonetic: "(Zen[16384]|zen)",
		},
		{
			pattern: "zie",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(Ze[16384]|ze[16384]|ze|tsi[128])",
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
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(s|ts[128]|ts[4096]|S[32768])",
		},
		{
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bdgv]"),
			},
			phonetic: "(z|dz[4096]|Z[32768])",
		},
		{
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ptckf]"),
			},
			phonetic: "(s|ts[4096]|S[32768])",
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
			pattern: "ii",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
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
			pattern: "iy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
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
			pattern: "yi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
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
			pattern: "yy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[áóéê]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[áóéê]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(e|je[131072])",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(e|EE[96])",
		},
		{
			pattern: "ą",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "om",
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
			pattern: "ę",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "em",
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
			pattern: "o",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phonetic: "(O|P[16384])",
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
			pattern: "ب",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "b",
		},
		{
			pattern:  "ب",
			phonetic: "b1",
		},
		{
			pattern: "ت",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "t",
		},
		{
			pattern:  "ت",
			phonetic: "t1",
		},
		{
			pattern: "ث",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "t",
		},
		{
			pattern:  "ث",
			phonetic: "t1",
		},
		{
			pattern: "ج",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(dZ|Z)",
		},
		{
			pattern:  "ج",
			phonetic: "(dZ1|Z1)",
		},
		{
			pattern: "ح",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern: "ح",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern:  "ح",
			phonetic: "(h1|1)",
		},
		{
			pattern: "خ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "x",
		},
		{
			pattern:  "خ",
			phonetic: "x1",
		},
		{
			pattern: "د",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "d",
		},
		{
			pattern:  "د",
			phonetic: "d1",
		},
		{
			pattern: "ذ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "d",
		},
		{
			pattern:  "ذ",
			phonetic: "d1",
		},
		{
			pattern: "ر",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "r",
		},
		{
			pattern:  "ر",
			phonetic: "r1",
		},
		{
			pattern: "ز",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "z",
		},
		{
			pattern:  "ز",
			phonetic: "z1",
		},
		{
			pattern: "س",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "s",
		},
		{
			pattern:  "س",
			phonetic: "s1",
		},
		{
			pattern: "ش",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "S",
		},
		{
			pattern:  "ش",
			phonetic: "S1",
		},
		{
			pattern: "ص",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "s",
		},
		{
			pattern:  "ص",
			phonetic: "s1",
		},
		{
			pattern: "ض",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "d",
		},
		{
			pattern:  "ض",
			phonetic: "d1",
		},
		{
			pattern: "ط",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "t",
		},
		{
			pattern:  "ط",
			phonetic: "t1",
		},
		{
			pattern: "ظ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "z",
		},
		{
			pattern:  "ظ",
			phonetic: "z1",
		},
		{
			pattern: "ع",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern: "ع",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern:  "ع",
			phonetic: "(h1|1)",
		},
		{
			pattern: "غ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "g",
		},
		{
			pattern:  "غ",
			phonetic: "g1",
		},
		{
			pattern: "ف",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "f",
		},
		{
			pattern:  "ف",
			phonetic: "f1",
		},
		{
			pattern: "ق",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "k",
		},
		{
			pattern:  "ق",
			phonetic: "k1",
		},
		{
			pattern: "ك",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "k",
		},
		{
			pattern:  "ك",
			phonetic: "k1",
		},
		{
			pattern: "ل",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "l",
		},
		{
			pattern:  "ل",
			phonetic: "l1",
		},
		{
			pattern: "م",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "m",
		},
		{
			pattern:  "م",
			phonetic: "m1",
		},
		{
			pattern: "ن",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "n",
		},
		{
			pattern:  "ن",
			phonetic: "n1",
		},
		{
			pattern: "ه",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern: "ه",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern:  "ه",
			phonetic: "(h1|1)",
		},
		{
			pattern: "و",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(u|v)",
		},
		{
			pattern:  "و",
			phonetic: "(u|v1)",
		},
		{
			pattern: "ي\u200e",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(i|j)",
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
			pattern: "с",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "с",
				suffix:   "",
			},
			phonetic: "",
		},
		{
			pattern: "гауз",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "haus",
		},
		{
			pattern: "гаус",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "haus",
		},
		{
			pattern: "гольц",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "holts",
		},
		{
			pattern: "геймер",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(hejmer|hajmer)",
		},
		{
			pattern: "гейм",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(hejm|hajm)",
		},
		{
			pattern: "гоф",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "hof",
		},
		{
			pattern: "гер",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ger",
		},
		{
			pattern: "ген",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "gen",
		},
		{
			pattern: "гин",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "gin",
		},
		{
			pattern: "г",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(й|ё|я|ю|ы|а|е|о|и|у)$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phonetic: "g",
		},
		{
			pattern: "г",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(а|е|о|и|у)"),
			},
			phonetic: "(g|h)",
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
			pattern: "ий",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(а|о|у)"),
			},
			phonetic: "j",
		},
		{
			pattern: "ый",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(а|о|у)"),
			},
			phonetic: "j",
		},
		{
			pattern: "ий",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "ый",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "ей",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(jej|ej)",
		},
		{
			pattern: "е",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(а|е|о|у)$"),
			},
			phonetic: "je",
		},
		{
			pattern: "е",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "je",
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
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
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
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiy]"),
			},
			phonetic: "ts",
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
			pattern: "th",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "t",
		},
		{
			pattern: "th",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[äöüaeiou]"),
			},
			phonetic: "(t|th)",
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
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "",
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
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
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
			pattern: "mc",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "mak",
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
			pattern: "cc",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[iey]"),
			},
			phonetic: "ks",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "c",
				suffix:   "",
			},
			phonetic: "",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[iey]"),
			},
			phonetic: "s",
		},
		{
			pattern: "gh",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "g",
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
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[iey]"),
			},
			phonetic: "(g|dZ)",
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
			pattern: "who",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "hu",
		},
		{
			pattern: "wh",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "w",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeiou]"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "H",
		},
		{
			pattern: "kn",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "n",
		},
		{
			pattern: "mb",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "m",
		},
		{
			pattern: "ng",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(N|ng)",
		},
		{
			pattern: "pn",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(pn|n)",
		},
		{
			pattern: "ps",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ps|s)",
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
			pattern: "wr",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "r",
		},
		{
			pattern: "x",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "z",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiouy]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
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
			pattern: "a",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "ej",
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
			pattern: "e",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "i",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(|E)",
		},
		{
			pattern:  "ie",
			phonetic: "i",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "aj",
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
			pattern: "o",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "ou",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "(ju|u)",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "r",
				suffix:   "",
			},
			phonetic: "(e|u)",
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
			pattern: "lt",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "u",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lt|)",
		},
		{
			pattern: "c",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "n",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(k|)",
		},
		{
			pattern: "d",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(t|)",
		},
		{
			pattern: "g",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "n",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(k|)",
		},
		{
			pattern: "p",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(p|)",
		},
		{
			pattern: "r",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "e",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(r|)",
		},
		{
			pattern: "t",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(t|)",
		},
		{
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(s|)",
		},
		{
			pattern: "ds",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ds|)",
		},
		{
			pattern: "ps",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ps|)",
		},
		{
			pattern: "rs",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "e",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(rs|)",
		},
		{
			pattern: "ts",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ts|)",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(s|)",
		},
		{
			pattern: "x",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "u",
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ks|)",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeéèêiou]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeéèêiou]"),
			},
			phonetic: "(s|)",
		},
		{
			pattern: "t",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeéèêiou]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[^aeéèêiou]"),
			},
			phonetic: "(t|)",
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
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiyéèê]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "gn",
			phonetic: "(n|gn)",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiy]"),
			},
			phonetic: "Z",
		},
		{
			pattern: "gue",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "k",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiy]"),
			},
			phonetic: "g",
		},
		{
			pattern: "aill",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "e",
				suffix:   "",
			},
			phonetic: "aj",
		},
		{
			pattern: "ll",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "e",
				suffix:   "",
			},
			phonetic: "(l|j)",
		},
		{
			pattern: "que",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "k",
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouyéèê]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiouyéèê]"),
			},
			phonetic: "z",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[bdgt]$"),
			},
			phonetic: "",
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiouy]"),
			},
			phonetic: "m",
		},
		{
			pattern: "m",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(m|n)",
		},
		{
			pattern: "ou",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeio]"),
			},
			phonetic: "v",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeio]"),
			},
			phonetic: "v",
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
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[ou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(e|)",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aoeu]"),
			},
			phonetic: "j",
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
			pattern: "ewitsch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "evitS",
		},
		{
			pattern: "owitsch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ovitS",
		},
		{
			pattern: "evitsch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "evitS",
		},
		{
			pattern: "ovitsch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ovitS",
		},
		{
			pattern: "witsch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "vitS",
		},
		{
			pattern: "vitsch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "vitS",
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
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiy]"),
			},
			phonetic: "ts",
		},
		{
			pattern: "sp",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "Sp",
		},
		{
			pattern: "st",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "St",
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
			pattern: "ewitz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "ewiz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "evitz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "eviz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "owitz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "owiz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "ovitz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "oviz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "witz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern: "wiz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern: "vitz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern: "viz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern:  "tz",
			phonetic: "ts",
		},
		{
			pattern: "thal",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "tal",
		},
		{
			pattern: "th",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "t",
		},
		{
			pattern: "th",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[äöüaeiou]"),
			},
			phonetic: "(t|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern: "rh",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "r",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouyäöü]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "H",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[äöüaeiouy]"),
			},
			phonetic: "(z|s)",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouyäöüj]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiouyäöü]"),
			},
			phonetic: "z",
		},
		{
			pattern:  "ß",
			phonetic: "s",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
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
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern:  "ie",
			phonetic: "I",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aoeu]"),
			},
			phonetic: "j",
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
			pattern: "αυ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "af",
		},
		{
			pattern: "αυ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phonetic: "af",
		},
		{
			pattern:  "αυ",
			phonetic: "av",
		},
		{
			pattern: "ευ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ef",
		},
		{
			pattern: "ευ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phonetic: "ef",
		},
		{
			pattern:  "ευ",
			phonetic: "ev",
		},
		{
			pattern: "ηυ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "if",
		},
		{
			pattern: "ηυ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(κ|π|σ|τ|φ|θ|χ|ψ)"),
			},
			phonetic: "if",
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
			pattern: "γγ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "(nj|j)",
		},
		{
			pattern: "γγ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "j",
		},
		{
			pattern: "γγ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "(ng|g)",
		},
		{
			pattern:  "γγ",
			phonetic: "g",
		},
		{
			pattern: "γκ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "g",
		},
		{
			pattern: "γκ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "(nj|j)",
		},
		{
			pattern: "γκ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(ε|ι|η)"),
			},
			phonetic: "j",
		},
		{
			pattern: "γκ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "(ng|g)",
		},
		{
			pattern:  "γκ",
			phonetic: "g",
		},
		{
			pattern: "γι",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phonetic: "j",
		},
		{
			pattern:  "γι",
			phonetic: "(gi|i)",
		},
		{
			pattern: "γε",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(α|ο|ω|υ)"),
			},
			phonetic: "j",
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
			pattern: "σ",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^(β|γ|δ|μ|ν|ρ)"),
			},
			phonetic: "z",
		},
		{
			pattern:  "μβ",
			phonetic: "(mb|b)",
		},
		{
			pattern: "μπ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "b",
		},
		{
			pattern: "μπ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "mb",
		},
		{
			pattern:  "μπ",
			phonetic: "b",
		},
		{
			pattern: "ντ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "d",
		},
		{
			pattern: "ντ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("(ε|ι|η|α|ο|ω|υ)$"),
			},
			phonetic: "(nd|nt)",
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
			pattern: "au",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "af",
		},
		{
			pattern: "au",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[kpstfh]"),
			},
			phonetic: "af",
		},
		{
			pattern:  "au",
			phonetic: "av",
		},
		{
			pattern: "eu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ef",
		},
		{
			pattern: "eu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[kpstfh]"),
			},
			phonetic: "ef",
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
			pattern: "gge",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(nje|je)",
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "(nj|j)",
		},
		{
			pattern: "ggi",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ni|i)",
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
			pattern: "gg",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ng|g)",
		},
		{
			pattern:  "gg",
			phonetic: "g",
		},
		{
			pattern: "gk",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "g",
		},
		{
			pattern: "gke",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(nje|je)",
		},
		{
			pattern: "gki",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ni|i)",
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
			pattern: "gk",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(ng|g)",
		},
		{
			pattern:  "gk",
			phonetic: "g",
		},
		{
			pattern: "nghi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
		},
		{
			pattern:  "nghi",
			phonetic: "(Ngi|Ni)",
		},
		{
			pattern: "nghe",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
		},
		{
			pattern:  "nghe",
			phonetic: "(Nje|Nge)",
		},
		{
			pattern: "ghi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "ghi",
			phonetic: "(gi|i)",
		},
		{
			pattern: "ghe",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
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
			pattern: "ngi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
		},
		{
			pattern:  "ngi",
			phonetic: "(Ngi|Ni)",
		},
		{
			pattern: "nge",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "Nj",
		},
		{
			pattern:  "nge",
			phonetic: "(Nje|Nge)",
		},
		{
			pattern: "gi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "gi",
			phonetic: "(gi|i)",
		},
		{
			pattern: "ge",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aouy]"),
			},
			phonetic: "j",
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
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
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
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bgdmnr]"),
			},
			phonetic: "z",
		},
		{
			pattern:  "mb",
			phonetic: "(mb|b)",
		},
		{
			pattern: "mp",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "b",
		},
		{
			pattern: "mp",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "mp",
		},
		{
			pattern:  "mp",
			phonetic: "b",
		},
		{
			pattern: "nt",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "d",
		},
		{
			pattern: "nt",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiouy]$"),
			},
			phonetic: "(nd|nt)",
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
			pattern: "ה",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
		},
		{
			pattern: "ה",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "1",
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
			pattern: "כ",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "K",
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
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[áo]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[áo]$"),
			},
			phonetic: "j",
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
			pattern: "gy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phonetic: "dj",
		},
		{
			pattern:  "gy",
			phonetic: "(d|gi)",
		},
		{
			pattern: "ny",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phonetic: "nj",
		},
		{
			pattern:  "ny",
			phonetic: "(n|ni)",
		},
		{
			pattern: "ty",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeouáéóúüöőű]"),
			},
			phonetic: "tj",
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
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
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
			pattern: "gn",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "(n|nj|gn)",
		},
		{
			pattern:  "gni",
			phonetic: "(ni|gni)",
		},
		{
			pattern: "gi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "dZ",
		},
		{
			pattern: "gg",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "dZ",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "dZ",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[bdgt]$"),
			},
			phonetic: "g",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "ci",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "tS",
		},
		{
			pattern: "ch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "k",
		},
		{
			pattern: "sc",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "S",
		},
		{
			pattern: "cc",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "tS",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "tS",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiou]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "z",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
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
			pattern: "u",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aei]"),
			},
			phonetic: "v",
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
			pattern: "ska",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ski",
		},
		{
			pattern: "cka",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "tski",
		},
		{
			pattern: "lowa",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lova|lof|l|el)",
		},
		{
			pattern: "kowa",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(kova|kof|k|ek)",
		},
		{
			pattern: "owa",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ova|of|)",
		},
		{
			pattern: "lowna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lovna|levna|l|el)",
		},
		{
			pattern: "kowna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(kovna|k|ek)",
		},
		{
			pattern: "owna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ovna|)",
		},
		{
			pattern: "lówna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(l|el)",
		},
		{
			pattern: "kówna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(k|ek)",
		},
		{
			pattern: "ówna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(a|i)",
		},
		{
			pattern:  "czy",
			phonetic: "tSi",
		},
		{
			pattern: "cze",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(tSe|tSF)",
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
			pattern: "cia",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(tSB|tsB)",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa|tsa)",
		},
		{
			pattern: "cią",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSom|tsom)",
		},
		{
			pattern:  "cią",
			phonetic: "(tSon|tson)",
		},
		{
			pattern: "cię",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSem|tsem)",
		},
		{
			pattern:  "cię",
			phonetic: "(tSen|tsen)",
		},
		{
			pattern: "cie",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(tSF|tsF)",
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
			pattern: "sia",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(SB|sB|sja)",
		},
		{
			pattern:  "sia",
			phonetic: "(Sa|sja)",
		},
		{
			pattern: "sią",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Som|som)",
		},
		{
			pattern:  "sią",
			phonetic: "(Son|son)",
		},
		{
			pattern: "się",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Sem|sem)",
		},
		{
			pattern:  "się",
			phonetic: "(Sen|sen)",
		},
		{
			pattern: "sie",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(SF|sF|se)",
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
			pattern: "zia",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(ZB|zB|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za|zja)",
		},
		{
			pattern: "zią",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zom|zom)",
		},
		{
			pattern:  "zią",
			phonetic: "(Zon|zon)",
		},
		{
			pattern: "zię",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zem|zem)",
		},
		{
			pattern:  "zię",
			phonetic: "(Zen|zen)",
		},
		{
			pattern: "zie",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(ZF|zF)",
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
			pattern: "że",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(Ze|ZF)",
		},
		{
			pattern: "że",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(Ze|ZF|ze|zF)",
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
			pattern: "rze",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "t",
			},
			phonetic: "(Se|re)",
		},
		{
			pattern:  "rze",
			phonetic: "(Ze|re|rZe)",
		},
		{
			pattern: "rzy",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "t",
			},
			phonetic: "(Si|ri)",
		},
		{
			pattern:  "rzy",
			phonetic: "(Zi|ri|rZi)",
		},
		{
			pattern: "rz",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "t",
			},
			phonetic: "(S|r)",
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
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "s",
				suffix:   "",
			},
			phonetic: "",
		},
		{
			pattern:  "ó",
			phonetic: "(u|o)",
		},
		{
			pattern: "ą",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "om",
		},
		{
			pattern: "ę",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bp]"),
			},
			phonetic: "em",
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
			pattern: "ij",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
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
			pattern: "ie",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "F",
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
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "B",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdgkpstwzż]"),
			},
			phonetic: "(E|F)",
		},
		{
			pattern: "o",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcćdgklłmnńrsśtwzźż]"),
			},
			phonetic: "P",
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
			pattern: "sc",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "s",
		},
		{
			pattern: "sç",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "ç",
			phonetic: "s",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "s",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "s",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aáuiíoóeéêy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "z",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[dglmnrv]"),
			},
			phonetic: "(Z|S)",
		},
		{
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(Z|s|S)",
		},
		{
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bdgv]"),
			},
			phonetic: "(Z|z)",
		},
		{
			pattern: "z",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ptckf]"),
			},
			phonetic: "(s|S|z)",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiu]"),
			},
			phonetic: "g",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ao]"),
			},
			phonetic: "gv",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "Z",
		},
		{
			pattern: "qu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[eiu]"),
			},
			phonetic: "k",
		},
		{
			pattern: "qu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ao]"),
			},
			phonetic: "kv",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o|u)",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aei]"),
			},
			phonetic: "v",
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
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[bdgt]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aáuiíoóeéêy]"),
			},
			phonetic: "(ez|eS|eks)",
		},
		{
			pattern: "ex",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[cs]"),
			},
			phonetic: "e",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aáuiíoóeéê]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeiíou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "m",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bcdfglnprstv]"),
			},
			phonetic: "(m|n)",
		},
		{
			pattern: "m",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(m|n)",
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
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aáuoóeéê]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
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
			pattern: "ch",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "k",
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
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "dZ",
		},
		{
			pattern:  "gh",
			phonetic: "g",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
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
			pattern: "yna",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(in|ina)",
		},
		{
			pattern: "ina",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(in|ina)",
		},
		{
			pattern: "liova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lof|lef)",
		},
		{
			pattern: "lova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(lof|lef|lova)",
		},
		{
			pattern: "ova",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(of|ova)",
		},
		{
			pattern: "eva",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(ef|ova)",
		},
		{
			pattern: "aia",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(aja|i)",
		},
		{
			pattern: "aja",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(aja|i)",
		},
		{
			pattern: "aya",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(aja|i)",
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
			pattern: "ger",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ger",
		},
		{
			pattern: "gen",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "gen",
		},
		{
			pattern: "gin",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "gin",
		},
		{
			pattern:  "gg",
			phonetic: "g",
		},
		{
			pattern: "g",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[jaeoiuy]$"),
			},
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeoiu]"),
			},
			phonetic: "g",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aeoiu]"),
			},
			phonetic: "(g|h)",
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
			pattern: "tz",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "ts",
		},
		{
			pattern:  "tz",
			phonetic: "(ts|tz)",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[iey]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "s",
				suffix:   "",
			},
			phonetic: "",
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
			pattern: "ij",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "io",
			phonetic: "(jo|e)",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[au]"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeou]$"),
			},
			phonetic: "j",
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e)",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[au]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "i",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "(je|E)",
		},
		{
			pattern:  "ee",
			phonetic: "(aje|i)",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aou]$"),
			},
			phonetic: "je",
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
			pattern: "ig",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiou]$"),
			},
			phonetic: "(tS|ig)",
		},
		{
			pattern: "ix",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[aeiou]$"),
			},
			phonetic: "S",
		},
		{
			pattern:  "tx",
			phonetic: "tS",
		},
		{
			pattern: "tj",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "tS",
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
			pattern: "h",
			leftContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("[dgt]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^$"),
			},
			phonetic: "",
		},
		{
			pattern: "m",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[bpvf]"),
			},
			phonetic: "(m|n)",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "s",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(g|gv)",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[ei]"),
			},
			phonetic: "(x|g|dZ)",
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
			pattern: "u",
			rightContext: &ruleMatcher{
				contains: "",
				prefix:   "",
				suffix:   "",
				pattern:  regexp.MustCompile("^[aei]"),
			},
			phonetic: "v",
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
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("^o’"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("^o'"),
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "mc",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "fitz",
			suffix:   "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ceau",
			prefix:   "",
			suffix:   "",
		},
		langs:  65600,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "eau",
			prefix:   "",
			suffix:   "",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "ault",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "oult",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "eux",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "eix",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "glou",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "uu",
			prefix:   "",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "tx",
			prefix:   "",
			suffix:   "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "witz",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "tz",
		},
		langs:  131232,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "tz",
			suffix:   "",
		},
		langs:  131104,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "poulos",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "pulos",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "iou",
			prefix:   "",
			suffix:   "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "sj",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "sj",
			suffix:   "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "güe",
			prefix:   "",
			suffix:   "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "güi",
			prefix:   "",
			suffix:   "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ghe",
			prefix:   "",
			suffix:   "",
		},
		langs:  66048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ghi",
			prefix:   "",
			suffix:   "",
		},
		langs:  66048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "escu",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "esco",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "vici",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "schi",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "ii",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "iy",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "yy",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "yi",
		},
		langs:  131072,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "rz",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "rz",
		},
		langs:  16512,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[bcdfgklmnpstwz]rz"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("rz[bcdfghklmnpstw]"),
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "cki",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "ska",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "cka",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ae",
			prefix:   "",
			suffix:   "",
		},
		langs:  131232,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "oe",
			prefix:   "",
			suffix:   "",
		},
		langs:  131312,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "th",
		},
		langs:  160,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "th",
			suffix:   "",
		},
		langs:  672,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "mann",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "cz",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "cy",
			prefix:   "",
			suffix:   "",
		},
		langs:  16896,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "niew",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "etti",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "eti",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "ati",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "ato",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[aoei]no$"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[aoei]ni$"),
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "esi",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "oli",
		},
		langs:  4096,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "field",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "stein",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "heim",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "heimer",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "thal",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "zweig",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[aeou]h"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "äh",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "öh",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "üh",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[ln]h[ao]$"),
		},
		langs:  32768,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("[ln]h[aou]"),
		},
		langs:  819416,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "chsch",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "tsch",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "sch",
		},
		langs:  131200,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "sch",
			suffix:   "",
		},
		langs:  131200,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "ck",
		},
		langs:  160,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "c",
		},
		langs:  608264,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "sz",
			prefix:   "",
			suffix:   "",
		},
		langs:  18432,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "cs",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "cs",
			suffix:   "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "dzs",
			prefix:   "",
			suffix:   "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "zs",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "zs",
			suffix:   "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "wl",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "wr",
			suffix:   "",
		},
		langs:  16560,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "gy",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gy[aeou]"),
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "gy",
			prefix:   "",
			suffix:   "",
		},
		langs:  133696,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "guy",
			prefix:   "",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gu[ei]"),
		},
		langs:  294976,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gu[ao]"),
		},
		langs:  294912,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("gi[aou]"),
		},
		langs:  4608,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ly",
			prefix:   "",
			suffix:   "",
		},
		langs:  150016,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ny",
			prefix:   "",
			suffix:   "",
		},
		langs:  412160,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ty",
			prefix:   "",
			suffix:   "",
		},
		langs:  150016,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ā",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ć",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ç",
			prefix:   "",
			suffix:   "",
		},
		langs:  819264,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "č",
			prefix:   "",
			suffix:   "",
		},
		langs:  8200,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ď",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ē",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ğ",
			prefix:   "",
			suffix:   "",
		},
		langs:  524288,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ģ",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ī",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ķ",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ļ",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ł",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ņ",
			prefix:   "",
			suffix:   "",
		},
		langs:  8192,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ń",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ñ",
			prefix:   "",
			suffix:   "",
		},
		langs:  262144,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ň",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ř",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ś",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ş",
			prefix:   "",
			suffix:   "",
		},
		langs:  589824,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "š",
			prefix:   "",
			suffix:   "",
		},
		langs:  8200,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ţ",
			prefix:   "",
			suffix:   "",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ť",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ź",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ž",
			prefix:   "",
			suffix:   "",
		},
		langs:  8200,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ż",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ß",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ä",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "á",
			prefix:   "",
			suffix:   "",
		},
		langs:  297480,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "â",
			prefix:   "",
			suffix:   "",
		},
		langs:  98368,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ă",
			prefix:   "",
			suffix:   "",
		},
		langs:  65536,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ą",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "à",
			prefix:   "",
			suffix:   "",
		},
		langs:  32768,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ã",
			prefix:   "",
			suffix:   "",
		},
		langs:  32768,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ę",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "é",
			prefix:   "",
			suffix:   "",
		},
		langs:  2632,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "è",
			prefix:   "",
			suffix:   "",
		},
		langs:  266304,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ê",
			prefix:   "",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ě",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ê",
			prefix:   "",
			suffix:   "",
		},
		langs:  32832,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "í",
			prefix:   "",
			suffix:   "",
		},
		langs:  297480,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "î",
			prefix:   "",
			suffix:   "",
		},
		langs:  65600,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ı",
			prefix:   "",
			suffix:   "",
		},
		langs:  524288,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ó",
			prefix:   "",
			suffix:   "",
		},
		langs:  317960,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ö",
			prefix:   "",
			suffix:   "",
		},
		langs:  526464,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ô",
			prefix:   "",
			suffix:   "",
		},
		langs:  32832,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "õ",
			prefix:   "",
			suffix:   "",
		},
		langs:  34816,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ò",
			prefix:   "",
			suffix:   "",
		},
		langs:  266240,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ű",
			prefix:   "",
			suffix:   "",
		},
		langs:  2048,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ú",
			prefix:   "",
			suffix:   "",
		},
		langs:  297480,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ü",
			prefix:   "",
			suffix:   "",
		},
		langs:  821376,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ù",
			prefix:   "",
			suffix:   "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ů",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ý",
			prefix:   "",
			suffix:   "",
		},
		langs:  520,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "а",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ё",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "о",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "е",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "и",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "у",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ы",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "э",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ю",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "я",
			prefix:   "",
			suffix:   "",
		},
		langs:  4,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "α",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ε",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "η",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ι",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ο",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "υ",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ω",
			prefix:   "",
			suffix:   "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ا",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ب",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ت",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ث",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ج",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ح",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("خ'"),
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "د",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ذ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ر",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ز",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "س",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ش",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ص",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ض",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ط",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ظ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ع",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "غ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ف",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ق",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ك",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ل",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "م",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ن",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ه",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "و",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ي",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "آ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "إ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "أ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ؤ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ئ",
			prefix:   "",
			suffix:   "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "א",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ב",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ג",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ד",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ה",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ו",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ז",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ח",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ט",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "י",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "כ",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ל",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "מ",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "נ",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ס",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ע",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "פ",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "צ",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ק",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ר",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ש",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "ת",
			prefix:   "",
			suffix:   "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			contains: "a",
			prefix:   "",
			suffix:   "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "o",
			prefix:   "",
			suffix:   "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "e",
			prefix:   "",
			suffix:   "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "i",
			prefix:   "",
			suffix:   "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "y",
			prefix:   "",
			suffix:   "",
		},
		langs:  75030,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "u",
			prefix:   "",
			suffix:   "",
		},
		langs:  1286,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "j",
			prefix:   "",
			suffix:   "",
		},
		langs:  4096,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("j[^aoeiuy]"),
		},
		langs:  295488,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "g",
			prefix:   "",
			suffix:   "",
		},
		langs:  8,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "k",
			prefix:   "",
			suffix:   "",
		},
		langs:  364608,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "q",
			prefix:   "",
			suffix:   "",
		},
		langs:  748056,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "v",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "w",
			prefix:   "",
			suffix:   "",
		},
		langs:  993864,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "x",
			prefix:   "",
			suffix:   "",
		},
		langs:  534552,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "dj",
			prefix:   "",
			suffix:   "",
		},
		langs:  786432,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("v[^aoeiu]"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("y[^aoeiu]"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("c[^aohk]"),
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "dzi",
			prefix:   "",
			suffix:   "",
		},
		langs:  524512,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "ou",
			prefix:   "",
			suffix:   "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("a[eiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("ö[eaiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("ü[eaiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("e[aiou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("i[aeou]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("o[aieu]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("u[aieo]"),
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "aj",
			prefix:   "",
			suffix:   "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "ej",
			prefix:   "",
			suffix:   "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "oj",
			prefix:   "",
			suffix:   "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "uj",
			prefix:   "",
			suffix:   "",
		},
		langs:  240,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "eu",
			prefix:   "",
			suffix:   "",
		},
		langs:  147456,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "ky",
			prefix:   "",
			suffix:   "",
		},
		langs:  16384,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "kie",
			prefix:   "",
			suffix:   "",
		},
		langs:  262720,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "gie",
			prefix:   "",
			suffix:   "",
		},
		langs:  360960,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("ch[aou]"),
		},
		langs:  4096,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "ch",
			prefix:   "",
			suffix:   "",
		},
		langs:  524288,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "son",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "",
			suffix:   "",
			pattern:  regexp.MustCompile("sc[ei]"),
		},
		langs:  64,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "sch",
			prefix:   "",
			suffix:   "",
		},
		langs:  280640,
		accept: false,
	},
	{
		match: ruleMatcher{
			contains: "",
			prefix:   "h",
			suffix:   "",
		},
		langs:  131072,
		accept: false,
	},
}

var genFinalRules = finalRules{
	approx: finalRule{
		first: []rule{
			{
				pattern: "h",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[fktSs]"),
				},
				phonetic: "p",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "p",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "p",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vgdZz]"),
				},
				phonetic: "b",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "b",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pktSs]"),
				},
				phonetic: "f",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "f",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "f",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vbgdZz]"),
				},
				phonetic: "v",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "v",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pftSs]"),
				},
				phonetic: "k",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "k",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "k",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vbdZz]"),
				},
				phonetic: "g",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "g",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pfkSs]"),
				},
				phonetic: "t",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "t",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "t",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vbgZz]"),
				},
				phonetic: "d",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "d",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "dZ",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "tS",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pfkSt]"),
				},
				phonetic: "s",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
			},
			{
				pattern: "ji",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "i",
			},
			{
				pattern: "jI",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "I",
			},
			{
				pattern: "a",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[aA]"),
				},
				phonetic: "",
			},
			{
				pattern: "a",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "A",
				},
				phonetic: "",
			},
			{
				pattern: "A",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "A",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "b",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "d",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "f",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "g",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "j",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "j",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "k",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "l",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "l",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "m",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "m",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "n",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "n",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "p",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "r",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "r",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "t",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "v",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "z",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "vanden",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(vanden|)",
			},
			{
				pattern: "vander",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(vander|)",
			},
			{
				pattern: "van",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[bp]"),
				},
				phonetic: "(vam|[16])",
			},
			{
				pattern: "van",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(van|[16])",
			},
			{
				pattern: "n",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[bp]"),
				},
				phonetic: "m",
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
				pattern: "sen",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[rmnl]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(zn|zon)",
			},
			{
				pattern: "sen",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(sn|son)",
			},
			{
				pattern: "sEn",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[rmnl]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(zn|zon)",
			},
			{
				pattern: "sEn",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(sn|son)",
			},
			{
				pattern: "e",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "i",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "E",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "I",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "Q",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "Y",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[BbdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "e",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(e|)",
			},
			{
				pattern: "i",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(i|)",
			},
			{
				pattern: "E",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(E|)",
			},
			{
				pattern: "I",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(I|)",
			},
			{
				pattern: "Q",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(Q|)",
			},
			{
				pattern: "Y",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(Y|)",
			},
			{
				pattern:  "lEs",
				phonetic: "(lEs|lz)",
			},
			{
				pattern: "lE",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				phonetic: "(lE|l)",
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
				pattern: "lYndEr",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lander",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lAndEr",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lAnder",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "landEr",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lender",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lEndEr",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lendEr",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lEnder",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "lYnder",
			},
			{
				pattern: "burk",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "bUrk",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "burg",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "bUrg",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "Burk",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "BUrk",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "Burg",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "BUrg",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[rmnl]"),
				},
				phonetic: "z",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[rmnl]"),
				},
				phonetic: "z",
			},
			{
				pattern: "s",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[rmnl]$"),
				},
				phonetic: "z",
			},
			{
				pattern: "S",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[rmnl]$"),
				},
				phonetic: "z",
			},
			{
				pattern: "dS",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "S",
			},
			{
				pattern: "dZ",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "S",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "S",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(S|s)",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "(S|s)",
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
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fktSs]"),
					},
					phonetic: "(p|f[262144])",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "p",
						suffix:   "",
					},
					phonetic: "",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(p|f[262144])",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[pktSs]"),
					},
					phonetic: "(f|p[262144])",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "f",
						suffix:   "",
					},
					phonetic: "",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(f|p[262144])",
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
					pattern: "t",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(t|[64])",
				},
				{
					pattern: "g",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "n",
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(g|[64])",
				},
				{
					pattern: "k",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "n",
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(k|[64])",
				},
				{
					pattern: "p",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(p|[64])",
				},
				{
					pattern: "r",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[Ee]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(r|[64])",
				},
				{
					pattern: "s",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(s|[64])",
				},
				{
					pattern: "t",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiouAEIOU]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					phonetic: "(t|[64])",
				},
				{
					pattern: "s",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiouAEIOU]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^aeiouAEIOU]"),
					},
					phonetic: "(s|[64])",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phonetic: "(Q[128]|i|D[32])",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk[128])",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts[128])",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(Q[128]|i)",
				},
				{
					pattern: "lEE",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(li|il[32])",
				},
				{
					pattern: "rEE",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(ri|ir[32])",
				},
				{
					pattern: "lE",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(li|il[32]|lY[128])",
				},
				{
					pattern: "rE",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(ri|ir[32]|rY[128])",
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
					pattern: "iA",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ia|io)",
				},
				{
					pattern:  "iA",
					phonetic: "(ia|io|iY[128])",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phonetic: "(a|o|Y[128]|D[32])",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					phonetic: "(i|Y[128]|[32])",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					phonetic: "(i|Y[128]|[32])",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoAOQY]"),
					},
					phonetic: "i",
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
					pattern: "O",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprstv]$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[oeiuQY]$"),
					},
					phonetic: "o",
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
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[oeiuQY]$"),
					},
					phonetic: "(a|o)",
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
					pattern: "U",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DoiuQY]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "Uk",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(uk|Qk[128])",
				},
				{
					pattern: "Uk",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "uk",
				},
				{
					pattern: "sUts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(suts|sQts[128])",
				},
				{
					pattern: "Uts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "uts",
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
					pattern: "e",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprstv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoAOQY]"),
					},
					phonetic: "i",
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
					pattern: "p",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "p",
				},
				{
					pattern:  "p",
					phonetic: "(p|b)",
				},
			},
			uint64(genrussian): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiEIou]$"),
					},
					phonetic: "i",
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
					pattern: "om",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(om|im)",
				},
				{
					pattern: "on",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(on|in)",
				},
				{
					pattern: "em",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|om)",
				},
				{
					pattern: "en",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|on)",
				},
				{
					pattern: "Em",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|Ym|om)",
				},
				{
					pattern: "En",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|Yn|on)",
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
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			uint64(gencyrillic): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiEIou]$"),
					},
					phonetic: "i",
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
					pattern: "om",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(om|im)",
				},
				{
					pattern: "on",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(on|in)",
				},
				{
					pattern: "em",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|om)",
				},
				{
					pattern: "en",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|on)",
				},
				{
					pattern: "Em",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|Ym|om)",
				},
				{
					pattern: "En",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|Yn|on)",
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
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
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
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^aEIeiou]e"),
					},
					phonetic: "(Q|i|D)",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aEIeiou]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern: "lE",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(il|li|lY)",
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
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("D[^aeiEIou]$"),
					},
					phonetic: "(i|)",
				},
				{
					pattern: "e",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("D[^aeiEIou]$"),
					},
					phonetic: "(i|)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiEuQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoQY]"),
					},
					phonetic: "i",
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
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiAEIOUouQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
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
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoAOUiuQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoAOQY]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aoAOUeiuQY]$"),
					},
					phonetic: "o",
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
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aoeOUiuQY]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern:  "A",
					phonetic: "(a|o|Y)",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuUQY]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "Uk",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(uk|Qk)",
				},
				{
					pattern: "Uk",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "uk",
				},
				{
					pattern: "sUts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(suts|sQts)",
				},
				{
					pattern: "Uts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "uts",
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
					pattern: "aiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(o|om|im)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(o|on|in)",
				},
				{
					pattern:  "B",
					phonetic: "o",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(i|im|om)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(i|in|on)",
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
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phonetic: "i",
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
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
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
					pattern: "aiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(o|om|im)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(o|on|in)",
				},
				{
					pattern:  "B",
					phonetic: "o",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[bp]"),
					},
					phonetic: "(i|im|om)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(i|in|on)",
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
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phonetic: "i",
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
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
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
				pattern: "h",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[fktSs]"),
				},
				phonetic: "p",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "p",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "p",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vgdZz]"),
				},
				phonetic: "b",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "b",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pktSs]"),
				},
				phonetic: "f",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "f",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "f",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vbgdZz]"),
				},
				phonetic: "v",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "v",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pftSs]"),
				},
				phonetic: "k",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "k",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "k",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vbdZz]"),
				},
				phonetic: "g",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "g",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pfkSs]"),
				},
				phonetic: "t",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "t",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "t",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[vbgZz]"),
				},
				phonetic: "d",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "d",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "dZ",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "tS",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pfkSt]"),
				},
				phonetic: "s",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
			},
			{
				pattern: "ji",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "i",
			},
			{
				pattern: "jI",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "I",
			},
			{
				pattern: "a",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[aA]"),
				},
				phonetic: "",
			},
			{
				pattern: "a",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "A",
				},
				phonetic: "",
			},
			{
				pattern: "A",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "A",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "b",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "d",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "f",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "g",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "j",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "j",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "k",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "l",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "l",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "m",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "m",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "n",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "n",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "p",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "r",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "r",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "t",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "v",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "z",
					suffix:   "",
				},
				phonetic: "",
			},
			{
				pattern:  "H",
				phonetic: "",
			},
			{
				pattern: "s",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[^t]$"),
				},
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[bgZd]"),
				},
				phonetic: "z",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[pfkst]"),
				},
				phonetic: "S",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "S",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^[bgzd]"),
				},
				phonetic: "Z",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("^$"),
				},
				phonetic: "s",
			},
			{
				pattern: "ji",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
			{
				pattern: "jI",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
			{
				pattern: "je",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
			{
				pattern: "jE",
				leftContext: &ruleMatcher{
					contains: "",
					prefix:   "",
					suffix:   "",
					pattern:  regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
		},
		second: map[uint64][]rule{
			uint64(genany): []rule{
				{
					pattern: "EE",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "e",
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
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fktSs]"),
					},
					phonetic: "p",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "p",
						suffix:   "",
					},
					phonetic: "",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "p",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[pktSs]"),
					},
					phonetic: "f",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "f",
						suffix:   "",
					},
					phonetic: "",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "f",
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
					pattern: "EE",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "e",
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
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[fktSs]"),
					},
					phonetic: "p",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "p",
						suffix:   "",
					},
					phonetic: "",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "p",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^[pktSs]"),
					},
					phonetic: "f",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "f",
						suffix:   "",
					},
					phonetic: "",
				},
				{
					pattern: "V",
					rightContext: &ruleMatcher{
						contains: "",
						prefix:   "",
						suffix:   "",
						pattern:  regexp.MustCompile("^$"),
					},
					phonetic: "f",
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
