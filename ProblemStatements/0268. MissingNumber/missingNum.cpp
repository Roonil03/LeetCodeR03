class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int n = nums.size(), s= n * (n + 1) / 2, t {0};
        for(int v : nums){
            t += v;
        }
        return s - t;
    }
};