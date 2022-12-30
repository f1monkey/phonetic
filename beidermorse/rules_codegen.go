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
	result.Rules = src.Rules
	result.FinalRules = src.FinalRules

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

func transformPattern(pattern string) DestRuleMatch {
	r := DestRuleMatch{}

	if containsRegexp.MatchString(pattern) {
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
				{
					pattern: {{ printf "%q" (index $rule.Patterns 0) }},
					{{- if ne (index $rule.Patterns 1) ""}}
						leftContext: regexp.MustCompile({{ printf "\"%s$\"" (index $rule.Patterns 1) }}),
					{{- end }}
					{{- if ne (index $rule.Patterns 2) ""}}
						rightContext: regexp.MustCompile({{ printf "\"^%s\"" (index $rule.Patterns 2) }}),
					{{- end }}
					phonetic: {{ printf "%q" (index $rule.Patterns 3) }},
				},
			{{- end }}
		},
	{{- end }}
}

var {{ .Mode }}LangRules = []langRule{
	{{- range $rule := .LangRules }}
		{
			match: ruleMatch{
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
				{
					pattern: {{ printf "%q" (index $rule.Patterns 0) }},
					{{- if ne (index $rule.Patterns 1) ""}}
						leftContext: regexp.MustCompile({{ printf "\"%s$\"" (index $rule.Patterns 1) }}),
					{{- end }}
					{{- if ne (index $rule.Patterns 2) ""}}
						rightContext: regexp.MustCompile({{ printf "\"^%s\"" (index $rule.Patterns 2) }}),
					{{- end }}
					phonetic: {{ printf "%q" (index $rule.Patterns 3) }},
				},
			{{- end }}
		},
		second: map[uint64][]rule{
			{{- range $secRule := .FinalRules.Approx.Second }}
				uint64({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): []rule{
					{{- range $rule := $secRule.Rules }}
						{
							pattern: {{ printf "%q" (index $rule.Patterns 0) }},
							{{- if ne (index $rule.Patterns 1) ""}}
								leftContext: regexp.MustCompile({{ printf "\"%s$\"" (index $rule.Patterns 1) }}),
							{{- end }}
							{{- if ne (index $rule.Patterns 2) ""}}
								rightContext: regexp.MustCompile({{ printf "\"^%s\"" (index $rule.Patterns 2) }}),
							{{- end }}
							phonetic: {{ printf "%q" (index $rule.Patterns 3) }},
						},
					{{- end }}
				},
			{{- end }}
		},
	},
	exact: finalRule{
		first: []rule{
			{{- range $rule := .FinalRules.Exact.First }}
				{
					pattern: {{ printf "%q" (index $rule.Patterns 0) }},
					{{- if ne (index $rule.Patterns 1) ""}}
						leftContext: regexp.MustCompile({{ printf "\"%s$\"" (index $rule.Patterns 1) }}),
					{{- end }}
					{{- if ne (index $rule.Patterns 2) ""}}
						rightContext: regexp.MustCompile({{ printf "\"^%s\"" (index $rule.Patterns 2) }}),
					{{- end }}
					phonetic: {{ printf "%q" (index $rule.Patterns 3) }},
				},
			{{- end }}
		},
		second: map[uint64][]rule{
			{{- range $secRule := .FinalRules.Exact.Second }}
				uint64({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): []rule{
					{{- range $rule := $secRule.Rules }}
						{
							pattern: {{ printf "%q" (index $rule.Patterns 0) }},
							{{- if ne (index $rule.Patterns 1) ""}}
								leftContext: regexp.MustCompile({{ printf "\"%s$\"" (index $rule.Patterns 1) }}),
							{{- end }}
							{{- if ne (index $rule.Patterns 2) ""}}
								rightContext: regexp.MustCompile({{ printf "\"^%s\"" (index $rule.Patterns 2) }}),
							{{- end }}
							phonetic: {{ printf "%q" (index $rule.Patterns 3) }},
						},
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
	Rules      map[string][]SrcRule
	FinalRules SrcFinalRules
	LangRules  []DestLangRule
	Discards   []string
}

type DestLangRule struct {
	Match  DestRuleMatch
	Langs  uint64
	Accept bool
}

type DestRuleMatch struct {
	Pattern  string
	Prefix   string
	Suffix   string
	Contains string
}
