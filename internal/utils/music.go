package utils

import (
	"math"
	"strings"
)

type NoteAttributes struct {
	Duration   float64
	Octave     int
	Alteration string
}

var defaultNoteAttributes = NoteAttributes{
	Duration:   1,
	Octave:     4,
	Alteration: "n",
}

var noteMapping = map[string]string{
	"A": "la",
	"B": "si",
	"C": "do",
	"D": "re",
	"E": "mi",
	"F": "fa",
	"G": "sol",
}

var notePositions = map[string]int{
	"do":  0,
	"re":  2,
	"mi":  4,
	"fa":  5,
	"sol": 7,
	"la":  9,
	"si":  11,
}

func CalculateNumeration(note string, octave int, alteration string) int {

	basePos := notePositions[strings.ToLower(note)]

	var alt int
	switch alteration {
	case "b":
		alt = -1
	case "#":
		alt = 1
	default:
		alt = 0
	}

	return basePos + alt + (12 * octave)
}

func CalculateFrequency(noteNum int) float64 {
	frequency := 440 * math.Pow(2, float64(noteNum-57)/12)
	return frequency
}
