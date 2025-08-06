class Solution {
public:
    int closedIsland(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size(), c {0};        
        function<void(int,int)> dfs = [&](int x,int y) {
            if(x<0 || x>=m || y<0 || y>=n || grid[x][y]){
                return;
            }
            grid[x][y] = 1;
            dfs(x+1,y);
            dfs(x-1,y);
            dfs(x,y+1);
            dfs(x,y-1);
        };        
        for(int i {0}; i<m; i++){
            dfs(i,0);
            dfs(i,n-1);
        }
        for(int j {0}; j<n; j++){
            dfs(0,j);
            dfs(m-1,j);
        }
        for(int i {1}; i<m-1; i++) {
            for(int j {1}; j<n-1; j++) {
                if(grid[i][j] == 0) {
                    c++;
                    dfs(i,j);
                }
            }
        }
        return c;
    }
};