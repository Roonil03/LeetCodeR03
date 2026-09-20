class Solution {
public:
    int maxDistance(vector<vector<int>>& arrays) {
        int res {0};
        int mn = arrays[0].front(), nm = arrays[0].back();
        for(int i {1}; i < arrays.size(); i++){
            res = max({res, arrays[i].back() - mn, nm - arrays[i].front()});
            mn = min(mn, arrays[i].front());
            nm = max(nm, arrays[i].back());
        }
        return res;
    }
};