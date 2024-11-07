# 2275. Largest Combination With Bitwise AND Greater Than Zero
## Question Level: Medium
### Description:
The bitwise ``AND`` of an array nums is the bitwise ``AND`` of all integers in nums.

- For example, for nums = [1, 5, 3], the bitwise AND is equal to 1 & 5 & 3 = 1.
- Also, for nums = [7], the bitwise AND is 7.

You are given an array of positive integers candidates. Evaluate the bitwise ``AND`` of every combination of numbers of candidates. Each number in candidates may only be used once in each combination.

Return the size of the largest combination of candidates with a bitwise AND greater than 0.

### Examples:
#### Example 1:

Input: candidates = ``[16,17,71,62,12,24,14]``<br>
Output: 4<br>
Explanation: The combination ``[16,17,62,24]`` has a bitwise AND of 16 & 17 & 62 & 24 = 16 > 0.<br>
The size of the combination is 4.<br>
It can be shown that no combination with a size greater than 4 has a bitwise AND greater than 0.<br>
Note that more than one combination may have the largest size.<br>
For example, the combination ``[62,12,24,14]`` has a bitwise AND of 62 & 12 & 24 & 14 = 8 > 0.<br>
#### Example 2:

Input: candidates = ``[8,8]``<br>
Output: 2<br>
Explanation: The largest combination ``[8,8]`` has a bitwise AND of 8 & 8 = 8 > 0.<br>
The size of the combination is 2, so we return 2.<br>

### Constraints:

- 1 <= ``candidates.length`` <= 10^5
- 1 <= ``candidates[i]`` <= 10^7

### <i>Concepts Used:
- Array
- Hash Table
- Bit Manipulation
- Counting