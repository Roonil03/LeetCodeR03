class Solution {
public:
    int maxProfit(int k, vector<int>& prices) {
        int n = prices.size();
        if (n == 0 || k == 0){
            return 0;
        }
        if (k >= n / 2) {
            int res = 0;
            for (int i = 1; i < n; ++i){
                res += max(0, prices[i] - prices[i - 1]);
            }
            return res;
        }
        vector<vector<int>> dp(k + 1, vector<int>(n, 0));
        for (int t = 1; t <= k; ++t) {
            int a = -prices[0];
            for (int d = 1; d < n; ++d) {
                dp[t][d] = max(dp[t][d - 1], prices[d] + a);
                a = max(a, dp[t - 1][d] - prices[d]);
            }
        }
        return dp[k][n - 1];
    }
};