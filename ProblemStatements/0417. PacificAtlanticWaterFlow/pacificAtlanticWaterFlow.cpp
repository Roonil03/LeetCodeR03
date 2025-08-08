class Solution {
public:
    vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
        int m = heights.size(), n = heights[0].size();
        vector<vector<bool>> a(m, vector<bool>(n, false));
        vector<vector<bool>> b(m, vector<bool>(n, false));
        for(int i{0}; i < m; i++){
            dfs(heights, a, i, 0);
            dfs(heights, b, i, n - 1);
        }
        for(int j {0}; j < n; j++){
            dfs(heights, a, 0, j);
            dfs(heights, b, m - 1, j);
        }
        vector<vector<int>> res;
        for (int i {0}; i < m; i++){
            for(int j {0}; j < n; j++){
                if(a[i][j] && b[i][j]){
                    res.push_back({i, j});
                }
            }
        }
        return res;
    }

    void dfs(vector<vector<int>>& heights, vector<vector<bool>>& vis, int i, int j){
        if(vis[i][j]){
            return;
        }
        vis[i][j] = true;
        int dirs[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        int m = heights.size(), n = heights[0].size();
        for (auto& d : dirs){
            int x = i + d[0], y = j + d[1];
            if(x >= 0 && x < m && y >= 0 && y < n && heights[x][y] >= heights[i][j]){
                dfs(heights, vis, x, y);
            }
        }
    }
};