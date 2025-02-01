# 827. Making A Large Island
## Question Level: Hard
### Description:
You are given an `n x n` binary matrix grid. You are allowed to change at most one 0 to be 1.

Return the size of the largest island in grid after applying this operation.

An island is a 4-directionally connected group of 1s.

### Examples:
#### Example 1:

Input: grid = `[[1,0],[0,1]]`<br>
Output: 3<br>
Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.<br>
#### Example 2:

Input: grid = `[[1,1],[1,0]]`<br>
Output: 4<br>
Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4.<br>
#### Example 3:

Input: grid = `[[1,1],[1,1]]`<br>
Output: 4<br>
Explanation: Can't change any 0 to 1, only one island with area = 4.<br>


### Constraints:

- n == `grid.length`
- n == `grid[i].length`
- 1 <= n <= 500
- `grid[i][j]` is either 0 or 1.


### <i>Concepts Used:
- Array
- Depth-First Search
- Breadth-First Search
- Union Find
- Matrix </i>