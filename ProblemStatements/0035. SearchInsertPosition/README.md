# 35. Search Insert Position
## Question Level: Easy

### Description:
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with ``O(log n)`` runtime complexity.

### Examples:
#### Example 1:

Input: nums = `[1,3,5,6]`, target = 5<br>
Output: 2<br>
#### Example 2:

Input: nums = `[1,3,5,6]`, target = 2<br>
Output: 1<br>
#### Example 3:

Input: nums = `[1,3,5,6]`, target = 7<br>
Output: 4<br>

### Constraints:

- 1 <= ``nums.length`` <= 104
- -10^4 <= `nums[i]` <= 10^4
- `nums` contains distinct values sorted in ascending order.
- -10^4 <= `target` <= 10^4

### <i>Concepts Used:
- Array
- Binary Search </i>