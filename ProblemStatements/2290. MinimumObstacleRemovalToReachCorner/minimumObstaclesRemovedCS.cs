using System;
using System.Collections.Generic;

public class Solution {    public int MinimumObstacles(int[][] grid) {
        int m = grid.Length, n = grid[0].Length;
        int[][] directions = new int[][] {
            new int[] {0, 1}, new int[] {1, 0}, new int[] {0, -1}, new int[] {-1, 0}
        };
        int[,] cost = new int[m, n];
        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++)
                cost[i, j] = int.MaxValue;
        var deque = new LinkedList<(int x, int y, int c)>();
        deque.AddFirst((0, 0, 0));
        cost[0, 0] = 0;
        while (deque.Count > 0) {
            var (x, y, currentCost) = deque.First.Value;
            deque.RemoveFirst();
            if (x == m - 1 && y == n - 1)
                return currentCost;
            foreach (var dir in directions) {
                int nx = x + dir[0], ny = y + dir[1];
                if (nx >= 0 && nx < m && ny >= 0 && ny < n) {
                    int newCost = currentCost + grid[nx][ny];
                    if (newCost < cost[nx, ny]) {
                        cost[nx, ny] = newCost;
                        if (grid[nx][ny] == 0)
                            deque.AddFirst((nx, ny, newCost)); 
                        else
                            deque.AddLast((nx, ny, newCost));
                    }
                }
            }
        }
        return -1;
    }
}
