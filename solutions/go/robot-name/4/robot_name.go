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

func idToName(id int) string {
	letters := id / 1000
	numbers := id % 1000

	return string([]byte{
		byte('A' + letters/26),
		byte('A' + letters%26),
		byte('0' + numbers/100),
		byte('0' + (numbers/10)%10),
		byte('0' + numbers%10),
	})
}

func robotIdGenerator() func() (int, error) {
	nCombinations := 26 * 26 * 10 * 10 * 10
	available := make([]int, nCombinations)

	for i := range nCombinations {
		available[i] = i
	}

	return func() (int, error) {
		var id int

		nAvailable := len(available)
		if nAvailable == 0 {
			return 0, ErrNoAvailableNames
		}

		// swap-delete for O(1)
		n := rand.Intn(nAvailable)
		id = available[n]
		last := nAvailable - 1
		available[n] = available[last]
		available = available[:last]

		return id, nil
	}
}

var generateId = robotIdGenerator()

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		id, err := generateId()
		if err != nil {
			return "", err
		}
		r.name = idToName(id)
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
