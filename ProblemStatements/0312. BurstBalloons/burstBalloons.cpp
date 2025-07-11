class Solution {
public:
    int maxCoins(vector<int>& nums) {
        int n = nums.size();
        vector<int>arr(n + 2,  1);
        for(int i {0}; i < n; i++){
            arr[i+1] = nums[i];
        }
        vector<vector<int>> dp(n + 2, vector<int>(n + 2, 0));
        for(int i {2}; i <= n + 1; i++){
            for(int l {0}; l + i <= n+1; l++){
                int r = l + i;
                for(int k = l +1 ; k < r; k++){
                    dp[l][r] = max(dp[l][r], dp[l][k] + dp[k][r] + arr[l] * arr[k] * arr[r]);
                }
            }
        }
        return dp[0][n+1];
    }
};