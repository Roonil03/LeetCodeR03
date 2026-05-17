# Weekly Contest 502
## Ranking:
## Question 1:
*Check Adjacent Digit Differences*
### Difficulty: Easy
### Points: 3
### Description:
You are given a string s consisting of digits.

Return true if the absolute difference between every pair of adjacent digits is at most 2, otherwise return false.

The absolute difference between a and b is defined as `abs(a - b)`.

### Examples:
#### Example 1:

Input: s = "`132`"

Output: true

Explanation:
- The absolute difference between digits at `s[0]` and `s[1]` is `abs(1 - 3)` = 2.
- The absolute difference between digits at `s[1]` and `s[2]` is `abs(3 - 2)` = 1.  
Since both differences are at most 2, the answer is true.
#### Example 2:

Input: s = "`129`"

Output: false

Explanation:
- The absolute difference between digits at `s[0]` and `s[1]` is `abs(1 - 2)` = 1.
- The absolute difference between digits at `s[1]` and `s[2]` is `abs(2 - 9)` = 7, which is greater than 2.  
Therefore, the answer is false.

### Constraints:

- 2 <= `s.length` <= 100
- `s` consists only of digits.

## Question 2:
*Count K-th Roots in a Range*
### Difficulty: Medium
### Points: 4
### Description:
You are given three integers l, r, and k.

An integer y is said to be a perfect kth power if there exists an integer x such that y = x<sup>k</sup>.

Return the number of integers y in the range `[l, r]` (inclusive) that are perfect k<sup>th</sup> powers.

### Examples:
#### Example 1:

Input: l = 1, r = 9, k = 3

Output: 2

Explanation:
The perfect cubes in the range `[1, 9]` are:
- 1 = 1<sup>3</sup>
- 8 = 2<sup>3</sup>  
Hence, the answer is 2.
#### Example 2:

Input: l = 8, r = 30, k = 2

Output: 3

Explanation:

The perfect squares in the range `[8, 30]` are:
- 9 = 3<sup>2</sup>
- 16 = 4<sup>2</sup>
- 25 = 5<sup>2</sup>  
Hence, the answer is 3.

### Constraints:

- 0 <= l <= r <= 10<sup>9</sup>
- 1 <= k <= 30

## Question 3:
*Largest Local Values in a Matrix II*
### Difficulty: Medium
### Points: 5
### Description:
You are given an n x m integer matrix matrix containing non-negative integers.

A non-zero cell (row, col) checks the cells near it as follows:
- Let x =` matrix[row][col]`.
- Consider every cell within x rows and x columns of `(row, col)`.
- Ignore cells that are outside the matrix.
- Ignore the cells where both the row distance and column distance are exactly x.

The cell `(row, col)` is a local maximum if it is non-zero and no considered cell has a value greater than x.

Return an integer denoting the number of local maximums in matrix.

### Examples:
#### ​​​​​​​Example 1:

Input: matrix = `[[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,2,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0]]`

Output: 1

​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/13/chatgpt-image-may-14-2026-01_53_19-am.png"><br>

Explanation:
- For the non-zero cell (3, 3), x = matrix[3][3] = 2.
- The highlighted cells are the considered cells within x rows and x columns of (3, 3).
- The four cells with both row and column distances equal to x = 2 are ignored.
- No considered cell has a value greater than 2, so (3, 3) is a local maximum.
- There are no other non-zero cells, so the answer is 1.
#### Example 2:

Input: matrix = `[[1,2],[3,4]]`

Output: 1

Explanation:

Only the cell with value 4 is a local maximum. Every other non-zero cell considers a cell with a greater value.

#### Example 3:

Input: matrix = `[[1,0,1],[0,1,0],[1,0,1]]`

Output: 5

Explanation:
- For a cell with value 1, the considered cells are the cell itself and its 4-directionally adjacent cells that are inside the matrix.
- Each of the five cells with value 1 only considers cells with values 0 or 1, so all five of them are local maximums.
#### Example 4:

Input: matrix = `[[1,1],[1,1]]`

Output: 4

Explanation:

All cells have the same value. Therefore, no cell considers another cell with a greater value, so all 4 cells are local maximums.

### Constraints:

- 1 <= `n` == `matrix.length` <= 200
- 1 <= `m` == `matrix[i].length` <= 200
- 0 <= `matrix[i][j]` <= 200

## Question 4
*Smallest Unique Subarray*
### Difficulty: Hard
### Points: 6
### Description:
You are given an integer array nums.

Find the minimum length of a subarray that is not identical to any other subarray in nums.

Return an integer denoting the minimum possible length of such a subarray.

A subarray is a contiguous non-empty sequence of elements within an array.

Two subarrays are considered identical if they have the same length and the same elements in corresponding positions.

### Examples:
#### Example 1:
Input: nums = `[3,3,3]`

Output: 3

Explanation:
- Subarrays of length 1: `[3]` → appears 3 times
- Subarrays of length 2: `[3, 3]` → appears 2 times
- Subarrays of length 3: `[3, 3, 3]` → appears once  
The subarray `[3, 3, 3]` is unique, so the smallest unique subarray length is 3.

#### Example 2:

Input: nums = `[2,1,2,3,3]`

Output: 1

Explanation:

Subarrays of length 1:
- `[2]` → appears 2 times
- `[1]` → appears once
- `[3]` → appears 2 times  
The subarray `[1]` is unique, so the smallest unique subarray length is 1.
#### Example 3:

Input: nums = `[1,1,2,2,1]`

Output: 2

Explanation:

Subarrays of length 1:
- `[1]` → appears 3 times
- `[2]` → appears 2 times

Subarrays of length 2:
- `[1, 1]` → appears once
- `[1, 2]` → appears once
- `[2, 2]` → appears once
- `[2, 1]` → appears once  

There is at least one subarray of length 2 that is unique, so the smallest unique subarray length is 2.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>5</sup>