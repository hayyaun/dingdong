package game

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
)

const (
	height = 20 // game box height
	padx   = 6  // padding
	spfx   = 6  // space factor horizontal
	spfy   = 1  // space factor vertical
	fps    = 10 // frames per second
)

func Play() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalln("Error initializing screen: %", err)
	}

	// Ensure the terminal gets restored on exit
	defer screen.Fini()

	ticker := time.NewTicker(time.Second / fps)
	defer ticker.Stop()

	// Channel for key events
	quit := make(chan struct{})

	// Goroutine for handling input events
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Rune() == 'q' || ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quit) // Signal the main loop to exit
					return
				}
				checkScore(ev)
			}
		}
	}()

	// Music
	go PlayMusic()

	// Game loop
	for {
		select {

		case <-ticker.C:
			// Redraw the screen periodically
			screen.Clear()
			updateCells() // matrix
			updateHits()  // matrix
			showCells(screen)
			showKeys(screen)
			score.show(screen)
			screen.Show()

		case <-quit:
			return // Exit when the quit signal is received
		}
	}
}
