package main

import (
	"bytes"
)

func solveSudoku(board [][]byte) {
	var dfs func(i, j int) bool
	conflict := func(i, j int, c byte) bool {
		if bytes.Contains(board[i], []byte{c}) {
			return true
		}

		for k := 0; k <= 8; k++ {
			if board[k][j] == c {
				return true
			}
		}

		for l := 0; l <= 2; l++ {
			for m := 0; m <= 2; m++ {
				if board[(i/3)*3+l][(j/3)*3+m] == c {
					return true
				}
			}
		}

		return false
	}

	dfs = func(i, j int) bool {
		if board[i][j] != '.' {
			if i == 8 && j == 8 {
				return true
			} else if j == 8 {
				return dfs(i+1, 0)
			}
			return dfs(i, j+1)
		}

		for c := '1'; c <= '9'; c++ {
			if conflict(i, j, byte(c)) {
				continue
			}

			board[i][j] = byte(c)
			if i == 8 && j == 8 {
				return true
			} else if j == 8 {
				if dfs(i+1, 0) {
					return true
				}
			} else if dfs(i, j+1) {
				return true
			}
			board[i][j] = '.'
		}

		return false
	}

	dfs(0, 0)
}

// add state cache
func solveSudoku2(board [][]byte) {
	var row_state, col_state [9][9]bool
	var box_state [3][3][9]bool

	var dfs func(id int) bool
	dfs = func(id int) bool {
		if id == 81 {
			return true
		}
		r, c := id/9, id%9
		if board[r][c] != '.' {
			return dfs(id + 1)
		}
		for d := 0; d < 9; d++ {
			if row_state[r][d] || col_state[d][c] || box_state[r/3][c/3][d] {
				continue
			}

			row_state[r][d], col_state[d][c], box_state[r/3][c/3][d] = true, true, true
			board[r][c] = byte('1' + d)
			if dfs(id + 1) {
				return true
			}
			board[r][c] = '.'
			row_state[r][d], col_state[d][c], box_state[r/3][c/3][d] = false, false, false
		}

		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			c := int(board[i][j]) - '1'
			row_state[i][c], col_state[c][j], box_state[i/3][j/3][c] = true, true, true
		}
	}

	dfs(0)
}
