package main

import (
	"encoding/json"
	"fmt"
	"main/moves"
	"main/starting"
	"main/types"
	"os"
)

func pretify(board types.Board) string {
	returning := "___________________\n"

	for _, column := range board {
		returning = returning + "|"
		for x, val := range column {
			if val.Value {
				returning = returning + "●"
			} else if val.Exists {
				returning = returning + "○"
			} else {
				returning = returning + " "
			}

			if x != 8 {
				returning = returning + " "
			}
		}

		returning = returning + "|\n"
	}

	returning += "‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾"

	return returning
}

func UniqueNonEmptyElementsOf(s []types.Board) []types.Board {
	unique := make(map[types.Board]bool, len(s))
	var us []types.Board
	for _, elem := range s {
		if !unique[elem] {
			us = append(us, elem)
			unique[elem] = true
		}
	}

	return us
}

func computeLayer(boards []types.Board, final chan<- []types.Board) {
	layer := []types.Board{}
	jobs := make(chan types.Board, len(boards))
	results := make(chan []types.Board, len(boards))

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for _, board := range boards {
		jobs <- board
	}

	close(jobs)

	for j := 0; j < len(boards); j++ {
		res := <-results
		layer = append(layer, res...)
	}
	close(results)

	if len(layer) == 0 {
		return
	}

	final <- layer

	go computeLayer(UniqueNonEmptyElementsOf(layer), final)
}

func main() {
	data := [][]types.Board{}
	final := make(chan []types.Board, 44)
	go computeLayer([]types.Board{starting.Board}, final)
	for cur := range final {
		fmt.Println(len(data)+1, "layers computed")
		data = append(data, cur)
	}
	close(final)

	file, _ := json.MarshalIndent(data, "", " ")

	os.WriteFile("test.json", file, 0644)
}

func worker(jobs <-chan types.Board, results chan<- []types.Board) {
	for n := range jobs {
		results <- moves.GetMoves(n)
	}
}
