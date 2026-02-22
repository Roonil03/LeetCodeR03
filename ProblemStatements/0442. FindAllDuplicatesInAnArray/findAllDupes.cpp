class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int> res;
        for(int i {0}; i < nums.size(); i++){
            int v = abs(nums[i]);
            int id = v - 1;
            if(nums[id] < 0){
                res.push_back(v);
            } else{
                nums[id] = -nums[id];
            }
        }
        return res;
    }
};