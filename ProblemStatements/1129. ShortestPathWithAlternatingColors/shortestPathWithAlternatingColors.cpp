class Solution {
public:
    vector<int> shortestAlternatingPaths(int n, vector<vector<int>>& redEdges, vector<vector<int>>& blueEdges) {
        vector<vector<int>> r(n), b(n);
        for (auto& e: redEdges){
            r[e[0]].push_back(e[1]);
        }
        for(auto& e : blueEdges){
            b[e[0]].push_back(e[1]);
        }
        queue<array<int, 3>> q;
        vector<vector<bool>> vis(n, vector<bool>(2, false));
        vector<int> res(n, -1);
        q.push({0, 0, 0});
        q.push({0, 1, 0});
        vis[0][0] = vis[0][1] = true;
        res[0] = 0;
        while(!q.empty()){
            auto[node, col, dist] = q.front();
            q.pop();
            int x = 1 - col;
            vector<int>& adj = (x == 0) ? r[node] : b[node];
            for(int f : adj){
                if(!vis[f][x]){
                    vis[f][x] = true;
                    if(res[f] == -1){
                        res[f] = dist + 1;
                    }
                    q.push({f, x, dist+1});
                }
            }
        }
        return res;
    }
};