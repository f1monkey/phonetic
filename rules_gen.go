// GENERATED CODE. DO NOT EDIT!
package bmpm

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
			patterns: [4]string{
				"yna",
				"",
				"$",
				"(in[131072]|ina)",
			},
		},
		{
			patterns: [4]string{
				"ina",
				"",
				"$",
				"(in[131072]|ina)",
			},
		},
		{
			patterns: [4]string{
				"liova",
				"",
				"$",
				"(lova|lof[131072]|lef[131072])",
			},
		},
		{
			patterns: [4]string{
				"lova",
				"",
				"$",
				"(lova|lof[131072]|lef[131072]|l[8]|el[8])",
			},
		},
		{
			patterns: [4]string{
				"kova",
				"",
				"$",
				"(kova|kof[131072]|k[8]|ek[8])",
			},
		},
		{
			patterns: [4]string{
				"ova",
				"",
				"$",
				"(ova|of[131072]|[8])",
			},
		},
		{
			patterns: [4]string{
				"ová",
				"",
				"$",
				"(ova|[8])",
			},
		},
		{
			patterns: [4]string{
				"eva",
				"",
				"$",
				"(eva|ef[131072])",
			},
		},
		{
			patterns: [4]string{
				"aia",
				"",
				"$",
				"(aja|i[131072])",
			},
		},
		{
			patterns: [4]string{
				"aja",
				"",
				"$",
				"(aja|i[131072])",
			},
		},
		{
			patterns: [4]string{
				"aya",
				"",
				"$",
				"(aja|i[131072])",
			},
		},
		{
			patterns: [4]string{
				"lowa",
				"",
				"$",
				"(lova|lof[16384]|l[16384]|el[16384])",
			},
		},
		{
			patterns: [4]string{
				"kowa",
				"",
				"$",
				"(kova|kof[16384]|k[16384]|ek[16384])",
			},
		},
		{
			patterns: [4]string{
				"owa",
				"",
				"$",
				"(ova|of[16384]|)",
			},
		},
		{
			patterns: [4]string{
				"lowna",
				"",
				"$",
				"(lovna|levna|l[16384]|el[16384])",
			},
		},
		{
			patterns: [4]string{
				"kowna",
				"",
				"$",
				"(kovna|k[16384]|ek[16384])",
			},
		},
		{
			patterns: [4]string{
				"owna",
				"",
				"$",
				"(ovna|[16384])",
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
				"á",
				"",
				"$",
				"(a|i[8])",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"$",
				"(a|i[16392])",
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
				"que",
				"",
				"$",
				"(k[64]|ke|kve)",
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
				"m",
				"",
				"[bfpv]",
				"(m|n)",
			},
		},
		{
			patterns: [4]string{
				"m",
				"[aeiouy]",
				"[aeiouy]",
				"m",
			},
		},
		{
			patterns: [4]string{
				"m",
				"[aeiouy]",
				"",
				"(m|n[32832])",
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
				"lio",
				"",
				"",
				"(lo|le[131072])",
			},
		},
		{
			patterns: [4]string{
				"lyo",
				"",
				"",
				"(lo|le[131072])",
			},
		},
		{
			patterns: [4]string{
				"lt",
				"u",
				"$",
				"(lt|[64])",
			},
		},
		{
			patterns: [4]string{
				"v",
				"^",
				"",
				"(v|f[128]|b[262144])",
			},
		},
		{
			patterns: [4]string{
				"ex",
				"",
				"[aáuiíoóeéêy]",
				"(ez[32768]|eS[32768]|eks|egz)",
			},
		},
		{
			patterns: [4]string{
				"ex",
				"",
				"[cs]",
				"(e[32768]|ek)",
			},
		},
		{
			patterns: [4]string{
				"x",
				"u",
				"$",
				"(ks|[64])",
			},
		},
		{
			patterns: [4]string{
				"ck",
				"",
				"",
				"(k|tsk[16392])",
			},
		},
		{
			patterns: [4]string{
				"cz",
				"",
				"",
				"(tS|tsz[8])",
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
				"dh",
				"^",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"bh",
				"^",
				"",
				"b",
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
				"kh",
				"",
				"",
				"(x[131104]|kh)",
			},
		},
		{
			patterns: [4]string{
				"lh",
				"",
				"",
				"(lh|l[32768])",
			},
		},
		{
			patterns: [4]string{
				"nh",
				"",
				"",
				"(nh|nj[32768])",
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
				"[aeiouy]",
				"[ei]",
				"(S|StS[131072]|sk[69632])",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"[aeiouy]",
				"",
				"(S|StS[131072])",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"[ei]",
				"(sk[69632]|S|StS[131072])",
			},
		},
		{
			patterns: [4]string{
				"sch",
				"",
				"",
				"(S|StS[131072])",
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
				"(S[131104]|sh)",
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
				"(Z[131104]|zh|tsh[128])",
			},
		},
		{
			patterns: [4]string{
				"chs",
				"",
				"",
				"(ks[128]|xs|tSs[131104])",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"[ei]",
				"(x|tS[393248]|k[69632]|S[32832])",
			},
		},
		{
			patterns: [4]string{
				"ch",
				"",
				"",
				"(x|tS[393248]|S[32832])",
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
				"(t[672]|th)",
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
				"gh",
				"",
				"[ei]",
				"(g[70144]|gh)",
			},
		},
		{
			patterns: [4]string{
				"ouh",
				"",
				"[aioe]",
				"(v[64]|uh)",
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
				"h",
				"",
				"$",
				"",
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
				"(h|x[66048]|H[381024])",
			},
		},
		{
			patterns: [4]string{
				"cia",
				"",
				"",
				"(tSa[16384]|tsa)",
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
				"(tSon[16384]|tson)",
			},
		},
		{
			patterns: [4]string{
				"cię",
				"",
				"[bp]",
				"(tSem[16384]|tsem)",
			},
		},
		{
			patterns: [4]string{
				"cię",
				"",
				"",
				"(tSen[16384]|tsen)",
			},
		},
		{
			patterns: [4]string{
				"cie",
				"",
				"",
				"(tSe[16384]|tse)",
			},
		},
		{
			patterns: [4]string{
				"cio",
				"",
				"",
				"(tSo[16384]|tso)",
			},
		},
		{
			patterns: [4]string{
				"ciu",
				"",
				"",
				"(tSu[16384]|tsu)",
			},
		},
		{
			patterns: [4]string{
				"sci",
				"",
				"$",
				"(Si[4096]|stsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
			},
		},
		{
			patterns: [4]string{
				"sc",
				"",
				"[ei]",
				"(S[4096]|sts[16392]|dZ[524288]|tS[81920]|s)",
			},
		},
		{
			patterns: [4]string{
				"ci",
				"",
				"$",
				"(tsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
			},
		},
		{
			patterns: [4]string{
				"cy",
				"",
				"",
				"(si|tsi[16384])",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"[ei]",
				"(ts[16392]|dZ[524288]|tS[81920]|k[512]|s)",
			},
		},
		{
			patterns: [4]string{
				"sç",
				"",
				"[aeiou]",
				"(s|stS[524288])",
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
				"^",
				"",
				"(S|s[2048])",
			},
		},
		{
			patterns: [4]string{
				"sz",
				"",
				"$",
				"(S|s[2048])",
			},
		},
		{
			patterns: [4]string{
				"sz",
				"",
				"",
				"(S|s[2048]|sts[128])",
			},
		},
		{
			patterns: [4]string{
				"ssp",
				"",
				"",
				"(Sp[128]|sp)",
			},
		},
		{
			patterns: [4]string{
				"sp",
				"",
				"",
				"(Sp[128]|sp)",
			},
		},
		{
			patterns: [4]string{
				"sst",
				"",
				"",
				"(St[128]|st)",
			},
		},
		{
			patterns: [4]string{
				"st",
				"",
				"",
				"(St[128]|st)",
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
				"sj",
				"^",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sj",
				"",
				"$",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sj",
				"",
				"",
				"(sj|S[16]|sx[262144]|sZ[589824])",
			},
		},
		{
			patterns: [4]string{
				"sia",
				"",
				"",
				"(Sa[16384]|sa[16384]|sja)",
			},
		},
		{
			patterns: [4]string{
				"sią",
				"",
				"[bp]",
				"(Som[16384]|som)",
			},
		},
		{
			patterns: [4]string{
				"sią",
				"",
				"",
				"(Son[16384]|son)",
			},
		},
		{
			patterns: [4]string{
				"się",
				"",
				"[bp]",
				"(Sem[16384]|sem)",
			},
		},
		{
			patterns: [4]string{
				"się",
				"",
				"",
				"(Sen[16384]|sen)",
			},
		},
		{
			patterns: [4]string{
				"sie",
				"",
				"",
				"(se|sje|Se[16384]|zi[128])",
			},
		},
		{
			patterns: [4]string{
				"sio",
				"",
				"",
				"(So[16384]|so)",
			},
		},
		{
			patterns: [4]string{
				"siu",
				"",
				"",
				"(Su[16384]|sju)",
			},
		},
		{
			patterns: [4]string{
				"si",
				"[äöëaáuiíoóeéêy]",
				"",
				"(Si[16384]|si|zi[37056])",
			},
		},
		{
			patterns: [4]string{
				"si",
				"",
				"",
				"(Si[16384]|si|zi[128])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aáuiíoóeéêy]",
				"[aáuíoóeéêy]",
				"(s|z[37056])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"[aeouäöë]",
				"(s|z[128])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aeiouy]",
				"[dglmnrv]",
				"(s|z|Z[32768]|[64])",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"[dglmnrv]",
				"(s|z|Z[32768])",
			},
		},
		{
			patterns: [4]string{
				"gue",
				"",
				"$",
				"(k[64]|gve)",
			},
		},
		{
			patterns: [4]string{
				"gu",
				"",
				"[ei]",
				"(g[64]|gv[294912])",
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
				"guy",
				"",
				"",
				"gi",
			},
		},
		{
			patterns: [4]string{
				"gli",
				"",
				"",
				"(glI|l[4096])",
			},
		},
		{
			patterns: [4]string{
				"gni",
				"",
				"",
				"(gnI|ni[4160])",
			},
		},
		{
			patterns: [4]string{
				"gn",
				"",
				"[aeou]",
				"(n[4160]|nj[4160]|gn)",
			},
		},
		{
			patterns: [4]string{
				"ggie",
				"",
				"",
				"(je[512]|dZe)",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"",
				"[aou]",
				"(j[512]|dZ)",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"[yaeiou]",
				"[aou]",
				"(gI|dZ[4096]|j[512])",
			},
		},
		{
			patterns: [4]string{
				"gge",
				"[yaeiou]",
				"",
				"(gE|xe[262144]|gZe[32832]|dZe[331808]|je[512])",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"[yaeiou]",
				"",
				"(gI|xi[262144]|gZi[32832]|dZi[331808]|i[512])",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"",
				"[aou]",
				"(gI|dZ[4096]|j[512])",
			},
		},
		{
			patterns: [4]string{
				"gie",
				"",
				"$",
				"(ge|gi[128]|ji[64]|dZe[4096])",
			},
		},
		{
			patterns: [4]string{
				"gie",
				"",
				"",
				"(ge|gi[128]|dZe[4096]|je[512])",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"[aou]",
				"(i[512]|dZ)",
			},
		},
		{
			patterns: [4]string{
				"ge",
				"[yaeiou]",
				"",
				"(gE|xe[262144]|Ze[32832]|dZe[331808])",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"[yaeiou]",
				"",
				"(gI|xi[262144]|Zi[32832]|dZi[331808])",
			},
		},
		{
			patterns: [4]string{
				"ge",
				"",
				"",
				"(gE|xe[262144]|hE[131072]|je[512]|Ze[32832]|dZe[331808])",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"",
				"(gI|xi[262144]|hI[131072]|i[512]|Zi[32832]|dZi[331808])",
			},
		},
		{
			patterns: [4]string{
				"gy",
				"",
				"[aeouáéóúüöőű]",
				"(gi|dj[2048])",
			},
		},
		{
			patterns: [4]string{
				"gy",
				"",
				"",
				"(gi|d[2048])",
			},
		},
		{
			patterns: [4]string{
				"g",
				"[yaeiou]",
				"[aouyei]",
				"g",
			},
		},
		{
			patterns: [4]string{
				"g",
				"",
				"[aouei]",
				"(g|h[131072])",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"",
				"(i|ej[16]|ix[262144]|iZ[622656])",
			},
		},
		{
			patterns: [4]string{
				"j",
				"",
				"[aoeiuy]",
				"(j|dZ[32]|x[262144]|Z[622656])",
			},
		},
		{
			patterns: [4]string{
				"rz",
				"t",
				"",
				"(S[16384]|r)",
			},
		},
		{
			patterns: [4]string{
				"rz",
				"",
				"",
				"(rz|rts[128]|Z[16384]|r[16384]|rZ[16384])",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"$",
				"(ts|tS[160])",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"^",
				"",
				"(ts[131232]|tS[160])",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"",
				"(ts[131232]|tz)",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"[bcdgkpstwzż]",
				"(Za[16384]|za[16384]|zja)",
			},
		},
		{
			patterns: [4]string{
				"zia",
				"",
				"",
				"(Za[16384]|zja)",
			},
		},
		{
			patterns: [4]string{
				"zią",
				"",
				"[bp]",
				"(Zom[16384]|zom)",
			},
		},
		{
			patterns: [4]string{
				"zią",
				"",
				"",
				"(Zon[16384]|zon)",
			},
		},
		{
			patterns: [4]string{
				"zię",
				"",
				"[bp]",
				"(Zem[16384]|zem)",
			},
		},
		{
			patterns: [4]string{
				"zię",
				"",
				"",
				"(Zen[16384]|zen)",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"[bcdgkpstwzż]",
				"(Ze[16384]|ze[16384]|ze|tsi[128])",
			},
		},
		{
			patterns: [4]string{
				"zie",
				"",
				"",
				"(ze|Ze[16384]|tsi[128])",
			},
		},
		{
			patterns: [4]string{
				"zio",
				"",
				"",
				"(Zo[16384]|zo)",
			},
		},
		{
			patterns: [4]string{
				"ziu",
				"",
				"",
				"(Zu[16384]|zju)",
			},
		},
		{
			patterns: [4]string{
				"zi",
				"",
				"",
				"(Zi[16384]|zi|tsi[128]|dzi[4096]|tsi[4096]|si[262144])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"$",
				"(s|ts[128]|ts[4096]|S[32768])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"[bdgv]",
				"(z|dz[4096]|Z[32768])",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"[ptckf]",
				"(s|ts[4096]|S[32768])",
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
				"(oue|ve[64])",
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
				"ae",
				"",
				"",
				"(Y[128]|aje[131072]|ae)",
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
				"au",
				"",
				"",
				"(au|o[64])",
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
				"ea",
				"",
				"",
				"(ea|ja[65536])",
			},
		},
		{
			patterns: [4]string{
				"ee",
				"",
				"",
				"(i[32]|aje[131072]|e)",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"(eu|Yj[128]|ej[128]|oj[128]|Y[16])",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: [4]string{
				"ia",
				"",
				"",
				"ja",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"(i[128]|e[16384]|ije[131072]|Q[16]|je)",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"io",
				"",
				"",
				"(jo|e[131072])",
			},
		},
		{
			patterns: [4]string{
				"iu",
				"",
				"",
				"ju",
			},
		},
		{
			patterns: [4]string{
				"iy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"oe",
				"",
				"",
				"(Y[128]|oje[131072]|u[16]|oe)",
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
				"(u[32]|o)",
			},
		},
		{
			patterns: [4]string{
				"ou",
				"",
				"",
				"(ou|u[576]|au[16])",
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
				"oy",
				"",
				"",
				"oj",
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
				"ua",
				"",
				"",
				"va",
			},
		},
		{
			patterns: [4]string{
				"ue",
				"",
				"",
				"(Q[128]|uje[131072]|ve)",
			},
		},
		{
			patterns: [4]string{
				"ui",
				"",
				"",
				"(uj|vi|Y[16])",
			},
		},
		{
			patterns: [4]string{
				"uu",
				"",
				"",
				"(u|Q[16])",
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
				"uy",
				"",
				"",
				"uj",
			},
		},
		{
			patterns: [4]string{
				"ya",
				"",
				"",
				"ja",
			},
		},
		{
			patterns: [4]string{
				"ye",
				"",
				"",
				"(je|ije[131072])",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"^",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yo",
				"",
				"",
				"(jo|e[131072])",
			},
		},
		{
			patterns: [4]string{
				"yu",
				"",
				"",
				"ju",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"i",
				"[áóéê]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"y",
				"[áóéê]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"e",
				"^",
				"",
				"(e|je[131072])",
			},
		},
		{
			patterns: [4]string{
				"e",
				"",
				"$",
				"(e|EE[96])",
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
				"ã",
				"",
				"",
				"(a|an)",
			},
		},
		{
			patterns: [4]string{
				"ă",
				"",
				"",
				"(e[65536]|a)",
			},
		},
		{
			patterns: [4]string{
				"ā",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"č",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ć",
				"",
				"",
				"(tS[16384]|ts)",
			},
		},
		{
			patterns: [4]string{
				"ç",
				"",
				"",
				"(s|tS[524288])",
			},
		},
		{
			patterns: [4]string{
				"ď",
				"",
				"",
				"(d|dj[8])",
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
				"ě",
				"",
				"",
				"(e|je[8])",
			},
		},
		{
			patterns: [4]string{
				"ē",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ģ",
				"",
				"",
				"(d|dj)",
			},
		},
		{
			patterns: [4]string{
				"ğ",
				"",
				"",
				"",
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
				"ī",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ı",
				"",
				"",
				"(i|e[524288]|[524288])",
			},
		},
		{
			patterns: [4]string{
				"ķ",
				"",
				"",
				"(k|t[8192]|tj[8192])",
			},
		},
		{
			patterns: [4]string{
				"ļ",
				"",
				"",
				"l",
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
				"(n|nj[16384])",
			},
		},
		{
			patterns: [4]string{
				"ñ",
				"",
				"",
				"(n|nj[262144])",
			},
		},
		{
			patterns: [4]string{
				"ņ",
				"",
				"",
				"(n|nj[8192])",
			},
		},
		{
			patterns: [4]string{
				"ó",
				"",
				"",
				"(u[16384]|o)",
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
				"õ",
				"",
				"",
				"(o|on[32768]|Y[2048])",
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
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: [4]string{
				"ř",
				"",
				"",
				"(r|rZ[8])",
			},
		},
		{
			patterns: [4]string{
				"ś",
				"",
				"",
				"(S[16384]|s)",
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
				"š",
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
				"ť",
				"",
				"",
				"(t|tj[8])",
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
				"(Q|u[294912])",
			},
		},
		{
			patterns: [4]string{
				"ū",
				"",
				"",
				"u",
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
				"ů",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ù",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ý",
				"",
				"",
				"i",
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
				"(Z[16384]|z)",
			},
		},
		{
			patterns: [4]string{
				"ž",
				"",
				"",
				"Z",
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
				"o",
				"",
				"[bcćdgklłmnńrsśtwzźż]",
				"(O|P[16384])",
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
				"B",
			},
		},
		{
			patterns: [4]string{
				"c",
				"",
				"",
				"(k|ts[16392]|dZ[524288])",
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
				"(h|x[65536]|H[299072])",
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
				"(j|x[262144]|Z[622656])",
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
				"(s|S[32768])",
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
				"V",
			},
		},
		{
			patterns: [4]string{
				"w",
				"",
				"",
				"(v|w[48])",
			},
		},
		{
			patterns: [4]string{
				"x",
				"",
				"",
				"(ks|gz|S[294912])",
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
				"(z|ts[128]|dz[4096]|ts[4096]|s[262144])",
			},
		},
	},
	genarabic: []rule{
		{
			patterns: [4]string{
				"ا",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"ب",
				"",
				"$",
				"b",
			},
		},
		{
			patterns: [4]string{
				"ب",
				"",
				"",
				"b1",
			},
		},
		{
			patterns: [4]string{
				"ت",
				"",
				"$",
				"t",
			},
		},
		{
			patterns: [4]string{
				"ت",
				"",
				"",
				"t1",
			},
		},
		{
			patterns: [4]string{
				"ث",
				"",
				"$",
				"t",
			},
		},
		{
			patterns: [4]string{
				"ث",
				"",
				"",
				"t1",
			},
		},
		{
			patterns: [4]string{
				"ج",
				"",
				"$",
				"(dZ|Z)",
			},
		},
		{
			patterns: [4]string{
				"ج",
				"",
				"",
				"(dZ1|Z1)",
			},
		},
		{
			patterns: [4]string{
				"ح",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ح",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ح",
				"",
				"",
				"(h1|1)",
			},
		},
		{
			patterns: [4]string{
				"خ",
				"",
				"$",
				"x",
			},
		},
		{
			patterns: [4]string{
				"خ",
				"",
				"",
				"x1",
			},
		},
		{
			patterns: [4]string{
				"د",
				"",
				"$",
				"d",
			},
		},
		{
			patterns: [4]string{
				"د",
				"",
				"",
				"d1",
			},
		},
		{
			patterns: [4]string{
				"ذ",
				"",
				"$",
				"d",
			},
		},
		{
			patterns: [4]string{
				"ذ",
				"",
				"",
				"d1",
			},
		},
		{
			patterns: [4]string{
				"ر",
				"",
				"$",
				"r",
			},
		},
		{
			patterns: [4]string{
				"ر",
				"",
				"",
				"r1",
			},
		},
		{
			patterns: [4]string{
				"ز",
				"",
				"$",
				"z",
			},
		},
		{
			patterns: [4]string{
				"ز",
				"",
				"",
				"z1",
			},
		},
		{
			patterns: [4]string{
				"س",
				"",
				"$",
				"s",
			},
		},
		{
			patterns: [4]string{
				"س",
				"",
				"",
				"s1",
			},
		},
		{
			patterns: [4]string{
				"ش",
				"",
				"$",
				"S",
			},
		},
		{
			patterns: [4]string{
				"ش",
				"",
				"",
				"S1",
			},
		},
		{
			patterns: [4]string{
				"ص",
				"",
				"$",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ص",
				"",
				"",
				"s1",
			},
		},
		{
			patterns: [4]string{
				"ض",
				"",
				"$",
				"d",
			},
		},
		{
			patterns: [4]string{
				"ض",
				"",
				"",
				"d1",
			},
		},
		{
			patterns: [4]string{
				"ط",
				"",
				"$",
				"t",
			},
		},
		{
			patterns: [4]string{
				"ط",
				"",
				"",
				"t1",
			},
		},
		{
			patterns: [4]string{
				"ظ",
				"",
				"$",
				"z",
			},
		},
		{
			patterns: [4]string{
				"ظ",
				"",
				"",
				"z1",
			},
		},
		{
			patterns: [4]string{
				"ع",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ع",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ع",
				"",
				"",
				"(h1|1)",
			},
		},
		{
			patterns: [4]string{
				"غ",
				"",
				"$",
				"g",
			},
		},
		{
			patterns: [4]string{
				"غ",
				"",
				"",
				"g1",
			},
		},
		{
			patterns: [4]string{
				"ف",
				"",
				"$",
				"f",
			},
		},
		{
			patterns: [4]string{
				"ف",
				"",
				"",
				"f1",
			},
		},
		{
			patterns: [4]string{
				"ق",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: [4]string{
				"ق",
				"",
				"",
				"k1",
			},
		},
		{
			patterns: [4]string{
				"ك",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: [4]string{
				"ك",
				"",
				"",
				"k1",
			},
		},
		{
			patterns: [4]string{
				"ل",
				"",
				"$",
				"l",
			},
		},
		{
			patterns: [4]string{
				"ل",
				"",
				"",
				"l1",
			},
		},
		{
			patterns: [4]string{
				"م",
				"",
				"$",
				"m",
			},
		},
		{
			patterns: [4]string{
				"م",
				"",
				"",
				"m1",
			},
		},
		{
			patterns: [4]string{
				"ن",
				"",
				"$",
				"n",
			},
		},
		{
			patterns: [4]string{
				"ن",
				"",
				"",
				"n1",
			},
		},
		{
			patterns: [4]string{
				"ه",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ه",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: [4]string{
				"ه",
				"",
				"",
				"(h1|1)",
			},
		},
		{
			patterns: [4]string{
				"و",
				"",
				"$",
				"(u|v)",
			},
		},
		{
			patterns: [4]string{
				"و",
				"",
				"",
				"(u|v1)",
			},
		},
		{
			patterns: [4]string{
				"ي\u200e",
				"",
				"$",
				"(i|j)",
			},
		},
		{
			patterns: [4]string{
				"ي\u200e",
				"",
				"",
				"(i|j1)",
			},
		},
	},
	gencyrillic: []rule{
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
				"с",
				"",
				"с",
				"",
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
				"(hejmer|hajmer)",
			},
		},
		{
			patterns: [4]string{
				"гейм",
				"",
				"$",
				"(hejm|hajm)",
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
				"ей",
				"^",
				"",
				"(jej|ej)",
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
				"ej",
			},
		},
		{
			patterns: [4]string{
				"ей",
				"",
				"",
				"ej",
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
				"ё",
				"",
				"",
				"(e|jo)",
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
	genczech: []rule{
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
				"qu",
				"",
				"",
				"(k|kv)",
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
				"ei",
				"",
				"",
				"(ej|aj)",
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
				"č",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"š",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"ž",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"ň",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"ť",
				"",
				"",
				"(t|tj)",
			},
		},
		{
			patterns: [4]string{
				"ď",
				"",
				"",
				"(d|dj)",
			},
		},
		{
			patterns: [4]string{
				"ř",
				"",
				"",
				"(r|rZ)",
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
				"ý",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ě",
				"",
				"",
				"(e|je)",
			},
		},
		{
			patterns: [4]string{
				"ů",
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
				"(h|g)",
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
				"(k|kv)",
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
				"z",
			},
		},
	},
	gendutch: []rule{
		{
			patterns: [4]string{
				"ssj",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"sj",
				"",
				"",
				"S",
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
				"[eiy]",
				"ts",
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
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"h",
				"[aeiouy]",
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
				"ou",
				"",
				"",
				"au",
			},
		},
		{
			patterns: [4]string{
				"ie",
				"",
				"",
				"(Q|i)",
			},
		},
		{
			patterns: [4]string{
				"uu",
				"",
				"",
				"(Q|u)",
			},
		},
		{
			patterns: [4]string{
				"ee",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"(Y|Yj)",
			},
		},
		{
			patterns: [4]string{
				"aa",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"oo",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"oe",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"ij",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"ui",
				"",
				"",
				"(Y|uj)",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"(ej|aj)",
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
				"i",
				"[aou]",
				"",
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
				"(g|x)",
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
				"(i|Q)",
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
				"(u|Q)",
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
				"(w|v)",
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
	genenglish: []rule{
		{
			patterns: [4]string{
				"�",
				"",
				"",
				"",
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
				"mc",
				"^",
				"",
				"mak",
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
				"x",
				"^",
				"",
				"z",
			},
		},
		{
			patterns: [4]string{
				"y",
				"^",
				"",
				"j",
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
				"yi",
				"^",
				"",
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
				"(aue|oue)",
			},
		},
		{
			patterns: [4]string{
				"ai",
				"",
				"",
				"(aj|ej|e)",
			},
		},
		{
			patterns: [4]string{
				"ay",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: [4]string{
				"a",
				"",
				"[^aeiou]e",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"(ej|aj|i)",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"(ej|aj|i)",
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
				"a",
				"",
				"",
				"(e|o|a)",
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
				"dZ",
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
				"(o|a)",
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
				"(u|a)",
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
				"(w|v)",
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
	genfrench: []rule{
		{
			patterns: [4]string{
				"lt",
				"u",
				"$",
				"(lt|)",
			},
		},
		{
			patterns: [4]string{
				"c",
				"n",
				"$",
				"(k|)",
			},
		},
		{
			patterns: [4]string{
				"d",
				"",
				"$",
				"(t|)",
			},
		},
		{
			patterns: [4]string{
				"g",
				"n",
				"$",
				"(k|)",
			},
		},
		{
			patterns: [4]string{
				"p",
				"",
				"$",
				"(p|)",
			},
		},
		{
			patterns: [4]string{
				"r",
				"e",
				"$",
				"(r|)",
			},
		},
		{
			patterns: [4]string{
				"t",
				"",
				"$",
				"(t|)",
			},
		},
		{
			patterns: [4]string{
				"z",
				"",
				"$",
				"(s|)",
			},
		},
		{
			patterns: [4]string{
				"ds",
				"",
				"$",
				"(ds|)",
			},
		},
		{
			patterns: [4]string{
				"ps",
				"",
				"$",
				"(ps|)",
			},
		},
		{
			patterns: [4]string{
				"rs",
				"e",
				"$",
				"(rs|)",
			},
		},
		{
			patterns: [4]string{
				"ts",
				"",
				"$",
				"(ts|)",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"$",
				"(s|)",
			},
		},
		{
			patterns: [4]string{
				"x",
				"u",
				"$",
				"(ks|)",
			},
		},
		{
			patterns: [4]string{
				"s",
				"[aeéèêiou]",
				"[^aeéèêiou]",
				"(s|)",
			},
		},
		{
			patterns: [4]string{
				"t",
				"[aeéèêiou]",
				"[^aeéèêiou]",
				"(t|)",
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
				"aill",
				"",
				"e",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ll",
				"",
				"e",
				"(l|j)",
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
				"m",
				"[aeiouy]",
				"[aeiouy]",
				"m",
			},
		},
		{
			patterns: [4]string{
				"m",
				"[aeiouy]",
				"",
				"(m|n)",
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
				"au",
				"",
				"",
				"(o|au)",
			},
		},
		{
			patterns: [4]string{
				"ai",
				"",
				"",
				"(e|aj)",
			},
		},
		{
			patterns: [4]string{
				"ay",
				"",
				"",
				"(e|aj)",
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
				"(oj|va)",
			},
		},
		{
			patterns: [4]string{
				"ei",
				"",
				"",
				"(aj|ej|e)",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"(aj|ej|e)",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"(ej|Y)",
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
				"(u|Q)",
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
	gengerman: []rule{
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
				"sch",
				"",
				"",
				"S",
			},
		},
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
				"(aj|ej)",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"(Yj|ej|aj|oj)",
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
	gengreek: []rule{
		{
			patterns: [4]string{
				"αυ",
				"",
				"$",
				"af",
			},
		},
		{
			patterns: [4]string{
				"αυ",
				"",
				"(κ|π|σ|τ|φ|θ|χ|ψ)",
				"af",
			},
		},
		{
			patterns: [4]string{
				"αυ",
				"",
				"",
				"av",
			},
		},
		{
			patterns: [4]string{
				"ευ",
				"",
				"$",
				"ef",
			},
		},
		{
			patterns: [4]string{
				"ευ",
				"",
				"(κ|π|σ|τ|φ|θ|χ|ψ)",
				"ef",
			},
		},
		{
			patterns: [4]string{
				"ευ",
				"",
				"",
				"ev",
			},
		},
		{
			patterns: [4]string{
				"ηυ",
				"",
				"$",
				"if",
			},
		},
		{
			patterns: [4]string{
				"ηυ",
				"",
				"(κ|π|σ|τ|φ|θ|χ|ψ)",
				"if",
			},
		},
		{
			patterns: [4]string{
				"ηυ",
				"",
				"",
				"iv",
			},
		},
		{
			patterns: [4]string{
				"ου",
				"",
				"",
				"u",
			},
		},
		{
			patterns: [4]string{
				"αι",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: [4]string{
				"ει",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"οι",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: [4]string{
				"ωι",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: [4]string{
				"ηι",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"υι",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"γγ",
				"(ε|ι|η|α|ο|ω|υ)",
				"(ε|ι|η)",
				"(nj|j)",
			},
		},
		{
			patterns: [4]string{
				"γγ",
				"",
				"(ε|ι|η)",
				"j",
			},
		},
		{
			patterns: [4]string{
				"γγ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"(ng|g)",
			},
		},
		{
			patterns: [4]string{
				"γγ",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"γκ",
				"^",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"γκ",
				"(ε|ι|η|α|ο|ω|υ)",
				"(ε|ι|η)",
				"(nj|j)",
			},
		},
		{
			patterns: [4]string{
				"γκ",
				"",
				"(ε|ι|η)",
				"j",
			},
		},
		{
			patterns: [4]string{
				"γκ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"(ng|g)",
			},
		},
		{
			patterns: [4]string{
				"γκ",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"γι",
				"",
				"(α|ο|ω|υ)",
				"j",
			},
		},
		{
			patterns: [4]string{
				"γι",
				"",
				"",
				"(gi|i)",
			},
		},
		{
			patterns: [4]string{
				"γε",
				"",
				"(α|ο|ω|υ)",
				"j",
			},
		},
		{
			patterns: [4]string{
				"γε",
				"",
				"",
				"(ge|je)",
			},
		},
		{
			patterns: [4]string{
				"κζ",
				"",
				"",
				"gz",
			},
		},
		{
			patterns: [4]string{
				"τζ",
				"",
				"",
				"dz",
			},
		},
		{
			patterns: [4]string{
				"σ",
				"",
				"(β|γ|δ|μ|ν|ρ)",
				"z",
			},
		},
		{
			patterns: [4]string{
				"μβ",
				"",
				"",
				"(mb|b)",
			},
		},
		{
			patterns: [4]string{
				"μπ",
				"^",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"μπ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"mb",
			},
		},
		{
			patterns: [4]string{
				"μπ",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"ντ",
				"^",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"ντ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"(nd|nt)",
			},
		},
		{
			patterns: [4]string{
				"ντ",
				"",
				"",
				"(nt|d)",
			},
		},
		{
			patterns: [4]string{
				"ά",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"έ",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ή",
				"",
				"",
				"(i|e)",
			},
		},
		{
			patterns: [4]string{
				"ί",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ό",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ύ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: [4]string{
				"ώ",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ΰ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: [4]string{
				"ϋ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: [4]string{
				"ϊ",
				"",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"α",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"β",
				"",
				"",
				"(v|b)",
			},
		},
		{
			patterns: [4]string{
				"γ",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"δ",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"ε",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ζ",
				"",
				"",
				"z",
			},
		},
		{
			patterns: [4]string{
				"η",
				"",
				"",
				"(i|e)",
			},
		},
		{
			patterns: [4]string{
				"ι",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"κ",
				"",
				"",
				"k",
			},
		},
		{
			patterns: [4]string{
				"λ",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"μ",
				"",
				"",
				"m",
			},
		},
		{
			patterns: [4]string{
				"ν",
				"",
				"",
				"n",
			},
		},
		{
			patterns: [4]string{
				"ξ",
				"",
				"",
				"ks",
			},
		},
		{
			patterns: [4]string{
				"ο",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"π",
				"",
				"",
				"p",
			},
		},
		{
			patterns: [4]string{
				"ρ",
				"",
				"",
				"r",
			},
		},
		{
			patterns: [4]string{
				"σ",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"ς",
				"",
				"",
				"s",
			},
		},
		{
			patterns: [4]string{
				"τ",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"υ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: [4]string{
				"φ",
				"",
				"",
				"f",
			},
		},
		{
			patterns: [4]string{
				"θ",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"χ",
				"",
				"",
				"x",
			},
		},
		{
			patterns: [4]string{
				"ψ",
				"",
				"",
				"ps",
			},
		},
		{
			patterns: [4]string{
				"ω",
				"",
				"",
				"o",
			},
		},
	},
	gengreeklatin: []rule{
		{
			patterns: [4]string{
				"au",
				"",
				"$",
				"af",
			},
		},
		{
			patterns: [4]string{
				"au",
				"",
				"[kpstfh]",
				"af",
			},
		},
		{
			patterns: [4]string{
				"au",
				"",
				"",
				"av",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"$",
				"ef",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"[kpstfh]",
				"ef",
			},
		},
		{
			patterns: [4]string{
				"eu",
				"",
				"",
				"ev",
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
				"gge",
				"[aeiouy]",
				"",
				"(nje|je)",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"[aeiouy]",
				"[aou]",
				"(nj|j)",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"[aeiouy]",
				"",
				"(ni|i)",
			},
		},
		{
			patterns: [4]string{
				"gge",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"ggi",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"gg",
				"[aeiouy]",
				"",
				"(ng|g)",
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
				"gk",
				"^",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"gke",
				"[aeiouy]",
				"",
				"(nje|je)",
			},
		},
		{
			patterns: [4]string{
				"gki",
				"[aeiouy]",
				"",
				"(ni|i)",
			},
		},
		{
			patterns: [4]string{
				"gke",
				"",
				"",
				"je",
			},
		},
		{
			patterns: [4]string{
				"gki",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"gk",
				"[aeiouy]",
				"",
				"(ng|g)",
			},
		},
		{
			patterns: [4]string{
				"gk",
				"",
				"",
				"g",
			},
		},
		{
			patterns: [4]string{
				"nghi",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: [4]string{
				"nghi",
				"",
				"",
				"(Ngi|Ni)",
			},
		},
		{
			patterns: [4]string{
				"nghe",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: [4]string{
				"nghe",
				"",
				"",
				"(Nje|Nge)",
			},
		},
		{
			patterns: [4]string{
				"ghi",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ghi",
				"",
				"",
				"(gi|i)",
			},
		},
		{
			patterns: [4]string{
				"ghe",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ghe",
				"",
				"",
				"(je|ge)",
			},
		},
		{
			patterns: [4]string{
				"ngh",
				"",
				"",
				"Ng",
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
				"ngi",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: [4]string{
				"ngi",
				"",
				"",
				"(Ngi|Ni)",
			},
		},
		{
			patterns: [4]string{
				"nge",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: [4]string{
				"nge",
				"",
				"",
				"(Nje|Nge)",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"gi",
				"",
				"",
				"(gi|i)",
			},
		},
		{
			patterns: [4]string{
				"ge",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ge",
				"",
				"",
				"(je|ge)",
			},
		},
		{
			patterns: [4]string{
				"ng",
				"",
				"",
				"Ng",
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
				"i",
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
				"y",
				"[aeou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"",
				"[aeou]",
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
				"ch",
				"",
				"",
				"x",
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
				"dh",
				"",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"dj",
				"",
				"",
				"dZ",
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
				"th",
				"",
				"",
				"t",
			},
		},
		{
			patterns: [4]string{
				"kz",
				"",
				"",
				"gz",
			},
		},
		{
			patterns: [4]string{
				"tz",
				"",
				"",
				"dz",
			},
		},
		{
			patterns: [4]string{
				"s",
				"",
				"[bgdmnr]",
				"z",
			},
		},
		{
			patterns: [4]string{
				"mb",
				"",
				"",
				"(mb|b)",
			},
		},
		{
			patterns: [4]string{
				"mp",
				"^",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"mp",
				"[aeiouy]",
				"",
				"mp",
			},
		},
		{
			patterns: [4]string{
				"mp",
				"",
				"",
				"b",
			},
		},
		{
			patterns: [4]string{
				"nt",
				"^",
				"",
				"d",
			},
		},
		{
			patterns: [4]string{
				"nt",
				"[aeiouy]",
				"",
				"(nd|nt)",
			},
		},
		{
			patterns: [4]string{
				"nt",
				"",
				"",
				"(nt|d)",
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
				"óu",
				"",
				"",
				"u",
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
				"ý",
				"",
				"",
				"(i|Q|u)",
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
				"(b|v)",
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
				"x",
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
				"(j|Z)",
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
				"ο",
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
				"(i|Q|u)",
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
	genhebrew: []rule{
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
	genhungarian: []rule{
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
				"(aj|ej)",
			},
		},
		{
			patterns: [4]string{
				"ey",
				"",
				"",
				"(aj|ej)",
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
				"(ej|e)",
			},
		},
		{
			patterns: [4]string{
				"ely",
				"",
				"",
				"(ej|eli)",
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
				"ú",
				"",
				"",
				"u",
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
	genitalian: []rule{
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
				"h",
				"",
				"$",
				"",
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
	genlatvian: []rule{
		{
			patterns: [4]string{
				"č",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ģ",
				"",
				"",
				"(d|dj)",
			},
		},
		{
			patterns: [4]string{
				"ķ",
				"",
				"",
				"(t|tj)",
			},
		},
		{
			patterns: [4]string{
				"ļ",
				"",
				"",
				"l",
			},
		},
		{
			patterns: [4]string{
				"š",
				"",
				"",
				"S",
			},
		},
		{
			patterns: [4]string{
				"ņ",
				"",
				"",
				"(n|nj)",
			},
		},
		{
			patterns: [4]string{
				"ž",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: [4]string{
				"ā",
				"",
				"",
				"a",
			},
		},
		{
			patterns: [4]string{
				"ē",
				"",
				"",
				"e",
			},
		},
		{
			patterns: [4]string{
				"ī",
				"",
				"",
				"i",
			},
		},
		{
			patterns: [4]string{
				"ū",
				"",
				"",
				"u",
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
				"ei",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: [4]string{
				"io",
				"",
				"",
				"jo",
			},
		},
		{
			patterns: [4]string{
				"iu",
				"",
				"",
				"ju",
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
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: [4]string{
				"ui",
				"",
				"",
				"uj",
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
	genpolish: []rule{
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
				"[aeou]",
				"",
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
	genportuguese: []rule{
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
				"h",
				"",
				"$",
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
	genromanian: []rule{
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
				"qu",
				"",
				"",
				"k",
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
				"(x|h)",
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
				"z",
			},
		},
	},
	genrussian: []rule{
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
				"qu",
				"",
				"",
				"(kv|k)",
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
				"[aeou]",
				"",
				"j",
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
				"[aeiou]",
				"",
				"j",
			},
		},
		{
			patterns: [4]string{
				"ii",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"iy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: [4]string{
				"yi",
				"",
				"$",
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
	genspanish: []rule{
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
				"h",
				"",
				"$",
				"",
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
				"b",
				"",
				"",
				"B",
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
				"(x|Z)",
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
				"V",
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
				"(ks|gz|S)",
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
				"z",
				"",
				"",
				"(z|s)",
			},
		},
	},
	genturkish: []rule{
		{
			patterns: [4]string{
				"ç",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: [4]string{
				"ğ",
				"",
				"",
				"",
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
				"ü",
				"",
				"",
				"Q",
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
				"ı",
				"",
				"",
				"(e|i|)",
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
				"dZ",
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
				"j",
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
}

var genLangRules = []langRule{
	{
		pattern: "^o’",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "^o'",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "^mc",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "^fitz",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "ceau",
		langs:   65600,
		accept:  true,
	},
	{
		pattern: "eau",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "ault$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "oult$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "eux$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "eix$",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "glou$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "uu",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "tx",
		langs:   262144,
		accept:  true,
	},
	{
		pattern: "witz",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "tz$",
		langs:   131232,
		accept:  true,
	},
	{
		pattern: "^tz",
		langs:   131104,
		accept:  true,
	},
	{
		pattern: "poulos$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "pulos$",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "iou",
		langs:   512,
		accept:  true,
	},
	{
		pattern: "sj$",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "^sj",
		langs:   16,
		accept:  true,
	},
	{
		pattern: "güe",
		langs:   262144,
		accept:  true,
	},
	{
		pattern: "güi",
		langs:   262144,
		accept:  true,
	},
	{
		pattern: "ghe",
		langs:   66048,
		accept:  true,
	},
	{
		pattern: "ghi",
		langs:   66048,
		accept:  true,
	},
	{
		pattern: "escu$",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "esco$",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "vici$",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "schi$",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "ii$",
		langs:   131072,
		accept:  true,
	},
	{
		pattern: "iy$",
		langs:   131072,
		accept:  true,
	},
	{
		pattern: "yy$",
		langs:   131072,
		accept:  true,
	},
	{
		pattern: "yi$",
		langs:   131072,
		accept:  true,
	},
	{
		pattern: "^rz",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "rz$",
		langs:   16512,
		accept:  true,
	},
	{
		pattern: "[bcdfgklmnpstwz]rz",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "rz[bcdfghklmnpstw]",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "cki$",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ska$",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "cka$",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ae",
		langs:   131232,
		accept:  true,
	},
	{
		pattern: "oe",
		langs:   131312,
		accept:  true,
	},
	{
		pattern: "th$",
		langs:   160,
		accept:  true,
	},
	{
		pattern: "^th",
		langs:   672,
		accept:  true,
	},
	{
		pattern: "mann",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "cz",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "cy",
		langs:   16896,
		accept:  true,
	},
	{
		pattern: "niew",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "etti$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "eti$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "ati$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "ato$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "[aoei]no$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "[aoei]ni$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "esi$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "oli$",
		langs:   4096,
		accept:  true,
	},
	{
		pattern: "field$",
		langs:   32,
		accept:  true,
	},
	{
		pattern: "stein",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "heim$",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "heimer$",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "thal",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "zweig",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "[aeou]h",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "äh",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "öh",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "üh",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "[ln]h[ao]$",
		langs:   32768,
		accept:  true,
	},
	{
		pattern: "[ln]h[aou]",
		langs:   819416,
		accept:  true,
	},
	{
		pattern: "chsch",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "tsch",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "sch$",
		langs:   131200,
		accept:  true,
	},
	{
		pattern: "^sch",
		langs:   131200,
		accept:  true,
	},
	{
		pattern: "ck$",
		langs:   160,
		accept:  true,
	},
	{
		pattern: "c$",
		langs:   608264,
		accept:  true,
	},
	{
		pattern: "sz",
		langs:   18432,
		accept:  true,
	},
	{
		pattern: "cs$",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "^cs",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "dzs",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "zs$",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "^zs",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "^wl",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "^wr",
		langs:   16560,
		accept:  true,
	},
	{
		pattern: "gy$",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "gy[aeou]",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "gy",
		langs:   133696,
		accept:  true,
	},
	{
		pattern: "guy",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "gu[ei]",
		langs:   294976,
		accept:  true,
	},
	{
		pattern: "gu[ao]",
		langs:   294912,
		accept:  true,
	},
	{
		pattern: "gi[aou]",
		langs:   4608,
		accept:  true,
	},
	{
		pattern: "ly",
		langs:   150016,
		accept:  true,
	},
	{
		pattern: "ny",
		langs:   412160,
		accept:  true,
	},
	{
		pattern: "ty",
		langs:   150016,
		accept:  true,
	},
	{
		pattern: "ā",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ć",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ç",
		langs:   819264,
		accept:  true,
	},
	{
		pattern: "č",
		langs:   8200,
		accept:  true,
	},
	{
		pattern: "ď",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ē",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ğ",
		langs:   524288,
		accept:  true,
	},
	{
		pattern: "ģ",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ī",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ķ",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ļ",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ł",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ņ",
		langs:   8192,
		accept:  true,
	},
	{
		pattern: "ń",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ñ",
		langs:   262144,
		accept:  true,
	},
	{
		pattern: "ň",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ř",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ś",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ş",
		langs:   589824,
		accept:  true,
	},
	{
		pattern: "š",
		langs:   8200,
		accept:  true,
	},
	{
		pattern: "ţ",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "ť",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ź",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ž",
		langs:   8200,
		accept:  true,
	},
	{
		pattern: "ż",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "ß",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "ä",
		langs:   128,
		accept:  true,
	},
	{
		pattern: "á",
		langs:   297480,
		accept:  true,
	},
	{
		pattern: "â",
		langs:   98368,
		accept:  true,
	},
	{
		pattern: "ă",
		langs:   65536,
		accept:  true,
	},
	{
		pattern: "ą",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "à",
		langs:   32768,
		accept:  true,
	},
	{
		pattern: "ã",
		langs:   32768,
		accept:  true,
	},
	{
		pattern: "ę",
		langs:   16384,
		accept:  true,
	},
	{
		pattern: "é",
		langs:   2632,
		accept:  true,
	},
	{
		pattern: "è",
		langs:   266304,
		accept:  true,
	},
	{
		pattern: "ê",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "ě",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ê",
		langs:   32832,
		accept:  true,
	},
	{
		pattern: "í",
		langs:   297480,
		accept:  true,
	},
	{
		pattern: "î",
		langs:   65600,
		accept:  true,
	},
	{
		pattern: "ı",
		langs:   524288,
		accept:  true,
	},
	{
		pattern: "ó",
		langs:   317960,
		accept:  true,
	},
	{
		pattern: "ö",
		langs:   526464,
		accept:  true,
	},
	{
		pattern: "ô",
		langs:   32832,
		accept:  true,
	},
	{
		pattern: "õ",
		langs:   34816,
		accept:  true,
	},
	{
		pattern: "ò",
		langs:   266240,
		accept:  true,
	},
	{
		pattern: "ű",
		langs:   2048,
		accept:  true,
	},
	{
		pattern: "ú",
		langs:   297480,
		accept:  true,
	},
	{
		pattern: "ü",
		langs:   821376,
		accept:  true,
	},
	{
		pattern: "ù",
		langs:   64,
		accept:  true,
	},
	{
		pattern: "ů",
		langs:   8,
		accept:  true,
	},
	{
		pattern: "ý",
		langs:   520,
		accept:  true,
	},
	{
		pattern: "а",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ё",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "о",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "е",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "и",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "у",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ы",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "э",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "ю",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "я",
		langs:   4,
		accept:  true,
	},
	{
		pattern: "α",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ε",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "η",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ι",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ο",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "υ",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ω",
		langs:   256,
		accept:  true,
	},
	{
		pattern: "ا",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ب",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ت",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ث",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ج",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ح",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "خ'",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "د",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ذ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ر",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ز",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "س",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ش",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ص",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ض",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ط",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ظ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ع",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "غ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ف",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ق",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ك",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ل",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "م",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ن",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ه",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "و",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ي",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "آ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "إ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "أ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ؤ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "ئ",
		langs:   2,
		accept:  true,
	},
	{
		pattern: "א",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ב",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ג",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ד",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ה",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ו",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ז",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ח",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ט",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "י",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "כ",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ל",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "מ",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "נ",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ס",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ע",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "פ",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "צ",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ק",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ר",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ש",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "ת",
		langs:   1024,
		accept:  true,
	},
	{
		pattern: "a",
		langs:   1286,
		accept:  false,
	},
	{
		pattern: "o",
		langs:   1286,
		accept:  false,
	},
	{
		pattern: "e",
		langs:   1286,
		accept:  false,
	},
	{
		pattern: "i",
		langs:   1286,
		accept:  false,
	},
	{
		pattern: "y",
		langs:   75030,
		accept:  false,
	},
	{
		pattern: "u",
		langs:   1286,
		accept:  false,
	},
	{
		pattern: "j",
		langs:   4096,
		accept:  false,
	},
	{
		pattern: "j[^aoeiuy]",
		langs:   295488,
		accept:  false,
	},
	{
		pattern: "g",
		langs:   8,
		accept:  false,
	},
	{
		pattern: "k",
		langs:   364608,
		accept:  false,
	},
	{
		pattern: "q",
		langs:   748056,
		accept:  false,
	},
	{
		pattern: "v",
		langs:   16384,
		accept:  false,
	},
	{
		pattern: "w",
		langs:   993864,
		accept:  false,
	},
	{
		pattern: "x",
		langs:   534552,
		accept:  false,
	},
	{
		pattern: "dj",
		langs:   786432,
		accept:  false,
	},
	{
		pattern: "v[^aoeiu]",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "y[^aoeiu]",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "c[^aohk]",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "dzi",
		langs:   524512,
		accept:  false,
	},
	{
		pattern: "ou",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "a[eiou]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "ö[eaiou]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "ü[eaiou]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "e[aiou]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "i[aeou]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "o[aieu]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "u[aieo]",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "aj",
		langs:   240,
		accept:  false,
	},
	{
		pattern: "ej",
		langs:   240,
		accept:  false,
	},
	{
		pattern: "oj",
		langs:   240,
		accept:  false,
	},
	{
		pattern: "uj",
		langs:   240,
		accept:  false,
	},
	{
		pattern: "eu",
		langs:   147456,
		accept:  false,
	},
	{
		pattern: "ky",
		langs:   16384,
		accept:  false,
	},
	{
		pattern: "kie",
		langs:   262720,
		accept:  false,
	},
	{
		pattern: "gie",
		langs:   360960,
		accept:  false,
	},
	{
		pattern: "ch[aou]",
		langs:   4096,
		accept:  false,
	},
	{
		pattern: "ch",
		langs:   524288,
		accept:  false,
	},
	{
		pattern: "son$",
		langs:   128,
		accept:  false,
	},
	{
		pattern: "sc[ei]",
		langs:   64,
		accept:  false,
	},
	{
		pattern: "sch",
		langs:   280640,
		accept:  false,
	},
	{
		pattern: "^h",
		langs:   131072,
		accept:  false,
	},
}

var genFinalRules = finalRules{
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
					"[aA]",
					"",
				},
			},
			{
				patterns: [4]string{
					"a",
					"A",
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
					"j",
					"",
					"j",
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
					"vanden",
					"^",
					"",
					"(vanden|)",
				},
			},
			{
				patterns: [4]string{
					"vander",
					"^",
					"",
					"(vander|)",
				},
			},
			{
				patterns: [4]string{
					"van",
					"^",
					"[bp]",
					"(vam|[16])",
				},
			},
			{
				patterns: [4]string{
					"van",
					"^",
					"",
					"(van|[16])",
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
					"sen",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: [4]string{
					"sen",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: [4]string{
					"sEn",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: [4]string{
					"sEn",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: [4]string{
					"e",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"i",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"E",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"I",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Q",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Y",
					"[BbdfgklmnprsStvzZ]",
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
					"burk",
					"",
					"$",
					"(burk|berk)",
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
					"burg",
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
					"Burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"BUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"Burg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"BUrg",
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
							"mb",
							"",
							"",
							"(mb|b[512])",
						},
					},
					{
						patterns: [4]string{
							"mp",
							"",
							"",
							"(mp|b[512])",
						},
					},
					{
						patterns: [4]string{
							"ng",
							"",
							"",
							"(ng|g[512])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[fktSs]",
							"(p|f[262144])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"p",
							"",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"$",
							"(p|f[262144])",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"[pktSs]",
							"(f|p[262144])",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"f",
							"",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"$",
							"(f|p[262144])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"(b|v[262144])",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"",
							"(v|b[262144])",
						},
					},
					{
						patterns: [4]string{
							"t",
							"",
							"$",
							"(t|[64])",
						},
					},
					{
						patterns: [4]string{
							"g",
							"n",
							"$",
							"(g|[64])",
						},
					},
					{
						patterns: [4]string{
							"k",
							"n",
							"$",
							"(k|[64])",
						},
					},
					{
						patterns: [4]string{
							"p",
							"",
							"$",
							"(p|[64])",
						},
					},
					{
						patterns: [4]string{
							"r",
							"[Ee]",
							"$",
							"(r|[64])",
						},
					},
					{
						patterns: [4]string{
							"s",
							"",
							"$",
							"(s|[64])",
						},
					},
					{
						patterns: [4]string{
							"t",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(t|[64])",
						},
					},
					{
						patterns: [4]string{
							"s",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(s|[64])",
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
							"(Q[128]|i|D[32])",
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
							"(ik|Qk[128])",
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
							"(sits|sQts[128])",
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
							"(Q[128]|i)",
						},
					},
					{
						patterns: [4]string{
							"lEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32])",
						},
					},
					{
						patterns: [4]string{
							"rEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32])",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32]|lY[128])",
						},
					},
					{
						patterns: [4]string{
							"rE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32]|rY[128])",
						},
					},
					{
						patterns: [4]string{
							"EE",
							"",
							"",
							"(i|)",
						},
					},
					{
						patterns: [4]string{
							"ea",
							"",
							"",
							"(D|a|i)",
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
							"eu",
							"",
							"",
							"(D|e|u)",
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
							"(ia|io|iY[128])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[128]|D[32])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
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
							"(i|Y[128])",
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
							"(o|Y[128])",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"",
							"o",
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
							"(a|o|Y[128])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"",
							"(a|o)",
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
							"(uk|Qk[128])",
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
							"(suts|sQts[128])",
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
							"(u|Q[128])",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"",
							"u",
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
							"(i|Y[128])",
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
				langs: 1,
				rules: []rule{
					{
						patterns: [4]string{
							"1a",
							"",
							"",
							"(D|a)",
						},
					},
					{
						patterns: [4]string{
							"1i",
							"",
							"",
							"(D|i|e)",
						},
					},
					{
						patterns: [4]string{
							"1u",
							"",
							"",
							"(D|u|o)",
						},
					},
					{
						patterns: [4]string{
							"j1",
							"",
							"",
							"(ja|je|jo|ju|j)",
						},
					},
					{
						patterns: [4]string{
							"1",
							"",
							"",
							"(a|e|i|o|u|)",
						},
					},
					{
						patterns: [4]string{
							"u",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"i",
							"",
							"",
							"(i|e)",
						},
					},
					{
						patterns: [4]string{
							"p",
							"",
							"$",
							"p",
						},
					},
					{
						patterns: [4]string{
							"p",
							"",
							"",
							"(p|b)",
						},
					},
				},
			},
			{
				langs: 17,
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
				langs: 6,
				rules: []rule{
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
				},
			},
			{
				langs: 3,
				rules: []rule{
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
				},
			},
			{
				langs: 4,
				rules: []rule{
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
				},
			},
			{
				langs: 5,
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
				langs: 7,
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
				langs: 8,
				rules: []rule{
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
				},
			},
			{
				langs: 9,
				rules: []rule{
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
							"N",
							"",
							"",
							"",
						},
					},
				},
			},
			{
				langs: 10,
				rules: []rule{},
			},
			{
				langs: 11,
				rules: []rule{
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
				},
			},
			{
				langs: 12,
				rules: []rule{
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
				},
			},
			{
				langs: 13,
				rules: []rule{
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
				},
			},
			{
				langs: 14,
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
				langs: 15,
				rules: []rule{
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
				},
			},
			{
				langs: 16,
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
				langs: 18,
				rules: []rule{
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
							"B",
							"",
							"",
							"(b|v)",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"",
							"(b|v)",
						},
					},
				},
			},
			{
				langs: 19,
				rules: []rule{
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
					"[aA]",
					"",
				},
			},
			{
				patterns: [4]string{
					"a",
					"A",
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
					"j",
					"",
					"j",
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
					"vanden",
					"^",
					"",
					"(vanden|)",
				},
			},
			{
				patterns: [4]string{
					"vander",
					"^",
					"",
					"(vander|)",
				},
			},
			{
				patterns: [4]string{
					"van",
					"^",
					"[bp]",
					"(vam|[16])",
				},
			},
			{
				patterns: [4]string{
					"van",
					"^",
					"",
					"(van|[16])",
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
					"sen",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: [4]string{
					"sen",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: [4]string{
					"sEn",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: [4]string{
					"sEn",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: [4]string{
					"e",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"i",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"E",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"I",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Q",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: [4]string{
					"Y",
					"[BbdfgklmnprsStvzZ]",
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
					"burk",
					"",
					"$",
					"(burk|berk)",
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
					"burg",
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
					"Burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"BUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"Burg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: [4]string{
					"BUrg",
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
							"mb",
							"",
							"",
							"(mb|b[512])",
						},
					},
					{
						patterns: [4]string{
							"mp",
							"",
							"",
							"(mp|b[512])",
						},
					},
					{
						patterns: [4]string{
							"ng",
							"",
							"",
							"(ng|g[512])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"[fktSs]",
							"(p|f[262144])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"p",
							"",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"$",
							"(p|f[262144])",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"[pktSs]",
							"(f|p[262144])",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"f",
							"",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"$",
							"(f|p[262144])",
						},
					},
					{
						patterns: [4]string{
							"B",
							"",
							"",
							"(b|v[262144])",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"",
							"(v|b[262144])",
						},
					},
					{
						patterns: [4]string{
							"t",
							"",
							"$",
							"(t|[64])",
						},
					},
					{
						patterns: [4]string{
							"g",
							"n",
							"$",
							"(g|[64])",
						},
					},
					{
						patterns: [4]string{
							"k",
							"n",
							"$",
							"(k|[64])",
						},
					},
					{
						patterns: [4]string{
							"p",
							"",
							"$",
							"(p|[64])",
						},
					},
					{
						patterns: [4]string{
							"r",
							"[Ee]",
							"$",
							"(r|[64])",
						},
					},
					{
						patterns: [4]string{
							"s",
							"",
							"$",
							"(s|[64])",
						},
					},
					{
						patterns: [4]string{
							"t",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(t|[64])",
						},
					},
					{
						patterns: [4]string{
							"s",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(s|[64])",
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
							"(Q[128]|i|D[32])",
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
							"(ik|Qk[128])",
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
							"(sits|sQts[128])",
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
							"(Q[128]|i)",
						},
					},
					{
						patterns: [4]string{
							"lEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32])",
						},
					},
					{
						patterns: [4]string{
							"rEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32])",
						},
					},
					{
						patterns: [4]string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32]|lY[128])",
						},
					},
					{
						patterns: [4]string{
							"rE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32]|rY[128])",
						},
					},
					{
						patterns: [4]string{
							"EE",
							"",
							"",
							"(i|)",
						},
					},
					{
						patterns: [4]string{
							"ea",
							"",
							"",
							"(D|a|i)",
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
							"eu",
							"",
							"",
							"(D|e|u)",
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
							"(ia|io|iY[128])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[128]|D[32])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
						},
					},
					{
						patterns: [4]string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
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
							"(i|Y[128])",
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
							"(o|Y[128])",
						},
					},
					{
						patterns: [4]string{
							"O",
							"",
							"",
							"o",
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
							"(a|o|Y[128])",
						},
					},
					{
						patterns: [4]string{
							"A",
							"",
							"",
							"(a|o)",
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
							"(uk|Qk[128])",
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
							"(suts|sQts[128])",
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
							"(u|Q[128])",
						},
					},
					{
						patterns: [4]string{
							"U",
							"",
							"",
							"u",
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
							"(i|Y[128])",
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
				langs: 1,
				rules: []rule{
					{
						patterns: [4]string{
							"1a",
							"",
							"",
							"(D|a)",
						},
					},
					{
						patterns: [4]string{
							"1i",
							"",
							"",
							"(D|i|e)",
						},
					},
					{
						patterns: [4]string{
							"1u",
							"",
							"",
							"(D|u|o)",
						},
					},
					{
						patterns: [4]string{
							"j1",
							"",
							"",
							"(ja|je|jo|ju|j)",
						},
					},
					{
						patterns: [4]string{
							"1",
							"",
							"",
							"(a|e|i|o|u|)",
						},
					},
					{
						patterns: [4]string{
							"u",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: [4]string{
							"i",
							"",
							"",
							"(i|e)",
						},
					},
					{
						patterns: [4]string{
							"p",
							"",
							"$",
							"p",
						},
					},
					{
						patterns: [4]string{
							"p",
							"",
							"",
							"(p|b)",
						},
					},
				},
			},
			{
				langs: 17,
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
				langs: 6,
				rules: []rule{
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
				},
			},
			{
				langs: 3,
				rules: []rule{
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
				},
			},
			{
				langs: 4,
				rules: []rule{
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
				},
			},
			{
				langs: 5,
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
				langs: 7,
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
				langs: 8,
				rules: []rule{
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
				},
			},
			{
				langs: 9,
				rules: []rule{
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
							"N",
							"",
							"",
							"",
						},
					},
				},
			},
			{
				langs: 10,
				rules: []rule{},
			},
			{
				langs: 11,
				rules: []rule{
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
				},
			},
			{
				langs: 12,
				rules: []rule{
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
				},
			},
			{
				langs: 13,
				rules: []rule{
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
				},
			},
			{
				langs: 14,
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
				langs: 15,
				rules: []rule{
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
				},
			},
			{
				langs: 16,
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
				langs: 18,
				rules: []rule{
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
							"B",
							"",
							"",
							"(b|v)",
						},
					},
					{
						patterns: [4]string{
							"V",
							"",
							"",
							"(b|v)",
						},
					},
				},
			},
			{
				langs: 19,
				rules: []rule{
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
				},
			},
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
