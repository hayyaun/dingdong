package game

import "github.com/gdamore/tcell"

var style = tcell.StyleDefault.Foreground(tcell.ColorWhite)
var styleBad = tcell.StyleDefault.Foreground(tcell.ColorOrangeRed)
var styleMeh = tcell.StyleDefault.Foreground(tcell.ColorYellow)
var styleGood = tcell.StyleDefault.Foreground(tcell.ColorGreen)

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
