package game

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

var iter = 0

func updateCells() {
	iter += 1

	n := weightedRandom(10, 100)
	if len(lines) < height || iter%2 == 0 {
		n = 0
	}

	line := []int{}
	for i := 0; i < n; i += 1 {
		col := rand.Intn(width)
		line = append(line, col)
	}

	lines = append(lines, line)

	// Remove first - shift
	if len(lines) > height {
		lines = lines[1:]
	}
}

func showCells(screen tcell.Screen) {
	for i, line := range lines {
		row := (height - i) * spfy // Reverse screen
		if i < 2 {
			screen.SetContent(0, row, '-', nil, styleGood) // Good
		} else if i < 5 {
			screen.SetContent(0, row, '-', nil, styleMeh) // Meh
		} else {
			screen.SetContent(0, row, '-', nil, styleBad) // Bad
		}
		for _, j := range line {
			col := j*spfx + padx
			screen.SetContent(col, row, 'X', nil, style) // Draw cell
		}
	}
}
