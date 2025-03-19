package main

import (
	"flag"
	"fmt"
	"meli-challenge/internal/validator"
)

func main() {
	melodyFlag := flag.String("melody", "", "A string representing the melody to validate")
	flag.Parse()

	isValid, err := validator.ValidateMelody(*melodyFlag)
	if !isValid {
		fmt.Printf("error at position %d", err)
	} else {
		fmt.Println("valid melody")
	}
}
