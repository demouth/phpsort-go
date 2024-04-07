package sortregular

import (
	"unsafe"
)

type compareFunc func(a, b unsafe.Pointer) int

type swapFunc func(a, b unsafe.Pointer)

// https://github.com/php/php-src/blob/0a0e8064e044b133da423952d8e78d50c4841a2e/Zend/zend_sort.c#L248
func ZendSort(base []string, cmp compareFunc, swp swapFunc) {
	nmemb := len(base)

	for {
		if nmemb <= 16 {
			zendInsertSort(base, cmp, swp)
			return
		} else {
			startIdx := 0
			endIdx := startIdx + nmemb
			pivotIdx := nmemb >> 1
			start := unsafe.Pointer(&base[startIdx])
			end := unsafe.Pointer(&base[endIdx-1])
			pivot := unsafe.Pointer(&base[pivotIdx])

			zendSort3(start, pivot, end, cmp, swp)
			swp(unsafe.Pointer(&base[0+1]), pivot)
			pivotIdx = 1
			pivot = unsafe.Pointer(&base[pivotIdx])
			i := pivotIdx + 1
			j := nmemb - 1

			for {
				for cmp(pivot, unsafe.Pointer(&base[i])) > 0 {
					i++
					if i == j {
						goto done
					}
				}
				j--
				if j == i {
					goto done
				}
				for cmp(unsafe.Pointer(&base[j]), pivot) > 0 {
					j--
					if j == i {
						goto done
					}
				}
				swp(unsafe.Pointer(&base[i]), unsafe.Pointer(&base[j]))
				i++
				if i == j {
					goto done
				}
			}
		done:
			swp(unsafe.Pointer(&base[pivotIdx]), unsafe.Pointer(&base[i-1]))
			if i-1-startIdx < endIdx-i {
				ZendSort(base[:i], cmp, swp)
				base = base[i:]
				nmemb = len(base)
			} else {
				ZendSort(base[i:], cmp, swp)
				nmemb = i - startIdx - 1
			}
		}
	}
}

func zendSort2(a, b unsafe.Pointer, cmp compareFunc, swp swapFunc) {
	if cmp(a, b) > 0 {
		swp(a, b)
	}
}

func zendSort3(a, b, c unsafe.Pointer, cmp compareFunc, swp swapFunc) {
	if cmp(a, b) > 0 {
		swp(a, b)
	}
	if cmp(b, c) > 0 {
		swp(b, c)
		if cmp(a, b) > 0 {
			swp(a, b)
		}
	}
}
func zendSort4(a, b, c, d unsafe.Pointer, cmp compareFunc, swp swapFunc) {
	zendSort3(a, b, c, cmp, swp)
	if cmp(c, d) > 0 {
		swp(c, d)
		if cmp(b, c) > 0 {
			swp(b, c)
			if cmp(a, b) > 0 {
				swp(a, b)
			}
		}
	}
}

func zendSort5(a, b, c, d, e unsafe.Pointer, cmp compareFunc, swp swapFunc) {
	zendSort4(a, b, c, d, cmp, swp)
	if cmp(d, e) > 0 {
		swp(d, e)
		if cmp(c, d) > 0 {
			swp(c, d)
			if cmp(b, c) > 0 {
				swp(b, c)
				if cmp(a, b) > 0 {
					swp(a, b)
				}
			}
		}
	}
}

func zendInsertSort(base []string, cmp compareFunc, swp swapFunc) {
	nmemb := len(base)
	switch nmemb {
	case 0, 1:
		// No need to sort
	case 2:
		zendSort2(
			unsafe.Pointer(&base[0]),
			unsafe.Pointer(&base[1]),
			cmp,
			swp,
		)
	case 3:
		zendSort3(
			unsafe.Pointer(&base[0]),
			unsafe.Pointer(&base[1]),
			unsafe.Pointer(&base[2]),
			cmp,
			swp,
		)
	case 4:
		zendSort4(
			unsafe.Pointer(&base[0]),
			unsafe.Pointer(&base[1]),
			unsafe.Pointer(&base[2]),
			unsafe.Pointer(&base[3]),
			cmp,
			swp,
		)
	case 5:
		zendSort5(
			unsafe.Pointer(&base[0]),
			unsafe.Pointer(&base[1]),
			unsafe.Pointer(&base[2]),
			unsafe.Pointer(&base[3]),
			unsafe.Pointer(&base[4]),
			cmp,
			swp,
		)

	default:
		start := 0
		end := nmemb
		sentry := 6
		siz2 := 2

		for i := 1; i < sentry; i += 1 {
			j := i - 1
			if cmp(unsafe.Pointer(&base[j]), unsafe.Pointer(&base[i])) > 0 {
				for j != start {
					j -= 1
					if cmp(unsafe.Pointer(&base[j]), unsafe.Pointer(&base[i])) <= 0 {
						j += 1
						break
					}
				}
				for k := i; k > j; k -= 1 {
					swp(unsafe.Pointer(&base[k]), unsafe.Pointer(&base[k-1]))
				}
			}
		}

		for i := sentry; i < end; i += 1 {
			j := i - 1
			if cmp(unsafe.Pointer(&base[j]), unsafe.Pointer(&base[i])) > 0 {
				for {
					j -= siz2
					if cmp(unsafe.Pointer(&base[j]), unsafe.Pointer(&base[i])) <= 0 {
						j += 1
						if cmp(unsafe.Pointer(&base[j]), unsafe.Pointer(&base[i])) <= 0 {
							j += 1
						}
						break
					}
					if j == start {
						break
					}
					if j == start+1 {
						j -= 1
						if cmp(unsafe.Pointer(&base[i]), unsafe.Pointer(&base[j])) > 0 {
							j += 1
						}
						break
					}
				}
				for k := i; k > j; k -= 1 {
					swp(unsafe.Pointer(&base[k]), unsafe.Pointer(&base[k-1]))
				}
			}
		}
	}
}
