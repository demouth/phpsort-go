# phpsort-go

`phpsort-go` is a Go language port of the PHP sort function.
It provides functionality similar to the `sort` function used in PHP.

## Usage

Import the package and use the Sort function:

```go
package main

import (
	"github.com/demouth/phpsort-go/v2"
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

	phpsort.Sort(strings)

	for _, s := range strings {
		println(s)
	}
}
```

When using the comparison function:

```go
package main

import (
	"github.com/demouth/phpsort-go/v2"
)

func main() {
	println(phpsort.ZendiSmartStrcmp("2", "1"))    // 1
	println(phpsort.ZendiSmartStrcmp("1", "2"))    // -1
	println(phpsort.ZendiSmartStrcmp("  10", "2")) // 1
	println(phpsort.ZendiSmartStrcmp("  1", "1"))  // 0
	println(phpsort.ZendiSmartStrcmp("A", "1"))    // 1
}
```

## Behavior of the sort function in PHP

For reference, here's the original `sort` function in PHP:

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

