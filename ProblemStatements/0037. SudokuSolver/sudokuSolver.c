#include <stdbool.h>

typedef struct {
    bool rows[9][9];
    bool cols[9][9];
    bool boxes[9][9];
} Sudoku;

bool isValid(Sudoku* s, int r, int c, int num) {
    int b = (r / 3) * 3 + c / 3;
    return !s->rows[r][num] && !s->cols[c][num] && !s->boxes[b][num];
}

void placeNumber(Sudoku* s, char** board, int r, int c, int num) {
    board[r][c] = num + '1';
    int b = (r / 3) * 3 + c / 3;
    s->rows[r][num] = s->cols[c][num] = s->boxes[b][num] = true;
}

void removeNumber(Sudoku* s, char** board, int r, int c, int num) {
    board[r][c] = '.';
    int b = (r / 3) * 3 + c / 3;
    s->rows[r][num] = s->cols[c][num] = s->boxes[b][num] = false;
}

bool backtrack(Sudoku* s, char** board, int r, int c) {
    if (r == 9) return true;
    if (c == 9) return backtrack(s, board, r + 1, 0);
    if (board[r][c] != '.') return backtrack(s, board, r, c + 1);
    for (int num = 0; num < 9; num++) {
        if (isValid(s, r, c, num)) {
            placeNumber(s, board, r, c, num);
            if (backtrack(s, board, r, c + 1)) return true;
            removeNumber(s, board, r, c, num);
        }
    }
    return false;
}

void solveSudoku(char** board, int boardSize, int* boardColSize) {
    Sudoku s = {0};
    for (int r = 0; r < boardSize; r++) {
        for (int c = 0; c < boardColSize[r]; c++) {
            if (board[r][c] != '.') {
                int num = board[r][c] - '1';
                int b = (r / 3) * 3 + c / 3;
                s.rows[r][num] = s.cols[c][num] = s.boxes[b][num] = true;
            }
        }
    }
    backtrack(&s, board, 0, 0);
}