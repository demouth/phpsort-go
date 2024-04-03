<?php

$a = [
    'lemon',
    'orange',
    'banana',
    'apple',
    '0     ',
    ' 1    ',
    '  2   ',
    '   3  ',
    '   04 ',
    '    05',
    '0000006',
    '7',
    '100',
    '201',
    '2001',
    '200X',
    'X001',
    'X002',
    'X0003',
    'X0030',
    'X0040',
];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(21) {
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
  string(7) "0000006"
  [7]=>
  string(1) "7"
  [8]=>
  string(3) "100"
  [9]=>
  string(3) "201"
  [10]=>
  string(4) "2001"
  [11]=>
  string(4) "200X"
  [12]=>
  string(5) "X0003"
  [13]=>
  string(4) "X001"
  [14]=>
  string(4) "X002"
  [15]=>
  string(5) "X0030"
  [16]=>
  string(5) "X0040"
  [17]=>
  string(5) "apple"
  [18]=>
  string(6) "banana"
  [19]=>
  string(5) "lemon"
  [20]=>
  string(6) "orange"
}
*/