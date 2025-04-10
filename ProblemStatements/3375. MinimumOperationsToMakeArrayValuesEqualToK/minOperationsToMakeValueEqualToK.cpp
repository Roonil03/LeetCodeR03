class Solution {
    public:
        int minOperations(vector<int>& nums, int k) {
            // sort(nums.begin(), nums.end());
            // int min1 {nums[0]};
            // if(min1 < k){
            //     return -1;
            // }
            // int res {0};
            // for(int i {0}; i < nums.size(); i++){
            //     if (nums[i] > k){
            //         res++;
            //         while(i+1 < nums.size() && nums[i] == nums[i+1]){
            //             i++;
            //         }
            //     }
            // }
            // return res;
            unordered_map<int, int>m;
            for(int i : nums){
                if(i < k){
                    return -1;
                } else if(i > k){
                    m[i]++;
                }
            }
            return m.size();
        }
    };