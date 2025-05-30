# 1004. Max Consecutive Ones III
## Question Level: Medium
### Description:
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

### Examples:
#### Example 1:

Input: nums = `[1,1,1,0,0,0,1,1,1,1,0]`, k = 2  
Output: 6  
Explanation: [1,1,1,0,0,<i><b>1</b>,1,1,1,1,<b>1</b></i>]  
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
#### Example 2:

Input: nums = `[0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]`, k = 3  
Output: 10  
Explanation: [0,0,<i>1,1,<b>1</b>,<b>1</b>,1,1,1,<b>1</b>,1,1</i>,0,0,0,1,1,1,1]  
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- `nums[i]` is either 0 or 1.
- 0 <= k <= `nums.length`

### <i>Concepts Used:
- Array
- Binary Search
- Sliding Window
- Prefix Sum</i>