package armstrongnumbers

import (
	"math"
	"strconv"
)

func intPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func IsNumber(n int) bool {
	nCpy := n
	length := len(strconv.Itoa(nCpy))
	sum := 0
	for range length {
		sum += intPow(nCpy%10, length)
		nCpy /= 10
	}
	return sum == n
}
