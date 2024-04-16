package sortregular

import (
	"math"
)

const (
	IS_LONG   = 1
	IS_DOUBLE = 2 // not supported

	longMinDigits = "-9223372036854775808"
)

// if s1 is larger, return 1
// if s2 is larger, return -1
// https://github.com/php/php-src/blob/98b43d07f9d0bea021c8fd6bda70bfdbbb7a6b7f/Zend/zend_operators.c#L3323
func ZendiSmartStrcmp(s1, s2 string) int {
	ret1, _, lval1, dval1, _ := isNumericStringEx(s1)
	ret2, _, lval2, dval2, _ := isNumericStringEx(s2)

	if ret1 != 0 && ret2 != 0 {
		if dval1-dval2 == 0 && ((!math.IsInf(dval1, 0) && (lval1 > math.MaxInt64 || lval1 < math.MinInt64)) || (!math.IsInf(dval1, 0) && (lval2 > math.MaxInt64 || lval2 < math.MinInt64))) {
			return zendNormalizeBool(dval1)
		}

		if lval1 > lval2 {
			return 1
		} else if lval1 < lval2 {
			return -1
		}
		return 0
	} else {
		return zendBinaryStrcmp(s1, s2)
	}
}

func zendBinaryStrcmp(s1, s2 string) int {
	return strcmp([]byte(s1), []byte(s2))
}

// If s1 is larger, return 1
// If s2 is larger, return -1
func strcmp(s1, s2 []byte) int {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return -1
		} else if s1[i] > s2[i] {
			return 1
		}
	}
	if len(s1) < len(s2) {
		return -1
	} else if len(s1) > len(s2) {
		return 1
	}
	return 0
}

// https://github.com/php/php-src/blob/7a3516cca5ad307ca7dcb63224448661f30d623e/Zend/zend_operators.c#L3507
func isNumericStringEx(str string) (uint8, int, int64, float64, bool) {
	// Check if the string is a numeric string
	length := len(str)
	if length == 0 {
		return 0, 0, 0, 0, false
	}

	ptr := 0
	// digits := 0
	// dpOrE := 0
	localDval := 0.0
	var tmpLval int64
	neg := false

	// Skip any whitespace
	for ptr < length && (str[ptr] == ' ' || str[ptr] == '\t' || str[ptr] == '\n' || str[ptr] == '\r' || str[ptr] == '\v' || str[ptr] == '\f') {
		ptr++
	}

	if ptr < length && str[ptr] == '-' {
		neg = true
		ptr++
	} else if ptr < length && str[ptr] == '+' {
		ptr++
	}

	if ptr < length && isDigit(str[ptr]) {
		for ptr < length && (isDigit(str[ptr])) {
			tmpLval = tmpLval*10 + int64(str[ptr]-'0')
			ptr++
		}
	} else {
		return 0, 0, 0, 0, false
	}

	if ptr < length {
		for ptr < length && (str[ptr] == ' ' || str[ptr] == '\t' || str[ptr] == '\n' || str[ptr] == '\r' || str[ptr] == '\v' || str[ptr] == '\f') {
			ptr++
		}
		if ptr < length {
			return 0, 0, 0, 0, true
		}
	}

	if neg {
		tmpLval = -tmpLval
	}

	return IS_LONG, 0, tmpLval, localDval, false
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func zendNormalizeBool(value float64) int {
	if value < 0 {
		return -1
	} else if value > 0 {
		return 1
	}
	return 0
}
