package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, x := range divisors {
			if x != 0 && i%x == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
