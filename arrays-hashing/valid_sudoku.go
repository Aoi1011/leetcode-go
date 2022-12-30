package arrayshashing

func isValidSudoku(board [][]byte) bool {
	type Square struct {
		x int
		y int
	}

	rows := make(map[int]byte)
	columns := make(map[int]byte)
	squares := make(map[Square]byte)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			square := Square{
				x: r / 3,
				y: c / 3,
			}
			if rows[r] == board[r][c] || columns[c] == board[r][c] || squares[square] == board[r][c] {
				return false
			}

			rows[r] = board[r][c]
			columns[c] = board[r][c]
			squares[square] = board[r][c]
		}
	}

	return true
}
