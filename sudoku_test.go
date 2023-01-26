package main

import "testing"

func TestIsValid(t *testing.T) {
	validPuzzle := [81]int{ 7, 0, 0, 0, 4, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 6, 0, 2, 0,
        9, 0, 8, 0, 0, 0, 3, 5, 0, 4, 9, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 2, 1, 0, 8,
        5, 0, 0, 0, 1, 0, 9, 0, 6, 0, 7, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 6, 0, 0, 0, 2, 0, 0,
        0, 8 }
	validSudoku := sudoku{cells: validPuzzle}
	var output bool
	output = validSudoku.isValid()
	if output != true {
		t.Error("isValid() expected true.")
	}

	invalidPuzzle := [81]int{ 7, 7, 0, 0, 4, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 6, 0,
		2, 0, 9, 0, 8, 0, 0, 0, 3, 5, 0, 4, 9, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 2,
		1, 0, 8, 5, 0, 0, 0, 1, 0, 9, 0, 6, 0, 7, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 6, 0, 0,
		0, 2, 0, 0, 0, 8 }
	invalidSudoku := sudoku{cells: invalidPuzzle}
	output = invalidSudoku.isValid()
	if output != false {
		t.Error("isValid() expected false.")
	}
}