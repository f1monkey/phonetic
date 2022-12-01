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
	"ash": "bmpm-php/ash.json",
	"gen": "bmpm-php/gen.json",
	"sep": "bmpm-php/sep.json",
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
		err = tpl.Execute(&buf, ruleSet)
		if err != nil {
			panic(err)
		}

		src, err := format.Source(buf.Bytes())
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile("rules_"+mode+".go", src, 0644); err != nil {
			panic(err)
		}
	}
}

type Rule struct {
	Pattens [4]string `json:"patterns"`
}

type SecondFinalRule struct {
	Langs int    `json:"langs"`
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
	Langs   int    `json:"langs"`
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
package bmpm

type {{.Mode}}Lang int


const (
	{{.Mode}}{{ (index .Languages 0) }} {{.Mode}}Lang = 1 << iota
	{{- range $i, $lang := .Languages}}
		{{- if ne $i 0 }}
			{{$.Mode}}{{ $lang }}
		{{- end}}
	{{- end}}
) 


`
