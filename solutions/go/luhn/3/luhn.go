package luhn

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^[0-9]+$`)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := len(id) - 1; i >= 0; i-- {
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
		double = !double
	}
	return sum%10 == 0
}
