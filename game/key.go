package game

import (
	"unicode"

	"github.com/gdamore/tcell"
	. "github.com/hayyaun/dingdong/utils"
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

func keyToWeightable(keys []WKey) []Weighable {
	weightables := make([]Weighable, len(keys))
	for i := range keys {
		weightables[i] = &keys[i]
	}
	return weightables
}

func mapRuneToIndex(char rune) int {
	for i, k := range keys {
		if k.r == char {
			return i
		}
	}
	return -1 // Return -1 if character is not in slice
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
