package common

import (
	"regexp"

	"github.com/f1monkey/phonetic/internal/exrunes"
)

// Mode which name mode to use for matching
type Mode string

const (
	Generic   Mode = "generic"   // for general usage
	Ashkenazi Mode = "ashkenazi" // for ashkenazi names
	Sephardic Mode = "sephardic" // for sephardic names
)

func (m Mode) Valid() bool {
	return m == Generic || m == Ashkenazi || m == Sephardic
}

// Accuracy exact or approximate matching
type Accuracy string

const (
	Exact  Accuracy = "exact"  // exact matching rules
	Approx Accuracy = "approx" // approx matching (results in more tokens)
)

func (a Accuracy) Valid() bool {
	return a == Exact || a == Approx
}

type Rules []Rule

func (r Rules) Apply(input *Tokens, lang Lang, ignoreLangs bool) *Tokens {
	if len(r) == 0 {
		return input
	}

	result := &Tokens{Buf: input.Buf}
	for _, tok := range input.Items {
		item := input.Buf.Get(tok.ID)

		id := input.Buf.AddWithSpace(nil, len(item))
		tok := Token{ID: id, Langs: tok.Langs}
		newTokens := &Tokens{Buf: input.Buf, Items: []Token{tok}}
		for j := 0; j < len(item); {
			var (
				applied bool
				offset  int
			)

			for _, rr := range r {
				var tmp []RuleToken
				tmp, offset = rr.ApplyTo(item, j)
				if len(tmp) > 0 {
					applied = true
					if newTokens.Len() == 0 {
						for k := range tmp {
							newTokens.Add(tmp[k].Text, tmp[k].Langs)
						}
					} else {
						newTokens.MergeRules(tmp, lang)
					}
					break
				}
			}

			if !applied {
				for k := range newTokens.Items {
					newTokens.Buf.Append(newTokens.Items[k].ID, []rune{item[j]})
				}
			}
			j += offset
		}

		result.Items = append(result.Items, newTokens.Items...)
	}

	if ignoreLangs {
		for i := range result.Items {
			result.Items[i].Langs = LangAny
		}
	}

	result.Deduplicate()
	return result
}

type Rule struct {
	Pattern      []rune
	LeftContext  *Matcher
	RightContext *Matcher
	Phonetic     []RuleToken
}

type RuleToken struct {
	Text  []rune
	Langs Lang
}

func (r Rule) ApplyTo(input []rune, position int) (result []RuleToken, offset int) {
	patternLength := len(r.Pattern)
	inputLength := len(input)
	offset = 1

	if len(r.Phonetic) == 0 {
		return
	}

	if patternLength > inputLength-position {
		return
	}

	if !exrunes.ContainsAt(input, r.Pattern, position) {
		return
	}

	if r.RightContext != nil {
		if !r.RightContext.Match((input[position+patternLength:])) {
			return
		}
	}

	if r.LeftContext != nil {
		if !r.LeftContext.Match(input[:position]) {
			return
		}
	}

	result = r.Phonetic
	offset = patternLength

	return
}

type FinalRules struct {
	Approx FinalRule
	Exact  FinalRule
}

type FinalRule struct {
	First  Rules
	Second map[Lang]Rules
}

type LangRule struct {
	Matcher *Matcher
	Langs   Lang
	Accept  bool
}

type Matcher struct {
	MatchEmptyString bool
	Pattern          *regexp.Regexp
	Prefix           [][]rune
	Suffix           [][]rune
	Contains         [][]rune
	Exact            [][]rune
}

func (r Matcher) Match(str []rune) bool {
	if r.MatchEmptyString && len(str) == 0 {
		return true
	}

	if len(r.Contains) > 0 {
		for i := range r.Contains {
			if exrunes.Contains(str, r.Contains[i]) {
				return true
			}
		}
		return false
	}

	if len(r.Exact) > 0 {
		for i := range r.Exact {
			if exrunes.Equal(str, r.Exact[i]) {
				return true
			}
		}
		return false
	}

	if len(r.Prefix) > 0 {
		for i := range r.Prefix {
			if exrunes.HasPrefix(str, r.Prefix[i]) {
				return true
			}
		}
		return false
	}

	if len(r.Suffix) > 0 {
		for i := range r.Suffix {
			if exrunes.HasSuffix(str, r.Suffix[i]) {
				return true
			}
		}
		return false
	}

	if r.Pattern != nil {
		return r.Pattern.MatchString(string(str)) // @todo optimize this
	}

	return false
}
