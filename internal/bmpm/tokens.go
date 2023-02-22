package bmpm

import (
	"github.com/f1monkey/phonetic/internal/exrunes"
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
	ID    int
	Langs Lang
}

type Tokens struct {
	Items []Token
	Buf   *exrunes.Buffer
}

func (t *Tokens) Len() int {
	return len(t.Items)
}

func (t *Tokens) Iterate(iter func(s []rune) bool) {
	for _, item := range t.Items {
		if !iter(t.Buf.Get(item.ID)) {
			return
		}
	}
}

func (t *Tokens) Add(data []rune, langs Lang) {
	t.Items = append(t.Items, Token{ID: t.Buf.Add(data), Langs: langs})
}

func (t *Tokens) MergeRules(rules []RuleToken, lang Lang) {
	if len(rules) == 0 {
		return
	}

	initialLen := t.Len()
	for i := 0; i < initialLen; i++ {
		for j := range rules {
			newLang := lang.merge(t.Items[i].Langs, rules[j].Langs)
			if newLang == LangInvalid {
				continue
			}

			copyID := t.Buf.Copy(t.Items[i].ID)
			t.Buf.Append(copyID, rules[j].Text)
			t.Items = append(t.Items, Token{ID: copyID, Langs: newLang})
		}
	}

	t.Items = slices.Delete(t.Items, 0, initialLen)
}

type tokenStr struct {
	text  string
	langs Lang
}

func (t *Tokens) Deduplicate() {
	uniq := make(map[tokenStr]struct{}, len(t.Items))

	for i := 0; i < len(t.Items); i++ {
		tt := tokenStr{string(t.Buf.Get(t.Items[i].ID)), t.Items[i].Langs}

		if _, ok := uniq[tt]; !ok {
			uniq[tt] = struct{}{}
			continue
		}

		t.Items = slices.Delete(t.Items, i, i+1)
		i--
	}
}
