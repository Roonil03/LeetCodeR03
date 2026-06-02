# Weekly Contest 503
## Ranking:
## Question 1:
*Limit Occurrences in Sorted Array*
### Difficulty: Easy
### Points: 3
### Description:
You are given a sorted integer array `nums` and an integer k.

Return an array such that each distinct element appears at most k times, while preserving the relative order of the elements in nums.

### Examples:
#### Example 1:

Input: nums = `[1,1,1,2,2,3]`, k = 2

Output: `[1,1,2,2,3]`

Explanation:

Each element can appear at most 2 times.
- The element 1 appears 3 times, so only 2 occurrences are kept.
- The element 2 appears 2 times, so both occurrences are kept.
- The element 3 appears 1 time, so it is kept.  
Thus, the resulting array is `[1, 1, 2, 2, 3]`.

#### Example 2:

Input: nums = `[1,2,3]`, k = 1

Output: `[1,2,3]`

Explanation:

All elements are distinct and already appear at most once, so the array remains unchanged.

### Constraints:

- 1 <= `nums.length` <= 100
- 1 <= `nums[i]` <= 100
- `nums` is sorted in non-decreasing order.
- 1 <= `k` <= `nums.length`

### <i>Follow-up:

- Can you solve this in-place using O(1) extra space?
- Note that the space used for returning or resizing the result does not count toward the space complexity mentioned above, as some languages do not support in-place resizing.</i>

## Question 2:
*Password Strength*
### Difficulty: Medium
### Points: 5
### Description:
You are given a string `password`.

The strength of the password is calculated based on the following rules:
- 1 point for each distinct lowercase letter ('`a`' to '`z`').
- 2 points for each distinct uppercase letter ('`A`' to '`Z`').
- 3 points for each distinct digit ('`0`' to '`9`').
- 5 points for each distinct special character from the set "`!@#$`".

Each character contributes at most once, even if it appears multiple times.

Return an integer denoting the strength of the password.

### Examples:
#### Example 1:

Input: password = "`aA1!`"

Output: 11

Explanation:

- The distinct characters are '`a`', '`A`', '`1`' and '`!`'.
- Thus, the strength = 1 + 2 + 3 + 5 = 11.

#### Example 2:

Input: password = "`bbB11#`"

Output: 11

Explanation:
- The distinct characters are '`b`', '`B`', '`1`' and '`#`'.
- Thus, the strength = 1 + 2 + 3 + 5 = 11.​​​​​​​

### Constraints:

- 1 <= `password.length` <= 10<sup>5</sup>
- `password` consists of lowercase and uppercase English letters, digits, and special characters from "`!@#$`".

## Question 3:
*Minimum Operations to Sort a Permutation*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer array nums of length n, where nums is a permutation of the numbers in the range `[0..n - 1]`.

You may perform only the following operations:
- Reverse the entire array.
- Rotate Left by One: Move the first element to the end of the array, and rest elements to left by one position.

Return the minimum number of operations required to sort the array in increasing order. If it is not possible to sort the array using only the given operations, return -1.

A permutation is a rearrangement of all the elements of an array.

### Examples:
#### Example 1:

Input: nums = `[0,2,1]`

Output: 2

Explanation:
- Rotate Left by one: `[2, 1, 0]`
- Reverse the array: `[0, 1, 2]`  
The array becomes sorted in 2 operations, which is minimal

#### Example 2:

Input: nums = `[1,0,2]`

Output: 2

Explanation:
- Reverse the array: `[2, 0, 1]`
- Rotate Left by one: `[0, 1, 2]`  
The array becomes sorted in 2 operations, which is minimal.

#### Example 3:

Input: nums = `[2,0,1,3]`

Output: -1

Explanation:

It is impossible to reach `[2, 0, 1, 3]`. Thus, the answer is -1.

### Constraints:
- 1 <= `n` == `nums.length` <= 10<sup>5</sup>
- 0 <= `nums[i]` <= `n - 1`
- `nums` is a permutation of integers from 0 to n - 1.

## Question 4:
*Number of Pairs After Increment*
### Difficulty: Hard
### Points: 6
### Description:
You are given two integer arrays `nums1` and `nums2`, and a 2D integer array `queries`.

Each `queries[i]` is one of the following types:
- `[1, x, y, val] `– Add val to every element in `nums2[x..y]`.
- `[2, tot]` – Compute the number of pairs `(j, k)` such that `nums1[j] + nums2[k] == tot`.

Return an integer array answer, where `answer[j]` is the number of pairs for the jth query of type 2.

### Examples:
#### Example 1:

Input: nums1 = `[1,2]`, nums2 = `[3,4]`, queries = `[[2,5],[1,0,0,2],[2,5]]`

Output: `[2,1]`

Explanation:
- `queries[0]` = `[2, 5]`: Valid pairs are `nums1[0] + nums2[1]` = `1 + 4` = 5 and `nums1[1] + nums2[0]` = `2 + 3` = 5.
- `queries[1]` = `[1, 0, 0, 2]`: Add 2 to `nums2[0]`, resulting in `nums2` = `[5, 4]`.
- `queries[2]` = `[2, 5]`: Valid pair is `nums1[0] + nums2[1]` = `1 + 4` = 5.
- Thus, the answer = `[2, 1]`.

#### Example 2:

Input: nums1 = `[1,1]`, nums2 = `[2,2,3]`, queries = `[[2,4],[1,0,1,1],[2,4]]`

Output: [2,6]

Explanation:
- `queries[0]` = `[2, 4]`: Valid pairs are `nums1[0] + nums2[2]` = `1 + 3` and `nums1[1] + nums2[2]` = `1 + 3`.
- `queries[1]` = `[1, 0, 1, 1]`: Add 1 to `nums2[0]` and `nums2[1]`, resulting in `nums2` = `[3, 3, 3]`.
- `queries[2]` = `[2, 4]`: Every element of `nums1` = `[1, 1]` pairs with every element of `nums2` = `[3, 3, 3]` as 1 + 3 = 4. That gives 2 × 3 = 6 pairs in total.
- Thus, the answer = `[2, 6]`.

#### Example 3:

Input: nums1 = `[2,5,8,4]`, nums2 = `[1,3,8]`, queries = `[[2,9],[1,1,2,1],[2,10]]`

Output: `[1,0]`

Explanation:
- `queries[0]` = `[2, 9]`: Only valid pair is `nums1[2] + nums2[0]` = `8 + 1` = 9.
- `queries[1]` = `[1, 1, 2, 1]`: Add 1 to `nums2[1]` and `nums2[2]`, resulting in​​​​​​​ `nums2` = `[1, 4, 9]`.
- `queries[2]` = `[2, 10]`: No pair sums to 10.
- Thus, the answer = `[1, 0]`.

### Constraints:

- 1 <= `nums1.length` <= 5
- 1 <= `nums2.length` <= 5 * 10<sup>4</sup>
- 1 <= `nums1[i]`, `nums2[i]` <= 10<sup>5</sup>
- 1 <= `queries.length` <= 5 * 10<sup>4</sup>
- `queries[i].length` == 2 or 4
    - `queries[i]` == `[1, x, y, val]`, or
    - `queries[i]` == `[2, tot]`
    - 0 <= `x` <= `y` < `nums2.length`
    - 1 <= `val` <= 10<sup>5</sup>
    - 1 <= `tot` <= 10<sup>9​​​​​​​</sup>
