package game

import (
	"fmt"

	"github.com/gdamore/tcell"
)

// dynamics
var score = 0
var combo = 0

const (
	GOOD_SCORE = 5
	MEH_SCORE  = 2
	BAD_SCORE  = -1
)

func showScore(screen tcell.Screen) {
	text := fmt.Sprintf("Score: %v \t\t\t Combo: %v", score, combo)
	drawText(screen, padx, height+4, text, nil)
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
			} else if i < 10 {
				// Wrong Time
				cell.status = Bad
				score = max(0, score+BAD_SCORE)
				combo = 0 // reset combo
				hit(r, Bad)
			}
			// else - forgive
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
