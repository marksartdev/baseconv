package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/baseconv/search"
)

// nolint:gochecknoglobals // It is constant.
var alphabet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func TestIndexOfRune(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		alphabet []rune
		name     string
		index    int
		target   rune
		status   bool
	}{
		{[]rune{}, "empty", 0, 'a', false},
		{[]rune{'a'}, "1-symbol:founded", 0, 'a', true},
		{[]rune{'a'}, "1-symbol:not-founded", 0, 'b', false},
		{[]rune{'a', 'b'}, "2-symbols:founded-0", 0, 'a', true},
		{[]rune{'a', 'b'}, "2-symbols:founded-1", 1, 'b', true},
		{[]rune{'a', 'b'}, "2-symbols:not-founded", 0, 'c', false},
		{[]rune{'a', 'b', 'c'}, "3-symbols:founded-0", 0, 'a', true},
		{[]rune{'a', 'b', 'c'}, "3-symbols:founded-1", 1, 'b', true},
		{[]rune{'a', 'b', 'c'}, "3-symbols:founded-2", 2, 'c', true},
		{[]rune{'a', 'b', 'c'}, "3-symbols:not-founded", 0, 'd', false},
		{alphabet, "full:founded-0", 0, '0', true},
		{alphabet, "full:founded-1", 5, '5', true},
		{alphabet, "full:founded-2", 9, '9', true},
		{alphabet, "full:founded-3", 10, 'A', true},
		{alphabet, "full:founded-4", 15, 'F', true},
		{alphabet, "full:founded-5", 22, 'M', true},
		{alphabet, "full:founded-6", 23, 'N', true},
		{alphabet, "full:founded-7", 28, 'S', true},
		{alphabet, "full:founded-8", 35, 'Z', true},
		{alphabet, "full:founded-9", 36, 'a', true},
		{alphabet, "full:founded-10", 41, 'f', true},
		{alphabet, "full:founded-11", 48, 'm', true},
		{alphabet, "full:founded-12", 49, 'n', true},
		{alphabet, "full:founded-13", 54, 's', true},
		{alphabet, "full:founded-14", 61, 'z', true},
		{alphabet, "full:not-founded", 0, '~', false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Helper()
			t.Parallel()

			idx, ok := search.IndexOfRune(tc.alphabet, tc.target)
			assert.Equal(t, tc.index, idx)
			assert.Equal(t, tc.status, ok)
		})
	}
}
