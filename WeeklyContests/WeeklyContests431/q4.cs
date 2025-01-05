/* Could not solve it entirely during the time of the contest,
 * it has been solved post the time limit of the contest.
 */

using System;
using System.Collections.Generic;
using System.Linq;

public class Solution {
    public int[] MaximumWeight(IList<IList<int>> intervals) {
        int[][] a = new int[intervals.Count][];
        for (int i = 0; i < a.Length; i++) {
            a[i] = new int[4];
            a[i][0] = intervals[i][0];
            a[i][1] = intervals[i][1];
            a[i][2] = intervals[i][2];
            a[i][3] = i;
        }
        Array.Sort(a, (x, y) => x[0] - y[0]);

        long?[][] dp = new long?[a.Length][];
        for (int i = 0; i < a.Length; i++) {
            dp[i] = new long?[4];
        }

        List<int> list = new List<int>();
        long best = Recur(0, 3, dp, a);

        int start = 0;
        int left = 3;
        long sub = 0L;
        int next = 0;
        for (int k = 0; k < 4; k++) {
            int cur = -1;
            for (int i = 0; i < a.Length; i++) {
                if (intervals[i][0] > start) {
                    int l = 0, r = a.Length - 1;
                    while (l < r) {
                        int m = (l + r) / 2;
                        if (a[m][0] <= intervals[i][1]) {
                            l = m + 1;
                        } else {
                            r = m;
                        }
                    }
                    if (l < a.Length && a[l][0] <= intervals[i][1]) l++;
                    if (best == (Recur(l, left - 1, dp, a) + intervals[i][2])) {
                        if (cur == -1 || i < cur) {
                            cur = i;
                            sub = intervals[i][2];
                            next = intervals[i][1];
                        }
                    }
                }
            }
            if (cur != -1) {
                list.Add(cur);
                start = next;
                best -= sub;
                left--;
            } else {
                break;
            }
            if (best == 0) break;
        }
        list.Sort();
        return list.ToArray();
    }

    public void Trace(List<int> arr, int ind, int left, long?[][] dp, int[][] a) {
        if (ind == a.Length) return;
        if (left == -1) return;
        long res = Recur(ind, left, dp, a);

        int l = ind + 1, r = a.Length - 1;
        while (l < r) {
            int m = (l + r) / 2;
            if (a[m][0] <= a[ind][1]) {
                l = m + 1;
            } else {
                r = m;
            }
        }
        if (l < a.Length && a[l][0] <= a[ind][1]) l++;
        if (res == Recur(l, left - 1, dp, a) + a[ind][2]) {
            arr.Add(a[ind][3]);
            Trace(arr, l, left - 1, dp, a);
        } else {
            Trace(arr, ind + 1, left, dp, a);
        }
    }

    public long Recur(int ind, int left, long?[][] dp, int[][] a) {
        if (left == -1) return 0;
        if (ind == a.Length) return 0;
        if (dp[ind][left] != null) return dp[ind][left].Value;

        long res = Math.Max(Recur(ind + 1, left, dp, a), Recur(ind + 1, left, dp, a));
        int l = ind + 1, r = a.Length - 1;
        while (l < r) {
            int m = (l + r) / 2;
            if (a[m][0] <= a[ind][1]) {
                l = m + 1;
            } else {
                r = m;
            }
        }
        if (l < a.Length && a[l][0] <= a[ind][1]) l++;
        res = Math.Max(res, Recur(l, left - 1, dp, a) + a[ind][2]);

        dp[ind][left] = res;
        return res;
    }
}

// The provided code defines a Solution class with a method MaximumWeight that calculates the maximum 
// weight of non-overlapping intervals from a given list of intervals, each represented by a start 
// time, end time, and weight. The method starts by sorting the intervals based on their start times 
// and initializing a dynamic programming (DP) table dp to store intermediate results. The Recur method 
// is used to compute the maximum weight using recursion and memoization, considering whether to 
// include or exclude each interval. The Trace method reconstructs the solution to find which intervals
// contribute to the maximum weight. The main MaximumWeight method iterates through potential starting 
// points and uses binary search to optimize the selection of intervals, updating the best weight and 
// tracking the selected intervals. The algorithm efficiently handles the interval selection and weight 
// calculation with a time complexity of (O(nlog n)) due to sorting and binary search operations, and a 
// space complexity of (O(n)) for the DP table and auxiliary data structures.