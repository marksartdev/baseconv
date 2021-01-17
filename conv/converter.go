// Package conv provides functions for converting numbers from one base to another.
// Supported 2-, 4-, 8-, 10-, 16-, 36-, 62-base numbers.
// Also is able to convert numbers with custom alphabets.
package conv

import (
	"fmt"
	"strconv"
)

const (
	// Base2 is a binary number system.
	Base2 = 2
	// Base4 is a quaternary number system.
	Base4 = 4
	// Base8 is a octal number system.
	Base8 = 8
	// Base10 is a decimal number system.
	Base10 = 10
	// Base16 is hexadecimal number system.
	Base16 = 16
	// Base36 is a hexatrigesimal number system.
	Base36 = 36
	// Base62 is a duosexagesimal number system.
	Base62 = 62
)

// Base of number.
type Base uint

// Alphabet for converting.
type Alphabet []rune

// Convert number from one base to another.
func Convert(num string, from, to Base) (string, error) {
	err := checkBases(from, to)
	if err != nil {
		return "", err
	}

	dec := 0

	if from == Base10 {
		dec, err = strconv.Atoi(num)
	} else {
		dec, err = convXbaseTo10base(num, alphabets[from])
	}

	if err != nil {
		return "", fmt.Errorf("converting error: %w", err)
	}

	if to == Base10 {
		return strconv.Itoa(dec), nil
	}

	return conv10baseToXbase(dec, alphabets[to]), nil
}

// ConvertFromCustomBase converts number from custom base using custom alphabet.
func ConvertFromCustomBase(num string, fromAlphabet Alphabet, to Base) (string, error) {
	err := checkBases(to)
	if err != nil {
		return "", err
	}

	dec, err := convXbaseTo10base(num, fromAlphabet)
	if err != nil {
		return "", fmt.Errorf("converting error: %w", err)
	}

	if to == Base10 {
		return strconv.Itoa(dec), nil
	}

	return conv10baseToXbase(dec, alphabets[to]), nil
}

// ConvertToCustomBase converts number to custom base using custom alphabet.
func ConvertToCustomBase(num string, from Base, toAlphabet Alphabet) (string, error) {
	err := checkBases(from)
	if err != nil {
		return "", err
	}

	dec := 0

	if from == Base10 {
		dec, err = strconv.Atoi(num)
	} else {
		dec, err = convXbaseTo10base(num, alphabets[from])
	}

	if err != nil {
		return "", fmt.Errorf("converting error: %w", err)
	}

	return conv10baseToXbase(dec, toAlphabet), nil
}

// Check support for bases.
func checkBases(bases ...Base) error {
	support := make(map[Base]bool)

	for _, base := range bases {
		support[base] = false
	}

	for base := range alphabets {
		if _, ok := support[base]; ok {
			support[base] = true
		}
	}

	for base, isSupported := range support {
		if !isSupported {
			return ErrUnsupportedBase{base}
		}
	}

	return nil
}
