<?php

$a = ['5','4','3','2','1'];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(5) {
  [0]=>
  string(1) "1"
  [1]=>
  string(1) "2"
  [2]=>
  string(1) "3"
  [3]=>
  string(1) "4"
  [4]=>
  string(1) "5"
}
*/