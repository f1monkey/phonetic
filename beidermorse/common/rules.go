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

func (r Rules) Apply(input Tokens, lang Lang, ignoreLangs bool) Tokens {
	if len(r) == 0 {
		return input
	}

	var result Tokens
	for _, tok := range input {
		newTokens := Tokens{{Text: nil, Langs: tok.Langs}}
		for j := 0; j < len(tok.Text); {
			var (
				applied bool
				offset  int
			)

			for _, rr := range r {
				var tmp Tokens
				tmp, offset = rr.ApplyTo(tok.Text, j)
				if len(tmp) > 0 {
					applied = true
					if len(newTokens) == 0 {
						newTokens = tmp
					} else {
						newTokens = newTokens.Merge(lang, tmp)
					}
					break
				}
			}

			if !applied {
				for k := range newTokens {
					newTokens[k].Text = append(newTokens[k].Text, tok.Text[j])
				}
			}
			j += offset
		}

		result = append(result, newTokens...)
	}

	if ignoreLangs {
		for i := range result {
			result[i].Langs = LangAny
		}
	}

	return result.Deduplicate()
}

type Rule struct {
	Pattern      exrunes.Runes
	LeftContext  *Matcher
	RightContext *Matcher
	Phonetic     Tokens
}

func (r Rule) ApplyTo(input exrunes.Runes, position int) (result []Token, offset int) {
	patternLength := len(r.Pattern)
	inputLength := len(input)
	offset = 1

	if len(r.Phonetic) == 0 {
		return
	}

	if patternLength > inputLength-position {
		return
	}

	if !input.ContainsAt(r.Pattern, position) {
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
	Prefix           exrunes.Runes
	Suffix           exrunes.Runes
	Contains         exrunes.Runes
}

func (r Matcher) Match(str exrunes.Runes) bool {
	if r.MatchEmptyString && len(str) == 0 {
		return true
	}

	if len(r.Contains) > 0 {
		return str.Contains(r.Contains)
	}

	if len(r.Prefix) > 0 {
		return str.HasPrefix(r.Prefix)
	}

	if len(r.Suffix) > 0 {
		return str.HasSuffix(r.Suffix)
	}

	if r.Pattern != nil {
		return r.Pattern.MatchString(string(str)) // @todo optimize this
	}

	return false
}
