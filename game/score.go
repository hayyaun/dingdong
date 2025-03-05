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
	col := mapRuneToIndex(r)

	for i, line := range lines {
		for _, cell := range line {
			if cell.Col != col || cell.Status != None {
				continue // no score - 404 or already used
			}
			if i < 2 {
				cell.Status = Good
				score += GOOD_SCORE
				hit(r, Good)
			} else if i < 5 {
				cell.Status = Bad
				score += MEH_SCORE
				hit(r, Meh)
			}
			combo += 1
			return // don't check further - key is used already
		}
	}

	// Wrong Key - Wrong Time
	score = max(0, score+BAD_SCORE)
	combo = 0
	hit(r, Bad)
}

func loseScore(line []*Cell) {
	// TODO - hard-mode only
	for _, cell := range line {
		if cell.Status != None {
			continue
		}
		// Miss
		score = max(0, score+BAD_SCORE)
		combo = 0
	}
}
