package main

import (
	"github.com/demouth/phpsort-go"
)

func main() {
	println(phpsort.ZendiSmartStrcmp("2", "1"))    // 1
	println(phpsort.ZendiSmartStrcmp("1", "2"))    // -1
	println(phpsort.ZendiSmartStrcmp("  10", "2")) // 1
	println(phpsort.ZendiSmartStrcmp("  1", "1"))  // 0
	println(phpsort.ZendiSmartStrcmp("A", "1"))    // 1
}
