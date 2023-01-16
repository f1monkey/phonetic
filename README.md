# Phonetic

Set of different phonetic encoders' implementations.

Now consists of:
* Beider-Morse encoder -  BMPM implementation. It's a Go port of [the original PHP library](https://stevemorse.org/phoneticinfo.htm)


## Install

Install:
```
$ go get -v github.com/f1monkey/phonetic
```

## Usage

### Beider-Morse encoder

```go
package main

import (
	"fmt"

	"github.com/f1monkey/phonetic/beidermorse"
)

func main() {
	e, err := beidermorse.NewEncoder()
	if err != nil {
		panic(err)
	}

	result := e.Encode("orange")

	fmt.Println(result)
    // [orangi oragi orongi orogi orYngi Yrangi Yrongi YrYngi oranxi oronxi orani oroni oranii oronii oranzi oronzi urangi urongi]
}

```
