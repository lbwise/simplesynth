package main

import "math"

var flatsMap = map[string]string{
	"A#": "A#",
	"Bb": "A#",
	"C#": "C#",
	"Db": "C#",
	"D#": "D#",
	"Eb": "D#",
	"F#": "F#",
	"Gb": "F#",
	"G#": "G#",
	"Ab": "G#",
}

type StaveNote struct {
	Note      string
	Octave    int
	Interval  int // ms
	Velocity  int
	frequency float64
	amplitude int16
}

type Stave []StaveNote

func (s *Stave) Generate() int {
	var sampleSize int
	for i := range *s {
		note := &(*s)[i] // reference to the actual element
		sampleSize += note.Interval * SAMPLE_RATE / 1000
		if note.Velocity == 0 {
			note.Velocity = 127
		}
		note.amplitude = int16(MAX_AMP * note.Velocity / 128)
		note.frequency = GetFrequency(note.Note, note.Octave)
	}
	return sampleSize
}

func GetFrequency(note string, octave int) float64 {
	const F0 = 440
	const OCTAVE0 = 4
	aIdx := 9

	if len(note) == 2 {
		note = flatsMap[note]
	}

	notes := [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	var localDiff int
	for i, n := range notes {
		if note == n {
			localDiff = i - aIdx
		}
	}
	semitoneDiff := ((octave - OCTAVE0) * 12) + localDiff
	return F0 * math.Pow(2, float64(semitoneDiff)/float64(12))
}
