// GENERATED CODE. DO NOT EDIT!
package bmpm

type ashLang int

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

var ashRules = map[ashLang][]rule{
	ashany: []rule{
		{
			patterns: []string{
				"yna",
				"",
				"$",
				"(in[512]|ina)",
			},
		},
		{
			patterns: []string{
				"ina",
				"",
				"$",
				"(in[512]|ina)",
			},
		},
		{
			patterns: []string{
				"liova",
				"",
				"$",
				"(lof[512]|lef[512]|lova)",
			},
		},
		{
			patterns: []string{
				"lova",
				"",
				"$",
				"(lof[512]|lef[512]|lova)",
			},
		},
		{
			patterns: []string{
				"ova",
				"",
				"$",
				"(of[512]|ova)",
			},
		},
		{
			patterns: []string{
				"eva",
				"",
				"$",
				"(ef[512]|eva)",
			},
		},
		{
			patterns: []string{
				"aia",
				"",
				"$",
				"(aja|i[512])",
			},
		},
		{
			patterns: []string{
				"aja",
				"",
				"$",
				"(aja|i[512])",
			},
		},
		{
			patterns: []string{
				"aya",
				"",
				"$",
				"(aja|i[512])",
			},
		},
		{
			patterns: []string{
				"lowa",
				"",
				"$",
				"(lova|lof[128]|l[128]|el[128])",
			},
		},
		{
			patterns: []string{
				"kowa",
				"",
				"$",
				"(kova|kof[128]|k[128]|ek[128])",
			},
		},
		{
			patterns: []string{
				"owa",
				"",
				"$",
				"(ova|of[128]|)",
			},
		},
		{
			patterns: []string{
				"lowna",
				"",
				"$",
				"(lovna|levna|l[128]|el[128])",
			},
		},
		{
			patterns: []string{
				"kowna",
				"",
				"$",
				"(kovna|k[128]|ek[128])",
			},
		},
		{
			patterns: []string{
				"owna",
				"",
				"$",
				"(ovna|[128])",
			},
		},
		{
			patterns: []string{
				"lówna",
				"",
				"$",
				"(l|el[128])",
			},
		},
		{
			patterns: []string{
				"kówna",
				"",
				"$",
				"(k|ek[128])",
			},
		},
		{
			patterns: []string{
				"ówna",
				"",
				"$",
				"",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"$",
				"(a|i[128])",
			},
		},
		{
			patterns: []string{
				"rh",
				"^",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"ssch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"chsch",
				"",
				"",
				"xS",
			},
		},
		{
			patterns: []string{
				"tsch",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"[ei]",
				"(sk[256]|S|StS[512])",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"",
				"(S|StS[512])",
			},
		},
		{
			patterns: []string{
				"ssh",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"sh",
				"",
				"[äöü]",
				"sh",
			},
		},
		{
			patterns: []string{
				"sh",
				"",
				"[aeiou]",
				"(S[516]|sh)",
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
				"(x[516]|kh)",
			},
		},
		{
			patterns: []string{
				"chs",
				"",
				"",
				"(ks[16]|xs|tSs[516])",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"[ei]",
				"(x|k[256]|tS[516])",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"(x|tS[516])",
			},
		},
		{
			patterns: []string{
				"ck",
				"",
				"",
				"(k|tsk[128])",
			},
		},
		{
			patterns: []string{
				"czy",
				"",
				"",
				"tSi",
			},
		},
		{
			patterns: []string{
				"cze",
				"",
				"[bcdgkpstwzż]",
				"(tSe|tSF)",
			},
		},
		{
			patterns: []string{
				"ciewicz",
				"",
				"",
				"(tsevitS|tSevitS)",
			},
		},
		{
			patterns: []string{
				"siewicz",
				"",
				"",
				"(sevitS|SevitS)",
			},
		},
		{
			patterns: []string{
				"ziewicz",
				"",
				"",
				"(zevitS|ZevitS)",
			},
		},
		{
			patterns: []string{
				"riewicz",
				"",
				"",
				"rjevitS",
			},
		},
		{
			patterns: []string{
				"diewicz",
				"",
				"",
				"djevitS",
			},
		},
		{
			patterns: []string{
				"tiewicz",
				"",
				"",
				"tjevitS",
			},
		},
		{
			patterns: []string{
				"iewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: []string{
				"ewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: []string{
				"owicz",
				"",
				"",
				"ovitS",
			},
		},
		{
			patterns: []string{
				"icz",
				"",
				"",
				"itS",
			},
		},
		{
			patterns: []string{
				"cz",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"cia",
				"",
				"[bcdgkpstwzż]",
				"(tSB[128]|tsB)",
			},
		},
		{
			patterns: []string{
				"cia",
				"",
				"",
				"(tSa[128]|tsa)",
			},
		},
		{
			patterns: []string{
				"cią",
				"",
				"[bp]",
				"(tSom[128]|tsom)",
			},
		},
		{
			patterns: []string{
				"cią",
				"",
				"",
				"(tSon[128]|tson)",
			},
		},
		{
			patterns: []string{
				"cię",
				"",
				"[bp]",
				"(tSem[128]|tsem)",
			},
		},
		{
			patterns: []string{
				"cię",
				"",
				"",
				"(tSen[128]|tsen)",
			},
		},
		{
			patterns: []string{
				"cie",
				"",
				"[bcdgkpstwzż]",
				"(tSF[128]|tsF)",
			},
		},
		{
			patterns: []string{
				"cie",
				"",
				"",
				"(tSe[128]|tse)",
			},
		},
		{
			patterns: []string{
				"cio",
				"",
				"",
				"(tSo[128]|tso)",
			},
		},
		{
			patterns: []string{
				"ciu",
				"",
				"",
				"(tSu[128]|tsu)",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"$",
				"(tsi[128]|tSi[384]|tS[256]|si)",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"",
				"(tsi[128]|tSi[384]|si)",
			},
		},
		{
			patterns: []string{
				"ce",
				"",
				"[bcdgkpstwzż]",
				"(tsF[128]|tSe[384]|se)",
			},
		},
		{
			patterns: []string{
				"ce",
				"",
				"",
				"(tSe[384]|tse[128]|se)",
			},
		},
		{
			patterns: []string{
				"cy",
				"",
				"",
				"(si|tsi[128])",
			},
		},
		{
			patterns: []string{
				"ssz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"sz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"ssp",
				"",
				"",
				"(Sp[16]|sp)",
			},
		},
		{
			patterns: []string{
				"sp",
				"",
				"",
				"(Sp[16]|sp)",
			},
		},
		{
			patterns: []string{
				"sst",
				"",
				"",
				"(St[16]|st)",
			},
		},
		{
			patterns: []string{
				"st",
				"",
				"",
				"(St[16]|st)",
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
				"sia",
				"",
				"[bcdgkpstwzż]",
				"(SB[128]|sB[128]|sja)",
			},
		},
		{
			patterns: []string{
				"sia",
				"",
				"",
				"(Sa[128]|sja)",
			},
		},
		{
			patterns: []string{
				"sią",
				"",
				"[bp]",
				"(Som[128]|som)",
			},
		},
		{
			patterns: []string{
				"sią",
				"",
				"",
				"(Son[128]|son)",
			},
		},
		{
			patterns: []string{
				"się",
				"",
				"[bp]",
				"(Sem[128]|sem)",
			},
		},
		{
			patterns: []string{
				"się",
				"",
				"",
				"(Sen[128]|sen)",
			},
		},
		{
			patterns: []string{
				"sie",
				"",
				"[bcdgkpstwzż]",
				"(SF[128]|sF|zi[16])",
			},
		},
		{
			patterns: []string{
				"sie",
				"",
				"",
				"(se|Se[128]|zi[16])",
			},
		},
		{
			patterns: []string{
				"sio",
				"",
				"",
				"(So[128]|so)",
			},
		},
		{
			patterns: []string{
				"siu",
				"",
				"",
				"(Su[128]|sju)",
			},
		},
		{
			patterns: []string{
				"si",
				"",
				"",
				"(Si[128]|si|zi[16])",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"[aeiouäöë]",
				"(s|z[16])",
			},
		},
		{
			patterns: []string{
				"gue",
				"",
				"",
				"ge",
			},
		},
		{
			patterns: []string{
				"gui",
				"",
				"",
				"gi",
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
				"gh",
				"",
				"[ei]",
				"(g[256]|gh)",
			},
		},
		{
			patterns: []string{
				"gauz",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: []string{
				"gaus",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: []string{
				"gol'ts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"golts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"gol'tz",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"goltz",
				"",
				"",
				"holts",
			},
		},
		{
			patterns: []string{
				"gol'ts",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: []string{
				"golts",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: []string{
				"gol'tz",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: []string{
				"goltz",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: []string{
				"gendler",
				"",
				"$",
				"hendler",
			},
		},
		{
			patterns: []string{
				"gejmer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"gejm",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"geymer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"geym",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"geimer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"geim",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"gof",
				"",
				"$",
				"hof",
			},
		},
		{
			patterns: []string{
				"ger",
				"",
				"$",
				"ger",
			},
		},
		{
			patterns: []string{
				"gen",
				"",
				"$",
				"gen",
			},
		},
		{
			patterns: []string{
				"gin",
				"",
				"$",
				"gin",
			},
		},
		{
			patterns: []string{
				"gie",
				"",
				"$",
				"(ge|gi[16]|ji[8])",
			},
		},
		{
			patterns: []string{
				"gie",
				"",
				"",
				"ge",
			},
		},
		{
			patterns: []string{
				"ge",
				"[yaeiou]",
				"",
				"(gE|xe[1024]|dZe[260])",
			},
		},
		{
			patterns: []string{
				"gi",
				"[yaeiou]",
				"",
				"(gI|xi[1024]|dZi[260])",
			},
		},
		{
			patterns: []string{
				"ge",
				"",
				"",
				"(gE|dZe[260]|hE[512]|xe[1024])",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"",
				"(gI|dZi[260]|hI[512]|xi[1024])",
			},
		},
		{
			patterns: []string{
				"gy",
				"",
				"[aeouáéóúüöőű]",
				"(gi|dj[64])",
			},
		},
		{
			patterns: []string{
				"gy",
				"",
				"",
				"(gi|d[64])",
			},
		},
		{
			patterns: []string{
				"g",
				"[jyaeiou]",
				"[aouyei]",
				"g",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[aouei]",
				"(g|h[512])",
			},
		},
		{
			patterns: []string{
				"ej",
				"",
				"",
				"(aj|eZ[264]|ex[1024])",
			},
		},
		{
			patterns: []string{
				"ej",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ly",
				"",
				"[au]",
				"(l|lj)",
			},
		},
		{
			patterns: []string{
				"li",
				"",
				"[au]",
				"(l|lj)",
			},
		},
		{
			patterns: []string{
				"lj",
				"",
				"[au]",
				"(l|lj)",
			},
		},
		{
			patterns: []string{
				"lio",
				"",
				"",
				"(lo|le[512])",
			},
		},
		{
			patterns: []string{
				"lyo",
				"",
				"",
				"(lo|le[512])",
			},
		},
		{
			patterns: []string{
				"ll",
				"",
				"",
				"(l|J[1024])",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"[aoeiuy]",
				"(j|dZ[4]|x[1024]|Z[264])",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"(j|x[1024])",
			},
		},
		{
			patterns: []string{
				"pf",
				"",
				"",
				"(pf|p|f)",
			},
		},
		{
			patterns: []string{
				"ph",
				"",
				"",
				"(ph|f)",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"(kv[16]|k)",
			},
		},
		{
			patterns: []string{
				"rze",
				"t",
				"",
				"(Se[128]|re)",
			},
		},
		{
			patterns: []string{
				"rze",
				"",
				"",
				"(rze|rtsE[16]|Ze[128]|re[128]|rZe[128])",
			},
		},
		{
			patterns: []string{
				"rzy",
				"t",
				"",
				"(Si[128]|ri)",
			},
		},
		{
			patterns: []string{
				"rzy",
				"",
				"",
				"(Zi[128]|ri[128]|rZi)",
			},
		},
		{
			patterns: []string{
				"rz",
				"t",
				"",
				"(S[128]|r)",
			},
		},
		{
			patterns: []string{
				"rz",
				"",
				"",
				"(rz|rts[16]|Z[128]|r[128]|rZ[128])",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"$",
				"(ts|tS[20])",
			},
		},
		{
			patterns: []string{
				"tz",
				"^",
				"",
				"(ts|tS[20])",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"",
				"(ts[532]|tz)",
			},
		},
		{
			patterns: []string{
				"zh",
				"",
				"",
				"(Z|zh[128]|tsh[16])",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"[bcdgkpstwzż]",
				"(ZB[128]|zB[128]|zja)",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"",
				"(Za[128]|zja)",
			},
		},
		{
			patterns: []string{
				"zią",
				"",
				"[bp]",
				"(Zom[128]|zom)",
			},
		},
		{
			patterns: []string{
				"zią",
				"",
				"",
				"(Zon[128]|zon)",
			},
		},
		{
			patterns: []string{
				"zię",
				"",
				"[bp]",
				"(Zem[128]|zem)",
			},
		},
		{
			patterns: []string{
				"zię",
				"",
				"",
				"(Zen[128]|zen)",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"[bcdgkpstwzż]",
				"(ZF[128]|zF[128]|ze|tsi[16])",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"",
				"(ze|Ze[128]|tsi[16])",
			},
		},
		{
			patterns: []string{
				"zio",
				"",
				"",
				"(Zo[128]|zo)",
			},
		},
		{
			patterns: []string{
				"ziu",
				"",
				"",
				"(Zu[128]|zju)",
			},
		},
		{
			patterns: []string{
				"zi",
				"",
				"",
				"(Zi[128]|zi|tsi[16])",
			},
		},
		{
			patterns: []string{
				"thal",
				"",
				"$",
				"tal",
			},
		},
		{
			patterns: []string{
				"th",
				"^",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"th",
				"",
				"[aeiou]",
				"(t[16]|th)",
			},
		},
		{
			patterns: []string{
				"th",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"vogel",
				"",
				"",
				"(vogel|fogel[16])",
			},
		},
		{
			patterns: []string{
				"v",
				"^",
				"",
				"(v|f[16])",
			},
		},
		{
			patterns: []string{
				"h",
				"[aeiouyäöü]",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"(h|x[384])",
			},
		},
		{
			patterns: []string{
				"h",
				"^",
				"",
				"(h|H[20])",
			},
		},
		{
			patterns: []string{
				"yi",
				" ",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"iy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"e",
				"in",
				"$",
				"(e|[8])",
			},
		},
		{
			patterns: []string{
				"yj",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"$",
				"i",
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
				"oue",
				"",
				"",
				"oue",
			},
		},
		{
			patterns: []string{
				"au",
				"",
				"",
				"(au|o[8])",
			},
		},
		{
			patterns: []string{
				"ou",
				"",
				"",
				"(ou|u[8])",
			},
		},
		{
			patterns: []string{
				"ue",
				"",
				"",
				"(Q|uje[512])",
			},
		},
		{
			patterns: []string{
				"ae",
				"",
				"",
				"(Y[16]|aje[512]|ae)",
			},
		},
		{
			patterns: []string{
				"oe",
				"",
				"",
				"(Y[16]|oje[512]|oe)",
			},
		},
		{
			patterns: []string{
				"ee",
				"",
				"",
				"(i[4]|aje[512]|e)",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"(aj[16]|oj[16]|eu)",
			},
		},
		{
			patterns: []string{
				"i",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"[bcdgkpstwzż]",
				"(i[16]|e[128]|ije[512]|je)",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"(i[16]|e[128]|ije[512]|je)",
			},
		},
		{
			patterns: []string{
				"ye",
				"",
				"",
				"(je|ije[512])",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: []string{
				"io",
				"",
				"",
				"(jo|e[512])",
			},
		},
		{
			patterns: []string{
				"yo",
				"",
				"",
				"(jo|e[512])",
			},
		},
		{
			patterns: []string{
				"ea",
				"",
				"",
				"(ea|ja[256])",
			},
		},
		{
			patterns: []string{
				"e",
				"^",
				"",
				"(e|je[512])",
			},
		},
		{
			patterns: []string{
				"oo",
				"",
				"",
				"(u[4]|o)",
			},
		},
		{
			patterns: []string{
				"uu",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ć",
				"",
				"",
				"(tS[128]|ts)",
			},
		},
		{
			patterns: []string{
				"ł",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"ń",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"ñ",
				"",
				"",
				"(n|nj[1024])",
			},
		},
		{
			patterns: []string{
				"ś",
				"",
				"",
				"(S[128]|s)",
			},
		},
		{
			patterns: []string{
				"ş",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"ţ",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: []string{
				"ż",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"ź",
				"",
				"",
				"(Z[128]|z)",
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
				"ą",
				"",
				"[bp]",
				"om",
			},
		},
		{
			patterns: []string{
				"ą",
				"",
				"",
				"on",
			},
		},
		{
			patterns: []string{
				"ä",
				"",
				"",
				"(Y|e)",
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
				"ă",
				"",
				"",
				"(e[256]|a)",
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
				"é",
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
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ę",
				"",
				"[bp]",
				"em",
			},
		},
		{
			patterns: []string{
				"ę",
				"",
				"",
				"en",
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
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"ő",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"ó",
				"",
				"",
				"(u[128]|o)",
			},
		},
		{
			patterns: []string{
				"ű",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: []string{
				"ü",
				"",
				"",
				"Q",
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
				"ű",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: []string{
				"ß",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"'",
				"",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"\"",
				"",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"[bcdgkpstwzż]",
				"(A|B[128])",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"[bcdgkpstwzż]",
				"(E|F[128])",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"[bcćdgklłmnńrsśtwzźż]",
				"(O|P[128])",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"A",
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
				"(k|ts[128])",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"j",
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
				"O",
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
				"U",
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
				"(ts[16]|z)",
			},
		},
	},
	ashcyrillic: []rule{
		{
			patterns: []string{
				"ця",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: []string{
				"цю",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: []string{
				"циа",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: []string{
				"цие",
				"",
				"",
				"tse",
			},
		},
		{
			patterns: []string{
				"цио",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: []string{
				"циу",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: []string{
				"сие",
				"",
				"",
				"se",
			},
		},
		{
			patterns: []string{
				"сио",
				"",
				"",
				"so",
			},
		},
		{
			patterns: []string{
				"зие",
				"",
				"",
				"ze",
			},
		},
		{
			patterns: []string{
				"зио",
				"",
				"",
				"zo",
			},
		},
		{
			patterns: []string{
				"гауз",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: []string{
				"гаус",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: []string{
				"гольц",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"геймер",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"гейм",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"гоф",
				"",
				"$",
				"hof",
			},
		},
		{
			patterns: []string{
				"гер",
				"",
				"$",
				"ger",
			},
		},
		{
			patterns: []string{
				"ген",
				"",
				"$",
				"gen",
			},
		},
		{
			patterns: []string{
				"гин",
				"",
				"$",
				"gin",
			},
		},
		{
			patterns: []string{
				"г",
				"(й|ё|я|ю|ы|а|е|о|и|у)",
				"(а|е|о|и|у)",
				"g",
			},
		},
		{
			patterns: []string{
				"г",
				"",
				"(а|е|о|и|у)",
				"(g|h)",
			},
		},
		{
			patterns: []string{
				"ля",
				"",
				"",
				"la",
			},
		},
		{
			patterns: []string{
				"лю",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: []string{
				"лё",
				"",
				"",
				"(le|lo)",
			},
		},
		{
			patterns: []string{
				"лио",
				"",
				"",
				"(le|lo)",
			},
		},
		{
			patterns: []string{
				"ле",
				"",
				"",
				"(lE|lo)",
			},
		},
		{
			patterns: []string{
				"ийе",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ие",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ыйе",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ые",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ий",
				"",
				"(а|о|у)",
				"j",
			},
		},
		{
			patterns: []string{
				"ый",
				"",
				"(а|о|у)",
				"j",
			},
		},
		{
			patterns: []string{
				"ий",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"ый",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"ё",
				"",
				"",
				"(e|jo)",
			},
		},
		{
			patterns: []string{
				"ей",
				"^",
				"",
				"(jaj|aj)",
			},
		},
		{
			patterns: []string{
				"е",
				"(а|е|о|у)",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"е",
				"^",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"эй",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ей",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ауе",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: []string{
				"ауэ",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: []string{
				"а",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"б",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"в",
				"",
				"",
				"v",
			},
		},
		{
			patterns: []string{
				"г",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"д",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"е",
				"",
				"",
				"E",
			},
		},
		{
			patterns: []string{
				"ж",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"з",
				"",
				"",
				"z",
			},
		},
		{
			patterns: []string{
				"и",
				"",
				"",
				"I",
			},
		},
		{
			patterns: []string{
				"й",
				"",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"к",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"л",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"м",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"н",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"о",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"п",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"р",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"с",
				"",
				"с",
				"",
			},
		},
		{
			patterns: []string{
				"с",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"т",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"у",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ф",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"х",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"ц",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: []string{
				"ч",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ш",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"щ",
				"",
				"",
				"StS",
			},
		},
		{
			patterns: []string{
				"ъ",
				"",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"ы",
				"",
				"",
				"I",
			},
		},
		{
			patterns: []string{
				"ь",
				"",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"э",
				"",
				"",
				"E",
			},
		},
		{
			patterns: []string{
				"ю",
				"",
				"",
				"ju",
			},
		},
		{
			patterns: []string{
				"я",
				"",
				"",
				"ja",
			},
		},
	},
	ashenglish: []rule{
		{
			patterns: []string{
				"tch",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"(tS|x)",
			},
		},
		{
			patterns: []string{
				"ck",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"cc",
				"",
				"[iey]",
				"ks",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"c",
				"",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[iey]",
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
				"gh",
				"^",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"gh",
				"",
				"",
				"(g|f|w)",
			},
		},
		{
			patterns: []string{
				"gn",
				"",
				"",
				"(gn|n)",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[iey]",
				"(g|dZ)",
			},
		},
		{
			patterns: []string{
				"th",
				"",
				"",
				"t",
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
				"ph",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"",
				"(S|sk)",
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
				"who",
				"^",
				"",
				"hu",
			},
		},
		{
			patterns: []string{
				"wh",
				"^",
				"",
				"w",
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
				"h",
				"",
				"[^aeiou]",
				"",
			},
		},
		{
			patterns: []string{
				"h",
				"^",
				"",
				"H",
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
				"j",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: []string{
				"kn",
				"^",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"mb",
				"",
				"$",
				"m",
			},
		},
		{
			patterns: []string{
				"ng",
				"",
				"$",
				"(N|ng)",
			},
		},
		{
			patterns: []string{
				"pn",
				"^",
				"",
				"(pn|n)",
			},
		},
		{
			patterns: []string{
				"ps",
				"^",
				"",
				"(ps|s)",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"kw",
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
				"tia",
				"",
				"",
				"(So|Sa)",
			},
		},
		{
			patterns: []string{
				"tio",
				"",
				"",
				"So",
			},
		},
		{
			patterns: []string{
				"wr",
				"^",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"(w|v)",
			},
		},
		{
			patterns: []string{
				"x",
				"^",
				"",
				"z",
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
				"yi",
				" ",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"y",
				"^",
				"[aeiouy]",
				"j",
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
				"oue",
				"",
				"",
				"(aue|oue)",
			},
		},
		{
			patterns: []string{
				"ai",
				"",
				"",
				"(aj|e)",
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
				"a",
				"",
				"[^aeiou]e",
				"aj",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"",
				"(e|o|a)",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"(aj|i)",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"(aj|i)",
			},
		},
		{
			patterns: []string{
				"ear",
				"",
				"",
				"ia",
			},
		},
		{
			patterns: []string{
				"ea",
				"",
				"",
				"(i|e)",
			},
		},
		{
			patterns: []string{
				"ee",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"[^aeiou]e",
				"i",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"$",
				"(|E)",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"",
				"E",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[^aeiou]e",
				"aj",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"I",
			},
		},
		{
			patterns: []string{
				"oa",
				"",
				"",
				"ou",
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
				"oo",
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
				"(u|ou)",
			},
		},
		{
			patterns: []string{
				"oy",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"[^aeiou]e",
				"ou",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"",
				"(o|a)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"[^aeiou]e",
				"(ju|u)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"r",
				"(e|u)",
			},
		},
		{
			patterns: []string{
				"u",
				"",
				"",
				"(u|a)",
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
	ashfrench: []rule{
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
				"aj",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"aj",
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
				"yi",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				"",
				"i",
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
				"E",
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
				"I",
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
	ashgerman: []rule{
		{
			patterns: []string{
				"ziu",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: []string{
				"zio",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: []string{
				"ssch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"chsch",
				"",
				"",
				"xS",
			},
		},
		{
			patterns: []string{
				"ewitsch",
				"",
				"$",
				"evitS",
			},
		},
		{
			patterns: []string{
				"owitsch",
				"",
				"$",
				"ovitS",
			},
		},
		{
			patterns: []string{
				"evitsch",
				"",
				"$",
				"evitS",
			},
		},
		{
			patterns: []string{
				"ovitsch",
				"",
				"$",
				"ovitS",
			},
		},
		{
			patterns: []string{
				"witsch",
				"",
				"$",
				"vitS",
			},
		},
		{
			patterns: []string{
				"vitsch",
				"",
				"$",
				"vitS",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"chs",
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
				"x",
			},
		},
		{
			patterns: []string{
				"ck",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[eiy]",
				"ts",
			},
		},
		{
			patterns: []string{
				"sp",
				"^",
				"",
				"Sp",
			},
		},
		{
			patterns: []string{
				"st",
				"^",
				"",
				"St",
			},
		},
		{
			patterns: []string{
				"ssp",
				"",
				"",
				"(Sp|sp)",
			},
		},
		{
			patterns: []string{
				"sp",
				"",
				"",
				"(Sp|sp)",
			},
		},
		{
			patterns: []string{
				"sst",
				"",
				"",
				"(St|st)",
			},
		},
		{
			patterns: []string{
				"st",
				"",
				"",
				"(St|st)",
			},
		},
		{
			patterns: []string{
				"pf",
				"",
				"",
				"(pf|p|f)",
			},
		},
		{
			patterns: []string{
				"ph",
				"",
				"",
				"(ph|f)",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"kv",
			},
		},
		{
			patterns: []string{
				"ewitz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: []string{
				"ewiz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: []string{
				"evitz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: []string{
				"eviz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: []string{
				"owitz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: []string{
				"owiz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: []string{
				"ovitz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: []string{
				"oviz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: []string{
				"witz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: []string{
				"wiz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: []string{
				"vitz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: []string{
				"viz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: []string{
				"thal",
				"",
				"$",
				"tal",
			},
		},
		{
			patterns: []string{
				"th",
				"^",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"th",
				"",
				"[äöüaeiou]",
				"(t|th)",
			},
		},
		{
			patterns: []string{
				"th",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"rh",
				"^",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"h",
				"[aeiouyäöü]",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"h",
				"^",
				"",
				"H",
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
				"s",
				"",
				"[äöüaeiouy]",
				"(z|s)",
			},
		},
		{
			patterns: []string{
				"s",
				"[aeiouyäöüj]",
				"[aeiouyäöü]",
				"z",
			},
		},
		{
			patterns: []string{
				"ß",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"$",
				"i",
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
				"ue",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: []string{
				"ae",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"oe",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"ü",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: []string{
				"ä",
				"",
				"",
				"(Y|e)",
			},
		},
		{
			patterns: []string{
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"(aj|oj)",
			},
		},
		{
			patterns: []string{
				"i",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"I",
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
				"ñ",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"ã",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"ő",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ű",
				"",
				"",
				"u",
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
				"a",
				"",
				"",
				"A",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"j",
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
				"O",
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
				"U",
			},
		},
		{
			patterns: []string{
				"v",
				"",
				"",
				"(f|v)",
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
				"ts",
			},
		},
	},
	ashhebrew: []rule{
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
				"TB",
			},
		},
	},
	ashhungarian: []rule{
		{
			patterns: []string{
				"sz",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"zs",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"cs",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ay",
				"",
				"",
				"(oj|aj)",
			},
		},
		{
			patterns: []string{
				"ai",
				"",
				"",
				"(oj|aj)",
			},
		},
		{
			patterns: []string{
				"aj",
				"",
				"",
				"(oj|aj)",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"y",
				"[áo]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"i",
				"[áo]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"ee",
				"",
				"",
				"(aj|e)",
			},
		},
		{
			patterns: []string{
				"ely",
				"",
				"",
				"(aj|eli)",
			},
		},
		{
			patterns: []string{
				"ly",
				"",
				"",
				"(j|li)",
			},
		},
		{
			patterns: []string{
				"gy",
				"",
				"[aeouáéóúüöőű]",
				"dj",
			},
		},
		{
			patterns: []string{
				"gy",
				"",
				"",
				"(d|gi)",
			},
		},
		{
			patterns: []string{
				"ny",
				"",
				"[aeouáéóúüöőű]",
				"nj",
			},
		},
		{
			patterns: []string{
				"ny",
				"",
				"",
				"(n|ni)",
			},
		},
		{
			patterns: []string{
				"ty",
				"",
				"[aeouáéóúüöőű]",
				"tj",
			},
		},
		{
			patterns: []string{
				"ty",
				"",
				"",
				"(t|ti)",
			},
		},
		{
			patterns: []string{
				"qu",
				"",
				"",
				"(ku|kv)",
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
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"ő",
				"",
				"",
				"Y",
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
				"Q",
			},
		},
		{
			patterns: []string{
				"ű",
				"",
				"",
				"Q",
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
				"ts",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"j",
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
				"(S|s)",
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
				"z",
			},
		},
	},
	ashpolish: []rule{
		{
			patterns: []string{
				"ska",
				"",
				"$",
				"ski",
			},
		},
		{
			patterns: []string{
				"cka",
				"",
				"$",
				"tski",
			},
		},
		{
			patterns: []string{
				"lowa",
				"",
				"$",
				"(lova|lof|l|el)",
			},
		},
		{
			patterns: []string{
				"kowa",
				"",
				"$",
				"(kova|kof|k|ek)",
			},
		},
		{
			patterns: []string{
				"owa",
				"",
				"$",
				"(ova|of|)",
			},
		},
		{
			patterns: []string{
				"lowna",
				"",
				"$",
				"(lovna|levna|l|el)",
			},
		},
		{
			patterns: []string{
				"kowna",
				"",
				"$",
				"(kovna|k|ek)",
			},
		},
		{
			patterns: []string{
				"owna",
				"",
				"$",
				"(ovna|)",
			},
		},
		{
			patterns: []string{
				"lówna",
				"",
				"$",
				"(l|el)",
			},
		},
		{
			patterns: []string{
				"kówna",
				"",
				"$",
				"(k|ek)",
			},
		},
		{
			patterns: []string{
				"ówna",
				"",
				"$",
				"",
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
				"czy",
				"",
				"",
				"tSi",
			},
		},
		{
			patterns: []string{
				"cze",
				"",
				"[bcdgkpstwzż]",
				"(tSe|tSF)",
			},
		},
		{
			patterns: []string{
				"ciewicz",
				"",
				"",
				"(tsevitS|tSevitS)",
			},
		},
		{
			patterns: []string{
				"siewicz",
				"",
				"",
				"(sevitS|SevitS)",
			},
		},
		{
			patterns: []string{
				"ziewicz",
				"",
				"",
				"(zevitS|ZevitS)",
			},
		},
		{
			patterns: []string{
				"riewicz",
				"",
				"",
				"rjevitS",
			},
		},
		{
			patterns: []string{
				"diewicz",
				"",
				"",
				"djevitS",
			},
		},
		{
			patterns: []string{
				"tiewicz",
				"",
				"",
				"tjevitS",
			},
		},
		{
			patterns: []string{
				"iewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: []string{
				"ewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: []string{
				"owicz",
				"",
				"",
				"ovitS",
			},
		},
		{
			patterns: []string{
				"icz",
				"",
				"",
				"itS",
			},
		},
		{
			patterns: []string{
				"cz",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"cia",
				"",
				"[bcdgkpstwzż]",
				"(tSB|tsB)",
			},
		},
		{
			patterns: []string{
				"cia",
				"",
				"",
				"(tSa|tsa)",
			},
		},
		{
			patterns: []string{
				"cią",
				"",
				"[bp]",
				"(tSom|tsom)",
			},
		},
		{
			patterns: []string{
				"cią",
				"",
				"",
				"(tSon|tson)",
			},
		},
		{
			patterns: []string{
				"cię",
				"",
				"[bp]",
				"(tSem|tsem)",
			},
		},
		{
			patterns: []string{
				"cię",
				"",
				"",
				"(tSen|tsen)",
			},
		},
		{
			patterns: []string{
				"cie",
				"",
				"[bcdgkpstwzż]",
				"(tSF|tsF)",
			},
		},
		{
			patterns: []string{
				"cie",
				"",
				"",
				"(tSe|tse)",
			},
		},
		{
			patterns: []string{
				"cio",
				"",
				"",
				"(tSo|tso)",
			},
		},
		{
			patterns: []string{
				"ciu",
				"",
				"",
				"(tSu|tsu)",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"",
				"(tSi|tsI)",
			},
		},
		{
			patterns: []string{
				"ć",
				"",
				"",
				"(tS|ts)",
			},
		},
		{
			patterns: []string{
				"ssz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"sz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"sia",
				"",
				"[bcdgkpstwzż]",
				"(SB|sB|sja)",
			},
		},
		{
			patterns: []string{
				"sia",
				"",
				"",
				"(Sa|sja)",
			},
		},
		{
			patterns: []string{
				"sią",
				"",
				"[bp]",
				"(Som|som)",
			},
		},
		{
			patterns: []string{
				"sią",
				"",
				"",
				"(Son|son)",
			},
		},
		{
			patterns: []string{
				"się",
				"",
				"[bp]",
				"(Sem|sem)",
			},
		},
		{
			patterns: []string{
				"się",
				"",
				"",
				"(Sen|sen)",
			},
		},
		{
			patterns: []string{
				"sie",
				"",
				"[bcdgkpstwzż]",
				"(SF|sF|se)",
			},
		},
		{
			patterns: []string{
				"sie",
				"",
				"",
				"(Se|se)",
			},
		},
		{
			patterns: []string{
				"sio",
				"",
				"",
				"(So|so)",
			},
		},
		{
			patterns: []string{
				"siu",
				"",
				"",
				"(Su|sju)",
			},
		},
		{
			patterns: []string{
				"si",
				"",
				"",
				"(Si|sI)",
			},
		},
		{
			patterns: []string{
				"ś",
				"",
				"",
				"(S|s)",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"[bcdgkpstwzż]",
				"(ZB|zB|zja)",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"",
				"(Za|zja)",
			},
		},
		{
			patterns: []string{
				"zią",
				"",
				"[bp]",
				"(Zom|zom)",
			},
		},
		{
			patterns: []string{
				"zią",
				"",
				"",
				"(Zon|zon)",
			},
		},
		{
			patterns: []string{
				"zię",
				"",
				"[bp]",
				"(Zem|zem)",
			},
		},
		{
			patterns: []string{
				"zię",
				"",
				"",
				"(Zen|zen)",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"[bcdgkpstwzż]",
				"(ZF|zF)",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"",
				"(Ze|ze)",
			},
		},
		{
			patterns: []string{
				"zio",
				"",
				"",
				"(Zo|zo)",
			},
		},
		{
			patterns: []string{
				"ziu",
				"",
				"",
				"(Zu|zju)",
			},
		},
		{
			patterns: []string{
				"zi",
				"",
				"",
				"(Zi|zI)",
			},
		},
		{
			patterns: []string{
				"że",
				"",
				"[bcdgkpstwzż]",
				"(Ze|ZF)",
			},
		},
		{
			patterns: []string{
				"że",
				"",
				"[bcdgkpstwzż]",
				"(Ze|ZF|ze|zF)",
			},
		},
		{
			patterns: []string{
				"że",
				"",
				"",
				"Ze",
			},
		},
		{
			patterns: []string{
				"źe",
				"",
				"",
				"(Ze|ze)",
			},
		},
		{
			patterns: []string{
				"ży",
				"",
				"",
				"Zi",
			},
		},
		{
			patterns: []string{
				"źi",
				"",
				"",
				"(Zi|zi)",
			},
		},
		{
			patterns: []string{
				"ż",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"ź",
				"",
				"",
				"(Z|z)",
			},
		},
		{
			patterns: []string{
				"rze",
				"t",
				"",
				"(Se|re)",
			},
		},
		{
			patterns: []string{
				"rze",
				"",
				"",
				"(Ze|re|rZe)",
			},
		},
		{
			patterns: []string{
				"rzy",
				"t",
				"",
				"(Si|ri)",
			},
		},
		{
			patterns: []string{
				"rzy",
				"",
				"",
				"(Zi|ri|rZi)",
			},
		},
		{
			patterns: []string{
				"rz",
				"t",
				"",
				"(S|r)",
			},
		},
		{
			patterns: []string{
				"rz",
				"",
				"",
				"(Z|r|rZ)",
			},
		},
		{
			patterns: []string{
				"lio",
				"",
				"",
				"(lo|le)",
			},
		},
		{
			patterns: []string{
				"ł",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"ń",
				"",
				"",
				"n",
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
				"s",
				"",
				"s",
				"",
			},
		},
		{
			patterns: []string{
				"ó",
				"",
				"",
				"(u|o)",
			},
		},
		{
			patterns: []string{
				"ą",
				"",
				"[bp]",
				"om",
			},
		},
		{
			patterns: []string{
				"ę",
				"",
				"[bp]",
				"em",
			},
		},
		{
			patterns: []string{
				"ą",
				"",
				"",
				"on",
			},
		},
		{
			patterns: []string{
				"ę",
				"",
				"",
				"en",
			},
		},
		{
			patterns: []string{
				"ije",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"yje",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"iie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"yie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"iye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"yye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"yj",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"yi",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"iy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"rie",
				"",
				"",
				"rje",
			},
		},
		{
			patterns: []string{
				"die",
				"",
				"",
				"dje",
			},
		},
		{
			patterns: []string{
				"tie",
				"",
				"",
				"tje",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"[bcdgkpstwzż]",
				"F",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"e",
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
				"au",
				"",
				"",
				"au",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ej",
				"",
				"",
				"aj",
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
				"aj",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"i",
				"[ou]",
				"",
				"j",
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
				"[aeou]",
				"j",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"[bcdgkpstwzż]",
				"B",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"[bcdgkpstwzż]",
				"(E|F)",
			},
		},
		{
			patterns: []string{
				"o",
				"",
				"[bcćdgklłmnńrsśtwzźż]",
				"P",
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
				"ts",
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
				"E",
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
				"(h|x)",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"",
				"I",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"j",
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
				"I",
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
	ashromanian: []rule{
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
				"ce",
				"",
				"",
				"tSe",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"",
				"(tSi|tS)",
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
				"ch",
				"",
				"",
				"x",
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
				"gi",
				"",
				"",
				"(dZi|dZ)",
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
				"gh",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"i",
				"[aou]",
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
				"ţ",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: []string{
				"ş",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"h",
				"",
				"",
				"(x|h)",
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
				"î",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ea",
				"",
				"",
				"ja",
			},
		},
		{
			patterns: []string{
				"ă",
				"",
				"",
				"(e|a)",
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
				"E",
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
				"i",
				"",
				"",
				"I",
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
	ashrussian: []rule{
		{
			patterns: []string{
				"yna",
				"",
				"$",
				"(in|ina)",
			},
		},
		{
			patterns: []string{
				"ina",
				"",
				"$",
				"(in|ina)",
			},
		},
		{
			patterns: []string{
				"liova",
				"",
				"$",
				"(lof|lef)",
			},
		},
		{
			patterns: []string{
				"lova",
				"",
				"$",
				"(lof|lef|lova)",
			},
		},
		{
			patterns: []string{
				"ova",
				"",
				"$",
				"(of|ova)",
			},
		},
		{
			patterns: []string{
				"eva",
				"",
				"$",
				"(ef|ova)",
			},
		},
		{
			patterns: []string{
				"aia",
				"",
				"$",
				"(aja|i)",
			},
		},
		{
			patterns: []string{
				"aja",
				"",
				"$",
				"(aja|i)",
			},
		},
		{
			patterns: []string{
				"aya",
				"",
				"$",
				"(aja|i)",
			},
		},
		{
			patterns: []string{
				"tsya",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: []string{
				"tsyu",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: []string{
				"tsia",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: []string{
				"tsie",
				"",
				"",
				"tse",
			},
		},
		{
			patterns: []string{
				"tsio",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: []string{
				"tsye",
				"",
				"",
				"tse",
			},
		},
		{
			patterns: []string{
				"tsyo",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: []string{
				"tsiu",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: []string{
				"sie",
				"",
				"",
				"se",
			},
		},
		{
			patterns: []string{
				"sio",
				"",
				"",
				"so",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"",
				"ze",
			},
		},
		{
			patterns: []string{
				"zio",
				"",
				"",
				"zo",
			},
		},
		{
			patterns: []string{
				"sye",
				"",
				"",
				"se",
			},
		},
		{
			patterns: []string{
				"syo",
				"",
				"",
				"so",
			},
		},
		{
			patterns: []string{
				"zye",
				"",
				"",
				"ze",
			},
		},
		{
			patterns: []string{
				"zyo",
				"",
				"",
				"zo",
			},
		},
		{
			patterns: []string{
				"gauz",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: []string{
				"gaus",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: []string{
				"gol'ts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"golts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"gol'tz",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"goltz",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: []string{
				"gejmer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"gejm",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"geimer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"geim",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"geymer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: []string{
				"geym",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: []string{
				"gendler",
				"",
				"$",
				"hendler",
			},
		},
		{
			patterns: []string{
				"gof",
				"",
				"$",
				"hof",
			},
		},
		{
			patterns: []string{
				"gojf",
				"",
				"$",
				"hojf",
			},
		},
		{
			patterns: []string{
				"goyf",
				"",
				"$",
				"hojf",
			},
		},
		{
			patterns: []string{
				"goif",
				"",
				"$",
				"hojf",
			},
		},
		{
			patterns: []string{
				"ger",
				"",
				"$",
				"ger",
			},
		},
		{
			patterns: []string{
				"gen",
				"",
				"$",
				"gen",
			},
		},
		{
			patterns: []string{
				"gin",
				"",
				"$",
				"gin",
			},
		},
		{
			patterns: []string{
				"gg",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"kog",
				"^",
				"[aeoiu]",
				"(kog|koh)",
			},
		},
		{
			patterns: []string{
				"kag",
				"^",
				"[aeoiu]",
				"(kag|kah)",
			},
		},
		{
			patterns: []string{
				"g",
				"[jaeoiuy]",
				"[aeoiu]",
				"g",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[aeoiu]",
				"(g|h)",
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
				"ch",
				"",
				"",
				"(tS|x)",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"",
				"(StS|S)",
			},
		},
		{
			patterns: []string{
				"ssh",
				"",
				"",
				"S",
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
				"zh",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"$",
				"ts",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"",
				"(ts|tz)",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[iey]",
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
				"qu",
				"",
				"",
				"(kv|k)",
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
				"",
				"s",
				"",
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
				"lya",
				"",
				"",
				"la",
			},
		},
		{
			patterns: []string{
				"lyu",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: []string{
				"lia",
				"",
				"",
				"la",
			},
		},
		{
			patterns: []string{
				"liu",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: []string{
				"lja",
				"",
				"",
				"la",
			},
		},
		{
			patterns: []string{
				"lju",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: []string{
				"le",
				"",
				"",
				"(lo|lE)",
			},
		},
		{
			patterns: []string{
				"lyo",
				"",
				"",
				"(lo|le)",
			},
		},
		{
			patterns: []string{
				"lio",
				"",
				"",
				"(lo|le)",
			},
		},
		{
			patterns: []string{
				"ije",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"iye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"iie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"yje",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"yye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"yie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"iy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"yj",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"yi",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: []string{
				"io",
				"",
				"",
				"(jo|e)",
			},
		},
		{
			patterns: []string{
				"i",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: []string{
				"i",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ej",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"yo",
				"",
				"",
				"(jo|e)",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"[aiou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"iy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"yi",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: []string{
				"yj",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"e",
				"^",
				"",
				"(je|E)",
			},
		},
		{
			patterns: []string{
				"ee",
				"",
				"",
				"(aje|i)",
			},
		},
		{
			patterns: []string{
				"e",
				"[aou]",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"y",
				"",
				"",
				"I",
			},
		},
		{
			patterns: []string{
				"oo",
				"",
				"",
				"(oo|u)",
			},
		},
		{
			patterns: []string{
				"'",
				"",
				"",
				"",
			},
		},
		{
			patterns: []string{
				"\"",
				"",
				"",
				"",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"",
				"j",
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
	ashspanish: []rule{
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
				"ch",
				"",
				"",
				"(tS|dZ)",
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
				"x",
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
				"ll",
				"",
				"",
				"(l|Z)",
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
				"(x|g)",
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
				"(i|j|S|Z)",
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
				"E",
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
				"I",
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

var ashLangRules = []langRule{
	{
		pattern: "zh",
		langs:   660,
		accept:  true,
	},
	{
		pattern: "eau",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "[aoeiuäöü]h",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "^vogel",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "vogel$",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "witz",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "tz$",
		langs:   532,
		accept:  true,
	},
	{
		pattern: "^tz",
		langs:   516,
		accept:  true,
	},
	{
		pattern: "güe",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "güi",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ghe",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ghi",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "vici$",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "schi$",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "chsch",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "tsch",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ssch",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "sch$",
		langs:   528,
		accept:  true,
	},
	{
		pattern: "^sch",
		langs:   528,
		accept:  true,
	},
	{
		pattern: "^rz",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "rz$",
		langs:   144,
		accept:  true,
	},
	{
		pattern: "[^aoeiuäöü]rz",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "rz[^aoeiuäöü]",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "cki$",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ska$",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "cka$",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ue",
		langs:   528,
		accept:  true,
	},
	{
		pattern: "ae",
		langs:   532,
		accept:  true,
	},
	{
		pattern: "oe",
		langs:   540,
		accept:  true,
	},
	{
		pattern: "th$",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "^th",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "th[^aoeiu]",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "mann",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "cz",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "cy",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "niew",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "stein",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "heim$",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "heimer$",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ii$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "iy$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "yy$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "yi$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "yj$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "ij$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gaus$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gauz$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gauz$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "goltz$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gol'tz$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "golts$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gol'ts$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "^goltz",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "^gol'tz",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "^golts",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "^gol'ts",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gendler$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gejmer$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gejm$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "geimer$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "geim$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "geymer",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "geym$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "gof$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "thal",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "zweig",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "ck$",
		langs:   20,
		accept:  true,
	},
	{
		pattern: "c$",
		langs:   448,
		accept:  true,
	},
	{
		pattern: "sz",
		langs:   192,
		accept:  true,
	},
	{
		pattern: "gue",
		langs:   1032,
		accept:  true,
	},
	{
		pattern: "gui",
		langs:   1032,
		accept:  true,
	},
	{
		pattern: "guy",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "cs$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "^cs",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "dzs",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "zs$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "^zs",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "^wl",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "^wr",
		langs:   148,
		accept:  true,
	},
	{
		pattern: "gy$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "gy[aeou]",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "gy",
		langs:   576,
		accept:  true,
	},
	{
		pattern: "ly",
		langs:   704,
		accept:  true,
	},
	{
		pattern: "ny",
		langs:   704,
		accept:  true,
	},
	{
		pattern: "ty",
		langs:   704,
		accept:  true,
	},
	{
		pattern: "â",
		langs:   264,
		accept:  true,
	},
	{
		pattern: "ă",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "à",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ä",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "á",
		langs:   1088,
		accept:  true,
	},
	{
		pattern: "ą",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ć",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ç",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ę",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "é",
		langs:   1096,
		accept:  true,
	},
	{
		pattern: "è",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ê",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "í",
		langs:   1088,
		accept:  true,
	},
	{
		pattern: "î",
		langs:   264,
		accept:  true,
	},
	{
		pattern: "ł",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ń",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ñ",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ó",
		langs:   1216,
		accept:  true,
	},
	{
		pattern: "ö",
		langs:   80,
		accept:  true,
	},
	{
		pattern: "õ",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "ş",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ś",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ţ",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ü",
		langs:   80,
		accept:  true,
	},
	{
		pattern: "ù",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ű",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "ú",
		langs:   1088,
		accept:  true,
	},
	{
		pattern: "ź",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ż",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ß",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "а",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ё",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "о",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "е",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "и",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "у",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ы",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "э",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ю",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "я",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "א",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ב",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ג",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ד",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ה",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ו",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ז",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ח",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ט",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "י",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "כ",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ל",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "מ",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "נ",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ס",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ע",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "פ",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "צ",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ק",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ר",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ש",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ת",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "a",
		langs:   34,
		accept:  false,
	},
	{
		pattern: "o",
		langs:   34,
		accept:  false,
	},
	{
		pattern: "e",
		langs:   34,
		accept:  false,
	},
	{
		pattern: "i",
		langs:   34,
		accept:  false,
	},
	{
		pattern: "y",
		langs:   290,
		accept:  false,
	},
	{
		pattern: "u",
		langs:   34,
		accept:  false,
	},
	{
		pattern: "v[^aoeiuäüö]",
		langs:   16,
		accept:  false,
	},
	{
		pattern: "y[^aoeiu]",
		langs:   16,
		accept:  false,
	},
	{
		pattern: "c[^aohk]",
		langs:   16,
		accept:  false,
	},
	{
		pattern: "dzi",
		langs:   28,
		accept:  false,
	},
	{
		pattern: "ou",
		langs:   16,
		accept:  false,
	},
	{
		pattern: "aj",
		langs:   28,
		accept:  false,
	},
	{
		pattern: "ej",
		langs:   28,
		accept:  false,
	},
	{
		pattern: "oj",
		langs:   28,
		accept:  false,
	},
	{
		pattern: "uj",
		langs:   28,
		accept:  false,
	},
	{
		pattern: "k",
		langs:   256,
		accept:  false,
	},
	{
		pattern: "v",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "ky",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "eu",
		langs:   640,
		accept:  false,
	},
	{
		pattern: "w",
		langs:   1864,
		accept:  false,
	},
	{
		pattern: "kie",
		langs:   1032,
		accept:  false,
	},
	{
		pattern: "gie",
		langs:   1288,
		accept:  false,
	},
	{
		pattern: "q",
		langs:   960,
		accept:  false,
	},
	{
		pattern: "sch",
		langs:   1224,
		accept:  false,
	},
	{
		pattern: "^h",
		langs:   512,
		accept:  false,
	},
}

var ashFinalRules = finalRules{
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
					"[gdZz]",
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
					"[bgdZz]",
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
					"[bdZz]",
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
					"[bgZz]",
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
					"jnm",
					"",
					"",
					"jm",
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
					"jI",
					"^",
					"",
					"I",
				},
			},
			{
				patterns: []string{
					"a",
					"",
					"[aAB]",
					"",
				},
			},
			{
				patterns: []string{
					"a",
					"[AB]",
					"",
					"",
				},
			},
			{
				patterns: []string{
					"A",
					"",
					"A",
					"",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"B",
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
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: []string{
					"kAg",
					"^",
					"[AEOIUaeoiu]",
					"(kOg|kO[512])",
				},
			},
			{
				patterns: []string{
					"kOg",
					"^",
					"[AEOIUaeoiu]",
					"(kAg|kA[512])",
				},
			},
			{
				patterns: []string{
					"kog",
					"^",
					"[AEOIUaeoiu]",
					"(kog|ko[512])",
				},
			},
			{
				patterns: []string{
					"kag",
					"^",
					"[AEOIUaeoiu]",
					"(kag|ka[512])",
				},
			},
			{
				patterns: []string{
					"h",
					"",
					"",
					"",
				},
			},
			{
				patterns: []string{
					"H",
					"",
					"",
					"(x|)",
				},
			},
			{
				patterns: []string{
					"F",
					"",
					"[bdgkpstvzZ]h",
					"e",
				},
			},
			{
				patterns: []string{
					"F",
					"",
					"[bdgkpstvzZ]x",
					"e",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"[bdgkpstvzZ]h",
					"a",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"[bdgkpstvzZ]x",
					"a",
				},
			},
			{
				patterns: []string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(e|)",
				},
			},
			{
				patterns: []string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(i|)",
				},
			},
			{
				patterns: []string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(E|)",
				},
			},
			{
				patterns: []string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(I|)",
				},
			},
			{
				patterns: []string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(F|)",
				},
			},
			{
				patterns: []string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Q|)",
				},
			},
			{
				patterns: []string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Y|)",
				},
			},
			{
				patterns: []string{
					"lEs",
					"",
					"",
					"(lEs|lz)",
				},
			},
			{
				patterns: []string{
					"lE",
					"[bdfgkmnprStvzZ]",
					"",
					"(lE|l)",
				},
			},
			{
				patterns: []string{
					"aue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AvE",
					"",
					"",
					"(D|AvE)",
				},
			},
			{
				patterns: []string{
					"Ave",
					"",
					"",
					"(D|Ave)",
				},
			},
			{
				patterns: []string{
					"avE",
					"",
					"",
					"(D|avE)",
				},
			},
			{
				patterns: []string{
					"ave",
					"",
					"",
					"(D|ave)",
				},
			},
			{
				patterns: []string{
					"OvE",
					"",
					"",
					"(D|OvE)",
				},
			},
			{
				patterns: []string{
					"Ove",
					"",
					"",
					"(D|Ove)",
				},
			},
			{
				patterns: []string{
					"ovE",
					"",
					"",
					"(D|ovE)",
				},
			},
			{
				patterns: []string{
					"ove",
					"",
					"",
					"(D|ove)",
				},
			},
			{
				patterns: []string{
					"ea",
					"",
					"",
					"(D|ea)",
				},
			},
			{
				patterns: []string{
					"EA",
					"",
					"",
					"(D|EA)",
				},
			},
			{
				patterns: []string{
					"Ea",
					"",
					"",
					"(D|Ea)",
				},
			},
			{
				patterns: []string{
					"eA",
					"",
					"",
					"(D|eA)",
				},
			},
			{
				patterns: []string{
					"aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
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
					"lYndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lander",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lAndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lAnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"landEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lender",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lEndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lendEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lEnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"bUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"bUrg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"burg",
					"",
					"$",
					"(burk|berk)",
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
		},
		second: []secondFinalRule{
			{
				langs: 0,
				rules: []rule{
					{
						patterns: []string{
							"b",
							"",
							"",
							"(b|v[1024])",
						},
					},
					{
						patterns: []string{
							"J",
							"",
							"",
							"z",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"AiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"OiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"UiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"AiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"OiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"UiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[bp]",
							"(o|om[128]|im[128])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[dgkstvz]",
							"(a|o|on[128]|in[128])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"AiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"OiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"UiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"AiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"OiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"UiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[bp]",
							"(i|im[128]|om[128])",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in[128]|on[128])",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiouAEIBFOUQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^aeiouAEBFIOU]e",
							"(Q[16]|i|D[4])",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk[16])",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts[16])",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(Q[16]|i)",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"$",
							"(li|il[4])",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[4]|lY[16])",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: []string{
							"Ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: []string{
							"iA",
							"",
							"$",
							"(ia|io)",
						},
					},
					{
						patterns: []string{
							"iA",
							"",
							"",
							"(ia|io|iY[16])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[16]|D[4])",
						},
					},
					{
						patterns: []string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: []string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"[fklmnprstv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"[fklmnprstv]$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"[oeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"",
							"(o|Y[16])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"[oeiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"",
							"(a|o|Y[16])",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"[DoiuQY]",
							"",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: []string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk[16])",
						},
					},
					{
						patterns: []string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: []string{
							"sUts",
							"",
							"$",
							"(suts|sQts[16])",
						},
					},
					{
						patterns: []string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"",
							"(u|Q[16])",
						},
					},
				},
			},
			{
				langs: 9,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: []string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: []string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: []string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: []string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: []string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 1,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: []string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: []string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: []string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: []string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: []string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 2,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"[^aEIeiou]e",
							"(Q|i|D)",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(il|li|lY)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"E",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: []string{
							"e",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiEuQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
				},
			},
			{
				langs: 3,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 4,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiAEIOUouQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(Q|i)",
						},
					},
					{
						patterns: []string{
							"AU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"aU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"Au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"OU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"oU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"Ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoAOUiuQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"[fklmnprst]$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"[aoAOUeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"",
							"(o|Y)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"[aoeOUiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"",
							"(a|o|Y)",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"[DaoiuUQY]",
							"",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: []string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk)",
						},
					},
					{
						patterns: []string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: []string{
							"sUts",
							"",
							"$",
							"(suts|sQts)",
						},
					},
					{
						patterns: []string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"",
							"(u|Q)",
						},
					},
				},
			},
			{
				langs: 5,
				rules: []rule{},
			},
			{
				langs: 6,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 7,
				rules: []rule{
					{
						patterns: []string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 8,
				rules: []rule{
					{
						patterns: []string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 10,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
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
					"[gdZz]",
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
					"[bgdZz]",
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
					"[bdZz]",
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
					"[bgZz]",
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
					"jnm",
					"",
					"",
					"jm",
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
					"jI",
					"^",
					"",
					"I",
				},
			},
			{
				patterns: []string{
					"a",
					"",
					"[aAB]",
					"",
				},
			},
			{
				patterns: []string{
					"a",
					"[AB]",
					"",
					"",
				},
			},
			{
				patterns: []string{
					"A",
					"",
					"A",
					"",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"B",
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
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: []string{
					"kAg",
					"^",
					"[AEOIUaeoiu]",
					"(kOg|kO[512])",
				},
			},
			{
				patterns: []string{
					"kOg",
					"^",
					"[AEOIUaeoiu]",
					"(kAg|kA[512])",
				},
			},
			{
				patterns: []string{
					"kog",
					"^",
					"[AEOIUaeoiu]",
					"(kog|ko[512])",
				},
			},
			{
				patterns: []string{
					"kag",
					"^",
					"[AEOIUaeoiu]",
					"(kag|ka[512])",
				},
			},
			{
				patterns: []string{
					"h",
					"",
					"",
					"",
				},
			},
			{
				patterns: []string{
					"H",
					"",
					"",
					"(x|)",
				},
			},
			{
				patterns: []string{
					"F",
					"",
					"[bdgkpstvzZ]h",
					"e",
				},
			},
			{
				patterns: []string{
					"F",
					"",
					"[bdgkpstvzZ]x",
					"e",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"[bdgkpstvzZ]h",
					"a",
				},
			},
			{
				patterns: []string{
					"B",
					"",
					"[bdgkpstvzZ]x",
					"a",
				},
			},
			{
				patterns: []string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(e|)",
				},
			},
			{
				patterns: []string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(i|)",
				},
			},
			{
				patterns: []string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(E|)",
				},
			},
			{
				patterns: []string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(I|)",
				},
			},
			{
				patterns: []string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(F|)",
				},
			},
			{
				patterns: []string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Q|)",
				},
			},
			{
				patterns: []string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Y|)",
				},
			},
			{
				patterns: []string{
					"lEs",
					"",
					"",
					"(lEs|lz)",
				},
			},
			{
				patterns: []string{
					"lE",
					"[bdfgkmnprStvzZ]",
					"",
					"(lE|l)",
				},
			},
			{
				patterns: []string{
					"aue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AvE",
					"",
					"",
					"(D|AvE)",
				},
			},
			{
				patterns: []string{
					"Ave",
					"",
					"",
					"(D|Ave)",
				},
			},
			{
				patterns: []string{
					"avE",
					"",
					"",
					"(D|avE)",
				},
			},
			{
				patterns: []string{
					"ave",
					"",
					"",
					"(D|ave)",
				},
			},
			{
				patterns: []string{
					"OvE",
					"",
					"",
					"(D|OvE)",
				},
			},
			{
				patterns: []string{
					"Ove",
					"",
					"",
					"(D|Ove)",
				},
			},
			{
				patterns: []string{
					"ovE",
					"",
					"",
					"(D|ovE)",
				},
			},
			{
				patterns: []string{
					"ove",
					"",
					"",
					"(D|ove)",
				},
			},
			{
				patterns: []string{
					"ea",
					"",
					"",
					"(D|ea)",
				},
			},
			{
				patterns: []string{
					"EA",
					"",
					"",
					"(D|EA)",
				},
			},
			{
				patterns: []string{
					"Ea",
					"",
					"",
					"(D|Ea)",
				},
			},
			{
				patterns: []string{
					"eA",
					"",
					"",
					"(D|eA)",
				},
			},
			{
				patterns: []string{
					"aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ajU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ojO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"OjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ejO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"EjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ujO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"UjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"ijO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"IjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: []string{
					"AjU",
					"",
					"",
					"D",
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
					"lYndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lander",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lAndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lAnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"landEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lender",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lEndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lendEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"lEnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: []string{
					"bUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"bUrg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"burg",
					"",
					"$",
					"(burk|berk)",
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
		},
		second: []secondFinalRule{
			{
				langs: 0,
				rules: []rule{
					{
						patterns: []string{
							"b",
							"",
							"",
							"(b|v[1024])",
						},
					},
					{
						patterns: []string{
							"J",
							"",
							"",
							"z",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"AiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"OiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"UiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"AiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"OiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"UiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[bp]",
							"(o|om[128]|im[128])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[dgkstvz]",
							"(a|o|on[128]|in[128])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"AiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"OiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"UiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"AiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"OiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"UiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[bp]",
							"(i|im[128]|om[128])",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in[128]|on[128])",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiouAEIBFOUQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^aeiouAEBFIOU]e",
							"(Q[16]|i|D[4])",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk[16])",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts[16])",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(Q[16]|i)",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"$",
							"(li|il[4])",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[4]|lY[16])",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: []string{
							"Ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: []string{
							"iA",
							"",
							"$",
							"(ia|io)",
						},
					},
					{
						patterns: []string{
							"iA",
							"",
							"",
							"(ia|io|iY[16])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[16]|D[4])",
						},
					},
					{
						patterns: []string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: []string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"[fklmnprstv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"[fklmnprstv]$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"[oeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"",
							"(o|Y[16])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"[oeiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"",
							"(a|o|Y[16])",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"[DoiuQY]",
							"",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: []string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk[16])",
						},
					},
					{
						patterns: []string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: []string{
							"sUts",
							"",
							"$",
							"(suts|sQts[16])",
						},
					},
					{
						patterns: []string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"",
							"(u|Q[16])",
						},
					},
				},
			},
			{
				langs: 9,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: []string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: []string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: []string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: []string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: []string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 1,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: []string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: []string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: []string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: []string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: []string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 2,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"[^aEIeiou]e",
							"(Q|i|D)",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(il|li|lY)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"E",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: []string{
							"e",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiEuQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
				},
			},
			{
				langs: 3,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 4,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiAEIOUouQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(Q|i)",
						},
					},
					{
						patterns: []string{
							"AU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"aU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"Au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"OU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"oU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"Ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoAOUiuQY]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"[fklmnprst]$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"[aoAOUeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"",
							"(o|Y)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"[aoeOUiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"",
							"(a|o|Y)",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"[DaoiuUQY]",
							"",
							"u",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: []string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk)",
						},
					},
					{
						patterns: []string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: []string{
							"sUts",
							"",
							"$",
							"(suts|sQts)",
						},
					},
					{
						patterns: []string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"",
							"(u|Q)",
						},
					},
				},
			},
			{
				langs: 5,
				rules: []rule{},
			},
			{
				langs: 6,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 7,
				rules: []rule{
					{
						patterns: []string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 8,
				rules: []rule{
					{
						patterns: []string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: []string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: []string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
				},
			},
			{
				langs: 10,
				rules: []rule{
					{
						patterns: []string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: []string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: []string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: []string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: []string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: []string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: []string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: []string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: []string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: []string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: []string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: []string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: []string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: []string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
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
