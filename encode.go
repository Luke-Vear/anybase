package anybase

// Encode a number in the base described by the alphabet.
func Encode(decimal uint, alphabet string) (string, error) {
	err := validateAlphabet(alphabet)
	if err != nil {
		return "", err
	}

	if decimal == 0 {
		return string(alphabet[0]), nil
	}

	value, base := decimal, uint(len(alphabet))
	var encoded []byte
	for value > 0 {
		rem := value % base
		value = (value - rem) / base
		encoded = append(encoded, alphabet[rem])
	}
	return reverse(string(encoded)), nil
}
