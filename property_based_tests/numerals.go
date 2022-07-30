package propertybasedtests

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}
type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabicNum int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabicNum >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabicNum -= numeral.Value
		}
	}

	// for arabicNum > 0 {
	// 	switch {
	// 	case arabicNum > 9:
	// 		result.WriteString("X")
	// 		arabicNum -= 10
	// 	case arabicNum > 8:
	// 		result.WriteString("IX")
	// 		arabicNum -= 9
	// 	case arabicNum > 4:
	// 		result.WriteString("V")
	// 		arabicNum -= 5
	// 	case arabicNum > 3:
	// 		result.WriteString("IV")
	// 		arabicNum -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabicNum--
	// 	}
	// }

	// for i := arabicNum; i > 0; i-- {
	// 	if arabicNum == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}
	// 	if arabicNum == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}
	// 	result.WriteString("I")
	// }

	return result.String()

}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if i+1 < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			// build the two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			// get the value of the two character string
			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total++
			}
		} else {
			total++
		}
	}
	return total
}
