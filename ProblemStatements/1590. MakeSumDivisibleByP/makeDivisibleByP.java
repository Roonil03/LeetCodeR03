import java.util.*;

class Solution {
    public int minSubarray(int[] nums, int p) {
        long totalSum = 0;
        for (int num : nums) {
            totalSum += num;
        }
        int rem = (int) (totalSum % p);
        if (rem == 0) {
            return 0;
        }
        HashMap<Integer, Integer> pre = new HashMap<>();
        pre.put(0, -1);
        long preSum = 0;
        int minLength = nums.length;
        for (int i = 0; i < nums.length; i++) {
            preSum += nums[i];
            int currentMod = (int) (preSum % p);
            int target = (currentMod - rem + p) % p;
            if (pre.containsKey(target)) {
                minLength = Math.min(minLength, i - pre.get(target));
            }
            pre.put(currentMod, i);
        }
        return minLength == nums.length ? -1 : minLength;
    }
}