class Solution:
    def numWays(self, words: List[str], target: str) -> int:
        MOD = 10**9 + 7
        m, n = len(words), len(words[0])
        dp = [0] * (len(target) + 1)
        dp[0] = 1        
        for k in range(n):
            count = [0] * 26
            for word in words:
                count[ord(word[k]) - ord('a')] += 1
            for i in range(len(target) - 1, -1, -1):
                dp[i + 1] = (dp[i + 1] + dp[i] * count[ord(target[i]) - ord('a')]) % MOD
        return dp[len(target)]