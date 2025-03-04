package game

import (
	"fmt"
	"slices"

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

func mapRuneToIndex(char rune) int {
	for i, c := range keys {
		if c == char {
			return i
		}
	}
	return -1 // Return -1 if character is not in slice
}

func checkScore(ev *tcell.EventKey, screen tcell.Screen) {

	r := ev.Rune()

	for i, line := range lines {
		if !slices.Contains(line, mapRuneToIndex(r)) {
			continue
		}
		if i < 2 {
			score += GOOD_SCORE
			hit(r, Good)
		} else if i < 5 {
			score += MEH_SCORE
			hit(r, Meh)
		}
		combo += 1
		return // don't check further - key is used already
	}

	// Wrong Key - Wrong Time
	score = max(0, score+BAD_SCORE)
	combo = 0
	hit(r, Bad)
}
