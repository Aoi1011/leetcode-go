package arrayshashing

import "testing"

func TestValidSudoku(t *testing.T) {
	board := [][]byte{}
	board = append(board, []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'})
	board = append(board, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'})
	board = append(board, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'})
	board = append(board, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'})
	board = append(board, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'})
	board = append(board, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'})
	board = append(board, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'})
	board = append(board, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'})
	board = append(board, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'})
	output := true

	res := isValidSudoku(board)

	if res != output {
		t.Errorf("Error: %t", res)
	}

}
