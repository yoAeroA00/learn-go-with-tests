package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var ret string
	for i := 0; i < repeatCount; i++ {
		ret += character
	}
	return ret
}

func Repeat2(character string) string {
	var ret strings.Builder
	for i := 0; i < repeatCount; i++ {
		ret.WriteString(character)
	}
	return ret.String()
}
