package track

import (
  "fmt"
  "log"
)

// Stores the individual instrument details that determine
// when to sound a particular sample and which sample to use
type Track struct {
  Name            string
  PathToSample    string
  Pattern         []bool
  Length          int
}

// Interprets the Track data
func LoadTrack(name string, pattern string, pathToSample string) Track {

  // Verify that the pattern to play is at least 1 beat
  if len(pattern) < 1 {
    var errorMessage = fmt.Sprintf("Malformed pattern in track %s detected", name)
    log.Fatal(errorMessage)
  }

  var t = Track{Name: name, PathToSample: pathToSample, Length:len(pattern)}

  // Interpret the input pattern to markers when to sound the sample
  var ticks = make([]bool, t.Length)
  for pos, char := range pattern {
    if char == '+' {
      ticks[pos] = true
    }
  }

  t.Pattern = ticks
  return t
}

// Determines if current track should be activated for current tick
func (t Track) ShouldSountAt(tickIndex int) (bool) {
  return t.Pattern[tickIndex]
}
