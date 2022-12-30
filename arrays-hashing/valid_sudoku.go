package arrayshashing

func isValidSudoku(board [][]byte) bool {
	type Column struct {
		index int
		value byte
	}
	type Loc struct {
		x int
		y int
	}

	columns := make(map[Column]bool)
	squares := make(map[Loc]byte)

	for index, chars := range board {
		for j, char := range chars {
			if char != '.' {
				rows := make(map[byte]bool)
				if _, found := rows[char]; found {
					return false
				}
				rows[char] = true

				column := Column{
					index: j,
					value: char,
				}
				if _, found := columns[column]; found {
					return false
				}
				columns[column] = true

				loc := Loc{
					x: (index + 1) / 3,
					y: (j + 1) / 3,
				}
				if _, found := squares[loc]; found {
					return false
				}
				squares[loc] = char
			}

		}
	}

	return true
}
