class Solution {
public:
    int maxRotateFunction(vector<int>& nums) {
        int n = nums.size();
        long long sum {0}, f{0};
        for(int i{0}; i < n; i++){
            sum += nums[i];
            f += (long long)i * nums[i];
        }
        long long res = f;
        for(int k = n - 1; k > 0; k--){
            f += sum - (long long)n * nums[k];
            if(f > res){
                res = f;
            }
        }
        return (int)res;
    }
};