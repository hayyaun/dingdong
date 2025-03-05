package game

import (
	"slices"
)

type Hit struct {
	r       rune
	status  Status
	persist int
}

var hits = []*Hit{}

func updateHits() {
	for i, hit := range hits {
		if hit == nil {
			continue
		}
		if hit.persist <= 0 {
			hits = slices.Delete(hits, i, i+1)
		}
		hit.persist -= 1
	}
}

func hit(r rune, status Status) {
	hits = append(hits, &Hit{r: r, status: status, persist: 3})
}
