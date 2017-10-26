package anybase

import (
	"testing"
)

var (
	base2Alphabet  = "01"
	base10Alphabet = base2Alphabet + "23456789"
	base16Alphabet = base10Alphabet + "abcdef"

	base56Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

	tt = []struct {
		alphabet     string
		uintValue    uint
		encodedValue string
	}{
		{
			alphabet:     base2Alphabet,
			uintValue:    0,
			encodedValue: "0",
		},
		{
			alphabet:     base2Alphabet,
			uintValue:    1,
			encodedValue: "1",
		},
		{
			alphabet:     base2Alphabet,
			uintValue:    2,
			encodedValue: "10",
		},
		{
			alphabet:     base10Alphabet,
			uintValue:    987,
			encodedValue: "987",
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    0,
			encodedValue: "2",
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    1,
			encodedValue: "3",
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    55,
			encodedValue: "z",
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    56,
			encodedValue: "32",
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    3135,
			encodedValue: "zz",
		},
		{
			alphabet:     base56Alphabet,
			uintValue:    0xFFFFFFFFFFFFFFFF,
			encodedValue: "36psTsTZwTUH",
		},
	}
)

func TestUintToString(t *testing.T) {
	for _, tc := range tt {
		gen, _ := NewGenerator(tc.alphabet)
		n := gen.NumberFromUint(tc.uintValue)
		if n.String() != tc.encodedValue {
			t.Errorf("actual: %v, encodedValue: %v, uintValue: %v, alphabet: %v", n.String(), tc.encodedValue, tc.uintValue, tc.alphabet)
		}
	}
}

func TestEncodedToUint(t *testing.T) {
	for _, tc := range tt {
		gen, _ := NewGenerator(tc.alphabet)
		n := gen.NumberFromEncoded(tc.encodedValue)
		if n.Decimal() != tc.uintValue {
			t.Errorf("actual: %v, encodedValue: %v, uintValue: %v, alphabet: %v", n.Decimal(), tc.encodedValue, tc.uintValue, tc.alphabet)
		}
	}
}
