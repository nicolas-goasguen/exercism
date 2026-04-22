package luhn

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^[0-9]+$`)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	length := len(id)
	if length <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := range length {
		r := id[i]

		if r < '0' || r > '9' {
			return false
		}

		digit := int(r - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}
	return sum%10 == 0
}
