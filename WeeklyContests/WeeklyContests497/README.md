# Weekly Contest 497
## Ranking:
## Question 1:
*Find the Degree of Each Vertex*
### Difficulty: Easy
### Points: 3
### Description:
You are given a 2D integer array matrix of size `n x n` representing the adjacency matrix of an undirected graph with n vertices labeled from 0 to `n - 1`.
- `matrix[i][j]` = 1 indicates that there is an edge between vertices i and j.
- `matrix[i][j]` = 0 indicates that there is no edge between vertices i and j.

The degree of a vertex is the number of edges connected to it.

Return an integer array ans of size n where `ans[i]` represents the degree of vertex i.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2026/03/14/g41f.png">

Input: matrix = `[[0,1,1],[1,0,1],[1,1,0]]`

Output: `[2,2,2]`

Explanation:
- Vertex 0 is connected to vertices 1 and 2, so its degree is 2.
- Vertex 1 is connected to vertices 0 and 2, so its degree is 2.
- Vertex 2 is connected to vertices 0 and 1, so its degree is 2.

Thus, the answer is `[2, 2, 2]`.

#### Example 2:

<img src="https://assets.leetcode.com/uploads/2026/03/14/g42f.png">

Input: matrix = `[[0,1,0],[1,0,0],[0,0,0]]`

Output: `[1,1,0]`

Explanation:
- Vertex 0 is connected to vertex 1, so its degree is 1.
- Vertex 1 is connected to vertex 0, so its degree is 1.
- Vertex 2 is not connected to any vertex, so its degree is 0.

Thus, the answer is `[1, 1, 0]`.

#### Example 3:

Input: matrix = `[[0]]`

Output: `[0]`

Explanation:

There is only one vertex and it has no edges connected to it. Thus, the answer is `[0]`.

### Constraints:

- 1 <= `n` == `matrix.length` == `matrix[i].length` <= 100​​​​​​​
- ​​​​​​`​matrix[i][i]` == 0
- `matrix[i][j]` is either 0 or 1
- `matrix[i][j]` == `matrix[j][i]`

## Question 2:
*Angles of a Triangle*
### Difficulty: Medium
### Points: 4
### Description:
You are given a positive integer array sides of length 3.

Determine if there exists a triangle with positive area whose three side lengths are given by the elements of sides.

If such a triangle exists, return an array of three floating-point numbers representing its internal angles (in degrees), sorted in non-decreasing order. Otherwise, return an empty array.

Answers within 10<sup>-5</sup> of the actual answer will be accepted.

### Example:
#### Example 1:

Input: sides = `[3,4,5]`

Output: `[36.86990,53.13010,90.00000]`

Explanation:

You can form a right-angled triangle with side lengths 3, 4, and 5. The internal angles of this triangle are approximately 36.869897646, 53.130102354, and 90 degrees respectively.

#### Example 2:

Input: sides = `[2,4,2]`

Output: `[]`

Explanation:

You cannot form a triangle with positive area using side lengths 2, 4, and 2.

### Constraints:

- `sides.length` == 3
- 1 <= `sides[i]` <= 1000

## Question 3:
*Longest Balanced Substring After One Swap*
### Difficulty: Medium
### Points: 5
### Description:
You are given a binary string s consisting only of characters '`0`' and '`1`'.

A string is balanced if it contains an equal number of '`0`'s and '`1`'s.

You can perform at most one swap between any two characters in s. Then, you select a balanced substring from s.

Return an integer representing the maximum length of the balanced substring you can select.

A substring is a contiguous sequence of characters within a string.

### Examples:
#### Example 1:

Input: s = "`100001`"

Output: 4

Explanation:
- Swap "`100001`". The string becomes "`101000`".
- Select the substring "`101000`", which is balanced because it has two '`0`'s and two '`1`'s.
#### Example 2:

Input: s = "`111`"

Output: 0

Explanation:
- Choose not to perform any swaps.
- Select the empty substring, which is balanced because it has zero '`0`'s and zero '`1`'s.

### Constraints:

- 1 <= `s.length` <= 10<sup>5</sup>
- s consists only of the characters '`0`' and '`1`'.

## Question 4:
*Good Subsequence Queries*
### Difficulty: Hard
### Points: 7
### Description:
You are given an integer array nums of length n and an integer p.

A non-empty subsequence of nums is called good if:
- Its length is strictly less than n.
- The greatest common divisor (GCD) of its elements is exactly p.

You are also given a 2D integer array queries of length q, where each `queries[i]` = `[ind`<sub>`i`</sub>`, val`<sub>`i`</sub>`]` indicates that you should update `nums[ind`<sub>`i`</sub>`]` to `val`<sub>`i`</sub>.

After each query, determine whether there exists any good subsequence in the current array.

Return the number of queries for which a good subsequence exists.

A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

The term `gcd(a, b)` denotes the greatest common divisor of a and b.

### Examples:
#### Example 1:

Input: nums = `[4,8,12,16]`, p = 2, queries = `[[0,3],[2,6]]`

Output: 1

Explanation:  
| i | [ind<sub>i</sub>, val<sub>i</sub>] | Operation | Updated nums | Any good Subsequence |
|---|---|---|---|---|
| 0 | [0, 3] | Update `nums[0]` to `3` | [3, 8, 12, 16] | No, as no subsequence has GCD exactly `p = 2` |
| 1 | [2, 6] | Update `nums[2]` to `6` | [3, 8, 6, 16] | Yes, subsequence `[8, 6]` has GCD exactly `p = 2` |

Thus, the answer is 1.

#### Example 2:
Input: nums = `[4,5,7,8]`, p = 3, queries = `[[0,6],[1,9],[2,3]]`

Output: 2

Explanation:  
| i | [ind<sub>i</sub>, val<sub>i</sub>] | Operation | Updated nums | Any good Subsequence |
|---|---|---|---|---|
| 0 | [0, 6] | Update `nums[0]` to `6` | [6, 5, 7, 8] | No, as no subsequence has GCD exactly `p = 3` |
| 1 | [1, 9] | Update `nums[1]` to `9` | [6, 9, 7, 8] | Yes, subsequence `[6, 9]` has GCD exactly `p = 3` |
| 2 | [2, 3] | Update `nums[2]` to `3` | [6, 9, 3, 8] | Yes, subsequence `[6, 9, 3]` has GCD exactly `p = 3` |

Thus, the answer is 2.

#### Example 3:

Input: nums = `[5,7,9]`, p = 2, queries = `[[1,4],[2,8]]`

Output: 0

Explanation:  
| i | [ind<sub>i</sub>, val<sub>i</sub>] | Operation | Updated nums | Any good Subsequence |
|---|---|---|---|---|
| 0 | [1, 4] | Update `nums[1]` to `4` | [5, 4, 9] | No, as no subsequence has GCD exactly `p = 2` |
| 1 | [2, 8] | Update `nums[2]` to `8` | [5, 4, 8] | No, as no subsequence has GCD exactly `p = 2` |

Thus, the answer is 0.

### Constraints:

- 2 <= `n` == `nums.length` <= 5 * 10<sup>4</sup>
- 1 <= `nums[i]` <= 5 * 10<sup>4</sup>
- 1 <= `queries.length` <= 5 * 10<sup>4</sup>
- `queries[i]` = `[ind`<sub>`i`</sub>`, val`<sub>`i`</sub>`]`
- 1 <= `val`<sub>`i`</sub>, p <= 5 * 10<sup>4</sup>
- 0 <= `ind`<sub>`i`</sub> <= n - 1