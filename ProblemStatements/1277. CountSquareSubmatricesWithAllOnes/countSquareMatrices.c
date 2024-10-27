int countSquares(int** matrix, int matrixSize, int* matrixColSize) {
    int m = matrixSize;
    int n = *matrixColSize;
    int totalSquares = 0;
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            if (matrix[i][j] == 1 && i > 0 && j > 0) {
                matrix[i][j] = 1 + (matrix[i-1][j] < matrix[i][j-1] ? (matrix[i-1][j] < matrix[i-1][j-1] ? matrix[i-1][j] : matrix[i-1][j-1]) : (matrix[i][j-1] < matrix[i-1][j-1] ? matrix[i][j-1] : matrix[i-1][j-1]));
            }
            totalSquares += matrix[i][j];
        }
    }

    return totalSquares;
}