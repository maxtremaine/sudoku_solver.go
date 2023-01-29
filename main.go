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
	startPuzzle := string(data)
	startSudoku, err := String(startPuzzle)
	if err != nil {
		log.Fatal(err)
	}
	if !startSudoku.isValid() {
		log.Fatal("The starting puzzle is not valid.")
	}
	// Run the branching and elmination process
	//var workingBranches []sudoku
	
	// Close timing
	elapsedNS := time.Since(t0)
	elapsedMS := float64(elapsedNS) * 0.000001
	log.Printf("Ran successfully in %f milliseconds.", elapsedMS)
}
