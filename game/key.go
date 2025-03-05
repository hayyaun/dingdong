package game

import (
	"unicode"

	"github.com/gdamore/tcell"
)

type WKey struct {
	v rune
	w int // Weight
}

var keys = []WKey{
	{v: 'a', w: 1}, {v: 's', w: 1}, {v: 'd', w: 2}, {v: 'f', w: 5}, // left hand
	{v: 'j', w: 5}, {v: 'k', w: 2}, {v: 'l', w: 1}, {v: ';', w: 1}, // right hand
}

func (key *WKey) Weight() int {
	return key.w
}

func (key *WKey) Value() any {
	return key.v
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
		if hit == nil || hit.Rune != r {
			continue
		}
		return hit.Status
	}
	return None
}

func showKeys(screen tcell.Screen) {
	end := (height + 1) * spfy
	for i, k := range keys {
		col := i*spfx + padx
		status := hitLookup(k.v)
		st := status.toStyle()
		screen.SetContent(col, end, unicode.ToUpper(k.v), nil, st)
	}
}
