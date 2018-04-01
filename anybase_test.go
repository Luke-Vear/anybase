package anybase

import "testing"

func TestUintPow(t *testing.T) {
	tt := []struct {
		description   string
		inputBase     uint
		inputExponent uint
		expected      uint
	}{
		{
			description:   "simple 2 pow 2",
			inputBase:     2,
			inputExponent: 2,
			expected:      4,
		},
		{
			description:   "exponent 0 - should be 1",
			inputBase:     100,
			inputExponent: 0,
			expected:      1,
		},
		{
			description:   "exponent 1 - should be same as base",
			inputBase:     100,
			inputExponent: 1,
			expected:      100,
		},
	}

	for _, tc := range tt {
		if actual := uintPow(tc.inputBase, tc.inputExponent); actual != tc.expected {
			t.Errorf("actual: %v, expected: %v, inputs: %v %v", actual, tc.expected, tc.inputBase, tc.inputExponent)
		}
	}
}
