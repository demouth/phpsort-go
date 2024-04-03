package sortregular

import (
	"reflect"
	"testing"
)

func TestSortStrings(t *testing.T) {
	input := []string{
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
	}
	expected := []string{
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
	}

	target := make([]string, len(input))
	copy(target, input)

	SortRegular(target)

	if !reflect.DeepEqual(target, expected) {
		t.Errorf("Expected: %#v, got: %#v", expected, target)
	}
}
