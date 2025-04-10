# 894. All Possible Full Binary Trees
## Question Level: Medium
### Description:
Given an integer n, return a list of all possible full binary trees with n nodes. Each node of each tree in the answer must have `Node.val` == 0.

Each element of the answer is the root node of one possible tree. You may return the final list of trees in any order.

A full binary tree is a binary tree where each node has exactly 0 or 2 children.
### Examples:
#### Example 1:

<img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/22/fivetrees.png"><br>
Input: n = 7  
Output: `[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]`  
#### Example 2:

Input: n = 3  
Output: `[[0,0,0]]`  
### Constraints:

- 1 <= n <= 20

### <i>Concepts Used:
- Dynamic Programming
- Tree
- Recursion
- Memoization
- Binary Tree </i>