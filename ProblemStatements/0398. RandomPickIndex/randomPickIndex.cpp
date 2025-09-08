class Solution {
public:

    vector<int> nums;
    random_device rd;
    mutable mt19937 gen;

    Solution(vector<int>& nums) : nums(nums), gen(rd()) {
        
    }
    
    int pick(int target) {
        int count {0}, res {-1};
        for(int i {0}; i < (int)nums.size(); i++){
            if(nums[i] == target){
                count++;
                uniform_int_distribution<> dis(1, count);
                if(dis(gen) == 1){
                    res = i;
                }
            }
        }
        return res;
    }
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(nums);
 * int param_1 = obj->pick(target);
 */