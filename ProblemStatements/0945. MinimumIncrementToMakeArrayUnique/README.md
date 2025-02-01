# 945. Minimum Increment to Make Array Unique
## Question Level: Medium
### Description:
You are given an integer array `nums`. In one move, you can pick an index i where 0 <= i < `nums.length` and increment `nums[i]` by 1.

Return the minimum number of moves to make every value in nums unique.

The test cases are generated so that the answer fits in a 32-bit integer.

### Examples:
#### Example 1:

Input: nums = `[1,2,2]`<br>
Output: 1<br>
Explanation: After 1 move, the array could be `[1, 2, 3]`.<br>
#### Example 2:

Input: nums = `[3,2,1,2,1,7]`<br>
Output: 6<br>
Explanation: After 6 moves, the array could be `[3, 4, 1, 2, 5, 7]`.<br>
It can be shown that it is impossible for the array to have all unique values with 5 or less moves.

### Constraints:

- 1 <= `nums.length` <= 10^5
- 0 <= `nums[i]` <= 10^5

### <i>Concepts Used:
- Array
- Greedy
- Sorting
- Counting </i>