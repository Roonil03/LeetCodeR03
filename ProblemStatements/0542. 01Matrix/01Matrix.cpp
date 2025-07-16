class Solution {
public:
    vector<vector<int>> updateMatrix(vector<vector<int>>& mat) {
        int m = mat.size();
        int n = mat[0].size();
        vector<vector<int>>res(m, vector<int>(n, INT_MAX));
        queue<pair<int, int>> q;
        for (int i {0}; i < m; i++){
            for (int j {0}; j < n; j++){
                if(mat[i][j] == 0){
                    res[i][j] = 0;
                    q.emplace(i, j);
                }
            }
        }
        int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0,1}};
        while(!q.empty()){
            auto[r, c] = q.front();
            q.pop();
            for(auto &d : dirs){
                int nr = r + d[0], nc = c + d[1];
                if(nr >= 0 && nr < m && nc >= 0 && nc < n){
                    if(res[nr][nc] > res[r][c] + 1){
                        res[nr][nc] = res[r][c] + 1;
                        q.emplace(nr, nc);
                    }
                }
            }
        }
        return res;
    }
};