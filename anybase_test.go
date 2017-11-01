package anybase

import "testing"

var (
	base2Alphabet  = "01"
	base10Alphabet = base2Alphabet + "23456789"
	base16Alphabet = base10Alphabet + "abcdef"

	base56Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

	tt = []struct {
		alphabet     string
		uintValue    uint
		encodedValue string
		outputErr    error
	}{
		{
			alphabet:     base2Alphabet,
			uintValue:    0,
			encodedValue: "0",
			outputErr:    nil,
		},
		{
			alphabet:     base2Alphabet,
			uintValue:    1,
			encodedValue: "1",
			outputErr:    nil,
		},
		{
			alphabet:     base2Alphabet,
			uintValue:    2,
			encodedValue: "10",
			outputErr:    nil,
		},
		{
			alphabet:     base10Alphabet,
			uintValue:    987,
			encodedValue: "987",
			outputErr:    nil,
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    0,
			encodedValue: "2",
			outputErr:    nil,
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    1,
			encodedValue: "3",
			outputErr:    nil,
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    55,
			encodedValue: "z",
			outputErr:    nil,
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    56,
			encodedValue: "32",
			outputErr:    nil,
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    3135,
			encodedValue: "zz",
			outputErr:    nil,
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    0xFFFFFFFFFFFFFFFF,
			encodedValue: "36psTsTZwTUH",
			outputErr:    nil,
		},
	}
)

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
