package boggle

import (
	"strings"
	"testing"
)

func TestBoardImportWhiteSpace(t *testing.T) {
	input := `P W Y R
						E N T H
						G S I Q
						O L S A`
	reader := strings.NewReader(input)
	var board Board
	err := board.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed to import board with whitespace: %s", err)
	}
}

func TestBoardImportNoWhiteSpace(t *testing.T) {
	input := `PWYR
						ENTH
						GSIQ
						OLSA`
	reader := strings.NewReader(input)
	var board Board
	err := board.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed to import board with no whitespace: %s", err)
	}
}

func TestBoardBadCharacters(t *testing.T) {
	input := `P2YR
						ENTH
						GSIQ
						OLSA`
	reader := strings.NewReader(input)
	var board Board
	err := board.ImportFromReader(reader)
	if err == nil {
		t.Error("Failed to return error for bad character.")
	}
}

func TestBoardImportTooShort(t *testing.T) {
	input := `P W Y R`
	reader := strings.NewReader(input)
	var board Board
	err := board.ImportFromReader(reader)
	if err == nil {
		t.Error("Failed to return error for insufficiant data.")
	}
}

func TestBoardImportTooLong(t *testing.T) {
	input := `P W Y R
						E N T H
						G S I Q
						O L S A
						G S I Q`
	reader := strings.NewReader(input)
	var board Board
	err := board.ImportFromReader(reader)
	if err == nil {
		t.Error("Failed to return error for excessive data.")
	}
}
