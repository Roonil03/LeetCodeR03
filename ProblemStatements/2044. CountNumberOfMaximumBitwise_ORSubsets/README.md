# 2044. Count Number of Maximum Bitwise-OR Subsets
## Question Level: Medium
### Description:
Given an integer array nums, find the maximum possible ``bitwise OR`` of a subset of nums and return the number of different non-empty subsets with the maximum bitwise OR.

An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b. Two subsets are considered different if the indices of the elements chosen are different.

The bitwise OR of an array a is equal to ``a[0] OR a[1]`` ``OR`` ... ``OR`` ``a[a.length - 1]`` (0-indexed).

### Examples:
<b>Example 1:</b><br>
Input: nums = [3,1]<br>
Output: 2<br>
Explanation: The maximum possible bitwise OR of a subset is 3. There are 2 subsets with a bitwise OR of 3:<br>
- [3]
- [3,1]

<b>Example 2:</b><br>
Input: nums = [2,2,2]<br>
Output: 7<br>
Explanation: All non-empty subsets of [2,2,2] have a bitwise OR of 2. There are 23 - 1 = 7 total subsets.<br>

<b>Example 3:</b><br>
Input: nums = [3,2,1,5]<br>
Output: 6<br>
Explanation: The maximum possible bitwise OR of a subset is 7. There are 6 subsets with a bitwise OR of 7:
- [3,5]
- [3,1,5]
- [3,2,5]
- [3,2,1,5]
- [2,5]
- [2,1,5]

### Constraints:
- 1 <= nums.length <= 16
- 1 <= nums[i] <= 10^5

### <i>Concepts used:
- Array
- Backtracking
- Bit Manipulation
- Enumeration</i>
