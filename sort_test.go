package sortregular

import (
	"reflect"
	"slices"
	"sort"
	"strconv"
	"testing"
)

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			"test1.php",
			[]string{
				"lemon",
				"orange",
				"banana",
				"apple",
				"0     ",
				" 1    ",
				"  2   ",
				"   3  ",
				"   04 ",
				"    05",
				"0000006",
				"7",
				"100",
				"201",
				"2001",
				"200X",
				"X001",
				"X002",
				"X0003",
				"X0030",
				"X0040",
			},
			[]string{
				"0     ",
				" 1    ",
				"  2   ",
				"   3  ",
				"   04 ",
				"    05",
				"0000006",
				"7",
				"100",
				"201",
				"2001",
				"200X",
				"X0003",
				"X001",
				"X002",
				"X0030",
				"X0040",
				"apple",
				"banana",
				"lemon",
				"orange",
			},
		},
		{
			"test2.php",
			[]string{
				"  001  ",
				" 002 ",
				"3",
				"004C",
				"004A",
				"004B",
				"005",
				"006",
				"07",
				"8",
				"",
				"  ",
				"  9  ",
				"  0010A  ",
				"  B0010  ",
				"  B0011  ",
				"B0012",
				"  B0021  ",
				"  B0022  ",
				"B0030",
				"B0030A",
				"B0030B",
				"B00310",
				"B0031A",
				"00031",
				"32ABC",
				"0033",
				"40",
				"!!111",
			},
			[]string{
				"",
				"  ",
				"  001  ",
				"  0010A  ",
				"  9  ",
				"  B0010  ",
				"  B0011  ",
				"  B0021  ",
				"  B0022  ",
				" 002 ",
				"!!111",
				"004A",
				"004B",
				"004C",
				"3",
				"005",
				"006",
				"07",
				"32ABC",
				"8",
				"00031",
				"0033",
				"40",
				"B0012",
				"B0030",
				"B0030A",
				"B0030B",
				"B00310",
				"B0031A",
			},
		},
		{
			"test3.php",
			[]string{
				"002",
				"1",
				"3",
				"3",
				"0",
				"-1",
				"-2",
				"-10",
				"0004",
				"12",
				"10",
				"200",
				"100",
				"20",
				"11",
				"+12",
				"9223372036854775807",
			},
			[]string{
				"-10",
				"-2",
				"-1",
				"0",
				"1",
				"002",
				"3",
				"3",
				"0004",
				"10",
				"11",
				"12",
				"+12",
				"20",
				"100",
				"200",
				"9223372036854775807",
			},
		},
		{
			"test4.php",
			[]string{
				"lemon",
				"orange",
				"banana2",
				"banana",
				"banana1",
				"banana20",
				"banana21",
				"banana30",
				"banana30 1",
				"banana30 2",
				"banana10",
				"orange",
				"apple",
				"banana",
				"lemon",
			},
			[]string{
				"apple",
				"banana",
				"banana",
				"banana1",
				"banana10",
				"banana2",
				"banana20",
				"banana21",
				"banana30",
				"banana30 1",
				"banana30 2",
				"lemon",
				"lemon",
				"orange",
				"orange",
			},
		},
		{
			"test5.php",
			[]string{
				"1",
				" 2",
				" A3",
			},
			[]string{
				"1",
				" 2",
				" A3",
			},
		},
		{
			"test6.php",
			[]string{
				" A3",
				"1",
				" 2",
			},
			[]string{
				" A3",
				"1",
				" 2",
			},
		},
		{
			"test7.php",
			[]string{
				"1",
				" A3",
				" 2",
			},
			[]string{
				" 2",
				" A3",
				"1",
			},
		},
		{
			"8.php",
			[]string{
				"2",
				"1",
			},
			[]string{
				"1",
				"2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			target := make([]string, len(test.input))
			copy(target, test.input)

			SortRegular(target)

			if !reflect.DeepEqual(target, test.expected) {
				t.Errorf("Expected: %#v, got: %#v", test.expected, target)
			}
		})
	}
}

func Benchmark100(b *testing.B) {
	strs := makeStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortRegular(strs)
	}
}

func BenchmarkSlicesPackage100(b *testing.B) {
	strs := makeStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.Sort(strs)
	}
}

func BenchmarkCmp100000(b *testing.B) {
	strs := makeStrings(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.SliceStable(strs, func(i, j int) bool {
			if ZendiSmartStrcmp(strs[i], strs[j]) < 0 {
				return true
			}
			return false
		})
	}
}

func makeStrings(n int) []string {
	base := make([]string, n)
	for i := 0; i < n; i++ {
		base[i] = strconv.Itoa(i)
	}
	return base
}
