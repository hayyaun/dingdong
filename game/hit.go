package game

import "slices"

type Status int

const (
	None Status = iota
	Good
	Meh
	Bad
)

type Hit struct {
	Rune    rune
	Status  Status
	Persist int
}

var hits = []*Hit{}

func updateHits() {
	for i, hit := range hits {
		if hit != nil && hit.Persist <= 0 {
			hits = slices.Delete(hits, i, i+1)
		}
		hit.Persist -= 1
	}
}

func hit(r rune, s Status) {
	hits = append(hits, &Hit{Rune: r, Status: s, Persist: 3})
}
