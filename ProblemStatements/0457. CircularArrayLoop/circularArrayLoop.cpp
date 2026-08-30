class Solution {
public:
    bool circularArrayLoop(vector<int>& nums) {
        int n = nums.size();
        auto nn = [&](int i){
            return ((i + nums[i]) % n + n) % n;
        };
        for(int i {0}; i < n; i++){
            if(nums[i] == 0){
                continue;
            }
            int f = i, s = i;
            bool fg = nums[i] > 0;
            auto f1 = [&](int id){
                return nums[id] != 0 && (nums[id] > 0) == fg;
            };
            while(f1(s) && f1(f) && f1(nn(f))){
                s = nn(s);
                f = nn(nn(f));
                if (f == s){
                    if (s == nn(s)){
                        break;
                    }
                    return true;
                }
            }
            int cur = i;
            while(f1(cur)){
                int j = nn(cur);
                nums[cur] = 0;
                cur = j;
            }
        }
        return false;
    }
};