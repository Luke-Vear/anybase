package anybase

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	for _, tc := range tt {
		decoded, err := Decode(tc.encodedValue, tc.alphabet)
		if !reflect.DeepEqual(tc.outputErr, err) {
			t.Errorf("actualErr: %v, outputErr: %v, uintValue: %v, alphabet: %v",
				err, tc.outputErr, tc.uintValue, tc.alphabet)
		}
		if decoded != tc.uintValue {
			t.Errorf("actual: %v, encodedValue: %v, uintValue: %v, alphabet: %v",
				decoded, tc.encodedValue, tc.uintValue, tc.alphabet)
		}
	}
}
