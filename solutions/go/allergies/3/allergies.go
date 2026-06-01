package allergies

// BenchmarkAllergies
// BenchmarkAllergies-10             480588              2200 ns/op            3136 B/op         76 allocs/op
// BenchmarkAllergicTo
// BenchmarkAllergicTo-10           3129865               391.0 ns/op             0 B/op          0 allocs/op

var allergensIdx = map[string]int{
	"eggs":         0,
	"peanuts":      1,
	"shellfish":    2,
	"strawberries": 3,
	"tomatoes":     4,
	"chocolate":    5,
	"pollen":       6,
	"cats":         7,
}

var allergensStr = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
	var result []string
	for i, allergen := range allergensStr {
		if allergies&(1<<i) != 0 {
			result = append(result, allergen)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	idx, ok := allergensIdx[allergen]
	return ok && allergies&(1<<idx) != 0
}
