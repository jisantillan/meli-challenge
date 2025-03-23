package utils

import (
	"fmt"
	"math"
	"meli-challenge/internal/model"
	"strconv"
	"strings"
)

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

func ParseMelody(input string) (*model.Melody, error) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return nil, fmt.Errorf("empty melody")
	}

	tempoValue, err := parseTempo(fields[0])
	if err != nil {
		return nil, err
	}
	melody := &model.Melody{
		Tempo: model.Tempo{
			Value: tempoValue,
			Unit:  "bpm",
		},
	}

	for _, part := range fields[1:] {
		if strings.HasPrefix(part, "S") {
			note, err := parseSilence(part)
			if err != nil {
				return nil, err
			}
			melody.Notes = append(melody.Notes, note)
		} else {
			note, err := parseNote(part)
			if err != nil {
				return nil, err
			}
			melody.Notes = append(melody.Notes, note)
		}
	}

	return melody, nil
}

func parseTempo(tempoStr string) (int, error) {
	tempo, err := strconv.Atoi(tempoStr)
	if err != nil {
		return 0, fmt.Errorf("invalid tempo value: %s", tempoStr)
	}
	return tempo, nil
}

func parseSilence(part string) (model.Note, error) {
	if len(part) == 1 {
		return model.Note{
			Type:     "silence",
			Duration: defaultNoteAttributes.Duration,
		}, nil
	}

	attributesStr := cleanBracesAttributes(part)
	attributes, err := parseAttributes(attributesStr)
	if err != nil {
		return model.Note{}, err
	}

	return model.Note{
		Type:     "silence",
		Duration: attributes.Duration,
	}, nil
}

func parseNote(part string) (model.Note, error) {
	if len(part) == 1 {
		noteName := noteMapping[part]
		noteNum := CalculateNumeration(noteName, defaultNoteAttributes.Octave, defaultNoteAttributes.Alteration)
		frequency := CalculateFrequency(noteNum)
		frequency = round(frequency, 2)
		return model.Note{
			Type:       "note",
			Name:       noteName,
			Octave:     defaultNoteAttributes.Octave,
			Alteration: defaultNoteAttributes.Alteration,
			Duration:   defaultNoteAttributes.Duration,
			Frequency:  model.Number(frequency),
		}, nil
	}

	attributesStr := cleanBracesAttributes(part)
	attributes, err := parseAttributes(attributesStr)
	if err != nil {
		return model.Note{}, err
	}

	noteName := noteMapping[string(part[0])]
	noteNum := CalculateNumeration(noteName, attributes.Octave, attributes.Alteration)
	frequency := CalculateFrequency(noteNum)
	frequency = round(frequency, 2)

	return model.Note{
		Type:       "note",
		Name:       noteName,
		Octave:     attributes.Octave,
		Alteration: attributes.Alteration,
		Duration:   attributes.Duration,
		Frequency:  model.Number(frequency),
	}, nil
}

func cleanBracesAttributes(part string) string {
	return strings.TrimSuffix(strings.TrimPrefix(part, part[:2]), "}")
}

func parseAttributes(attributesStr string) (NoteAttributes, error) {
	attributes := strings.Split(attributesStr, ";")
	result := NoteAttributes{
		Duration:   defaultNoteAttributes.Duration,
		Octave:     defaultNoteAttributes.Octave,
		Alteration: defaultNoteAttributes.Alteration,
	}

	for _, attr := range attributes {
		parts := strings.SplitN(attr, "=", 2)
		key, value := parts[0], parts[1]

		switch key {
		case "d":
			if duration, err := parseDuration(value); err == nil {
				result.Duration = duration
			} else {
				return result, err
			}
		case "o":
			if octave, err := strconv.Atoi(value); err == nil {
				result.Octave = octave
			} else {
				return result, fmt.Errorf("invalid octave value")
			}
		case "a":
			result.Alteration = value
		}
	}

	return result, nil
}

func parseDuration(value string) (float64, error) {
	if strings.Contains(value, "/") {
		parts := strings.Split(value, "/")
		if len(parts) == 2 {
			numerator, err1 := strconv.Atoi(parts[0])
			denominator, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil || denominator == 0 {
				return -1, fmt.Errorf("invalid duration fraction")
			}
			return math.Round((float64(numerator)/float64(denominator))*100) / 100, nil
		}
	} else {
		if duration, err := strconv.ParseFloat(value, 64); err == nil {
			return duration, nil
		}
		return -1, fmt.Errorf("invalid duration value")
	}
	return -1, fmt.Errorf("invalid duration value")
}

func round(val float64, precision int) float64 {
	rounder := math.Pow(10, float64(precision))
	return math.Round(val*rounder) / rounder
}
