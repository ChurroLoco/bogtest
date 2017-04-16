package boggle

import (
	"io"
	"text/scanner"
)

type Dictionary struct {
	words       map[string]bool // A unique list of acceptable boggle words
	prefixIndex map[string]int  // A map of prefix to how many words use that prefix
}

// ImportFromReader validates and adds words from the read into the dictionary.
// Additionally this method builds an index of prefixes that can be used
// to quickly determine if any of the words in the dictionary start with
// a specific sub-string.
func (d *Dictionary) ImportFromReader(r io.Reader) error {
	if d.words == nil {
		d.words = make(map[string]bool)
	}
	if d.prefixIndex == nil {
		d.prefixIndex = make(map[string]int)
	}
	var s scanner.Scanner
	s.Init(r)
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		if validateDictionaryWord(s.TokenText()) {
			word := s.TokenText()
			d.words[word] = true
			// Build an index of prefixes for each word
			for i := 1; i < len(word); i++ {
				prefix := word[:i]
				d.prefixIndex[prefix] = d.prefixIndex[prefix] + 1
			}
		}
	}
	return nil
}

// CheckWord return true if the given string is found in this dictionary.
func (d Dictionary) CheckWord(w string) bool {
	_, exists := d.words[w]
	return exists
}

// SearchPerfix returns an array of strings in the dictionary that start with the prefix
func (d Dictionary) SearchPrefix(p string) int {
	count, _ := d.prefixIndex[p]
	return count
}

// validateDictionaryWord returns true if a word can be included in the boggle dictionary.
func validateDictionaryWord(w string) bool {
	wordLength := len(w)
	// All words over 17 characters can't be used in boggle. (Considering 'q' as 'qu')
	if wordLength > 17 {
		return false
	}
	// Boggle words have to be at least 3 characters long.
	if wordLength < 3 {
		return false
	}
	return true
}
