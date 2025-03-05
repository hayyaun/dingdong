package game

import (
	"fmt"

	"github.com/gdamore/tcell"
)

const (
	GOOD_SCORE = 5
	MEH_SCORE  = 2
	BAD_SCORE  = -1
)

func drawText(screen tcell.Screen, x, y int, text string, color *tcell.Color) {
	c := tcell.ColorGreen
	if color != nil {
		c = *color
	}
	style := tcell.StyleDefault.Foreground(c)

	for i, c := range text {
		screen.SetContent(x+i, y, rune(c), nil, style) // Convert byte to rune
	}
	screen.Show()
}

func showScore(screen tcell.Screen) {
	drawText(screen, padx, height+4, fmt.Sprintf("Score: %v \t\t\t Combo: %v", score, combo), nil)
}

func checkScore(ev *tcell.EventKey) {

	r := ev.Rune()

	for i, line := range lines {
		for _, cell := range line {
			if cell.r != r || cell.status != None {
				continue // no score - 404 or already used
			}
			if i < 2 {
				// Good Time
				cell.status = Good
				score += GOOD_SCORE
				combo += 1
				hit(r, Good)
			} else if i < 5 {
				// Meh Time
				cell.status = Meh
				score += MEH_SCORE
				combo += 1
				hit(r, Meh)
			} else {
				// Wrong Time
				cell.status = Bad
				score = max(0, score+BAD_SCORE)
				combo = 0
				hit(r, Bad)
			}
			return // don't check further - key is used already
		}
	}

	// Wrong Key
	score = max(0, score+BAD_SCORE)
	combo = 0
	hit(r, Bad)
}

// Checks outgoing cells
// TODO - hard-mode only
func loseScore(line []*Cell) {
	for _, cell := range line {
		if cell.status != None {
			hit(cell.r, cell.status)
			continue
		}
		// Miss
		score = max(0, score+BAD_SCORE)
		combo = 0
		hit(cell.r, Bad)
	}
}
