# 3151. Special Array I
## Question Level: Easy
### Description:
An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers nums. Return true if nums is a special array, otherwise, return false.

### Examples:
#### Example 1:
Input: nums = `[1]`<br>
Output: true<br>
Explanation:<br>
There is only one element. So the answer is true.<br>

#### Example 2:
Input: nums = `[2,1,4]`<br>
Output: true<br>
Explanation:<br>
There is only two pairs: `(2,1)` and `(1,4)`, and both of them contain numbers with different parity. So the answer is true.

#### Example 3:
Input: nums = `[4,3,1,6]`<br>
Output: false<br>
Explanation:<br>
`nums[1]` and `nums[2]` are both odd. So the answer is false.

### Constraints:

1 <= `nums.length` <= 100
1 <= `nums[i]` <= 100

### <i>Concepts Used:
- Array</i>