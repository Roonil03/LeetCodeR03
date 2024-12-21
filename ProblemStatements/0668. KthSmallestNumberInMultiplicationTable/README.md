# 668. Kth Smallest Number in Multiplication Table
## Question Level: Hard
### Description:
Nearly everyone has used the Multiplication Table. The multiplication table of size m x n is an integer matrix mat where `mat[i][j]` == `i * j` (1-indexed).

Given three integers `m`, `n`, and `k`, return the kth smallest element in the `m x n` multiplication table.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2021/05/02/multtable1-grid.jpg"><br>
Input: m = 3, n = 3, k = 5<br>
Output: 3<br>
Explanation: The 5th smallest number is 3.<br>
#### Example 2:
<img src="https://assets.leetcode.com/uploads/2021/05/02/multtable2-grid.jpg"><br>
Input: m = 2, n = 3, k = 6<br>
Output: 6<br>
Explanation: The 6th smallest number is 6.

### Constraints:

- 1 <= `m, n` <= `3 * 10^4`
- 1 <= `k` <= `m * n`

### <i>Concepts Used:
- Math
- Binary Search </i>