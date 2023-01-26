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

```go
package main

import (
	"fmt"

	"github.com/f1monkey/phonetic/beidermorse"
	"github.com/f1monkey/phonetic/beidermorse/beidermorseash"
	"github.com/f1monkey/phonetic/beidermorse/beidermorsesep"
	"github.com/f1monkey/phonetic/beidermorse/common"
)

func main() {

	// USE ENCODER WITH "GENERIC" RULESET WITH "APPROX" ACCURACY
	genEncoder, err := beidermorse.NewEncoder()
	if err != nil {
		panic(err)
	}
	result := genEncoder.Encode("orange")
	fmt.Println(result)
	// prints: [orangi oragi orongi orogi orYngi Yrangi Yrongi YrYngi oranxi oronxi orani oroni oranii oronii oranzi oronzi urangi urongi]

	// USE ENCODER WITH "GENERIC" RULESET WITH "EXACT" ACCURACY
	genEncoder, err = beidermorse.NewEncoder(beidermorse.WithAccuracy(common.Exact))
	if err != nil {
		panic(err)
	}
	result = genEncoder.Encode("orange")
	fmt.Println(result)
	// prints: [orange oranxe oranhe oranje oranZe orandZe]

	// USE ENCODER WITH "GENERIC" RULESET WITH "EXACT" ACCURACY
	// AND "ENGLISH" LANGUAGE
	genEncoder, err = beidermorse.NewEncoder(
		beidermorse.WithAccuracy(common.Exact),
		beidermorse.WithLang(beidermorse.English),
	)
	if err != nil {
		panic(err)
	}
	result = genEncoder.Encode("orange")
	fmt.Println(result)
	// prints: [orenk orenge orendS orendZe oronk oronge orondS orondZe orank orange orandS orandZe arenk arenge arendS arendZe aronk aronge arondS arondZe arank arange arandS arandZe]


	// USE ENCODER WITH "ASHKENAZI" RULESET
	ashEncoder, err := beidermorseash.NewEncoder()
	if err != nil {
		panic(err)
	}
	result = ashEncoder.Encode("orange")
	fmt.Println(result)
	// prints: [orangi orongi orYngi Yrangi Yrongi YrYngi oranzi oronzi orani oroni oranxi oronxi urangi urongi]

	// USE ENCODER WITH "SEPHARDIC" RULESET
	sepEncoder, err := beidermorsesep.NewEncoder()
	if err != nil {
		panic(err)
	}
	result = sepEncoder.Encode("orange")
	fmt.Println(result)
	// prints: [uranzi uranz uranS uranzi uranz uranhi uranh]
}
```