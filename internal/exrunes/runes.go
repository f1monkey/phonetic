package exrunes

// Contains checks if haystack contains passed substring
func Contains(haystack []rune, needle []rune) bool {
	return ContainsAt(haystack, needle, -1)
}

// HasPrefix checks if passed runes slice is a prefix of the haystack
func HasPrefix(haystack []rune, needle []rune) bool {
	return ContainsAt(haystack, needle, 0)
}

// HasPrefix checks if passed runes slice is a suffix of the haystack
func HasSuffix(haystack []rune, needle []rune) bool {
	return ContainsAt(haystack, needle, len(haystack)-len(needle))
}

// ContainsAt check if haystack has the provided substring in the specific position
func ContainsAt(haystack []rune, needle []rune, index int) bool {
	if len(needle) == 0 {
		return false
	}

	if len(haystack) < index-1 {
		return false
	}

	matchCnt := 0

	start := index
	if index < 0 {
		start = 0
	}
	for i := start; i < len(haystack); i++ {
		if needle[matchCnt] == haystack[i] {
			matchCnt++
		} else if index < 0 {
			matchCnt = 0
		} else {
			break
		}
		if matchCnt >= len(needle) {
			return true
		}

		continue
	}

	return false
}
