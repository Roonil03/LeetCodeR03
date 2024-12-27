public class Solution {
    public int MaxScoreSightseeingPair(int[] values) {
        int ans = 0;
        int prev = values[0];
        for(int i = 1; i< values.Length; i++){
            ans = Math.Max(ans, prev + values[i] - i);
            prev = Math.Max(prev, values[i]+i);
        }
        return ans;
    }
}