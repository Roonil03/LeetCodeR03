# 85. Maximal Rectangle
## Question Level: Hard
### Description:
Given a `rows` x `cols` binary matrix filled with `0`'s and `1`'s, find the largest rectangle containing only 1's and return its area.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2020/09/14/maximal.jpg"><br>
Input: matrix = `[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]`<br>
Output: 6<br>
Explanation: The maximal rectangle is shown in the above picture.<br>
#### Example 2:

Input: matrix = `[["0"]]`<br>
Output: 0<br>
#### Example 3:

Input: matrix = `[["1"]]`<br>
Output: 1<br>

### Constraints:

- `rows` == `matrix.length`
- `cols` == `matrix[i].length`
- 1 <= `row`, `cols` <= 200
- `matrix[i][j]` is '0' or '1'.

### <i>Concepts Used:
- Array
- Dynamic Programming
- Stack
- Matrix
- Monotonic Stack </i>