package service

import (
	"errors"
	"log"
	"meli-challenge/internal"
	"meli-challenge/internal/model"
	"meli-challenge/internal/validator"

	"strconv"
)

type MelodyService struct{}

func NewMelodyService() *MelodyService {
	return &MelodyService{}
}

// Validate a melody in text format and returns its representation.
//
// Parameters:
//
//	input (string): The melody in text format to validate.
//
// Returns:
//
//	*model.Melody: The validated melody.
//	error: An error if the melody is not valid or cannot be parsed.
func (s *MelodyService) Validate(input string) (*model.Melody, error) {

	log.Println("Starting melody validation")

	isValid, errPos := validator.ValidateMelody(input)
	if !isValid {
		log.Printf("Melody is invalid, error at position: %d", errPos)

		return nil, errors.New("error at position " + strconv.Itoa(errPos))
	}
	log.Println("Melody is valid, proceeding to parse")

	melody, err := internal.ParseMelody(input)

	if err != nil {
		log.Printf("Error parsing melody: %v", err)
		return nil, err
	}

	log.Println("Melody parsed successfully")

	return melody, nil
}
