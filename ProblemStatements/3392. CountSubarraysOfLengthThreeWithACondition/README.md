# 3392. Count Subarrays of Length Three With a Condition
## Question Level: Easy
### Description:
Given an integer array nums, return the number of subarrays of length 3 such that the sum of the first and third numbers equals exactly half of the second number.

A subarray is a contiguous non-empty sequence of elements within an array.

### Examples:
#### Example 1:

Input: nums = `[1,2,1,4,1]`<br>
Output: 1<br>
Explanation:<br>
Only the subarray `[1,4,1]` contains exactly 3 elements where the sum of the first and third numbers equals half the middle number.

#### Example 2:
Input: nums = `[1,1,1]`<br>
Output: 0<br>
Explanation:<br>
`[1,1,1]` is the only subarray of length 3. However, its first and third numbers do not add to half the middle number.


### Constraints:

- 3 <= `nums.length` <= 100
- -100 <= `nums[i]` <= 100

### <i>Biweekly Contest Question</i>