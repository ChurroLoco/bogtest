package main

import (
	"./boggle"
	"strings"
	"testing"
)

const TestDictionary = `entry
								hint
								hints
								his
								hit
								hits
								its
								line
								lines
								list
								log
								logs
								lose
								losing
								loss
								lost
								new
								pen
								quit
								quits
								saint
								sent
								sit
								sits
								thin
								thing
								things
								this
								tin
								try
								went
								west
								quadricentennials`

func TestBoardOne(t *testing.T) {
	reader := strings.NewReader(TestDictionary)
	var dict boggle.Dictionary
	err := dict.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed to import TestDictionary: %s", err)
	}

	boardInput := `P W Y R
								 E N T H
								 G S I Q
								 O L S A`
	reader = strings.NewReader(boardInput)
	var board boggle.Board
	err = board.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed to import board: %s", err)
	}
	words, err := board.FindAllWords(&dict)
	if err != nil {
		t.Errorf("Failed to find words: %s", err)
	}

	actualWordCount := len(words)
	expectedWordCount := 32
	if actualWordCount != expectedWordCount {
		t.Errorf("Wrong number of words: Found '%d' and expect '%d'", actualWordCount, expectedWordCount)
	}
}

func TestCrazyLongWordBoard(t *testing.T) {
	reader := strings.NewReader(TestDictionary)
	var dict boggle.Dictionary
	err := dict.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed to import TestDictionary: %s", err)
	}

	boardInput := `Q A D R
								 N E C I
								 T E I A
								 N N S L`

	reader = strings.NewReader(boardInput)
	var board boggle.Board
	err = board.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed to import board: %s", err)
	}
	words, err := board.FindAllWords(&dict)
	if err != nil {
		t.Errorf("Failed to find words: %s", err)
	}

	crazyWord := "quadricentennials"
	var foundCrazyWord bool
	for _, word := range words {
		if word == crazyWord {
			foundCrazyWord = true
			break
		}
	}
	if !foundCrazyWord {
		t.Errorf("Failed to find crazy long word '%s'", crazyWord)
	}
}
