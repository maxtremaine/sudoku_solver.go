package main

import (
	"os"
	"log"
)

func main() {
	data, err := os.ReadFile("io/start.sudoku")
	if err != nil {
		log.Fatal(err)
	}
	startPuzzle := string(data)
	log.Println(startPuzzle)
}