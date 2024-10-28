import java.util.*;

class Solution {
    public int longestSquareStreak(int[] nums) {
        if (nums == null || nums.length < 2) return -1;
        Arrays.sort(nums);
        HashSet<Integer> numSet = new HashSet<>();
        for (int num : nums) {
           // if (num > 0) {
                numSet.add(num);
            //}
        }

        int longestStreak = -1;
        for (int num : nums) {
            if (num <= 0) continue;

            int current = num;
            int streakLength = 1;

            while (current <= Math.sqrt(Integer.MAX_VALUE) && numSet.contains(current * current)) {
                current *= current;
                streakLength++;
            }
            if (streakLength >= 2) {
                longestStreak = Math.max(longestStreak, streakLength);
            }
        }

        return longestStreak;
    }
}
