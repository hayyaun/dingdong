package game

import (
	"unicode"

	"github.com/gdamore/tcell"
)

type WKey struct {
	v rune
	w int // Weight
}

type WKeySlice []WKey

var keys = WKeySlice{
	{v: 'a', w: 1}, {v: 's', w: 1}, {v: 'd', w: 2}, {v: 'f', w: 5}, // left hand
	{v: 'j', w: 5}, {v: 'k', w: 2}, {v: 'l', w: 1}, {v: ';', w: 1}, // right hand
}

// Method to convert WKeySlice to []Weightable
func (keys WKeySlice) ToWeightables() []Weightable {
	weightables := make([]Weightable, len(keys))
	for i, k := range keys {
		weightables[i] = Weightable{v: k.v, w: k.w}
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
		st := style
		if status == Good {
			st = styleGood
		} else if status == Meh {
			st = styleMeh
		} else if status == Bad {
			st = styleBad
		}
		screen.SetContent(col, end, unicode.ToUpper(k.v), nil, st)
	}
}
