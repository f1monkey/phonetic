package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Regular expression pattern
	pattern := "[ln]h[aou]"

	// Escape any special characters
	basicPattern := regexp.QuoteMeta(pattern)

	// Print the result
	fmt.Println(basicPattern)
}
