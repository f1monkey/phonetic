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

	tpl, err := template.New("rules").Parse(rulesTemplate)
	if err != nil {
		panic(err)
	}

	for mode, ruleSet := range rules {
		var buf bytes.Buffer
		err := tpl.Execute(&buf, transformRuleSet(ruleSet))
		if err != nil {
			panic(err)
		}

		src := buf.Bytes()
		src, err = format.Source(src)
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile("beidermorse/rules_"+mode+".go", src, 0644); err != nil {
			panic(err)
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
	pattern: []rune({{ printf "%q" .Pattern }}),
	{{- if ne .LeftContext nil}}
		leftContext: {{- template "rulematchertpl" .LeftContext }},
	{{- end }}
	{{- if ne .RightContext nil}}
		rightContext: {{- template "rulematchertpl" .RightContext }},
	{{- end }}
	phoneticRules: []token{
		{{- range $i, $p := .PhoneticRules }}
			{
				{{- if ne $p.Text ""}}
					text: []rune({{ printf "%q" $p.Text }}),
				{{- else }}
					text: nil,
				{{- end }}
				langs: {{ $p.Langs }},
			},
		{{- end }}
	},
},
{{- end }}
{{- define "rulematchertpl" }}
	&ruleMatcher{
		matchEmptyString: {{ .MatchEmptyString }},
		{{- if ne .Contains ""}}
			contains: []rune({{ printf "%q" .Contains }}),
		{{- end }}
		{{- if ne .Prefix ""}}
			prefix: []rune({{ printf "%q" .Prefix }}),
		{{- end }}
		{{- if ne .Suffix ""}}
			suffix: []rune({{ printf "%q" .Suffix }}),
		{{- end }}
		{{- if ne .Pattern ""}}
			pattern: regexp.MustCompile({{ printf "%q" .Pattern }}),
		{{- end }}
	}
{{- end }}

// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type {{ .Mode }}Lang languageID

const (
	{{.Mode}}{{ (index .Languages 0) }} {{ .Mode }}Lang = 1 << iota
	{{- range $i, $lang := .Languages }}
		{{- if ne $i 0 }}
			{{ $.Mode }}{{ $lang }}
		{{- end }}
	{{- end }}
)

func (l {{ .Mode }}Lang) String() string {
	switch l {
		{{- range $i, $lang := .Languages }}
	case {{ $.Mode }}{{ $lang }}:
		return {{ printf "%q" $lang }}
		{{- end }}
	}
	return ""
}

const {{ .Mode }}All = {{- range $i, $lang := .Languages }}
	{{- if ne $i 0 }}
		{{- if ne $i 1 }}+{{- end }}
		{{ $.Mode }}{{ $lang }}
	{{- end }}
{{- end}}

var {{ .Mode }}Rules = map[{{ .Mode }}Lang]rules{
	{{- range $lang, $rules := .Rules }}
		{{ $.Mode }}{{ $lang}}: rules{
			{{- range $rule := $rules }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
	{{- end }}
}

var {{ .Mode }}LangRules = []langRule{
	{{- range $rule := .LangRules }}
		{
			match: {{- template "rulematchertpl" $rule.Match }},
			langs: {{ $rule.Langs }},
			accept: {{ $rule.Accept }},
		},
	{{- end }}
}

var {{ .Mode }}FinalRules = finalRules{
	approx: finalRule{
		first: rules{
			{{- range $rule := .FinalRules.Approx.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		second: map[languageID]rules{
			{{- range $secRule := .FinalRules.Approx.Second }}
				languageID({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): rules{
					{{- range $rule := $secRule.Rules }}
						{{- template "ruletpl" $rule }}
					{{- end }}
				},
			{{- end }}
		},
	},
	exact: finalRule{
		first: rules{
			{{- range $rule := .FinalRules.Exact.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		second: map[languageID]rules{
			{{- range $secRule := .FinalRules.Exact.Second }}
				languageID({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): rules{
					{{- range $rule := $secRule.Rules }}
						{{- template "ruletpl" $rule }}
					{{- end }}
				},
			{{- end }}
		},
	},
}


var {{ .Mode }}Discards = []string{
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
