# 2257. Count Unguarded Cells in the Grid
## Question Level: Medium
### Description:
You are given two integers m and n representing a 0-indexed m x n grid. You are also given two 2D integer arrays guards and walls where `guards[i]` = `[rowi, coli]` and `walls[j]` = `[rowj, colj]` represent the positions of the ith guard and jth wall respectively.

A guard can see every cell in the four cardinal directions (north, east, south, or west) starting from their position unless obstructed by a wall or another guard. A cell is guarded if there is at least one guard that can see it.

Return the number of unoccupied cells that are not guarded.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2022/03/10/example1drawio2.png"><br>
Input: m = 4, n = 6, guards = `[[0,0],[1,1],[2,3]]`, walls = `[[0,1],[2,2],[1,4]]`<br>
Output: 7<br>
Explanation: <br>The guarded and unguarded cells are shown in red and green respectively in the above diagram.<br>
There are a total of 7 unguarded cells, so we return 7.

#### Example 2:
<img src="https://assets.leetcode.com/uploads/2022/03/10/example2drawio.png"><br>
Input: m = 3, n = 3, guards = `[[1,1]]`, walls = `[[0,1],[1,0],[2,1],[1,2]]`<br>
Output: 4<br>
Explanation: <br>The unguarded cells are shown in green in the above diagram.<br>
There are a total of 4 unguarded cells, so we return 4.<br>

### Constraints:

- 1 <= m, n <= 10^5
- 2 <= m * n <= 10^5
- 1 <= `guards.length`, `walls.length` <= 5 * 10^4
- 2 <= `guards.length` + `walls.length` <= m * n
- `guards[i].length` == `walls[j].length` == 2
- 0 <= `rowi`, `rowj` < m
- 0 <= `coli`, `colj` < n
- All the positions in guards and walls are unique.

### <i>Concepts Used:
- Array
- Matrix
- Simulation </i>