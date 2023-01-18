package beidermorse

import "golang.org/x/exp/slices"

type languageID int64

const (
	langsUnitialized languageID = -1
	langsInvalid     languageID = 0
	langsAny         languageID = 1
)

func (l languageID) merge(src ...languageID) languageID {
	if len(src) == 0 {
		return l
	}

	result := l
	for _, lang := range src {
		if result == langsUnitialized || result == langsAny {
			result = lang
			continue
		}
		result &= lang
	}

	return result
}

type token struct {
	text  runes
	langs languageID
}

type tokens []token

func (t tokens) merge(lang languageID, src ...tokens) tokens {
	if len(src) == 0 {
		return t
	}

	result := t
	i := 1
	for i < len(src)+1 { // use while-loop instead of recursion here
		newResult := make([]token, 0, len(result)*len(src[i-1]))
		for _, r1 := range result {
			for _, r2 := range src[i-1] {
				lang := lang.merge(r1.langs, r2.langs)
				if lang == langsInvalid {
					continue
				}

				newResult = append(newResult, token{
					text:  append(r1.text, r2.text...),
					langs: lang,
				})
			}
		}
		result = newResult
		i++
	}

	return result
}

type tokenStr struct {
	text  string
	langs languageID
}

func (t tokens) deduplicate() tokens {
	uniq := make(map[tokenStr]struct{}, len(t))

	for i := 0; i < len(t); i++ {
		tt := tokenStr{string(t[i].text), t[i].langs}

		if _, ok := uniq[tt]; !ok {
			uniq[tt] = struct{}{}
			continue
		}

		t = slices.Delete(t, i, i+1)
		i--
	}

	return t
}
