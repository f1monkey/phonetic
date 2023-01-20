# Phonetic

Set of different phonetic encoders' implementations.

Now consists of:
* [Beider-Morse encoder](#beider-morse-encoder) -  BMPM implementation. It's a Go port of [the original PHP library](https://stevemorse.org/phoneticinfo.htm)


## Installion

To install:
```
$ go get -v github.com/f1monkey/phonetic
```

## Usage

### Beider-Morse encoder
Contains a HUGE amount of different rules to transform a word to it's phonetic rperesentation.
To reduce outcoming binary size, the three rulesets were split into different packages:
* `github.com/f1monkey/phonetic/beidermorse` - generic rules (for general usage)
* `github.com/f1monkey/phonetic/beidermorse/beidermorseash` - ashkenazi rules
* `github.com/f1monkey/phonetic/beidermorse/beidermorsesep` - sephardic rules

Each package contains `exact` and `approx` (default) rulesets. To use `exact` ruleset, you should pass a special option to encoder (see in example).

Code example:

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
