class Solution:
    def beautifulSplits(self, nums: list[int]) -> int:
        n = len(nums)
        pres = [[0] * (n + 1) for a in range(n + 1)]        
        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                pres[i][j] = pres[i + 1][j + 1] + 1 if nums[i] == nums[j] else 0
        ans = 0
        for i in range(1, n - 1):
            for j in range(i + 1, n):
                len1 = i
                len2 = j - i
                len3 = n - j
                ok = 0
                if len1 <= len2 and pres[0][i] >= len1:
                    ok = 1
                elif len2 <= len3 and pres[i][j] >= len2:
                    ok = 1
                ans += ok
        return ans
