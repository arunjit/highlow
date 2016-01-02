package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arunjit/highlow"
)

var (
	minFlag = flag.Int("min", 1, "Lowerbound of the guessing range")
	maxFlag = flag.Int("max", 1001, "Upperbound of the guessing range")
)

func main() {
	flag.Parse()
	min, max := *minFlag, *maxFlag
	if min < 1 || min >= max {
		fmt.Errorf("min must be >0 and <max")
		os.Exit(1)
	}
	secret := highlow.GenerateSecret(min, max)
	guesses := highlow.CountGuesses(secret, min, max)
	fmt.Printf("%d guesses to guess %d\n", guesses, secret)
}
