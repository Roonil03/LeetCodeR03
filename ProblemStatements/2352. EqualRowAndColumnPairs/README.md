# 2352. Equal Row and Column Pairs
## Question Level: Medium
### Description:
Given a 0-indexed n x n integer matrix grid, return the number of pairs (r<sub>i</sub>, c<sub>j</sub>) such that row r<sub>i</sub> and column c<sub>j</sub> are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2022/06/01/ex1.jpg"><br>
Input: grid = `[[3,2,1],[1,7,6],[2,7,7]]`  
Output: 1  
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): `[2,7,7]`
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2022/06/01/ex2.jpg"><br>
Input: grid = `[[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]`  
Output: 3  
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): `[3,1,2,2]`
- (Row 2, Column 2): `[2,4,2,2]`
- (Row 3, Column 2): `[2,4,2,2]`

### Constraints:

- `n` == `grid.length` == `grid[i].length`
- 1 <= `n` <= 200
- 1 <= `grid[i][j]` <= 10<sup>5</sup>

### <i>Concepts Used:
- Array
- Hash Table
- Matrix
- Simulation</i>