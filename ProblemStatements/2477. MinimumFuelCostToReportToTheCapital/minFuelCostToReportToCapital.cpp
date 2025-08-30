class Solution {
public:
    long long minimumFuelCost(vector<vector<int>>& roads, int seats) {
        int n = roads.size() + 1;
        vector<vector<int>> adj(n);
        for(auto& r : roads){
            adj[r[0]].push_back(r[1]);
            adj[r[1]].push_back(r[0]);
        }
        long long res {0};
        dfs(0, -1, adj, seats, &res);
        return res;
    }

    int dfs(int node, int par, vector<vector<int>>& adj, int seats, long long* res){
        int a {1};
        for(int n : adj[node]){
            if(n == par){
                continue;
            }
            int b = dfs(n, node, adj, seats, res);
            a += b;
            int car = (b + seats - 1) / seats;
            *res += car;
        }
        return a;
    }
};