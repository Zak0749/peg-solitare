package main

import (
	"fmt"
	"peg-solitare/moves"
	"peg-solitare/starting"
	"peg-solitare/types"
	"peg-solitate/pretify"
)

func main() {
	jobs := make(chan types.Board, 1)
	results := make(chan []types.Board, 1)

	go worker(jobs, results)

	for i := 0; i < 1; i++ {
		jobs <- starting.Board
	}

	for j := 0; j < 1; j++ {
		res := <-results
		for _, board := range res {
			fmt.Println(pretify.Board(board))
		}
	}
}

func worker(jobs <-chan types.Board, results chan<- []types.Board) {
	for n := range jobs {
		results <- moves.GetMoves(n)
	}
}
