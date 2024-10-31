import sys

class Solution(object):
   def minimumTotalDistance(self, robot, factory):
    robot.sort()
    factory.sort()
    n = len(robot)
    m = len(factory)
    dp = [[sys.maxsize] * (m + 1) for _ in range(n + 1)]
    dp[0][0] = 0
    for j in range(1, m + 1):
        posj, limitj = factory[j - 1]
        dp[0][j] = 0 
        for i in range(1, n+1):
            dp[i][j] = dp[i][j-1]
            sum = 0
            for k in range(1, min(i, limitj) + 1):
                sum += abs(robot[i - k] - posj)
                dp[i][j] = min(dp[i][j], dp[i-k][j-1] + sum)                
    return dp[n][m]

        