package wordy

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`(\-?\d+)([^\-\d\?]*)`)
)

var operations = map[string]func(a, b int) (int, error){
	"plus":          add,
	"minus":         subtract,
	"multiplied by": multiply,
	"divided by":    divide,
}

func add(a, b int) (int, error) {
	return a + b, nil
}

func subtract(a, b int) (int, error) {
	return a - b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

var ErrDevisionByZero = errors.New("division by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDevisionByZero
	}
	return a / b, nil
}

func tokenize(question string) ([]int, []func(a, b int) (int, error)) {
	var nums []int
	var ops []func(a, b int) (int, error)

	matches := re.FindAllStringSubmatch(question, -1)
	for i, match := range matches {
		if num, err := strconv.Atoi(match[1]); err == nil {
			nums = append(nums, num)
		} else {
			return nil, nil
		}
		if i < len(matches)-1 {
			op, ok := operations[strings.TrimSpace(match[2])]
			if !ok {
				return nil, nil
			}
			ops = append(ops, op)
		} else if strings.TrimSpace(match[2]) != "" {
			return nil, nil
		}
	}
	if len(ops) != len(nums)-1 {
		return nil, nil
	}

	return nums, ops
}

func Answer(question string) (int, bool) {
	nums, ops := tokenize(question)

	if len(nums) == 0 {
		return 0, false
	}

	result := nums[0]
	nums = nums[1:]

	for i, num := range nums {
		op := ops[i]
		r, err := op(result, num)
		if err != nil {
			return 0, false
		}
		result = r
	}

	return result, true
}
