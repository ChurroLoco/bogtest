 
 
 
 # BoggleTest 
 
 ## Submission Notes
 - Tested in Go version 'go1.7.4 darwin/amd64'

 ## Programming Rules:
  1. Any version of Go 1.7 or higher can be used.
     A. Please document which version was used.
  2. Any core language feature may be used.
  3. No 3rd party packages can be used.
     A. If you have to use "go get", you can't use it.
  4. Sample board and dictionary files provided.
  5. Submit all source files needed to run the solution.

 ## Game Rules:
  Boggle is a word game.  The goal is to make as many words as possible
  out of the given set of letters laid out in a 4x4 grid.  Words are
  formed by starting with any letter and moving to an adjacent letter
  (up, down, left, right, or diagonal) and so-forth on.  Once a letter
  is used in a word, it can not be used again.  All words must be a
  minimum of 3 characters in length.  Finally, in this version, a "q"
  present on the board will always represent "qu", since "u" nearly always
  follows "q" in English.  If a word may be formed multiple ways on the same
  board, it only counts once.

## Example:

 ***Board:***
 ```
P W Y R
E N T H
G S I Q
O L S A
```

 ***A few possible words:***
   - pen
   - peg
   - quit
   - hit
   - slit
   - slits

 ### Command line arguments:
 ```bash
  boggletest <dictionary_filename> <board_filename> <output_filename>
 ```

 ### Dictionary:
  The dictionary file is an ASCII text file that lists acceptable words.  Each word is
  new line separated.  Words are in alphabetical order and all lowercase, utilizing
  only letters 'a' to 'z'.

 ### Board:
  The board file is an ASCII text file that is 4 lines of 4 characters.  These
  represent the game board, a 4x4 matrix of characters.  These may be mixed case.
  Whitespace is optional and should be ignored.  Only letters 'a' to 'z' or 'A'
  to 'Z' are used.

 ### Output:
  The output should be an ASCII text file (in alphabetical order) of all legal words
  possible to spell on the current board that are in the given dictionary.  These
  should be all lowercase and newline separated (same format as the dictionary).

 ### Notes:
  Your final solution should be PRODUCTION QUALITY, as if it is getting checked
  in to live production code.