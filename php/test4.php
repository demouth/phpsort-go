<?php

$a = [
  'lemon',
  'orange',
  'banana2',
  'banana',
  'banana1',
  'banana20',
  'banana21',
  'banana30',
  'banana30 1',
  'banana30 2',
  'banana10',
  'orange',
  'apple',
  'banana',
  'lemon',
];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(15) {
  [0]=>
  string(5) "apple"
  [1]=>
  string(6) "banana"
  [2]=>
  string(6) "banana"
  [3]=>
  string(7) "banana1"
  [4]=>
  string(8) "banana10"
  [5]=>
  string(7) "banana2"
  [6]=>
  string(8) "banana20"
  [7]=>
  string(8) "banana21"
  [8]=>
  string(8) "banana30"
  [9]=>
  string(10) "banana30 1"
  [10]=>
  string(10) "banana30 2"
  [11]=>
  string(5) "lemon"
  [12]=>
  string(5) "lemon"
  [13]=>
  string(6) "orange"
  [14]=>
  string(6) "orange"
}
*/