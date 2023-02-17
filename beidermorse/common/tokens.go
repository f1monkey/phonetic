package common

import (
	"golang.org/x/exp/slices"
)

type Lang int64

const (
	LangUnitialized Lang = -1
	LangInvalid     Lang = 0
	LangAny         Lang = 1
)

func (l Lang) merge(src ...Lang) Lang {
	if len(src) == 0 {
		return l
	}

	result := l
	for _, lang := range src {
		if result == LangUnitialized || result == LangAny {
			result = lang
			continue
		}
		result &= lang
	}

	return result
}

type Token struct {
	Text  []rune
	Langs Lang
}

type Tokens []Token

func (t Tokens) Merge(lang Lang, src Tokens) Tokens {
	if len(src) == 0 {
		return t
	}

	result := make([]Token, 0, len(t)*len(src))
	for _, r1 := range t {
		for _, r2 := range src {
			lang := lang.merge(r1.Langs, r2.Langs)
			if lang == LangInvalid {
				continue
			}

			text := make([]rune, 0, len(r1.Text)+len(r2.Text))
			text = append(text, r1.Text...)
			text = append(text, r2.Text...)

			result = append(result, Token{
				Text:  text,
				Langs: lang,
			})
		}
	}

	return result
}

type tokenStr struct {
	text  string
	langs Lang
}

func (t Tokens) Deduplicate() Tokens {
	uniq := make(map[tokenStr]struct{}, len(t))

	for i := 0; i < len(t); i++ {
		tt := tokenStr{string(t[i].Text), t[i].Langs}

		if _, ok := uniq[tt]; !ok {
			uniq[tt] = struct{}{}
			continue
		}

		t = slices.Delete(t, i, i+1)
		i--
	}

	return t
}
