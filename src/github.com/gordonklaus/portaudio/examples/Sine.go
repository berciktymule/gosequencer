package main

import (
	"github.com/gordonklaus/portaudio"
	"math"
	"time"
	"fmt"
)

var sampleRate float64
var frequency = 80.0
var step float64

func getNextPhase(phase float64) float64 {
	_, phase = math.Modf(phase + step)
	return float64(math.Sin(2 * math.Pi * phase))
}

func nextValue(phase int) float32 {
	return float32(math.Sin(math.Pi * float64(phase) * float64(step)))
}

func main() {
	sampleRate = 22050.0
	step = frequency / sampleRate
	fmt.Println(step)
	var phase float64 = 0
	portaudio.Initialize()
	defer portaudio.Terminate()
	h, err := portaudio.DefaultHostApi()
	chk(err)
	stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), func(out []int32) {
		for i := range out {
			phase = getNextPhase(phase)
			out[i] = int32(phase * 18446744073709551615)
		}
	})
	chk(err)
	defer stream.Close()
	chk(stream.Start())
	time.Sleep(1*time.Second)
	chk(stream.Stop())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
