package api

import (
	"net/http"
)

// NewRouter creates and returns an *http.ServeMux with predefined routes
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /melody/validate", ValidateMelodyHandler)

	return mux
}
