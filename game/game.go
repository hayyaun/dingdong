package game

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
)

func Play() {
	configMode()

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}

	// Ensure the terminal gets restored on exit
	defer screen.Fini()

	ticker := time.NewTicker(time.Second / time.Duration(fps))
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
	go playMusic()

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
			showScore(screen)
			screen.Show()

		case <-quit:
			return // Exit when the quit signal is received
		}
	}
}
