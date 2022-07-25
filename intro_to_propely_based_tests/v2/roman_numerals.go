package romannumerals

import (
	"strings"
)

type RomanNum struct {
	symbol string
	value  int
}

var allRomanNumerals = []RomanNum{
	{symbol: "M", value: 1000},
	{symbol: "CM", value: 900},
	{symbol: "D", value: 500},
	{symbol: "CD", value: 400},
	{symbol: "C", value: 100},
	{symbol: "XC", value: 90},
	{symbol: "L", value: 50},
	{symbol: "XL", value: 40},
	{symbol: "X", value: 10},
	{symbol: "IX", value: 9},
	{symbol: "V", value: 5},
	{symbol: "IV", value: 4},
	{symbol: "I", value: 1},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.value {
			result.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	var result int
	for _, numeral := range allRomanNumerals {
		if len(numeral.symbol) <= len(roman) {
			for roman[:len(numeral.symbol)] == numeral.symbol {
				result += numeral.value
				roman = roman[len(numeral.symbol):]
				if len(roman) == 0 {
					return result
				}
			}
		}
	}
	return 0
}
