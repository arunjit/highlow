package highlow

// Guess returns the number of guesses to guess the secret number within the range.
func Guess(secret, min, max int) int {
	guesses := 0
	low, high := min, max
	for low < high {
		mid := (low + high) / 2
		guesses = guesses + 1
		if mid < secret {
			low = mid
		} else if mid > secret {
			high = mid
		} else {
			break
		}
	}
	return guesses
}
