# Weekly Contest 511
## Ranking: 1718 / 40827
## Question 1:
*Even Number of Knight Moves*
### Difficulty: Easy
### Points: 3
### Description:
You are given two integer arrays start and target, where each array is of the form `[x, y]` representing a cell on a standard 8 x 8 chessboard.

Return true if a knight can move from start to target in an even number of moves. Otherwise, return false.

Note: A valid knight move consists of moving two squares in one direction and one square perpendicular to it. The figure below illustrates all eight possible moves from a cell.

<img src="https://assets.leetcode.com/uploads/2018/10/12/knight.png"><br>

### Examples:
#### Example 1:

Input: start = `[1,1]`, target = `[2,2]`

Output: true

Explanation:

One possible sequence of moves is `(1, 1) -> (3, 2) -> (2, 4) -> (4, 3) -> (2, 2)`.

The knight reaches the target in 4 moves, which is even. Thus, the answer is true.

#### Example 2:

Input: start = `[4,5]`, target = `[6,6]`

Output: false

Explanation:​​​​​​​

It is impossible to reach target = `[6, 6]` from start = `[4, 5]` in an even number of moves. Thus, the answer is false.

### Constraints:

- `start.length` == `target.length` == 2
- 0 <= `start[i]`, `target[i]` <= 7

## Question 2:
*Count Dominant Nodes in a Binary Tree*
### Difficulty: Medium
### Points: 4
### Description:
You are given the root of a complete binary tree.

A node x is called dominant if its value is equal to the maximum value among all nodes in the subtree rooted at x.

Return the number of dominant nodes in the tree.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2026/06/13/tnew.png"><br>

Input: root = `[5,3,8,2,4,7,1]`

Output: 5

Explanation:

- The leaf nodes with values 2, 4, 7, and 1 are dominant.
- The node with value 8 is dominant because its value is the maximum value in its subtree `[8, 7, 1]`.
- Thus, the answer is 5.
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2026/06/15/t9.png"><br>

Input: root = `[1,2,3,1,2]`

Output: 4

Explanation:

- The leaf nodes with values 1, 2, and 3 are dominant.
- The node with value 2 whose subtree is `[2, 1, 2]` is dominant because its value is the maximum value in its subtree.
- Thus, the answer is 4.

### Constraints:

- The number of nodes in the tree is in the range `[1, 10`<sup>`5`</sup>`]`.
- 1 <= `Node.val` <= 10<sup>9</sup>
- The tree is guaranteed to be a complete binary tree.

## Question 3: 
*Transform Binary String Using Subsequence Sort*
### Difficulty: Medium
### Points: 5
### Description:
You are given a binary string s.

You are also given an array of strings `strs`, where each `strs[i]` has the same length as s and consists of characters '`0`', '`1`', and '`?`'. Each '`?`' can be replaced by either '`0`' or '`1`'.

You may perform the following operation any number of times (including zero):
- Choose any subsequence sub of s.
- Sort sub in non-decreasing order.
- Replace the chosen subsequence in s with the sorted sub, keeping all other characters unchanged.
- Return a boolean array ans, where `ans[i]` is true if it's possible to replace all '`?`' in `strs[i]` with '`0`' or '`1`' and transform s into the resulting string using the allowed operation above, otherwise return false.

### Examples:
#### Example 1:

Input: s = "101", strs = `["1?1","0?1","0?0"]`

Output: `[true,true,false]`

Explanation:

| `i` | `strs[i]` | Replacement | Result | `strs[i]` | Operation(s) | Result |
| :-- | :-- | :-- | :-- | :-- | :-- | :-- |
| 0 | "`1?1`" | ? → 0 | "`101`" | Matches s. | true |
| 1 | "`0?1`" | ? → 1 | "`011`" | Select the subsequence at indices `[0..2]` of s → "`101`". Sort "`101`" to get "`011`" = `strs[i]`. | true |
| 2 | "`0?0`" | ? → 0 or 1 | "`000`" or "`010`" | Not feasible. | false |
Thus, ans = `[true, true, false]`.

#### Example 2:

Input: s = "`1100`", strs = `["0011","11?1","1?1?"]`

Output: `[true,false,true]`

Explanation:

| `i` | `strs[i]` | Replacement | Result | `strs[i]` | Operation(s) | Result |
| :-- | :-- | :-- | :-- | :-- | :-- | :-- |
| 0 | "`0011`" | - | "`0011`" | Select the subsequence at indices `[0..3]` of s → "`1100`". Sort "`1100`" to get "`0011`" = `strs[i]`. | true |
| 1 | "`11?1`" | ? → 0 | "`1101`" | Not feasible. | false |
| 2 | "`1?1?`" | First `? → 0`. Second` ? → 0` | "`1010`" | Select the subsequence at indices `[1, 2]` of s → "`10`". Sort "`10`" to get "`01`", so s = "`1010`". | true |
Thus, ans = `[true, false, true]`.

#### Example 3:

Input: s = "`1010`", strs = `["0011"]`

Output: `[true]`

Explanation:

| `i` | `strs[i]` | Replacement | Result | `strs[i]` | Operation(s) | Result |
| :-- | :-- | :-- | :-- | :-- | :-- | :-- |
| 0 | "`0011`" | - | "`0011`" | Select the subsequence at indices `[0, 2, 3]` of s → "`110`". Sort "`110`" to get "`011`", so s = "`0011`" = `strs[i]`. | true |
Thus, ans = `[true]`.

### Constraints:

- 1 <= `n` == `s.length` <= 2000
- `s[i]` is either '`0`' or '`1`'.
- 1 <= `strs.length` <= 2000
- `strs[i].length` == n
- `strs[i]` is either '`0`', '`1`', or '`?`'​​​​​​​.

## Question 4: 
*Minimum Number of String Groups Through Transformations*
### Difficulty: Hard
### Points: 6
### Description:
You are given an array of strings words.

Define a transformation on a string s as follows:
- Let E be the subsequence of characters at even indices of s.
- Let O be the subsequence of characters at odd indices of s.
- Independently cyclically shift E and O by any number of positions to the right, possibly zero.
- Reconstruct the string by placing the shifted E characters back into even indices and the shifted O characters back into odd indices.

Two strings are equivalent if one can be transformed into the other by a single transformation.

Partition words into the minimum number of groups such that:
- Every string belongs to exactly one group.
- Every pair of strings in the same group are equivalent.

Return an integer denoting the minimum number of groups.

### Examples:
#### Example 1:
Input: words = `["ntgwz","zwntg"]`

Output: 1

Explanation:

- For "`ntgwz`", the even-index subsequence is "`ngz`" and the odd-index subsequence is "`tw`".
- Shift "`ngz`" right by 1 position to obtain "`zng`", and shift "`tw`" right by 1 position to obtain "`wt`".
- After reconstructing the string, we obtain "`zwntg`".
- Therefore, both strings are equivalent and belong to the same group.
#### Example 2:

Input: words = `["abc","cab","bac","acb","bca","cba"]`

Output: 3

Explanation:

The strings can be partitioned into the following groups:
- `["abc","cba"]`
- `["cab","bac"]`
- `["acb","bca"]`
#### Example 3:

Input: words = `["leet","abb","bab","deed","edde","code","bba"]`

Output: 5

Explanation:

The strings can be partitioned into the following groups:
- `["abb","bba"]`
- `["deed","edde"]`
- `["leet"]`
- `["bab"]`
- `["code"]`  
​​​​​​​​​​​​​​All pairs of strings in each group are equivalent.


### Constraints:

- 1 <= `words.length` <= 10<sup>5</sup>
- 1 <= `words[i].length` <= 5 * 10<sup>5</sup>
- The sum of `words[i].length` does not exceed 5 * 10<sup>5</sup>.
- `words[i]` consist of lowercase English letters.
 