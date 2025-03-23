package api

import (
	"encoding/json"
	"log"
	"meli-challenge/api/mapper"
	"meli-challenge/api/request"
	"meli-challenge/api/response"
	"meli-challenge/api/service"
	"net/http"
)

func ValidateMelodyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to validate melody")

	w.Header().Set("Content-Type", "application/json")

	var melodyReq request.ValidateMelodyRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&melodyReq)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		response := response.ValidateMelodyErrorResponse{
			Cause: "Invalid JSON",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	melodyService := service.NewMelodyService()
	log.Printf("Validating melody: %s", melodyReq.Melody)

	validateMelodyResponse, err := melodyService.Validate(melodyReq.Melody)
	if err != nil {
		log.Printf("Error validating melody: %v", err)
		response := response.ValidateMelodyErrorResponse{
			Cause: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("Melody validated successfully")

	json.NewEncoder(w).Encode(validateMelodyResponse)
}

func PlayMelodyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to play melody")

	var melodyReq request.PlayMelodyRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&melodyReq)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	melody := mapper.PlayMelodyRequestToMelody(melodyReq)
	melodyService := service.NewMelodyService()
	log.Printf("Playing melody")

	err = melodyService.Play(melody.Tempo, melody.Notes)
	if err != nil {
		log.Printf("Error playing melody: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("Melody played successfully")

	w.WriteHeader(http.StatusAccepted)
}
