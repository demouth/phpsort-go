package phpsort

import (
	"testing"
)

func TestIsNumericStringEx(t *testing.T) {
	type (
		expected struct {
			r    uint8
			lval int64
			dval float64
		}
		got expected
	)

	tests := []struct {
		input string
		expected
	}{
		{
			"",
			expected{0, 0, 0},
		},
		{
			"A",
			expected{0, 0, 0},
		},
		{
			"1A",
			expected{0, 0, 0},
		},
		{
			"002A",
			expected{0, 0, 0},
		},

		// IS_LONG
		{
			"1",
			expected{IS_LONG, 1, 0},
		},
		{
			"00",
			expected{IS_LONG, 0, 0},
		},
		{
			"002",
			expected{IS_LONG, 2, 0},
		},
		{
			"-10",
			expected{IS_LONG, -10, 0},
		},
		{
			"  +10  ",
			expected{IS_LONG, 10, 0},
		},
		{
			"  100  ",
			expected{IS_LONG, 100, 0},
		},

		// IS_DOUBLE
		{
			"1.0",
			expected{IS_DOUBLE, 0, 1.0},
		},
		{
			"-10.5",
			expected{IS_DOUBLE, 0, -10.5},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {

			r, _, lval, dval, _ := isNumericStringEx(test.input)
			if r != test.expected.r || lval != test.expected.lval || dval != test.expected.dval {
				t.Errorf("Expected: %#v, got: %#v", test.expected, got{r, lval, dval})
			}

		})
	}
}
