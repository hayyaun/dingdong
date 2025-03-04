package game

import "math/rand"

func weightedRandom(p int, max int) int {
	r := rand.Intn(max)

	switch {
	case r < p/10:
		return 4
	case r < p/10*2:
		return 3
	case r < p/10*5:
		return 2
	case r < p:
		return 1
	default:
		return 0
	}
}
