package game

import "math/rand"

func weightedRandom(p int, max int) int {
	r := rand.Intn(max)

	switch {
	case r < p/5:
		return 4
	case r < p/5*2:
		return 3
	case r < p/5*3:
		return 2
	case r < p:
		return 1
	default:
		return 0
	}
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
