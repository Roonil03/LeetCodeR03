class Solution {
public:
    int minFallingPathSum(vector<vector<int>>& matrix) {
        int n = matrix.size();
        vector<vector<int>> dp(n, vector<int>(n, 0));
        for (int j {0}; j < n; j++){
            dp[n-1][j] = matrix[n-1][j];
        }
        for (int i = n-2; i >= 0; --i) {
            for (int j {0}; j < n; ++j) {
                int mn = dp[i+1][j];
                if (j > 0) mn = min(mn, dp[i+1][j-1]);
                if (j < n-1) mn = min(mn, dp[i+1][j+1]);
                dp[i][j] = matrix[i][j] + mn;
            }
        }
        return *min_element(dp[0].begin(), dp[0].end());
    }
};