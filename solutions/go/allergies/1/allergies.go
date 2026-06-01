package allergies

import (
	"slices"
	"strconv"
)

var (
	allergiesIdx = map[string]int{
		"eggs":         0,
		"peanuts":      1,
		"shellfish":    2,
		"strawberries": 3,
		"tomatoes":     4,
		"chocolate":    5,
		"pollen":       6,
		"cats":         7,
	}
	allergiesStr = []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats",
	}
)

func toBits(allergies uint) []rune {
	bits := []rune(strconv.FormatUint(uint64(allergies), 2))
	slices.Reverse(bits)
	return bits
}

func Allergies(allergies uint) []string {
	bits := toBits(allergies)

	var result []string
	for i, b := range bits {
		if i >= len(allergiesStr) {
			break
		}
		if b == '1' {
			result = append(result, allergiesStr[i])
		}
	}

	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	i, ok := allergiesIdx[allergen]
	if !ok {
		return false
	}
	bits := toBits(allergies)
	if i >= len(bits) {
		return false
	}
	return bits[i] == '1'
}
