class Solution {
    public:
        int maxFrequency(vector<int>& nums, int k) {
            sort(begin(nums), end(nums));
            long i {0}, j {0}, n = nums.size(), ans {1}, sum {nums[0]};
            for(; i < n;i++){
                while(j < n && (j-i+1) * nums[j]-sum <= k){
                    ans = max(ans, j-i+1);
                    j++;
                    if(j < n){
                        sum += nums[j];
                    }
                }
                sum -= nums[i];
            }
            return ans;
        }
    };