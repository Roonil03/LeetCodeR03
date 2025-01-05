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