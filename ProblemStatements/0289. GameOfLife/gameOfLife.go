func gameOfLife(board [][]int)  {
    m, n := len(board), len(board[0])    
    d := [8][2]int{
        {-1, -1}, {-1, 0}, {-1, 1},{0, -1},{0, 1},{1, -1}, {1, 0}, {1, 1}
    }    
    cb := make([][]int, m)
    for i := range board {
        cb[i] = make([]int, n)
        copy(cb[i], board[i])
    }    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            ln := 0            
            for _, dir := range d {
                ni, nj := i + dir[0], j + dir[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n && cb[ni][nj] == 1 {
                    ln++
                }
            }            
            if cb[i][j] == 1 && (ln < 2 || ln > 3) {
                board[i][j] = 0
            } else if cb[i][j] == 0 && ln == 3 {
                board[i][j] = 1
            }
        }
    }
}