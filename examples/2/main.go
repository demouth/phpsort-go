package main

import (
	"github.com/demouth/sortregular-go"
)

func main() {
	println(sortregular.ZendiSmartStrcmp("2", "1"))    // 1
	println(sortregular.ZendiSmartStrcmp("1", "2"))    // -1
	println(sortregular.ZendiSmartStrcmp("  10", "2")) // 1
	println(sortregular.ZendiSmartStrcmp("  1", "1"))  // 0
	println(sortregular.ZendiSmartStrcmp("A", "1"))    // 1
}
