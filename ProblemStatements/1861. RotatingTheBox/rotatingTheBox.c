char** rotateTheBox(char** box, int boxSize, int* boxColSize, int* returnSize, int** returnColumnSizes) {
    int m = boxSize, n = *boxColSize;
    char** mat = (char**)malloc(sizeof(char*) * n);
    for (int i = 0; i < n; i++) {
        mat[i] = (char*)malloc(sizeof(char) * m);
    }
    for (int j = 0; j < n; j++) {
        for (int i = m - 1; i >= 0; i--) {
            mat[j][m - 1 - i] = box[i][j];
        }
    }
    for (int i = 0; i < m; i++) {
        int idx = n - 1;
        for (int j = n - 1; j >= 0; j--) {
            if (mat[j][i] == '#') {
                mat[j][i] = '.';
                mat[idx--][i] = '#';
            } else if (mat[j][i] == '*') {
                idx = j - 1;
            }
        }
    }
    *returnSize = n;
    *returnColumnSizes = (int*)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i++) {
        (*returnColumnSizes)[i] = m;
    }
    return mat;
}
