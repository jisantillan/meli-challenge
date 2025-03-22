package validator

import (
	"strconv"
	"strings"
)

func ValidateDuration(duration string) bool {

	durationRegex := `^(\d+/\d+|\d+)$` // checks for a fraction (n/n) or a number

	if !matchesPattern(duration, durationRegex) {
		return false
	}

	if strings.Contains(duration, "/") {
		if !ValidateFractionalBetween(duration, "/", 0, 4) {
			return false
		}
	} else {
		if !ValidateIntegerBetween(duration, 0, 4) {
			return false
		}
	}

	return true
}

func ValidateAlteration(value string, note string) bool {
	if note == "E" || note == "B" {
		if value != "n" {
			return false
		}
	}

	if note == "C" || note == "F" {
		if value == "b" {
			return false
		}
	}

	if value != "#" && value != "b" && value != "n" {
		return false
	}

	return true
}

func ValidateOctave(value string) bool {
	octave, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if octave < 0 || octave > 8 {
		return false
	}

	return true
}

func ExtractAttributes(part string) (string, error) {
	var attributesStr string
	if strings.HasPrefix(part[1:], "{") && strings.HasSuffix(part, "}") {
		attributesStr = part[2 : len(part)-1]
	} else {
		attributesStr = part[2:]
	}
	return attributesStr, nil
}
