package hamming

import "fmt"

type ErrDifferentLengths struct {
	lenA int
	lenB int
}

func (e *ErrDifferentLengths) Error() string {
	return fmt.Sprintf("len(a): %d not equal to len(b): %d", e.lenA, e.lenB)
}

func Distance(a, b string) (int, error) {
	var dist int
	if len(a) != len(b) {
		return 0, &ErrDifferentLengths{len(a), len(b)}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist += 1
		}
	}
	return dist, nil
}
