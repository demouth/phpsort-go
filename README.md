# SORT_REGULAR in Go

This repository contains a Go implementation of the `SORT_REGULAR` from PHP.

## Usage

Import the package and use the sort_regular function:

```go
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

	sortregular.SortRegular(strings)

	for _, s := range strings {
		println(s)
	}
}
```

When using the comparison function:

```go
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
```

## PHP SORT_REGULAR

For reference, here's the original `SORT_REGULAR` in PHP:

```php
<?php

$a = [
    'lemon',
    'apple',
    '0     ',
    ' 1    ',
    '  2   ',
    '   3  ',
    '   04 ',
    '    05',
];
sort($a);
var_dump($a);
/*
array(8) {
  [0]=>
  string(6) "0     "
  [1]=>
  string(6) " 1    "
  [2]=>
  string(6) "  2   "
  [3]=>
  string(6) "   3  "
  [4]=>
  string(6) "   04 "
  [5]=>
  string(6) "    05"
  [6]=>
  string(5) "apple"
  [7]=>
  string(5) "lemon"
}
*/
```

