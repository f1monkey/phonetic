//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Println(rules)
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
	Languages  []string          `json:"languages"`
	Rules      map[string][]Rule `json:"rules"`
	FinalRules FinalRules        `json:"finalRules"`
	LangRules  []LangRule        `json:"langRules"`
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
