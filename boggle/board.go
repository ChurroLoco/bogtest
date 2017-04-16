package boggle

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
)

const BoardSize = 4

type Board struct {
	cells [BoardSize][BoardSize]string
}

// ImportFromReader validates and sets the proper runes in the board cells.
func (b *Board) ImportFromReader(r io.Reader) error {
	var valuesRead int
	expectedValueCount := BoardSize * BoardSize
	var values []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Err() != nil {

		}
		line := scanner.Text()
		line = strings.Join(strings.Fields(line), "")
		for _, rune := range line {
			if !unicode.IsLetter(rune) {
				return fmt.Errorf("Unacceptable character found in board '%c'", rune)
			}
			//fmt.Println(string(rune))
			values = append(values, string(unicode.ToLower(rune)))
			valuesRead++
			if valuesRead > expectedValueCount {
				break
			}
		}
	}
	// Make sure the data received is realistic before continuing.
	if valuesRead != expectedValueCount {
		return fmt.Errorf(fmt.Sprintf("Found at least %d runes but expected %d", valuesRead, expectedValueCount))
	}

	// Set the cells,
	// The cells is organize by [x,y] and [0,0] is the bottom left.
	for sIndex, s := range values {
		row := BoardSize - 1 - sIndex/BoardSize
		col := 0
		if sIndex != 0 {
			col = sIndex % BoardSize
		}
		b.cells[row][col] = s
	}
	return nil
}

// FindAllWords searches for words on the board given a dictionary of potential words.
func (b *Board) FindAllWords(dict *Dictionary) ([]string, error) {
	foundWords := make(map[string]bool)
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			b.searchRecursive(dict, "", foundWords, nil, location{X: x, Y: y})
		}
	}

	var results []string
	for word := range foundWords {
		results = append(results, word)
	}
	sort.Strings(results)
	return results, nil
}

// searchRecursive is an internal method to start search from a specifc cell on the board.
func (b *Board) searchRecursive(dict *Dictionary, currentWord string, foundWords map[string]bool, searched map[location]bool, loc location) (map[string]bool, error) {
	err := b.validateLocation(loc)
	if err != nil {
		return nil, err
	}

	currentVal := b.cells[loc.X][loc.Y]
	if currentVal == "q" {
		currentVal = "qu"
	}

	currentWord += string(currentVal)

	if dict.CheckWord(currentWord) {
		foundWords[currentWord] = true
	}

	if searched == nil {
		searched = make(map[location]bool)
	}

	// If there are more word that could be found given the current prefix, go for it!
	// Otherwise just stop now.
	if dict.SearchPrefix(currentWord) > 0 {
		// Make sure not sure search the current location ever again
		searched[loc] = true
		locationsToSearch := loc.GetNeighbors()

		// Remove the locations already searched
		for searchedLoc := range searched {
			for curIndex, curLoc := range locationsToSearch {
				if curLoc == searchedLoc {
					locationsToSearch = append(locationsToSearch[:curIndex], locationsToSearch[curIndex+1:]...)
				}
			}
		}

		// Recursively search the other locations for more words and aggragate the results
		for _, locToSearch := range locationsToSearch {
			newSearched := make(map[location]bool)
			for key, val := range searched {
				newSearched[key] = val
			}
			results, err := b.searchRecursive(dict, currentWord, foundWords, newSearched, locToSearch)
			if err != nil {
				return nil, err
			}

			for word := range results {
				foundWords[word] = true
			}
		}
	}
	return foundWords, nil
}

// validateLocation return true if the location fits within the bounds of the board
func (b *Board) validateLocation(l location) error {
	if l.X < 0 || l.X >= BoardSize || l.Y < 0 || l.Y >= BoardSize {
		return fmt.Errorf("Location not valid on this board %+v", l)
	}
	return nil
}

// location is an internal structure used by the boggle board to represent
// the two-dimensional location of a cell on the board
type location struct {
	X int
	Y int
}

// GetNeighbors returns all the valid neighbor locations that this cell has
func (l *location) GetNeighbors() []location {
	xMin := l.X - 1
	if xMin < 0 {
		xMin = 0
	}
	xMax := l.X + 1
	if xMax >= BoardSize {
		xMax = BoardSize - 1
	}
	yMin := l.Y - 1
	if yMin < 0 {
		yMin = 0
	}
	yMax := l.Y + 1
	if yMax >= BoardSize {
		yMax = BoardSize - 1
	}
	var results []location
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			results = append(results, location{X: x, Y: y})
		}
	}
	return results
}
