package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/churroloco/boggle"
)

func main() {
	flag.Parse()
	if flag.NArg() != 3 {
		log.Fatal("Usage: boggletest <dictionary_filename> <board_filename> <output_filename>")
	}
	dictFilePath := flag.Args()[0]
	boardFilePath := flag.Args()[1]
	outputPath := flag.Args()[2]

	// Attempt import the dictionary from the given filepath.
	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		panic(fmt.Errorf("Failed to read dictionary file: %s", err))
	}
	defer dictFile.Close()
	var dict boggle.Dictionary
	err = dict.ImportFromReader(dictFile)
	if err != nil {
		panic(fmt.Errorf("Failed to import dictionary file: %s", err))
	}

	// Attempt import the board from the given filepath.
	boardFile, err := os.Open(boardFilePath)
	if err != nil {
		panic(fmt.Errorf("Failed to read board file: %s", err))
	}
	defer boardFile.Close()
	var board boggle.Board
	err = board.ImportFromReader(boardFile)
	if err != nil {
		panic(fmt.Errorf("Failed to import board file: %s", err))
	}
	words, err := board.FindAllWords(&dict)
	if err != nil {
		panic(fmt.Errorf("Failed to find puzzle solution: %s", err))
	}

	// Attempt to write the output to the given output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Errorf("Failed to open output file: %s", err))
	}
	defer outFile.Close()
	for _, word := range words {
		outFile.WriteString(fmt.Sprintf("%s\n", word))
	}
}
