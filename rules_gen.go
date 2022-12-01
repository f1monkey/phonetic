// GENERATED CODE. DO NOT EDIT!
package bmpm

type genLang int

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

var genRules = map[genLang][]rule{
	genany: []rule{
		{
			patterns: []string{
				"yna",
				"",
				"$",
				"(in[131072]|ina)",
			},
		},
		{
			patterns: []string{
				"ina",
				"",
				"$",
				"(in[131072]|ina)",
			},
		},
		{
			patterns: []string{
				"liova",
				"",
				"$",
				"(lova|lof[131072]|lef[131072])",
			},
		},
		{
			patterns: []string{
				"lova",
				"",
				"$",
				"(lova|lof[131072]|lef[131072]|l[8]|el[8])",
			},
		},
		{
			patterns: []string{
				"kova",
				"",
				"$",
				"(kova|kof[131072]|k[8]|ek[8])",
			},
		},
		{
			patterns: []string{
				"ova",
				"",
				"$",
				"(ova|of[131072]|[8])",
			},
		},
		{
			patterns: []string{
				"ová",
				"",
				"$",
				"(ova|[8])",
			},
		},
		{
			patterns: []string{
				"eva",
				"",
				"$",
				"(eva|ef[131072])",
			},
		},
		{
			patterns: []string{
				"aia",
				"",
				"$",
				"(aja|i[131072])",
			},
		},
		{
			patterns: []string{
				"aja",
				"",
				"$",
				"(aja|i[131072])",
			},
		},
		{
			patterns: []string{
				"aya",
				"",
				"$",
				"(aja|i[131072])",
			},
		},
		{
			patterns: []string{
				"lowa",
				"",
				"$",
				"(lova|lof[16384]|l[16384]|el[16384])",
			},
		},
		{
			patterns: []string{
				"kowa",
				"",
				"$",
				"(kova|kof[16384]|k[16384]|ek[16384])",
			},
		},
		{
			patterns: []string{
				"owa",
				"",
				"$",
				"(ova|of[16384]|)",
			},
		},
		{
			patterns: []string{
				"lowna",
				"",
				"$",
				"(lovna|levna|l[16384]|el[16384])",
			},
		},
		{
			patterns: []string{
				"kowna",
				"",
				"$",
				"(kovna|k[16384]|ek[16384])",
			},
		},
		{
			patterns: []string{
				"owna",
				"",
				"$",
				"(ovna|[16384])",
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
				"á",
				"",
				"$",
				"(a|i[8])",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"$",
				"(a|i[16392])",
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
				"que",
				"",
				"$",
				"(k[64]|ke|kve)",
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
				"m",
				"",
				"[bfpv]",
				"(m|n)",
			},
		},
		{
			patterns: []string{
				"m",
				"[aeiouy]",
				"[aeiouy]",
				"m",
			},
		},
		{
			patterns: []string{
				"m",
				"[aeiouy]",
				"",
				"(m|n[32832])",
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
				"lio",
				"",
				"",
				"(lo|le[131072])",
			},
		},
		{
			patterns: []string{
				"lyo",
				"",
				"",
				"(lo|le[131072])",
			},
		},
		{
			patterns: []string{
				"lt",
				"u",
				"$",
				"(lt|[64])",
			},
		},
		{
			patterns: []string{
				"v",
				"^",
				"",
				"(v|f[128]|b[262144])",
			},
		},
		{
			patterns: []string{
				"ex",
				"",
				"[aáuiíoóeéêy]",
				"(ez[32768]|eS[32768]|eks|egz)",
			},
		},
		{
			patterns: []string{
				"ex",
				"",
				"[cs]",
				"(e[32768]|ek)",
			},
		},
		{
			patterns: []string{
				"x",
				"u",
				"$",
				"(ks|[64])",
			},
		},
		{
			patterns: []string{
				"ck",
				"",
				"",
				"(k|tsk[16392])",
			},
		},
		{
			patterns: []string{
				"cz",
				"",
				"",
				"(tS|tsz[8])",
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
				"dh",
				"^",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"bh",
				"^",
				"",
				"b",
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
				"kh",
				"",
				"",
				"(x[131104]|kh)",
			},
		},
		{
			patterns: []string{
				"lh",
				"",
				"",
				"(lh|l[32768])",
			},
		},
		{
			patterns: []string{
				"nh",
				"",
				"",
				"(nh|nj[32768])",
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
				"[aeiouy]",
				"[ei]",
				"(S|StS[131072]|sk[69632])",
			},
		},
		{
			patterns: []string{
				"sch",
				"[aeiouy]",
				"",
				"(S|StS[131072])",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"[ei]",
				"(sk[69632]|S|StS[131072])",
			},
		},
		{
			patterns: []string{
				"sch",
				"",
				"",
				"(S|StS[131072])",
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
				"(S[131104]|sh)",
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
				"(Z[131104]|zh|tsh[128])",
			},
		},
		{
			patterns: []string{
				"chs",
				"",
				"",
				"(ks[128]|xs|tSs[131104])",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"[ei]",
				"(x|tS[393248]|k[69632]|S[32832])",
			},
		},
		{
			patterns: []string{
				"ch",
				"",
				"",
				"(x|tS[393248]|S[32832])",
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
				"(t[672]|th)",
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
				"gh",
				"",
				"[ei]",
				"(g[70144]|gh)",
			},
		},
		{
			patterns: []string{
				"ouh",
				"",
				"[aioe]",
				"(v[64]|uh)",
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
				"h",
				"",
				"$",
				"",
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
				"(h|x[66048]|H[381024])",
			},
		},
		{
			patterns: []string{
				"cia",
				"",
				"",
				"(tSa[16384]|tsa)",
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
				"(tSon[16384]|tson)",
			},
		},
		{
			patterns: []string{
				"cię",
				"",
				"[bp]",
				"(tSem[16384]|tsem)",
			},
		},
		{
			patterns: []string{
				"cię",
				"",
				"",
				"(tSen[16384]|tsen)",
			},
		},
		{
			patterns: []string{
				"cie",
				"",
				"",
				"(tSe[16384]|tse)",
			},
		},
		{
			patterns: []string{
				"cio",
				"",
				"",
				"(tSo[16384]|tso)",
			},
		},
		{
			patterns: []string{
				"ciu",
				"",
				"",
				"(tSu[16384]|tsu)",
			},
		},
		{
			patterns: []string{
				"sci",
				"",
				"$",
				"(Si[4096]|stsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
			},
		},
		{
			patterns: []string{
				"sc",
				"",
				"[ei]",
				"(S[4096]|sts[16392]|dZ[524288]|tS[81920]|s)",
			},
		},
		{
			patterns: []string{
				"ci",
				"",
				"$",
				"(tsi[16392]|dZi[524288]|tSi[81920]|tS[65536]|si)",
			},
		},
		{
			patterns: []string{
				"cy",
				"",
				"",
				"(si|tsi[16384])",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"[ei]",
				"(ts[16392]|dZ[524288]|tS[81920]|k[512]|s)",
			},
		},
		{
			patterns: []string{
				"sç",
				"",
				"[aeiou]",
				"(s|stS[524288])",
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
				"^",
				"",
				"(S|s[2048])",
			},
		},
		{
			patterns: []string{
				"sz",
				"",
				"$",
				"(S|s[2048])",
			},
		},
		{
			patterns: []string{
				"sz",
				"",
				"",
				"(S|s[2048]|sts[128])",
			},
		},
		{
			patterns: []string{
				"ssp",
				"",
				"",
				"(Sp[128]|sp)",
			},
		},
		{
			patterns: []string{
				"sp",
				"",
				"",
				"(Sp[128]|sp)",
			},
		},
		{
			patterns: []string{
				"sst",
				"",
				"",
				"(St[128]|st)",
			},
		},
		{
			patterns: []string{
				"st",
				"",
				"",
				"(St[128]|st)",
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
				"sj",
				"^",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"sj",
				"",
				"$",
				"S",
			},
		},
		{
			patterns: []string{
				"sj",
				"",
				"",
				"(sj|S[16]|sx[262144]|sZ[589824])",
			},
		},
		{
			patterns: []string{
				"sia",
				"",
				"",
				"(Sa[16384]|sa[16384]|sja)",
			},
		},
		{
			patterns: []string{
				"sią",
				"",
				"[bp]",
				"(Som[16384]|som)",
			},
		},
		{
			patterns: []string{
				"sią",
				"",
				"",
				"(Son[16384]|son)",
			},
		},
		{
			patterns: []string{
				"się",
				"",
				"[bp]",
				"(Sem[16384]|sem)",
			},
		},
		{
			patterns: []string{
				"się",
				"",
				"",
				"(Sen[16384]|sen)",
			},
		},
		{
			patterns: []string{
				"sie",
				"",
				"",
				"(se|sje|Se[16384]|zi[128])",
			},
		},
		{
			patterns: []string{
				"sio",
				"",
				"",
				"(So[16384]|so)",
			},
		},
		{
			patterns: []string{
				"siu",
				"",
				"",
				"(Su[16384]|sju)",
			},
		},
		{
			patterns: []string{
				"si",
				"[äöëaáuiíoóeéêy]",
				"",
				"(Si[16384]|si|zi[37056])",
			},
		},
		{
			patterns: []string{
				"si",
				"",
				"",
				"(Si[16384]|si|zi[128])",
			},
		},
		{
			patterns: []string{
				"s",
				"[aáuiíoóeéêy]",
				"[aáuíoóeéêy]",
				"(s|z[37056])",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"[aeouäöë]",
				"(s|z[128])",
			},
		},
		{
			patterns: []string{
				"s",
				"[aeiouy]",
				"[dglmnrv]",
				"(s|z|Z[32768]|[64])",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"[dglmnrv]",
				"(s|z|Z[32768])",
			},
		},
		{
			patterns: []string{
				"gue",
				"",
				"$",
				"(k[64]|gve)",
			},
		},
		{
			patterns: []string{
				"gu",
				"",
				"[ei]",
				"(g[64]|gv[294912])",
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
				"guy",
				"",
				"",
				"gi",
			},
		},
		{
			patterns: []string{
				"gli",
				"",
				"",
				"(glI|l[4096])",
			},
		},
		{
			patterns: []string{
				"gni",
				"",
				"",
				"(gnI|ni[4160])",
			},
		},
		{
			patterns: []string{
				"gn",
				"",
				"[aeou]",
				"(n[4160]|nj[4160]|gn)",
			},
		},
		{
			patterns: []string{
				"ggie",
				"",
				"",
				"(je[512]|dZe)",
			},
		},
		{
			patterns: []string{
				"ggi",
				"",
				"[aou]",
				"(j[512]|dZ)",
			},
		},
		{
			patterns: []string{
				"ggi",
				"[yaeiou]",
				"[aou]",
				"(gI|dZ[4096]|j[512])",
			},
		},
		{
			patterns: []string{
				"gge",
				"[yaeiou]",
				"",
				"(gE|xe[262144]|gZe[32832]|dZe[331808]|je[512])",
			},
		},
		{
			patterns: []string{
				"ggi",
				"[yaeiou]",
				"",
				"(gI|xi[262144]|gZi[32832]|dZi[331808]|i[512])",
			},
		},
		{
			patterns: []string{
				"ggi",
				"",
				"[aou]",
				"(gI|dZ[4096]|j[512])",
			},
		},
		{
			patterns: []string{
				"gie",
				"",
				"$",
				"(ge|gi[128]|ji[64]|dZe[4096])",
			},
		},
		{
			patterns: []string{
				"gie",
				"",
				"",
				"(ge|gi[128]|dZe[4096]|je[512])",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"[aou]",
				"(i[512]|dZ)",
			},
		},
		{
			patterns: []string{
				"ge",
				"[yaeiou]",
				"",
				"(gE|xe[262144]|Ze[32832]|dZe[331808])",
			},
		},
		{
			patterns: []string{
				"gi",
				"[yaeiou]",
				"",
				"(gI|xi[262144]|Zi[32832]|dZi[331808])",
			},
		},
		{
			patterns: []string{
				"ge",
				"",
				"",
				"(gE|xe[262144]|hE[131072]|je[512]|Ze[32832]|dZe[331808])",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"",
				"(gI|xi[262144]|hI[131072]|i[512]|Zi[32832]|dZi[331808])",
			},
		},
		{
			patterns: []string{
				"gy",
				"",
				"[aeouáéóúüöőű]",
				"(gi|dj[2048])",
			},
		},
		{
			patterns: []string{
				"gy",
				"",
				"",
				"(gi|d[2048])",
			},
		},
		{
			patterns: []string{
				"g",
				"[yaeiou]",
				"[aouyei]",
				"g",
			},
		},
		{
			patterns: []string{
				"g",
				"",
				"[aouei]",
				"(g|h[131072])",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"",
				"(i|ej[16]|ix[262144]|iZ[622656])",
			},
		},
		{
			patterns: []string{
				"j",
				"",
				"[aoeiuy]",
				"(j|dZ[32]|x[262144]|Z[622656])",
			},
		},
		{
			patterns: []string{
				"rz",
				"t",
				"",
				"(S[16384]|r)",
			},
		},
		{
			patterns: []string{
				"rz",
				"",
				"",
				"(rz|rts[128]|Z[16384]|r[16384]|rZ[16384])",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"$",
				"(ts|tS[160])",
			},
		},
		{
			patterns: []string{
				"tz",
				"^",
				"",
				"(ts[131232]|tS[160])",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"",
				"(ts[131232]|tz)",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"[bcdgkpstwzż]",
				"(Za[16384]|za[16384]|zja)",
			},
		},
		{
			patterns: []string{
				"zia",
				"",
				"",
				"(Za[16384]|zja)",
			},
		},
		{
			patterns: []string{
				"zią",
				"",
				"[bp]",
				"(Zom[16384]|zom)",
			},
		},
		{
			patterns: []string{
				"zią",
				"",
				"",
				"(Zon[16384]|zon)",
			},
		},
		{
			patterns: []string{
				"zię",
				"",
				"[bp]",
				"(Zem[16384]|zem)",
			},
		},
		{
			patterns: []string{
				"zię",
				"",
				"",
				"(Zen[16384]|zen)",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"[bcdgkpstwzż]",
				"(Ze[16384]|ze[16384]|ze|tsi[128])",
			},
		},
		{
			patterns: []string{
				"zie",
				"",
				"",
				"(ze|Ze[16384]|tsi[128])",
			},
		},
		{
			patterns: []string{
				"zio",
				"",
				"",
				"(Zo[16384]|zo)",
			},
		},
		{
			patterns: []string{
				"ziu",
				"",
				"",
				"(Zu[16384]|zju)",
			},
		},
		{
			patterns: []string{
				"zi",
				"",
				"",
				"(Zi[16384]|zi|tsi[128]|dzi[4096]|tsi[4096]|si[262144])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"$",
				"(s|ts[128]|ts[4096]|S[32768])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"[bdgv]",
				"(z|dz[4096]|Z[32768])",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"[ptckf]",
				"(s|ts[4096]|S[32768])",
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
				"(oue|ve[64])",
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
				"ae",
				"",
				"",
				"(Y[128]|aje[131072]|ae)",
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
				"au",
				"",
				"",
				"(au|o[64])",
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
				"ea",
				"",
				"",
				"(ea|ja[65536])",
			},
		},
		{
			patterns: []string{
				"ee",
				"",
				"",
				"(i[32]|aje[131072]|e)",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"(eu|Yj[128]|ej[128]|oj[128]|Y[16])",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: []string{
				"ia",
				"",
				"",
				"ja",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"(i[128]|e[16384]|ije[131072]|Q[16]|je)",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"io",
				"",
				"",
				"(jo|e[131072])",
			},
		},
		{
			patterns: []string{
				"iu",
				"",
				"",
				"ju",
			},
		},
		{
			patterns: []string{
				"iy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"oe",
				"",
				"",
				"(Y[128]|oje[131072]|u[16]|oe)",
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
				"(u[32]|o)",
			},
		},
		{
			patterns: []string{
				"ou",
				"",
				"",
				"(ou|u[576]|au[16])",
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
				"oy",
				"",
				"",
				"oj",
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
				"ua",
				"",
				"",
				"va",
			},
		},
		{
			patterns: []string{
				"ue",
				"",
				"",
				"(Q[128]|uje[131072]|ve)",
			},
		},
		{
			patterns: []string{
				"ui",
				"",
				"",
				"(uj|vi|Y[16])",
			},
		},
		{
			patterns: []string{
				"uu",
				"",
				"",
				"(u|Q[16])",
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
				"uy",
				"",
				"",
				"uj",
			},
		},
		{
			patterns: []string{
				"ya",
				"",
				"",
				"ja",
			},
		},
		{
			patterns: []string{
				"ye",
				"",
				"",
				"(je|ije[131072])",
			},
		},
		{
			patterns: []string{
				"yi",
				"^",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"yi",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"yo",
				"",
				"",
				"(jo|e[131072])",
			},
		},
		{
			patterns: []string{
				"yu",
				"",
				"",
				"ju",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"i",
				"[áóéê]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"y",
				"[áóéê]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"e",
				"^",
				"",
				"(e|je[131072])",
			},
		},
		{
			patterns: []string{
				"e",
				"",
				"$",
				"(e|EE[96])",
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
				"ã",
				"",
				"",
				"(a|an)",
			},
		},
		{
			patterns: []string{
				"ă",
				"",
				"",
				"(e[65536]|a)",
			},
		},
		{
			patterns: []string{
				"ā",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"č",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ć",
				"",
				"",
				"(tS[16384]|ts)",
			},
		},
		{
			patterns: []string{
				"ç",
				"",
				"",
				"(s|tS[524288])",
			},
		},
		{
			patterns: []string{
				"ď",
				"",
				"",
				"(d|dj[8])",
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
				"ě",
				"",
				"",
				"(e|je[8])",
			},
		},
		{
			patterns: []string{
				"ē",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ģ",
				"",
				"",
				"(d|dj)",
			},
		},
		{
			patterns: []string{
				"ğ",
				"",
				"",
				"",
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
				"ī",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ı",
				"",
				"",
				"(i|e[524288]|[524288])",
			},
		},
		{
			patterns: []string{
				"ķ",
				"",
				"",
				"(k|t[8192]|tj[8192])",
			},
		},
		{
			patterns: []string{
				"ļ",
				"",
				"",
				"l",
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
				"(n|nj[16384])",
			},
		},
		{
			patterns: []string{
				"ñ",
				"",
				"",
				"(n|nj[262144])",
			},
		},
		{
			patterns: []string{
				"ņ",
				"",
				"",
				"(n|nj[8192])",
			},
		},
		{
			patterns: []string{
				"ó",
				"",
				"",
				"(u[16384]|o)",
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
				"õ",
				"",
				"",
				"(o|on[32768]|Y[2048])",
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
				"ö",
				"",
				"",
				"Y",
			},
		},
		{
			patterns: []string{
				"ř",
				"",
				"",
				"(r|rZ[8])",
			},
		},
		{
			patterns: []string{
				"ś",
				"",
				"",
				"(S[16384]|s)",
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
				"š",
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
				"ť",
				"",
				"",
				"(t|tj[8])",
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
				"(Q|u[294912])",
			},
		},
		{
			patterns: []string{
				"ū",
				"",
				"",
				"u",
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
				"ů",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ù",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ý",
				"",
				"",
				"i",
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
				"(Z[16384]|z)",
			},
		},
		{
			patterns: []string{
				"ž",
				"",
				"",
				"Z",
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
				"o",
				"",
				"[bcćdgklłmnńrsśtwzźż]",
				"(O|P[16384])",
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
				"B",
			},
		},
		{
			patterns: []string{
				"c",
				"",
				"",
				"(k|ts[16392]|dZ[524288])",
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
				"(h|x[65536]|H[299072])",
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
				"(j|x[262144]|Z[622656])",
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
				"(s|S[32768])",
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
				"V",
			},
		},
		{
			patterns: []string{
				"w",
				"",
				"",
				"(v|w[48])",
			},
		},
		{
			patterns: []string{
				"x",
				"",
				"",
				"(ks|gz|S[294912])",
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
				"(z|ts[128]|dz[4096]|ts[4096]|s[262144])",
			},
		},
	},
	genarabic: []rule{
		{
			patterns: []string{
				"ا",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"ب",
				"",
				"$",
				"b",
			},
		},
		{
			patterns: []string{
				"ب",
				"",
				"",
				"b1",
			},
		},
		{
			patterns: []string{
				"ت",
				"",
				"$",
				"t",
			},
		},
		{
			patterns: []string{
				"ت",
				"",
				"",
				"t1",
			},
		},
		{
			patterns: []string{
				"ث",
				"",
				"$",
				"t",
			},
		},
		{
			patterns: []string{
				"ث",
				"",
				"",
				"t1",
			},
		},
		{
			patterns: []string{
				"ج",
				"",
				"$",
				"(dZ|Z)",
			},
		},
		{
			patterns: []string{
				"ج",
				"",
				"",
				"(dZ1|Z1)",
			},
		},
		{
			patterns: []string{
				"ح",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: []string{
				"ح",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: []string{
				"ح",
				"",
				"",
				"(h1|1)",
			},
		},
		{
			patterns: []string{
				"خ",
				"",
				"$",
				"x",
			},
		},
		{
			patterns: []string{
				"خ",
				"",
				"",
				"x1",
			},
		},
		{
			patterns: []string{
				"د",
				"",
				"$",
				"d",
			},
		},
		{
			patterns: []string{
				"د",
				"",
				"",
				"d1",
			},
		},
		{
			patterns: []string{
				"ذ",
				"",
				"$",
				"d",
			},
		},
		{
			patterns: []string{
				"ذ",
				"",
				"",
				"d1",
			},
		},
		{
			patterns: []string{
				"ر",
				"",
				"$",
				"r",
			},
		},
		{
			patterns: []string{
				"ر",
				"",
				"",
				"r1",
			},
		},
		{
			patterns: []string{
				"ز",
				"",
				"$",
				"z",
			},
		},
		{
			patterns: []string{
				"ز",
				"",
				"",
				"z1",
			},
		},
		{
			patterns: []string{
				"س",
				"",
				"$",
				"s",
			},
		},
		{
			patterns: []string{
				"س",
				"",
				"",
				"s1",
			},
		},
		{
			patterns: []string{
				"ش",
				"",
				"$",
				"S",
			},
		},
		{
			patterns: []string{
				"ش",
				"",
				"",
				"S1",
			},
		},
		{
			patterns: []string{
				"ص",
				"",
				"$",
				"s",
			},
		},
		{
			patterns: []string{
				"ص",
				"",
				"",
				"s1",
			},
		},
		{
			patterns: []string{
				"ض",
				"",
				"$",
				"d",
			},
		},
		{
			patterns: []string{
				"ض",
				"",
				"",
				"d1",
			},
		},
		{
			patterns: []string{
				"ط",
				"",
				"$",
				"t",
			},
		},
		{
			patterns: []string{
				"ط",
				"",
				"",
				"t1",
			},
		},
		{
			patterns: []string{
				"ظ",
				"",
				"$",
				"z",
			},
		},
		{
			patterns: []string{
				"ظ",
				"",
				"",
				"z1",
			},
		},
		{
			patterns: []string{
				"ع",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: []string{
				"ع",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: []string{
				"ع",
				"",
				"",
				"(h1|1)",
			},
		},
		{
			patterns: []string{
				"غ",
				"",
				"$",
				"g",
			},
		},
		{
			patterns: []string{
				"غ",
				"",
				"",
				"g1",
			},
		},
		{
			patterns: []string{
				"ف",
				"",
				"$",
				"f",
			},
		},
		{
			patterns: []string{
				"ف",
				"",
				"",
				"f1",
			},
		},
		{
			patterns: []string{
				"ق",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: []string{
				"ق",
				"",
				"",
				"k1",
			},
		},
		{
			patterns: []string{
				"ك",
				"",
				"$",
				"k",
			},
		},
		{
			patterns: []string{
				"ك",
				"",
				"",
				"k1",
			},
		},
		{
			patterns: []string{
				"ل",
				"",
				"$",
				"l",
			},
		},
		{
			patterns: []string{
				"ل",
				"",
				"",
				"l1",
			},
		},
		{
			patterns: []string{
				"م",
				"",
				"$",
				"m",
			},
		},
		{
			patterns: []string{
				"م",
				"",
				"",
				"m1",
			},
		},
		{
			patterns: []string{
				"ن",
				"",
				"$",
				"n",
			},
		},
		{
			patterns: []string{
				"ن",
				"",
				"",
				"n1",
			},
		},
		{
			patterns: []string{
				"ه",
				"^",
				"",
				"1",
			},
		},
		{
			patterns: []string{
				"ه",
				"",
				"$",
				"1",
			},
		},
		{
			patterns: []string{
				"ه",
				"",
				"",
				"(h1|1)",
			},
		},
		{
			patterns: []string{
				"و",
				"",
				"$",
				"(u|v)",
			},
		},
		{
			patterns: []string{
				"و",
				"",
				"",
				"(u|v1)",
			},
		},
		{
			patterns: []string{
				"ي\u200e",
				"",
				"$",
				"(i|j)",
			},
		},
		{
			patterns: []string{
				"ي\u200e",
				"",
				"",
				"(i|j1)",
			},
		},
	},
	gencyrillic: []rule{
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
				"с",
				"",
				"с",
				"",
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
				"(hejmer|hajmer)",
			},
		},
		{
			patterns: []string{
				"гейм",
				"",
				"$",
				"(hejm|hajm)",
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
				"ей",
				"^",
				"",
				"(jej|ej)",
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
				"ej",
			},
		},
		{
			patterns: []string{
				"ей",
				"",
				"",
				"ej",
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
				"ё",
				"",
				"",
				"(e|jo)",
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
	genczech: []rule{
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
				"qu",
				"",
				"",
				"(k|kv)",
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
				"ei",
				"",
				"",
				"(ej|aj)",
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
				"č",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"š",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"ž",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"ň",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"ť",
				"",
				"",
				"(t|tj)",
			},
		},
		{
			patterns: []string{
				"ď",
				"",
				"",
				"(d|dj)",
			},
		},
		{
			patterns: []string{
				"ř",
				"",
				"",
				"(r|rZ)",
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
				"ý",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ě",
				"",
				"",
				"(e|je)",
			},
		},
		{
			patterns: []string{
				"ů",
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
				"(h|g)",
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
				"(k|kv)",
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
				"z",
			},
		},
	},
	gendutch: []rule{
		{
			patterns: []string{
				"ssj",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"sj",
				"",
				"",
				"S",
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
				"[eiy]",
				"ts",
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
				"ss",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"h",
				"[aeiouy]",
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
				"ou",
				"",
				"",
				"au",
			},
		},
		{
			patterns: []string{
				"ie",
				"",
				"",
				"(Q|i)",
			},
		},
		{
			patterns: []string{
				"uu",
				"",
				"",
				"(Q|u)",
			},
		},
		{
			patterns: []string{
				"ee",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"(Y|Yj)",
			},
		},
		{
			patterns: []string{
				"aa",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"oo",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"oe",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"ij",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: []string{
				"ui",
				"",
				"",
				"(Y|uj)",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"(ej|aj)",
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
				"i",
				"[aou]",
				"",
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
				"(g|x)",
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
				"(i|Q)",
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
				"(u|Q)",
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
				"(w|v)",
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
	genenglish: []rule{
		{
			patterns: []string{
				"�",
				"",
				"",
				"",
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
				"mc",
				"^",
				"",
				"mak",
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
				"x",
				"^",
				"",
				"z",
			},
		},
		{
			patterns: []string{
				"y",
				"^",
				"",
				"j",
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
				"yi",
				"^",
				"",
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
				"(aue|oue)",
			},
		},
		{
			patterns: []string{
				"ai",
				"",
				"",
				"(aj|ej|e)",
			},
		},
		{
			patterns: []string{
				"ay",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: []string{
				"a",
				"",
				"[^aeiou]e",
				"ej",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"(ej|aj|i)",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"(ej|aj|i)",
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
				"a",
				"",
				"",
				"(e|o|a)",
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
				"dZ",
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
				"(o|a)",
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
				"(u|a)",
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
				"(w|v)",
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
	genfrench: []rule{
		{
			patterns: []string{
				"lt",
				"u",
				"$",
				"(lt|)",
			},
		},
		{
			patterns: []string{
				"c",
				"n",
				"$",
				"(k|)",
			},
		},
		{
			patterns: []string{
				"d",
				"",
				"$",
				"(t|)",
			},
		},
		{
			patterns: []string{
				"g",
				"n",
				"$",
				"(k|)",
			},
		},
		{
			patterns: []string{
				"p",
				"",
				"$",
				"(p|)",
			},
		},
		{
			patterns: []string{
				"r",
				"e",
				"$",
				"(r|)",
			},
		},
		{
			patterns: []string{
				"t",
				"",
				"$",
				"(t|)",
			},
		},
		{
			patterns: []string{
				"z",
				"",
				"$",
				"(s|)",
			},
		},
		{
			patterns: []string{
				"ds",
				"",
				"$",
				"(ds|)",
			},
		},
		{
			patterns: []string{
				"ps",
				"",
				"$",
				"(ps|)",
			},
		},
		{
			patterns: []string{
				"rs",
				"e",
				"$",
				"(rs|)",
			},
		},
		{
			patterns: []string{
				"ts",
				"",
				"$",
				"(ts|)",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"$",
				"(s|)",
			},
		},
		{
			patterns: []string{
				"x",
				"u",
				"$",
				"(ks|)",
			},
		},
		{
			patterns: []string{
				"s",
				"[aeéèêiou]",
				"[^aeéèêiou]",
				"(s|)",
			},
		},
		{
			patterns: []string{
				"t",
				"[aeéèêiou]",
				"[^aeéèêiou]",
				"(t|)",
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
				"aill",
				"",
				"e",
				"aj",
			},
		},
		{
			patterns: []string{
				"ll",
				"",
				"e",
				"(l|j)",
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
				"m",
				"[aeiouy]",
				"[aeiouy]",
				"m",
			},
		},
		{
			patterns: []string{
				"m",
				"[aeiouy]",
				"",
				"(m|n)",
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
				"au",
				"",
				"",
				"(o|au)",
			},
		},
		{
			patterns: []string{
				"ai",
				"",
				"",
				"(e|aj)",
			},
		},
		{
			patterns: []string{
				"ay",
				"",
				"",
				"(e|aj)",
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
				"(oj|va)",
			},
		},
		{
			patterns: []string{
				"ei",
				"",
				"",
				"(aj|ej|e)",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"(aj|ej|e)",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"(ej|Y)",
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
				"(u|Q)",
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
	gengerman: []rule{
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
				"sch",
				"",
				"",
				"S",
			},
		},
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
				"(aj|ej)",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"(aj|ej)",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"(Yj|ej|aj|oj)",
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
	gengreek: []rule{
		{
			patterns: []string{
				"αυ",
				"",
				"$",
				"af",
			},
		},
		{
			patterns: []string{
				"αυ",
				"",
				"(κ|π|σ|τ|φ|θ|χ|ψ)",
				"af",
			},
		},
		{
			patterns: []string{
				"αυ",
				"",
				"",
				"av",
			},
		},
		{
			patterns: []string{
				"ευ",
				"",
				"$",
				"ef",
			},
		},
		{
			patterns: []string{
				"ευ",
				"",
				"(κ|π|σ|τ|φ|θ|χ|ψ)",
				"ef",
			},
		},
		{
			patterns: []string{
				"ευ",
				"",
				"",
				"ev",
			},
		},
		{
			patterns: []string{
				"ηυ",
				"",
				"$",
				"if",
			},
		},
		{
			patterns: []string{
				"ηυ",
				"",
				"(κ|π|σ|τ|φ|θ|χ|ψ)",
				"if",
			},
		},
		{
			patterns: []string{
				"ηυ",
				"",
				"",
				"iv",
			},
		},
		{
			patterns: []string{
				"ου",
				"",
				"",
				"u",
			},
		},
		{
			patterns: []string{
				"αι",
				"",
				"",
				"aj",
			},
		},
		{
			patterns: []string{
				"ει",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: []string{
				"οι",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: []string{
				"ωι",
				"",
				"",
				"oj",
			},
		},
		{
			patterns: []string{
				"ηι",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: []string{
				"υι",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"γγ",
				"(ε|ι|η|α|ο|ω|υ)",
				"(ε|ι|η)",
				"(nj|j)",
			},
		},
		{
			patterns: []string{
				"γγ",
				"",
				"(ε|ι|η)",
				"j",
			},
		},
		{
			patterns: []string{
				"γγ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"(ng|g)",
			},
		},
		{
			patterns: []string{
				"γγ",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"γκ",
				"^",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"γκ",
				"(ε|ι|η|α|ο|ω|υ)",
				"(ε|ι|η)",
				"(nj|j)",
			},
		},
		{
			patterns: []string{
				"γκ",
				"",
				"(ε|ι|η)",
				"j",
			},
		},
		{
			patterns: []string{
				"γκ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"(ng|g)",
			},
		},
		{
			patterns: []string{
				"γκ",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"γι",
				"",
				"(α|ο|ω|υ)",
				"j",
			},
		},
		{
			patterns: []string{
				"γι",
				"",
				"",
				"(gi|i)",
			},
		},
		{
			patterns: []string{
				"γε",
				"",
				"(α|ο|ω|υ)",
				"j",
			},
		},
		{
			patterns: []string{
				"γε",
				"",
				"",
				"(ge|je)",
			},
		},
		{
			patterns: []string{
				"κζ",
				"",
				"",
				"gz",
			},
		},
		{
			patterns: []string{
				"τζ",
				"",
				"",
				"dz",
			},
		},
		{
			patterns: []string{
				"σ",
				"",
				"(β|γ|δ|μ|ν|ρ)",
				"z",
			},
		},
		{
			patterns: []string{
				"μβ",
				"",
				"",
				"(mb|b)",
			},
		},
		{
			patterns: []string{
				"μπ",
				"^",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"μπ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"mb",
			},
		},
		{
			patterns: []string{
				"μπ",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"ντ",
				"^",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"ντ",
				"(ε|ι|η|α|ο|ω|υ)",
				"",
				"(nd|nt)",
			},
		},
		{
			patterns: []string{
				"ντ",
				"",
				"",
				"(nt|d)",
			},
		},
		{
			patterns: []string{
				"ά",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"έ",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ή",
				"",
				"",
				"(i|e)",
			},
		},
		{
			patterns: []string{
				"ί",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ό",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ύ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: []string{
				"ώ",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ΰ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: []string{
				"ϋ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: []string{
				"ϊ",
				"",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"α",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"β",
				"",
				"",
				"(v|b)",
			},
		},
		{
			patterns: []string{
				"γ",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"δ",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"ε",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ζ",
				"",
				"",
				"z",
			},
		},
		{
			patterns: []string{
				"η",
				"",
				"",
				"(i|e)",
			},
		},
		{
			patterns: []string{
				"ι",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"κ",
				"",
				"",
				"k",
			},
		},
		{
			patterns: []string{
				"λ",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"μ",
				"",
				"",
				"m",
			},
		},
		{
			patterns: []string{
				"ν",
				"",
				"",
				"n",
			},
		},
		{
			patterns: []string{
				"ξ",
				"",
				"",
				"ks",
			},
		},
		{
			patterns: []string{
				"ο",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"π",
				"",
				"",
				"p",
			},
		},
		{
			patterns: []string{
				"ρ",
				"",
				"",
				"r",
			},
		},
		{
			patterns: []string{
				"σ",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"ς",
				"",
				"",
				"s",
			},
		},
		{
			patterns: []string{
				"τ",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"υ",
				"",
				"",
				"(Q|i|u)",
			},
		},
		{
			patterns: []string{
				"φ",
				"",
				"",
				"f",
			},
		},
		{
			patterns: []string{
				"θ",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"χ",
				"",
				"",
				"x",
			},
		},
		{
			patterns: []string{
				"ψ",
				"",
				"",
				"ps",
			},
		},
		{
			patterns: []string{
				"ω",
				"",
				"",
				"o",
			},
		},
	},
	gengreeklatin: []rule{
		{
			patterns: []string{
				"au",
				"",
				"$",
				"af",
			},
		},
		{
			patterns: []string{
				"au",
				"",
				"[kpstfh]",
				"af",
			},
		},
		{
			patterns: []string{
				"au",
				"",
				"",
				"av",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"$",
				"ef",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"[kpstfh]",
				"ef",
			},
		},
		{
			patterns: []string{
				"eu",
				"",
				"",
				"ev",
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
				"gge",
				"[aeiouy]",
				"",
				"(nje|je)",
			},
		},
		{
			patterns: []string{
				"ggi",
				"[aeiouy]",
				"[aou]",
				"(nj|j)",
			},
		},
		{
			patterns: []string{
				"ggi",
				"[aeiouy]",
				"",
				"(ni|i)",
			},
		},
		{
			patterns: []string{
				"gge",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"ggi",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"gg",
				"[aeiouy]",
				"",
				"(ng|g)",
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
				"gk",
				"^",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"gke",
				"[aeiouy]",
				"",
				"(nje|je)",
			},
		},
		{
			patterns: []string{
				"gki",
				"[aeiouy]",
				"",
				"(ni|i)",
			},
		},
		{
			patterns: []string{
				"gke",
				"",
				"",
				"je",
			},
		},
		{
			patterns: []string{
				"gki",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"gk",
				"[aeiouy]",
				"",
				"(ng|g)",
			},
		},
		{
			patterns: []string{
				"gk",
				"",
				"",
				"g",
			},
		},
		{
			patterns: []string{
				"nghi",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: []string{
				"nghi",
				"",
				"",
				"(Ngi|Ni)",
			},
		},
		{
			patterns: []string{
				"nghe",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: []string{
				"nghe",
				"",
				"",
				"(Nje|Nge)",
			},
		},
		{
			patterns: []string{
				"ghi",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: []string{
				"ghi",
				"",
				"",
				"(gi|i)",
			},
		},
		{
			patterns: []string{
				"ghe",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: []string{
				"ghe",
				"",
				"",
				"(je|ge)",
			},
		},
		{
			patterns: []string{
				"ngh",
				"",
				"",
				"Ng",
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
				"ngi",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: []string{
				"ngi",
				"",
				"",
				"(Ngi|Ni)",
			},
		},
		{
			patterns: []string{
				"nge",
				"",
				"[aouy]",
				"Nj",
			},
		},
		{
			patterns: []string{
				"nge",
				"",
				"",
				"(Nje|Nge)",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: []string{
				"gi",
				"",
				"",
				"(gi|i)",
			},
		},
		{
			patterns: []string{
				"ge",
				"",
				"[aouy]",
				"j",
			},
		},
		{
			patterns: []string{
				"ge",
				"",
				"",
				"(je|ge)",
			},
		},
		{
			patterns: []string{
				"ng",
				"",
				"",
				"Ng",
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
				"i",
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
				"y",
				"[aeou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"yi",
				"",
				"[aeou]",
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
				"ch",
				"",
				"",
				"x",
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
				"dh",
				"",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"dj",
				"",
				"",
				"dZ",
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
				"th",
				"",
				"",
				"t",
			},
		},
		{
			patterns: []string{
				"kz",
				"",
				"",
				"gz",
			},
		},
		{
			patterns: []string{
				"tz",
				"",
				"",
				"dz",
			},
		},
		{
			patterns: []string{
				"s",
				"",
				"[bgdmnr]",
				"z",
			},
		},
		{
			patterns: []string{
				"mb",
				"",
				"",
				"(mb|b)",
			},
		},
		{
			patterns: []string{
				"mp",
				"^",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"mp",
				"[aeiouy]",
				"",
				"mp",
			},
		},
		{
			patterns: []string{
				"mp",
				"",
				"",
				"b",
			},
		},
		{
			patterns: []string{
				"nt",
				"^",
				"",
				"d",
			},
		},
		{
			patterns: []string{
				"nt",
				"[aeiouy]",
				"",
				"(nd|nt)",
			},
		},
		{
			patterns: []string{
				"nt",
				"",
				"",
				"(nt|d)",
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
				"óu",
				"",
				"",
				"u",
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
				"ý",
				"",
				"",
				"(i|Q|u)",
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
				"(b|v)",
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
				"x",
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
				"(j|Z)",
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
				"ο",
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
				"(i|Q|u)",
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
	genhebrew: []rule{
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
	genhungarian: []rule{
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
				"(aj|ej)",
			},
		},
		{
			patterns: []string{
				"ey",
				"",
				"",
				"(aj|ej)",
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
				"(ej|e)",
			},
		},
		{
			patterns: []string{
				"ely",
				"",
				"",
				"(ej|eli)",
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
				"ú",
				"",
				"",
				"u",
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
	genitalian: []rule{
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
				"h",
				"",
				"$",
				"",
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
	genlatvian: []rule{
		{
			patterns: []string{
				"č",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ģ",
				"",
				"",
				"(d|dj)",
			},
		},
		{
			patterns: []string{
				"ķ",
				"",
				"",
				"(t|tj)",
			},
		},
		{
			patterns: []string{
				"ļ",
				"",
				"",
				"l",
			},
		},
		{
			patterns: []string{
				"š",
				"",
				"",
				"S",
			},
		},
		{
			patterns: []string{
				"ņ",
				"",
				"",
				"(n|nj)",
			},
		},
		{
			patterns: []string{
				"ž",
				"",
				"",
				"Z",
			},
		},
		{
			patterns: []string{
				"ā",
				"",
				"",
				"a",
			},
		},
		{
			patterns: []string{
				"ē",
				"",
				"",
				"e",
			},
		},
		{
			patterns: []string{
				"ī",
				"",
				"",
				"i",
			},
		},
		{
			patterns: []string{
				"ū",
				"",
				"",
				"u",
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
				"ei",
				"",
				"",
				"ej",
			},
		},
		{
			patterns: []string{
				"io",
				"",
				"",
				"jo",
			},
		},
		{
			patterns: []string{
				"iu",
				"",
				"",
				"ju",
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
				"o",
				"",
				"",
				"o",
			},
		},
		{
			patterns: []string{
				"ui",
				"",
				"",
				"uj",
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
	genpolish: []rule{
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
				"[aeou]",
				"",
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
	genportuguese: []rule{
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
				"h",
				"",
				"$",
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
	genromanian: []rule{
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
				"qu",
				"",
				"",
				"k",
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
				"(x|h)",
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
				"z",
			},
		},
	},
	genrussian: []rule{
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
				"qu",
				"",
				"",
				"(kv|k)",
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
				"[aeou]",
				"",
				"j",
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
				"[aeiou]",
				"",
				"j",
			},
		},
		{
			patterns: []string{
				"ii",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"iy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"yy",
				"",
				"$",
				"i",
			},
		},
		{
			patterns: []string{
				"yi",
				"",
				"$",
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
	genspanish: []rule{
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
				"h",
				"",
				"$",
				"",
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
				"b",
				"",
				"",
				"B",
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
				"(x|Z)",
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
				"V",
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
				"(ks|gz|S)",
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
				"z",
				"",
				"",
				"(z|s)",
			},
		},
	},
	genturkish: []rule{
		{
			patterns: []string{
				"ç",
				"",
				"",
				"tS",
			},
		},
		{
			patterns: []string{
				"ğ",
				"",
				"",
				"",
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
				"ü",
				"",
				"",
				"Q",
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
				"ı",
				"",
				"",
				"(e|i|)",
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
				"dZ",
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
				"j",
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
					"[aA]",
					"",
				},
			},
			{
				patterns: []string{
					"a",
					"A",
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
					"j",
					"",
					"j",
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
					"vanden",
					"^",
					"",
					"(vanden|)",
				},
			},
			{
				patterns: []string{
					"vander",
					"^",
					"",
					"(vander|)",
				},
			},
			{
				patterns: []string{
					"van",
					"^",
					"[bp]",
					"(vam|[16])",
				},
			},
			{
				patterns: []string{
					"van",
					"^",
					"",
					"(van|[16])",
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
					"sen",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: []string{
					"sen",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: []string{
					"sEn",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: []string{
					"sEn",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: []string{
					"e",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"i",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"E",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"I",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Q",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Y",
					"[BbdfgklmnprsStvzZ]",
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
					"burk",
					"",
					"$",
					"(burk|berk)",
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
					"burg",
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
					"Burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"BUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"Burg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"BUrg",
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
							"mb",
							"",
							"",
							"(mb|b[512])",
						},
					},
					{
						patterns: []string{
							"mp",
							"",
							"",
							"(mp|b[512])",
						},
					},
					{
						patterns: []string{
							"ng",
							"",
							"",
							"(ng|g[512])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[fktSs]",
							"(p|f[262144])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"p",
							"",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"$",
							"(p|f[262144])",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"[pktSs]",
							"(f|p[262144])",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"f",
							"",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"$",
							"(f|p[262144])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"(b|v[262144])",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"",
							"(v|b[262144])",
						},
					},
					{
						patterns: []string{
							"t",
							"",
							"$",
							"(t|[64])",
						},
					},
					{
						patterns: []string{
							"g",
							"n",
							"$",
							"(g|[64])",
						},
					},
					{
						patterns: []string{
							"k",
							"n",
							"$",
							"(k|[64])",
						},
					},
					{
						patterns: []string{
							"p",
							"",
							"$",
							"(p|[64])",
						},
					},
					{
						patterns: []string{
							"r",
							"[Ee]",
							"$",
							"(r|[64])",
						},
					},
					{
						patterns: []string{
							"s",
							"",
							"$",
							"(s|[64])",
						},
					},
					{
						patterns: []string{
							"t",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(t|[64])",
						},
					},
					{
						patterns: []string{
							"s",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(s|[64])",
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
							"(Q[128]|i|D[32])",
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
							"(ik|Qk[128])",
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
							"(sits|sQts[128])",
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
							"(Q[128]|i)",
						},
					},
					{
						patterns: []string{
							"lEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32])",
						},
					},
					{
						patterns: []string{
							"rEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32])",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32]|lY[128])",
						},
					},
					{
						patterns: []string{
							"rE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32]|rY[128])",
						},
					},
					{
						patterns: []string{
							"EE",
							"",
							"",
							"(i|)",
						},
					},
					{
						patterns: []string{
							"ea",
							"",
							"",
							"(D|a|i)",
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
							"eu",
							"",
							"",
							"(D|e|u)",
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
							"(ia|io|iY[128])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[128]|D[32])",
						},
					},
					{
						patterns: []string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
						},
					},
					{
						patterns: []string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
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
							"(i|Y[128])",
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
							"(o|Y[128])",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"",
							"o",
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
							"(a|o|Y[128])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"",
							"(a|o)",
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
							"(uk|Qk[128])",
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
							"(suts|sQts[128])",
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
							"(u|Q[128])",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"",
							"u",
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
							"(i|Y[128])",
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
				langs: 1,
				rules: []rule{
					{
						patterns: []string{
							"1a",
							"",
							"",
							"(D|a)",
						},
					},
					{
						patterns: []string{
							"1i",
							"",
							"",
							"(D|i|e)",
						},
					},
					{
						patterns: []string{
							"1u",
							"",
							"",
							"(D|u|o)",
						},
					},
					{
						patterns: []string{
							"j1",
							"",
							"",
							"(ja|je|jo|ju|j)",
						},
					},
					{
						patterns: []string{
							"1",
							"",
							"",
							"(a|e|i|o|u|)",
						},
					},
					{
						patterns: []string{
							"u",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"i",
							"",
							"",
							"(i|e)",
						},
					},
					{
						patterns: []string{
							"p",
							"",
							"$",
							"p",
						},
					},
					{
						patterns: []string{
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
				langs: 6,
				rules: []rule{
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
				},
			},
			{
				langs: 3,
				rules: []rule{
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
				},
			},
			{
				langs: 4,
				rules: []rule{
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
				},
			},
			{
				langs: 5,
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
				langs: 7,
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
				langs: 8,
				rules: []rule{
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
				},
			},
			{
				langs: 9,
				rules: []rule{
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
				},
			},
			{
				langs: 12,
				rules: []rule{
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
				},
			},
			{
				langs: 13,
				rules: []rule{
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
				},
			},
			{
				langs: 14,
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
				langs: 15,
				rules: []rule{
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
				},
			},
			{
				langs: 16,
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
				langs: 18,
				rules: []rule{
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
							"B",
							"",
							"",
							"(b|v)",
						},
					},
					{
						patterns: []string{
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
					"[aA]",
					"",
				},
			},
			{
				patterns: []string{
					"a",
					"A",
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
					"j",
					"",
					"j",
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
					"vanden",
					"^",
					"",
					"(vanden|)",
				},
			},
			{
				patterns: []string{
					"vander",
					"^",
					"",
					"(vander|)",
				},
			},
			{
				patterns: []string{
					"van",
					"^",
					"[bp]",
					"(vam|[16])",
				},
			},
			{
				patterns: []string{
					"van",
					"^",
					"",
					"(van|[16])",
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
					"sen",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: []string{
					"sen",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: []string{
					"sEn",
					"[rmnl]",
					"$",
					"(zn|zon)",
				},
			},
			{
				patterns: []string{
					"sEn",
					"",
					"$",
					"(sn|son)",
				},
			},
			{
				patterns: []string{
					"e",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"i",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"E",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"I",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Q",
					"[BbdfgklmnprsStvzZ]",
					"[ln]$",
					"",
				},
			},
			{
				patterns: []string{
					"Y",
					"[BbdfgklmnprsStvzZ]",
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
					"burk",
					"",
					"$",
					"(burk|berk)",
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
					"burg",
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
					"Burk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"BUrk",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"Burg",
					"",
					"$",
					"(burk|berk)",
				},
			},
			{
				patterns: []string{
					"BUrg",
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
							"mb",
							"",
							"",
							"(mb|b[512])",
						},
					},
					{
						patterns: []string{
							"mp",
							"",
							"",
							"(mp|b[512])",
						},
					},
					{
						patterns: []string{
							"ng",
							"",
							"",
							"(ng|g[512])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"[fktSs]",
							"(p|f[262144])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"p",
							"",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"$",
							"(p|f[262144])",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"[pktSs]",
							"(f|p[262144])",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"f",
							"",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"$",
							"(f|p[262144])",
						},
					},
					{
						patterns: []string{
							"B",
							"",
							"",
							"(b|v[262144])",
						},
					},
					{
						patterns: []string{
							"V",
							"",
							"",
							"(v|b[262144])",
						},
					},
					{
						patterns: []string{
							"t",
							"",
							"$",
							"(t|[64])",
						},
					},
					{
						patterns: []string{
							"g",
							"n",
							"$",
							"(g|[64])",
						},
					},
					{
						patterns: []string{
							"k",
							"n",
							"$",
							"(k|[64])",
						},
					},
					{
						patterns: []string{
							"p",
							"",
							"$",
							"(p|[64])",
						},
					},
					{
						patterns: []string{
							"r",
							"[Ee]",
							"$",
							"(r|[64])",
						},
					},
					{
						patterns: []string{
							"s",
							"",
							"$",
							"(s|[64])",
						},
					},
					{
						patterns: []string{
							"t",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(t|[64])",
						},
					},
					{
						patterns: []string{
							"s",
							"[aeiouAEIOU]",
							"[^aeiouAEIOU]",
							"(s|[64])",
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
							"(Q[128]|i|D[32])",
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
							"(ik|Qk[128])",
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
							"(sits|sQts[128])",
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
							"(Q[128]|i)",
						},
					},
					{
						patterns: []string{
							"lEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32])",
						},
					},
					{
						patterns: []string{
							"rEE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32])",
						},
					},
					{
						patterns: []string{
							"lE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(li|il[32]|lY[128])",
						},
					},
					{
						patterns: []string{
							"rE",
							"[bdfgkmnprsStvzZ]",
							"",
							"(ri|ir[32]|rY[128])",
						},
					},
					{
						patterns: []string{
							"EE",
							"",
							"",
							"(i|)",
						},
					},
					{
						patterns: []string{
							"ea",
							"",
							"",
							"(D|a|i)",
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
							"eu",
							"",
							"",
							"(D|e|u)",
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
							"(ia|io|iY[128])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"[^aeiouAEBFIOU]e",
							"(a|o|Y[128]|D[32])",
						},
					},
					{
						patterns: []string{
							"E",
							"i[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
						},
					},
					{
						patterns: []string{
							"E",
							"a[^aeiouAEIOU]",
							"",
							"(i|Y[128]|[32])",
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
							"(i|Y[128])",
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
							"(o|Y[128])",
						},
					},
					{
						patterns: []string{
							"O",
							"",
							"",
							"o",
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
							"(a|o|Y[128])",
						},
					},
					{
						patterns: []string{
							"A",
							"",
							"",
							"(a|o)",
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
							"(uk|Qk[128])",
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
							"(suts|sQts[128])",
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
							"(u|Q[128])",
						},
					},
					{
						patterns: []string{
							"U",
							"",
							"",
							"u",
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
							"(i|Y[128])",
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
				langs: 1,
				rules: []rule{
					{
						patterns: []string{
							"1a",
							"",
							"",
							"(D|a)",
						},
					},
					{
						patterns: []string{
							"1i",
							"",
							"",
							"(D|i|e)",
						},
					},
					{
						patterns: []string{
							"1u",
							"",
							"",
							"(D|u|o)",
						},
					},
					{
						patterns: []string{
							"j1",
							"",
							"",
							"(ja|je|jo|ju|j)",
						},
					},
					{
						patterns: []string{
							"1",
							"",
							"",
							"(a|e|i|o|u|)",
						},
					},
					{
						patterns: []string{
							"u",
							"",
							"",
							"(o|u)",
						},
					},
					{
						patterns: []string{
							"i",
							"",
							"",
							"(i|e)",
						},
					},
					{
						patterns: []string{
							"p",
							"",
							"$",
							"p",
						},
					},
					{
						patterns: []string{
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
				langs: 6,
				rules: []rule{
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
				},
			},
			{
				langs: 3,
				rules: []rule{
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
				},
			},
			{
				langs: 4,
				rules: []rule{
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
				},
			},
			{
				langs: 5,
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
				langs: 7,
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
				langs: 8,
				rules: []rule{
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
				},
			},
			{
				langs: 9,
				rules: []rule{
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
				},
			},
			{
				langs: 12,
				rules: []rule{
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
				},
			},
			{
				langs: 13,
				rules: []rule{
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
				},
			},
			{
				langs: 14,
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
				langs: 15,
				rules: []rule{
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
				},
			},
			{
				langs: 16,
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
				langs: 18,
				rules: []rule{
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
							"B",
							"",
							"",
							"(b|v)",
						},
					},
					{
						patterns: []string{
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
