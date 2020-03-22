# Drum Machine

## The Objective
The goal is to create a simple [drum machine sequencer](http://en.wikipedia.org/wiki/Drum_machine) that allows you to "play" your own version of the famous [four-on-the-floor](http://en.wikipedia.org/wiki/Four_on_the_floor_(music)) pattern, as shown below:

![https://raw.githubusercontent.com/mattetti/sm-808/master/Four_to_the_floor_Roland_TR-707.jpg](https://raw.githubusercontent.com/mattetti/sm-808/master/Four_to_the_floor_Roland_TR-707.jpg)


## Installation
This solution uses [portaudio](http://www.portaudio.com/) for playback.
more specifically [gordonklaus golang port](https://github.com/gordonklaus/portaudio)

Please follow their installation steps if necessary.

The Linux 64bit library is already part of this repo as I have used Ubuntu.

## Usage
Before running please make sure that the project is in Go path:
`export GOPATH=$(pwd)`
### Running
`go run main.go -song examples/four_on_floor.trk` should render Four on Floor playback
#### Parameters
`-song` Path to the song to be played. ***Required***

`-tempo` Playback speed (BPM). Defaults to 500BPM

`-loop` Number of times to loop through the song pattern. Defaults to 3

### Song of the year
Make sure to run
`go run main.go -song examples/blue_monday.trk -loop 1`


## Samples
Current samples contain:
1. [Kick drum](samples/kick.aiff)
1. [Snare drum](samples/snare.aiff)
1. [Hi-Hat](samples/hihat.aiff)

### Adding samples

Currently the code supports only one specific format for playback.

If you wish to add additional samples please run:

`go run src/github.com/gordonklaus/portaudio/examples/record.go samples/new_sample_name`

## Song Format
I've chosen to create my own simple format to easily edit the track.
It seems easy to understand and has the ability to visually align all the tracks.
```
HiHat   --+---+---+---+-    samples/hihat.aiff
Snare   ----+-------+---    samples/snare.aiff
Kick    +---+---+---+---    samples/kick.aiff
```
First column is the track label used for display during playback (keep it short;)
Following is the pattern for the current track.

`+` means sound, `-` means no sound.

The last column has the path to the samples to use (relative to pwd).

The columns are separated by any number of white spaces so that it is easy to align the track patterns for tracks with labels of different lengths.
