MOD = 10**9+7
comb2 = lambda x: x * (x-1) // 2

class Solution:
    def subsequencesWithMiddleMode(self, nums: List[int]) -> int:
        mp = {x: i for i, x in enumerate(set(nums))}
        nums = [mp[x] for x in nums]
        n = len(nums)
        pres = [[0] * len(mp)]
        for x in nums:
            pres.append(pres[-1].copy())
            pres[-1][x] += 1
        ans = 0
        for i in range(2, n-2):
            c = nums[i]
            l = pres[i][c]
            r = pres[n][c] - pres[i+1][c]
            res = comb2(i) * comb2(n-1-i) - comb2(i-l) * comb2(n-1-i-r)
            for y in range(len(mp)):
                if y == c: continue
                ly = pres[i][y]
                ry = pres[n][y] - pres[i+1][y]
                lz = i - l - ly
                rz = n-1-i - r - ry
                res -= l*ly*ry*rz + lz*ly*r*ry + l*(i-l)*comb2(ry) + r*(n-1-i-r)*comb2(ly)
            ans = (ans + res) % MOD
        return ans