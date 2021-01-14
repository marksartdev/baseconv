package conv

import "fmt"

// ErrUnsupportedBase is a error which is threw when base is unsupported.
type ErrUnsupportedBase struct {
	base Base
}

func (e ErrUnsupportedBase) Error() string {
	return fmt.Sprintf("base %d is unsupported", e.base)
}

// ErrEmptyNumberString is a error which is threw when string containing number is empty.
type ErrEmptyNumberString struct{}

func (ErrEmptyNumberString) Error() string {
	return "number string is empty"
}

// ErrIncorrectRune is a error which is threw when rune is not in alphabet.
type ErrIncorrectRune struct {
	r rune
}

func (e ErrIncorrectRune) Error() string {
	return fmt.Sprintf("rune %s is not found in alphabet", string(e.r))
}
