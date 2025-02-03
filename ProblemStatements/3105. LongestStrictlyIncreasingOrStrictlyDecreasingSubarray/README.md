# 3105. Longest Strictly Increasing or Strictly Decreasing Subarray
## Question Level: Easy
### Description:
You are given an array of integers nums. Return the length of the longest subarray of nums which is either strictly increasing or strictly decreasing.

### Examples:
#### Example 1:
Input: nums = `[1,4,3,3,2]`<br>
Output: 2<br>
Explanation:<br>
The strictly increasing subarrays of nums are `[1]`, `[2]`, `[3]`, `[3]`, `[4]`, and `[1,4]`.<br>
The strictly decreasing subarrays of nums are `[1]`, `[2]`, `[3]`, `[3]`, `[4]`, `[3,2]`, and `[4,3]`.<br>
Hence, we return 2.

#### Example 2:
Input: nums = `[3,3,3,3]`<br>
Output: 1<br>
Explanation:<br>
The strictly increasing subarrays of nums are `[3]`, `[3]`, `[3]`, and `[3]`.<br>
The strictly decreasing subarrays of nums are `[3]`, `[3]`, `[3]`, and `[3]`.<br>
Hence, we return 1.

#### Example 3:
Input: nums = `[3,2,1]`<br>
Output: 3<br>
Explanation:<br>
The strictly increasing subarrays of nums are `[3]`, `[2]`, and `[1]`.<br>
The strictly decreasing subarrays of nums are `[3]`, `[2]`, `[1]`, `[3,2]`, `[2,1]`, and `[3,2,1]`.<br>
Hence, we return 3.


### Constraints:

- 1 <= `nums.length` <= 50
- 1 <= `nums[i]` <= 50

### <i>Concepts Used:
- Array </i>