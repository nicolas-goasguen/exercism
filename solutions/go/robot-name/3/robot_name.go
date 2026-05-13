package robotname

import (
	"errors"
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
	available := make([]string, 0, nLetters*nNumbers)
	for l := range nLetters {
		for n := range nNumbers {
			key := string([]byte{
				byte('A' + l/26),
				byte('A' + l%26),
				byte('0' + n/100),
				byte('0' + (n/10)%10),
				byte('0' + n%10),
			})
			available = append(available, key)
		}
	}
	return func() (string, error) {
		var key string
		nAvailable := len(available)
		if nAvailable == 0 {
			return "", ErrNoAvailableNames
		}
		n := rand.Intn(nAvailable)

		// swap-delete for O(1)
		key = available[n]
		last := nAvailable - 1
		available[n] = available[last]
		available = available[:last]

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
