package game

import "github.com/gdamore/tcell"

func showKeys(screen tcell.Screen) {
	end := (height + 1) * spfy
	for i, c := range keys {
		col := i*spfx + padx
		screen.SetContent(col, end, rune(c), nil, style)
	}
}
