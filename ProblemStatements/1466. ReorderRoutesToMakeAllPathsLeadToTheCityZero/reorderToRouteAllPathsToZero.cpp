class Solution {
public:
    int minReorder(int n, vector<vector<int>>& connections) {
        vector<vector<pair<int, int>>> adj(n);
        for(auto& con : connections){
            int u = con[0], v = con[1];
            adj[u].push_back({v, 1});
            adj[v].push_back({u, 0});
        }
        vector<bool> vis(n, false);
        int res {0};
        dfs(0, adj, vis, res);
        return res;
    }

    void dfs(int n, vector<vector<pair<int, int>>>& adj, vector<bool>& vis, int& res){
        vis[n] = true;
        for(auto& [a, b] : adj[n]){
            if(!vis[a]){
                res += b;
                dfs(a, adj, vis, res);
            }
        }
    }
};