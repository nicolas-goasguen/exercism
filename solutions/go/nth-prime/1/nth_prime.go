package nthprime

import (
	"fmt"
)

type ErrNegative struct {
	n int
}

func (e *ErrNegative) Error() string {
	return fmt.Sprintf("'n' is equal or less than zero: %d", e.n)
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, &ErrNegative{n: n}
	}
	count := 0
	candidate := 1

	for count < n {
		candidate++
		isPrime := true
		for i := 2; i*i <= candidate; i++ {
			if candidate%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	return candidate, nil
}
