package validator

import (
	"regexp"
	"strings"
)

var allowedNoteAttributes = map[string]bool{
	"d": true, // duration
	"a": true, // alteration
	"o": true, // octave
}

var allowedSilenceAttributes = map[string]bool{
	"d": true, // duration
}

func matchesPattern(value, regex string) bool {
	matched, _ := regexp.MatchString(regex, value)
	return matched
}

func validateCloseFormat(part string, position int) (bool, int) {
	if !strings.HasSuffix(part, "}") {
		return false, position + len(part)
	}
	return true, -1
}

func validateAttributeValues(attributesStr string, allowedAttributes map[string]bool, position int, part string) (bool, int) {
	attributes := strings.Split(strings.TrimSuffix(attributesStr, ";"), ";")
	seenAttributes := make(map[string]bool)

	currPos := 2

	for i, attr := range attributes {
		if attr == "" {
			return false, position + currPos
		}

		attrStr := splitAttributeValue(attr)
		parts := strings.SplitN(attrStr, "=", 2)
		if len(parts) != 2 {
			return false, position + currPos
		}
		key, value := parts[0], parts[1]

		if !allowedAttributes[key] {
			return false, position + currPos
		}

		if seenAttributes[key] {
			return false, position + currPos
		}
		seenAttributes[key] = true

		switch key {
		case "d":
			if valid := ValidateDuration(value); !valid {
				return false, position + currPos + len(key) + 1
			}
		case "o":
			if valid := ValidateOctave(value); !valid {
				return false, position + currPos + len(key) + 1
			}
		case "a":
			note := string(part[0])
			if valid := ValidateAlteration(value, note); !valid {
				return false, position + currPos + len(key) + 1
			}
		}

		if i < len(attributes)-1 && part[currPos+len(attrStr)] != ';' {
			return false, position + currPos + len(attrStr)
		}

		// next position (+1 due `;`)
		currPos += len(attr) + 1
	}

	if strings.HasSuffix(attributesStr, ";") {
		return false, position + len(attributesStr) + 1
	}

	return true, -1
}

func splitAttributeValue(attr string) string {
	idx := strings.Index(attr, "=")
	if idx == -1 {
		return ""
	}

	value := attr[idx+1:]

	if equalsIdx := strings.Index(value, "="); equalsIdx != -1 {
		value = value[:equalsIdx-2]
	}

	return attr[:idx+1] + value
}

func validateOpenFormat(part string, position int) (bool, int) {
	if !strings.HasPrefix(part[1:], "{") {
		return false, position + 1
	}
	return true, -1
}

func validateAttributes(part string, position int, allowedAttributes map[string]bool) (bool, int) {

	if valid, errPos := validateOpenFormat(part, position); !valid {
		return false, errPos
	}

	attributesStr, err := ExtractAttributes(part)
	if err != nil {
		return false, -1
	}

	if valid, errPos := validateAttributeValues(attributesStr, allowedAttributes, position, part); !valid {
		return false, errPos
	}

	if valid, errPos := validateCloseFormat(part, position); !valid {
		return false, errPos
	}

	return true, -1
}

func validateSilence(part string, position int) (bool, int) {
	if !(string(part[0]) == "S") {
		return false, position
	}

	if len(part) == 1 {
		return true, -1
	}

	valid, errPos := validateAttributes(part, position, allowedSilenceAttributes)
	if !valid {
		return false, errPos
	}

	return true, -1
}

func validateNote(part string, position int) (bool, int) {

	noteRegex := `^[A-G]$` // checks for a note from A to G

	if !matchesPattern(string(part[0]), noteRegex) {
		return false, position
	}

	if len(part) == 1 {
		return true, -1
	}

	valid, errPos := validateAttributes(part, position, allowedNoteAttributes)
	if !valid {
		return false, errPos
	}

	return true, -1
}

func validateTempo(tempo string, position int) (bool, int) {
	tempoRegex := `^\d+$` // checks if it consists only of numbers
	if !matchesPattern(tempo, tempoRegex) {
		return false, position
	}
	return true, -1
}

func ValidateMelody(melody string) (bool, int) {
	if len(melody) == 0 {
		return false, 0
	}

	tempo, rest := extractFirstPart(melody)
	if len(rest) == 0 {
		return false, len(tempo) + 1
	}

	if valid, err := validateTempo(tempo, 0); !valid {
		return false, err
	}

	position := len(tempo)
	if rest[0] != ' ' {
		return false, position
	}

	position++
	prevSpace := true

	for len(rest) > 0 {
		part, newRest := extractFirstPart(rest)

		if len(part) == 0 {
			return false, position
		}
		if !prevSpace {
			return false, position - 1
		}

		if errPos := validatePart(part, position); errPos != -1 {
			return false, errPos
		}

		rest = newRest
		position += len(part) + 1 // +1 due space
		prevSpace = (len(rest) > 0 && rest[0] == ' ')
	}

	return true, -1
}

func extractFirstPart(melody string) (string, string) {
	melody = strings.TrimSpace(melody)

	if len(melody) == 0 {
		return "", ""
	}

	firstSpace := strings.Index(melody, " ")
	closingBrace := strings.Index(melody, "}")

	if firstSpace == -1 && closingBrace == -1 {
		return melody, ""
	}

	if firstSpace == -1 {
		firstSpace = len(melody)
	}
	if closingBrace != -1 && (firstSpace == -1 || closingBrace < firstSpace) {
		return melody[:closingBrace+1], melody[closingBrace+1:]
	}

	return melody[:firstSpace], melody[firstSpace:]
}

func validatePart(part string, position int) int {
	if strings.HasPrefix(part, "S") {
		if valid, errPos := validateSilence(part, position); !valid {
			return errPos
		}
	} else {
		if valid, errPos := validateNote(part, position); !valid {
			return errPos
		}
	}
	return -1
}
