class Solution:
    # def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     res = 0
    #     for i in range(n):
    #         c = nums[i]
    #         op = 0
    #         for j in range(i, n):
    #             a = max(0, c - nums[j])
    #             op += a
    #             if op <= k:
    #                 res += 1
    #                 c = max(c, nums[j] + a) 
    #             else:
    #                 break       
    #     return res
# class Solution:
    def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
        from itertools import accumulate        
        tmp = list(accumulate(nums, initial=0))
        n = len(nums)
        stk = [n]
        acc = [0]        
        def f(idx):
            l, r = 0, len(stk) - 1
            while l <= r:
                m = (l + r) // 2
                if stk[m] > idx: l = m + 1
                else: r = m - 1
            return acc[-1] - acc[l] + nums[stk[l]] * (idx + 1 - stk[l]) - (tmp[idx + 1] - tmp[i])        
        ans = 0
        for i in range(n - 1, -1, -1):
            while stk[-1] != n and nums[stk[-1]] <= nums[i]:
                acc.pop()
                stk.pop()
            acc.append((stk[-1] - i) * nums[i] + acc[-1])
            stk.append(i)            
            l = i
            r = n - 1
            while l <= r:
                m = (l + r) // 2
                if f(m) > k: r = m - 1
                else: l = m + 1            
            ans += r - i + 1
        return ans