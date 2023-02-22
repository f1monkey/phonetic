// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorseash

import (
	"fmt"
	"math/bits"

	"github.com/f1monkey/phonetic/internal/bmpm"
	"github.com/f1monkey/phonetic/internal/exrunes"
)

var ErrInvalidMode = fmt.Errorf("invalid name mode")
var ErrInvalidAccuracy = fmt.Errorf("invalid accuracy value")

// Accuracy exact or approximate matching
type Accuracy bmpm.Accuracy

const (
	Exact  Accuracy = "exact"  // exact matching rules
	Approx Accuracy = "approx" // approx matching (results in more tokens)
)

func (a Accuracy) Valid() bool {
	return a == Exact || a == Approx
}

type Encoder struct {
	accuracy         Accuracy
	lang             Lang
	useBufferStorage bool
}

// NewEncoder create new encoder instance
func NewEncoder(opts ...EncoderOption) (*Encoder, error) {
	result := &Encoder{
		accuracy: Approx,
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
	lang := bmpm.Lang(e.lang)
	if lang == 0 {
		lang = langDetector(input)
	}

	main, final1, final2 := getRules(e.accuracy, lang)

	var buf *exrunes.Buffer
	if e.useBufferStorage {
		buf = exrunes.BufferGet(200)
		defer exrunes.BufferFree(buf)
	} else {
		buf = exrunes.NewBuffer(200)
	}

	tokens := bmpm.MakeTokens(
		input, bmpm.Ashkenazi,
		bmpm.Accuracy(e.accuracy),
		bmpm.Ruleset{Main: main, Final1: final1, Final2: final2, Discards: Discards, DetectLang: langDetector},
		lang,
		false,
		buf,
	)

	if tokens == nil {
		return nil
	}

	result := make([]string, 0, tokens.Len())
	tokens.Iterate(func(s []rune) bool {
		result = append(result, string(s))
		return true
	})

	return result
}

// SetOption set encoder option. Method is not safe for concurrent usage
func (e *Encoder) SetOption(opt EncoderOption) error {
	return opt(e)
}

// EncoderOption func to provide encoder configuration parameter
type EncoderOption func(e *Encoder) error

// WithAccuracy Set encoder accuracy
func WithAccuracy(a Accuracy) EncoderOption {
	return func(e *Encoder) error {
		if !a.Valid() {
			return fmt.Errorf("%w: %q", ErrInvalidAccuracy, a)
		}
		e.accuracy = a
		return nil
	}
}

// WithLang Set encoder default lang (see lang constants)
func WithLang(l Lang) EncoderOption {
	return func(e *Encoder) error {
		e.lang = l
		return nil
	}
}

// WithBufferReuse reuse buffers to reduce GC pressure
// Leads to an increase in constant memory consumption, especially under heavy loads
func WithBufferReuse(value bool) EncoderOption {
	return func(e *Encoder) error {
		e.useBufferStorage = value
		return nil
	}
}

func getRules(
	accuracy Accuracy,
	lang bmpm.Lang,
) (bmpm.Rules, bmpm.Rules, bmpm.Rules) {
	var main, final1, final2 bmpm.Rules

	langCount := bits.OnesCount64(uint64(lang))
	if langCount > 1 {
		lang = bmpm.Lang(Any)
	}
	main = Rules[lang]

	if accuracy == Approx {
		final1 = FinalRules.Approx.First
		final2 = FinalRules.Approx.Second[lang]
	} else if accuracy == Exact {
		final1 = FinalRules.Exact.First
		final2 = FinalRules.Exact.Second[lang]
	}

	return main, final1, final2
}

func detectLangFunc() bmpm.DetectLangFunc {
	return func(input string) bmpm.Lang {
		all := All
		rules := LangRules

		runes := []rune(input)
		remaining := bmpm.Lang(all)
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
