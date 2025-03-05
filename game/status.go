package game

import "github.com/gdamore/tcell"

type Status int

const (
	None Status = iota
	Good
	Meh
	Bad
)

func (status Status) toStyle() tcell.Style {
	switch status {
	case Good:
		return styleGood
	case Meh:
		return styleMeh
	case Bad:
		return styleBad
	default:
		return style
	}
}
