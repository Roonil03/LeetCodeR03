# 3432. Count Partitions with Even Sum Difference
## Question Level: Easy
### Description:
You are given an integer array nums of length n.

A partition is defined as an index i where 0 <= i < n - 1, splitting the array into two non-empty subarrays such that:
- Left subarray contains indices `[0, i]`.
- Right subarray contains indices `[i + 1, n - 1]`.

Return the number of partitions where the difference between the sum of the left and right subarrays is even.

### Examples:
#### Example 1:
Input: nums = `[10,10,3,7,6]`<br>
Output: 4<br>
Explanation:<br>
The 4 partitions are:
- `[10]`, `[10, 3, 7, 6]` with a sum difference of 10 - 26 = -16, which is even.
- `[10, 10]`, `[3, 7, 6]` with a sum difference of 20 - 16 = 4, which is even.
- `[10, 10, 3]`, `[7, 6]` with a sum difference of 23 - 13 = 10, which is even.
- `[10, 10, 3, 7]`, `[6]` with a sum difference of 30 - 6 = 24, which is even.
#### Example 2:
Input: nums = `[1,2,2]`<br>
Output: 0<br>
Explanation:<br>
No partition results in an even sum difference.<br>

#### Example 3:
Input: nums = `[2,4,6,8]`<br>
Output: 3<br>
Explanation:<br>
All partitions result in an even sum difference.<br>


### Constraints:

2 <= n == `nums.length` <= 100
1 <= `nums[i]` <= 100

### <i>Weekly Contest Question</i>