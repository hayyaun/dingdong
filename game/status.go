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
	if status == Good {
		return styleGood
	} else if status == Meh {
		return styleMeh
	} else if status == Bad {
		return styleBad
	}
	return style
}
