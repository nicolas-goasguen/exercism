package reversestring

import "slices"

func Reverse(input string) string {
	output := []rune(input)
	slices.Reverse(output)
	return string(output)
}
