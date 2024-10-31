#include <stdlib.h>
#include <stdio.h>

int max(int a, int b){
    return a > b ? a : b;
}

int findMaxMoves(int** grid, int gridSize, int* gridColSize, int row, int col, int** memo) {
    if (col == gridColSize[row] - 1)
        return 0;
    if (memo[row][col] != -1)
        return memo[row][col];
    int max_move = 0;
    if (row > 0 && grid[row - 1][col + 1] > grid[row][col]) {
        max_move = max(max_move, 1 + findMaxMoves(grid, gridSize, gridColSize, row - 1, col + 1, memo));
    }
    if (grid[row][col + 1] > grid[row][col]) {
        max_move = max(max_move, 1 + findMaxMoves(grid, gridSize, gridColSize, row, col + 1, memo));
    }
    if (row < gridSize - 1 && grid[row + 1][col + 1] > grid[row][col]) {
        max_move = max(max_move, 1 + findMaxMoves(grid, gridSize, gridColSize, row + 1, col + 1, memo));
    }
    memo[row][col] = max_move;
    return max_move;
}

int maxMoves(int** grid, int gridSize, int* gridColSize) {
    int** memo = (int**)malloc(gridSize * sizeof(int*));
    for (int i = 0; i < gridSize; i++) {
        memo[i] = (int*)malloc(gridColSize[i] * sizeof(int));
        for (int j = 0; j < gridColSize[i]; j++) {
            memo[i][j] = -1;
        }
    }
    int max_moves = 0;
    for (int i = 0; i < gridSize; i++) {
        max_moves = max(max_moves, findMaxMoves(grid, gridSize, gridColSize, i, 0, memo));
    }
    for (int i = 0; i < gridSize; i++) {
        free(memo[i]);
    }
    free(memo);
    return max_moves;
}
