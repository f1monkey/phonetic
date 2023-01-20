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
	"regexp"
	"strconv"
	"strings"
	"text/template"
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

var (
	containsRegexp = regexp.MustCompile(`^\p{L}+$`)
	prefixRegexp   = regexp.MustCompile(`^\^\p{L}+$`)
	suffixRegexp   = regexp.MustCompile(`^\p{L}+\$$`)
)

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

	if pattern == "^$" {
		r.MatchEmptyString = true
	} else if containsRegexp.MatchString(pattern) {
		r.Contains = pattern
	} else if prefixRegexp.MatchString(pattern) {
		r.Prefix = strings.ReplaceAll(pattern, "^", "")
	} else if suffixRegexp.MatchString(pattern) {
		r.Suffix = strings.ReplaceAll(pattern, "$", "")
	} else {
		r.Pattern = pattern
	}

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
	Phonetic: []common.Token{
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
	&common.Matcher{
		MatchEmptyString: {{ .MatchEmptyString }},
		{{- if ne .Contains ""}}
			Contains: []rune({{ printf "%q" .Contains }}),
		{{- end }}
		{{- if ne .Prefix ""}}
			Prefix: []rune({{ printf "%q" .Prefix }}),
		{{- end }}
		{{- if ne .Suffix ""}}
			Suffix: []rune({{ printf "%q" .Suffix }}),
		{{- end }}
		{{- if ne .Pattern ""}}
			Pattern: regexp.MustCompile({{ printf "%q" .Pattern }}),
		{{- end }}
	}
{{- end }}

// GENERATED CODE. DO NOT EDIT!
package beidermorse{{- if ne .Mode "gen" }}{{ .Mode }}{{- end }}

import (
	"regexp"
	"github.com/f1monkey/phonetic/beidermorse/common"
)

type Lang common.Lang

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

var Rules = map[common.Lang]common.Rules{
	{{- range $lang, $rules := .Rules }}
		common.Lang({{ $lang | Title }}): {
			{{- range $rule := $rules }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
	{{- end }}
}

var LangRules = []common.LangRule{
	{{- range $rule := .LangRules }}
		{
			Matcher: {{- template "rulematchertpl" $rule.Match }},
			Langs: {{ $rule.Langs }},
			Accept: {{ $rule.Accept }},
		},
	{{- end }}
}

var FinalRules = common.FinalRules{
	Approx: common.FinalRule{
		First: common.Rules{
			{{- range $rule := .FinalRules.Approx.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		Second: map[common.Lang]common.Rules{
			{{- range $secRule := .FinalRules.Approx.Second }}
				common.Lang({{ (index $.Languages $secRule.Lang) | Title }}): {
					{{- range $rule := $secRule.Rules }}
						{{- template "ruletpl" $rule }}
					{{- end }}
				},
			{{- end }}
		},
	},
	Exact: common.FinalRule{
		First: common.Rules{
			{{- range $rule := .FinalRules.Exact.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		Second: map[common.Lang]common.Rules{
			{{- range $secRule := .FinalRules.Exact.Second }}
				common.Lang({{ (index $.Languages $secRule.Lang) | Title }}): {
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
	Prefix           string
	Suffix           string
	Contains         string
}

func (m DestRuleMatch) IsEmpty() bool {
	return m.Contains == "" &&
		m.Prefix == "" &&
		m.Suffix == "" &&
		m.Pattern == "" &&
		m.MatchEmptyString == false
}

const encoderTemplate = `
// GENERATED CODE. DO NOT EDIT!
package beidermorse{{- if ne .Mode "gen" }}{{ .Mode }}{{- end }}

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
		input,
		{{- if eq .Mode "gen" }}common.Generic,{{- end }}
		{{- if eq .Mode "ash" }}common.Ashkenazi,{{- end }}
		{{- if eq .Mode "sep" }}common.Sephardic,{{- end }}
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
`
