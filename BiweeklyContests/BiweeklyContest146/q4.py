# Could not solve it entirely during the time of the contest, 
# it has been solved post the time limit of the contest.

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
    

# The subsequencesWithMiddleMode function computes the number of subsequences of an array nums that have a specific "middle mode" 
# structure, where the middle element of the subsequence is the most frequent element in that subsequence. The program begins by 
# compressing the values of nums into indices via a mapping (mp), which simplifies subsequent operations. A prefix sum array pres 
# is built to keep track of the frequency of each element in nums up to each index. The function then iterates over possible middle 
# elements (from index 2 to n-2), where each element is treated as the "mode" of a subsequence. For each candidate middle element c, 
# the function calculates the number of valid subsequences based on the combinatorial selection of elements on both sides of c and 
# checks for overlaps. It calculates contributions for subsequences where other elements are the mode, adjusting the result for 
# overcounts. The result is accumulated in ans, which is returned after applying the modulo operation (MOD = 10^9 + 7) to handle 
# large numbers. This algorithm relies on prefix sum arrays for frequency counting, combinatorial counting functions (comb2), and a 
# greedy approach to efficiently count valid subsequences.

# Data structures used:
#   Dictionary (mp): Maps the unique elements of nums to consecutive indices for easier access.
#   List (pres): A 2D list to store the prefix frequency counts of each element at every index in nums.
#   Integer (ans): Holds the final result, the count of valid subsequences.