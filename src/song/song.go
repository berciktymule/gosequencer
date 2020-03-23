package song

import (
  "fmt"
  "time"
  "strings"
  "bufio"
  "os"
  "log"
  "track"
  "path/filepath"
  "playback"
)

type Song struct {
  Name            string
  PathToInput     string
  Length          int
  TrackList       []track.Track
}

// Prints the details of the Song to the console
func (s Song) Details() (string) {
  return fmt.Sprintf("Song: %s, Length: %d", s.Name, s.Length)
}

// Play the song
func (currentSong Song) Play(tempo int, loopCount int ) {
  // Loop the pattern the requested number of times
  for currentLoop := 1; currentLoop <= loopCount; currentLoop++ {
    for tickIndex := 0; tickIndex < currentSong.Length; tickIndex++ {

      var label, paths = currentSong.GetSamplesForTick(tickIndex)

      playCurrentSamples(paths)
      displayTickLabel(label, tickIndex)
      waitTillNextTick(tempo)
    }
  }
}

// Returns concatenated names of all the instruments
// that should be played at current tick
func (s Song) GetSamplesForTick(tickIndex int) (string, []string) {
  var trackNames []string
  var samplePaths []string

  for _, currentTrack := range s.TrackList {
    if currentTrack.ShouldSountAt(tickIndex) {
      trackNames = append(trackNames, currentTrack.Name)
      samplePaths = append(samplePaths, currentTrack.PathToSample)
    }
  }
  if len(trackNames) < 1 {
    return "-", nil
  }
  //Join the name of all the samples
  return strings.Join(trackNames, "+"), samplePaths
}

// Plays all the supplied samples simultaneously
func playCurrentSamples(samplePaths []string) {
  for _, samplePath := range samplePaths {
    go playback.PlayFile(samplePath)
  }
}

// Formats and displays current tick
func displayTickLabel(label string, tickIndex int) {
  // display the labels of the instruments that play at current time
  fmt.Printf("| %-17s", label)

  // format to 4 bars per line
  if (tickIndex % 4) == 3 {
    fmt.Printf("|\n")
  }
}

// Pauses execution until it's time to play the next tick
func waitTillNextTick(tempo int) {
  // Delay between ticks is computed by inverting Beats Per Second
  var tickDurationInMilliseconds = time.Duration(60000.0/tempo)
  // wait until next beat based on the tempo
  time.Sleep(tickDurationInMilliseconds * time.Millisecond)
}

// Read the file from the disk and interpret the song
func LoadSongFromFile(pathToInput string) Song {
  // make sure that the path exists
  file, err := os.Open(pathToInput)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  // define the new song
  newSong := Song{Name:filepath.Base(pathToInput), PathToInput:pathToInput, Length: -1}

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var rawLine = scanner.Text()
    //split the line into fields
    var line  = strings.Fields(rawLine)
    //Make sure we have all the data presend in the track
    if len(line) < 3 {
      log.Fatal("Malformed track detected: " + rawLine)
    }

    // interpret the parts of the stored track
    name, pattern, pathToSample := line[0], line[1], line[2]

    // get current track length and validate it matches
    var trackLength = len(pattern)
    if newSong.Length > 0 && newSong.Length != trackLength {
      log.Fatal("Mismatch on track length compared to previous track: " + rawLine)
    }

    var newTrack =  track.LoadTrack(name, pattern, pathToSample)

    // Since all tracks should be equal length
    // we can use the track's length as the whole song's length
    newSong.Length = trackLength

    // Ad the track to the song
    newSong.TrackList = append(newSong.TrackList, newTrack)
  }

  //log any read errors
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return newSong
}
