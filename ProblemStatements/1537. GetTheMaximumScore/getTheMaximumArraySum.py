class Solution:
    def maxSum(self, nums1: List[int], nums2: List[int]) -> int:
        mod = 10**9 + 7
        a ,b ,i ,j =0, 0, 0, 0
        while i < len(nums1) or j < len(nums2):
            if i < len(nums1) and (j == len(nums2) or nums1[i] < nums2[j]):
                a += nums1[i]
                i +=1
            elif j < len(nums2) and (i == len(nums1) or nums1[i] > nums2[j]):
                b += nums2[j]
                j += 1
            else:
                a = b = max(a, b) + nums1[i]
                i +=1
                j +=1
        return max(a, b) % mod