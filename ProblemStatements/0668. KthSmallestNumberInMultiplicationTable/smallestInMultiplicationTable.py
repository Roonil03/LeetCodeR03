class Solution:
    def findKthNumber(self, m: int, n: int, k: int) -> int:
        def count(x):
            c = 0
            for i in range (1, m+1):
                c += min(x // i,n)
            return c
        l, r = 1, m*n
        while(l < r):
            mid = (l + r) // 2
            if(count(mid) < k):
                l = mid + 1
            else:
                r = mid
        return l