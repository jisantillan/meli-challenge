package request

type PlayMelodyRequest struct {
	Tempo Tempo  `json:"tempo"`
	Notes []Note `json:"notes"`
}

type Tempo struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Note struct {
	Type       string  `json:"type"`
	Name       string  `json:"name,omitempty"`
	Octave     int     `json:"octave,omitempty"`
	Alteration string  `json:"alteration,omitempty"`
	Duration   float64 `json:"duration"`
	Frequency  float64 `json:"frequency,omitempty"`
}
