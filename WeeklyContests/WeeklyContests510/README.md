# Weekly Contest 510
## Ranking: 651 / 40113
## Question 1:
*Number of Elapsed Seconds Between Two Times*
### Difficulty: Easy
### Points: 3
### Description:
You are given two valid times `startTime` and `endTime`, each represented as a string in the format "`HH:MM:SS`".

Return the number of seconds that have elapsed from `startTime` to `endTime`.

### Examples:
#### Example 1:

Input: startTime = "`01:00:00`", endTime = "`01:00:25`"

Output: 25

Explanation:

endTime is 25 seconds ahead of startTime.
#### Example 2:

Input: startTime = "`12:34:56`", endTime = "`13:00:00`"

Output: 1504

Explanation:

endTime is 25 minutes and 4 seconds ahead of startTime, which equals 1504 seconds.

### Constraints:

- `startTime.length` == 8
- `endTime.length` == 8
- `startTime` and `endTime` are valid times in the format "`HH:MM:SS`"
- 00 <= `HH` <= 23
- 00 <= `MM` <= 59
- 00 <= `SS` <= 59
- `endTime` is not earlier than `startTime`

## Question 2:
*Minimum Total Cost to Process All Elements*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums and an integer k.

Initially, you have k units of resources.

You must process the elements of nums from left to right. To process the ith element, you need `nums[i]` resources.

If your available resources are less than `nums[i]`, you may perform an operation that increases your available resources by k. The value of k is fixed and does not change throughout the process. The first such operation incurs a cost of 1, the second incurs a cost of 2, and so on.

After processing the ith element, your available resources decrease by `nums[i]`.

Return an integer denoting the minimum total cost required to process all elements. Since the answer may be very large, return it modulo 10<sup>9</sup> + 7.

### Examples:
#### Example 1:

Input: nums = `[1,2,3,4]`, k = 4

Output: 3

Explanation:

- After processing `nums[0]`, we have 4 - 1 = 3 units of resources left.
- After processing `nums[1]`, we have 3 - 2 = 1 unit of resources left.
- Since `nums[2]` = 3 and only 1 unit of resources is available, we perform the first operation costing 1. After processing `nums[2]`, we have 1 + 4 - 3 = 2 units of resources left.
- Since `nums[3]` = 4 and only 2 units of resources are available, we perform the second operation costing 2, to have 2 + 4 = 6 units of resources, which is enough to process `nums[3]`.
- Thus, the total cost is 1 + 2 = 3.
#### Example 2:

Input: nums = `[1,1,7,14]`, k = 4

Output: 15

Explanation:

- After processing `nums[0]`, we have 4 - 1 = 3 units of resources left.
- After processing `nums[1]`, we have 3 - 1 = 2 units of resources left.
- Since `nums[2]` = 7 and only 2 units of resources are available, we perform two operations costing 1 + 2 = 3. After processing `nums[2]`, we have 2 + 4 + 4 - 7 = 3 units of resources left.
- Since `nums[3]` = 14 and only 3 units of resources are available, we perform three operations costing 3 + 4 + 5 = 12, to have 3 + 4 + 4 + 4 = 15 units of resources, which is enough to process `nums[3]`.
- Thus, the total cost is 3 + 12 = 15.
#### Example 3:

Input: nums = `[1,2,3,4]`, k = 10

Output: 0

Explanation:
- To process all elements, we can use the initial 10 units of resources without performing any operations. Thus, the total cost required is 0.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>9</sup>
- 1 <= `k` <= 10<sup>9</sup>

## Question 3:
*Create Grid With Exactly K Paths I*
### Difficulty: Medium
### Points: 5
### Description:
You are given three integers m, n, and k.

Construct any m x n grid consisting only of the characters '`.`' and '`#`', where:
- '`.`' represents a free cell.
- '`#`' represents an obstacle cell.

A valid path is a sequence of free cells that:
- Starts at the top-left cell `(0, 0)`.
- Ends at the bottom-right cell `(m - 1, n - 1)`.
- Moves only:
    - Right, from `(i, j)` to `(i, j + 1)`, or
    - Down, from `(i, j)` to `(i + 1, j)`.

Return any grid such that there are exactly k valid paths from the top-left cell to the bottom-right cell. If no such grid exists, return an empty array.

### Examples:
#### Example 1:

Input: m = 2, n = 3, k = 2

Output: `["...","#.."]`

Explanation:

<img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-27-at-113554am.png"><br>

There are exactly k = 2 valid paths from (0, 0) to (1, 2):
- `(0, 0) → (0, 1) → (0, 2) → (1, 2)`
- `(0, 0) → (0, 1) → (1, 1) → (1, 2)`
#### Example 2:

Input: m = 3, n = 3, k = 4

Output: `["..#","...","#.."]`

Explanation:

<img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-27-at-113452am.png"><br>

There are exactly k = 4 valid paths from `(0, 0)` to `(2, 2)`:
- `(0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)`
- `(0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)`
- `(0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)`
- `(0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2)`
#### Example 3:

Input: m = 1, n = 4, k = 2

Output: `[]`

Explanation:​

No grid exists with exactly k = 2 valid paths for a `1 x 4` grid, so the answer is an empty array.

### Constraints:

- 1 <= `m`, `n` <= 10
- 1 <= `k` <= 4

## Question 4:
*Maximum Consistent Columns in a Grid*
### Difficulty: Hard
### Points: 6
### Description:
You are given a 2D integer array grid of size `m x n`, and an integer `limit`.

You may remove zero or more columns from the grid, but at least one column must remain. The relative order of the remaining columns must be preserved.

A grid is called consistent if for every row i, and for every pair of adjacent remaining columns a and b with `a < b`, the following holds: `|grid[i][b] - grid[i][a]|` <= `limit`.

Return the maximum number of columns that can remain such that the resulting grid is consistent.

### Example 1:

Input: grid = `[[-2,0,3]]`, limit = 2

Output: 2

Explanation:
- Remove column 2 and keep columns 0 and 1, which gives `|grid[0][1] − grid[0][0]|` = `|0 − (−2)|` = `2` <= `limit`.
- Thus, the maximum number of columns that can remain is 2.
#### Example 2:

Input: grid = `[[1,-1,1],[2,2,2]]`, limit = 1

Output: 2

Explanation:
- Remove column 1 and keep columns 0 and 2, which gives
    - `|grid[0][2] − grid[0][0]|` = `|1 − 1|` = `0` <= `limit` and
    - `|grid[1][2] − grid[1][0]|` = `|2 − 2|` = `0` <= `limit`.
- Thus, the maximum number of columns that can remain is 2.
#### Example 3:

Input: grid = `[[-5,5]]`, limit = 9

Output: 1

Explanation:
- Remove either column 0 or column 1, since `|grid[0][1] − grid[0][0]|` = `|5 − (−5)|` = `10` > `limit`.
- Thus, the maximum number of columns that can remain is 1.

### Constraints:

- 1 <= `m` == `grid.length` <= 250
- 1 <= `n` == `grid[i].length` <= 250
- -10<sup>5</sup> <= `grid[i][j]` <= 10<sup>5</sup>
- 0 <= `limit` <= 10<sup>5</sup>
