class Solution {
public:
    int minMoves(vector<string>& matrix) {
        // int n = matrix.size(), m = matrix[0].size();
        // unordered_map<char, vector<pair<int,int>>> temp;
        // for(int i {0}; i < n; i++){
        //     for(int j {0}; j < m; j++){
        //         char c = matrix[i][j];
        //         if(c >= 'A' && c <= 'Z'){
        //             temp[c].push_back({i, j});
        //         }
        //     }
        // }
        // vector<vector<bool>> vis(n, vector<bool>(m, false));
        // unordered_set<char> u;
        // queue<tuple<int,int,int>> q;
        // q.emplace(0, 0, 0);
        // vis[0][0] = true;
        // int dirs[4][2] = {{1,0}, {-1,0}, {0,1}, {0,-1}};
        // while(!q.empty()){
        //     auto [x, y, dist] = q.front(); 
        //     q.pop();
        //     if(x == n-1 && y == m-1){
        //         return dist;
        //     }
        //     for(auto& dir : dirs){
        //         int nx = x + dir[0], ny = y + dir[1];
        //         if(nx >= 0 && nx < n && ny >= 0 && ny < m && matrix[nx][ny] != '#' && !vis[nx][ny]){
        //             vis[nx][ny] = true;
        //             q.emplace(nx, ny, dist + 1);
        //         }
        //     }
        //     char c = matrix[x][y];
        //     if(c >= 'A' && c <= 'Z' && u.find(c) == u.end()){
        //         u.insert(c);
        //         for(auto& [px, py] : temp[c]) {
        //             if(px == x && py == y){
        //                 continue;
        //             }
        //             if(!vis[px][py]){
        //                 vis[px][py] = true;
        //                 q.emplace(px, py, dist);
        //             }
        //         }
        //     }
        // }
        // return -1;

        int n = matrix.size(), m = matrix[0].size();
        unordered_map<char, vector<pair<int, int>>> cells;
        for(int i = 0; i < n; i++) {
            for(int j = 0; j < m; j++) {
                if(matrix[i][j] != '.' && matrix[i][j] != '#') {
                    cells[matrix[i][j]].push_back({i, j});
                }
            }
        }
        if(matrix[n - 1][m - 1] == '#') return -1;
        deque<tuple<int, int, int>> dq;
        vector<vector<int>> dist(n, vector<int> (m, INT_MAX));
        unordered_set<char> used;        
        dq.push_back({0, 0, 0});
        dist[0][0] = 0;
        int dx[4] = {0, 0, -1, 1};
        int dy[4] = {-1, 1, 0, 0};
        while(!dq.empty()) {
            auto [curDist, x, y] = dq.front();
            dq.pop_front();            
            if(curDist > dist[x][y]) continue;
            if(x == n - 1 && y == m - 1) return curDist;
            if(isupper(matrix[x][y]) && used.find(matrix[x][y]) == used.end()) {
                used.insert(matrix[x][y]);
                for(auto &[newX, newY]: cells[matrix[x][y]]) {
                    if(curDist < dist[newX][newY]) {
                        dist[newX][newY] = curDist;
                        dq.push_front({curDist, newX, newY});
                    }
                }
            }            
            for(int k = 0; k < 4; k++) {
                int nextX = x + dx[k], nextY = y + dy[k];
                if(isValid(nextX, nextY, n, m, matrix) && 1 + curDist < dist[nextX][nextY]) {
                    dist[nextX][nextY] = 1 + curDist;
                    dq.push_back({1 + curDist, nextX, nextY});
                }
            }
        }
        return -1;
    }

    bool isValid(int i, int j, int n, int m, vector<string> &matrix) {
        if(i < 0 || j < 0 || i >= n || j >= m) return false;
        if(matrix[i][j] == '#') return false;
        return true;
    }
};