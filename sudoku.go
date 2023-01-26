package main

type sudoku struct {
	cells [81]uint8
}

func (puzzle sudoku) isValid() bool {
	return true
}