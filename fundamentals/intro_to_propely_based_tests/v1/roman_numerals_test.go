package romannumerals

import "testing"

func TestConvertToRoman(t *testing.T) {
	testCases := []struct {
		description string
		arabic      int
		want        string
	}{
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
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := ConvertToRoman(tc.arabic)
			if got != tc.want {
				t.Errorf("got %q but want %q", got, tc.want)
			}
		})
	}
}
