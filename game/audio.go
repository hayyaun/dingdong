package game

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func PlayMusic() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Caught a panic:", r)
		}
	}()

	// Open the audio file
	f, err := os.Open("./assets/music.mp3")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Decode the file
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
	defer streamer.Close()

	// Loop the audio indefinitely
	loopedStreamer := beep.Loop(-1, streamer)

	// Initialize the speaker
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("Error initializing speaker:", err)
		return
	}

	// Play the looped audio
	speaker.Play(loopedStreamer)

	// Keep the program running so the audio keeps playing
	select {} // Infinite blocking
}
