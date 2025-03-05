package game

import (
	"slices"

	"github.com/gdamore/tcell"
)

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

func (status Status) toStyle() tcell.Style {
	if status == Good {
		return styleGood
	} else if status == Meh {
		return styleMeh
	} else if status == Bad {
		return styleBad
	}
	return style
}

func updateHits() {
	for i, hit := range hits {
		if hit == nil {
			continue
		}
		if hit.Persist <= 0 {
			hits = slices.Delete(hits, i, i+1)
		}
		hit.Persist -= 1
	}
}

func hit(r rune, s Status) {
	hits = append(hits, &Hit{Rune: r, Status: s, Persist: 3})
}
