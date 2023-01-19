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

func (t tokens) merge(lang languageID, src tokens) tokens {
	if len(src) == 0 {
		return t
	}

	result := make([]token, 0, len(t)*len(src))
	for _, r1 := range t {
		for _, r2 := range src {
			lang := lang.merge(r1.langs, r2.langs)
			if lang == langsInvalid {
				continue
			}

			text := make([]rune, 0, len(r1.text)+len(r2.text))
			text = append(text, r1.text...)
			text = append(text, r2.text...)

			result = append(result, token{
				text:  text,
				langs: lang,
			})
		}
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
