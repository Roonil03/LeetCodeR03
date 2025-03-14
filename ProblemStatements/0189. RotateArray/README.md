# 189. Rotate Array
## Question Level: Medium
### Description:
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
### Examples:
#### Example 1:

Input: nums = `[1,2,3,4,5,6,7]`, k = 3<br>
Output: `[5,6,7,1,2,3,4]`<br>
Explanation:<br>
rotate 1 steps to the right: `[7,1,2,3,4,5,6]`<br>
rotate 2 steps to the right: `[6,7,1,2,3,4,5]`<br>
rotate 3 steps to the right: `[5,6,7,1,2,3,4]`<br>
#### Example 2:

Input: nums = `[-1,-100,3,99]`, k = 2<br>
Output: `[3,99,-1,-100]`<br>
Explanation: <br>
rotate 1 steps to the right: `[99,-1,-100,3]`<br>
rotate 2 steps to the right: `[3,99,-1,-100]`<br>


### Constraints:

- 1 <= `nums.length` <= 10^5
- -2^31 <= `nums[i]` <= 2^31 - 1
- 0 <= k <= 10^5


### <i>Follow up:

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with `O(1)` extra space?

### Concepts Used:
- Array
- Math
- Two Pointers </i>