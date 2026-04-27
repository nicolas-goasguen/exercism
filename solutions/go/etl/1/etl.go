package etl

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int, 26)
	for amount, letters := range in {
		for _, l := range letters {
			out[l] = amount
		}
	}
	return out
}
