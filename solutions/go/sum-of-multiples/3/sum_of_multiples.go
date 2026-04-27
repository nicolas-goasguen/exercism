package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]struct{})
	for _, d := range divisors {
		if d != 0 {
			n := limit / d
			if limit%d == 0 {
				n--
			}
			for i := range n {
				multiples[(i+1)*d] = struct{}{}
			}
		}
	}

	sum := 0
	for m := range multiples {
		sum += m
	}

	return sum
}
