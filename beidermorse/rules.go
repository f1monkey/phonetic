package beidermorse

import (
	"regexp"
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

// Ruleset exact or approximate matching
type Ruleset string

const (
	Exact  Ruleset = "exact"  // exact matching rules
	Approx Ruleset = "approx" // approx matching (results in more tokens)
)

func (r Ruleset) Valid() bool {
	return r == Exact || r == Approx
}

type rules []rule

func (r rules) apply(input tokens, lang languageID, ignoreLangs bool) tokens {
	if len(r) == 0 {
		return input
	}

	var result tokens
	for _, tok := range input {
		tokenRunes := tok.text
		newTokens := tokens{{text: nil, langs: tok.langs}}
		for j := 0; j < len(tokenRunes); {
			var (
				applied bool
				offset  int
			)

			for _, rr := range r {
				var tmp tokens
				tmp, offset = rr.applyTo(tokenRunes, j)
				if len(tmp) > 0 {
					applied = true
					if len(newTokens) == 0 {
						newTokens = tmp
					} else {
						newTokens = newTokens.merge(lang, tmp)
					}
					break
				}
			}

			if !applied {
				for k := range newTokens {
					newTokens[k].text = append(newTokens[k].text, tokenRunes[j])
				}
			}
			j += offset
		}

		result = append(result, newTokens...)
	}

	if ignoreLangs {
		for i := range result {
			result[i].langs = langsAny
		}
	}

	return result.deduplicate()
}

type rule struct {
	pattern       runes
	leftContext   *ruleMatcher
	rightContext  *ruleMatcher
	phoneticRules tokens
}

func (r rule) applyTo(input runes, position int) (result []token, offset int) {
	patternLength := len(r.pattern)
	inputLength := len(input)
	offset = 1

	if len(r.phoneticRules) == 0 {
		return
	}

	if patternLength > inputLength-position {
		return
	}

	if !input.containsAt(r.pattern, position) {
		return
	}

	if r.rightContext != nil {
		if !r.rightContext.matches((input[position+patternLength:])) {
			return
		}
	}

	if r.leftContext != nil {
		if !r.leftContext.matches(input[:position+1]) {
			return
		}
	}

	result = r.phoneticRules
	offset = patternLength

	return
}

type finalRules struct {
	approx finalRule
	exact  finalRule
}

type finalRule struct {
	first  rules
	second map[languageID]rules
}

type langRule struct {
	match  *ruleMatcher
	langs  languageID
	accept bool
}

type ruleMatcher struct {
	matchEmptyString bool
	pattern          *regexp.Regexp
	prefix           runes
	suffix           runes
	contains         runes
}

func (r ruleMatcher) matches(str runes) bool {
	if r.matchEmptyString && len(str) == 0 {
		return true
	}

	if len(r.contains) > 0 {
		return str.contains(r.contains)
	}

	if len(r.prefix) > 0 {
		return str.hasPrefix(r.prefix)
	}

	if len(r.suffix) > 0 {
		return str.hasSuffix(r.suffix)
	}

	if r.pattern != nil {
		return r.pattern.MatchString(string(str)) // @todo optimize this
	}

	return false
}
