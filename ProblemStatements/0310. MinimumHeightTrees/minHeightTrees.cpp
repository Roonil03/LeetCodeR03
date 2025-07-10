class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        if (n == 1){
            return {0};
        }
        vector<vector<int>> adj(n);
        vector<int> deg(n, 0);
        for(auto& e : edges){
            int u = e[0], v = e[1];
            adj[u].push_back(v);
            adj[v].push_back(u);
            deg[u]++;
            deg[v]++;
        }
        queue<int> q;
        for (int i{0}; i < n; i++){
            if (deg[i] == 1){
                q.push(i);
            }
        }
        int rem = n;
        while(rem > 2){
            int l = q.size();
            rem -= l;
            for (int i {0}; i < l; i++){
                int a = q.front();
                q.pop();
                for (int n : adj[a]){
                    deg[n]--;
                    if(deg[n] == 1){
                        q.push(n);
                    }
                }
            }
        }
        vector<int>res;
        while(!q.empty()){
            res.push_back(q.front());
            q.pop();
        }
        return res;
    }
};