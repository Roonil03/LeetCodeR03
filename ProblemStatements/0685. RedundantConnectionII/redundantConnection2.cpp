class Solution {
public:
    vector<int> findRedundantDirectedConnection(vector<vector<int>>& edges) {
        int n = edges.size();
        vector<int> par(n + 1), c1, c2;
        for(auto& e : edges){
            if(par[e[1]]){
                c1 = {par[e[1]], e[1]};
                c2 = e;
                break;
            }
            par[e[1]] = e[0];
        }
        vector<int> root(n + 1);
        iota(root.begin(), root.end(), 0);
        auto find = [&](int i){
            while(root[i] != i){
                i = root[i] = root[root[i]];
            }
            return i;
        };
        for(auto& e : edges){
            if(e == c2){
                continue;
            }
            int u = find(e[0]), v = find(e[1]);
            if(u == v){
                return c1.empty() ? e : c1;
            }
            root[u] = v;
        }
        return c2;
    }
};