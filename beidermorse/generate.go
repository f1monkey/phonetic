//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/f1monkey/phonetic/internal/regexparser"
)

var ruleSources = map[string]string{
	"ash": "beidermorse/bmpm-rules/ash.json",
	"gen": "beidermorse/bmpm-rules/gen.json",
	"sep": "beidermorse/bmpm-rules/sep.json",
}

func main() {
	rules, err := loadRules()
	if err != nil {
		panic(err)
	}
	fixRules(rules)

	rulesTpl, err := template.New("rules").Funcs(template.FuncMap{
		"Title": strings.Title,
	}).Parse(rulesTemplate)
	if err != nil {
		panic(err)
	}

	encoderTpl, err := template.New("encoder").Parse(encoderTemplate)
	if err != nil {
		panic(err)
	}

	for mode, ruleSet := range rules {
		dir := "beidermorse"
		if mode != "gen" {
			dir = path.Join(dir, "beidermorse"+mode)
		}

		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}

		data := transformRuleSet(ruleSet)

		for filename, tpl := range map[string]*template.Template{
			"rules.go":   rulesTpl,
			"encoder.go": encoderTpl,
		} {

			var buf bytes.Buffer
			err := tpl.Execute(&buf, data)
			if err != nil {
				panic(err)
			}

			src := buf.Bytes()
			src, err = format.Source(src)
			if err != nil {
				panic(err)
			}

			if err := os.WriteFile(path.Join(dir, filename), src, 0644); err != nil {
				panic(err)
			}
		}
	}
}

func transformRuleSet(src SrcRuleSet) DestRuleSet {
	result := DestRuleSet{}

	result.Mode = src.Mode
	result.Languages = src.Languages

	result.Rules = make(map[string][]DestRule, len(src.Rules))
	for code, rules := range src.Rules {
		result.Rules[code] = transformRules(rules)
	}

	result.FinalRules = DestFinalRules{
		Approx: transformFinalRule(src.FinalRules.Approx),
		Exact:  transformFinalRule(src.FinalRules.Exact),
	}

	result.LangRules = make([]DestLangRule, len(src.LangRules))
	for i, lr := range src.LangRules {
		r := DestLangRule{
			Langs:  lr.Langs,
			Accept: lr.Accept,
			Match:  transformPattern(lr.Pattern),
		}

		result.LangRules[i] = r
	}

	result.Discards = src.Discards

	return result
}

func transformFinalRule(src SrcFinalRule) DestFinalRule {
	result := DestFinalRule{
		First:  transformRules(src.First),
		Second: make([]DestSecondFinalRule, len(src.Second)),
	}

	for i, sr := range src.Second {
		result.Second[i] = DestSecondFinalRule{
			Lang:  sr.Lang,
			Rules: transformRules(sr.Rules),
		}
	}

	return result
}

func transformRules(src []SrcRule) []DestRule {
	result := make([]DestRule, len(src))
	for i, s := range src {
		result[i] = transformRule(s)
	}
	return result
}

func transformRule(src SrcRule) DestRule {
	result := DestRule{
		Pattern:       src.Patterns[0],
		Phonetic:      src.Patterns[3],
		PhoneticRules: transformRulesPhonetic(src.Patterns[3]),
	}

	l := ""
	if src.Patterns[1] != "" {
		l = src.Patterns[1] + "$"
	}
	if lm := transformPattern(l); !lm.IsEmpty() {
		result.LeftContext = &lm
	}

	r := ""
	if src.Patterns[2] != "" {
		r = "^" + src.Patterns[2]
	}
	if rm := transformPattern(r); !rm.IsEmpty() {
		result.RightContext = &rm
	}

	return result
}

func transformRulesPhonetic(src string) []DestRulePhonetic {
	if strings.Index(src, "(") == -1 {
		return []DestRulePhonetic{transformRulePhonetic(src)}
	}

	if strings.Index(src, ")") == -1 {
		panic(fmt.Errorf("invalid rule %q: must contain closing bracket ')'", src))
	}

	src = strings.Trim(src, "()")
	parts := strings.Split(src, "|")
	result := make([]DestRulePhonetic, len(parts))
	for i, part := range parts {
		result[i] = transformRulePhonetic(part)
	}
	return result
}

func transformRulePhonetic(src string) DestRulePhonetic {
	langs := -1 // apply to all langs
	if start := strings.Index(src, "["); start != -1 {
		end := strings.Index(src, "]")
		if end == -1 {
			panic(fmt.Errorf("invalid rule %q: must contain closing bracket ']'", src))
		}

		l, err := strconv.Atoi(src[start+1 : end])
		if err != nil {
			panic(fmt.Errorf("invalid rule %q: unable to parse lang attibute as integer", src))
		}
		langs = l
		src = src[0:start]
	}

	return DestRulePhonetic{
		Text:  src,
		Langs: langs,
	}
}

func transformPattern(pattern string) DestRuleMatch {
	r := DestRuleMatch{}

	parsed, ok := regexparser.Parse(pattern)
	if !ok {
		r.Pattern = pattern
		return r
	}

	if len(parsed.Items) == 0 {
		r.MatchEmptyString = true
		return r
	}

	if parsed.IsPrefix {
		r.Prefix = parsed.Items
		return r
	}

	if parsed.IsSuffix {
		r.Suffix = parsed.Items
		return r
	}

	if parsed.IsExact {
		r.Exact = parsed.Items
		return r
	}

	r.Contains = parsed.Items
	return r
}

func loadRules() (map[string]SrcRuleSet, error) {
	result := make(map[string]SrcRuleSet)
	for mode, filename := range ruleSources {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		var rules SrcRuleSet
		if err := json.Unmarshal(data, &rules); err != nil {
			return nil, fmt.Errorf("%q unmarshal err: %w", mode, err)
		}
		result[mode] = rules
	}

	return result, nil
}

func fixRules(rules map[string]SrcRuleSet) {
	// remove "/" in lang rules' patterns
	for mode, ruleSet := range rules {
		for i := range ruleSet.LangRules {
			ruleSet.LangRules[i].Pattern = strings.Trim(ruleSet.LangRules[i].Pattern, "/")
		}

		ruleSet.Mode = mode
		rules[mode] = ruleSet
	}
}

const rulesTemplate = `
{{- define "ruletpl" }}
{
	Pattern: []rune({{ printf "%q" .Pattern }}),
	{{- if ne .LeftContext nil}}
		LeftContext: {{- template "rulematchertpl" .LeftContext }},
	{{- end }}
	{{- if ne .RightContext nil}}
		RightContext: {{- template "rulematchertpl" .RightContext }},
	{{- end }}
	Phonetic: []bmpm.RuleToken{
		{{- range $i, $p := .PhoneticRules }}
			{
				{{- if ne $p.Text ""}}
					Text: []rune({{ printf "%q" $p.Text }}),
				{{- else }}
					Text: nil,
				{{- end }}
				Langs: {{ $p.Langs }},
			},
		{{- end }}
	},
},
{{- end }}
{{- define "rulematchertpl" }}
	&bmpm.Matcher{
		MatchEmptyString: {{ .MatchEmptyString }},
		{{- if .Contains }}
			Contains: [][]rune{
				{{- range $p := .Contains }}
					[]rune({{ printf "%q" $p }}),
				{{- end }}
			},
		{{- end }}
		{{- if .Exact }}
			Exact: [][]rune{
				{{- range $p := .Exact }}
					[]rune({{ printf "%q" $p }}),
				{{- end }}
			},
		{{- end }}
		{{- if .Prefix }}
			Prefix: [][]rune{
				{{- range $p := .Prefix }}
					[]rune({{ printf "%q" $p }}),
				{{- end }}
			},
		{{- end }}
		{{- if .Suffix }}
			Suffix: [][]rune{
				{{- range $p := .Suffix }}
					[]rune({{ printf "%q" $p }}),
				{{- end }}
			},
		{{- end }}
		{{- if ne .Pattern ""}}
			Pattern: regexp.MustCompile({{ printf "%q" .Pattern }}),
		{{- end }}
	}
{{- end }}

// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorse{{- if ne .Mode "gen" }}{{ .Mode }}{{- end }}

import (
	"regexp"
	"github.com/f1monkey/phonetic/internal/bmpm"
)

type Lang bmpm.Lang

const (
	{{ (index .Languages 0) | Title }} Lang = 1 << iota
	{{- range $i, $lang := .Languages }}
		{{- if ne $i 0 }}
			{{ $lang | Title }}
		{{- end }}
	{{- end }}
)

func (l Lang) String() string {
	switch l {
		{{- range $i, $lang := .Languages }}
	case {{ $lang | Title }}:
		return {{ printf "%q" $lang }}
		{{- end }}
	}
	return ""
}

const All = {{- range $i, $lang := .Languages }}
	{{- if ne $i 0 }}
		{{- if ne $i 1 }}+{{- end }}
		{{ $lang | Title }}
	{{- end }}
{{- end}}

var Rules = map[bmpm.Lang]bmpm.Rules{
	{{- range $lang, $rules := .Rules }}
		bmpm.Lang({{ $lang | Title }}): {
			{{- range $rule := $rules }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
	{{- end }}
}

var LangRules = []bmpm.LangRule{
	{{- range $rule := .LangRules }}
		{
			Matcher: {{- template "rulematchertpl" $rule.Match }},
			Langs: {{ $rule.Langs }},
			Accept: {{ $rule.Accept }},
		},
	{{- end }}
}

var FinalRules = bmpm.FinalRules{
	Approx: bmpm.FinalRule{
		First: bmpm.Rules{
			{{- range $rule := .FinalRules.Approx.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		Second: map[bmpm.Lang]bmpm.Rules{
			{{- range $secRule := .FinalRules.Approx.Second }}
				bmpm.Lang({{ (index $.Languages $secRule.Lang) | Title }}): {
					{{- range $rule := $secRule.Rules }}
						{{- template "ruletpl" $rule }}
					{{- end }}
				},
			{{- end }}
		},
	},
	Exact: bmpm.FinalRule{
		First: bmpm.Rules{
			{{- range $rule := .FinalRules.Exact.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		Second: map[bmpm.Lang]bmpm.Rules{
			{{- range $secRule := .FinalRules.Exact.Second }}
				bmpm.Lang({{ (index $.Languages $secRule.Lang) | Title }}): {
					{{- range $rule := $secRule.Rules }}
						{{- template "ruletpl" $rule }}
					{{- end }}
				},
			{{- end }}
		},
	},
}


var Discards = []string{
	{{- range $item := .Discards }}
		{{ printf "%q" $item }},
	{{- end }}
}
`

type SrcRule struct {
	Patterns [4]string `json:"patterns"`
}

type SrcSecondFinalRule struct {
	Lang  int64     `json:"lang"`
	Rules []SrcRule `json:"rules"`
}

type SrcFinalRule struct {
	First  []SrcRule            `json:"first"`
	Second []SrcSecondFinalRule `json:"second"`
}

type SrcFinalRules struct {
	Approx SrcFinalRule `json:"approx"`
	Exact  SrcFinalRule `json:"exact"`
}

type SrcLangRule struct {
	Pattern string `json:"pattern"`
	Langs   int64  `json:"langs"`
	Accept  bool   `json:"accept"`
}

type SrcRuleSet struct {
	Mode       string               `json:"-"`
	Languages  []string             `json:"languages"`
	Rules      map[string][]SrcRule `json:"rules"`
	FinalRules SrcFinalRules        `json:"finalRules"`
	LangRules  []SrcLangRule        `json:"langRules"`
	Discards   []string             `json:"discards"`
}

type DestRuleSet struct {
	Mode       string
	Languages  []string
	Rules      map[string][]DestRule
	FinalRules DestFinalRules
	LangRules  []DestLangRule
	Discards   []string
}

type DestRule struct {
	Pattern       string
	LeftContext   *DestRuleMatch
	RightContext  *DestRuleMatch
	Phonetic      string
	PhoneticRules []DestRulePhonetic
}

type DestRulePhonetic struct {
	Text  string
	Langs int
}

type DestLangRule struct {
	Match  DestRuleMatch
	Langs  int64
	Accept bool
}

type DestSecondFinalRule struct {
	Lang  int64      `json:"lang"`
	Rules []DestRule `json:"rules"`
}

type DestFinalRule struct {
	First  []DestRule            `json:"first"`
	Second []DestSecondFinalRule `json:"second"`
}

type DestFinalRules struct {
	Approx DestFinalRule `json:"approx"`
	Exact  DestFinalRule `json:"exact"`
}

type DestRuleMatch struct {
	MatchEmptyString bool
	Pattern          string
	Exact            []string
	Prefix           []string
	Suffix           []string
	Contains         []string
}

func (m DestRuleMatch) IsEmpty() bool {
	return len(m.Contains) == 0 &&
		len(m.Prefix) == 0 &&
		len(m.Suffix) == 0 &&
		len(m.Exact) == 0 &&
		m.Pattern == "" &&
		m.MatchEmptyString == false
}

const encoderTemplate = `
// THE FOLLOWING CODE WAS GENERATED USING "beidermorse/generate.go" COMMAND.
// DO NOT EDIT!
package beidermorse{{- if ne .Mode "gen" }}{{ .Mode }}{{- end }}

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
	accuracy Accuracy
	lang Lang
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
		input,
		{{- if eq .Mode "gen" }}bmpm.Generic,{{- end }}
		{{- if eq .Mode "ash" }}bmpm.Ashkenazi,{{- end }}
		{{- if eq .Mode "sep" }}bmpm.Sephardic,{{- end }}
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
		result = append(result,  string(s))
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
`
