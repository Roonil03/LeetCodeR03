class Solution {
public:
    int longestPalindromeSubseq(string s) {
        int n = s.size();
        vector<vector<int>> dp(n, vector<int>(n, 0));
        for(int i{0}; i < n;i++){
            dp[i][i] = 1;
        }
        for(int j{2}; j <= n; j++){
            for(int i{0}; i <= n - j; i++){
                int a = i + j - 1;
                if(s[i] == s[a]){
                    dp[i][a] = 2 + (i + 1 <= a - 1 ? dp[i+1][a-1] : 0);
                } else{
                    dp[i][a] = max(dp[i+1][a], dp[i][a-1]);
                }
            }
        }
        return dp[0][n-1];
    }
};