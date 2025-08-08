class Solution {
public:
    int minPatches(vector<int>& nums, int n) {
        long long miss {1};
        int i {0}, res {0};
        while(miss <= n){
            if(i < nums.size() && nums[i] <= miss){
                miss += nums[i++];
            } else{
                miss += miss;
                res++;
            }
        }
        return res;
    }
};