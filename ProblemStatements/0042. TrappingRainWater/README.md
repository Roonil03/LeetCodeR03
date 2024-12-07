# 42. Trapping Rain Water
## Question Level: Hard
### Description:
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png"><br>
Input: height = `[0,1,0,2,1,0,1,3,2,1,2,1]`<br>
Output: 6<br>
Explanation: The above elevation map (black section) is represented by array `[0,1,0,2,1,0,1,3,2,1,2,1]`. In this case, 6 units of rain water (blue section) are being trapped.
#### Example 2:

Input: height = `[4,2,0,3,2,5]`<br>
Output: 9<br>


### Constraints:
- n == `height.length`
- 1 <= `n` <= 2 * 10^4
- 0 <= `height[i]` <= 10^5

### <i>Concepts Used:
- Array
- Two Pointers
- Dynamic Programming
- Stack
- Monotonic Stack </i>