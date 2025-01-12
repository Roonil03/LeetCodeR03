# Weekly Contests 432
## Ranking:
## Question 1:
<i>Zigzag Grid Traversal With Skip</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given an `m x n` 2D array grid of positive integers.

Your task is to traverse grid in a zigzag pattern while skipping every alternate cell.

Zigzag pattern traversal is defined as following the below actions:
- Start at the top-left cell `(0, 0)`.
- Move right within a row until the end of the row is reached.
- Drop down to the next row, then traverse left until the beginning of the row is reached.
- Continue alternating between right and left traversal until every row has been traversed.

Note that you must skip every alternate cell during the traversal.

Return an array of integers result containing, in order, the value of the cells visited during the zigzag traversal with skips.

### Testcases:
#### Example 1:
Input: grid = `[[1,2],[3,4]]`<br>
Output: `[1,4]`<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example0.png"><br>
#### Example 2:
Input: grid = `[[2,1],[2,1],[2,1]]`<br>
Output: `[2,1,2]`<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example1.png"><br>
#### Example 3:
Input: grid = `[[1,2,3],[4,5,6],[7,8,9]]`<br>
Output: `[1,3,5,7,9]`<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example2.png"><br>


### Constraints:

- 2 <= n == `grid.length` <= 50
- 2 <= m == `grid[i].length` <= 50
- 1 <= `grid[i][j]` <= 2500

## Question 2:
<i>Maximum Amount of Money Robot Can Earn</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an `m x n` grid. A robot starts at the top-left corner of the grid `(0, 0)` and wants to reach the bottom-right corner `(m - 1, n - 1)`. The robot can move either right or down at any point in time.

The grid contains a value `coins[i][j]` in each cell:

If `coins[i][j]` >= 0, the robot gains that many coins.
If `coins[i][j]` < 0, the robot encounters a robber, and the robber steals the absolute value of `coins[i][j]` coins.
The robot has a special ability to neutralize robbers in at most 2 cells on its path, preventing them from stealing coins in those cells.

Note: The robot's total coins can be negative.

Return the maximum profit the robot can gain on the route.

### Testcases:
#### Example 1:

Input: coins = `[[0,1,-1],[1,-2,3],[2,-3,4]]`<br>
Output: 8<br>
Explanation:<br>
An optimal path for maximum coins is:<br>
1. Start at `(0, 0)` with 0 coins (total coins = 0).
2. Move to `(0, 1)`, gaining 1 coin (total coins = 0 + 1 = 1).
3. Move to `(1, 1)`, where there's a robber stealing 2 coins. The robot uses one neutralization here, avoiding the robbery (total coins = 1).
4. Move to `(1, 2)`, gaining 3 coins (total coins = 1 + 3 = 4).
5. Move to `(2, 2)`, gaining 4 coins (total coins = 4 + 4 = 8).

#### Example 2:
Input: coins = `[[10,10,10],[10,10,10]]`<br>
Output: 40<br>
Explanation:<br>
An optimal path for maximum coins is:<br>
1. Start at `(0, 0)` with 10 coins (total coins = 10).
2. Move to `(0, 1)`, gaining 10 coins (total coins = 10 + 10 = 20).
3. Move to `(0, 2)`, gaining another 10 coins (total coins = 20 + 10 = 30).
4. Move to `(1, 2)`, gaining the final 10 coins (total coins = 30 + 10 = 40).

### Constraints:

- m == `coins.length`
- n == `coins[i].length`
- 1 <= m, n <= 500
- -1000 <= `coins[i][j]` <= 1000

## Question 3:
<i>Minimize the Maximum Edge Weight of Graph</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given two integers, n and threshold, as well as a directed weighted graph of n nodes numbered from 0 to n - 1. The graph is represented by a 2D integer array edges, where `edges[i]` = `[Ai, Bi, Wi]` indicates that there is an edge going from node `Ai` to node `Bi` with weight `Wi`.

You have to remove some edges from this graph (possibly none), so that it satisfies the following conditions:
1. Node 0 must be reachable from all other nodes.
2. The maximum edge weight in the resulting graph is minimized.
3. Each node has at most threshold outgoing edges.

Return the minimum possible value of the maximum edge weight after removing the necessary edges. If it is impossible for all conditions to be satisfied, return -1.

### Testcases:
#### Example 1:
Input: n = 5, edges = `[[1,0,1],[2,0,2],[3,0,1],[4,3,1],[2,1,1]]`, threshold = 2<br>
Output: 1<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/12/09/s-1.png"><br>
Remove the edge `2 -> 0`. The maximum weight among the remaining edges is 1.

#### Example 2:
Input: n = 5, edges = `[[0,1,1],[0,2,2],[0,3,1],[0,4,1],[1,2,1],[1,4,1]]`, threshold = 1<br>
Output: -1<br>
Explanation: <br>
It is impossible to reach node 0 from node 2.

#### Example 3:
Input: n = 5, edges = `[[1,2,1],[1,3,3],[1,4,5],[2,3,2],[3,4,2],[4,0,1]]`, threshold = 1<br>
Output: 2<br>
Explanation: <br>
<img src="https://assets.leetcode.com/uploads/2024/12/09/s2-1.png"><br>
Remove the edges `1 -> 3` and `1 -> 4`. The maximum weight among the remaining edges is 2.

#### Example 4:
Input: n = 5, edges = `[[1,2,1],[1,3,3],[1,4,5],[2,3,2],[4,0,1]]`, threshold = 1<br>
Output: -1<br>

### Constraints:

- 2 <= n <= 10^5
- 1 <= `threshold` <= n - 1
- 1 <= `edges.length` <= `min(105, n * (n - 1) / 2)`.
- `edges[i].length` == 3
- 0 <= `Ai`, `Bi` < n
- `Ai` != `Bi`
- 1 <= `Wi` <= 10^6
- There may be multiple edges between a pair of nodes, but they must have unique weights.

## Question 4:
<i>Count Non-Decreasing Subarrays After K Operations</i>
## Difficulty: Hard
### Points: 7
### Description:
You are given an array nums of n integers and an integer k.

For each subarray of nums, you can apply up to k operations on it. In each operation, you increment any element of the subarray by 1.

Note that each subarray is considered independently, meaning changes made to one subarray do not persist to another.

Return the number of subarrays that you can make non-decreasing ​​​​​after performing at most k operations.

An array is said to be non-decreasing if each element is greater than or equal to its previous element, if it exists

### Testcases:
#### Example 1:
Input: nums = `[6,3,1,2,4,4]`, k = 7<br>
Output: 17<br>
Explanation:<br>
Out of all 21 possible subarrays of nums, only the subarrays `[6, 3, 1]`, `[6, 3, 1, 2]`, `[6, 3, 1, 2, 4]` and `[6, 3, 1, 2, 4, 4]` cannot be made non-decreasing after applying up to k = 7 operations. Thus, the number of non-decreasing subarrays is `21 - 4` = `17`.

#### Example 2:
Input: nums = `[6,3,1,3,6]`, k = 4<br>
Output: 12<br>
Explanation:<br>
The subarray `[3, 1, 3, 6]` along with all subarrays of nums with three or fewer elements, except `[6, 3, 1]`, can be made non-decreasing after k operations. There are 5 subarrays of a single element, 4 subarrays of two elements, and 2 subarrays of three elements except `[6, 3, 1]`, so there are `1 + 5 + 4 + 2` = `12` subarrays that can be made non-decreasing.

### Constraints:

- 1 <= `nums.length` <= 10^5
- 1 <= `nums[i]` <= 10^9
- 1 <= k <= 10^9