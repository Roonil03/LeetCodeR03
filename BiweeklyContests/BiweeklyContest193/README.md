# Biweekly Contest 193
## Ranking:
## Question 1:
*Maximum Product Pair With Target Sum*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array nums and an integer target.

A pair of distinct indices `(i, j)` is valid if:
- `nums[i] + nums[j]` == `target`
- `nums[i]` > `nums[j]`

Return a valid pair `[i, j]` whose product `nums[i] * nums[j]` is maximum among all valid pairs. If no valid pair exists, return `[-1, -1]`.

If multiple valid pairs achieve the maximum product, you may return any of them.

### Examples:
#### Example 1:

Input: nums = `[1,2,3,4]`, target = 5

Output: [2,1]

Explanation:

There are 2 valid pairs:  
| No. | $(i, j)$ | $\text{nums}[i]$ | $\text{nums}[j]$ | Sum | Product |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | $(3, 0)$ | $4$ | $1$ | $5$ | $4$ |
| 2 | $(2, 1)$ | $3$ | $2$ | $5$ | $6$ |

Both pairs sum to 5 and satisfy `nums[i] > nums[j]`. The second pair has the larger product, 3 * 2 = 6, so the answer is `[2, 1]`.

#### Example 2:

Input: nums = `[-3,-1,4,2]`, target = 1

Output: `[3,1]`

Explanation:

There are 2 valid pairs:  
| No. | $(i, j)$ | $\text{nums}[i]$ | $\text{nums}[j]$ | Sum | Product |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | $(2, 0)$ | $4$ | $-3$ | $1$ | $-12$ |
| 2 | $(3, 1)$ | $2$ | $-1$ | $1$ | $-2$ |

Since -2 > -12, the pair at indices (3, 1) is chosen and the answer is `[3, 1]`.

#### Example 3:

Input: nums = `[3,3,5]`, target = 6

Output: `[-1,-1]`

Explanation:

​​​​​​​No valid pair exists, since `nums[0]` and `nums[1]` sum to 6 but are equal.
### Constraints:

- 2 <= `nums.length` <= 100
- -100 <= `nums[i]`, `target` <= 100​​​​​​​

## Question 2:
*Longest Resilient Subarray I*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums and an integer k.

A subarray is resilient if, for every position in it, deleting the element at that position leaves the remaining elements with a sum divisible by k.

For a subarray of length 1, the remaining sum is 0, which is divisible by k.

Return the length of the longest resilient subarray.

### Examples:
#### Example 1:

Input: nums = `[2,4,6,3]`, k = 2

Output: 3

Explanation:

The subarray `[2, 4, 6]` sums to 12. Deleting 2, 4, or 6 leaves a sum of 10, 8, or 6, respectively, each of which is divisible by 2. Therefore, the answer is 3.

#### Example 2:

Input: nums = `[1,4,7,1]`, k = 3

Output: 4

Explanation:

The subarray `[1, 4, 7, 1]` sums to 13. Deleting the first 1, the 4, the 7, or the last 1 leaves a sum of 12, 9, 6, or 12, respectively, each of which is divisible by 3. Therefore, the answer is 4.

#### Example 3:

Input: nums = `[5,5,4,8]`, k = 4

Output: 2

Explanation:

The subarray `[5, 5]` is not resilient, since deleting either 5 leaves a sum of 5, which is not divisible by 4.

The subarray `[4, 8]` sums to 12. Deleting 4 or 8 leaves a sum of 8 or 4, respectively, each of which is divisible by 4. Therefore, the answer is 2.

### Constraints:

- 1 <= `nums.length` <= 2500
- 1 <= `nums[i]` <= 10<sup>5</sup>
- 1 <= `k` <= 10<sup>5</sup>

## Question 3:
*Longest Resilient Subarray II*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer array nums and an integer k.

A subarray is resilient if, for every position in it, deleting the element at that position leaves the remaining elements with a sum divisible by k.

For a subarray of length 1, the remaining sum is 0, which is divisible by k.

Return the length of the longest resilient subarray.

### Examples:
#### Example 1:

Input: nums = `[2,4,6,3]`, k = 2

Output: 3

Explanation:

The subarray `[2, 4, 6]` sums to 12. Deleting 2, 4, or 6 leaves a sum of 10, 8, or 6, respectively, each of which is divisible by 2. Therefore, the answer is 3.

#### Example 2:

Input: nums = `[1,4,7,1]`, k = 3

Output: 4

Explanation:

The subarray `[1, 4, 7, 1]` sums to 13. Deleting the first 1, the 4, the 7, or the last 1 leaves a sum of 12, 9, 6, or 12, respectively, each of which is divisible by 3. Therefore, the answer is 4.

#### Example 3:

Input: nums = `[5,5,4,8]`, k = 4

Output: 2

Explanation:

The subarray `[5, 5]` is not resilient, since deleting either 5 leaves a sum of 5, which is not divisible by 4.

The subarray `[4, 8]` sums to 12. Deleting 4 or 8 leaves a sum of 8 or 4, respectively, each of which is divisible by 4. Therefore, the answer is 2.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>5</sup>
- 1 <= `k` <= 10<sup>5</sup>

## Question 4:
*Minimize Inversions in a Prefix Tree*
### Difficulty: Hard
### Points: 7
### Description:
You are given an array words of n unique strings containing only lowercase English letters from '`a`' to '`j`'. You are also given an integer array target containing each index from 0 to n - 1 exactly once.

Build a prefix tree from words:
- The root represents the empty string "".
- Each other node represents a different non-empty prefix of a word. The parent of a node is the prefix obtained by removing its last character.
- The node representing the complete word `words[i]` is marked with index i. A marked node may also have children.

You may choose the order of the children of each node independently. You cannot change which node is a child's parent.

Traverse the tree starting at the root. When visiting a node, first record its index if it is marked. Then visit its children in the chosen order, completing the traversal of each child before moving to the next. This is a pre-order traversal.

The recorded indices form an array traversal of length n. Let `pos[i]` be the position of index i in target.

An inversion is a pair of positions `(a, b)` such that `0 <= a < b < n` and `pos[traversal[a]] > pos[traversal[b]]`.

Return the minimum possible number of inversions over all choices of child orderings.


### Examples:
#### Example 1:

​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/04/screenshot-2026-03-04-at-55332pm.png"><br>

Input: words =` ["bad","bag","fig","fed"]`, target = `[3,1,0,2]`

Output: 2

Explanation:

Visit the root's "f" child before its "b" child. Within the "f" branch, visit "fed" before "fig". Within the "b" branch, visit "bag" before "bad".

This produces traversal = `[3,2,1,0]`:  
| Traversal position | Word index | Word | Position in target |
| :---: | :---: | :---: | :---: |
| 0 | 3 | "fed" | 0 |
| 1 | 2 | "fig" | 3 |
| 2 | 1 | "bag" | 1 |
| 3 | 0 | "bad" | 2 |

The positions in target are `[0,3,1,2]`. The inversions are the pairs of traversal positions `(1,2)` and `(1,3)`, since `3 > 1` and `3 > 2`.

No child ordering produces fewer than 2 inversions.

#### Example 2:

Input: words = `["ab","ac","ba","bc"]`, target = `[0,3,1,2]`

Output: 1

Explanation:

Visit the root's "`a`" child before its "`b`" child. Within these branches, visit "`ab`" before "`ac`", and "`bc`" before "`ba`".

This produces traversal = `[0,1,3,2]`:  
| Traversal position | Word index | Word | Position in target |
| :---: | :---: | :---: | :---: |
| 0 | 0 | "ab" | 0 |
| 1 | 1 | "ac" | 2 |
| 2 | 3 | "bc" | 1 |
| 3 | 2 | "ba" | 3 |

The positions in target are `[0,2,1,3]`. The only inversion is the pair of traversal positions `(1,2)`, since `2 > 1`.

No child ordering produces zero inversions.

### Constraints:

- 1 <= `n` == `words.length` <= 10<sup>4</sup>
- 1 <= `words[i].length` <= 20
- `words[i]` contains only lowercase English letters from '`a`' to '`j`'​​​​​​​
- All `words[i]` are unique
-0 <= `target[i]` <= `n - 1`
- `target` is a permutation of integers from 0 to `n - 1`
