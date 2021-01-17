package conv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/baseconv/conv"
)

type testCase struct {
	name string
	num  string
	from conv.Base
	to   conv.Base
	res  string
	err  string
}

func TestConvert(t *testing.T) {
	t.Parallel()

	testCases := append(getTestCasesForBase2(), getTestCasesForBase4()...)
	testCases = append(testCases, getTestCasesForBase8()...)
	testCases = append(testCases, getTestCasesForBase10()...)
	testCases = append(testCases, getTestCasesForBase16()...)
	testCases = append(testCases, getTestCasesForBase36()...)
	testCases = append(testCases, getTestCasesForBase62()...)
	testCases = append(testCases, getTestCasesForErrors()...)
	testCases = append(testCases, getTestCasesForRareCases()...)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := conv.Convert(tc.num, tc.from, tc.to)
			assert.Equal(t, tc.res, res)

			if tc.err != "" || err != nil {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}

func getTestCasesForBase2() []testCase {
	return []testCase{
		{"base2-to-base2", "1111001010001101", conv.Base2, conv.Base2, "1111001010001101", ""},
		{"base2-to-base4", "1111001010001101", conv.Base2, conv.Base4, "33022031", ""},
		{"base2-to-base8", "1111001010001101", conv.Base2, conv.Base8, "171215", ""},
		{"base2-to-base10", "1111001010001101", conv.Base2, conv.Base10, "62093", ""},
		{"base2-to-base16", "1111001010001101", conv.Base2, conv.Base16, "F28D", ""},
		{"base2-to-base36", "1111001010001101", conv.Base2, conv.Base36, "1BWT", ""},
		{"base2-to-base62", "1111001010001101", conv.Base2, conv.Base62, "G9V", ""},
	}
}

func getTestCasesForBase4() []testCase {
	return []testCase{
		{"base4-to-base2", "11311320", conv.Base4, conv.Base2, "101110101111000", ""},
		{"base4-to-base4", "11311320", conv.Base4, conv.Base4, "11311320", ""},
		{"base4-to-base8", "11311320", conv.Base4, conv.Base8, "56570", ""},
		{"base4-to-base10", "11311320", conv.Base4, conv.Base10, "23928", ""},
		{"base4-to-base16", "11311320", conv.Base4, conv.Base16, "5D78", ""},
		{"base4-to-base36", "11311320", conv.Base4, conv.Base36, "IGO", ""},
		{"base4-to-base62", "11311320", conv.Base4, conv.Base62, "6Dw", ""},
	}
}

func getTestCasesForBase8() []testCase {
	return []testCase{
		{"base8-to-base2", "67601", conv.Base8, conv.Base2, "110111110000001", ""},
		{"base8-to-base4", "67601", conv.Base8, conv.Base4, "12332001", ""},
		{"base8-to-base8", "67601", conv.Base8, conv.Base8, "67601", ""},
		{"base8-to-base10", "67601", conv.Base8, conv.Base10, "28545", ""},
		{"base8-to-base16", "67601", conv.Base8, conv.Base16, "6F81", ""},
		{"base8-to-base36", "67601", conv.Base8, conv.Base36, "M0X", ""},
		{"base8-to-base62", "67601", conv.Base8, conv.Base62, "7QP", ""},
	}
}

func getTestCasesForBase10() []testCase {
	return []testCase{
		{"base10-to-base2", "17434", conv.Base10, conv.Base2, "100010000011010", ""},
		{"base10-to-base4", "17434", conv.Base10, conv.Base4, "10100122", ""},
		{"base10-to-base8", "17434", conv.Base10, conv.Base8, "42032", ""},
		{"base10-to-base10", "17434", conv.Base10, conv.Base10, "17434", ""},
		{"base10-to-base16", "17434", conv.Base10, conv.Base16, "441A", ""},
		{"base10-to-base36", "17434", conv.Base10, conv.Base36, "DGA", ""},
		{"base10-to-base62", "17434", conv.Base10, conv.Base62, "4XC", ""},
	}
}

func getTestCasesForBase16() []testCase {
	return []testCase{
		{"base16-to-base2", "552B", conv.Base16, conv.Base2, "101010100101011", ""},
		{"base16-to-base4", "552B", conv.Base16, conv.Base4, "11110223", ""},
		{"base16-to-base8", "552B", conv.Base16, conv.Base8, "52453", ""},
		{"base16-to-base10", "552B", conv.Base16, conv.Base10, "21803", ""},
		{"base16-to-base16", "552B", conv.Base16, conv.Base16, "552B", ""},
		{"base16-to-base36", "552B", conv.Base16, conv.Base36, "GTN", ""},
		{"base16-to-base62", "552B", conv.Base16, conv.Base62, "5ff", ""},
	}
}

func getTestCasesForBase36() []testCase {
	return []testCase{
		{"base36-to-base2", "HQO", conv.Base36, conv.Base2, "101100111010000", ""},
		{"base36-to-base4", "HQO", conv.Base36, conv.Base4, "11213100", ""},
		{"base36-to-base8", "HQO", conv.Base36, conv.Base8, "54720", ""},
		{"base36-to-base10", "HQO", conv.Base36, conv.Base10, "22992", ""},
		{"base36-to-base16", "HQO", conv.Base36, conv.Base16, "59D0", ""},
		{"base36-to-base36", "HQO", conv.Base36, conv.Base36, "HQO", ""},
		{"base36-to-base62", "HQO", conv.Base36, conv.Base62, "5yq", ""},
	}
}

func getTestCasesForBase62() []testCase {
	return []testCase{
		{"base62-to-base2", "6zr", conv.Base62, conv.Base2, "110100100010011", ""},
		{"base62-to-base4", "6zr", conv.Base62, conv.Base4, "12210103", ""},
		{"base62-to-base8", "6zr", conv.Base62, conv.Base8, "64423", ""},
		{"base62-to-base10", "6zr", conv.Base62, conv.Base10, "26899", ""},
		{"base62-to-base16", "6zr", conv.Base62, conv.Base16, "6913", ""},
		{"base62-to-base36", "6zr", conv.Base62, conv.Base36, "KR7", ""},
		{"base62-to-base62", "6zr", conv.Base62, conv.Base62, "6zr", ""},
	}
}

func getTestCasesForRareCases() []testCase {
	return []testCase{
		{"zero", "0", conv.Base10, conv.Base2, "0", ""},
		{"negative", "-19530", conv.Base10, conv.Base2, "100110001001010", ""},
	}
}

func getTestCasesForErrors() []testCase {
	return []testCase{
		{"empty-num-string", "", conv.Base2, conv.Base10, "", "converting error: number string is empty"},
		{"bad-num-string", "A", conv.Base2, conv.Base10, "", "converting error: rune A is not found in alphabet"},
		{"bad-dec-string", "A", conv.Base10, conv.Base2, "", "converting error: strconv.Atoi: parsing \"A\": invalid syntax"},
		{"unsupported-from-base", "100", 11111, conv.Base10, "", "base 11111 is unsupported"},
		{"unsupported-to-base", "100", conv.Base10, 22222, "", "base 22222 is unsupported"},
	}
}

func TestConvertFromCustomBase(t *testing.T) {
	t.Parallel()

	alphabet := conv.Alphabet{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'h'}

	testCases := []struct {
		name string
		num  string
		from conv.Alphabet
		to   conv.Base
		res  string
		err  string
	}{
		{"base2", "GhGAB", alphabet, conv.Base2, "110111110000001", ""},
		{"base4", "GhGAB", alphabet, conv.Base4, "12332001", ""},
		{"base8", "GhGAB", alphabet, conv.Base8, "67601", ""},
		{"base10", "GhGAB", alphabet, conv.Base10, "28545", ""},
		{"base16", "GhGAB", alphabet, conv.Base16, "6F81", ""},
		{"base36", "GhGAB", alphabet, conv.Base36, "M0X", ""},
		{"base62", "GhGAB", alphabet, conv.Base62, "7QP", ""},
		{"empty-num-string", "", alphabet, conv.Base10, "", "converting error: number string is empty"},
		{"bad-num-string", "H", alphabet, conv.Base10, "", "converting error: rune H is not found in alphabet"},
		{"unsupported-base", "GhGAB", alphabet, 33333, "", "base 33333 is unsupported"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := conv.ConvertFromCustomBase(tc.num, tc.from, tc.to)
			assert.Equal(t, tc.res, res)

			if tc.err != "" || err != nil {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}

func TestConvertToCustomBase(t *testing.T) {
	t.Parallel()

	alphabet := conv.Alphabet{'A', 'B', 'C', 'D', 'E', 'F', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	testCases := []struct {
		name string
		num  string
		from conv.Base
		to   conv.Alphabet
		res  string
		err  string
	}{
		{"base2", "101010100101011", conv.Base2, alphabet, "FFC5", ""},
		{"base4", "11110223", conv.Base4, alphabet, "FFC5", ""},
		{"base8", "52453", conv.Base8, alphabet, "FFC5", ""},
		{"base10", "21803", conv.Base10, alphabet, "FFC5", ""},
		{"base16", "552B", conv.Base16, alphabet, "FFC5", ""},
		{"base36", "GTN", conv.Base36, alphabet, "FFC5", ""},
		{"base62", "5ff", conv.Base62, alphabet, "FFC5", ""},
		{"empty-num-string", "", conv.Base2, alphabet, "", "converting error: number string is empty"},
		{"bad-num-string", "h", conv.Base2, alphabet, "", "converting error: rune h is not found in alphabet"},
		{"bad-dec-string", "h", conv.Base10, alphabet, "", "converting error: strconv.Atoi: parsing \"h\": invalid syntax"},
		{"unsupported-base", "5ff", 44444, alphabet, "", "base 44444 is unsupported"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := conv.ConvertToCustomBase(tc.num, tc.from, tc.to)
			assert.Equal(t, tc.res, res)

			if tc.err != "" || err != nil {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}
