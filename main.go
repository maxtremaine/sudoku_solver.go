package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// Start timing
	t0 := time.Now()
	// Read and create the starting puzzle
	data, err := os.ReadFile("io/start.sudoku")
	if err != nil {
		log.Fatal(err)
	}
	startString := string(data)
	startPuzzle, err := SudokuFile(startString)
	if err != nil {
		log.Fatal(err)
	}
	if !startPuzzle.isValid() {
		log.Fatal("The starting puzzle is not valid.")
	}
	// Run the branching and elmination process
	startingBlankCells := startPuzzle.getBlankCells()
	workingBranches := []sudoku{startPuzzle}
	for runIndex, _ := range startingBlankCells {
		runNumber := runIndex + 1
		var newBranches []sudoku
		for _, branch := range workingBranches {
			blankCells := branch.getBlankCells()
			lowestDFCell := blankCells[0]
			for _, possibleValue := range lowestDFCell.possibleValues {
				newBranch := branch.changeCell(lowestDFCell.index, possibleValue)
				newBranches = append(newBranches, newBranch)
			}
		}
		workingBranches = newBranches
		log.Printf("Completed run %d with %d branches.", runNumber, len(newBranches))
	}
	// Output
	solvedFile := workingBranches[0].toSudokuFile()
	err = os.WriteFile("io/finish.sudoku", []byte(solvedFile), 0666)
	if err != nil {
		log.Fatal(err)
	}
	// Close timing
	elapsedNS := time.Since(t0)
	elapsedMS := float64(elapsedNS) * 0.000001
	log.Printf("Ran successfully in %f milliseconds.", elapsedMS)
}
