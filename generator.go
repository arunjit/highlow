package highlow

import (
	"math/rand"
	"time"
)

// GenerateSecret generates a random secret number within the given range
func GenerateSecret(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
