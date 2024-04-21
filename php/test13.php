<?php

$a = [
  '-4',
  '-2',
  '2',
  '4',
  '  -2  ',
  '   4  ',
  '  -4  ',
];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(7) {
  [0]=>
  string(2) "-4"
  [1]=>
  string(6) "  -4  "
  [2]=>
  string(2) "-2"
  [3]=>
  string(6) "  -2  "
  [4]=>
  string(1) "2"
  [5]=>
  string(1) "4"
  [6]=>
  string(6) "   4  "
}
*/