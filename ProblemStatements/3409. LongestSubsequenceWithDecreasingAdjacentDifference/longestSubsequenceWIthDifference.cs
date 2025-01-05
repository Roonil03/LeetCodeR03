public class Solution {
    int n;
    int[] a = new int[10005];
    int[] lst = new int[305];
    int[,] dp = new int[10005, 305];
    int[,] pd = new int[10005, 305];
    public int LongestSubsequence(int[] nums) {
        n = nums.Length;
        int ans = 0;        
        for(int i = 0; i < n; ++i) a[i + 1] = nums[i];        
        for(int i = 1; i <= n; ++i) {
            dp[i, 300] = 1;
            for(int j = 0; j <= 300; ++j) {
                if(lst[j] == 0) continue;
                int now = Math.Abs(j - a[i]);
                dp[i, now] = Math.Max(dp[i, now], pd[lst[j], now] + 1);
            }
            pd[i, 300] = dp[i, 300];
            for(int j = 299; j >= 0; --j) pd[i, j] = Math.Max(dp[i, j], pd[i, j + 1]);
            lst[a[i]] = i;
        }        
        for(int i = 1; i <= n; ++i) {
            for(int j = 0; j <= 300; ++j) {
                ans = Math.Max(ans, dp[i, j]);
            }
        }        
        return ans;
    }
}