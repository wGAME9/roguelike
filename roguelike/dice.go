package roguelike

import "math/rand/v2"

func getRandomIntBetween(min, max int) int {
	return getDiceRoll(max-min) + min
}

// getRandomInt returns an integer from 0 to the number - 1
func getRandomInt(num int) int {
	return rand.IntN(num)
}

// getDiceRoll returns an integer from 1 to the number
func getDiceRoll(num int) int {
	return rand.IntN(num) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
