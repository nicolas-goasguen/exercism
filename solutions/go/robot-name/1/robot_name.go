package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var ErrNoAvailableNames = errors.New("no available robot names")

func robotNameGenerator() func() (string, error) {
	nLetters := 26 * 26
	nNumbers := 10 * 10 * 10
	available := make(map[string]struct{}, nLetters*nNumbers)
	for l := range nLetters {
		letter1 := string(rune('A' + l/26))
		letter2 := string(rune('A' + l%26))
		for n := range nNumbers {
			number := fmt.Sprintf("%03d", n)
			key := letter1 + letter2 + number
			available[key] = struct{}{}
		}
	}
	return func() (string, error) {
		var key string
		nAvailable := len(available)
		if nAvailable == 0 {
			return "", ErrNoAvailableNames
		}
		n := rand.Intn(nAvailable)
		for k := range available {
			if n == 0 {
				key = k
			}
			n--
		}
		delete(available, key)
		return key, nil
	}
}

var generateName = robotNameGenerator()

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := generateName()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
