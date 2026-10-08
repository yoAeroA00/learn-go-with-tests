package roman

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
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

func ConvertToRoman(nums int) string {
	var result strings.Builder

	// for i := nums; i > 0; i-- {
	// 	if i == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}
	// 	if i == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}
	// 	result.WriteString("I")
	// }

	// for nums > 0 {
	// 	switch {
	// 	case nums > 9:
	// 		result.WriteString("X")
	// 		nums -= 10
	// 	case nums > 8:
	// 		result.WriteString("IX")
	// 		nums -= 9
	// 	case nums > 4:
	// 		result.WriteString("V")
	// 		nums -= 5
	// 	case nums > 3:
	// 		result.WriteString("IV")
	// 		nums -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		nums--
	// 	}
	// }

	for _, numeral := range allRomanNumerals {
		for nums >= numeral.Value {
			result.WriteString(numeral.Symbol)
			nums -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) (result int) {
	// if roman == "III" {
	// 	return 3
	// }
	// if roman == "II" {
	// 	return 2
	// }
	// return 1

	// for range roman {
	// 	result++
	// }

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			result += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return
}
