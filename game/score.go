package game

import (
	"fmt"

	"github.com/gdamore/tcell"
)

func drawText(screen tcell.Screen, x, y int, text string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)

	for i, c := range text {
		screen.SetContent(x+i, y, rune(c), nil, style) // Convert byte to rune
	}
	screen.Show()
}

func showScore(screen tcell.Screen) {
	drawText(screen, padx, height+4, fmt.Sprintf("Score: %v \t\t\t Combo: %v", score, combo))
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
	drawText(screen, padx, height+5, fmt.Sprintf("Key: %c", ev.Rune())) // Debug
	for i, line := range lines {
		addScore := 0
		if i < 2 {
			addScore = 5
		} else if i < 5 {
			addScore = 2
		}
		if contains(line, mapRuneToIndex(ev.Rune())) {
			score += addScore
			combo += 1
			return // don't check further - key is used already
		}
	}

	// Wrong Key - Wrong Time
	score = max(0, score-1)
	combo = 0
}
