using System;
using System.Collections.Generic;

public class Solution {
    public int SlidingPuzzle(int[][] board) {
        string target = "123450";
        string start = string.Join("", board[0]) + string.Join("", board[1]);
        int[][] dirs = new int[][] {
            new int[] {1, 3}, new int[] {0, 2, 4}, new int[] {1, 5},
            new int[] {0, 4}, new int[] {1, 3, 5}, new int[] {2, 4}
        };
        Queue<string> q = new Queue<string>();
        HashSet<string> seen = new HashSet<string>();
        q.Enqueue(start);
        seen.Add(start);
        int steps = 0;
        while (q.Count > 0) {
            int n = q.Count;
            for (int i = 0; i < n; i++) {
                string cur = q.Dequeue();
                if (cur == target) return steps;
                int idx = cur.IndexOf('0');
                foreach (int d in dirs[idx]) {
                    char[] arr = cur.ToCharArray();
                    (arr[idx], arr[d]) = (arr[d], arr[idx]);
                    string next = new string(arr);
                    if (!seen.Contains(next)) {
                        q.Enqueue(next);
                        seen.Add(next);
                    }
                }
            }
            steps++;
        }
        return -1;
    }
}
