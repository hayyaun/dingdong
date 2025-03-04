package game

import (
	"unicode"

	"github.com/gdamore/tcell"
	"github.com/hayyaun/dingdong/utils"
)

type Cell struct {
	r      rune
	status Status
}

var lines = [][]*Cell{}

var iter = 0 // used for adding gap between rows

func updateCells() {
	iter += 1
	ignore := iter%3 != 0 // ignore some rows

	// Chance of 0 - 1 - 2 - 3 - 4 cells
	n := utils.WeightedRandom(chanceToWeightable(chances)).(int)
	if len(lines) < height || ignore {
		n = 0
	}

	// Fill a line with keys
	line := []*Cell{}
	for i := 0; i < n; i += 1 {
		r := utils.WeightedRandom(keyToWeightable(keys)).(rune)
		line = append(line, &Cell{r: r, status: None})
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
		// Show Bars
		if i < 2 {
			screen.SetContent(0, row, '+', nil, styleGood) // Good
		} else if i < 5 {
			screen.SetContent(0, row, '+', nil, styleMeh) // Meh
		} else {
			screen.SetContent(0, row, '-', nil, styleBad) // Bad
		}
		// Show Particles
		for _, cell := range line {
			i := mapRuneToIndex(cell.r)
			col := i*spfx + padx
			r := unicode.ToUpper(cell.r)
			st := cell.status.toStyle()
			screen.SetContent(col, row, r, nil, st) // Draw cell
		}
	}
}
