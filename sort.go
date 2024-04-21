package phpsort

type compareFunc func(a, b string) int
type compareStableFunc func(a, b string, ai, bi int) int

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

	m := make(map[*string]int, len(strings))
	for i, s := range strings {
		m[&s] = i
	}

	ext := make([]int, len(strings))
	for i := range strings {
		ext[i] = i
	}

	swp := func(i, j int) {
		strings[i], strings[j] = strings[j], strings[i]
		ext[i], ext[j] = ext[j], ext[i]
	}

	// https://github.com/php/php-src/blob/98b43d07f9d0bea021c8fd6bda70bfdbbb7a6b7f/ext/standard/array.c#L105
	cmpStable := func(a, b string, ai, bi int) int {
		r := ZendiSmartStrcmp(a, b)
		if r == 0 {
			if ext[ai] > ext[bi] {
				return 1
			} else if ext[ai] < ext[bi] {
				return -1
			}
			return 0
		}
		return r
	}

	zendSort(strings, 0, len(strings)-1, cmpStable, swp)
}

// https://github.com/php/php-src/blob/0a0e8064e044b133da423952d8e78d50c4841a2e/Zend/zend_sort.c#L248
func zendSort(base []string, start, end int, cmp compareStableFunc, swp swapFunc) {
	for {
		nmemb := end - start + 1
		if nmemb <= 16 {
			zendInsertSort(base, start, end, cmp, swp)
			return
		} else {
			startIdx := start
			endIdx := startIdx + nmemb
			pivotIdx := start + (nmemb >> 1)

			if nmemb>>10 != 0 {
				offset := nmemb >> 1
				delta := offset >> 1
				zendSort5(base, startIdx, startIdx+delta, pivotIdx, pivotIdx+delta, endIdx-1, cmp, swp)
			} else {
				zendSort3(base, startIdx, pivotIdx, endIdx-1, cmp, swp)
			}
			swp(start+1, pivotIdx)
			pivotIdx = start + 1
			pivot := base[pivotIdx]
			i := pivotIdx + 1
			j := endIdx - 1

			for {
				for cmp(pivot, base[i], pivotIdx, i) >= 0 {
					i++
					if i == j {
						goto done
					}
				}
				j--
				if j == i {
					goto done
				}
				for cmp(base[j], pivot, j, pivotIdx) > 0 {
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

func zendSort2(base []string, a, b int, cmp compareStableFunc, swp swapFunc) {
	if cmp(base[a], base[b], a, b) > 0 {
		swp(a, b)
	}
}

func zendSort3(base []string, a, b, c int, cmp compareStableFunc, swp swapFunc) {
	if !(cmp(base[a], base[b], a, b) > 0) {
		if !(cmp(base[b], base[c], b, c) > 0) {
			return
		}
		swp(b, c)
		if cmp(base[a], base[b], a, b) > 0 {
			swp(a, b)
		}
		return
	}
	if !(cmp(base[c], base[b], c, b) > 0) {
		swp(a, c)
		return
	}
	swp(a, b)
	if cmp(base[b], base[c], b, c) > 0 {
		swp(b, c)
	}
}
func zendSort4(base []string, a, b, c, d int, cmp compareStableFunc, swp swapFunc) {
	zendSort3(base, a, b, c, cmp, swp)
	if cmp(base[c], base[d], c, d) > 0 {
		swp(c, d)
		if cmp(base[b], base[c], b, c) > 0 {
			swp(b, c)
			if cmp(base[a], base[b], a, b) > 0 {
				swp(a, b)
			}
		}
	}
}

func zendSort5(base []string, a, b, c, d, e int, cmp compareStableFunc, swp swapFunc) {
	zendSort4(base, a, b, c, d, cmp, swp)
	if cmp(base[d], base[e], d, e) > 0 {
		swp(d, e)
		if cmp(base[c], base[d], c, d) > 0 {
			swp(c, d)
			if cmp(base[b], base[c], b, c) > 0 {
				swp(b, c)
				if cmp(base[a], base[b], a, b) > 0 {
					swp(a, b)
				}
			}
		}
	}
}

func zendInsertSort(base []string, start, end int, cmp compareStableFunc, swp swapFunc) {
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
			if !(cmp(base[j], base[i], j, i) > 0) {
				continue
			}
			for j != start {
				j -= 1
				if !(cmp(base[j], base[i], j, i) > 0) {
					j += 1
					break
				}
			}
			for k := i; k > j; k -= 1 {
				swp(k, k-1)
			}
		}

		for i := sentry; i < end+1; i += 1 {
			j := i - 1
			if !(cmp(base[j], base[i], j, i) > 0) {
				continue
			}
			for {
				j -= siz2
				if !(cmp(base[j], base[i], j, i) > 0) {
					j += 1
					if !(cmp(base[j], base[i], j, i) > 0) {
						j += 1
					}
					break
				}
				if j == start {
					break
				}
				if j == start+1 {
					j -= 1
					if cmp(base[i], base[j], i, j) > 0 {
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
