<?php

$a = ["1"," 2"," A3"];
sort($a);
var_dump($a);

/*
PHP 8.3.2 (cli)
array(3) {
  [0]=>
  string(1) "1"
  [1]=>
  string(2) " 2"
  [2]=>
  string(3) " A3"
}
*/