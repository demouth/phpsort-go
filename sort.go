package sortregular

import (
	"sort"
	"strconv"
	"strings"
)

func SortRegular(strings []string) {
	sort.SliceStable(strings, func(i, j int) bool {
		return cmp(strings[i], strings[j])
	})
}

func cmp(s1, s2 string) bool {
	i1, a1 := is_numeric_string_ex(s1)
	i2, a2 := is_numeric_string_ex(s2)
	if i1 > 0 && i2 > 0 {
		return i1 < i2
	}
	return a1 < a2
}

// https://github.com/php/php-src/blob/7a3516cca5ad307ca7dcb63224448661f30d623e/Zend/zend_operators.h#L145
func ZEND_IS_DIGIT(s string) bool {
	for _, c := range s {
		if '0' <= c && c <= '9' {
			continue
		}
		return false
	}
	return true
}

// is_numeric_string_ex
// https://github.com/php/php-src/blob/7a3516cca5ad307ca7dcb63224448661f30d623e/Zend/zend_operators.c#L3507
func is_numeric_string_ex(s string) (int, string) {
	s = strings.TrimSpace(s)

	if ZEND_IS_DIGIT(s) {
		l := len(s)
		i := 0
	L:
		for i < l {
			switch s[i] {
			case '0':
				i++
			default:
				break L
			}
		}
		s = s[i:]
		conv, _ := strconv.Atoi(s)
		return conv, ""
	} else {
		return 0, s
	}
}
