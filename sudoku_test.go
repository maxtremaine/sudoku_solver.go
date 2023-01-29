package main

import (
	"reflect"
	"testing"
)

var validFile = `  abc def ghi
1 7__|_4_|__1
2 __1|___|2__
3 _6_|2_9|_8_
  -----------
4 __3|5_4|9__
5 1__|___|__4
6 __2|1_8|5__
  -----------
7 _1_|9_6|_7_
8 __8|___|4__
9 6__|_2_|__8`

var validPuzzle = sudoku{7, 0, 0, 0, 4, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 6, 0, 2, 0,
	9, 0, 8, 0, 0, 0, 3, 5, 0, 4, 9, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 2, 1, 0, 8, 5, 0,
	0, 0, 1, 0, 9, 0, 6, 0, 7, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0, 8}

func TestString(t *testing.T) {
	validSudoku, err := String(validFile)
	if err != nil {
		t.Error("File was valid, String returned error.")
	}
	if validSudoku != validPuzzle {
		t.Errorf("String expected %d, received %d.", validPuzzle, validSudoku)
	}
}

func TestIsValid(t *testing.T) {
	var output bool
	output = validPuzzle.isValid()
	if output != true {
		t.Error("isValid() expected true.")
	}

	invalidPuzzle := sudoku{7, 7, 0, 0, 4, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 6, 0, 2,
		0, 9, 0, 8, 0, 0, 0, 3, 5, 0, 4, 9, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 2, 1, 0, 8, 5,
		0, 0, 0, 1, 0, 9, 0, 6, 0, 7, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0, 8}
	output = invalidPuzzle.isValid()
	if output != false {
		t.Error("isValid() expected false.")
	}
}

func TestGetRelatedCells(t *testing.T) {
	relatedToValid := validPuzzle.getRelatedCells(1)
	relatedToOne := []int{1, 4, 6, 7}
	if !reflect.DeepEqual(relatedToValid, relatedToOne) {
		t.Errorf("getRelatedCells expected %d, got %d.", relatedToOne, relatedToValid)
	}
}

func TestChangeCell(t *testing.T) {
	puzzleWithChange := sudoku{7, 1, 0, 0, 4, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 6, 0, 2, 0,
		9, 0, 8, 0, 0, 0, 3, 5, 0, 4, 9, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 2, 1, 0, 8, 5, 0,
		0, 0, 1, 0, 9, 0, 6, 0, 7, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0, 8}
	newPuzzle := validPuzzle.changeCell(1, 1)
	if puzzleWithChange != newPuzzle {
		t.Errorf("changeCell expected %d, got %d.", puzzleWithChange, newPuzzle)
	}
	if validPuzzle == newPuzzle {
		t.Error("changeCell changed the original puzzle in place.")
	}
}
