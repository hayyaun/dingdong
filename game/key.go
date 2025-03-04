package game

import (
	"unicode"

	"github.com/gdamore/tcell"
)

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
	for i, r := range keys {
		col := i*spfx + padx
		status := hitLookup(r)
		st := style
		if status == Good {
			st = styleGood
		} else if status == Meh {
			st = styleMeh
		} else if status == Bad {
			st = styleBad
		}
		screen.SetContent(col, end, unicode.ToUpper(r), nil, st)
	}
}
