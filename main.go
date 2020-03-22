package main

import (
	"flag"
	"fmt"
	"song"
	"os"
	"playback"
	"time"
)

func main() {
  // Capture the parameters
	songPtr  := flag.String("song", "", "Path of the song to play. (Required)")
	tempoPtr := flag.Int("tempo", 500, "Tempo to use for playback in BPM")
	loopPtr  := flag.Int("loop", 3, "Number of times to loop the song pattern. (1 => single run)")

	flag.Parse()

	// quit if no song has been specified
	if *songPtr == "" {
	    flag.PrintDefaults()
	    os.Exit(1)
	}

  // Load the song
  s := song.LoadSongFromFile(*songPtr)

  // Display the playback details
	fmt.Printf("%s, Using tempo: %d\n", s.Details(), *tempoPtr)

	// Initialize audio interface once before playback of any samples
  playback.InitializeAudio()
	defer playback.ShutdownAudio()

  // Play the song
  s.Play(*tempoPtr, *loopPtr)

  fmt.Println("Done!")
	time.Sleep(300 * time.Millisecond)
}
