
<?php

$a = [
    '1.0',
    ' 1',
    '1 ',
    '+1.0',
];
sort($a);
var_dump($a);

/*
php:8.3.2-cli
array(4) {
  [0]=>
  string(3) "1.0"
  [1]=>
  string(2) " 1"
  [2]=>
  string(2) "1 "
  [3]=>
  string(4) "+1.0"
}


php:7.4.33-cli
array(4) {
  [0]=>
  string(3) "1.0"
  [1]=>
  string(2) " 1"
  [2]=>
  string(4) "+1.0"
  [3]=>
  string(2) "1 "
}

*/