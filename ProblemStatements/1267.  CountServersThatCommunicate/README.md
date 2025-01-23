# 1267. Count Servers that Communicate
## Question Level: Medium
### Description:
You are given a map of a server center, represented as a `m * n` integer matrix grid, where 1 means that on that cell there is a server and 0 means that it is no server. Two servers are said to communicate if they are on the same row or on the same column.

Return the number of servers that communicate with any other server.

### Examples:
#### Example 1:


<img src="https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-6.jpg"><br>
Input: grid = `[[1,0],[0,1]]`<br>
Output: 0<br>
Explanation: No servers can communicate with others.<br>
#### Example 2:


<img src="https://assets.leetcode.com/uploads/2019/11/13/untitled-diagram-4.jpg"><br>
Input: grid = `[[1,0],[1,1]]`<br>
Output: 3<br>
Explanation: All three servers can communicate with at least one other server.
#### Example 3:


<img src="https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-1-3.jpg"><br>
Input: grid = `[[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]`<br>
Output: 4<br>
Explanation: The two servers in the first row can communicate with each other. The two servers in the third column can communicate with each other. The server at right bottom corner can't communicate with any other server.<br>

### Constraints:

- m == `grid.length`
- n == `grid[i].length`
- 1 <= m <= 250
- 1 <= n <= 250
- `grid[i][j]` == 0 or 1

### <i>Concepts Used:
- Array
- Depth-First Search
- Breadth-First Search
- Union Find
- Matrix
- Counting </i>