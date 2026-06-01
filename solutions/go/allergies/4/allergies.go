package allergies

// BenchmarkAllergies
// BenchmarkAllergies-10             544321              2370 ns/op            3136 B/op         76 allocs/op
// BenchmarkAllergicTo
// BenchmarkAllergicTo-10           5959016               205.3 ns/op             0 B/op          0 allocs/op

type allergen struct {
	name string
	mask uint
}

var allergens = []allergen{
	{"eggs", 1 << 0},
	{"peanuts", 1 << 1},
	{"shellfish", 1 << 2},
	{"strawberries", 1 << 3},
	{"tomatoes", 1 << 4},
	{"chocolate", 1 << 5},
	{"pollen", 1 << 6},
	{"cats", 1 << 7},
}

func Allergies(allergies uint) []string {
	var result []string
	for _, a := range allergens {
		if allergies&a.mask != 0 {
			result = append(result, a.name)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, a := range allergens {
		if a.name == allergen {
			return allergies&a.mask != 0
		}
	}
	return false
}
