package main

import (
	"fmt"
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
)

// Command: go run cli.go

func main() {
	// i := puzzle.Generate("hard")
	i := board.Basic()
	sudoku := puzzle.New(i)
	sudoku.Grade()
	sudoku.SlowSolve()
	board, status, dif := sudoku.Display()
	fmt.Println(board, status, dif)
}