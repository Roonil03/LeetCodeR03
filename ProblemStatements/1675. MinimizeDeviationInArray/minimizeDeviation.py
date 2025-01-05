class Solution:
    def minimumDeviation(self, nums: List[int]) -> int:
        nums = [-num*2 if num % 2 else -num for num in nums]
        heapq.heapify(nums)
        mn = -max(nums)
        md = float('inf')        
        while True:
            maxN = -heapq.heappop(nums)
            md = min(md, maxN - mn)
            if maxN % 2:
                break
            mn = min(mn, maxN // 2)
            heapq.heappush(nums, -maxN // 2)
        
        return md