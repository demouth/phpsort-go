package sortregular

import (
	"math"
	"sort"
)

const (
	IS_LONG   = 1
	IS_DOUBLE = 2

	longMinDigits = "-9223372036854775808"
)

func SortRegular(strings []string) {
	sort.SliceStable(strings, func(i, j int) bool {
		r := zendiSmartStrcmp(strings[i], strings[j])
		return r < 0
	})
}

func zendBinaryStrcmp(s1, s2 string) int {
	return strcmp([]byte(s1), []byte(s2))
}

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

func isNumericStringEx(str string) (uint8, int, int64, float64, bool) {
	// Check if the string is a numeric string
	length := len(str)
	if length == 0 {
		return 0, 0, 0, 0, false
	}

	ptr := 0
	digits := 0
	dpOrE := 0
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
		for ptr < length && (isDigit(str[ptr]) || (str[ptr] == '.' && dpOrE < 1) || ((str[ptr] == 'e' || str[ptr] == 'E') && dpOrE < 2)) {
			if isDigit(str[ptr]) {
				tmpLval = tmpLval*10 + int64(str[ptr]-'0')
				digits++
				// } else if str[ptr] == '.' && dpOrE < 1 {
				// 	dpOrE = 1
				// } else if (str[ptr] == 'e' || str[ptr] == 'E') && dpOrE < 2 {
				// 	dpOrE = 2
			} else {
				if neg {
					tmpLval = -tmpLval
				}

				return IS_LONG, 0, tmpLval, localDval, false
			}
			ptr++
		}

		if digits >= 19 {
			return IS_DOUBLE, 0, 0, 0, true
		}
	} else if ptr < length && str[ptr] == '.' && ptr+1 < length && isDigit(str[ptr+1]) {
		dpOrE = 1
		ptr++
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

	if digits == 19 {
		if cmp := strcmp([]byte(str[:digits]), []byte(longMinDigits)); cmp >= 0 {
			return IS_DOUBLE, 0, 0, 0, true
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

func zendiSmartStrcmp(s1, s2 string) int {
	ret1, _, lval1, dval1, _ := isNumericStringEx(s1)
	ret2, _, lval2, dval2, _ := isNumericStringEx(s2)

	if ret1 != 0 && ret2 != 0 {
		if dval1-dval2 == 0 && ((!math.IsInf(dval1, 0) && (lval1 > math.MaxInt64 || lval1 < math.MinInt64)) || (!math.IsInf(dval1, 0) && (lval2 > math.MaxInt64 || lval2 < math.MinInt64))) {
			return zendNormalizeBool(dval1)
		}

		if ret1 == IS_DOUBLE || ret2 == IS_DOUBLE {
			if ret1 != IS_DOUBLE {
				if lval2 > math.MaxInt64 {
					return -1
				}
				dval1 = float64(lval1)
			} else if ret2 != IS_DOUBLE {
				if lval1 > math.MaxInt64 {
					return 1
				}
				dval2 = float64(lval2)
			} else if dval1 == dval2 && !math.IsInf(dval1, 0) {
				return zendBinaryStrcmp(s1, s2)
			}
			return zendNormalizeBool(dval1 - dval2)
		} else {
			if lval1 > lval2 {
				return 1
			} else if lval1 < lval2 {
				return -1
			}
			return 0
		}
	} else {
		return zendBinaryStrcmp(s1, s2)
	}
}
