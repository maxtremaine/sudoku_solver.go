package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	data, err := os.ReadFile("io/start.sudoku")
	if err != nil {
		log.Fatal(err)
	}
	startPuzzle := string(data)
	fmt.Println(startPuzzle)
}