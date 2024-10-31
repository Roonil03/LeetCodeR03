# 1671. Minimum Number of Removals to Make Mountain Array
## Question Level: Hard
### Description:
You may recall that an array arr is a mountain array if and only if:

1. ``arr.length`` >= 3
2. There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
    - ``arr[0]`` < ``arr[1]`` < ... < ``arr[i - 1]`` < ``arr[i]``
    - ``arr[i]`` > ``arr[i + 1]`` > ... > ``arr[arr.length - 1]``

Given an integer array nums​​​, return the minimum number of elements to remove to make nums​​​ a mountain array.

### Examples:
<b>Example 1:</b><br>
Input: nums = [1,3,1]<br>
Output: 0<br>
Explanation: The array itself is a mountain array so we do not need to remove any elements.<br>

<b>Example 2:</b><br>
Input: nums = [2,1,1,5,6,2,3,1]<br>
Output: 3<br>
Explanation: One solution is to remove the elements at indices 0, 1, and 5, making the array nums = [1,5,6,3,1].<br>

### Constraints:
- 3 <= ``nums.length`` <= 1000
- 1 <= ``nums[i]`` <= 10^9
- It is guaranteed that you can make a mountain array out of ``nums``.

### <i>Concepts Used:
- Array
- Binary Search
- Dynamic Programming
- Greedy </i>
 