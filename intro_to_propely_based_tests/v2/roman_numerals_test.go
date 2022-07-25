package romannumerals

import "testing"

type testCase struct {
	description string
	arabic      int
	roman       string
}

var testCases = []testCase{
	{"1 to I", 1, "I"},
	{"2 to II", 2, "II"},
	{"3 to III", 3, "III"},
	{"4 to IV", 4, "IV"},
	{"5 to V", 5, "V"},
	{"6 to VI", 6, "VI"},
	{"7 to VII", 7, "VII"},
	{"8 to VIII", 8, "VIII"},
	{"9 to IX", 9, "IX"},
	{"10 to X", 10, "X"},
	{"11 to XI", 11, "XI"},
	{"12 to XII", 12, "XII"},
	{"14 to XIV", 14, "XIV"},
	{"18 to XVIII", 18, "XVIII"},
	{"20 to XX", 20, "XX"},
	{"39 to XXXIX", 39, "XXXIX"},
	{"40 to XL", 40, "XL"},
	{"47 to XLVII", 47, "XLVII"},
	{"49 to XLIX", 49, "XLIX"},
	{"50 to L", 50, "L"},
	{"100 to C", 100, "C"},
	{"90 to XC", 90, "XC"},
	{"400 to CD", 400, "CD"},
	{"500 to D", 500, "D"},
	{"900 to CM", 900, "CM"},
	{"1000 to M", 1000, "M"},
	{"1984 to L", 1984, "MCMLXXXIV"},
	{"3999 to MMMCMXCIX", 3999, "MMMCMXCIX"},
	{"2014 to MMXIV", 2014, "MMXIV"},
	{"1006 to MVI", 1006, "MVI"},
	{"798 to DCCXCVIII", 798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := ConvertToRoman(tc.arabic)
			if got != tc.roman {
				t.Errorf("got %q but want %q", got, tc.roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := ConvertToArabic(tc.roman)
			if got != tc.arabic {
				t.Errorf("got %d but want %d", got, tc.arabic)
			}
		})
	}
}
