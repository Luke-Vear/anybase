package anybase

import "errors"

var (
	ErrAlphabetCharactersNotUnique = errors.New("alphabet characters must all be unique")
	ErrAlphabetTooSmall            = errors.New("alphabet cannot be len < 2")
)
