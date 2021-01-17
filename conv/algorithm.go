package conv

import (
	"math"

	"github.com/marksartdev/baseconv/search"
)

// Convert 10-base number to custom base number.
func conv10baseToXbase(num int, alphabet Alphabet) string {
	base := len(alphabet)

	if num < 0 {
		num = -num
	}

	remainders := make([]int, 0)

	if num == 0 {
		remainders = append(remainders, num)
	}

	for num > 0 {
		remainder := num % base
		num = (num - remainder) / base
		remainders = append([]int{remainder}, remainders...)
	}

	return convRemaindersToString(remainders, alphabet)
}

// Convert indexes of alphabet symbols to string.
func convRemaindersToString(idxs []int, alphabet Alphabet) string {
	res := ""

	for _, idx := range idxs {
		res += string(alphabet[idx])
	}

	return res
}

// Convert custom base number to 10-base number.
func convXbaseTo10base(num string, alphabet Alphabet) (int, error) {
	if num == "" {
		return 0, ErrEmptyNumberString{}
	}

	base := len(alphabet)

	res := 0
	pow := len(num) - 1

	for _, r := range num {
		i, ok := search.IndexOfRune(alphabet, r)
		if !ok {
			return 0, ErrIncorrectRune{r}
		}

		res += i * int(math.Pow(float64(base), float64(pow)))
		pow--
	}

	return res, nil
}
