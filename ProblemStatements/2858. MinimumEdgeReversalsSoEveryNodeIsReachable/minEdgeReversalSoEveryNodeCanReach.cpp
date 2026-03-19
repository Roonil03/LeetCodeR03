class Solution {
public:
    vector<vector<pair<int, int>>> adj;
    vector<int> res;

    int dfs1(int u, int p){
        int cost {0};
        for(auto&[v, w] : adj[u]){
            if(v == p){
                continue;
            }
            cost += w + dfs1(v, u);
        }
        return cost;
    }

    void dfs2(int u, int p){
        for(auto&[v, w] : adj[u]){
            if(v == p){
                continue;
            }
            if(w == 0){
                res[v] = res[u] + 1;
            } else{
                res[v] = res[u] - 1;
            }
            dfs2(v, u);
        }
    }

    vector<int> minEdgeReversals(int n, vector<vector<int>>& edges) {
        adj.resize(n);
        for(auto& i: edges){
            int u = i[0], v = i[1];
            adj[u].push_back({v, 0});
            adj[v].push_back({u, 1});
        }
        res.resize(n);
        res[0] = dfs1(0, -1);
        dfs2(0, -1);
        return res;
    }
};