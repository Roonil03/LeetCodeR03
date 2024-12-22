# 3397. Maximum Number of Distinct Elements After Operations
## Question Level: Medium
### Description:
You are given an integer array `nums` and an integer k.

You are allowed to perform the following operation on each element of the array at most once:
- Add an integer in the range `[-k, k]` to the element.

Return the maximum possible number of distinct elements in nums after performing the operations.

### Examples:
#### Example 1:
Input: nums = `[1,2,2,3,3,4]`, k = 2<br>
Output: 6<br>
Explanation:<br>
nums changes to `[-1, 0, 1, 2, 3, 4]` after performing operations on the first four elements.

#### Example 2:
Input: nums = `[4,4,4,4]`, k = 1<br>
Output: 3<br>
Explanation:<br>
By adding -1 to `nums[0]` and 1 to `nums[1]`, nums changes to `[3, 5, 4, 4]`.


### Constraints:

- 1 <= `nums.length` <= 10^5
- 1 <= `nums[i]` <= 10^9
- 0 <= `k` <= 10^9

### <i>Weekly Contest Question</i>