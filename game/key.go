package game

import (
	"unicode"

	"github.com/gdamore/tcell"
)

func showKeys(screen tcell.Screen) {
	end := (height + 1) * spfy
	for i, c := range keys {
		col := i*spfx + padx
		screen.SetContent(col, end, unicode.ToUpper(c), nil, style)
	}
}
