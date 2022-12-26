package phonetic

import "github.com/f1monkey/phonetic/beidermorse"

func NewBeiderMorseEncoder() *beidermorse.Encoder {
	return beidermorse.NewEncoder()
}
