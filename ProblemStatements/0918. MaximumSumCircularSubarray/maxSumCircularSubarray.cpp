class Solution {
public:
    int maxSubarraySumCircular(vector<int>& nums) {
        int n = nums.size(), tot {0}, a {nums[0]}, b {0}, c {nums[0]}, d {0};
        for (int n : nums){
            b = max(b + n, n);
            a = max(a, b);
            d = min(d + n, n);
            c = min(d, c);
            tot += n;
        }
        if (a <0){
            return a;
        }
        return max(a, tot - c);
    }
};