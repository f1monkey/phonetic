// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorseash

import (
	"fmt"
	"math/bits"

	"github.com/f1monkey/phonetic/beidermorse/common"
)

var ErrInvalidMode = fmt.Errorf("invalid name mode")
var ErrInvalidAccuracy = fmt.Errorf("invalid accuracy value")

type Encoder struct {
	accuracy common.Accuracy
}

// NewEncoder create new encoder instance
func NewEncoder(opts ...EncoderOption) (*Encoder, error) {
	result := &Encoder{
		accuracy: common.Approx,
	}

	for _, opt := range opts {
		if err := opt(result); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// MustNewEncoder create new encoder instance. Panics on error
func MustNewEncoder(opts ...EncoderOption) *Encoder {
	result, err := NewEncoder(opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// Encode transform a passed string to a slice of phonetic tokens
func (e *Encoder) Encode(input string) []string {
	langDetector := detectLangFunc()
	lang := langDetector(input)

	main, final1, final2 := getRules(e.accuracy, lang)

	tokens := common.MakeTokens(
		input, common.Ashkenazi,
		e.accuracy,
		common.Ruleset{Main: main, Final1: final1, Final2: final2, Discards: Discards, DetectLang: langDetector},
		lang,
		false,
	)

	result := make([]string, len(tokens))
	for i, t := range tokens {
		result[i] = string(t.Text)
	}

	return result
}

// SetOption set encoder option. Method is not safe for concurrent usage
func (e *Encoder) SetOption(opt EncoderOption) error {
	return opt(e)
}

// EncoderOption func to provide encoder configuration parameter
type EncoderOption func(e *Encoder) error

// WithAccuracy Set encoder accuracy
func WithAccuracy(r common.Accuracy) EncoderOption {
	return func(e *Encoder) error {
		if !r.Valid() {
			return fmt.Errorf("%w: %q", ErrInvalidAccuracy, r)
		}
		e.accuracy = r
		return nil
	}
}

func getRules(
	accuracy common.Accuracy,
	lang common.Lang,
) (common.Rules, common.Rules, common.Rules) {
	var main, final1, final2 common.Rules

	langCount := bits.OnesCount64(uint64(lang))
	if langCount > 1 {
		lang = common.Lang(Any)
	}
	main = Rules[lang]

	if accuracy == common.Approx {
		final1 = FinalRules.Approx.First
		final2 = FinalRules.Approx.Second[lang]
	} else if accuracy == common.Exact {
		final1 = FinalRules.Exact.First
		final2 = FinalRules.Exact.Second[lang]
	}

	return main, final1, final2
}

func detectLangFunc() common.DetectLangFunc {
	return func(input string) common.Lang {
		all := All
		rules := LangRules

		runes := common.Runes(input)
		remaining := common.Lang(all)
		for _, rule := range rules {
			if rule.Matcher == nil {
				continue
			}

			if !rule.Matcher.Match(runes) {
				continue
			}

			if rule.Accept {
				remaining &= rule.Langs
			} else {
				remaining &= ^rule.Langs % (remaining + 1)
			}

			if remaining == 0 {
				return remaining
			}
		}

		return remaining
	}
}
