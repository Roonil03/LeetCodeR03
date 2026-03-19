class Solution {
public:
    int maxNumberOfFamilies(int n, vector<vector<int>>& reservedSeats) {
        unordered_map<int, int> mp;
        for(auto& i : reservedSeats){
            int r = i[0], s = i[1];
            if(s >= 2 && s <= 9){
                mp[r] |= (1 << s);
            }
        }
        int res {0};
        for(auto& [r, m] : mp){
            bool l = !(m & ((1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)));
            bool mid = !(m & ((1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)));
            bool rt = !(m & ((1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)));
            if(l && rt){
                res += 2;
            } else if(l || mid || rt){
                res++;
            }
        }
        res += (n - mp.size()) * 2;
        return res;
    }
};