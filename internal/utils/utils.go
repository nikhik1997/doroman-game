package utils

import (
	"math/rand"
)

// RandomInt returns a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
