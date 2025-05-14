class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int n = nums.size();
        if (n == 0){
            return 0;
        }
        int a {nums[0]}, b {nums[0]}, res {nums[0]};
        for (int i {1}; i < n; i++){
            int cur = nums[i];
            if(cur < 0){
                swap(a, b);
            }
            a = max(cur, a * cur);
            b = min(cur, b * cur);
            res = max(res, a);
        }
        return res;
    }
};