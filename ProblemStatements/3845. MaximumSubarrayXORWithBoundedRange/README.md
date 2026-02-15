# 3845. Maximum Subarray XOR with Bounded Range
## Question Level: Hard
### Description:
You are given a non-negative integer array nums and an integer k.

You must select a subarray of nums such that the difference between its maximum and minimum elements is at most k. The value of this subarray is the bitwise XOR of all elements in the subarray.

Return an integer denoting the maximum possible value of the selected subarray.

A subarray is a contiguous non-empty sequence of elements within an array.

### Examples:
#### Example 1:

Input: nums = `[5,4,5,6]`, k = 2

Output: 7

Explanation:
- Select the subarray [5, **4, 5, 6**].
- The difference between its maximum and minimum elements is 6 - 4 = 2 <= k.
- The value is 4 XOR 5 XOR 6 = 7.
#### Example 2:

Input: nums = `[5,4,5,6]`, k = 1

Output: 6

Explanation:
- Select the subarray [5, 4, 5, **6**].
- The difference between its maximum and minimum elements is 6 - 6 = 0 <= k.
- The value is 6.

### Constraints:

- 1 <= `nums.length` <= 4 * 10<sup>4</sup>
- 0 <= `nums[i]` < 2<sup>15</sup>
- 0 <= `k` < 2<sup>15</sup>

### *Weekly Contest Question*