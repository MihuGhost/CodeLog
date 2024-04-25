package main

func main() {}

func solveSudoku(board [][]byte) {
	dfs(board)
}

func dfs(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}

			for k := '1'; k <= '9'; k++ {
				if isValid(board, byte(k), i, j) {
					board[i][j] = byte(k)
					if dfs(board) {
						return true
					}
					board[i][j] = '.'
				}
			}

			return false
		}
	}
	return true
}

func isValid(board [][]byte, num byte, row, col int) bool {
	//同一行
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	//同一列
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	//小宫格
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}
