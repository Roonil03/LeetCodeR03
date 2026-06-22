# Biweekly Contest 185
## Ranking:
## Question 1:
*Create Grid With Exactly One Path*
### Difficulty: Easy
### Points: 3
### Description:
You are given two integers m and n, representing the number of rows and columns of a grid.

Construct any m x n grid consisting only of the characters '.' and '#', where:
- '`.`' represents a free cell.
- '`#`' represents an obstacle cell.

A valid path is a sequence of free cells that:
- Starts at the top-left cell `(0, 0)`.
- Ends at the bottom-right cell `(m - 1, n - 1)`.
- Moves only:
    - Right, from `(i, j)` to `(i, j + 1)`, or
    - Down, from `(i, j)` to `(i + 1, j)`.

Return any grid such that there is exactly one valid path from the top-left cell to the bottom-right cell.

### Examples:
#### Example 1:

Input: m = 2, n = 3

Output: `["..#","#.."]`

Explanation:

<img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-26-at-61005pm.png"><br>

The only valid path is: `(0,0) → (0,1) → (1,1) → (1,2)`

#### Example 2:

Input: m = 3, n = 3

Output: `["..#","#..","##."]`

Explanation:

<img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-26-at-61129pm.png"><br>

The only valid path is: `(0,0) → (0,1) → (1,1) → (1,2) → (2,2)`

#### Example 3:

Input: m = 1, n = 4

Output: `["...."]`

Explanation:

The only valid path is: `(0,0) → (0,1) → (0,2) → (0,3)`

### Constraints:

- 1 <= m, n <= 25

## Question 2:
*Minimum Lights to Illuminate a Road*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer array lights of length n, representing positions 0 through n - 1 on a road.

For each position i:
- If `lights[i]` = v, where v > 0, there is a working bulb at position i that illuminates every position from `max(0, i - v)` to `min(n - 1, i + v)`, inclusive.
- If `lights[i]` = 0, there is no working bulb at position i.

A position is visible if it is illuminated by at least one working bulb.

You may install additional bulbs at any positions. Each additional bulb installed at position j illuminates positions from `max(0, j - 1)` to `min(n - 1, j + 1)`, inclusive.

Return the minimum number of additional bulbs required to make every position on the road visible.

### Examples:
#### Example 1:

Input: lights = `[0,0,0,0]`

Output: 2

Explanation:

One optimal placement is:
- Install an additional bulb at position 1, illuminating positions `[0, 1, 2]`.
- Install an additional bulb at position 3, illuminating positions `[2, 3]`.  
Therefore, the minimum number of additional bulbs required is 2.

#### Example 2:

Input: lights = `[0,0,0,2,0]`

Output: 1

Explanation:

- Since `lights[3]` = 2, the working bulb at position 3 illuminates positions `[1, 2, 3, 4]`.
- Installing an additional bulb at position 1 illuminates positions `[0, 1, 2]`, making every position visible.
- Therefore, the minimum number of additional bulbs required is 1.

### Constraints:

- 1 <= `n` == `lights.length` <= 10<sup>5</sup>
- 0 <= `lights[i]` <= `n`

## Question 3:
*Finish Time of Tasks I*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer n representing the number of tasks in a project, numbered from 0 to n - 1. These tasks are connected as a tree rooted at task 0. This is represented by a 2D integer array edges of length n - 1, where `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`]` indicates that task `u`<sub>`i`</sub> is the parent of task `v`<sub>`i`</sub>.

You are also given an array baseTime of length n, where `baseTime[i]` represents the time to complete task i.

The finish time of each task is calculated as follows:

- Leaf task: The finish time is `baseTime[i]`.
- Non-leaf task:
    - Let `earliest` be the minimum finish time among its children, and latest be the maximum finish time among its children.
    - Let `ownDuration` be `(latest - earliest) + baseTime[i]`.
    - The finish time of task i is `latest + ownDuration`.

Return the finish time of the root task 0.

### Examples:
#### Example 1:

Input: n = 3, edges = `[[0,1],[1,2]]`, baseTime = `[9,5,3]`

Output: 17

Explanation:  
- Task 2 is a leaf, so its finish time is `baseTime[2]` = 3.
- Task 1 has one child task 2:
    - `earliest = latest = 3`
    - `ownDuration = (latest - earliest) + baseTime[1] = 5`
    - Finish time of task 1 is 3 + 5 = 8
- Task 0 has one child with finish time 8:
    - `earliest = latest = 8`
    - `ownDuration = (latest - earliest) + baseTime[0] = 9`
    - Finish time of task 0 is 8 + 9 = 17
#### Example 2:

Input: n = 3, edges = `[[0,1],[0,2]]`, baseTime = `[4,7,6]`

Output: 12

Explanation:

- Task 1 is a leaf, so its finish time is `baseTime[1]` = 7.
- Task 2 is a leaf, so its finish time is `baseTime[2]` = 6.
- Task 0 has two children with finish times 7 and 6:
    - `earliest = 6`, `latest = 7`
    - `ownDuration = (latest - earliest) + baseTime[0] = (7 - 6) + 4 = 5`
    - Finish time of task 0 is `latest + ownDuration = 7 + 5 = 12`
#### Example 3:

Input: n = 4, edges = `[[0,1],[0,2],[2,3]]`, baseTime = `[5,8,2,1]`

Output: 18

Explanation:
- Task 1 is a leaf, so its finish time is `baseTime[1]` = 8.
- Task 3 is a leaf, so its finish time is `baseTime[3]` = 1.
- Task 2 has one child task 3:
    - `earliest = latest = 1`
    - `ownDuration = (latest - earliest) + baseTime[2] = 0 + 2 = 2`
    - Finish time of task 2 is `latest + ownDuration = 1 + 2 = 3`
- Task 0 has two children with finish times 8 and 3:
    - `earliest = 3`, `latest = 8`
    - `ownDuration = (latest - earliest) + baseTime[0] = (8 - 3) + 5 = 10`
    - Finish time of task 0 is `latest + ownDuration = 8 + 10 = 18`

### Constraints:

- 1 <= `n` <= 10<sup>5</sup>
- `edges.length` = `n - 1`
- `edges[i]` == `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`]`
- 0 <= `u`<sub>`i`</sub>, `v`<sub>`i`</sub> <= `n - 1`
- `u`<sub>`i`</sub> != `v`<sub>`i`</sub>
- The input is generated such that edges represents a valid tree.
- `baseTime.length` == `n`
- 1 <= `baseTime[i]` <= 10<sup>5</sup>​​​​​​​

## Question 4:
*Count Good Integers in a Range*
### Difficulty: Hard
### Points: 6
### Description:
You are given three integers l, r and k.

A number is considered good if the absolute difference between every pair of adjacent digits is at most k.

Return the number of good integers in the range `[l, r]` (inclusive).

The absolute difference between values x and y is defined as `abs(x - y)`.

### Examples:
#### Example 1:

Input: l = 10, r = 15, k = 1

Output: 3

Explanation:
- The good integers in the range are 10, 11, and 12.
- For 10, `abs(1 - 0)` = 1.
- For 11, `abs(1 - 1)` = 0.
- For 12, `abs(1 - 2)` = 1.
- All these differences are at most k = 1. Thus, the answer is 3.
#### Example 2:

Input: l = 201, r = 204, k = 2

Output: 2

Explanation:
- The good integers in the range are 201 and 202.
- For 201, `abs(2 - 0) = 2` and `abs(0 - 1) = 1`.
- For 202, `abs(2 - 0) = 2` and `abs(0 - 2) = 2`.
- Thus, the answer is 2.

### Constraints:

- 10 <= l <= r <= 10<sup>15</sup>
- 0 <= k <= 9