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
array(29) {
  [0]=>
  string(0) ""
  [1]=>
  string(2) "  "
  [2]=>
  string(7) "  001  "
  [3]=>
  string(9) "  0010A  "
  [4]=>
  string(5) "  9  "
  [5]=>
  string(9) "  B0010  "
  [6]=>
  string(9) "  B0011  "
  [7]=>
  string(9) "  B0021  "
  [8]=>
  string(9) "  B0022  "
  [9]=>
  string(5) " 002 "
  [10]=>
  string(5) "!!111"
  [11]=>
  string(4) "004A"
  [12]=>
  string(4) "004B"
  [13]=>
  string(4) "004C"
  [14]=>
  string(1) "3"
  [15]=>
  string(3) "005"
  [16]=>
  string(3) "006"
  [17]=>
  string(2) "07"
  [18]=>
  string(5) "32ABC"
  [19]=>
  string(1) "8"
  [20]=>
  string(5) "00031"
  [21]=>
  string(4) "0033"
  [22]=>
  string(2) "40"
  [23]=>
  string(5) "B0012"
  [24]=>
  string(5) "B0030"
  [25]=>
  string(6) "B0030A"
  [26]=>
  string(6) "B0030B"
  [27]=>
  string(6) "B00310"
  [28]=>
  string(6) "B0031A"
}
*/