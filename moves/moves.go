package moves

import (
	"main/deep"
	"main/types"
)

func GetMoves(board types.Board) []types.Board {
	moves := []types.Board{}
	for y, row := range board {
		for x, val := range row {
			if !val.Exists || !val.Value {
				continue
			}

			if y > 1 && board[y-1][x].Value && board[y-2][x].Exists && !board[y-2][x].Value {
				newBoard := deep.Copy(board)
				newBoard[y][x].Value = false
				newBoard[y-1][x].Value = false
				newBoard[y-2][x].Value = true
				moves = append(moves, newBoard)
			}

			if y < 7 && board[y+1][x].Value && board[y+2][x].Exists && !board[y+2][x].Value {
				newBoard := deep.Copy(board)
				newBoard[y][x].Value = false
				newBoard[y+1][x].Value = false
				newBoard[y+2][x].Value = true
				moves = append(moves, newBoard)
			}

			if x > 1 && row[x-1].Value && row[x-2].Exists && !row[x-2].Value {
				newBoard := deep.Copy(board)
				newBoard[y][x].Value = false
				newBoard[y][x-1].Value = false
				newBoard[y][x-2].Value = true
				moves = append(moves, newBoard)
			}

			if x < 7 && row[x+1].Value && row[x+2].Exists && !row[x+2].Value {
				newBoard := deep.Copy(board)
				newBoard[y][x].Value = false
				newBoard[y][x+1].Value = false
				newBoard[y][x+2].Value = true
				moves = append(moves, newBoard)
			}

		}
	}

	return moves
}
