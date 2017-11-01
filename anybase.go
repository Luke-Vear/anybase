package anybase

func validateAlphabet(alphabet string) error {
	if len(alphabet) < 2 {
		return ErrAlphabetTooSmall
	}
	seen := make(map[rune]bool)
	for _, c := range alphabet {
		if seen[c] {
			return ErrAlphabetCharactersNotUnique
		}
		seen[c] = true
	}
	return nil
}

func reverse(input string) string {
	rs, n := make([]rune, len(input)), 0
	for _, r := range input {
		rs[n] = r
		n++
	}
	rs = rs[0:n]
	for i := 0; i < n/2; i++ {
		rs[i], rs[n-1-i] = rs[n-1-i], rs[i]
	}
	return string(rs)
}

func uintPow(base, exponent uint) uint {
	if exponent == 0 {
		return 1
	}
	result := uint(1)
	for i := uint(0); i < exponent; i++ {
		result *= base
	}
	return result
}
