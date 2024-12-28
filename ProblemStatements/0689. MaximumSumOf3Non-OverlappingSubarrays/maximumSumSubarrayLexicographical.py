class Solution:
    def maxSumOfThreeSubarrays(self, nums: List[int], k: int) -> List[int]:
        n = len(nums)
        prefixSum = [0] * (n + 1)
        for i in range(n):
            prefixSum[i + 1] = prefixSum[i] + nums[i]        
        leftMaxIndex = [0] * n
        currentMaxSum = prefixSum[k] - prefixSum[0]
        for i in range(k, n):
            sumSubarray = prefixSum[i + 1] - prefixSum[i + 1 - k]
            if sumSubarray > currentMaxSum:
                leftMaxIndex[i] = i + 1 - k
                currentMaxSum = sumSubarray
            else:
                leftMaxIndex[i] = leftMaxIndex[i - 1]        
        rightMaxIndex = [0] * n
        rightMaxIndex[n - k] = n - k
        currentMaxSum = prefixSum[n] - prefixSum[n - k]
        for i in range(n - k - 1, -1, -1):
            sumSubarray = prefixSum[i + k] - prefixSum[i]
            if sumSubarray >= currentMaxSum:
                rightMaxIndex[i] = i
                currentMaxSum = sumSubarray
            else:
                rightMaxIndex[i] = rightMaxIndex[i + 1]        
        ans = []
        maxSum = 0
        for i in range(k, n - 2 * k + 1):
            left = leftMaxIndex[i - 1]
            right = rightMaxIndex[i + k]
            totalSum = (prefixSum[i + k] - prefixSum[i] + prefixSum[left + k] - prefixSum[left] + prefixSum[right + k] - prefixSum[right])
            if totalSum > maxSum:
                maxSum = totalSum
                ans = [left, i, right]        
        return ans
