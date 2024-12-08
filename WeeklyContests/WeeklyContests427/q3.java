/* Could not solve it entirely during the time of the contest, 
 * it has been solved post the time limit of the contest.
 */
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
/*
The solution to the problem of finding the maximum sum of a subarray whose size is divisible by k uses the concepts of prefix sums and modular arithmetic. Here's an explanation of how the algorithm works:
Prefix Sum:
The prefix sum technique is used to efficiently calculate the sum of any subarray. The prefix sum array p stores the sum of elements from the start of the array up to each index. This way, the sum of any subarray nums[i...j] can be calculated as p[j+1] - p[i] in constant time, which avoids recalculating sums repeatedly.

Modulo Operation:
The key idea is that we need to find subarrays whose size is divisible by k. A subarray's size is divisible by k if the difference in its starting and ending indices, j - i + 1, is divisible by k. This is equivalent to ensuring that the prefix sum indices i and j belong to the same modulo class when divided by k. In other words, we look for subarrays where both i % k == j % k.

Tracking Minimum Prefix Sums:
We maintain an array m of size k where each element corresponds to the smallest prefix sum encountered for a specific modulo class. This allows us to efficiently compute the maximum sum of subarrays whose size is divisible by k without iterating over all possible subarrays.

Algorithm:
First, we compute the prefix sum for each element in the array.
We then initialize the m array with Long.MAX_VALUE to track the smallest prefix sum for each modulo class.
As we iterate through the prefix sum array, for each index i, we check if the minimum prefix sum for the current modulo class has been recorded. If it has, we calculate the sum of the subarray from the previous occurrence of that modulo class to the current index.
The sum of the subarray is the difference between the current prefix sum and the smallest recorded prefix sum for the modulo class. If this sum is greater than the current maximum, we update the result.
Finally, we update the smallest prefix sum for the current modulo class if the current prefix sum is smaller.

Efficiency:
The time complexity of the solution is O(n), where n is the length of the array. This is because we only need to compute the prefix sums once and iterate over the array a single time. The use of modular arithmetic and dynamic tracking of minimum prefix sums avoids the need for nested loops, making the solution much more efficient compared to brute force approaches.

In summary, this solution leverages prefix sums to compute subarray sums in constant time and uses modular arithmetic to ensure we only consider subarrays whose size is divisible by k. The dynamic tracking of minimum prefix sums allows us to efficiently calculate the maximum sum of such subarrays.
 */