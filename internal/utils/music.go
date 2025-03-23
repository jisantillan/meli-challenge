package utils

import (
	"math"
	"strings"
)

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
