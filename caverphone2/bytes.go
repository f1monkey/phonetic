package caverphone2

import (
	"golang.org/x/exp/slices"
)

func bytesReplacePrefix(src []byte, from []byte, to []byte) []byte {
	if len(src) == 0 || len(from) == 0 || len(from) > len(src) {
		return src
	}
	diffLen := len(from) - len(to)
	if diffLen < 0 {
		panic("unable to replace if len(to) is greater than len(from)")
	}

	matchCnt := 0
	for i := 0; i < len(src); i++ {
		if i > 0 && matchCnt == 0 {
			break // no match for prefix
		}

		if src[i] == from[matchCnt] {
			matchCnt++
			if matchCnt == len(from) {
				copy(src[i-len(from)+1:], to)
				if diffLen != 0 {
					src = slices.Delete(src, i-diffLen+1, i+1)
				}
				i -= diffLen
				matchCnt = 0
				break
			}
		} else {
			break
		}
	}

	return src
}

func bytesReplaceSuffix(src []byte, from []byte, to []byte) []byte {
	if len(src) == 0 || len(from) == 0 || len(from) > len(src) {
		return src
	}
	diffLen := len(from) - len(to)
	if diffLen < 0 {
		panic("unable to replace if len(to) is greater than len(from)")
	}

	matchCnt := 0
	start := len(src) - len(from)
	for i := start; i < len(src); i++ {
		if i > start && matchCnt == 0 {
			break // no match for prefix
		}

		if src[i] == from[matchCnt] {
			matchCnt++
			if matchCnt == len(from) {
				copy(src[i-len(from)+1:], to)
				if diffLen != 0 {
					src = slices.Delete(src, i-diffLen+1, i+1)
				}
				i -= diffLen
				matchCnt = 0
				break
			}
		} else {
			break
		}
	}

	return src
}

func bytesReplaceAll(src []byte, from []byte, to []byte) []byte {
	if len(src) == 0 || len(from) == 0 || len(from) > len(src) {
		return src
	}
	diffLen := len(from) - len(to)
	if diffLen < 0 {
		panic("unable to replace if len(to) is greater than len(from)")
	}

	matchCnt := 0
	for i := 0; i < len(src); i++ {
		if src[i] == from[matchCnt] {
			matchCnt++
			if matchCnt == len(from) {
				copy(src[i-len(from)+1:], to)
				if diffLen != 0 {
					src = slices.Delete(src, i-diffLen+1, i+1)
				}
				i -= diffLen
				matchCnt = 0
			}
		} else {
			matchCnt = 0
		}
	}
	return src
}

func bytesReplaceSequence(src []byte, sequence []byte, to []byte) []byte {
	if len(src) == 0 || len(sequence) == 0 || len(sequence) > len(src) {
		return src
	}
	diffLen := len(sequence) - len(to)
	if diffLen < 0 {
		panic("unable to replace if len(to) is greater than len(from)")
	}

	matchCnt := 0
	sequences := 0
	for i := 0; i < len(src); i++ {
		if src[i] == sequence[matchCnt] {
			matchCnt++
			if matchCnt == len(sequence) {
				sequences++
				matchCnt = 0
			}
		} else {
			if sequences > 0 {
				idx := i - len(sequence)*sequences

				copy(src[idx:], to)
				toDel := len(sequence)*sequences - len(to)
				if toDel != 0 {
					src = slices.Delete(src, i-toDel, i)
				}
				i -= diffLen
			}

			sequences = 0
			matchCnt = 0
		}
	}

	if sequences > 0 {
		i := len(src)
		idx := i - len(sequence)*sequences

		copy(src[idx:], to)
		toDel := len(sequence)*sequences - len(to)
		if toDel != 0 {
			src = slices.Delete(src, i-toDel, i)
		}
		i -= diffLen
	}

	return src
}
