package main

import (
	"log"
	"meli-challenge/api"
	"net/http"
)

func main() {
	mux := api.NewRouter()

	port := ":8080"
	log.Println("Starting server on port " + port)
	http.ListenAndServe(port, mux)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}
