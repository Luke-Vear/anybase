package anybase

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
