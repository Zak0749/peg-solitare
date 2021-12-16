package deep

import "peg-solitare/types"

func Copy(board types.Board) types.Board {
	var newBoard types.Board

	for y, row := range board {
		for x, val := range row {
			newBoard[y][x] = types.Tile{Exists: val.Exists, Value: val.Value}
		}
	}

	return newBoard
}
