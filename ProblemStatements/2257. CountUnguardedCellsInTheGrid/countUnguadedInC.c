int countUnguarded(int m, int n, int** guards, int guardsSize, int* guardsColSize, int** walls, int wallsSize, int* wallsColSize) {
    int total = m * n, guarded = 0;
    char* grid = calloc(m * n, sizeof(char));
    for (int i = 0; i < guardsSize; i++) grid[guards[i][0] * n + guards[i][1]] = 'G';
    for (int i = 0; i < wallsSize; i++) grid[walls[i][0] * n + walls[i][1]] = 'W';
    for (int i = 0; i < guardsSize; i++) {
        int x = guards[i][0], y = guards[i][1];
        for (int d = 0; d < 4; d++) {
            int dx = d == 0 ? -1 : d == 1 ? 1 : 0, dy = d < 2 ? 0 : d == 2 ? -1 : 1;
            int nx = x + dx, ny = y + dy;
            while (nx >= 0 && ny >= 0 && nx < m && ny < n) {
                if (grid[nx * n + ny] == 'W' || grid[nx * n + ny] == 'G') break;
                if (grid[nx * n + ny] == 0) {
                    grid[nx * n + ny] = 'V';
                    guarded++;
                }
                nx += dx;
                ny += dy;
            }
        }
    }
    free(grid);
    return total - guarded - guardsSize - wallsSize;
}
