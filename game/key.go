package game

import (
	"unicode"

	"github.com/gdamore/tcell"
)

func hitLookup(r rune) Status {
	for _, hit := range hits {
		if hit.Rune != r {
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
		if status == Good {
			screen.SetContent(col, end, unicode.ToUpper(r), nil, styleGood)
		} else if status == Meh {
			screen.SetContent(col, end, unicode.ToUpper(r), nil, styleMeh)
		} else if status == Bad {
			screen.SetContent(col, end, unicode.ToUpper(r), nil, styleBad)
		} else {
			screen.SetContent(col, end, unicode.ToUpper(r), nil, style)
		}
	}
}
