package main

import "strings"

func main() {
}

func solveNQueens(n int) [][]string {
	var res [][]string
	chessBoard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "."
		}
	}

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(chessBoard[i], "")
			}
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			if isValid(chessBoard, i, row, n) {
				chessBoard[row][i] = "Q"
				dfs(row + 1)
				chessBoard[row][i] = "."
			}
		}
	}
	dfs(0)
	return res
}

func isValid(chessBoard [][]string, col, row, n int) bool {
	for i := 0; i < row; i++ {
		if chessBoard[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	return true
}
