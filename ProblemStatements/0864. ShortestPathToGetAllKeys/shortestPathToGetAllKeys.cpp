class Solution {
public:
    int shortestPathAllKeys(vector<string>& grid) {
        int m = grid.size(), n = grid[0].size();
        int tot {0}, x1 {0}, y1 {0};
        for(int i {0}; i < m; i++){
            for(int j {0}; j < n; j++){
                if(grid[i][j] == '@'){
                    x1 = i, y1 = j;
                } else if (grid[i][j] >= 'a' && grid[i][j] <= 'f'){
                    tot = max(tot, grid[i][j] - 'a' + 1);
                }
            }
        }
        int tar = (1 << tot) - 1;
        vector<vector<vector<bool>>> vis(m, vector<vector<bool>>(n, vector<bool>(1<<tot, false)));
        queue<tuple<int, int, int>> q;
        q.emplace(x1, y1, 0);
        vis[x1][y1][0] = true;
        int s {0};
        int dir[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
        while(!q.empty()){
            int sz = q.size();
            while(sz--){
                auto[x, y, key] = q.front();
                q.pop();
                if(key == tar){
                    return s;
                }
                for(int d {0}; d < 4; d++){
                    int nx = x + dir[d][0], ny = y + dir[d][1];
                    if(nx < 0 || nx >= m  || ny < 0 || ny >= n){
                        continue;
                    }
                    char c = grid[nx][ny];
                    if (c == '#'){
                        continue;
                    }
                    int nk = key;
                    if(c >= 'a' && c <= 'f'){
                        nk |= (1 << (c - 'a'));
                    }
                    if(c >= 'A' && c <= 'F' && !(nk & (1 << (c - 'A')))){
                        continue;
                    }
                    if(!vis[nx][ny][nk]){
                        vis[nx][ny][nk] = true;
                        q.emplace(nx, ny, nk);
                    }
                }
            }
            s++;
        }
        return -1;
    }
};