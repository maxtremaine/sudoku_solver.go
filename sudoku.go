package main

import (
	"sort"
	"strconv"
	"strings"
)

var groups = [][]int{
	// Rows
	{0, 1, 2, 3, 4, 5, 6, 7, 8},
	{9, 10, 11, 12, 13, 14, 15, 16, 17},
	{18, 19, 20, 21, 22, 23, 24, 25, 26},
	{27, 28, 29, 30, 31, 32, 33, 34, 35},
	{36, 37, 38, 39, 40, 41, 42, 43, 44},
	{45, 46, 47, 48, 49, 50, 51, 52, 53},
	{54, 55, 56, 57, 58, 59, 60, 61, 62},
	{63, 64, 65, 66, 67, 68, 69, 70, 71},
	{72, 73, 74, 75, 76, 77, 78, 79, 80},
	// Columns
	{0, 9, 18, 27, 36, 45, 54, 63, 72},
	{1, 10, 19, 28, 37, 46, 55, 64, 73},
	{2, 11, 20, 29, 38, 47, 56, 65, 74},
	{3, 12, 21, 30, 39, 48, 57, 66, 75},
	{4, 13, 22, 31, 40, 49, 58, 67, 76},
	{5, 14, 23, 32, 41, 50, 59, 68, 77},
	{6, 15, 24, 33, 42, 51, 60, 69, 78},
	{7, 16, 25, 34, 43, 52, 61, 70, 79},
	{8, 17, 26, 35, 44, 53, 62, 71, 80},
	// Boxes
	{0, 1, 2, 9, 10, 11, 18, 19, 20},
	{3, 4, 5, 12, 13, 14, 21, 22, 23},
	{6, 7, 8, 15, 16, 17, 24, 25, 26},
	{27, 28, 29, 36, 37, 38, 45, 46, 47},
	{30, 31, 32, 39, 40, 41, 48, 49, 50},
	{33, 34, 35, 42, 43, 44, 51, 52, 53},
	{54, 55, 56, 63, 64, 65, 72, 73, 74},
	{57, 58, 59, 66, 67, 68, 75, 76, 77},
	{60, 61, 62, 69, 70, 71, 78, 79, 80},
}

var fileToStringConversionIndexes = []int{16, 17, 18, 20, 21, 22, 24, 25, 26, 30, 31, 32, 34, 35,
	36, 38, 39, 40, 44, 45, 46, 48, 49, 50, 52, 53, 54, 72, 73, 74, 76, 77, 78, 80, 81, 82, 86,
	87, 88, 90, 91, 92, 94, 95, 96, 100, 101, 102, 104, 105, 106, 108, 109, 110, 128, 129, 130,
	132, 133, 134, 136, 137, 138, 142, 143, 144, 146, 147, 148, 150, 151, 152, 156, 157, 158, 160,
	161, 162, 164, 165, 166}

type sudoku [81]int

func SudokuFile(s string) (sudoku, error) {
	var output sudoku
	for outputIndex, stringIndex := range fileToStringConversionIndexes {
		cellValue := string(s[stringIndex])
		if cellValue != "_" {
			intValue, err := strconv.Atoi(cellValue)
			if err != nil {
				return sudoku{}, err
			}
			output[outputIndex] = intValue
		}
	}
	return output, nil
}

func (puzzle sudoku) isValid() bool {
	for _, group := range groups {
		var groupValues []int
		for _, cellIndex := range group {
			cellValue := puzzle[cellIndex]
			if cellValue != 0 {
				groupValues = append(groupValues, cellValue)
			}
		}
		if hasDuplicates(groupValues) {
			return false
		}
	}

	return true
}

func (puzzle sudoku) getRelatedCells(refIndex int) []int {
	var relatedIndexes []int
	for _, group := range groups {
		for _, i := range group {
			if i == refIndex {
				relatedIndexes = append(relatedIndexes, group...)
			}
		}
	}
	relatedIndexes = uniqueNumbers(relatedIndexes)
	var relatedValues []int
	for _, i := range relatedIndexes {
		value := puzzle[i]
		if value != 0 {
			relatedValues = append(relatedValues, value)
		}
	}
	relatedValues = uniqueNumbers(relatedValues)
	sort.Ints(relatedValues)
	return relatedValues
}

func (puzzle sudoku) changeCell(index int, value int) sudoku {
	var newPuzzle = puzzle
	newPuzzle[index] = value
	return newPuzzle
}

func (puzzle sudoku) getBlankCells() []blankCell {
	var blankCells []blankCell
	for i, cell := range puzzle {
		if cell == 0 {
			relatedCells := puzzle.getRelatedCells(i)
			possibleValues := getMissingDigits(relatedCells)
			newBlankCell := blankCell{i, possibleValues}
			blankCells = append(blankCells, newBlankCell)
		}
	}
	sort.Slice(blankCells, func(i, j int) bool {
		return len(blankCells[i].possibleValues) < len(blankCells[j].possibleValues)
	})
	return blankCells
}

func (puzzle sudoku) toSudokuFile() string {
	output := []string{" ", " ", "a", "b", "c", " ", "d", "e", "f", " ", "g", "h", "i", "\n",
		"1", " ", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "\n", "2", " ", "_", "_",
		"_", "|", "_", "_", "_", "|", "_", "_", "_", "\n", "3", " ", "_", "_", "_", "|", "_", "_",
		"_", "|", "_", "_", "_", "\n", " ", " ", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-",
		"-", "\n", "4", " ", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "\n", "5", " ",
		"_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "\n", "6", " ", "_", "_", "_", "|",
		"_", "_", "_", "|", "_", "_", "_", "\n", " ", " ", "-", "-", "-", "-", "-", "-", "-", "-",
		"-", "-", "-", "\n", "7", " ", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "\n",
		"8", " ", "_", "_", "_", "|", "_", "_", "_", "|", "_", "_", "_", "\n", "9", " ", "_", "_",
		"_", "|", "_", "_", "_", "|", "_", "_", "_"}

	for puzzleIndex, fileIndex := range fileToStringConversionIndexes {
		stringValue := strconv.Itoa(puzzle[puzzleIndex])
		if stringValue == "0" {
			output[fileIndex] = "_"
		} else {
			output[fileIndex] = stringValue
		}
	}
	return strings.Join(output, "")
}
