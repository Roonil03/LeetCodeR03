# 3202. Find the Maximum Length of Valid Subsequence II
## Question Level: Medium
### Description:
You are given an integer array nums and a positive integer k.  
A subsequence sub of nums with length x is called valid if it satisfies:
- `(sub[0] + sub[1]) % k == (sub[1] + sub[2]) % k` == ... == `(sub[x - 2] + sub[x - 1]) % k`.  

Return the length of the longest valid subsequence of nums.

### Examples:
#### Example 1:

Input: nums = `[1,2,3,4,5]`, k = 2

Output: 5

Explanation:

The longest valid subsequence is `[1, 2, 3, 4, 5]`.

#### Example 2:

Input: nums = `[1,4,2,3,1,4]`, k = 3

Output: 4

Explanation:

The longest valid subsequence is `[1, 4, 1, 4]`.

### Constraints:

- 2 <= `nums.length` <= 10<sup>3</sup>
- 1 <= `nums[i]` <= 10<sup>7</sup>
- 1 <= `k` <= 10<sup>3</sup>

### <i>Concepts Used:
- Array
- Dynamic Programming</i>