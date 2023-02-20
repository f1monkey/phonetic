package regexparser

import (
	"regexp"
	"strings"
)

type Result struct {
	IsPrefix bool
	IsSuffix bool
	IsExact  bool
	Items    []string
}

var (
	containsPrefix = regexp.MustCompile(`^\^`)
	containsSuffix = regexp.MustCompile(`\$$`)
	symbol         = regexp.MustCompile(`\p{L}`)
)

type state int

const (
	stateDefault state = iota
	stateCharacterClass
	stateCapturingGroup
)

// Parse transforms simple regular expressions of bmpm rules to all possible strings
// It is required to reduce allocation count while matching
func Parse(src string) (Result, bool) {
	pattern := src
	result := Result{}

	if pattern == "" {
		return result, false
	}

	if pattern == "^$" {
		return result, true
	}

	if containsPrefix.MatchString(pattern) {
		pattern, _ = strings.CutPrefix(pattern, "^")
		result.IsPrefix = true
	}

	if containsSuffix.MatchString(pattern) {
		pattern, _ = strings.CutSuffix(pattern, "$")
		result.IsSuffix = true
	}

	if result.IsPrefix && result.IsSuffix {
		result.IsExact = true
		result.IsPrefix = false
		result.IsSuffix = false
	}

	groups := [][]string{}
	var group []string
	str := ""
	appendGroup := func() {
		if str != "" {
			group = append(group, str)
			str = ""
		}
		if len(group) != 0 {
			groups = append(groups, group)
			group = nil
		}
	}
	curState := stateDefault
	for _, rr := range pattern {
		if rr == '[' {
			appendGroup()
			curState = stateCharacterClass
			continue
		}

		if rr == ']' {
			appendGroup()
			curState = stateDefault
			continue
		}

		if rr == '(' {
			appendGroup()
			curState = stateCapturingGroup
			continue
		}

		if rr == ')' {
			appendGroup()
			curState = stateDefault
			continue
		}

		if rr != '|' && !symbol.MatchString(string(rr)) {
			return result, false // unable to transfrom such regexp
		}

		switch curState {
		case stateDefault:
			str += string(rr)
		case stateCharacterClass:
			group = append(group, string(rr))
		case stateCapturingGroup:
			if rr == '|' {
				if str != "" {
					group = append(group, str)
					str = ""
				}
			} else {
				str += string(rr)
			}
		}
	}
	appendGroup()

	for i := 0; i < len(groups); i++ {
		group := groups[i]
		if len(result.Items) == 0 {
			result.Items = group
			continue
		}

		newItems := make([]string, 0, len(group)*len(result.Items))
		for _, item := range result.Items {
			for _, g := range group {
				newItems = append(newItems, item+g)
			}
		}
		result.Items = newItems
	}

	return result, true
}
