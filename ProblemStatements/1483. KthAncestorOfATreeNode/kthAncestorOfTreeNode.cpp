class TreeAncestor {
public:
    vector<vector<int>> anc;
    int logn;

    TreeAncestor(int n, vector<int>& parent) {
         logn = 32 - __builtin_clz(n);
         anc.assign(n, vector<int>(logn, -1));
         for(int i {0}; i < n;i++){
            if(parent[i] != -1){
                anc[i][0] = parent[i];
            }
         }
         for(int j {1}; j < logn; j++){
            for(int i {0}; i < n; i++){
                if(anc[i][j-1] != -1){
                    anc[i][j] = anc[anc[i][j-1]][j-1];
                }
            }
         }
    }
    
    int getKthAncestor(int node, int k) {
        for(int i {0}; i < logn; i++){
            if ((k >> i) & 1){
                if(node == -1){
                    return -1;
                }
                node = anc[node][i];
            }
        }
        return node;
    }
};

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * TreeAncestor* obj = new TreeAncestor(n, parent);
 * int param_1 = obj->getKthAncestor(node,k);
 */