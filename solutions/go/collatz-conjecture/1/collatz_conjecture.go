package collatzconjecture

import "fmt"

type ErrNegativeValue struct {
	Value int
}

func (e *ErrNegativeValue) Error() string {
	return fmt.Sprintf("input value cannot be negative or null: %d", e.Value)
}

func CollatzConjecture(n int) (int, error) {
	var count int
	if n <= 0 {
		return 0, &ErrNegativeValue{n}
	}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		count++
	}
	return count, nil
}
