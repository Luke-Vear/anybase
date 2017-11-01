package anybase

// Decode a number in the base described by the alphabet.
// TODO: reject > max
// TODO: expand test cases with errors
func Decode(encoded string, alphabet string) (uint, error) {
	err := validateAlphabet(alphabet)
	if err != nil {
		return 0, err
	}

	values := make(map[byte]uint, len(alphabet))
	for i := 0; i < len(alphabet); i++ {
		values[alphabet[i]] = uint(i)
	}

	rev, decimal := reverse(encoded), uint(0)
	encodedLen, base := uint(len(encoded)), uint(len(alphabet))
	for i := uint(0); i < encodedLen; i++ {
		decimal += values[rev[i]] * uintPow(base, i)
	}
	return decimal, nil
}
