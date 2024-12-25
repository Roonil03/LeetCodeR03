# 887. Super Egg Drop
## Question Level: Hard
### Description:
You are given `k` identical eggs and you have access to a building with `n` floors labeled from 1 to n.

You know that there exists a floor f where `0 <= f <= n` such that any egg dropped at a floor higher than f will break, and any egg dropped at or below floor f will not break.

Each move, you may take an unbroken egg and drop it from any floor x (where `1 <= x <= n`). If the egg breaks, you can no longer use it. However, if the egg does not break, you may reuse it in future moves.

Return the minimum number of moves that you need to determine with certainty what the value of f is.

### Examples:
#### Example 1:

Input: k = 1, n = 2<br>
Output: 2<br>
Explanation: <br>
Drop the egg from floor 1. If it breaks, we know that f = 0.<br>
Otherwise, drop the egg from floor 2. If it breaks, we know that f = 1.<br>
If it does not break, then we know f = 2.<br>
Hence, we need at minimum 2 moves to determine with certainty what the value of f is.
#### Example 2:

Input: k = 2, n = 6<br>
Output: 3<br>
### Example 3:

Input: k = 3, n = 14<br>
Output: 4<br>

### Constraints:

- 1 <= k <= 100
- 1 <= n <= 10^4

### <i>Concepts Used:
- Math
- Binary Search
- Dynamic Programming</i>