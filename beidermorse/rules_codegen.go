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
		err := tpl.Execute(&buf, ruleSet)
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

type Rule struct {
	Patterns [4]string `json:"patterns"`
}

type SecondFinalRule struct {
	Lang  uint64 `json:"lang"`
	Rules []Rule `json:"rules"`
}

type FinalRule struct {
	First  []Rule            `json:"first"`
	Second []SecondFinalRule `json:"second"`
}

type FinalRules struct {
	Approx FinalRule `json:"approx"`
	Exact  FinalRule `json:"exact"`
}

type LangRule struct {
	Pattern string `json:"pattern"`
	Langs   uint64 `json:"langs"`
	Accept  bool   `json:"accept"`
}

type RuleSet struct {
	Mode       string            `json:"-"`
	Languages  []string          `json:"languages"`
	Rules      map[string][]Rule `json:"rules"`
	FinalRules FinalRules        `json:"finalRules"`
	LangRules  []LangRule        `json:"langRules"`
	Discards   []string          `json:"discards"`
}

func loadRules() (map[string]RuleSet, error) {
	result := make(map[string]RuleSet)
	for mode, filename := range ruleSources {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		var rules RuleSet
		if err := json.Unmarshal(data, &rules); err != nil {
			return nil, fmt.Errorf("%q unmarshal err: %w", mode, err)
		}
		result[mode] = rules
	}

	return result, nil
}

func fixRules(rules map[string]RuleSet) {
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
					patterns: [4]string{
					{{- range $pattern := $rule.Patterns }}
						{{ printf "%q" $pattern }},
					{{- end }}
					},
				},
			{{- end }}
		},
	{{- end }}
}

var {{ .Mode }}LangRules = []langRule{
	{{- range $rule := .LangRules }}
		{
			pattern: regexp.MustCompile({{ printf "%q" $rule.Pattern }}),
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
					patterns: [4]string{
					{{- range $pattern := $rule.Patterns }}
						{{ printf "%q" $pattern }},
					{{- end }}
					},
				},
			{{- end }}
		},
		second: map[uint64][]rule{
			{{- range $secRule := .FinalRules.Approx.Second }}
				uint64({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): []rule{
					{{- range $rule := $secRule.Rules }}
						{
							patterns: [4]string{
							{{- range $pattern := $rule.Patterns }}
								{{ printf "%q" $pattern }},
							{{- end }}
							},
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
					patterns: [4]string{
					{{- range $pattern := $rule.Patterns }}
						{{ printf "%q" $pattern }},
					{{- end }}
					},
				},
			{{- end }}
		},
		second: map[uint64][]rule{
			{{- range $secRule := .FinalRules.Exact.Second }}
				uint64({{ $.Mode }}{{ index $.Languages $secRule.Lang }}): []rule{
					{{- range $rule := $secRule.Rules }}
						{
							patterns: [4]string{
							{{- range $pattern := $rule.Patterns }}
								{{ printf "%q" $pattern }},
							{{- end }}
							},
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
