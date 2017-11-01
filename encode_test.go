package anybase

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tc := range tt {
		encoded, err := Encode(tc.uintValue, tc.alphabet)
		if !reflect.DeepEqual(tc.outputErr, err) {
			t.Errorf("actualErr: %v, outputErr: %v, uintValue: %v, alphabet: %v",
				err, tc.outputErr, tc.uintValue, tc.alphabet)
		}
		if encoded != tc.encodedValue {
			t.Errorf("actual: %v, encodedValue: %v, uintValue: %v, alphabet: %v",
				encoded, tc.encodedValue, tc.uintValue, tc.alphabet)
		}
	}
}
