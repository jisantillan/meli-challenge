package validator

import (
	"strconv"
	"strings"
)

func ValidateFractionalBetween(value string, delimiter string, min float64, max float64) bool {
	parts := strings.Split(value, delimiter)
	if len(parts) != 2 {
		return false
	}

	numerator, err1 := strconv.Atoi(parts[0])
	denominator, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || denominator == 0 {
		return false
	}

	result := float64(numerator) / float64(denominator)
	return result >= min && result <= max
}

func ValidateIntegerBetween(value string, min int, max int) bool {
	num, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return num >= min && num <= max
}

func ValidateDecimalBetween(value string, min float64, max float64) bool {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	return num >= min && num <= max
}
