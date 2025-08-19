class UnionFind {
private:
    vector<int> parent;
    vector<int> rank;
    int components;

public:
    UnionFind(int n) : parent(n), rank(n, 0), components(n) {
        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }
    }
    
    int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }
    
    void unite(int x, int y) {
        int px = find(x);
        int py = find(y);        
        if (px == py){
            return;
        }
        if (rank[px] < rank[py]) {
            swap(px, py);
        }
        parent[py] = px;
        if (rank[px] == rank[py]) {
            rank[px]++;
        }
        components--;
    }
    
    int getComponents() const {
        return components;
    }
};

class Solution {
public:
    int numSimilarGroups(vector<string>& strs) {
        int n = strs.size();
        if (n <= 1){
            return n;
        }
        UnionFind uf(n);
        for(int i {0}; i < n; i++){
            for(int j = i + 1; j < n; j++){
                if(h1(strs[i], strs[j])){
                    uf.unite(i, j);
                }
            }
        }
        return uf.getComponents();
    }

    bool h1(string& s1, string& s2){
        if(s1.length() != s2.length()){
            return false;
        }
        vector<int> d;
        for(int i {0}; i < s1.length(); i++){
            if(s1[i] != s2[i]){
                d.push_back(i);
            }
        }
        if(d.size() == 0){
            return true;
        }
        if(d.size() == 2){
            int i = d[0];
            int j = d[1];
            return s1[i] == s2[j] && s1[j] == s2[i];
        }
        return false;
    }
};