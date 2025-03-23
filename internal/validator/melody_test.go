package validator

import (
	"testing"
)

func TestValidateMelody_Valid(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if !valid {
		t.Errorf("Expected melody to be valid, but got error at position %d", errPos)
	}
}

func TestValidateMelody_Invalid_EmptyInput(t *testing.T) {
	melody := ""
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to empty input, but it passed")
	} else {
		if errPos != 0 {
			t.Errorf("Expected error at position 0, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_Tempo(t *testing.T) {
	melody := "A{d=7/4;o=3;a=#}B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to missing tempo, but it passed")
	} else {
		if errPos != 0 {
			t.Errorf("Expected error at position 0, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_NoSpaceBetweenNotes(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#}B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to missing space, but it passed")
	} else {
		if errPos != 19 {
			t.Errorf("Expected error at position 19, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidDuration_FormatError(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=?3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid duration 'd=0?3', but it passed")
	} else {
		if errPos != 59 {
			t.Errorf("Expected error at position 59, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidDuration_DecimalOutOfRange(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=5.3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid duration 'd=5.3', but it passed")
	} else {
		if errPos != 59 {
			t.Errorf("Expected error at position 59, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidDuration_RangeError(t *testing.T) {
	melody := "60 A{d=5;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=2}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid duration 'd=5', but it passed")
	} else {
		if errPos != 7 {
			t.Errorf("Expected error at position 7, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidAlteration(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=p} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid alteration 'a=p', but it passed")
	} else {
		if errPos != 43 {
			t.Errorf("Expected error at position 43, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidAlteration_BNote(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} B{a=#;o=2;d=1/2} S A{d=2;a=b} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid alteration to B note 'a=#', but it passed")
	} else {
		if errPos != 24 {
			t.Errorf("Expected error at position 24, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidAlteration_CbemolNote(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} C{a=b;o=2;d=1/2} S A{d=2;a=b} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid alteration to C note 'a=b', but it passed")
	} else {
		if errPos != 24 {
			t.Errorf("Expected error at position 24, but got %d", errPos)
		}
	}
}

func TestValidateMelody_InvalidOctave(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{o=99;a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid octave 'o=99', but it passed")
	} else {
		if errPos != 50 {
			t.Errorf("Expected error at position 50, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_DuplicateAttributes(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=#;d=1/2} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to duplicate 'd' attribute, but it passed")
	} else {
		if errPos != 19 {
			t.Errorf("Expected error at position 19, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_CommaInsteadOfSemicolon(t *testing.T) {
	melody := "60 A{d=7/4,o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to ',' instead of ';' in A{d=7/4,o=3;a=#}, but it passed")
	} else {
		if errPos != 10 {
			t.Errorf("Expected error at position 10, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_TrailingSemicolon(t *testing.T) {
	melody := "60 A{d=7/4;} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to trailing semicolon, but it passed")
	} else {
		if errPos != 10 {
			t.Errorf("Expected error at position 10 due to trailing semicolon, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_NoCloseBrace(t *testing.T) {
	melody := "60 A{d=7/4;o=3;a=# B{o=2;d=1/4} S G{d=2}"
	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to no close brace '}' in A{d=7/4;o=3;a=# but it passed")
	} else {
		if errPos != 18 {
			t.Errorf("Expected error at position 18, but got %d", errPos)
		}
	}
}

func TestValidateMelody_Invalid_NoteDoesNotExist(t *testing.T) {
	melody := "60 Z{d=7/4;o=3;a=#} B{o=2;d=1/4} S G{d=2}"

	valid, errPos := ValidateMelody(melody)

	if valid {
		t.Errorf("Expected melody to be invalid due to invalid note 'Z', but it passed")
	} else {
		if errPos != 3 {
			t.Errorf("Expected error at position 3 for invalid note 'Z', but got %d", errPos)
		}
	}
}
