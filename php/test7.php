<?php

$a = ["1"," A3"," 2"];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(3) {
  [0]=>
  string(2) " 2"
  [1]=>
  string(3) " A3"
  [2]=>
  string(1) "1"
}
*/