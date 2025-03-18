func solve(board [][]byte) {
	r, c := len(board), len(board[0])
	for i := 0; i < c; i++ {
		dfs(board, 0, i)
		dfs(board, r-1, i)
	}
	for i := 0; i < r; i++ {
		dfs(board, i, 0)
		dfs(board, i, c-1)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row, col int) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != 'O' {
		return
	}
	board[row][col] = 'T'
	dfs(board, row+1, col)
	dfs(board, row-1, col)
	dfs(board, row, col+1)
	dfs(board, row, col-1)
}