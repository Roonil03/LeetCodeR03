# Could not solve it entirely during the time of the contest, 
# it has been solved post the time limit of the contest.


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


# 1. Preprocessing:
#   The code creates a 2D array pres to store the number of pairs of equal elements within subarrays.
#   pres[i][j] represents the number of pairs of equal elements in the subarray nums[i:j+1].
#   This preprocessing step efficiently calculates the number of pairs for all possible subarrays.

# 2. Counting Beautiful Splits:
#   The code iterates through all possible split points i (1 <= i <= n-2).
#   For each split point i, it iterates through all possible endpoints j (i+1 <= j <= n-1) of the second subarray.
#   For each pair of i and j, it checks if the split is "beautiful" based on the following conditions:
#       The length of the first subarray (len1) is less than or equal to the length of the second subarray (len2) and the number of pairs within the first subarray (pres[0][i]) is greater than or equal to len1.
#       The length of the second subarray (len2) is less than or equal to the length of the third subarray (len3) and the number of pairs within the second subarray (pres[i][j]) is greater than or equal to len2.
#   If the split is "beautiful," the count of beautiful splits (ans) is incremented.

# 3. Return Result:
#   The code returns the total count of beautiful splits (ans).

# In essence, the algorithm precomputes the number of pairs within subarrays and then efficiently iterates through all possible split points to count the number of "beautiful" splits that satisfy the given conditions.