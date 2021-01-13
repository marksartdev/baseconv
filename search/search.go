// Package search searches elements in slices.
package search

// IndexOfRune return index of rune in alphabet.
// Attention! Alphabet must be sorted.
func IndexOfRune(alphabet []rune, r rune) (int, bool) {
	return binarySearchRune(alphabet, r, 0, len(alphabet)-1)
}

// Return index of rune in alphabet with binary search.
func binarySearchRune(alphabet []rune, r rune, from, to int) (int, bool) {
	if from > to {
		return 0, false
	}

	mid := from + (to-from)/2 //nolint:gomnd // it is average
	if r == alphabet[mid] {
		return mid, true
	}

	if r > alphabet[mid] {
		return binarySearchRune(alphabet, r, mid+1, to)
	}

	return binarySearchRune(alphabet, r, from, mid-1)
}
