# Biweekly Contest 146
## Ranking: 5227
## Question 1:
<i>Count Subarrays of Length Three With a Condition</i>
### Difficulty: Easy
### Points: 3
### Description:
Given an integer array nums, return the number of subarrays of length 3 such that the sum of the first and third numbers equals exactly half of the second number.

A subarray is a contiguous non-empty sequence of elements within an array.

### Testcases:
#### Example 1:

Input: nums = `[1,2,1,4,1]`<br>
Output: 1<br>
Explanation:<br>
Only the subarray `[1,4,1]` contains exactly 3 elements where the sum of the first and third numbers equals half the middle number.

#### Example 2:
Input: nums = `[1,1,1]`<br>
Output: 0<br>
Explanation:<br>
`[1,1,1]` is the only subarray of length 3. However, its first and third numbers do not add to half the middle number.


### Constraints:

- 3 <= `nums.length` <= 100
- -100 <= `nums[i]` <= 100

## Question 2:
<i>Count Paths With the Given XOR Value</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given a 2D integer array grid with size `m x n`. You are also given an integer k.

Your task is to calculate the number of paths you can take from the top-left cell `(0, 0)` to the bottom-right cell `(m - 1, n - 1)` satisfying the following constraints:
- You can either move to the right or down. Formally, from the cell (i, j) you may move to the cell `(i, j + 1)` or to the cell `(i + 1, j)` if the target cell exists.
- The XOR of all the numbers on the path must be equal to k.

Return the total number of such paths.

Since the answer can be very large, return the result modulo `10^9 + 7`.


### Testcases:
#### Example 1:
Input: grid = `[[2, 1, 5], [7, 10, 0], [12, 6, 4]]`, k = 11<br>
Output: 3<br>
Explanation: <br>
The 3 paths are:<br>
```
(0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2)
(0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
(0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
```
#### Example 2:
Input: grid = `[[1, 3, 3, 3], [0, 3, 3, 2], [3, 0, 1, 1]]`, k = 2<br>
Output: 5<br>
Explanation:<br>
The 5 paths are:<br>
```
(0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3)
(0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2) → (2, 3)
(0, 0) → (1, 0) → (1, 1) → (1, 2) → (1, 3) → (2, 3)
(0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2) → (2, 3)
(0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2) → (2, 3)
```
#### Example 3:
Input: grid = `[[1, 1, 1, 2], [3, 0, 3, 2], [3, 0, 2, 2]]`, k = 10<br>
Output: 0<br>

### Constraints:

- 1 <= m == `grid.length` <= 300
- 1 <= n == `grid[r].length` <= 300
- 0 <= `grid[r][c]` < 16
- 0 <= k < 16

## Question 3:
<i>Check if Grid can be Cut into Sections</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer n representing the dimensions of an `n x n` grid, with the origin at the bottom-left corner of the grid. You are also given a 2D array of coordinates rectangles, where rectangles[i] is in the form `[startx, starty, endx, endy]`, representing a rectangle on the grid. Each rectangle is defined as follows:
- `(startx, starty)`: The bottom-left corner of the rectangle.
- `(endx, endy)`: The top-right corner of the rectangle.

Note that the rectangles do not overlap. Your task is to determine if it is possible to make either two horizontal or two vertical cuts on the grid such that:
- Each of the three resulting sections formed by the cuts contains at least one rectangle.
- Every rectangle belongs to exactly one section.

Return true if such cuts can be made; otherwise, return false.

### Testcases:
#### Example 1:
Input: n = 5, rectangles = `[[1,0,5,2],[0,2,2,4],[3,2,5,3],[0,4,4,5]]`<br>
Output: true<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/10/23/tt1drawio.png"><br>
The grid is shown in the diagram. We can make horizontal cuts at y = 2 and y = 4. Hence, output is true.

#### Example 2:
Input: n = 4, rectangles = `[[0,0,1,1],[2,0,3,4],[0,2,2,3],[3,0,4,3]]`<br>
Output: true<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/10/23/tc2drawio.png">
We can make vertical cuts at x = 2 and x = 3. Hence, output is true.

#### Example 3:
Input: n = 4, rectangles = `[[0,2,2,4],[1,0,3,2],[2,2,3,4],[3,0,4,2],[3,2,4,4]]`<br>
Output: false<br>
Explanation:<br>
We cannot make two horizontal or two vertical cuts that satisfy the conditions. Hence, output is false.

### Constraints:

- 3 <= `n` <= 10^9
- 3 <= `rectangles.length` <= 10^5
- 0 <= `rectangles[i][0]` < `rectangles[i][2]` <= n
- 0 <= `rectangles[i][1]` < `rectangles[i][3]` <= n
- No two rectangles overlap.

## Question 4:
<i>Subsequences with a Unique Middle Mode I</i>
### Difficulty: Hard
### Points: 7
### Description:
Given an integer array nums, find the number of subsequences of size 5 of nums with a unique middle mode.

Since the answer may be very large, return it modulo `10^9 + 7`.

A mode of a sequence of numbers is defined as the element that appears the maximum number of times in the sequence.

A sequence of numbers contains a unique mode if it has only one mode.

A sequence of numbers seq of size 5 contains a unique middle mode if the middle element (`seq[2]`) is a unique mode.

A subsequence is a non-empty array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

### Testcases:
#### Example 1:

Input: nums = `[1,1,1,1,1,1]`
Output: 6
Explanation:
`[1, 1, 1, 1, 1]` is the only subsequence of size 5 that can be formed, and it has a unique middle mode of 1. This subsequence can be formed in 6 different ways, so the output is 6. 

#### Example 2:
Input: nums = [1,2,2,3,3,4]
Output: 4
Explanation:
`[1, 2, 2, 3, 4]` and `[1, 2, 3, 3, 4]` each have a unique middle mode because the number at index 2 has the greatest frequency in the subsequence. `[1, 2, 2, 3, 3]` does not have a unique middle mode because 2 and 3 appear twice.

#### Example 3:
Input: nums = `[0,1,2,3,4,5,6,7,8]`
Output: 0
Explanation:
There is no subsequence of length 5 with a unique middle mode.

### Constraints:

- 5 <= `nums.length` <= 1000
- -10^9 <= `nums[i]` <= 10^9

