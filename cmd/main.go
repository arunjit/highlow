package main

import (
	"fmt"

	"github.com/arunjit/highlow"
)

const (
	min, max = 1, 1001
)

func main() {
	secret := highlow.GenerateSecret(min, max)
	guesses := highlow.Guess(secret, min, max)
	fmt.Printf("%d guesses to guess %d\n", guesses, secret)
}
