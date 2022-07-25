package romannumerals

import (
	"errors"
	"strings"
)

var ErrArabicOverflow = errors.New("unacceptable number (more than 3999)")

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", ErrArabicOverflow
	}
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.value {
			result.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}
	return result.String(), nil
}

func ConvertToArabic(roman string) (result uint16) {
	for _, symbol := range windowedRoman(roman).Symbols() {
		result += allRomanNumerals.ValueOf(symbol...)
	}
	return
}

type RomanNumeral struct {
	symbol string
	value  uint16
}

var allRomanNumerals = RomanNumerals{
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

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.symbol == symbol {
			return s.value
		}
	}

	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.symbol == symbol {
			return true
		}
	}

	return false
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		notAtEnd := i+1 < len(w)

		if isSubtractive(w[i]) && notAtEnd && allRomanNumerals.Exists(w[i], w[i+1]) {
			symbols = append(symbols, []byte{w[i], w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{w[i]})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
