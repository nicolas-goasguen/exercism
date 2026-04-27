package sumofmultiples

import (
	"slices"
)

func findMultiples(limit int, divisor int) []int {
	if divisor == 0 || divisor >= limit {
		return nil
	}

	n := limit / divisor
	if limit%divisor == 0 {
		n--
	}

	multiples := make([]int, 0, n)
	for i := range n {
		multiples = append(multiples, (i+1)*divisor)
	}

	return multiples
}

func SumMultiples(limit int, divisors ...int) int {
	var multiples []int
	for _, d := range divisors {
		divMultiples := findMultiples(limit, d)
		multiples = append(multiples, divMultiples...)
	}

	slices.Sort(multiples)
	multiples = slices.Compact(multiples)

	sum := 0
	for _, v := range multiples {
		sum += v
	}

	return sum
}
