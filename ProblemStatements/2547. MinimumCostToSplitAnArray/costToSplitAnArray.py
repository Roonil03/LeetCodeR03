class Solution:
    def minCost(self, nums: List[int], k: int) -> int:
        dp = [0]
        nums = [-1] + nums
        def ra(i):
            m = defaultdict(int)
            s = 0
            res = 10000000000000
            while i > 0:
                if m[nums[i]] > 0:
                    s += 1
                    if m[nums[i]] == 1:
                        s +=1
                m[nums[i]] += 1
                res = min(dp[i-1]+s, res)
                i -=1
            return res+k
        for i in range(1, len(nums)):
            dp.append(ra(i))
        return dp[-1]