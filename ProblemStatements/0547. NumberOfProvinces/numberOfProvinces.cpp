class UnionFind {
private:
    vector<int> parent, rank;
    int count;

public:
    UnionFind(int size) {
        parent.resize(size);
        rank.resize(size, 0);
        count = size;
        for (int i = 0; i < size; i++) {
            parent[i] = i;
        }
    }

    int find(int x) {
        if (parent[x] != x){
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }

    void union_set(int x, int y) {
        int xset = find(x), yset = find(y);
        if (rank[xset] < rank[yset]) {
            parent[xset] = yset;
        } else if (rank[xset] > rank[yset]) {
            parent[yset] = xset;
        } else {
            parent[yset] = xset;
            rank[xset]++;
        }
    }
    
};

class Solution {
public:
    int findCircleNum(vector<vector<int>>& isConnected) {
        int n = isConnected.size();
    //     vector<bool>vis(n, false);
    //     int res {0};
    //     for(int i {0} ; i < n; i++){
    //         if(!vis[i]){
    //             dfs(i, isConnected, vis);
    //             res++;
    //         }
    //     }
    //     return res;
    // }

    // void dfs(int a, vector<vector<int>>& c, vector<bool>& vis){
    //     vis[a] = true;
    //     for(int j {0}; j < c.size(); j++){
    //         if(c[a][j] && !vis[j]){
    //             dfs(j, c, vis);
    //         }
    //     }
        UnionFind dsu(n);
        int res = n;
        for (int i {0}; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (isConnected[i][j] && dsu.find(i) != dsu.find(j)) {
                    res--;
                    dsu.union_set(i, j);
                }
            }
        }
        return res;
    }
};