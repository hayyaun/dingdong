package game

import "math/rand"

type Weightable struct {
	v any // value
	w int // wright
}

func weightedRandom(items []Weightable) any {

	// Calculate total weight
	total := 0
	for _, item := range items {
		total += item.w
	}

	// Generate a random weight
	r := rand.Intn(total)

	// Find random position inside items range
	cumu := 0
	for _, item := range items {
		cumu += item.w
		if r < cumu {
			return item.v
		}
	}

	// Fallback value - never
	return items[0].v
}
