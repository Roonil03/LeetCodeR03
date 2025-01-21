# 1537. Get the Maximum Score
## Question Level: Hard
### Description:
You are given two sorted arrays of distinct integers `nums1` and `nums2`.

A valid path is defined as follows:
- Choose array `nums1` or `nums2` to traverse (from index-0).
- Traverse the current array from left to right.
- If you are reading any value that is present in `nums1` and `nums2` you are allowed to change your path to the other array. (Only one repeated value is considered in the valid path).
- The score is defined as the sum of unique values in a valid path.

Return the maximum score you can obtain of all possible valid paths. Since the answer may be too large, return it modulo 10^9 + 7.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/07/16/sample_1_1893.png"><br>
Input: nums1 = `[2,4,5,8,10]`, nums2 = `[4,6,8,9]`<br>
Output: 30<br>
Explanation: Valid paths:<br>
`[2,4,5,8,10]`, `[2,4,5,8,9]`, `[2,4,6,8,9]`, `[2,4,6,8,10]`,  (starting from nums1)<br>
`[4,6,8,9]`, `[4,5,8,10]`, `[4,5,8,9]`, `[4,6,8,10]`    (starting from nums2)<br>
The maximum is obtained with the path in green `[2,4,6,8,10]`.<br>
#### Example 2:

Input: nums1 = `[1,3,5,7,9]`, nums2 = `[3,5,100]`<br>
Output: 109<br>
Explanation: Maximum sum is obtained with the path `[1,3,5,100]`.<br>
#### Example 3:

Input: nums1 = `[1,2,3,4,5]`, nums2 = `[6,7,8,9,10]`<br>
Output: 40<br>
Explanation: There are no common elements between `nums1` and `nums2`.<br>
Maximum sum is obtained with the path `[6,7,8,9,10]`.<br>

### Constraints:

- 1 <= `nums1.length`, `nums2.length` <= 10^5
- 1 <= `nums1[i]`, `nums2[i]` <= 10^7
- `nums1` and `nums2` are strictly increasing.

### <i>Concepts Used:
- Array
- Two Pointers
- Dynamic Programming
- Greedy </i>