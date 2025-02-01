# 324. Wiggle Sort II
## Question Level: Medium
### Description:
Given an integer array nums, reorder it such that `nums[0]` < `nums[1]` > `nums[2]` < `nums[3]`....

You may assume the input array always has a valid answer.

### Examples
#### Example 1:

Input: nums = `[1,5,1,1,6,4]`<br>
Output: `[1,6,1,5,1,4]`<br>
Explanation: `[1,4,1,5,1,6]` is also accepted.<br>
#### Example 2:

Input: nums = `[1,3,2,2,3,1]`<br>
Output: `[2,3,1,3,1,2]`<br>

### Constraints:

- 1 <= `nums.length` <= 5 * 10^4
- 0 <= `nums[i]` <= 5000
- It is guaranteed that there will be an answer for the given input nums.


### <i>Follow Up: 
Can you do it in `O(n)` time and/or in-place with `O(1)` extra space?

### Concepts Used:
- Array
- Divide and Conquer
- Greedy
- Sorting
- Quickselect </i>