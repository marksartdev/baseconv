// Package conv convert number base.
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

// Convert number from one base to another base.
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

	return conv10baseToXbase(dec, alphabets[to])
}

// Check supported bases.
func checkBases(from, to Base) error {
	fromIsSupported, toIsSupported := false, false

	for base := range alphabets {
		if from == base {
			fromIsSupported = true
		}

		if to == base {
			toIsSupported = true
		}

		if fromIsSupported && toIsSupported {
			return nil
		}
	}

	if !fromIsSupported {
		return ErrUnsupportedBase{from}
	}

	return ErrUnsupportedBase{to}
}
