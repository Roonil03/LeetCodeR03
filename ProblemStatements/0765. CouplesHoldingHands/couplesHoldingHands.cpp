class Solution {
public:
    int minSwapsCouples(vector<int>& row) {
        int n = row.size() >> 1, k = n;
        vector<int>p (n);
        iota(p.begin(), p.end(), 0);
        auto find = [&](int x){
            while(x != p[x]){
                x = p[x] = p[p[x]];
            }
            return x;
        };
        for(int i {0}; i < n; i++){
            int u = find(row[i << 1] >> 1), v = find(row[(i<<1) | 1] >> 1);
            if(u != v){
                p[u] = v;
                k--;
            }
        }
        return n - k;
    }
};