package romannumerals

import (
	"testing"
	"testing/quick"
)

func TestPropertiesOfConversion(t *testing.T) {
	t.Run("general and overflow check", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			t.Log("testing", arabic)
			if arabic > 3999 {
				_, err := ConvertToRoman(arabic)
				return err == ErrArabicOverflow
			}
			roman, _ := ConvertToRoman(arabic)
			fromRoman := ConvertToArabic(roman)
			return fromRoman == arabic
		}

		if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
			t.Error("failed checks", err)
		}
	})
	t.Run("Check that only I (1), X (10) and C (100) can be \"subtractors\"", func(t *testing.T) {
		assertion := func(arabic uint8) bool {
			if arabic == 'I' || arabic == 'X' || arabic == 'C' {
				return isSubtractive(arabic) == true
			} else {
				return isSubtractive(arabic) == false
			}
		}
		if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
			t.Error("failed checks", err)
		}
	})
	t.Run("Can't have more than 3 consecutive symbols", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			roman, _ := ConvertToRoman(arabic)
			count := 1
			for i := 0; i < len(roman)-1; i++ {
				if roman[i] == roman[i+1] {
					count++
				} else {
					count = 1
				}
			}
			return count < 4
		}
		if err := quick.Check(assertion, nil); err != nil {
			t.Error("failed checks", err)
		}
	})
}
