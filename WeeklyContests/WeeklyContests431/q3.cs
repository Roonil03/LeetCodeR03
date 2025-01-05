/* Could not solve it entirely during the time of the contest,
 * it has been solved post the time limit of the contest.
 */

 public class Solution {
    public long MaximumCoins(int[][] coins, int k) {
        Array.Sort(coins, (a, b) => a[0].CompareTo(b[0]));
        int n = coins.Length;
        int[] starts = new int[n];
        int[] ends = new int[n];
        int[] values = new int[n];
        for (int i = 0; i < n; i++) {
            starts[i] = coins[i][0];
            ends[i] = coins[i][1];
            values[i] = coins[i][2];
        }
        long[] prefixSum = new long[n];
        prefixSum[0] = (long)(ends[0] - starts[0] + 1) * values[0];
        for (int i = 1; i < n; i++) {
            prefixSum[i] = prefixSum[i-1] + (long)(ends[i] - starts[i] + 1) * values[i];
        }
        HashSet<int> candidateL = new HashSet<int>();
        for (int i = 0; i < n; i++) {
            candidateL.Add(starts[i]);
            int adjL = ends[i] - k + 1;
            if (adjL >= 1) {
                candidateL.Add(adjL);
            }
        }
        var sortedL = candidateL.ToList();
        sortedL.Sort();
        long maxCoins = 0;
        foreach (var L in sortedL) {
            int R = L + k - 1;
            int leftIdx = Array.BinarySearch(ends, L);
            if (leftIdx < 0) leftIdx = ~leftIdx;
            int rightIdx = Array.BinarySearch(starts, R);
            if (rightIdx < 0) rightIdx = ~rightIdx - 1;
            if (leftIdx > rightIdx || leftIdx == n) {
                continue;
            }
            long total = prefixSum[rightIdx];
            if (leftIdx > 0) {
                total -= prefixSum[leftIdx - 1];
            }
            if (starts[leftIdx] < L) {
                long overlapLen = L - starts[leftIdx];
                total -= overlapLen * values[leftIdx];
            }
            if (ends[rightIdx] > R) {
                long overlapLen = ends[rightIdx] - R;
                total -= overlapLen * values[rightIdx];
            }
            if (total > maxCoins) {
                maxCoins = total;
            }
        }
        return maxCoins;
    }
}

//  The provided code defines a Solution class with a method MaximumCoins that calculates the 
//  maximum number of coins that can be collected from a set of intervals within a given length
//  k. The method first sorts the intervals by their starting points and then extracts the 
//  start, end, and value of each interval into separate arrays. It computes prefix sums for 
//  the interval values to facilitate quick range sum calculations. The method then generates 
//  candidate starting points for the collection interval, considering both the original start 
//  points and adjusted points to ensure all possible intervals of length k are considered. 
//  For each candidate start point L, it calculates the corresponding end point R and determines 
//  the total value of coins that can be collected within the interval [L, R] using binary 
//  search on the start and end points and adjusting for overlaps. The maximum collected coins 
//  value is updated accordingly. The time complexity is (O(nlog n)) due to sorting and binary 
//  search operations, and the space complexity is (O(n)) for storing interval data and prefix 
//  sums.