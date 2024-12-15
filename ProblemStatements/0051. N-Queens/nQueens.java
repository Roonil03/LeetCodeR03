import java.util.*;

class Solution {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> solutions = new ArrayList<>();
        char[][] board = new char[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                board[i][j] = '.';
            }
        }
        backtrack(0, board, solutions, n);
        return solutions;
    }

    private void backtrack(int row, char[][] board, List<List<String>> solutions, int n) {
        if (row == n) 
        {
            solutions.add(constructSolution(board));
            return;
        }
        for (int col = 0; col < n; col++) 
        {
            if (isValid(board, row, col, n)) 
            {
                board[row][col] = 'Q';
                backtrack(row + 1, board, solutions, n);
                board[row][col] = '.';
            }
        }
    }

    private boolean isValid(char[][] board, int row, int col, int n) {
        for (int i = 0; i < row; i++) 
        {
            if (board[i][col] == 'Q') 
            {
                return false;
            }
        }
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--)
        {
            if (board[i][j] == 'Q')
            {
                return false;
            }
        }
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++)
        {
            if (board[i][j] == 'Q')
            {
                return false;
            }
        }
        return true;
    }

    private List<String> constructSolution(char[][] board) {
        List<String> solution = new ArrayList<>();
        for (char[] row : board)
        {
            solution.add(new String(row));
        }
        return solution;
    }
}