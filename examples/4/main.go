package main

import (
	"strconv"

	"github.com/demouth/phpsort-go"
)

func main() {
	strs := makeStrings(10000)
	phpsort.Sort(strs)
}

func makeStrings(n int) []string {
	base := make([]string, n)
	for i := 0; i < n; i++ {
		base[n-1-i] = strconv.Itoa(i)
	}
	return base
}
