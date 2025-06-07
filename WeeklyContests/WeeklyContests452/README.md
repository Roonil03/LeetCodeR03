# Weekly Contest 452
## Ranking: 1262 / 28356
## Question 1:
*Partition Array into Two Equal Product Subsets*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums containing distinct positive integers and an integer target.

Determine if you can partition nums into two non-empty disjoint subsets, with each element belonging to exactly one subset, such that the product of the elements in each subset is equal to target.

Return true if such a partition exists and false otherwise.

A subset of an array is a selection of elements of the array.

### Testcases:
#### Example 1:

Input: nums = `[3,1,6,8,4]`, target = 24

Output: true

Explanation: The subsets `[3, 8]` and `[1, 6, 4]` each have a product of 24. Hence, the output is true.

#### Example 2:

Input: nums = `[2,5,3,7]`, target = 15

Output: false

Explanation: There is no way to partition nums into two non-empty disjoint subsets such that both subsets have a product of 15. Hence, the output is false.

### Constraints:

- 3 <= `nums.length` <= 12
- 1 <= `target` <= 10<sup>15</sup>
- 1 <= `nums[i]` <= 100
- All elements of nums are distinct.

## Question 2:
*Minimum Absolute Difference in Sliding Submatrix*
### Difficulty: Medium
### Points: 4
### Description:
You are given an `m x n` integer matrix grid and an integer k.

For every contiguous `k x k` submatrix of grid, compute the minimum absolute difference between any two distinct values within that submatrix.

Return a 2D array ans of size `(m - k + 1)` x `(n - k + 1)`, where `ans[i][j]` is the minimum absolute difference in the submatrix whose top-left corner is `(i, j)` in grid.

Note: If all elements in the submatrix have the same value, the answer will be 0.

A submatrix (x<sub>1</sub>, y<sub>1</sub>, x<sub>2</sub>, y<sub>2</sub>) is a matrix that is formed by choosing all cells `matrix[x][y]` where x<sub>1</sub> <= x <= x<sub>2</sub> and y<sub>1</sub> <= y <= y<sub>2</sub>.

### Testcases:
#### Example 1:

Input: grid = `[[1,8],[3,-2]]`, k = 2

Output: `[[2]]`

Explanation:

- There is only one possible `k x k` submatrix: `[[1, 8], [3, -2]]`.
- Distinct values in the submatrix are `[1, 8, 3, -2]`.
- The minimum absolute difference in the submatrix is `|1 - 3|` = 2. Thus, the answer is `[[2]]`.
#### Example 2:

Input: grid = `[[3,-1]]`, k = 1

Output: `[[0,0]]`

Explanation:

- Both k x k submatrix has only one distinct element.
- Thus, the answer is [[0, 0]].
#### Example 3:

Input: grid = `[[1,-2,3],[2,3,5]]`, k = 2

Output: `[[1,2]]`

Explanation:

- There are two possible k × k submatrix:
    - Starting at (0, 0): `[[1, -2], [2, 3]]`.
        - Distinct values in the submatrix are `[1, -2, 2, 3]`.
        - The minimum absolute difference in the submatrix is `|1 - 2|` = 1.
    - Starting at (0, 1): `[[-2, 3], [3, 5]]`.
        - Distinct values in the submatrix are `[-2, 3, 5]`.
        - The minimum absolute difference in the submatrix is `|3 - 5|` = 2.
- Thus, the answer is `[[1, 2]]`.

### Constraints:

- 1 <= `m` == `grid.length` <= 30
- 1 <= `n` == `grid[i].length` <= 30
- -10<sup>5</sup> <= `grid[i][j]` <= 10<sup>5</sup>
- 1 <= k <= `min(m, n)`

## Question 3: 
*Minimum Moves to Clean the Classroom*
## Difficulty: Medium
## Points: 5
### Description:
You are given an m x n grid classroom where a student volunteer is tasked with cleaning up litter scattered around the room. Each cell in the grid is one of the following:

- '`S`': Starting position of the student
- '`L`': Litter that must be collected (once collected, the cell becomes empty)
- '`R`': Reset area that restores the student's energy to full capacity, regardless of their current energy level (can be used multiple times)
- '`X`': Obstacle the student cannot pass through
- `'.':` Empty space

You are also given an integer energy, representing the student's maximum energy capacity. The student starts with this energy from the starting position '`S`'.

Each move to an adjacent cell (up, down, left, or right) costs 1 unit of energy. If the energy reaches 0, the student can only continue if they are on a reset area '`R`', which resets the energy to its maximum capacity energy.

Return the minimum number of moves required to collect all litter items, or -1 if it's impossible.

### Testcases:
#### Example 1:

Input: classroom = `["S.", "XL"]`, energy = 2

Output: 2

Explanation:

- The student starts at cell (0, 0) with 2 units of energy.
- Since cell (1, 0) contains an obstacle '`X`', the student cannot move directly downward.
- A valid sequence of moves to collect all litter is as follows:
    - Move 1: From (0, 0) → (0, 1) with 1 unit of energy and 1 unit remaining.
    - Move 2: From (0, 1) → (1, 1) to collect the litter '`L`'.
- The student collects all the litter using 2 moves. Thus, the output is 2.
#### Example 2:

Input: classroom = `["LS", "RL"]`, energy = 4

Output: 3

Explanation:

- The student starts at cell (0, 1) with 4 units of energy.
- A valid sequence of moves to collect all litter is as follows:
    - Move 1: From (0, 1) → (0, 0) to collect the first litter '`L`' with 1 unit of energy used and 3 units remaining.
    - Move 2: From (0, 0) → (1, 0) to 'R' to reset and restore energy back to 4.
    - Move 3: From (1, 0) → (1, 1) to collect the second litter '`L`'.
- The student collects all the litter using 3 moves. Thus, the output is 3.
#### Example 3:

Input: classroom = `["L.S", "RXL"]`, energy = 3

Output: -1

Explanation:

- No valid path collects all '`L`'.

### Constraints:

- 1 <= `m` == `classroom.length` <= 20
- 1 <= `n` == `classroom[i].length` <= 20
- `classroom[i][j]` is one of '`S`', '`L`', '`R`', '`X`', or '`.`'
- 1 <= `energy` <= 50
- There is exactly one '`S`' in the grid.
- There are at most 10 '`L`' cells in the grid.

## Question 4:
*Maximize Count of Distinct Primes After Split*
## Difficulty: Hard
## Points: 7
### Description:
You are given an integer array `nums` containing distinct positive integers and an integer target.

Determine if you can partition `nums` into two non-empty disjoint subsets, with each element belonging to exactly one subset, such that the product of the elements in each subset is equal to target.

Return true if such a partition exists and false otherwise.

A subset of an array is a selection of elements of the array.

### Testcases:
#### Example 1:

Input: nums = `[3,1,6,8,4]`, target = 24

Output: true

Explanation: The subsets `[3, 8]` and `[1, 6, 4]` each have a product of 24. Hence, the output is true.

#### Example 2:

Input: nums = `[2,5,3,7]`, target = 15

Output: false

Explanation: There is no way to partition nums into two non-empty disjoint subsets such that both subsets have a product of 15. Hence, the output is false.

### Constraints:

- 3 <= `nums.length` <= 12
- 1 <= `target` <= 10<sup>15</sup>
- 1 <= `nums[i]` <= 100
- All elements of `nums` are distinct.