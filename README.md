# Phonetic

Set of different phonetic encoders' implementations.

Provided encoders:
* [Soundex](#soundex)
* [Cologne phonetics](#cologne-phonetics)
* [Caverphone2](#caverphone2)
* [Beider-Morse](#beider-morse)

## Installion

To install:
```
$ go get -v github.com/f1monkey/phonetic
```

## Usage

### Soundex
The fastest algorithm in this library. Soundex is used to encode words into a phonetic code for matching similar sounding words with different spellings. It was developed for indexing English language names. [Wiki page](https://en.wikipedia.org/wiki/Soundex).

Code example:
```go
package main

import (
	"fmt"

	"github.com/f1monkey/phonetic/soundex"
)

func main() {
	e := soundex.NewEncoder()
	result := e.Encode("orange")
	fmt.Println(result)
	// prints: O652
}
```

### Cologne phonetics
Cologne phonetics (Kölner Phonetik) is a phonetic algorithm used for indexing German words by their sound, allowing for name and word matching in German language databases. [Wiki page](https://en.wikipedia.org/wiki/Cologne_phonetics)

Code example:

```go
package main

import (
	"fmt"

	"github.com/f1monkey/phonetic/cologne"
)

func main() {
	e := cologne.NewEncoder()
	result := e.Encode("Großtraktor")
	fmt.Println(result)
	// prints: 47827427
}
```

### Caverphone2
Caverphone2 is a phonetic algorithm used for indexing and matching names, particularly in English and New Zealand languages. [Wiki page](https://en.wikipedia.org/wiki/Caverphone)

```go
package main

import (
	"fmt"

	"github.com/f1monkey/phonetic/caverphone2"
)

func main() {
	e := caverphone2.NewEncoder()
	result := e.Encode("orange")
	fmt.Println(result)
	// prints: ARNK111111
}
```


### Beider-Morse
It's a Go port of [the original PHP library](https://stevemorse.org/phoneticinfo.htm)
BMPM is a phonetic algorithm used for indexing and matching names in multiple languages. Contains a huge amount of different rules to transform a word to it's phonetic representation. Current implementation is relatively slow.


To reduce outcoming binary size, the three rulesets were split into different packages:
* `github.com/f1monkey/phonetic/beidermorse` - generic rules (for general usage)
* `github.com/f1monkey/phonetic/beidermorse/beidermorseash` - ashkenazi rules
* `github.com/f1monkey/phonetic/beidermorse/beidermorsesep` - sephardic rules

Each package contains `exact` and `approx` (default) rulesets. To use `exact` ruleset, you should pass a special option to encoder (see in example).

Code examples:

* `generic` ruleset with `approx` accuracy
	```go
	import (
		"fmt"
		"github.com/f1monkey/phonetic/beidermorse"
	)

	func main() {
		encoder, _ := beidermorse.NewEncoder()
		result := encoder.Encode("orange")
		fmt.Println(result)
		// prints: [orangi oragi orongi orogi orYngi Yrangi Yrongi YrYngi oranxi oronxi orani oroni oranii oronii oranzi oronzi urangi urongi]
	}
	```
* `generic` ruleset with `exact` accuracy
	```go
	import (
		"fmt"
		"github.com/f1monkey/phonetic/beidermorse"
	)

	func main() {
		encoder, _ := beidermorse.NewEncoder(beidermorse.WithAccuracy(beidermorse.Exact))
		result := encoder.Encode("orange")
		fmt.Println(result)
		// prints: [orange oranxe oranhe oranje oranZe orandZe]

	}
	```
* `generic` ruleset with `exact` accuracy and `english` language with buffer reusing (to reduce GC pressure)
	```go
	import (
		"fmt"
		"github.com/f1monkey/phonetic/beidermorse"
	)

	func main() {
		encoder, err = beidermorse.NewEncoder(
			beidermorse.WithAccuracy(beidermorse.Exact),
			beidermorse.WithLang(beidermorse.English),
			beidermorse.WithBufferReuse(true),
		)
		result := encoder.Encode("orange")
		fmt.Println(result)
		// prints: [orenk orenge orendS orendZe oronk oronge orondS orondZe orank orange orandS orandZe arenk arenge arendS arendZe aronk aronge arondS arondZe arank arange arandS arandZe]

	}
	```
* `ashkenazi` ruleset with `approx` accuracy
	```go
		import (
			"fmt"
			"github.com/f1monkey/phonetic/beidermorseash"
		)

		func main() {
			encoder, _ := beidermorseash.NewEncoder()
			result := encoder.Encode("orange")
			fmt.Println(result)
			// prints: [orangi orongi orYngi Yrangi Yrongi YrYngi oranzi oronzi orani oroni oranxi oronxi urangi urongi]
		}
	```
* `sephardic` ruleset with `approx` accuracy
	```go
		import (
			"fmt"
			"github.com/f1monkey/phonetic/beidermorsesep"
		)

		func main() {
			encoder, _ := beidermorsesep.NewEncoder()
			result := encoder.Encode("orange")
			fmt.Println(result)
			// prints: [uranzi uranz uranS uranzi uranz uranhi uranh]
		}
	```