package track

import (
  "testing"
  "fmt"
)

var testPath = "test_path"
var testName = "test_name"
var testPattern = "-----+---"
var testTrack = LoadTrack(testName, testPattern, testPath)

func ExpectToSountAt(testTrack Track, index int, shouldI bool) string {
  var shouldSountAt = testTrack.ShouldSountAt(index)
  if shouldSountAt  != shouldI {
    return fmt.Sprintf("ShouldSountAt incorrectly said %v, for offset: %d.", shouldSountAt, index)
  }
  return ""
}

func TestCorrectTiming(t *testing.T) {
  var encounteredError = ExpectToSountAt(testTrack, 2, false)

  if encounteredError != "" {
    t.Errorf(encounteredError)
  }

  encounteredError = ExpectToSountAt(testTrack, 5, true)

  if encounteredError != "" {
    t.Errorf(encounteredError)
  }
}

func TestCorrectName(t *testing.T) {
  if testTrack.Name != testName {
    t.Errorf("Name was incorrect, got: %s, want: %s.", testName, testTrack.Name)
  }
}

func TestCorrectPath(t *testing.T) {
  if testTrack.PathToSample != testPath {
    t.Errorf("PathToSample was incorrect, got: %s, want: %s.", testPath, testTrack.PathToSample)
  }
}
