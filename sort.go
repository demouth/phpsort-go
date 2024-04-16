package phpsort

import "reflect"

type compareFunc func(a, b string) int

type swapFunc func(i int, j int)

type options struct {
	cmp compareFunc
}

type option func(*options)

func WithSortRegular() option {
	return func(o *options) {
		o.cmp = ZendiSmartStrcmp
	}
}

func Sort(strings []string, opts ...option) {
	o := &options{
		cmp: ZendiSmartStrcmp,
	}
	for _, opt := range opts {
		opt(o)
	}
	swp := reflect.Swapper(strings)
	zendSort(strings, 0, len(strings)-1, o.cmp, swp)
}

// https://github.com/php/php-src/blob/0a0e8064e044b133da423952d8e78d50c4841a2e/Zend/zend_sort.c#L248
func zendSort(base []string, start, end int, cmp compareFunc, swp swapFunc) {
	for {
		nmemb := end - start + 1
		if nmemb <= 16 {
			zendInsertSort(base, start, end, cmp, swp)
			return
		} else {
			startIdx := start
			endIdx := startIdx + nmemb
			pivotIdx := start + (nmemb >> 1)

			zendSort3(base, startIdx, pivotIdx, endIdx-1, cmp, swp)
			swp(start+1, pivotIdx)
			pivotIdx = start + 1
			pivot := base[pivotIdx]
			i := pivotIdx + 1
			j := startIdx + nmemb - 1

			for {
				for cmp(pivot, base[i]) > 0 {
					i++
					if i == j {
						goto done
					}
				}
				j--
				if j == i {
					goto done
				}
				for cmp(base[j], pivot) > 0 {
					j--
					if j == i {
						goto done
					}
				}
				swp(i, j)
				i++
				if i == j {
					goto done
				}
			}
		done:
			swp(pivotIdx, i-1)
			if i-1-startIdx < endIdx-i {
				zendSort(base, start, i-1, cmp, swp)
				start = i
			} else {
				zendSort(base, i, end, cmp, swp)
				end = i - 1
			}
		}
	}
}

func zendSort2(base []string, a, b int, cmp compareFunc, swp swapFunc) {
	if cmp(base[a], base[b]) > 0 {
		swp(a, b)
	}
}

func zendSort3(base []string, a, b, c int, cmp compareFunc, swp swapFunc) {
	if !(cmp(base[a], base[b]) > 0) {
		if !(cmp(base[b], base[c]) > 0) {
			return
		}
		swp(b, c)
		if cmp(base[a], base[b]) > 0 {
			swp(a, b)
		}
		return
	}
	if !(cmp(base[c], base[b]) > 0) {
		swp(a, c)
		return
	}
	swp(a, b)
	if cmp(base[b], base[c]) > 0 {
		swp(b, c)
	}
}
func zendSort4(base []string, a, b, c, d int, cmp compareFunc, swp swapFunc) {
	zendSort3(base, a, b, c, cmp, swp)
	if cmp(base[c], base[d]) > 0 {
		swp(c, d)
		if cmp(base[b], base[c]) > 0 {
			swp(b, c)
			if cmp(base[a], base[b]) > 0 {
				swp(a, b)
			}
		}
	}
}

func zendSort5(base []string, a, b, c, d, e int, cmp compareFunc, swp swapFunc) {
	zendSort4(base, a, b, c, d, cmp, swp)
	if cmp(base[d], base[e]) > 0 {
		swp(d, e)
		if cmp(base[c], base[d]) > 0 {
			swp(c, d)
			if cmp(base[b], base[c]) > 0 {
				swp(b, c)
				if cmp(base[a], base[b]) > 0 {
					swp(a, b)
				}
			}
		}
	}
}

func zendInsertSort(base []string, start, end int, cmp compareFunc, swp swapFunc) {
	nmemb := end - start + 1
	switch nmemb {
	case 0, 1:
		// No need to sort
	case 2:
		zendSort2(
			base,
			0+start,
			1+start,
			cmp,
			swp,
		)
	case 3:
		zendSort3(
			base,
			0+start,
			1+start,
			2+start,
			cmp,
			swp,
		)
	case 4:
		zendSort4(
			base,
			0+start,
			1+start,
			2+start,
			3+start,
			cmp,
			swp,
		)
	case 5:
		zendSort5(
			base,
			0+start,
			1+start,
			2+start,
			3+start,
			4+start,
			cmp,
			swp,
		)

	default:
		sentry := 6 + start
		siz2 := 2

		for i := 1 + start; i < sentry; i += 1 {
			j := i - 1
			if cmp(base[j], base[i]) > 0 {
				for j != start {
					j -= 1
					if cmp(base[j], base[i]) <= 0 {
						j += 1
						break
					}
				}
				for k := i; k > j; k -= 1 {
					swp(k, k-1)
				}
			}
		}

		for i := sentry; i < end+1; i += 1 {
			j := i - 1
			if cmp(base[j], base[i]) > 0 {
				for {
					j -= siz2
					if cmp(base[j], base[i]) <= 0 {
						j += 1
						if cmp(base[j], base[i]) <= 0 {
							j += 1
						}
						break
					}
					if j == start {
						break
					}
					if j == start+1 {
						j -= 1
						if cmp(base[i], base[j]) > 0 {
							j += 1
						}
						break
					}
				}
				for k := i; k > j; k -= 1 {
					swp(k, k-1)
				}
			}
		}
	}
}
