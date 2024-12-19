# 3388. Count Beautiful Splits in an Array
## Question Level: Medium
### Description:
You are given an array `nums`.

A split of an array `nums` is beautiful if:
- The array `nums` is split into three non-empty subarrays: `nums1`, `nums2`, and `nums3`, such that nums can be formed by concatenating `nums1`, `nums2`, and `nums3` in that order.
- The subarray nums1 is a prefix of `nums2` OR `nums2` is a prefix of `nums3`.

Return the number of ways you can make this split.

- A subarray is a contiguous non-empty sequence of elements within an array.
- A prefix of an array is a subarray that starts from the beginning of the array and extends to any point within it.

### Test Cases:
#### Example 1:
Input: nums = `[1,1,2,1]`
Output: 2
Explanation:
The beautiful splits are:
- A split with `nums1` = `[1]`, `nums2` = `[1,2]`, `nums3` = `[1]`.
- A split with `nums1` = `[1]`, `nums2` = `[1]`, `nums3` = `[2,1]`.

#### Example 2:
Input: nums = `[1,2,3,4]`
Output: 0
Explanation:
- There are 0 beautiful splits.

### Constraints:
- 1 <= `nums.length` <= 5000
- 0 <= `nums[i]` <= 50

### <i>Weekly Contest Question </i>