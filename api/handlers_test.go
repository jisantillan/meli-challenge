package api

import (
	"bytes"
	"encoding/json"
	"meli-challenge/api/request"
	"meli-challenge/api/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateMelodyHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    request.ValidateMelodyRequest
		expectedCode   int
		expectedResult *response.ValidateMelodyResponse
		expectedError  *response.ValidateMelodyErrorResponse
	}{
		{
			name: "Valid Melody",
			requestBody: request.ValidateMelodyRequest{
				Melody: "60 A{d=7/4;o=3;a=#} B{o=2;d=1/4} S",
			},
			expectedCode: http.StatusOK,
			expectedResult: &response.ValidateMelodyResponse{
				Tempo: response.Tempo{
					Value: 60,
					Unit:  "bpm",
				},
				Notes: []response.Note{
					{
						Type:       "note",
						Name:       "la",
						Octave:     3,
						Alteration: "#",
						Duration:   1.75,
						Frequency:  233.08,
					},
					{
						Type:       "note",
						Name:       "si",
						Octave:     2,
						Alteration: "n",
						Duration:   0.25,
						Frequency:  123.47,
					},
					{
						Type:     "silence",
						Duration: 1,
					},
				},
			},
		},
		{
			name: "Invalid Melody",
			requestBody: request.ValidateMelodyRequest{
				Melody: "60 A{d=7/4;o=3;a=# B{o=2;d=1/4} S",
			},
			expectedCode: http.StatusBadRequest,
			expectedError: &response.ValidateMelodyErrorResponse{
				Cause: "error at position 18",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Fatalf("Error marshaling request: %v", err)
			}

			req := httptest.NewRequest("POST", "/melody/validate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			ValidateMelodyHandler(rr, req)

			res := rr.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, res.StatusCode)
			}

			if tt.expectedResult != nil {
				var responseBody response.ValidateMelodyResponse
				json.NewDecoder(res.Body).Decode(&responseBody)

				if !compareMelodyResponses(&responseBody, tt.expectedResult) {
					t.Errorf("Expected response %+v, got %+v", tt.expectedResult, responseBody)
				}
			}

			if tt.expectedError != nil {
				var errorResponse response.ValidateMelodyErrorResponse
				json.NewDecoder(res.Body).Decode(&errorResponse)

				if errorResponse.Cause != tt.expectedError.Cause {
					t.Errorf("Expected error '%s', got '%s'", tt.expectedError.Cause, errorResponse.Cause)
				}
			}
		})
	}
}

func compareMelodyResponses(current, expected *response.ValidateMelodyResponse) bool {
	if current.Tempo != expected.Tempo || len(current.Notes) != len(expected.Notes) {
		return false
	}

	for i := range current.Notes {
		if current.Notes[i].Type != expected.Notes[i].Type ||
			current.Notes[i].Name != expected.Notes[i].Name ||
			current.Notes[i].Octave != expected.Notes[i].Octave ||
			current.Notes[i].Alteration != expected.Notes[i].Alteration ||
			current.Notes[i].Duration != expected.Notes[i].Duration ||
			current.Notes[i].Frequency != expected.Notes[i].Frequency {
			return false
		}
	}

	return true
}

func TestPlayMelodyHandler(t *testing.T) {
	requestBody := request.PlayMelodyRequest{
		Tempo: request.Tempo{
			Value: 60,
			Unit:  "bpm",
		},
		Notes: []request.Note{
			{
				Type:       "note",
				Name:       "la",
				Octave:     3,
				Alteration: "#",
				Duration:   1.75,
				Frequency:  233.08,
			},
			{
				Type:       "note",
				Name:       "si",
				Octave:     2,
				Alteration: "n",
				Duration:   0.25,
				Frequency:  123.47,
			},
			{
				Type:     "silence",
				Duration: 1,
			},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshaling request: %v", err)
	}

	req := httptest.NewRequest("POST", "/melody/play", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	PlayMelodyHandler(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		t.Errorf("Expected status %d, got %d", http.StatusAccepted, res.StatusCode)
	}
}
