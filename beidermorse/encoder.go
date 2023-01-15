package beidermorse

import (
	"fmt"
)

var ErrInvalidMode = fmt.Errorf("invalid name mode")
var ErrInvalidRuleset = fmt.Errorf("invalid ruleset")

type Encoder struct {
	mode    Mode
	ruleset Ruleset
}

// NewEncoder create new encoder instance
func NewEncoder(opts ...EncoderOption) (*Encoder, error) {
	result := &Encoder{
		mode:    Generic,
		ruleset: Approx,
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
	tokens := makeTokens(input, e.mode, e.ruleset, detectLang(input, e.mode), false)

	result := make([]string, len(tokens))
	for i, t := range tokens {
		result[i] = t.text
	}

	return result
}

// SetOption set encoder option. Method is not safe for concurrent usage
func (e *Encoder) SetOption(opt EncoderOption) error {
	return opt(e)
}

// EncoderOption func to provide encoder configuration parameter
type EncoderOption func(e *Encoder) error

// WithNameMode Set encoder name mode
func WithNameMode(m Mode) EncoderOption {
	return func(e *Encoder) error {
		if !m.Valid() {
			return fmt.Errorf("%w: %q", ErrInvalidMode, m)
		}
		e.mode = m
		return nil
	}
}

// WithNameMode Set encoder ruleset
func WithRuleset(r Ruleset) EncoderOption {
	return func(e *Encoder) error {
		if !r.Valid() {
			return fmt.Errorf("%w: %q", ErrInvalidRuleset, r)
		}
		e.ruleset = r
		return nil
	}
}
