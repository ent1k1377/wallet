package math

import "math/rand"

func RandomInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
