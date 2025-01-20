# 2661. First Completely Painted Row or Column
## Question Level: Medium
### Description:
You are given a 0-indexed integer array arr, and an `m x n` integer matrix `mat`. `arr` and `mat` both contain all the integers in the range `[1, m * n]`.

Go through each index i in arr starting from index 0 and paint the cell in mat containing the integer `arr[i]`.

Return the smallest index i at which either a row or a column will be completely painted in `mat`.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2023/01/18/grid1.jpg"><br>
Input: arr = `[1,3,4,2]`, mat = `[[1,4],[2,3]]`<br>
Output: 2<br>
Explanation: The moves are shown in order, and both the first row and second column of the matrix become fully painted at `arr[2]`.<br>
#### Example 2:
<img src="https://assets.leetcode.com/uploads/2023/01/18/grid2.jpg"><br>
Input: arr = `[2,8,7,4,1,3,5,6,9]`, mat = `[[3,2,5],[1,4,6],[8,7,9]]`<br>
Output: 3<br>
Explanation: The second column becomes fully painted at `arr[3]`.<br>

### Constraints:

- m == `mat.length`
- n = `mat[i].length`
- `arr.length` == m * n
- 1 <= m, n <= 10^5
- 1 <= m * n <= 10^5
- 1 <= `arr[i]`, `mat[r][c]` <= m * n
- All the integers of arr are unique.
- All the integers of mat are unique.

### <i>Concepts Used:
- Array
- Hash Table
- Matrix </i>
 