package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/demouth/phpsort-go/v2"
)

func main() {
	strs := makeStrings(1000)
	sort.SliceStable(strs, func(i, j int) bool {
		if phpsort.ZendiSmartStrcmp(strs[i], strs[j]) < 0 {
			return true
		}
		return false
	})
	fmt.Printf("%#v\n", strs)
}

func makeStrings(n int) []string {
	base := make([]string, n)
	for i := 0; i < n; i++ {
		base[i] = strconv.Itoa(i)
	}
	return base
}
