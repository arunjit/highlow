package highlow

import (
	"math/rand"
	"time"
)

// GenerateSecret generates a random secret number within the given range
func GenerateSecret(min, max int) int {
	return rand.Intn(max-min) + min
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
