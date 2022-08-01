package propertybasedtests

import (
	"fmt"
	"testing"
)

var (
	cases = []struct {
		ArabicNum int
		Roman     string
	}{
		{ArabicNum: 1, Roman: "I"},
		{ArabicNum: 2, Roman: "II"},
		{ArabicNum: 3, Roman: "III"},
		{ArabicNum: 4, Roman: "IV"},
		{ArabicNum: 5, Roman: "V"},
		{ArabicNum: 6, Roman: "VI"},
		{ArabicNum: 7, Roman: "VII"},
		{ArabicNum: 8, Roman: "VIII"},
		{ArabicNum: 9, Roman: "IX"},
		{ArabicNum: 10, Roman: "X"},
		{ArabicNum: 14, Roman: "XIV"},
		{ArabicNum: 18, Roman: "XVIII"},
		{ArabicNum: 20, Roman: "XX"},
		{ArabicNum: 39, Roman: "XXXIX"},
		{ArabicNum: 40, Roman: "XL"},
		{ArabicNum: 47, Roman: "XLVII"},
		{ArabicNum: 49, Roman: "XLIX"},
		{ArabicNum: 50, Roman: "L"},
		{ArabicNum: 100, Roman: "C"},
		{ArabicNum: 90, Roman: "XC"},
		{ArabicNum: 400, Roman: "CD"},
		{ArabicNum: 500, Roman: "D"},
		{ArabicNum: 900, Roman: "CM"},
		{ArabicNum: 1000, Roman: "M"},
		{ArabicNum: 1984, Roman: "MCMLXXXIV"},
		{ArabicNum: 3999, Roman: "MMMCMXCIX"},
		{ArabicNum: 2014, Roman: "MMXIV"},
		{ArabicNum: 1006, Roman: "MVI"},
		{ArabicNum: 798, Roman: "DCCXCVIII"},
		{ArabicNum: 1984, Roman: "MCMLXXXIV"},
	}
)

func TestRomanNumerals(t *testing.T) {
	// cases := []struct {
	// 	ArabicNum int
	// 	Roman     string
	// }{
	// 	{ArabicNum: 1, Roman: "I"},
	// 	{ArabicNum: 2, Roman: "II"},
	// 	{ArabicNum: 3, Roman: "III"},
	// 	{ArabicNum: 4, Roman: "IV"},
	// 	{ArabicNum: 5, Roman: "V"},
	// 	{ArabicNum: 6, Roman: "VI"},
	// 	{ArabicNum: 7, Roman: "VII"},
	// 	{ArabicNum: 8, Roman: "VIII"},
	// 	{ArabicNum: 9, Roman: "IX"},
	// 	{ArabicNum: 10, Roman: "X"},
	// 	{ArabicNum: 14, Roman: "XIV"},
	// 	{ArabicNum: 18, Roman: "XVIII"},
	// 	{ArabicNum: 20, Roman: "XX"},
	// 	{ArabicNum: 39, Roman: "XXXIX"},
	// 	{ArabicNum: 40, Roman: "XL"},
	// 	{ArabicNum: 47, Roman: "XLVII"},
	// 	{ArabicNum: 49, Roman: "XLIX"},
	// 	{ArabicNum: 50, Roman: "L"},
	// 	{ArabicNum: 100, Roman: "C"},
	// 	{ArabicNum: 90, Roman: "XC"},
	// 	{ArabicNum: 400, Roman: "CD"},
	// 	{ArabicNum: 500, Roman: "D"},
	// 	{ArabicNum: 900, Roman: "CM"},
	// 	{ArabicNum: 1000, Roman: "M"},
	// 	{ArabicNum: 1984, Roman: "MCMLXXXIV"},
	// 	{ArabicNum: 3999, Roman: "MMMCMXCIX"},
	// 	{ArabicNum: 2014, Roman: "MMXIV"},
	// 	{ArabicNum: 1006, Roman: "MVI"},
	// 	{ArabicNum: 798, Roman: "DCCXCVIII"},
	// 	{ArabicNum: 1984, Roman: "MCMLXXXIV"},
	// }

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.ArabicNum, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.ArabicNum)
			if got != test.Roman {
				t.Errorf("got: %q want: %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.ArabicNum), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.ArabicNum {
				t.Errorf("got %d, want %d", got, test.ArabicNum)
			}
		})
	}
}
