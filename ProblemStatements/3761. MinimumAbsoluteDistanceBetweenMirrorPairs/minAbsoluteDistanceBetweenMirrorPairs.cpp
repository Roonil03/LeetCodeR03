class Solution {
    using ll = long long;

public:
    int minMirrorPairDistance(vector<int>& nums) {
        int n = nums.size();
        int res = INT_MAX;
        unordered_map<int, int> mp;
        for(int i {0}; i < n; i++){
            if(mp.count(nums[i])){
                res = min(res, i - mp[nums[i]]);
            }
            int r = h1(nums[i]);
            mp[r] = i;
        }
        return res == INT_MAX ? -1 : res;
    }

    int h1(int num){
        int res {0};
        while(num > 0){
            res = res * 10 + num % 10;
            num /= 10;
        }
        return res;
    }
};