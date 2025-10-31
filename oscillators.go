package main

import "math"

type Oscillator interface {
	Generate(int16, int) (int, error)
}

type SquareOscillator struct {
	Note StaveNote
}

func (s *SquareOscillator) Generate(buf []int16, writeIdx int) (int, error) {

	numSamples := s.Note.Interval * SAMPLE_RATE / 1000
	amp := s.Note.amplitude

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(SAMPLE_RATE)
		y := float64(amp) * math.Sin(2*math.Pi*s.Note.frequency*t)
		if y >= 0 {
			buf[writeIdx+i] = amp
		} else {
			buf[writeIdx+i] = -amp
		}
	}
	return numSamples, nil
}

type SinOscillator struct {
	Note StaveNote
}

func (s *SinOscillator) Generate(buf []int16, writeIdx int) (int, error) {
	numSamples := s.Note.Interval * SAMPLE_RATE / 1000
	amp := s.Note.amplitude

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(SAMPLE_RATE)
		y := float64(amp) * math.Sin(2*math.Pi*s.Note.frequency*t)
		buf[writeIdx+i] = int16(y)
	}
	return numSamples, nil
}

type TriangleOscillator struct {
}

func (s *TriangleOscillator) generate(buf int16) error {
	return nil
}
