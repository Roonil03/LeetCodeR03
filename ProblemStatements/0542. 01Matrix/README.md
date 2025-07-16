# 542. 01 Matrix
## Question Level: Medium
### Description:
Given an `m x n` binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two cells sharing a common edge is 1.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/04/24/01-1-grid.jpg"><br>
Input: mat = `[[0,0,0],[0,1,0],[0,0,0]]`  
Output: `[[0,0,0],[0,1,0],[0,0,0]]`  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/04/24/01-2-grid.jpg"><br>
Input: mat = `[[0,0,0],[0,1,0],[1,1,1]]`  
Output: `[[0,0,0],[0,1,0],[1,2,1]]`  

### Constraints:

- `m` == `mat.length`
- `n` == `mat[i].length`
- 1 <= `m`, `n` <= 10<sup>4</sup>
- 1 <= `m * n` <= 10<sup>4</sup>
- `mat[i][j]` is either 0 or 1.
- There is at least one 0 in `mat`.

### <i>Concepts Used:
- Array
- Dynamic Programming
- Breadth-First Search
- Matrix</i> 