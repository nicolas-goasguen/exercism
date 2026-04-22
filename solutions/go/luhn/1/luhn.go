package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^[0-9]+$`)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	length := len(id)

	if !re.MatchString(id) {
		return false
	}
	if length <= 1 {
		return false
	}

	var sum int
	var offset int
	offset = length % 2

	for i := range length {
		numStr := id[i]
		num, err := strconv.Atoi(string(numStr))
		if err != nil {
			return false
		}

		if i%2 == offset {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	return sum%10 == 0
}
