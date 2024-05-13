package main

import (
	"fmt"

	"github.com/demouth/phpsort-go/v2"
)

func main() {
	var strings []string

	// PHP8
	strings = makeStrings()
	phpsort.Sort(strings)
	fmt.Printf("%#v\n", strings) // []string{"1.0", " 1", "1 ", "+1.0"}

	// PHP7
	strings = makeStrings()
	phpsort.Sort(strings, phpsort.WithPHP7Mode())
	fmt.Printf("%#v\n", strings) // []string{"1.0", " 1", "+1.0", "1 "}
}

func makeStrings() []string {
	return []string{
		"1.0",
		" 1",
		"1 ",
		"+1.0",
	}
}
