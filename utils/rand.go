package utils

import "math/rand"

type Weightable interface {
	Value() any
	Weight() int
}

func WeightedRandom(items []Weightable) any {

	// Calculate total weight
	total := 0
	for _, item := range items {
		total += item.Weight()
	}

	// Generate a random weight
	r := rand.Intn(total)

	// Find random position inside items range
	cumu := 0
	for _, item := range items {
		cumu += item.Weight()
		if r < cumu {
			return item.Value()
		}
	}

	// Fallback value - never
	return items[0].Value()
}
