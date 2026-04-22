package hamming

import "fmt"

type ErrDifferentLengths struct {
	LenA int
	LenB int
}

func (e *ErrDifferentLengths) Error() string {
	return fmt.Sprintf("len(a): %d not equal to len(b): %d", e.LenA, e.LenB)
}

func Distance(a, b string) (int, error) {
	runesA := []rune(a)
	runesB := []rune(b)

	if len(runesA) != len(runesB) {
		return 0, &ErrDifferentLengths{len(runesA), len(runesB)}
	}

	var dist int
	for i := 0; i < len(a); i++ {
		if runesA[i] != runesB[i] {
			dist++
		}
	}

	return dist, nil
}
