package bmpm

import (
	"regexp"
	"sync"
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

func detectLang(word string, langRules []langRule, availLangs uint64) uint64 {
	remaining := availLangs

	for _, rule := range langRules {
		if !regCache.get(rule.pattern).MatchString(word) {
			continue
		}

		if rule.accept {
			remaining &= rule.langs
		} else {
			remaining &= (availLangs ^ rule.langs) % (remaining + 1)
		}

		if remaining == 0 {
			return remaining
		}
	}

	return remaining
}

var regCache = regexpCache{
	data: make(map[string]*regexp.Regexp),
}

type regexpCache struct {
	data map[string]*regexp.Regexp
	mtx  sync.RWMutex
}

func (c *regexpCache) get(pattern string) *regexp.Regexp {
	c.mtx.RLock()
	r, ok := c.data[pattern]
	c.mtx.RUnlock()
	if ok {
		return r
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()
	r = regexp.MustCompile(pattern)
	c.data[pattern] = r

	return r
}
