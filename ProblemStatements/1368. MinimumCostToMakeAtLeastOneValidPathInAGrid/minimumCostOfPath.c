int minCost(int** grid, int gridSize, int* gridColSize) {
    int n = gridSize;
    int m = gridColSize[0];
    int inf = INT_MAX;
    int k = 0;    
    int** dp = (int**)malloc(n * sizeof(int*));
    for (int i = 0; i < n; i++) {
        dp[i] = (int*)malloc(m * sizeof(int));
        for (int j = 0; j < m; j++) {
            dp[i][j] = inf;
        }
    }    
    int dirt[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
    int bfsCapacity = n * m;
    int bfsSize = 0;
    int** bfs = (int**)malloc(bfsCapacity * sizeof(int*));    
    for (int i = 0; i < bfsCapacity; i++) {
        bfs[i] = (int*)malloc(2 * sizeof(int));
    }

    void dfs(int x, int y) {
        if (x < 0 || x >= n || y < 0 || y >= m || dp[x][y] != inf) {
            return;
        }
        dp[x][y] = k;
        bfs[bfsSize][0] = x;
        bfs[bfsSize][1] = y;
        bfsSize++;
        int dir = grid[x][y] - 1;
        dfs(x + dirt[dir][0], y + dirt[dir][1]);
    }    
    dfs(0, 0);    
    while (bfsSize > 0) {
        k++;
        int** bfs2 = (int**)malloc(bfsSize * sizeof(int*));
        for (int i = 0; i < bfsSize; i++) {
            bfs2[i] = (int*)malloc(2 * sizeof(int));
            bfs2[i][0] = bfs[i][0];
            bfs2[i][1] = bfs[i][1];
        }
        int bfs2Size = bfsSize;
        bfsSize = 0;        
        for (int i = 0; i < bfs2Size; i++) {
            int x = bfs2[i][0];
            int y = bfs2[i][1];
            for (int j = 0; j < 4; j++) {
                int nx = x + dirt[j][0];
                int ny = y + dirt[j][1];
                dfs(nx, ny);
            }
            free(bfs2[i]);
        }
        free(bfs2);
    }
    int result = dp[n - 1][m - 1];
    for (int i = 0; i < n; i++) {
        free(dp[i]);
    }
    free(dp);    
    for (int i = 0; i < bfsCapacity; i++) {
        free(bfs[i]);
    }
    free(bfs);
    return result;
}
