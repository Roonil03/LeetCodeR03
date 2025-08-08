class Solution {
public:
    int longestIncreasingPath(vector<vector<int>>& matrix) {
        if(matrix.empty() || matrix[0].empty()){
            return 0;
        }
        m = matrix.size();
        n = matrix[0].size();
        dp = vector<vector<int>>(m, vector<int>(n, 0));
        int res {0};
        for (int i {0}; i < m; i++){
            for(int j {0}; j < n; j++){
                res = max(res, dfs(matrix, i, j));
            }
        }
        return res;
    }

    int m, n;
    vector<vector<int>> dp;
    int dirs[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    
    int dfs(vector<vector<int>>& mat, int i, int j){
        if(dp[i][j] != 0){
            return dp[i][j];
        }
        int res {1};
        for(auto& d : dirs){
            int x = i + d[0], y = j + d[1];
            if(x >= 0 && x < m && y >= 0 && y < n && mat[x][y] > mat[i][j]){
                res = max(res, 1 + dfs(mat, x, y));
            }
        }
        dp[i][j] = res;
        return res;
    }
};