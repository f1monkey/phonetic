package beidermorse

type runes []rune

// // contains checks if runes contains passed substring
func (r runes) contains(search []rune) bool {
	return r.containsAt(search, -1)
}

// // checks if passed runes slice is a prefix of the current slice
func (r runes) hasPrefix(search runes) bool {
	return r.containsAt(search, 0)
}

func (r runes) hasSuffix(search []rune) bool {
	return r.containsAt(search, len(r)-len(search))
}

// containsAt check if runes has the provided substring in the specific position
func (r runes) containsAt(search runes, index int) bool {
	if len(search) == 0 {
		return false
	}

	if len(r) < index-1 {
		return false
	}

	matchCnt := 0

	start := index
	if index < 0 {
		start = 0
	}
	for i := start; i < len(r); i++ {
		if search[matchCnt] == r[i] {
			matchCnt++
		} else if index < 0 {
			matchCnt = 0
		} else {
			break
		}
		if matchCnt >= len(search) {
			return true
		}

		continue
	}

	return false
}
