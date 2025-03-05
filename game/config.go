package game

const (
	height = 20 // game box height
	padx   = 6  // padding
	spfx   = 6  // space factor horizontal
	spfy   = 1  // space factor vertical

	MODE_HARD   = "hard"
	MODE_NORMAL = "normal"
	MODE_EASY   = "easy"
)

var fps uint = 10 // frames per second
var Modes = []string{MODE_EASY, MODE_NORMAL, MODE_HARD}
var Mode = MODE_NORMAL

func configMode() {
	switch Mode {
	case MODE_EASY:
		fps = 6
	case MODE_HARD:
		fps = 12
	}
}
