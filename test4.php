<?php

$a = [
  'lemon',
  'orange',
  'banana',
  'orange',
  'apple',
  'banana',
  'lemon',
];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(7) {
  [0]=>
  string(5) "apple"
  [1]=>
  string(6) "banana"
  [2]=>
  string(6) "banana"
  [3]=>
  string(5) "lemon"
  [4]=>
  string(5) "lemon"
  [5]=>
  string(6) "orange"
  [6]=>
  string(6) "orange"
}
*/