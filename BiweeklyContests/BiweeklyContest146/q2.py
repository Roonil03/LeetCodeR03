MOD = 10**9 + 7

class Solution:
    def countPathsWithXorValue(self, grid, k) ->int:
        m, n = len(grid), len(grid[0])        
        dp = [[{} for _ in range(n)] for _ in range(m)]
        dp[0][0][grid[0][0]] = 1        
        for i in range(m):
            for j in range(n):
                for x in list(dp[i][j].keys()):
                    if j + 1 < n:
                        nx = x ^ grid[i][j + 1]
                        dp[i][j + 1][nx] = (dp[i][j + 1].get(nx, 0) + dp[i][j].get(x, 0)) % MOD
                    if i + 1 < m:
                        nx = x ^ grid[i + 1][j]
                        dp[i + 1][j][nx] = (dp[i + 1][j].get(nx, 0) + dp[i][j].get(x, 0)) % MOD        
        return dp[m - 1][n - 1].get(k, 0)