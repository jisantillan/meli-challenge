package utils

import (
	"fmt"
	"meli-challenge/internal/model"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

func PlayMelody(tempo model.Tempo, notes []model.Note) error {
	const sampleRate = beep.SampleRate(48000)

	err := speaker.Init(sampleRate, 4800)
	if err != nil {
		return fmt.Errorf("error inicializando speaker: %w", err)
	}

	ch := make(chan struct{})

	var sounds []beep.Streamer
	for _, note := range notes {
		freq := note.Frequency

		if note.Type == "silence" {
			durationInSeconds := (float64(note.Duration) * 60) / float64(tempo.Value)               //  note duration to time in seconds.
			sampleDuration := sampleRate.N(time.Duration(durationInSeconds * float64(time.Second))) //  note duration in seconds to the corresponding number of audio samples.

			sounds = append(sounds, beep.Silence(sampleDuration))
			continue
		}
		sineWave, err := generators.SinTone(sampleRate, int(freq))
		if err != nil {
			return fmt.Errorf("error generando tono: %w", err)
		}

		durationInSeconds := (float64(note.Duration) * 60) / float64(tempo.Value)
		sampleDuration := sampleRate.N(time.Duration(durationInSeconds * float64(time.Second)))
		sounds = append(sounds,
			beep.Take(sampleDuration, sineWave),
		)
	}

	sounds = append(sounds, beep.Callback(func() {
		ch <- struct{}{}
	}),
	)

	speaker.Play(beep.Seq(sounds...))

	<-ch
	return nil
}
