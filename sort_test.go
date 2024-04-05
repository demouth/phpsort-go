package sortregular

import (
	"reflect"
	"testing"
)

func TestSortStrings(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
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
			[]string{
				"001",
				"002",
				"3",
				"004C",
				"004A",
				"004B",
				"005",
				"006",
				"07",
				"8",
				"00031",
				"32ABC",
				"0033",
				"40",
				"!!111",
			},
			[]string{
				"!!111",
				"001",
				"002",
				"004A",
				"004B",
				"004C",
				"3",
				"005",
				"006",
				"07",
				"8",
				"00031",
				"0033",
				"32ABC",
				"40",
			},
		},
		{
			[]string{
				"002",
				"1",
				"3",
				"0004",
				"12",
				"10",
				"200",
				"100",
				"20",
				"11",
			},
			[]string{
				"1",
				"002",
				"3",
				"0004",
				"10",
				"11",
				"12",
				"20",
				"100",
				"200",
			},
		},
		{
			[]string{
				"lemon",
				"orange",
				"banana",
				"orange",
				"apple",
				"banana",
				"lemon",
			},
			[]string{
				"apple",
				"banana",
				"banana",
				"lemon",
				"lemon",
				"orange",
				"orange",
			},
		},
	}

	for _, test := range tests {
		target := make([]string, len(test.input))
		copy(target, test.input)

		SortRegular(target)

		if !reflect.DeepEqual(target, test.expected) {
			t.Errorf("Expected: %#v, got: %#v", test.expected, target)
		}
	}
}
