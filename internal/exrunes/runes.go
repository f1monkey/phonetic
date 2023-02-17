package exrunes

type Runes []rune

// // contains checks if runes contains passed substring
func (r Runes) Contains(search []rune) bool {
	return r.ContainsAt(search, -1)
}

// // checks if passed runes slice is a prefix of the current slice
func (r Runes) HasPrefix(search Runes) bool {
	return r.ContainsAt(search, 0)
}

func (r Runes) HasSuffix(search []rune) bool {
	return r.ContainsAt(search, len(r)-len(search))
}

// containsAt check if runes has the provided substring in the specific position
func (r Runes) ContainsAt(search Runes, index int) bool {
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
