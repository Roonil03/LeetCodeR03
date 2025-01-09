# 77. Combinations
## Question Level: Medium
### Description:
Given two integers n and k, return all possible combinations of k numbers chosen from the range `[1, n]`.

You may return the answer in any order.

### Examples:
#### Example 1:

Input: n = 4, k = 2<br>
Output: `[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]`<br>
Explanation: There are 4 choose 2 = 6 total combinations.<br>
Note that combinations are unordered, i.e., `[1,2]` and `[2,1]` are considered to be the same combination.
#### Example 2:

Input: n = 1, k = 1<br>
Output: `[[1]]`<br>
Explanation: There is 1 choose 1 = 1 total combination.

### Constraints:

- 1 <= n <= 20
- 1 <= k <= n

### <i>Concepts Used:
- Backtracking </i>