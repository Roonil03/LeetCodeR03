public class Solution {
    public int MincostTickets(int[] days, int[] costs) {
        HashSet<int> hs = new HashSet<int>(days);
        int last = days[days.Length-1];
        int [] dp = new int[last+1];
        for(int i = 1; i<= last; i++){
            if(!hs.Contains(i)){
                dp[i] = dp[i-1];
                continue;
            }
            int d1 = dp[i-1] + costs[0];
            int d7 = dp[Math.Max(0, i - 7)] + costs[1];
            int d30 = dp[Math.Max(0, i - 30)] + costs[2];
            dp[i] = Math.Min(d1, Math.Min(d7, d30));
        }
        return dp[last];
    }
}