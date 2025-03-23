package mapper

import (
	"meli-challenge/api/request"
	"meli-challenge/internal/model"
)

func PlayMelodyRequestToMelody(rMelodyRequest request.PlayMelodyRequest) model.Melody {
	notes := make([]model.Note, len(rMelodyRequest.Notes))
	for i, rNote := range rMelodyRequest.Notes {
		notes[i] = newNote(rNote)
	}

	return model.Melody{
		Tempo: newTempo(rMelodyRequest.Tempo),
		Notes: notes,
	}
}

func newNote(rNote request.Note) model.Note {
	return model.Note{
		Type:       rNote.Type,
		Name:       rNote.Name,
		Octave:     rNote.Octave,
		Alteration: rNote.Alteration,
		Duration:   rNote.Duration,
		Frequency:  model.Number(rNote.Frequency),
	}
}

func newTempo(rTempo request.Tempo) model.Tempo {
	return model.Tempo{
		Value: rTempo.Value,
		Unit:  rTempo.Unit,
	}
}
