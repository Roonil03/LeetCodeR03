class Solution {
public:
    int getMoneyAmount(int n) {
        vector<vector<int>> dp(n + 1, vector<int>(n + 1, 0));
        for(int l {2}; l <= n; l++){
            for(int s {1}; s <= n - l + 1; s++){
                int e = s + l - 1;
                dp[s][e] = INT_MAX;
                for(int p = s; p < e; p++){
                    int c = p + max(dp[s][p - 1], dp[p + 1][e]);
                    dp[s][e] = min(dp[s][e], c);
                }
            }
        }
        return dp[1][n];
    }
};