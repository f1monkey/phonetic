package bmpm

import (
	"regexp"
)

type Mode string

const (
	Generic   Mode = "generic"   // for general usage
	Ashkenazi Mode = "ashkenazi" // for ashkenazi names
	Sephardic Mode = "sephardic" // for sephardic names
)

type Result struct {
	Tokens []string
}

// // Encode encodes a word using generic rules
// func Encode(word string, lang Lang) Result {
// 	return encode(Generic, word, lang)
// }

// // EncodeAshkenazi encodes a word using ashkenazi rules
// func EncodeAshkenazi(word string, lang Lang) Result {
// 	return encode(Ashkenazi, word, lang)
// }

// // EncodeSephardic encodes a word using sephardic rules
// func EncodeSephardic(word string, lang Lang) Result {
// 	return encode(Sephardic, word, lang)
// }

// func encode(mode Mode, word string, lang Lang) Result {
// 	if lang == None {
// 		lang = detectLang(word, mode)
// 	}

// 	// todo
// 	return word
// }

func detectLang(word string, langRules []langRule, allLangs uint64) uint64 {
	remaining := allLangs

	for _, rule := range langRules {
		r := regexp.MustCompile(rule.pattern) // @todo regex cache
		if !r.MatchString(word) {
			continue
		}

		if rule.accept {
			remaining &= rule.langs
		} else {
			remaining &= (allLangs ^ rule.langs) % (remaining + 1)
		}

		if remaining == 0 {
			return remaining
		}
	}

	return remaining
}
