package game

import (
	"github.com/gdamore/tcell"
)

type Cell struct {
	Col    int
	Status Status
}

var lines = [][]*Cell{}

var iter = 0 // used for adding gap between rows

func updateCells() {
	iter += 1
	ignore := iter%2 == 0 // ignore even rows

	// Chance of 1 - 2 - 3 - 4 - 0 elements
	n := weightedRandom(chanceToWeightable(chances)).(int)
	if len(lines) < height || ignore {
		n = 0
	}

	// Fill a line with keys
	line := []*Cell{}
	for i := 0; i < n; i += 1 {
		r := weightedRandom(keyToWeightable(keys)).(rune)
		col := mapRuneToIndex(r)
		line = append(line, &Cell{Col: col, Status: None})
	}

	lines = append(lines, line)

	// Remove first - shift
	if len(lines) > height {
		loseScore(lines[0])
		lines = lines[1:]
	}
}

func showCells(screen tcell.Screen) {
	for i, line := range lines {
		row := (height - i) * spfy // Reverse screen
		if i < 2 {
			screen.SetContent(0, row, '+', nil, styleGood) // Good
		} else if i < 5 {
			screen.SetContent(0, row, '+', nil, styleMeh) // Meh
		} else {
			screen.SetContent(0, row, '-', nil, styleBad) // Bad
		}
		for _, cell := range line {
			col := cell.Col*spfx + padx
			screen.SetContent(col, row, 'X', nil, style) // Draw cell
		}
	}
}
