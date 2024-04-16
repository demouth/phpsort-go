package main

import (
	"github.com/demouth/sortregular-go"
)

func main() {
	strings := []string{
		"lemon",
		"apple",
		"0     ",
		" 1    ",
		"  2   ",
		"   3  ",
		"   04 ",
		"    05",
	}

	sortregular.Sort(strings)

	for _, s := range strings {
		println(s)
	}
}
