package game

import (
	"unicode"

	"github.com/gdamore/tcell"
)

type WKey struct {
	r rune
	w int // Weight
}

var keys = []WKey{
	{r: 'a', w: 1}, {r: 's', w: 1}, {r: 'd', w: 2}, {r: 'f', w: 5}, // left hand
	{r: 'j', w: 5}, {r: 'k', w: 2}, {r: 'l', w: 1}, {r: ';', w: 1}, // right hand
}

func (key *WKey) Weight() int {
	return key.w
}

func (key *WKey) Value() any {
	return key.r
}

func keyToWeightable(keys []WKey) []Weightable {
	weightables := make([]Weightable, len(keys))
	for i := range keys {
		weightables[i] = &keys[i]
	}
	return weightables
}

func hitLookup(r rune) Status {
	for _, hit := range hits {
		if hit == nil || hit.r != r {
			continue
		}
		return hit.status
	}
	return None
}

func showKeys(screen tcell.Screen) {
	end := (height + 1) * spfy
	for i, k := range keys {
		col := i*spfx + padx
		status := hitLookup(k.r)
		st := status.toStyle()
		screen.SetContent(col, end, unicode.ToUpper(k.r), nil, st)
	}
}
