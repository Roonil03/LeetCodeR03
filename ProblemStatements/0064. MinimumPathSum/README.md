# 64. Minimum Path Sum
## Question Level: Medium
### Description:
Given a `m x n` grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2020/11/05/minpath.jpg"><br>
Input: grid = `[[1,3,1],[1,5,1],[4,2,1]]`<br>
Output: 7<br>
Explanation: Because the path `1 → 3 → 1 → 1 → 1` minimizes the sum.
#### Example 2:

Input: grid = `[[1,2,3],[4,5,6]]`<br>
Output: 12<br>

###  Constraints:

- m == `grid.length`
- n == `grid[i].length`
- 1 <= `m`, `n` <= 200
- 0 <= `grid[i][j]` <= 200

### <i>Concepts Used:
- Array
- Dynamic Programming
- Matrix </i>