class Solution {
public:
    vector<int> sortItems(int n, int m, vector<int>& group, vector<vector<int>>& beforeItems) {
        int next = m;
        for(int i = 0; i < n; i++){
            if(group[i] == -1){
                group[i] = next++;
            }
        }
        int g = next;
        vector<vector<int>> gadj(g);
        vector<int> gindeg(g, 0);
        vector<vector<int>> iadj(n);
        vector<int> iindeg(n, 0);
        for(int i = 0; i < n; i++){
            int gi = group[i];
            for(int prev : beforeItems[i]){
                int gj = group[prev];
                if(gj != gi){
                    gadj[gj].push_back(gi);
                }
                iadj[prev].push_back(i);
            }
        }
        for(int i = 0; i < g; i++){
            for(int j : gadj[i]){
                gindeg[j]++;
            }
        }
        for(int i = 0; i < n; i++){
            for(int j : iadj[i]){
                iindeg[j]++;
            }
        }
        queue<int> q;
        for(int i = 0; i < g; i++){
            if(gindeg[i] == 0){
                q.push(i);
            }
        }
        vector<int> go;
        while(!q.empty()){
            int temp = q.front();
            q.pop();
            go.push_back(temp);
            for(int nxt : gadj[temp]){
                if(--gindeg[nxt] == 0){
                    q.push(nxt);
                }
            }
        }
        if(go.size() != g){
            return {};
        }
        queue<int> qi;
        for(int i = 0; i < n; i++){
            if(iindeg[i] == 0){
                qi.push(i);
            }
        }
        vector<int> io;
        while(!qi.empty()){
            int u = qi.front();
            qi.pop();
            io.push_back(u);
            for(int v : iadj[u]){
                if(--iindeg[v] == 0){
                    qi.push(v);
                }
            }
        }
        if(io.size() != n){
            return {};
        }
        vector<vector<int>> a(g);
        for(int i : io){
            a[group[i]].push_back(i);
        }
        vector<int> res;
        for(int i : go){
            for(int u : a[i]){
                res.push_back(u);
            }
        }
        return res;
    }
};