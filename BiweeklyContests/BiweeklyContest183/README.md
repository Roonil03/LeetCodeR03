# Biweekly Contest 183
## Ranking: 
## Question 1:
*Minimum Swaps to Move Zeros to End*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array `nums`.

In one operation, you can choose any two distinct indices i and j and swap `nums[i]` and `nums[j]`.

Return an integer denoting the minimum number of operations required to move all 0s to the end of the array.

### Examples:
#### Example 1:

Input: nums = `[0,1,0,3,12]`

Output: 2

Explanation:

We perform the following swap operations:
- Swap `nums[0]` and `nums[3]`, giving nums = `[3, 1, 0, 0, 12]`.
- Swap `nums[2]` and `nums[4]`, giving nums = `[3, 1, 12, 0, 0]`.  
Thus, the answer is 2.

#### Example 2:

Input: nums = `[0,1,0,2]`

Output: 1

Explanation:

We perform the following swap operations:
- Swap `nums[0]` and `nums[3]`, giving nums = `[2, 1, 0, 0]`.  
Thus, the answer is 1.

#### Example 3:

Input: nums = `[1,2,0]`

Output: 0

Explanation:

The array already satisfies the condition. Therefore, no swap operations are needed.

### Constraints:

- 1 <= `nums.length` <= 100
- 0 <= `nums[i]` <= 100


## Question 2:
*Minimum Operations to Make Array Modulo Alternating I*
### Difficulty: Medium
### Points: 4
### Description:

You are given an integer array nums and an integer k.

In one operation, you can increase or decrease any element of nums by 1.

An array is called modulo alternating if there exist two distinct integers x and y (`0 <= x`, `y < k`) such that:
- For every even index i, `nums[i] % k == x`
- For every odd index i, `nums[i] % k == y`

Return the minimum number of operations required to make nums modulo alternating.

### Examples:
#### Example 1:

Input: nums = `[1,4,2,8]`, k = 3

Output: 2

Explanation:
- Let's choose x = 1 for even indices and y = 2 for odd indices.
- Perform the following operations:
    - Increment `nums[1]` = 4 by 1, giving nums = `[1, 5, 2, 8]`.
    - Decrement `nums[2]` = 2 by 1, giving nums = `[1, 5, 1, 8]`.
- Now, for even indices, nums[i] % k = 1, and for odd indices, `nums[i] % k` = 2.
- Thus, the total number of operations required is 2.

#### Example 2:

Input: nums = `[1,1,1]`, k = 3

Output: 1

Explanation:
- Incrementing `nums[1]` by 1 gives nums = `[1, 2, 1]`, which satisfies the condition with x = 1 and y = 2.
- Thus, the total number of operations required is 1.

### Constraints:

- 1 <= `nums.length` <= 100
- 1 <= `nums[i]` <= 10<sup>9</sup>
- 2 <= `k` <= 100

## Question 3:
*Maximum Path Intersection Sum in a Grid*
### Difficulty: Medium
### Points: 5
### Description:
You are given an m x n integer matrix grid.

Two players move across the grid:

- Player 1 starts at the top-left cell `(0, 0)` and can move only right or down. Their destination is the bottom-right cell (m - 1, n - 1).
- Player 2 starts at the bottom-left cell (`m - 1, 0)` and can move only right or up. Their destination is the top-right cell (0, n - 1).

Each player must choose a valid path from their respective starting cell to their destination.

A cell is called shared if it belongs to both chosen paths.

Return an integer denoting the maximum possible sum of values of all shared cells.

### Examples:
#### Example 1:
​​​​​​​​​​​​​​​​​​​​​
<img src="https://assets.leetcode.com/uploads/2026/05/19/image.png"><br>
Input: grid = `[[1,2,0,-3],[1,-2,1,0],[-4,2,-1,3],[3,-3,3,-2],[-1,-5,0,1]]`

Output: 4

Explanation:  
The diagram shows one optimal choice of paths.
- Player 1 follows the red/purple path from the top-left cell to the bottom-right cell:
    - (`0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3) → (3, 3) → (4, 3)`
- Player 2 follows the blue/purple path from the bottom-left cell to the top-right cell:
    - `(4, 0) → (4, 1) → (3, 1) → (2, 1) → (2, 2) → (2, 3) → (1, 3) → (0, 3)`
- The shared cells are `(2, 1)`, `(2, 2)`, and `(2, 3)`.
- The sum is `2 + (-1) + 3` = 4, which is the maximum possible sum.

#### Example 2:

<img src="https://assets.leetcode.com/uploads/2026/05/19/chatgpt-image-may-19-2026-01_39_39-pm.png"><br>
Input: grid = `[[4,-2,-3],[-1,-3,-1],[-4,2,-1]]`

Output: 3

Explanation:

One optimal pair of paths is shown in the diagram.  
- Player 1 follows the red/purple path:
    - `(0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)`
- Player 2 follows the blue/purple path:
    - `(2, 0) → (1, 0) → (0, 0) → (0, 1) → (0, 2)`
- The shared cells are `(0, 0)` and `(1, 0)`.
- The sum is `4 + (-1) = 3`, which is the maximum possible.

### Constraints:
- `m` == `grid.length`
- `n` == `grid[i].length`
- 2 <= `m`, `n` <= 1000
- 4 <= `m * n` <= 5 * 10<sup>5</sup>
- -100 <= `grid[i][j]` <= 100

## Question 4:
*Count Non Adjacent Subsets in a Rooted Tree*
### Difficulty: Hard
### Points: 6
### Description:
You are given a rooted tree with n nodes labeled from 0 to n - 1, represented by an integer array parent of length n, where:
- `parent[0]` = -1 (node 0 is the root).
- For each 1 <= i < n, `parent[i]` is the parent of node i (0 <= `parent[i]` < i).

You are also given an integer array nums of length n, where nums[i] is the value of node i, and an integer k.

A non-empty subset of nodes is called valid if:
- The sum of the values of the selected nodes is divisible by k.
- No two selected nodes are adjacent in the tree (no node and its direct parent are both included in the subset).

Return the number of valid subsets modulo 10<sup>9</sup> + 7.

### Examples:
#### Example 1:

Input: parent = `[-1,0,1]`, nums = `[1,2,3]`, k = 3

Output: 1

Explanation:

​​​​​​​<img src="https://assets.leetcode.com/uploads/2025/07/17/image1.png"><br>

The only valid subset is `{2}`. It contains node 2 with value 3, which is divisible by 3.

#### Example 2:

Input: parent = `[-1,0,0,0]`, nums = `[2,1,2,1]`, k = 3

Output: 2

Explanation:

​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2025/07/17/image2.png"><br>

The valid subsets are:
- `{1, 2}`: Nodes 1 and 2 are both children of node 0 and not directly connected to each other. Their values sum to 1 + 2 = 3, which is divisible by 3.
- `{2, 3}`: Nodes 2 and 3 are also non-adjacent. Their values sum to 2 + 1 = 3, which is divisible by 3.

No other subset satisfies both conditions. Therefore, the answer is 2.

### Constraints:
- `n` == `parent.length` == `nums.length`
- 1 <= `n` <= 1000
- `parent[0]` == -1
- For all 1 <= i < n:
    - 0 <= `parent[i]` < i
- 1 <= `nums[i]` <= 10<sup>9</sup>
- 1 <= `k` <= 100​​​​​​​​​​​​​​​​​​​​​
- parent describes a valid rooted tree.



