# 3409. Longest Subsequence With Decreasing Adjacent Difference
## Question Level: Medium
### Description:
You are given an array of integers nums.

Your task is to find the length of the longest subsequence seq of nums, such that the absolute differences between consecutive elements form a non-increasing sequence of integers. In other words, for a subsequence `seq0`, `seq1`, `seq2`, ..., `seqm` of nums, `|seq1 - seq0|` >= `|seq2 - seq1|` >= ... >= `|seqm - seqm - 1|`.

Return the length of such a subsequence.

A subsequence is an non-empty array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.


### Examples:
#### Example 1:
Input: nums = `[16,6,3]`<br>
Output: 3<br>
Explanation: <br>
The longest subsequence is `[16, 6, 3]` with the absolute adjacent differences `[10, 3]`.<br>

#### Example 2:
Input: nums = `[6,5,3,4,2,1]`<br>
Output: 4<br>
Explanation:<br>
The longest subsequence is `[6, 4, 2, 1]` with the absolute adjacent differences `[2, 2, 1]`.<br>

#### Example 3:
Input: nums = `[10,20,10,19,10,20]`<br>
Output: 5<br>
Explanation: <br>
The longest subsequence is `[10, 20, 10, 19, 10]` with the absolute adjacent differences `[10, 10, 9, 9]`.<br>

### Constraints:

- 2 <= `nums.length` <= 10^4
- 1 <= `nums[i]` <= 300

### <i>Biweekly Contest Question</i>