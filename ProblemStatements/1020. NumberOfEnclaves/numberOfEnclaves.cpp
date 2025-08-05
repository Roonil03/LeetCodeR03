class Solution {
public:
    int numEnclaves(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<bool>> vis(m, vector<bool>(n, false));
        queue<pair<int, int>> q;
        for (int i{0}; i < m;i++){
            for(int j : {0, n - 1}){
                if(grid[i][j] == 1 && !vis[i][j]){
                    q.emplace(i, j);
                    vis[i][j] = true;
                }
            }
        }
        for (int j {0}; j < n; j++){
            for (int i : {0, m - 1}){
                if (grid[i][j] == 1 && !vis[i][j]){
                    q.emplace(i, j);
                    vis[i][j] = true;
                }
            }
        }
        int dir[4][2] = {{0,1},{1,0},{0,-1},{-1,0}};
        while (!q.empty()){
            auto[x,y] = q.front();
            q.pop();
            for(auto& d : dir){
                int nx = x + d[0];
                int ny = y + d[1];
                if(nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 && !vis[nx][ny]){
                    vis[nx][ny] = true;
                    q.emplace(nx, ny);
                }
            }
        }
        int res {0};
        for(int i {0}; i < m; i++){
            for (int j {0}; j < n; j++){
                if(grid[i][j] == 1 && !vis[i][j]){
                    res++;
                }
            }
        }
        return res;
    }
};