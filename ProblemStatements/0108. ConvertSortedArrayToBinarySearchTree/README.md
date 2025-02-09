# 108. Convert Sorted Array to Binary Search Tree
## Question Level: Easy
### Description:
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg"><br>
Input: nums = `[-10,-3,0,5,9]`<br>
Output: `[0,-3,9,-10,null,5]`<br>
Explanation: `[0,-10,5,null,-3,null,9]` is also accepted:<br>
<img src="https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg">
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/02/18/btree.jpg"><br>
Input: nums = `[1,3]`<br>
Output: `[3,1]`<br>
Explanation: `[1,null,3]` and `[3,1]` are both height-balanced BSTs.<br>

### Constraints:

- 1 <= `nums.length` <= 10^4
- -10^4 <= `nums[i]` <= 10^4
- `nums` is sorted in a strictly increasing order.

### <i>Concepts Used:
- Array
- Divide and Conquer
- Tree
- Binary Search Tree
- Binary Tree</i>