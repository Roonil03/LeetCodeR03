# Weekly Contest 454
## Ranking: 
## Question 1:
*Generate Tag for Video Caption*
### Description:
You are given a string caption representing the caption for a video.

The following actions must be performed in order to generate a valid tag for the video:

1. Combine all words in the string into a single camelCase string prefixed with '`#`'. A camelCase string is one where the first letter of all words except the first one is capitalized. All characters after the first character in each word must be lowercase.
2. Remove all characters that are not an English letter, except the first '`#`'.
3. Truncate the result to a maximum of 100 characters.

Return the tag after performing the actions on caption.

### Testcases:
#### Example 1:

Input: caption = "`Leetcode daily streak achieved`"

Output: "`#leetcodeDailyStreakAchieved`"

Explanation:

The first letter for all words except "`leetcode`" should be capitalized.

#### Example 2:

Input: caption = "`can I Go There`"

Output: "`#canIGoThere`"

Explanation:

The first letter for all words except "`can`" should be capitalized.

#### Example 3:

Input: caption = "`hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh`"

Output: "`#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh`"

Explanation:

Since the first word has length 101, we need to truncate the last two letters from the word.

### Constraints:

- 1 <= `caption.length` <= 150
- `caption` consists only of English letters and `' '`.

## Question 2:
*Count Special Triplets*
### Description:
You are given an integer array nums.

A special triplet is defined as a triplet of indices `(i, j, k)` such that:
- 0 <= i < j < k < n, where n = `nums.length`
- `nums[i]` == `nums[j]` * 2
- `nums[k]` == `nums[j]` * 2

Return the total number of special triplets in the array.

Since the answer may be large, return it modulo 10<sup>9</sup> + 7.

### Testcases:
#### Example 1:

Input: nums = `[6,3,6]`

Output: 1

Explanation:

The only special triplet is `(i, j, k)` = `(0, 1, 2)`, where:
- `nums[0]` = 6, `nums[1]` = 3, `nums[2]` = 6
- `nums[0]` = `nums[1]` * 2 = 3 * 2 = 6
- `nums[2]` = `nums[1]` * 2 = 3 * 2 = 6
#### Example 2:

Input: nums = `[0,1,0,0]`

Output: 1

Explanation:

The only special triplet is `(i, j, k)` = `(0, 2, 3)`, where:
- `nums[0]` = 0, `nums[2]` = 0, `nums[3]` = 0
- `nums[0]` = `nums[2]` * 2 = 0 * 2 = 0
- `nums[3]` = `nums[2]` * 2 = 0 * 2 = 0
#### Example 3:

Input: nums = `[8,4,2,8,4]`

Output: 2

Explanation:

There are exactly two special triplets:
- `(i, j, k)` = `(0, 1, 3)`
    - `nums[0]` = 8, `nums[1]` = 4, `nums[3]` = 8
    - `nums[0]` = `nums[1]` * 2 = 4 * 2 = 8
    - `nums[3]` = `nums[1]` * 2 = 4 * 2 = 8
- `(i, j, k)` = `(1, 2, 4)`
    - `nums[1]` = 4, `nums[2]` = 2, `nums[4]` = 4
    - `nums[1]` = `nums[2]` * 2 = 2 * 2 = 4
    - `nums[4]` = `nums[2]` * 2 = 2 * 2 = 4

### Constraints:

- 3 <= `n` == `nums.length` <= 10<sup>5</sup>
- 0 <= `nums[i]` <= 10<sup>5</sup>

## Question 3:
*Maximum Product of First and Last Elements of a Subsequence*
### Description:
You are given an integer array nums and an integer m.

Return the maximum product of the first and last elements of any subsequence of nums of size m.

A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

### Testcases:
#### Example 1:

Input: nums = `[-1,-9,2,3,-2,-3,1]`, m = 1

Output: 81

Explanation:

The subsequence `[-9]` has the largest product of the first and last elements: -9 * -9 = 81. Therefore, the answer is 81.

#### Example 2:

Input: nums = `[1,3,-5,5,6,-4]`, m = 3

Output: 20

Explanation:

The subsequence `[-5, 6, -4]` has the largest product of the first and last elements.

#### Example 3:

Input: nums = `[2,-1,2,-6,5,2,-5,7]`, m = 2

Output: 35

Explanation:

The subsequence `[5, 7]` has the largest product of the first and last elements.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- -10<sup>5</sup> <= `nums[i]` <= 10<sup>5</sup>
- 1 <= `m` <= `nums.length`

## Question 4:
*Find Weighted Median Node in Tree*
### Description:
You are given an integer n and an undirected, weighted tree rooted at node 0 with n nodes numbered from 0 to n - 1. This is represented by a 2D array edges of length n - 1, where `edges[i]` = `[u`<sup>`i`</sup>`, v`<sup>`i`</sup>`, w`<sup>`i`</sup>`]` indicates an edge from node u<sup>i</sup> to v<sup>i</sup> with weight w<sup>i</sup>.

The weighted median node is defined as the first node x on the path from u<sup>i</sup> to v<sup>i</sup> such that the sum of edge weights from ui to x is greater than or equal to half of the total path weight.

You are given a 2D integer array queries. For each `queries[j]` = `[u`<sup>`j`</sup>`, v`<sup>`j`</sup>`]`, determine the weighted median node along the path from u<sup>j</sup> to v<sup>j</sup>.

Return an array ans, where `ans[j]` is the node index of the weighted median for `queries[j]`.

### Testcases:
#### Example 1:

Input: n = 2, edges = `[[0,1,7]]`, queries = `[[1,0],[0,1]]`

Output: `[0,1]`

Explanation:  
<img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193447.png"><br>
| Query | Path | Edge Weights | Total Path Weight | Half | Explanation | Answer |
|-------|------|--------------|-------------------|------|-------------|--------|
| `[1, 0]` | `1 → 0` | `[7]` | 7 | 3.5 | Sum from 1 → 0 = 7 >= 3.5, median is node 0. | 0 |
| `[0, 1]` | `0 → 1` | `[7]` | 7 | 3.5 | Sum from 0 → 1 = 7 >= 3.5, median is node 1. | 1 |

#### Example 2:

Input: n = 3, edges = `[[0,1,2],[2,0,4]]`, queries = `[[0,1],[2,0],[1,2]]`

Output: `[1,0,2]`

Explanation:  
<img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193610.png"><br>
| Query | Path | Edge Weights | Total Path Weight | Half | Explanation | Answer |
|-------|------|--------------|-------------------|------|-------------|--------|
| `[0, 1]` | `0 → 1` | `[2]` | 2 | 1 | Sum from 0 → 1 = 2 >= 1, median is node 1. | 1 |
| `[2, 0]` | `2 → 0` | `[4]` | 4 | 2 | Sum from 2 → 0 = 4 >= 2, median is node 0. | 0 |
| `[1, 2]` | `1 → 0 → 2` | `[2, 4]` | 6 | 3 | Sum from 1 → 0 = 2 < 3. <br>Sum from 1 → 2 = 2 + 4 = 6 >= 3, median is node 2. | 2 |


### Example 3:

Input: n = 5, edges = `[[0,1,2],[0,2,5],[1,3,1],[2,4,3]]`, queries = `[[3,4],[1,2]]`

Output: `[2,2]`

Explanation:  
<img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193857.png"><br>
| Query | Path | Edge Weights | Total Path Weight | Half | Explanation | Answer |
|-------|------|--------------|-------------------|------|-------------|--------|
| `[3, 4]` | `3 → 1 → 0 → 2 → 4` | `[1, 2, 5, 3]` | 11 | 5.5 | Sum from 3 → 1 = 1 < 5.5. <br>Sum from 3 → 0 = 1 + 2 = 3 < 5.5. <br>Sum from 3 → 2 = 1 + 2 + 5 = 8 >= 5.5, median is node 2. | 2 |
| `[1, 2]` | `1 → 0 → 2` | `[2, 5]` | 7 | 3.5 | Sum from 1 → 0 = 2 < 3.5. <br>Sum from 1 → 2 = 2 + 5 = 7 >= 3.5, median is node 2. | 2 |


### Constraints:

- 2 <= `n` <= 10<sup>5</sup>
- `edges.length` == n - 1
- `edges[i]` == `[u`<sup>`i`</sup>`, v`<sup>`i`</sup>`, w`<sup>`i`</sup>`]`
- 0 <= u<sup>i</sup>, v<sup>i</sup> < n
- 1 <= w<sup>i</sup> <= 10<sup>9</sup>
- 1 <= `queries.length` <= 10<sup>5</sup>
- `queries[j]` == `[u`<sup>`j`</sup>`, v`<sup>`j`</sup>`]`
- 0 <= u<sup>j</sup>, v<sup>j</sup> < n
- The input is generated such that `edges` represents a valid tree.