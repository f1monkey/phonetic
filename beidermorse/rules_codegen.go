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
		Pattern:  src.Patterns[0],
		Phonetic: src.Patterns[3],
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
	pattern: {{ printf "%q" .Pattern }},
	{{- if ne .LeftContext nil}}
		leftContext: &ruleMatcher{
			matchEmptyString: {{ .LeftContext.MatchEmptyString }},
			contains: {{ printf "%q" .LeftContext.Contains }},
			prefix: {{ printf "%q" .LeftContext.Prefix }},
			suffix: {{ printf "%q" .LeftContext.Suffix }},
			{{- if ne .LeftContext.Pattern ""}}
				pattern: regexp.MustCompile({{ printf "%q" .LeftContext.Pattern }}),
			{{- end }}
		},
	{{- end }}
	{{- if ne .RightContext nil}}
		rightContext: &ruleMatcher{
			matchEmptyString: {{ .RightContext.MatchEmptyString }},
			contains: {{ printf "%q" .RightContext.Contains }},
			prefix: {{ printf "%q" .RightContext.Prefix }},
			suffix: {{ printf "%q" .RightContext.Suffix }},
			{{- if ne .RightContext.Pattern ""}}
				pattern: regexp.MustCompile({{ printf "%q" .RightContext.Pattern }}),
			{{- end }}
		},
	{{- end }}
	phonetic: {{ printf "%q" .Phonetic }},
},
{{- end }}

// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type {{ .Mode }}Lang uint64

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
	{{- if ne $i 0 }}+{{- end }}
	{{ $.Mode }}{{ $lang }}
{{- end}}

var {{ .Mode }}Rules = map[{{ .Mode }}Lang][]rule{
	{{- range $lang, $rules := .Rules }}
		{{ $.Mode }}{{ $lang}}: []rule{
			{{- range $rule := $rules }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
	{{- end }}
}

var {{ .Mode }}LangRules = []langRule{
	{{- range $rule := .LangRules }}
		{
			match: ruleMatcher{
				contains: {{ printf "%q" $rule.Match.Contains }},
				prefix: {{ printf "%q" $rule.Match.Prefix }},
				suffix: {{ printf "%q" $rule.Match.Suffix }},
				{{- if ne $rule.Match.Pattern ""}}
					pattern: regexp.MustCompile({{ printf "%q" $rule.Match.Pattern }}),
				{{- end }}
			},
			langs: {{ $rule.Langs }},
			accept: {{ $rule.Accept }},
		},
	{{- end }}
}

var {{ .Mode }}FinalRules = finalRules{
	approx: finalRule{
		first: []rule{
			{{- range $rule := .FinalRules.Approx.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		second: map[uint64][]rule{
			{{- range $secRule := .FinalRules.Approx.Second }}
				uint64({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): []rule{
					{{- range $rule := $secRule.Rules }}
						{{- template "ruletpl" $rule }}
					{{- end }}
				},
			{{- end }}
		},
	},
	exact: finalRule{
		first: []rule{
			{{- range $rule := .FinalRules.Exact.First }}
				{{- template "ruletpl" $rule }}
			{{- end }}
		},
		second: map[uint64][]rule{
			{{- range $secRule := .FinalRules.Exact.Second }}
				uint64({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): []rule{
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
	Lang  uint64    `json:"lang"`
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
	Langs   uint64 `json:"langs"`
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
	Pattern      string
	LeftContext  *DestRuleMatch
	RightContext *DestRuleMatch
	Phonetic     string
}

type DestLangRule struct {
	Match  DestRuleMatch
	Langs  uint64
	Accept bool
}

type DestSecondFinalRule struct {
	Lang  uint64     `json:"lang"`
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
