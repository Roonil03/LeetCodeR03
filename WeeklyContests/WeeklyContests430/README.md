# Weekly Contests 430
## Rank: 5483 / 22451
## Question 1:
<i>Minimum Operations to Make Columns Strictly Increasing</i>
### Difficulty: Easy
### Points: 3
### Description:
You are given a `m x n` matrix grid consisting of non-negative integers.

In one operation, you can increment the value of any `grid[i][j]` by 1.

Return the minimum number of operations needed to make all columns of grid strictly increasing.

### Testcases:
#### Example 1:
Input: grid = `[[3,2],[1,3],[3,4],[0,1]]`<br>
Output: 15<br>
Explanation:<br>
- To make the 0th column strictly increasing, we can apply 3 operations on `grid[1][0]`, 2 operations on `grid[2][0]`, and 6 operations on `grid[3][0]`.
- To make the 1st column strictly increasing, we can apply 4 operations on `grid[3][1]`.

#### Example 2:
Input: grid = `[[3,2,1],[2,1,0],[1,2,3]]`<br>
Output: 12<br>
Explanation:
- To make the 0th column strictly increasing, we can apply 2 operations on `grid[1][0]`, and 4 operations on `grid[2][0]`.
- To make the 1st column strictly increasing, we can apply 2 operations on `grid[1][1]`, and 2 operations on `grid[2][1]`.
- To make the 2nd column strictly increasing, we can apply 2 operations on `grid[1][2]`.


### Constraints:

- m == `grid.length`
- n == `grid[i].length`
- 1 <= `m`, `n` <= 50
- 0 <= `grid[i][j]` < 2500

## Question 2:
<i>Find the Lexicographically Largest String From the Box I</i>
### Difficulty: Medium
### Points: 4
### Description:
You are given a string `word`, and an integer `numFriends`.

Alice is organizing a game for her `numFriends` friends. There are multiple rounds in the game, where in each round:
- `word` is split into `numFriends` non-empty strings, such that no previous round has had the exact same split.
- All the split words are put into a box.

Find the lexicographically largest string from the box after all the rounds are finished.

A string a is lexicographically smaller than a string b if in the first position where a and b differ, string a has a letter that appears earlier in the alphabet than the corresponding letter in b.<br>
If the first min(a.length, b.length) characters do not differ, then the shorter string is the lexicographically smaller one.

### Testcases:
#### Example 1:
Input: word = "`dbca`", numFriends = 2<br>
Output: "`dbc`"<br>
Explanation: <br>
All possible splits are:
- "`d`" and "`bca`".
- "`db`" and "`ca`".
- "`dbc`" and "`a`".
#### Example 2:
Input: word = "`gggg`", numFriends = 4<br>
Output: "`g`"<br>
Explanation: <br>
The only possible split is: "`g`", "`g`", "`g`", and "`g`".

## Constraints:

- 1 <= `word.length` <= 5 * 10^3
- `word` consists only of lowercase English letters.
- 1 <= `numFriends` <= `word.length`

## Question 3:
<i>Count Special Subsequences</i>
### Difficulty: Medium
### Points: 5
### Description:
You are given an array nums consisting of positive integers.

A special subsequence is defined as a subsequence of length 4, represented by indices `(p, q, r, s)`, where `p < q < r < s`. This subsequence must satisfy the following conditions:
- `nums[p]` * `nums[r]` == `nums[q]` * `nums[s]`
- There must be at least one element between each pair of indices. In other words, `q - p > 1`, `r - q > 1` and `s - r > 1`.

A subsequence is a sequence derived from the array by deleting zero or more elements without changing the order of the remaining elements.

Return the number of different special subsequences in nums.

### Testcases:
#### Example 1:
Input: nums = `[1,2,3,4,3,6,1]`<br>
Output: 1<br>
Explanation:<br>
There is one special subsequence in nums.
- `(p, q, r, s)` = `(0, 2, 4, 6)`:
    - This corresponds to elements `(1, 3, 3, 1)`.
    - `nums[p]` * `nums[r]` = `nums[0]` * `nums[4]` = 1 * 3 = 3
    - `nums[q]` * `nums[s]` = `nums[2]` * `nums[6]` = 3 * 1 = 3

#### Example 2:
Input: nums = `[3,4,3,4,3,4,3,4]`<br>
Output: 3<br>
Explanation:<br>
There are three special subsequences in nums.
- `(p, q, r, s)` = `(0, 2, 4, 6)`:
    - This corresponds to elements `(3, 3, 3, 3)`.
    - `nums[p]` * `nums[r]` = `nums[0]` * `nums[4]` = 3 * 3 = 9
    - `nums[q]` * `nums[s]` = `nums[2]` * `nums[6]` = 3 * 3 = 9
- `(p, q, r, s)` = `(1, 3, 5, 7)`:
    - This corresponds to elements `(4, 4, 4, 4)`.
    - `nums[p]` * `nums[r]` = `nums[1]` * `nums[5]` = 4 * 4 = 16
    - `nums[q]` * `nums[s]` = `nums[3]` * `nums[7]` = 4 * 4 = 16
- `(p, q, r, s)` = `(0, 2, 5, 7)`:
    - This corresponds to elements `(3, 3, 4, 4)`.
    - `nums[p]` * `nums[r]` = `nums[0]` * `nums[5]` = 3 * 4 = 12
    - `nums[q]` * `nums[s]` = `nums[2]` * `nums[7]` = 3 * 4 = 12

### Constraints:

- 7 <= `nums.length` <= 1000
- 1 <= `nums[i]`<= 1000

## Question 4:
<i>Count the Number of Arrays with K Matching Adjacent Elements</i>
### Difficulty: Hard
### Points: 6
### Description:
You are given three integers n, m, k. A good array arr of size n is defined as follows:
- Each element in arr is in the inclusive range `[1, m]`.
- Exactly k indices i (where 1 <= i < n) satisfy the condition `arr[i - 1]` == `arr[i]`.

Return the number of good arrays that can be formed.

Since the answer may be very large, return it modulo `10^9 + 7`

### Testcases:
#### Example 1:
Input: n = 3, m = 2, k = 1<br>
Output: 4<br>
Explanation:
- There are 4 good arrays. They are `[1, 1, 2]`, `[1, 2, 2]`, `[2, 1, 1]` and `[2, 2, 1]`.<br>
Hence, the answer is 4.

#### Example 2:
Input: n = 4, m = 2, k = 2<br>
Output: 6<br>
Explanation:
- The good arrays are `[1, 1, 1, 2]`, `[1, 1, 2, 2]`, `[1, 2, 2, 2]`, `[2, 1, 1, 1]`, `[2, 2, 1, 1]` and `[2, 2, 2, 1]`.<br>
Hence, the answer is 6.

#### Example 3:
Input: n = 5, m = 2, k = 0<br>
Output: 2<br>
Explanation:
- The good arrays are `[1, 2, 1, 2, 1]` and `[2, 1, 2, 1, 2]`.<br> Hence, the answer is 2.

### Constraints:

- 1 <= `n` <= 10^5
- 1 <= `m` <= 10^5
- 0 <= `k` <= `n - 1`