package main

type sudoku struct {
	cells [81]int
}

func (puzzle sudoku) isValid() bool {
	return true
}