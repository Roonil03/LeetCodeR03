// Could not solve during the time period of the contest
// Has been solved later

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

// The provided code defines a Solution class with a method LongestSubsequence that determines the length of the longest 
// subsequence in an array nums, where the absolute difference between consecutive elements is at most 300. The solution 
// employs dynamic programming, initializing arrays a, lst, dp, and pd to store the input elements, last occurrence indices, 
// and DP states, respectively. It iterates through the input array to update these arrays, using cumulative maximum values 
// for efficient look-up. Finally, it calculates the result by finding the maximum value in the dp array. The time complexity 
// is (O(n*301)) and the space complexity is also (O(n*301))