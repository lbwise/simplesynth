package main

import "fmt"

const MAX_AMP = 32760
const SAMPLE_RATE = 44100

func main() {
	stave := Stave{
		StaveNote{Note: "C", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "D", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "E", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "F", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "G", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "A", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "B", Octave: 3, Interval: 1000, Velocity: 100},
		StaveNote{Note: "C", Octave: 4, Interval: 1000, Velocity: 100},
	}

	// Prepares the stave to be played
	sampleSize := stave.Generate()
	fmt.Println(stave)

	buf := make([]int16, sampleSize)
	var writeIdx int
	for _, note := range stave {
		osc := SquareOscillator{note}
		n, err := osc.Generate(buf, writeIdx)
		writeIdx += n
		if err != nil {
			panic("COULD NOT GENERATE SOUND")
		}
	}

	CreateWAV("first sound", buf, sampleSize)
}
