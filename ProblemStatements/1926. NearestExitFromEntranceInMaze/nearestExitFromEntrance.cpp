class Solution {
public:
    int nearestExit(vector<vector<char>>& maze, vector<int>& entrance) {
        int m = maze.size(), n = maze[0].size();
        queue<pair<int, int>> q;
        q.push({entrance[0], entrance[1]});
        vector<vector<bool>> vis(m, vector<bool>(n, false));
        int sp {0};
        int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
        while(!q.empty()){
            int sz = q.size();
            sp++;
            for(int i{0}; i <sz; i++){
                auto[r, c] = q.front();
                q.pop();
                for(auto& d : dirs){
                    int nr = r + d[0], nc = c + d[1];
                    if(nr < 0 || nr >= m || nc < 0 || nc >= n){
                        continue;
                    }
                    if(maze[nr][nc] == '+' || vis[nr][nc]){
                        continue;
                    }
                    if(nr == 0 || nr == m - 1 || nc == 0 || nc == n - 1){
                        if(!(nr == entrance[0] && nc == entrance[1])){
                            return sp;
                        }
                    }
                    vis[nr][nc] = true;
                    q.push({nr, nc});
                }
            }
        }
        return -1;
    }
};