package pretify

import "peg-solitare/types"

func Board(board types.Board) string {
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
