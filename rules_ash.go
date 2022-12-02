// GENERATED CODE. DO NOT EDIT!
package bmpm

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
			patterns: [4]string{
				"yna",
				"",
				"$",
				"(in[512]|ina)",
			},
		},
		{
			patterns: [4]string{
				"ina",
				"",
				"$",
				"(in[512]|ina)",
			},
		},
		{
			patterns: [4]string{
				"liova",
				"",
				"$",
				"(lof[512]|lef[512]|lova)",
			},
		},
		{
			patterns: [4]string{
				"lova",
				"",
				"$",
				"(lof[512]|lef[512]|lova)",
			},
		},
		{
			patterns: [4]string{
				"ova",
				"",
				"$",
				"(of[512]|ova)",
			},
		},
		{
			patterns: [4]string{
				"eva",
				"",
				"$",
				"(ef[512]|eva)",
			},
		},
		{
			patterns: [4]string{
				"aia",
				"",
				"$",
				"(aja|i[512])",
			},
		},
		{
			patterns: [4]string{
				"aja",
				"",
				"$",
				"(aja|i[512])",
			},
		},
		{
			patterns: [4]string{
				"aya",
				"",
				"$",
				"(aja|i[512])",
			},
		},
		{
			patterns: [4]string{
				"lowa",
				"",
				"$",
				"(lova|lof[128]|l[128]|el[128])",
			},
		},
		{
			patterns: [4]string{
				"kowa",
				"",
				"$",
				"(kova|kof[128]|k[128]|ek[128])",
			},
		},
		{
			patterns: [4]string{
				"owa",
				"",
				"$",
				"(ova|of[128]|)",
			},
		},
		{
			patterns: [4]string{
				"lowna",
				"",
				"$",
				"(lovna|levna|l[128]|el[128])",
			},
		},
		{
			patterns: [4]string{
				"kowna",
				"",
				"$",
				"(kovna|k[128]|ek[128])",
			},
		},
		{
			patterns: [4]string{
				"owna",
				"",
				"$",
				"(ovna|[128])",
			},
		},
		{
			patterns: [4]string{
				"lówna",
				"",
				"$",
				"(l|el[128])",
			},
		},
		{
			patterns: [4]string{
				"kówna",
				"",
				"$",
				"(k|ek[128])",
			},
		},
		{
			patterns: [4]string{
				"ówna",
				"",
				"$",
				"",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"$",
				"(a|i[128])",
			},
		},
		{
			patterns: [4]string{
				"rh",
				"^",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"ssch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"chsch",
				"",
				"",
				"xS",
			},
		},
		{
			patterns: [4]string{
				"tsch",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"[ei]",
				"(sk[256]|S|StS[512])",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"",
				"(S|StS[512])",
			},
		},
		{
			patterns: [4]string{
				"ssh",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sh",
				"",
				"[äöü]",
				"sh",
			},
		},
		{
			patterns: [4]string{
				"sh",
				"",
				"[aeiou]",
				"(S[516]|sh)",
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
				"(x[516]|kh)",
			},
		},
		{
			patterns: [4]string{
				"chs",
				"",
				"",
				"(ks[16]|xs|tSs[516])",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"[ei]",
				"(x|k[256]|tS[516])",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"(x|tS[516])",
			},
		},
		{
			patterns: [4]string{
				"ck",
				"",
				"",
				"(k|tsk[128])",
			},
		},
		{
			patterns: [4]string{
				"czy",
				"",
				"",
				"tSi",
			},
		},
		{
			patterns: [4]string{
				"cze",
				"",
				"[bcdgkpstwzż]",
				"(tSe|tSF)",
			},
		},
		{
			patterns: [4]string{
				"ciewicz",
				"",
				"",
				"(tsevitS|tSevitS)",
			},
		},
		{
			patterns: [4]string{
				"siewicz",
				"",
				"",
				"(sevitS|SevitS)",
			},
		},
		{
			patterns: [4]string{
				"ziewicz",
				"",
				"",
				"(zevitS|ZevitS)",
			},
		},
		{
			patterns: [4]string{
				"riewicz",
				"",
				"",
				"rjevitS",
			},
		},
		{
			patterns: [4]string{
				"diewicz",
				"",
				"",
				"djevitS",
			},
		},
		{
			patterns: [4]string{
				"tiewicz",
				"",
				"",
				"tjevitS",
			},
		},
		{
			patterns: [4]string{
				"iewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: [4]string{
				"ewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: [4]string{
				"owicz",
				"",
				"",
				"ovitS",
			},
		},
		{
			patterns: [4]string{
				"icz",
				"",
				"",
				"itS",
			},
		},
		{
			patterns: [4]string{
				"cz",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"cia",
				"",
				"[bcdgkpstwzż]",
				"(tSB[128]|tsB)",
			},
		},
		{
			patterns: [4]string{
				"cia",
				"",
				"",
				"(tSa[128]|tsa)",
			},
		},
		{
			patterns: [4]string{
				"cią",
				"",
				"[bp]",
				"(tSom[128]|tsom)",
			},
		},
		{
			patterns: [4]string{
				"cią",
				"",
				"",
				"(tSon[128]|tson)",
			},
		},
		{
			patterns: [4]string{
				"cię",
				"",
				"[bp]",
				"(tSem[128]|tsem)",
			},
		},
		{
			patterns: [4]string{
				"cię",
				"",
				"",
				"(tSen[128]|tsen)",
			},
		},
		{
			patterns: [4]string{
				"cie",
				"",
				"[bcdgkpstwzż]",
				"(tSF[128]|tsF)",
			},
		},
		{
			patterns: [4]string{
				"cie",
				"",
				"",
				"(tSe[128]|tse)",
			},
		},
		{
			patterns: [4]string{
				"cio",
				"",
				"",
				"(tSo[128]|tso)",
			},
		},
		{
			patterns: [4]string{
				"ciu",
				"",
				"",
				"(tSu[128]|tsu)",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"$",
				"(tsi[128]|tSi[384]|tS[256]|si)",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"",
				"(tsi[128]|tSi[384]|si)",
			},
		},
		{
			patterns: [4]string{
				"ce",
				"",
				"[bcdgkpstwzż]",
				"(tsF[128]|tSe[384]|se)",
			},
		},
		{
			patterns: [4]string{
				"ce",
				"",
				"",
				"(tSe[384]|tse[128]|se)",
			},
		},
		{
			patterns: [4]string{
				"cy",
				"",
				"",
				"(si|tsi[128])",
			},
		},
		{
			patterns: [4]string{
				"ssz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"ssp",
				"",
				"",
				"(Sp[16]|sp)",
			},
		},
		{
			patterns: [4]string{
				"sp",
				"",
				"",
				"(Sp[16]|sp)",
			},
		},
		{
			patterns: [4]string{
				"sst",
				"",
				"",
				"(St[16]|st)",
			},
		},
		{
			patterns: [4]string{
				"st",
				"",
				"",
				"(St[16]|st)",
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
				"sia",
				"",
				"[bcdgkpstwzż]",
				"(SB[128]|sB[128]|sja)",
			},
		},
		{
			patterns: [4]string{
				"sia",
				"",
				"",
				"(Sa[128]|sja)",
			},
		},
		{
			patterns: [4]string{
				"sią",
				"",
				"[bp]",
				"(Som[128]|som)",
			},
		},
		{
			patterns: [4]string{
				"sią",
				"",
				"",
				"(Son[128]|son)",
			},
		},
		{
			patterns: [4]string{
				"się",
				"",
				"[bp]",
				"(Sem[128]|sem)",
			},
		},
		{
			patterns: [4]string{
				"się",
				"",
				"",
				"(Sen[128]|sen)",
			},
		},
		{
			patterns: [4]string{
				"sie",
				"",
				"[bcdgkpstwzż]",
				"(SF[128]|sF|zi[16])",
			},
		},
		{
			patterns: [4]string{
				"sie",
				"",
				"",
				"(se|Se[128]|zi[16])",
			},
		},
		{
			patterns: [4]string{
				"sio",
				"",
				"",
				"(So[128]|so)",
			},
		},
		{
			patterns: [4]string{
				"siu",
				"",
				"",
				"(Su[128]|sju)",
			},
		},
		{
			patterns: [4]string{
				"si",
				"",
				"",
				"(Si[128]|si|zi[16])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"[aeiouäöë]",
				"(s|z[16])",
			},
		},
		{
			patterns: [4]string{
				"gue",
				"",
				"",
				"ge",
			},
		},
		{
			patterns: [4]string{
				"gui",
				"",
				"",
				"gi",
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
				"gh",
				"",
				"[ei]",
				"(g[256]|gh)",
			},
		},
		{
			patterns: [4]string{
				"gauz",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: [4]string{
				"gaus",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: [4]string{
				"gol'ts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"golts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"gol'tz",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"goltz",
				"",
				"",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"gol'ts",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"golts",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"gol'tz",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"goltz",
				"^",
				"",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"gendler",
				"",
				"$",
				"hendler",
			},
		},
		{
			patterns: [4]string{
				"gejmer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"gejm",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"geymer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"geym",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"geimer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"geim",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"gof",
				"",
				"$",
				"hof",
			},
		},
		{
			patterns: [4]string{
				"ger",
				"",
				"$",
				"ger",
			},
		},
		{
			patterns: [4]string{
				"gen",
				"",
				"$",
				"gen",
			},
		},
		{
			patterns: [4]string{
				"gin",
				"",
				"$",
				"gin",
			},
		},
		{
			patterns: [4]string{
				"gie",
				"",
				"$",
				"(ge|gi[16]|ji[8])",
			},
		},
		{
			patterns: [4]string{
				"gie",
				"",
				"",
				"ge",
			},
		},
		{
			patterns: [4]string{
				"ge",
				"[yaeiou]",
				"",
				"(gE|xe[1024]|dZe[260])",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"[yaeiou]",
				"",
				"(gI|xi[1024]|dZi[260])",
			},
		},
		{
			patterns: [4]string{
				"ge",
				"",
				"",
				"(gE|dZe[260]|hE[512]|xe[1024])",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"",
				"(gI|dZi[260]|hI[512]|xi[1024])",
			},
		},
		{
			patterns: [4]string{
				"gy",
				"",
				"[aeouáéóúüöőű]",
				"(gi|dj[64])",
			},
		},
		{
			patterns: [4]string{
				"gy",
				"",
				"",
				"(gi|d[64])",
			},
		},
		{
			patterns: [4]string{
				"g",
				"[jyaeiou]",
				"[aouyei]",
				"g",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[aouei]",
				"(g|h[512])",
			},
		},
		{
			patterns: [4]string{
				"ej",
				"",
				"",
				"(aj|eZ[264]|ex[1024])",
			},
		},
		{
			patterns: [4]string{
				"ej",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ly",
				"",
				"[au]",
				"(l|lj)",
			},
		},
		{
			patterns: [4]string{
				"li",
				"",
				"[au]",
				"(l|lj)",
			},
		},
		{
			patterns: [4]string{
				"lj",
				"",
				"[au]",
				"(l|lj)",
			},
		},
		{
			patterns: [4]string{
				"lio",
				"",
				"",
				"(lo|le[512])",
			},
		},
		{
			patterns: [4]string{
				"lyo",
				"",
				"",
				"(lo|le[512])",
			},
		},
		{
			patterns: [4]string{
				"ll",
				"",
				"",
				"(l|J[1024])",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"[aoeiuy]",
				"(j|dZ[4]|x[1024]|Z[264])",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"(j|x[1024])",
			},
		},
		{
			patterns: [4]string{
				"pf",
				"",
				"",
				"(pf|p|f)",
			},
		},
		{
			patterns: [4]string{
				"ph",
				"",
				"",
				"(ph|f)",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"(kv[16]|k)",
			},
		},
		{
			patterns: [4]string{
				"rze",
				"t",
				"",
				"(Se[128]|re)",
			},
		},
		{
			patterns: [4]string{
				"rze",
				"",
				"",
				"(rze|rtsE[16]|Ze[128]|re[128]|rZe[128])",
			},
		},
		{
			patterns: [4]string{
				"rzy",
				"t",
				"",
				"(Si[128]|ri)",
			},
		},
		{
			patterns: [4]string{
				"rzy",
				"",
				"",
				"(Zi[128]|ri[128]|rZi)",
			},
		},
		{
			patterns: [4]string{
				"rz",
				"t",
				"",
				"(S[128]|r)",
			},
		},
		{
			patterns: [4]string{
				"rz",
				"",
				"",
				"(rz|rts[16]|Z[128]|r[128]|rZ[128])",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"$",
				"(ts|tS[20])",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"^",
				"",
				"(ts|tS[20])",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"",
				"(ts[532]|tz)",
			},
		},
		{
			patterns: [4]string{
				"zh",
				"",
				"",
				"(Z|zh[128]|tsh[16])",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"[bcdgkpstwzż]",
				"(ZB[128]|zB[128]|zja)",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"",
				"(Za[128]|zja)",
			},
		},
		{
			patterns: [4]string{
				"zią",
				"",
				"[bp]",
				"(Zom[128]|zom)",
			},
		},
		{
			patterns: [4]string{
				"zią",
				"",
				"",
				"(Zon[128]|zon)",
			},
		},
		{
			patterns: [4]string{
				"zię",
				"",
				"[bp]",
				"(Zem[128]|zem)",
			},
		},
		{
			patterns: [4]string{
				"zię",
				"",
				"",
				"(Zen[128]|zen)",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"[bcdgkpstwzż]",
				"(ZF[128]|zF[128]|ze|tsi[16])",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"",
				"(ze|Ze[128]|tsi[16])",
			},
		},
		{
			patterns: [4]string{
				"zio",
				"",
				"",
				"(Zo[128]|zo)",
			},
		},
		{
			patterns: [4]string{
				"ziu",
				"",
				"",
				"(Zu[128]|zju)",
			},
		},
		{
			patterns: [4]string{
				"zi",
				"",
				"",
				"(Zi[128]|zi|tsi[16])",
			},
		},
		{
			patterns: [4]string{
				"thal",
				"",
				"$",
				"tal",
			},
		},
		{
			patterns: [4]string{
				"th",
				"^",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"th",
				"",
				"[aeiou]",
				"(t[16]|th)",
			},
		},
		{
			patterns: [4]string{
				"th",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"vogel",
				"",
				"",
				"(vogel|fogel[16])",
			},
		},
		{
			patterns: [4]string{
				"v",
				"^",
				"",
				"(v|f[16])",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[aeiouyäöü]",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"(h|x[384])",
			},
		},
		{
			patterns: [4]string{
				"h",
				"^",
				"",
				"(h|H[20])",
			},
		},
		{
			patterns: [4]string{
				"yi",
				" ",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"iy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"e",
				"in",
				"$",
				"(e|[8])",
			},
		},
		{
			patterns: [4]string{
				"yj",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"$",
				"i",
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
				"oue",
				"",
				"",
				"oue",
			},
		},
		{
			patterns: [4]string{
				"au",
				"",
				"",
				"(au|o[8])",
			},
		},
		{
			patterns: [4]string{
				"ou",
				"",
				"",
				"(ou|u[8])",
			},
		},
		{
			patterns: [4]string{
				"ue",
				"",
				"",
				"(Q|uje[512])",
			},
		},
		{
			patterns: [4]string{
				"ae",
				"",
				"",
				"(Y[16]|aje[512]|ae)",
			},
		},
		{
			patterns: [4]string{
				"oe",
				"",
				"",
				"(Y[16]|oje[512]|oe)",
			},
		},
		{
			patterns: [4]string{
				"ee",
				"",
				"",
				"(i[4]|aje[512]|e)",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"(aj[16]|oj[16]|eu)",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"[bcdgkpstwzż]",
				"(i[16]|e[128]|ije[512]|je)",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"(i[16]|e[128]|ije[512]|je)",
			},
		},
		{
			patterns: [4]string{
				"ye",
				"",
				"",
				"(je|ije[512])",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"io",
				"",
				"",
				"(jo|e[512])",
			},
		},
		{
			patterns: [4]string{
				"yo",
				"",
				"",
				"(jo|e[512])",
			},
		},
		{
			patterns: [4]string{
				"ea",
				"",
				"",
				"(ea|ja[256])",
			},
		},
		{
			patterns: [4]string{
				"e",
				"^",
				"",
				"(e|je[512])",
			},
		},
		{
			patterns: [4]string{
				"oo",
				"",
				"",
				"(u[4]|o)",
			},
		},
		{
			patterns: [4]string{
				"uu",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ć",
				"",
				"",
				"(tS[128]|ts)",
			},
		},
		{
			patterns: [4]string{
				"ł",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"ń",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"ñ",
				"",
				"",
				"(n|nj[1024])",
			},
		},
		{
			patterns: [4]string{
				"ś",
				"",
				"",
				"(S[128]|s)",
			},
		},
		{
			patterns: [4]string{
				"ş",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"ţ",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: [4]string{
				"ż",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"ź",
				"",
				"",
				"(Z[128]|z)",
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
				"ą",
				"",
				"[bp]",
				"om",
			},
		},
		{
			patterns: [4]string{
				"ą",
				"",
				"",
				"on",
			},
		},
		{
			patterns: [4]string{
				"ä",
				"",
				"",
				"(Y|e)",
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
				"ă",
				"",
				"",
				"(e[256]|a)",
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
				"é",
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
				"ê",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ę",
				"",
				"[bp]",
				"em",
			},
		},
		{
			patterns: [4]string{
				"ę",
				"",
				"",
				"en",
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
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"ő",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"ó",
				"",
				"",
				"(u[128]|o)",
			},
		},
		{
			patterns: [4]string{
				"ű",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: [4]string{
				"ü",
				"",
				"",
				"Q",
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
				"ű",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: [4]string{
				"ß",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"'",
				"",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"\"",
				"",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"[bcdgkpstwzż]",
				"(A|B[128])",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"[bcdgkpstwzż]",
				"(E|F[128])",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"[bcćdgklłmnńrsśtwzźż]",
				"(O|P[128])",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"A",
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
				"(k|ts[128])",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"j",
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
				"O",
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
				"U",
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
				"(ts[16]|z)",
			},
		},
	},
	ashcyrillic: []rule{
		{
			patterns: [4]string{
				"ця",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: [4]string{
				"цю",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: [4]string{
				"циа",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: [4]string{
				"цие",
				"",
				"",
				"tse",
			},
		},
		{
			patterns: [4]string{
				"цио",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: [4]string{
				"циу",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: [4]string{
				"сие",
				"",
				"",
				"se",
			},
		},
		{
			patterns: [4]string{
				"сио",
				"",
				"",
				"so",
			},
		},
		{
			patterns: [4]string{
				"зие",
				"",
				"",
				"ze",
			},
		},
		{
			patterns: [4]string{
				"зио",
				"",
				"",
				"zo",
			},
		},
		{
			patterns: [4]string{
				"гауз",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: [4]string{
				"гаус",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: [4]string{
				"гольц",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"геймер",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"гейм",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"гоф",
				"",
				"$",
				"hof",
			},
		},
		{
			patterns: [4]string{
				"гер",
				"",
				"$",
				"ger",
			},
		},
		{
			patterns: [4]string{
				"ген",
				"",
				"$",
				"gen",
			},
		},
		{
			patterns: [4]string{
				"гин",
				"",
				"$",
				"gin",
			},
		},
		{
			patterns: [4]string{
				"г",
				"(й|ё|я|ю|ы|а|е|о|и|у)",
				"(а|е|о|и|у)",
				"g",
			},
		},
		{
			patterns: [4]string{
				"г",
				"",
				"(а|е|о|и|у)",
				"(g|h)",
			},
		},
		{
			patterns: [4]string{
				"ля",
				"",
				"",
				"la",
			},
		},
		{
			patterns: [4]string{
				"лю",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: [4]string{
				"лё",
				"",
				"",
				"(le|lo)",
			},
		},
		{
			patterns: [4]string{
				"лио",
				"",
				"",
				"(le|lo)",
			},
		},
		{
			patterns: [4]string{
				"ле",
				"",
				"",
				"(lE|lo)",
			},
		},
		{
			patterns: [4]string{
				"ийе",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ие",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ыйе",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ые",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ий",
				"",
				"(а|о|у)",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ый",
				"",
				"(а|о|у)",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ий",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ый",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ё",
				"",
				"",
				"(e|jo)",
			},
		},
		{
			patterns: [4]string{
				"ей",
				"^",
				"",
				"(jaj|aj)",
			},
		},
		{
			patterns: [4]string{
				"е",
				"(а|е|о|у)",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"е",
				"^",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"эй",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ей",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ауе",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: [4]string{
				"ауэ",
				"",
				"",
				"aue",
			},
		},
		{
			patterns: [4]string{
				"а",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"б",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"в",
				"",
				"",
				"v",
			},
		},
		{
			patterns: [4]string{
				"г",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"д",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"е",
				"",
				"",
				"E",
			},
		},
		{
			patterns: [4]string{
				"ж",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"з",
				"",
				"",
				"z",
			},
		},
		{
			patterns: [4]string{
				"и",
				"",
				"",
				"I",
			},
		},
		{
			patterns: [4]string{
				"й",
				"",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"к",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"л",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"м",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"н",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"о",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"п",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"р",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"с",
				"",
				"с",
				"",
			},
		},
		{
			patterns: [4]string{
				"с",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"т",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"у",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ф",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"х",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"ц",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: [4]string{
				"ч",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ш",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"щ",
				"",
				"",
				"StS",
			},
		},
		{
			patterns: [4]string{
				"ъ",
				"",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"ы",
				"",
				"",
				"I",
			},
		},
		{
			patterns: [4]string{
				"ь",
				"",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"э",
				"",
				"",
				"E",
			},
		},
		{
			patterns: [4]string{
				"ю",
				"",
				"",
				"ju",
			},
		},
		{
			patterns: [4]string{
				"я",
				"",
				"",
				"ja",
			},
		},
	},
	ashenglish: []rule{
		{
			patterns: [4]string{
				"tch",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"(tS|x)",
			},
		},
		{
			patterns: [4]string{
				"ck",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"cc",
				"",
				"[iey]",
				"ks",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"c",
				"",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[iey]",
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
				"gh",
				"^",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"gh",
				"",
				"",
				"(g|f|w)",
			},
		},
		{
			patterns: [4]string{
				"gn",
				"",
				"",
				"(gn|n)",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[iey]",
				"(g|dZ)",
			},
		},
		{
			patterns: [4]string{
				"th",
				"",
				"",
				"t",
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
				"ph",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"",
				"(S|sk)",
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
				"who",
				"^",
				"",
				"hu",
			},
		},
		{
			patterns: [4]string{
				"wh",
				"^",
				"",
				"w",
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
				"h",
				"",
				"[^aeiou]",
				"",
			},
		},
		{
			patterns: [4]string{
				"h",
				"^",
				"",
				"H",
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
				"j",
				"",
				"",
				"dZ",
			},
		},
		{
			patterns: [4]string{
				"kn",
				"^",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"mb",
				"",
				"$",
				"m",
			},
		},
		{
			patterns: [4]string{
				"ng",
				"",
				"$",
				"(N|ng)",
			},
		},
		{
			patterns: [4]string{
				"pn",
				"^",
				"",
				"(pn|n)",
			},
		},
		{
			patterns: [4]string{
				"ps",
				"^",
				"",
				"(ps|s)",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"kw",
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
				"tia",
				"",
				"",
				"(So|Sa)",
			},
		},
		{
			patterns: [4]string{
				"tio",
				"",
				"",
				"So",
			},
		},
		{
			patterns: [4]string{
				"wr",
				"^",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"(w|v)",
			},
		},
		{
			patterns: [4]string{
				"x",
				"^",
				"",
				"z",
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
				"yi",
				" ",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"y",
				"^",
				"[aeiouy]",
				"j",
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
				"oue",
				"",
				"",
				"(aue|oue)",
			},
		},
		{
			patterns: [4]string{
				"ai",
				"",
				"",
				"(aj|e)",
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
				"a",
				"",
				"[^aeiou]e",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"",
				"(e|o|a)",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"(aj|i)",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"(aj|i)",
			},
		},
		{
			patterns: [4]string{
				"ear",
				"",
				"",
				"ia",
			},
		},
		{
			patterns: [4]string{
				"ea",
				"",
				"",
				"(i|e)",
			},
		},
		{
			patterns: [4]string{
				"ee",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"[^aeiou]e",
				"i",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"$",
				"(|E)",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"",
				"E",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[^aeiou]e",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"I",
			},
		},
		{
			patterns: [4]string{
				"oa",
				"",
				"",
				"ou",
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
				"oo",
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
				"(u|ou)",
			},
		},
		{
			patterns: [4]string{
				"oy",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"[^aeiou]e",
				"ou",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"",
				"(o|a)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"[^aeiou]e",
				"(ju|u)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"r",
				"(e|u)",
			},
		},
		{
			patterns: [4]string{
				"u",
				"",
				"",
				"(u|a)",
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
	ashfrench: []rule{
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
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"aj",
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
				"yi",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				"",
				"i",
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
				"E",
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
				"I",
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
	ashgerman: []rule{
		{
			patterns: [4]string{
				"ziu",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: [4]string{
				"zio",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: [4]string{
				"ssch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"chsch",
				"",
				"",
				"xS",
			},
		},
		{
			patterns: [4]string{
				"ewitsch",
				"",
				"$",
				"evitS",
			},
		},
		{
			patterns: [4]string{
				"owitsch",
				"",
				"$",
				"ovitS",
			},
		},
		{
			patterns: [4]string{
				"evitsch",
				"",
				"$",
				"evitS",
			},
		},
		{
			patterns: [4]string{
				"ovitsch",
				"",
				"$",
				"ovitS",
			},
		},
		{
			patterns: [4]string{
				"witsch",
				"",
				"$",
				"vitS",
			},
		},
		{
			patterns: [4]string{
				"vitsch",
				"",
				"$",
				"vitS",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"chs",
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
				"x",
			},
		},
		{
			patterns: [4]string{
				"ck",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[eiy]",
				"ts",
			},
		},
		{
			patterns: [4]string{
				"sp",
				"^",
				"",
				"Sp",
			},
		},
		{
			patterns: [4]string{
				"st",
				"^",
				"",
				"St",
			},
		},
		{
			patterns: [4]string{
				"ssp",
				"",
				"",
				"(Sp|sp)",
			},
		},
		{
			patterns: [4]string{
				"sp",
				"",
				"",
				"(Sp|sp)",
			},
		},
		{
			patterns: [4]string{
				"sst",
				"",
				"",
				"(St|st)",
			},
		},
		{
			patterns: [4]string{
				"st",
				"",
				"",
				"(St|st)",
			},
		},
		{
			patterns: [4]string{
				"pf",
				"",
				"",
				"(pf|p|f)",
			},
		},
		{
			patterns: [4]string{
				"ph",
				"",
				"",
				"(ph|f)",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"kv",
			},
		},
		{
			patterns: [4]string{
				"ewitz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: [4]string{
				"ewiz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: [4]string{
				"evitz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: [4]string{
				"eviz",
				"",
				"$",
				"(evits|evitS)",
			},
		},
		{
			patterns: [4]string{
				"owitz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: [4]string{
				"owiz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: [4]string{
				"ovitz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: [4]string{
				"oviz",
				"",
				"$",
				"(ovits|ovitS)",
			},
		},
		{
			patterns: [4]string{
				"witz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: [4]string{
				"wiz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: [4]string{
				"vitz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: [4]string{
				"viz",
				"",
				"$",
				"(vits|vitS)",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: [4]string{
				"thal",
				"",
				"$",
				"tal",
			},
		},
		{
			patterns: [4]string{
				"th",
				"^",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"th",
				"",
				"[äöüaeiou]",
				"(t|th)",
			},
		},
		{
			patterns: [4]string{
				"th",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"rh",
				"^",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[aeiouyäöü]",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"h",
				"^",
				"",
				"H",
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
				"s",
				"",
				"[äöüaeiouy]",
				"(z|s)",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aeiouyäöüj]",
				"[aeiouyäöü]",
				"z",
			},
		},
		{
			patterns: [4]string{
				"ß",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"$",
				"i",
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
				"ue",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: [4]string{
				"ae",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"oe",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"ü",
				"",
				"",
				"Q",
			},
		},
		{
			patterns: [4]string{
				"ä",
				"",
				"",
				"(Y|e)",
			},
		},
		{
			patterns: [4]string{
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"(aj|oj)",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"I",
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
				"ñ",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"ã",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"ő",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ű",
				"",
				"",
				"u",
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
				"a",
				"",
				"",
				"A",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"j",
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
				"O",
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
				"U",
			},
		},
		{
			patterns: [4]string{
				"v",
				"",
				"",
				"(f|v)",
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
				"ts",
			},
		},
	},
	ashhebrew: []rule{
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
				"TB",
			},
		},
	},
	ashhungarian: []rule{
		{
			patterns: [4]string{
				"sz",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"zs",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"cs",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ay",
				"",
				"",
				"(oj|aj)",
			},
		},
		{
			patterns: [4]string{
				"ai",
				"",
				"",
				"(oj|aj)",
			},
		},
		{
			patterns: [4]string{
				"aj",
				"",
				"",
				"(oj|aj)",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[áo]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[áo]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ee",
				"",
				"",
				"(aj|e)",
			},
		},
		{
			patterns: [4]string{
				"ely",
				"",
				"",
				"(aj|eli)",
			},
		},
		{
			patterns: [4]string{
				"ly",
				"",
				"",
				"(j|li)",
			},
		},
		{
			patterns: [4]string{
				"gy",
				"",
				"[aeouáéóúüöőű]",
				"dj",
			},
		},
		{
			patterns: [4]string{
				"gy",
				"",
				"",
				"(d|gi)",
			},
		},
		{
			patterns: [4]string{
				"ny",
				"",
				"[aeouáéóúüöőű]",
				"nj",
			},
		},
		{
			patterns: [4]string{
				"ny",
				"",
				"",
				"(n|ni)",
			},
		},
		{
			patterns: [4]string{
				"ty",
				"",
				"[aeouáéóúüöőű]",
				"tj",
			},
		},
		{
			patterns: [4]string{
				"ty",
				"",
				"",
				"(t|ti)",
			},
		},
		{
			patterns: [4]string{
				"qu",
				"",
				"",
				"(ku|kv)",
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
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"ő",
				"",
				"",
				"Y",
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
				"Q",
			},
		},
		{
			patterns: [4]string{
				"ű",
				"",
				"",
				"Q",
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
				"ts",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"j",
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
				"(S|s)",
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
				"z",
			},
		},
	},
	ashpolish: []rule{
		{
			patterns: [4]string{
				"ska",
				"",
				"$",
				"ski",
			},
		},
		{
			patterns: [4]string{
				"cka",
				"",
				"$",
				"tski",
			},
		},
		{
			patterns: [4]string{
				"lowa",
				"",
				"$",
				"(lova|lof|l|el)",
			},
		},
		{
			patterns: [4]string{
				"kowa",
				"",
				"$",
				"(kova|kof|k|ek)",
			},
		},
		{
			patterns: [4]string{
				"owa",
				"",
				"$",
				"(ova|of|)",
			},
		},
		{
			patterns: [4]string{
				"lowna",
				"",
				"$",
				"(lovna|levna|l|el)",
			},
		},
		{
			patterns: [4]string{
				"kowna",
				"",
				"$",
				"(kovna|k|ek)",
			},
		},
		{
			patterns: [4]string{
				"owna",
				"",
				"$",
				"(ovna|)",
			},
		},
		{
			patterns: [4]string{
				"lówna",
				"",
				"$",
				"(l|el)",
			},
		},
		{
			patterns: [4]string{
				"kówna",
				"",
				"$",
				"(k|ek)",
			},
		},
		{
			patterns: [4]string{
				"ówna",
				"",
				"$",
				"",
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
				"czy",
				"",
				"",
				"tSi",
			},
		},
		{
			patterns: [4]string{
				"cze",
				"",
				"[bcdgkpstwzż]",
				"(tSe|tSF)",
			},
		},
		{
			patterns: [4]string{
				"ciewicz",
				"",
				"",
				"(tsevitS|tSevitS)",
			},
		},
		{
			patterns: [4]string{
				"siewicz",
				"",
				"",
				"(sevitS|SevitS)",
			},
		},
		{
			patterns: [4]string{
				"ziewicz",
				"",
				"",
				"(zevitS|ZevitS)",
			},
		},
		{
			patterns: [4]string{
				"riewicz",
				"",
				"",
				"rjevitS",
			},
		},
		{
			patterns: [4]string{
				"diewicz",
				"",
				"",
				"djevitS",
			},
		},
		{
			patterns: [4]string{
				"tiewicz",
				"",
				"",
				"tjevitS",
			},
		},
		{
			patterns: [4]string{
				"iewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: [4]string{
				"ewicz",
				"",
				"",
				"evitS",
			},
		},
		{
			patterns: [4]string{
				"owicz",
				"",
				"",
				"ovitS",
			},
		},
		{
			patterns: [4]string{
				"icz",
				"",
				"",
				"itS",
			},
		},
		{
			patterns: [4]string{
				"cz",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"cia",
				"",
				"[bcdgkpstwzż]",
				"(tSB|tsB)",
			},
		},
		{
			patterns: [4]string{
				"cia",
				"",
				"",
				"(tSa|tsa)",
			},
		},
		{
			patterns: [4]string{
				"cią",
				"",
				"[bp]",
				"(tSom|tsom)",
			},
		},
		{
			patterns: [4]string{
				"cią",
				"",
				"",
				"(tSon|tson)",
			},
		},
		{
			patterns: [4]string{
				"cię",
				"",
				"[bp]",
				"(tSem|tsem)",
			},
		},
		{
			patterns: [4]string{
				"cię",
				"",
				"",
				"(tSen|tsen)",
			},
		},
		{
			patterns: [4]string{
				"cie",
				"",
				"[bcdgkpstwzż]",
				"(tSF|tsF)",
			},
		},
		{
			patterns: [4]string{
				"cie",
				"",
				"",
				"(tSe|tse)",
			},
		},
		{
			patterns: [4]string{
				"cio",
				"",
				"",
				"(tSo|tso)",
			},
		},
		{
			patterns: [4]string{
				"ciu",
				"",
				"",
				"(tSu|tsu)",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"",
				"(tSi|tsI)",
			},
		},
		{
			patterns: [4]string{
				"ć",
				"",
				"",
				"(tS|ts)",
			},
		},
		{
			patterns: [4]string{
				"ssz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sz",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sia",
				"",
				"[bcdgkpstwzż]",
				"(SB|sB|sja)",
			},
		},
		{
			patterns: [4]string{
				"sia",
				"",
				"",
				"(Sa|sja)",
			},
		},
		{
			patterns: [4]string{
				"sią",
				"",
				"[bp]",
				"(Som|som)",
			},
		},
		{
			patterns: [4]string{
				"sią",
				"",
				"",
				"(Son|son)",
			},
		},
		{
			patterns: [4]string{
				"się",
				"",
				"[bp]",
				"(Sem|sem)",
			},
		},
		{
			patterns: [4]string{
				"się",
				"",
				"",
				"(Sen|sen)",
			},
		},
		{
			patterns: [4]string{
				"sie",
				"",
				"[bcdgkpstwzż]",
				"(SF|sF|se)",
			},
		},
		{
			patterns: [4]string{
				"sie",
				"",
				"",
				"(Se|se)",
			},
		},
		{
			patterns: [4]string{
				"sio",
				"",
				"",
				"(So|so)",
			},
		},
		{
			patterns: [4]string{
				"siu",
				"",
				"",
				"(Su|sju)",
			},
		},
		{
			patterns: [4]string{
				"si",
				"",
				"",
				"(Si|sI)",
			},
		},
		{
			patterns: [4]string{
				"ś",
				"",
				"",
				"(S|s)",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"[bcdgkpstwzż]",
				"(ZB|zB|zja)",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"",
				"(Za|zja)",
			},
		},
		{
			patterns: [4]string{
				"zią",
				"",
				"[bp]",
				"(Zom|zom)",
			},
		},
		{
			patterns: [4]string{
				"zią",
				"",
				"",
				"(Zon|zon)",
			},
		},
		{
			patterns: [4]string{
				"zię",
				"",
				"[bp]",
				"(Zem|zem)",
			},
		},
		{
			patterns: [4]string{
				"zię",
				"",
				"",
				"(Zen|zen)",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"[bcdgkpstwzż]",
				"(ZF|zF)",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"",
				"(Ze|ze)",
			},
		},
		{
			patterns: [4]string{
				"zio",
				"",
				"",
				"(Zo|zo)",
			},
		},
		{
			patterns: [4]string{
				"ziu",
				"",
				"",
				"(Zu|zju)",
			},
		},
		{
			patterns: [4]string{
				"zi",
				"",
				"",
				"(Zi|zI)",
			},
		},
		{
			patterns: [4]string{
				"że",
				"",
				"[bcdgkpstwzż]",
				"(Ze|ZF)",
			},
		},
		{
			patterns: [4]string{
				"że",
				"",
				"[bcdgkpstwzż]",
				"(Ze|ZF|ze|zF)",
			},
		},
		{
			patterns: [4]string{
				"że",
				"",
				"",
				"Ze",
			},
		},
		{
			patterns: [4]string{
				"źe",
				"",
				"",
				"(Ze|ze)",
			},
		},
		{
			patterns: [4]string{
				"ży",
				"",
				"",
				"Zi",
			},
		},
		{
			patterns: [4]string{
				"źi",
				"",
				"",
				"(Zi|zi)",
			},
		},
		{
			patterns: [4]string{
				"ż",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"ź",
				"",
				"",
				"(Z|z)",
			},
		},
		{
			patterns: [4]string{
				"rze",
				"t",
				"",
				"(Se|re)",
			},
		},
		{
			patterns: [4]string{
				"rze",
				"",
				"",
				"(Ze|re|rZe)",
			},
		},
		{
			patterns: [4]string{
				"rzy",
				"t",
				"",
				"(Si|ri)",
			},
		},
		{
			patterns: [4]string{
				"rzy",
				"",
				"",
				"(Zi|ri|rZi)",
			},
		},
		{
			patterns: [4]string{
				"rz",
				"t",
				"",
				"(S|r)",
			},
		},
		{
			patterns: [4]string{
				"rz",
				"",
				"",
				"(Z|r|rZ)",
			},
		},
		{
			patterns: [4]string{
				"lio",
				"",
				"",
				"(lo|le)",
			},
		},
		{
			patterns: [4]string{
				"ł",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"ń",
				"",
				"",
				"n",
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
				"s",
				"",
				"s",
				"",
			},
		},
		{
			patterns: [4]string{
				"ó",
				"",
				"",
				"(u|o)",
			},
		},
		{
			patterns: [4]string{
				"ą",
				"",
				"[bp]",
				"om",
			},
		},
		{
			patterns: [4]string{
				"ę",
				"",
				"[bp]",
				"em",
			},
		},
		{
			patterns: [4]string{
				"ą",
				"",
				"",
				"on",
			},
		},
		{
			patterns: [4]string{
				"ę",
				"",
				"",
				"en",
			},
		},
		{
			patterns: [4]string{
				"ije",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"yje",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"iie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"yie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"iye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"yye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yj",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"iy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"rie",
				"",
				"",
				"rje",
			},
		},
		{
			patterns: [4]string{
				"die",
				"",
				"",
				"dje",
			},
		},
		{
			patterns: [4]string{
				"tie",
				"",
				"",
				"tje",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"[bcdgkpstwzż]",
				"F",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"e",
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
				"au",
				"",
				"",
				"au",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ej",
				"",
				"",
				"aj",
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
				"aj",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[ou]",
				"",
				"j",
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
				"[aeou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"[bcdgkpstwzż]",
				"B",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"[bcdgkpstwzż]",
				"(E|F)",
			},
		},
		{
			patterns: [4]string{
				"o",
				"",
				"[bcćdgklłmnńrsśtwzźż]",
				"P",
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
				"ts",
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
				"E",
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
				"(h|x)",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"",
				"I",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"j",
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
				"I",
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
	ashromanian: []rule{
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
				"ce",
				"",
				"",
				"tSe",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"",
				"(tSi|tS)",
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
				"ch",
				"",
				"",
				"x",
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
				"gi",
				"",
				"",
				"(dZi|dZ)",
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
				"gh",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aou]",
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
				"ţ",
				"",
				"",
				"ts",
			},
		},
		{
			patterns: [4]string{
				"ş",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"h",
				"",
				"",
				"(x|h)",
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
				"î",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ea",
				"",
				"",
				"ja",
			},
		},
		{
			patterns: [4]string{
				"ă",
				"",
				"",
				"(e|a)",
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
				"E",
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
				"i",
				"",
				"",
				"I",
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
	ashrussian: []rule{
		{
			patterns: [4]string{
				"yna",
				"",
				"$",
				"(in|ina)",
			},
		},
		{
			patterns: [4]string{
				"ina",
				"",
				"$",
				"(in|ina)",
			},
		},
		{
			patterns: [4]string{
				"liova",
				"",
				"$",
				"(lof|lef)",
			},
		},
		{
			patterns: [4]string{
				"lova",
				"",
				"$",
				"(lof|lef|lova)",
			},
		},
		{
			patterns: [4]string{
				"ova",
				"",
				"$",
				"(of|ova)",
			},
		},
		{
			patterns: [4]string{
				"eva",
				"",
				"$",
				"(ef|ova)",
			},
		},
		{
			patterns: [4]string{
				"aia",
				"",
				"$",
				"(aja|i)",
			},
		},
		{
			patterns: [4]string{
				"aja",
				"",
				"$",
				"(aja|i)",
			},
		},
		{
			patterns: [4]string{
				"aya",
				"",
				"$",
				"(aja|i)",
			},
		},
		{
			patterns: [4]string{
				"tsya",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: [4]string{
				"tsyu",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: [4]string{
				"tsia",
				"",
				"",
				"tsa",
			},
		},
		{
			patterns: [4]string{
				"tsie",
				"",
				"",
				"tse",
			},
		},
		{
			patterns: [4]string{
				"tsio",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: [4]string{
				"tsye",
				"",
				"",
				"tse",
			},
		},
		{
			patterns: [4]string{
				"tsyo",
				"",
				"",
				"tso",
			},
		},
		{
			patterns: [4]string{
				"tsiu",
				"",
				"",
				"tsu",
			},
		},
		{
			patterns: [4]string{
				"sie",
				"",
				"",
				"se",
			},
		},
		{
			patterns: [4]string{
				"sio",
				"",
				"",
				"so",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"",
				"ze",
			},
		},
		{
			patterns: [4]string{
				"zio",
				"",
				"",
				"zo",
			},
		},
		{
			patterns: [4]string{
				"sye",
				"",
				"",
				"se",
			},
		},
		{
			patterns: [4]string{
				"syo",
				"",
				"",
				"so",
			},
		},
		{
			patterns: [4]string{
				"zye",
				"",
				"",
				"ze",
			},
		},
		{
			patterns: [4]string{
				"zyo",
				"",
				"",
				"zo",
			},
		},
		{
			patterns: [4]string{
				"gauz",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: [4]string{
				"gaus",
				"",
				"$",
				"haus",
			},
		},
		{
			patterns: [4]string{
				"gol'ts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"golts",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"gol'tz",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"goltz",
				"",
				"$",
				"holts",
			},
		},
		{
			patterns: [4]string{
				"gejmer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"gejm",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"geimer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"geim",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"geymer",
				"",
				"$",
				"hajmer",
			},
		},
		{
			patterns: [4]string{
				"geym",
				"",
				"$",
				"hajm",
			},
		},
		{
			patterns: [4]string{
				"gendler",
				"",
				"$",
				"hendler",
			},
		},
		{
			patterns: [4]string{
				"gof",
				"",
				"$",
				"hof",
			},
		},
		{
			patterns: [4]string{
				"gojf",
				"",
				"$",
				"hojf",
			},
		},
		{
			patterns: [4]string{
				"goyf",
				"",
				"$",
				"hojf",
			},
		},
		{
			patterns: [4]string{
				"goif",
				"",
				"$",
				"hojf",
			},
		},
		{
			patterns: [4]string{
				"ger",
				"",
				"$",
				"ger",
			},
		},
		{
			patterns: [4]string{
				"gen",
				"",
				"$",
				"gen",
			},
		},
		{
			patterns: [4]string{
				"gin",
				"",
				"$",
				"gin",
			},
		},
		{
			patterns: [4]string{
				"gg",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"kog",
				"^",
				"[aeoiu]",
				"(kog|koh)",
			},
		},
		{
			patterns: [4]string{
				"kag",
				"^",
				"[aeoiu]",
				"(kag|kah)",
			},
		},
		{
			patterns: [4]string{
				"g",
				"[jaeoiuy]",
				"[aeoiu]",
				"g",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[aeoiu]",
				"(g|h)",
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
				"ch",
				"",
				"",
				"(tS|x)",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"",
				"(StS|S)",
			},
		},
		{
			patterns: [4]string{
				"ssh",
				"",
				"",
				"S",
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
				"zh",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"$",
				"ts",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"",
				"(ts|tz)",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[iey]",
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
				"qu",
				"",
				"",
				"(kv|k)",
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
				"",
				"s",
				"",
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
				"lya",
				"",
				"",
				"la",
			},
		},
		{
			patterns: [4]string{
				"lyu",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: [4]string{
				"lia",
				"",
				"",
				"la",
			},
		},
		{
			patterns: [4]string{
				"liu",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: [4]string{
				"lja",
				"",
				"",
				"la",
			},
		},
		{
			patterns: [4]string{
				"lju",
				"",
				"",
				"lu",
			},
		},
		{
			patterns: [4]string{
				"le",
				"",
				"",
				"(lo|lE)",
			},
		},
		{
			patterns: [4]string{
				"lyo",
				"",
				"",
				"(lo|le)",
			},
		},
		{
			patterns: [4]string{
				"lio",
				"",
				"",
				"(lo|le)",
			},
		},
		{
			patterns: [4]string{
				"ije",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"iye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"iie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"yje",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"yye",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"yie",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"iy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yj",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"",
				"[aou]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"io",
				"",
				"",
				"(jo|e)",
			},
		},
		{
			patterns: [4]string{
				"i",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[aou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ej",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"yo",
				"",
				"",
				"(jo|e)",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"[au]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[aiou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"iy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"",
				" ",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yj",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"e",
				"^",
				"",
				"(je|E)",
			},
		},
		{
			patterns: [4]string{
				"ee",
				"",
				"",
				"(aje|i)",
			},
		},
		{
			patterns: [4]string{
				"e",
				"[aou]",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"y",
				"",
				"",
				"I",
			},
		},
		{
			patterns: [4]string{
				"oo",
				"",
				"",
				"(oo|u)",
			},
		},
		{
			patterns: [4]string{
				"'",
				"",
				"",
				"",
			},
		},
		{
			patterns: [4]string{
				"\"",
				"",
				"",
				"",
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
				"E",
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
				"I",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"",
				"j",
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
	ashspanish: []rule{
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
				"ch",
				"",
				"",
				"(tS|dZ)",
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
				"x",
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
				"ll",
				"",
				"",
				"(l|Z)",
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
				"(x|g)",
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
				"(i|j|S|Z)",
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
				"E",
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
				"I",
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
					"[gdZz]",
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
					"[bgdZz]",
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
					"[bdZz]",
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
					"[bgZz]",
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
					"jnm",
					"",
					"",
					"jm",
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
					"jI",
					"^",
					"",
					"I",
				},
			},
			{
				patterns: [4]string{
					"a",
					"",
					"[aAB]",
					"",
				},
			},
			{
				patterns: [4]string{
					"a",
					"[AB]",
					"",
					"",
				},
			},
			{
				patterns: [4]string{
					"A",
					"",
					"A",
					"",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"B",
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
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: [4]string{
					"kAg",
					"^",
					"[AEOIUaeoiu]",
					"(kOg|kO[512])",
				},
			},
			{
				patterns: [4]string{
					"kOg",
					"^",
					"[AEOIUaeoiu]",
					"(kAg|kA[512])",
				},
			},
			{
				patterns: [4]string{
					"kog",
					"^",
					"[AEOIUaeoiu]",
					"(kog|ko[512])",
				},
			},
			{
				patterns: [4]string{
					"kag",
					"^",
					"[AEOIUaeoiu]",
					"(kag|ka[512])",
				},
			},
			{
				patterns: [4]string{
					"h",
					"",
					"",
					"",
				},
			},
			{
				patterns: [4]string{
					"H",
					"",
					"",
					"(x|)",
				},
			},
			{
				patterns: [4]string{
					"F",
					"",
					"[bdgkpstvzZ]h",
					"e",
				},
			},
			{
				patterns: [4]string{
					"F",
					"",
					"[bdgkpstvzZ]x",
					"e",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"[bdgkpstvzZ]h",
					"a",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"[bdgkpstvzZ]x",
					"a",
				},
			},
			{
				patterns: [4]string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(e|)",
				},
			},
			{
				patterns: [4]string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(i|)",
				},
			},
			{
				patterns: [4]string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(E|)",
				},
			},
			{
				patterns: [4]string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(I|)",
				},
			},
			{
				patterns: [4]string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(F|)",
				},
			},
			{
				patterns: [4]string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Q|)",
				},
			},
			{
				patterns: [4]string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Y|)",
				},
			},
			{
				patterns: [4]string{
					"lEs",
					"",
					"",
					"(lEs|lz)",
				},
			},
			{
				patterns: [4]string{
					"lE",
					"[bdfgkmnprStvzZ]",
					"",
					"(lE|l)",
				},
			},
			{
				patterns: [4]string{
					"aue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AvE",
					"",
					"",
					"(D|AvE)",
				},
			},
			{
				patterns: [4]string{
					"Ave",
					"",
					"",
					"(D|Ave)",
				},
			},
			{
				patterns: [4]string{
					"avE",
					"",
					"",
					"(D|avE)",
				},
			},
			{
				patterns: [4]string{
					"ave",
					"",
					"",
					"(D|ave)",
				},
			},
			{
				patterns: [4]string{
					"OvE",
					"",
					"",
					"(D|OvE)",
				},
			},
			{
				patterns: [4]string{
					"Ove",
					"",
					"",
					"(D|Ove)",
				},
			},
			{
				patterns: [4]string{
					"ovE",
					"",
					"",
					"(D|ovE)",
				},
			},
			{
				patterns: [4]string{
					"ove",
					"",
					"",
					"(D|ove)",
				},
			},
			{
				patterns: [4]string{
					"ea",
					"",
					"",
					"(D|ea)",
				},
			},
			{
				patterns: [4]string{
					"EA",
					"",
					"",
					"(D|EA)",
				},
			},
			{
				patterns: [4]string{
					"Ea",
					"",
					"",
					"(D|Ea)",
				},
			},
			{
				patterns: [4]string{
					"eA",
					"",
					"",
					"(D|eA)",
				},
			},
			{
				patterns: [4]string{
					"aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
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
					"lYndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lander",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lAndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lAnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"landEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lender",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lEndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lendEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lEnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"bUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"bUrg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"burg",
					"",
					"$",
					"(burk|berk)",
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
		},
		second: []secondFinalRule{
			{
				langs: 0,
				rules: []rule{
					{
						patterns: [4]string{
							"b",
							"",
							"",
							"(b|v[1024])",
						},
					},
					{
						patterns: [4]string{
							"J",
							"",
							"",
							"z",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"AiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"OiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"UiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"AiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"OiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"UiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[bp]",
							"(o|om[128]|im[128])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[dgkstvz]",
							"(a|o|on[128]|in[128])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"AiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"OiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"UiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"AiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"OiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"UiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[bp]",
							"(i|im[128]|om[128])",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in[128]|on[128])",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiouAEIBFOUQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^aeiouAEBFIOU]e",
							"(Q[16]|i|D[4])",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk[16])",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts[16])",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(Q[16]|i)",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"$",
							"(li|il[4])",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[4]|lY[16])",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: [4]string{
							"Ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: [4]string{
							"iA",
							"",
							"$",
							"(ia|io)",
						},
					},
					{
						patterns: [4]string{
							"iA",
							"",
							"",
							"(ia|io|iY[16])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[16]|D[4])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"[fklmnprstv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"[fklmnprstv]$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"[oeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"",
							"(o|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"[oeiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"",
							"(a|o|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"[DoiuQY]",
							"",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk[16])",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: [4]string{
							"sUts",
							"",
							"$",
							"(suts|sQts[16])",
						},
					},
					{
						patterns: [4]string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: [4]string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: [4]string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: [4]string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: [4]string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: [4]string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: [4]string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: [4]string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: [4]string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: [4]string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: [4]string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"[^aEIeiou]e",
							"(Q|i|D)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(il|li|lY)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"E",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiEuQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiAEIOUouQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(Q|i)",
						},
					},
					{
						patterns: [4]string{
							"AU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"aU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"Au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"OU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"oU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"Ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoAOUiuQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"[fklmnprst]$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"[aoAOUeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"",
							"(o|Y)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"[aoeOUiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"",
							"(a|o|Y)",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"[DaoiuUQY]",
							"",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: [4]string{
							"sUts",
							"",
							"$",
							"(suts|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
					"[gdZz]",
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
					"[bgdZz]",
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
					"[bdZz]",
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
					"[bgZz]",
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
					"jnm",
					"",
					"",
					"jm",
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
					"jI",
					"^",
					"",
					"I",
				},
			},
			{
				patterns: [4]string{
					"a",
					"",
					"[aAB]",
					"",
				},
			},
			{
				patterns: [4]string{
					"a",
					"[AB]",
					"",
					"",
				},
			},
			{
				patterns: [4]string{
					"A",
					"",
					"A",
					"",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"B",
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
					"n",
					"",
					"[bp]",
					"m",
				},
			},
			{
				patterns: [4]string{
					"kAg",
					"^",
					"[AEOIUaeoiu]",
					"(kOg|kO[512])",
				},
			},
			{
				patterns: [4]string{
					"kOg",
					"^",
					"[AEOIUaeoiu]",
					"(kAg|kA[512])",
				},
			},
			{
				patterns: [4]string{
					"kog",
					"^",
					"[AEOIUaeoiu]",
					"(kog|ko[512])",
				},
			},
			{
				patterns: [4]string{
					"kag",
					"^",
					"[AEOIUaeoiu]",
					"(kag|ka[512])",
				},
			},
			{
				patterns: [4]string{
					"h",
					"",
					"",
					"",
				},
			},
			{
				patterns: [4]string{
					"H",
					"",
					"",
					"(x|)",
				},
			},
			{
				patterns: [4]string{
					"F",
					"",
					"[bdgkpstvzZ]h",
					"e",
				},
			},
			{
				patterns: [4]string{
					"F",
					"",
					"[bdgkpstvzZ]x",
					"e",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"[bdgkpstvzZ]h",
					"a",
				},
			},
			{
				patterns: [4]string{
					"B",
					"",
					"[bdgkpstvzZ]x",
					"a",
				},
			},
			{
				patterns: [4]string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"e",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(e|)",
				},
			},
			{
				patterns: [4]string{
					"i",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(i|)",
				},
			},
			{
				patterns: [4]string{
					"E",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(E|)",
				},
			},
			{
				patterns: [4]string{
					"I",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(I|)",
				},
			},
			{
				patterns: [4]string{
					"F",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(F|)",
				},
			},
			{
				patterns: [4]string{
					"Q",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Q|)",
				},
			},
			{
				patterns: [4]string{
					"Y",
					"[bdfgklmnprsStvzZ]",
					"[ln][bdfgklmnprsStvzZ]",
					"(Y|)",
				},
			},
			{
				patterns: [4]string{
					"lEs",
					"",
					"",
					"(lEs|lz)",
				},
			},
			{
				patterns: [4]string{
					"lE",
					"[bdfgkmnprStvzZ]",
					"",
					"(lE|l)",
				},
			},
			{
				patterns: [4]string{
					"aue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oue",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AvE",
					"",
					"",
					"(D|AvE)",
				},
			},
			{
				patterns: [4]string{
					"Ave",
					"",
					"",
					"(D|Ave)",
				},
			},
			{
				patterns: [4]string{
					"avE",
					"",
					"",
					"(D|avE)",
				},
			},
			{
				patterns: [4]string{
					"ave",
					"",
					"",
					"(D|ave)",
				},
			},
			{
				patterns: [4]string{
					"OvE",
					"",
					"",
					"(D|OvE)",
				},
			},
			{
				patterns: [4]string{
					"Ove",
					"",
					"",
					"(D|Ove)",
				},
			},
			{
				patterns: [4]string{
					"ovE",
					"",
					"",
					"(D|ovE)",
				},
			},
			{
				patterns: [4]string{
					"ove",
					"",
					"",
					"(D|ove)",
				},
			},
			{
				patterns: [4]string{
					"ea",
					"",
					"",
					"(D|ea)",
				},
			},
			{
				patterns: [4]string{
					"EA",
					"",
					"",
					"(D|EA)",
				},
			},
			{
				patterns: [4]string{
					"Ea",
					"",
					"",
					"(D|Ea)",
				},
			},
			{
				patterns: [4]string{
					"eA",
					"",
					"",
					"(D|eA)",
				},
			},
			{
				patterns: [4]string{
					"aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Oji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Oje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Eji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Eje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Uji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Uje",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Iji",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjI",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ije",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjE",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ajU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ajo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ojO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Oja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ojo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"OjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ejO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Eja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ejo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"EjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ujO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Uja",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ujo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"UjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"ijO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ija",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjA",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Ijo",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"IjO",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"Aju",
					"",
					"",
					"D",
				},
			},
			{
				patterns: [4]string{
					"AjU",
					"",
					"",
					"D",
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
					"lYndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lander",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lAndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lAnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"landEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lender",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lEndEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lendEr",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"lEnder",
					"",
					"$",
					"lYnder",
				},
			},
			{
				patterns: [4]string{
					"bUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"bUrg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"burg",
					"",
					"$",
					"(burk|berk)",
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
		},
		second: []secondFinalRule{
			{
				langs: 0,
				rules: []rule{
					{
						patterns: [4]string{
							"b",
							"",
							"",
							"(b|v[1024])",
						},
					},
					{
						patterns: [4]string{
							"J",
							"",
							"",
							"z",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"AiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"OiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"UiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"AiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"OiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"UiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[bp]",
							"(o|om[128]|im[128])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[dgkstvz]",
							"(a|o|on[128]|in[128])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"AiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"OiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"UiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"AiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"OiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"UiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[bp]",
							"(i|im[128]|om[128])",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in[128]|on[128])",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiouAEIBFOUQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^aeiouAEBFIOU]e",
							"(Q[16]|i|D[4])",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk[16])",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts[16])",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(Q[16]|i)",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"$",
							"(li|il[4])",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[4]|lY[16])",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: [4]string{
							"Ei",
							"",
							"",
							"(D|i)",
						},
					},
					{
						patterns: [4]string{
							"iA",
							"",
							"$",
							"(ia|io)",
						},
					},
					{
						patterns: [4]string{
							"iA",
							"",
							"",
							"(ia|io|iY[16])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[16]|D[4])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[16]|[4])",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"[fklmnprstv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuAOIUQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"",
							"(i|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"[fklmnprstv]$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"[oeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"",
							"(o|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"[oeiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"",
							"(a|o|Y[16])",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"[DoiuQY]",
							"",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk[16])",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: [4]string{
							"sUts",
							"",
							"$",
							"(suts|sQts[16])",
						},
					},
					{
						patterns: [4]string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: [4]string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: [4]string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: [4]string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: [4]string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: [4]string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiEIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"om",
							"",
							"[bp]",
							"(om|im)",
						},
					},
					{
						patterns: [4]string{
							"on",
							"",
							"[dgkstvz]",
							"(on|in)",
						},
					},
					{
						patterns: [4]string{
							"em",
							"",
							"[bp]",
							"(im|om)",
						},
					},
					{
						patterns: [4]string{
							"en",
							"",
							"[dgkstvz]",
							"(in|on)",
						},
					},
					{
						patterns: [4]string{
							"Em",
							"",
							"[bp]",
							"(im|Ym|om)",
						},
					},
					{
						patterns: [4]string{
							"En",
							"",
							"[dgkstvz]",
							"(in|Yn|on)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"[^aEIeiou]e",
							"(Q|i|D)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(il|li|lY)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"E",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"D[^aeiEIou]",
							"",
							"(i|)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiEuQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiAEIOUouQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(Q|i)",
						},
					},
					{
						patterns: [4]string{
							"AU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"aU",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"Au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"OU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"oU",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"Ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"Ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"Oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"Ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoAOUiuQY]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoAOQY]",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"",
							"(Y|i)",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"[fklmnprst]$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"ts$",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"[aoAOUeiuQY]",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"",
							"(o|Y)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[fklmnprst]$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"ts$",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"[aoeOUiuQY]",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"",
							"(a|o|Y)",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"[DaoiuUQY]",
							"",
							"u",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"[^k]$",
							"u",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"[lr]",
							"$",
							"(uk|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Uk",
							"",
							"$",
							"uk",
						},
					},
					{
						patterns: [4]string{
							"sUts",
							"",
							"$",
							"(suts|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Uts",
							"",
							"$",
							"uts",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"aiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiB",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[bp]",
							"(o|om|im)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[dgkstvz]",
							"(o|on|in)",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"o",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[bp]",
							"(D|Dm)",
						},
					},
					{
						patterns: [4]string{
							"aiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"oiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"uiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"eiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"EiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"iiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"IiF",
							"",
							"[dgkstvz]",
							"(D|Dn)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[bp]",
							"(i|im|om)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"[dgkstvz]",
							"(i|in|on)",
						},
					},
					{
						patterns: [4]string{
							"F",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"P",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aeiAEBFIou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprst]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[DaoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
						patterns: [4]string{
							"I",
							"",
							"$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"[aEIeiou]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"[^k]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"[lr]",
							"$",
							"(ik|Qk)",
						},
					},
					{
						patterns: [4]string{
							"Ik",
							"",
							"$",
							"ik",
						},
					},
					{
						patterns: [4]string{
							"sIts",
							"",
							"$",
							"(sits|sQts)",
						},
					},
					{
						patterns: [4]string{
							"Its",
							"",
							"$",
							"its",
						},
					},
					{
						patterns: [4]string{
							"I",
							"",
							"",
							"(i|Q)",
						},
					},
					{
						patterns: [4]string{
							"au",
							"",
							"",
							"(D|a|u)",
						},
					},
					{
						patterns: [4]string{
							"ou",
							"",
							"",
							"(D|o|u)",
						},
					},
					{
						patterns: [4]string{
							"ai",
							"",
							"",
							"(D|a|i)",
						},
					},
					{
						patterns: [4]string{
							"oi",
							"",
							"",
							"(D|o|i)",
						},
					},
					{
						patterns: [4]string{
							"ui",
							"",
							"",
							"(D|u|i)",
						},
					},
					{
						patterns: [4]string{
							"a",
							"",
							"",
							"(a|o)",
						},
					},
					{
						patterns: [4]string{
							"e",
							"",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[fklmnprsStv]$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"ts$",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"[aoiuQ]",
							"",
							"i",
						},
					},
					{
						patterns: [4]string{
							"E",
							"",
							"[aoQ]",
							"i",
						},
					},
					{
						patterns: [4]string{
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
