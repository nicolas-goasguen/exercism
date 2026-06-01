package allergies

// BenchmarkAllergies
// BenchmarkAllergies-10             276694              4325 ns/op            3136 B/op         76 allocs/op
// BenchmarkAllergicTo
// BenchmarkAllergicTo-10           3236265               370.2 ns/op             0 B/op          0 allocs/op

var allergens = map[string]uint{
	"eggs":         1 << 0,
	"peanuts":      1 << 1,
	"shellfish":    1 << 2,
	"strawberries": 1 << 3,
	"tomatoes":     1 << 4,
	"chocolate":    1 << 5,
	"pollen":       1 << 6,
	"cats":         1 << 7,
}

func Allergies(allergies uint) []string {
	var result []string
	for allergen, mask := range allergens {
		if allergies&mask != 0 {
			result = append(result, allergen)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	mask, ok := allergens[allergen]
	return ok && allergies&mask != 0
}
