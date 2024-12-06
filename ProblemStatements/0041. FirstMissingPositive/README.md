# 41. First Missing Positive
## Question Level: Hard
### Description:
Given an unsorted integer array `nums`. Return the smallest positive integer that is not present in `nums`.

You must implement an algorithm that runs in <i>O(n)</i> time and uses <i>O(1)</i> auxiliary space.

### Examples:
#### Example 1:

Input: nums = `[1,2,0]`<br>
Output: 3<br>
Explanation: The numbers in the range `[1,2]` are all in the array.
#### Example 2:

Input: nums = `[3,4,-1,1]`<br>
Output: 2<br>
Explanation: 1 is in the array but 2 is missing.
#### Example 3:

Input: nums = `[7,8,9,11,12]`<br>
Output: 1<br>
Explanation: The smallest positive integer 1 is missing.

### Constraints:

- 1 <= `nums.length` <= 10^5
- `-2^31` <= `nums[i]` <= `2^31 - 1`

### <i>Concepts Used:
- Array
- Hash Table </i>