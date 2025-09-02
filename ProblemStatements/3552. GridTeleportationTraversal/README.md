# 3552. Grid Teleportation Traversal
## Question Level: Medium
### Description:
You are given a 2D character grid matrix of size m x n, represented as an array of strings, where matrix[i][j] represents the cell at the intersection of the ith row and jth column. Each cell is one of the following:

- '`.`' representing an empty cell.
- '`#`' representing an obstacle.
- An uppercase letter `('A'-'Z')` representing a teleportation portal.

You start at the top-left cell `(0, 0)`, and your goal is to reach the bottom-right cell `(m - 1, n - 1)`. You can move from the current cell to any adjacent cell `(up, down, left, right)` as long as the destination cell is within the grid bounds and is not an obstacle.

If you step on a cell containing a portal letter and you haven't used that portal letter before, you may instantly teleport to any other cell in the grid with the same letter. This teleportation does not count as a move, but each portal letter can be used at most once during your journey.

Return the minimum number of moves required to reach the bottom-right cell. If it is not possible to reach the destination, return -1.

### Examples:
#### Example 1:

Input: matrix = `["A..",".A.","..."]`

Output: 2

Explanation:

<img src="https://assets.leetcode.com/uploads/2025/03/15/example04140.png"><br>

Before the first move, teleport from `(0, 0)` to `(1, 1)`.
In the first move, move from `(1, 1)` to `(1, 2)`.
In the second move, move from `(1, 2)` to `(2, 2)`.
#### Example 2:

Input: matrix = `[".#...",".#.#.",".#.#.","...#."]`

Output: 13

Explanation:

<img src="https://assets.leetcode.com/uploads/2025/03/15/ezgifcom-animated-gif-maker.gif"><br>

### Constraints:

- 1 <= `m` == `matrix.length` <= 10<sup>3</sup>
- 1 <= `n` == `matrix[i].length` <= 10<sup>3</sup>
- `matrix[i][j]` is either '`#`', '`.`', or an uppercase English letter.
- `matrix[0][0]` is not an obstacle.

### <i>Concepts Used:
- Array
- Hash Table
- Breadth-First Search
- Matrix</i>