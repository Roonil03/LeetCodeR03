# 951. Flip Equivalent Binary Trees
## Question Level: Medium
### Description:
For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.

A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.

Given the roots of two binary trees root1 and root2, return true if the two trees are flip equivalent or false otherwise.


### Examples:
<b>Example 1:</b><br>
<img src="https://assets.leetcode.com/uploads/2018/11/29/tree_ex.png"><br>
Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]<br>
Output: true<br>
Explanation: We flipped at nodes with values 1, 3, and 5.<br>

<b>Example 2:</b><br>
Input: root1 = [], root2 = []<br>
Output: true<br>

<b>Example 3:</b><br>
Input: root1 = [], root2 = [1]<br>
Output: false<br>

### Constraints:
- The number of nodes in each tree is in the range [0, 100].
- Each tree will have unique node values in the range [0, 99].

### <i>Concepts Used:
- Tree
- Depth-First Search
- Binary Tree </i>
