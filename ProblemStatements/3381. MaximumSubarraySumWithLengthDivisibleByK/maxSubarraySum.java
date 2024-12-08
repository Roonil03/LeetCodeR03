import java.util.*;

class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        int n = nums.length;
        long[] p = new long[n + 1];
        for (int i = 0; i < n; ++i) {
            p[i + 1] = p[i] + nums[i];
        }
        long[] m = new long[k];
        Arrays.fill(m, Long.MAX_VALUE);
        long s = Long.MIN_VALUE;
        for (int i = 0; i <= n; ++i) {
            int mod = i % k;
            if (m[mod] != Long.MAX_VALUE) {
                long c = p[i] - m[mod];
                s = Math.max(s, c);
            }
            m[mod] = Math.min(m[mod], p[i]);
        }
        return s;
    }
}
