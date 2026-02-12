class Solution {
public:
    int numberOfArithmeticSlices(vector<int>& nums) {
        int n = nums.size();
        int res {0}, cur {0};;
        for(int i {2} ; i < n; i++){
            if(nums[i - 1] - nums[i-2] == nums[i] - nums[i-1]){
                cur++;
                res += cur;
            } else{
                cur = 0;
            }
        }
        return res;
    }
};