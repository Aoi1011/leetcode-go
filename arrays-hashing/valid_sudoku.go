package arrayshashing

func isValidSudoku(board [][]byte) bool {
	type Square struct {
		x, y interface{}
	}

	rows := make(map[int]byte)
	columns := make(map[int]byte)
	squares := make(map[Square]byte)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			if _, rowOk := rows[r]; rowOk {
				return false
			} else {
				rows[r] = board[r][c]
			}

			if _, colOk := columns[c]; colOk {
				return false
			} else {
				columns[c] = board[r][c]
			}

			if _, squOk := squares[Square{r / 3, c / 3}]; squOk {
				return false
			} else {
				squares[Square{r / 3, c / 3}] = board[r][c]
			}

		}
	}

	return true
}

func isValidSudoku1(board [][]byte) bool {
	// Time Complexity O(n2)

	hashMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			row := i
			column := j

			current_val := string(board[i][j])

			if current_val == "." {
				continue
			}
			_, ok1 := hashMap[current_val+"found in row"+string(rune(row))]
			_, ok2 := hashMap[current_val+"found in column"+string(rune(column))]
			_, ok3 := hashMap[current_val+"found in grid"+string(rune(i/3))+"-"+string(rune(j/3))]

			if ok1 || ok2 || ok3 {

				return false
			} else {
				hashMap[current_val+"found in row"+string(rune(row))] = true
				hashMap[current_val+"found in column"+string(rune(column))] = true
				hashMap[current_val+"found in grid"+string(rune(i/3))+"-"+string(rune(j/3))] = true
			}
		}

	}
	return true

}
