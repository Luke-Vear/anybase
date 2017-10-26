package anybase

import (
	"errors"
)

var (
	ErrAlphabetTooSmall = errors.New("alphabet cannot be len < 2")
)

type Generator interface {
	NumberFromUint(decimal uint) Number
	NumberFromEncoded(encoded string) Number
}

type generator struct {
	alphabet string
	values   map[byte]uint
	//max string
}

func (ng *generator) NumberFromUint(decimal uint) Number {
	return &number{
		decimal:  decimal,
		alphabet: ng.alphabet,
		values:   ng.values,
	}
}

// TODO: not yet working
// TODO: reject > max || recover
func (ng *generator) NumberFromEncoded(encoded string) Number {
	rev, decimal, base := reverse(encoded), uint(0), uint(len(encoded))
	for i := uint(0); i < base; i++ {
		decimal += ng.values[rev[i]] * uintPow(base, i)
	}
	return &number{
		decimal:  decimal,
		alphabet: ng.alphabet,
		values:   ng.values,
	}
}

func uintPow(base, exponent uint) uint {
	var result uint = 1
	for i := uint(0); i < exponent; i++ {
		result *= base
	}
	return result
}

type Number interface {
	Decimal() uint
	String() string
}

type number struct {
	decimal  uint
	alphabet string
	values   map[byte]uint
}

func (n *number) Decimal() uint { return n.decimal }

func (n *number) String() string {
	if n.decimal == 0 {
		return string(n.alphabet[0])
	}
	value, base := n.decimal, uint(len(n.alphabet))
	var encoded []byte
	for value > 0 {
		rem := value % base
		value = (value - rem) / base
		encoded = append(encoded, n.alphabet[rem])
	}
	return reverse(string(encoded))
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

// TODO: check all characters unique.
func NewGenerator(alphabet string) (Generator, error) {
	if len(alphabet) < 2 {
		return nil, ErrAlphabetTooSmall
	}
	values := make(map[byte]uint, len(alphabet))
	for i := 0; i < len(alphabet); i++ {
		values[alphabet[i]] = uint(i)
	}
	return &generator{
		alphabet: alphabet,
		values:   values,
	}, nil
}
