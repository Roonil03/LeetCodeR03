func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	rn := []byte(word)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(board, rn, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, w []byte, i, j, id int) bool {
	if id == len(w) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != w[id] {
		return false
	}
	temp := board[i][j]
	board[i][j] = '#'
	found := dfs(board, w, i+1, j, id+1) || dfs(board, w, i-1, j, id+1) || dfs(board, w, i, j+1, id+1) || dfs(board, w, i, j-1, id+1)
	board[i][j] = temp
	return found
}