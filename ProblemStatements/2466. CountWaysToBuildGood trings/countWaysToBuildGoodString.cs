public class Solution {
    public int CountGoodStrings(int low, int high, int zero, int one) {
        int mod = 1000000007;
        int[] dp = new int[high + 1];
        dp[0] = 1;
        for (int i = 1; i <= high; i++) {
            if (i >= zero){
                dp[i] = (dp[i] + dp[i - zero]) % mod;
            }
            if (i >= one){
                dp[i] = (dp[i] + dp[i - one]) % mod;
            }
        }
        int res = 0;
        for (int i = low; i <= high; i++) {
            res = (res + dp[i]) % mod;
        }
        return res;
    }
}