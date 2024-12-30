class Solution:
    def medianOfUniquenessArray(self, nums: List[int]) -> int:
        n = len(nums)
        total = n * (n + 1) // 2
        def most(k):
            res = 0
            cnt = Counter()
            i = 0
            for j in range(n):
                cnt[nums[j]] += 1
                while len(cnt) > k:
                    cnt[nums[i]] -= 1
                    if cnt[nums[i]] == 0:
                        del cnt[nums[i]]
                    i += 1
                res += j - i + 1
            return res
        l, r = 1, len(set(nums))
        while l < r:
            k = (l + r) // 2
            if most(k)*2 >= total:
                r = k
            else:
                l = k + 1
        return l