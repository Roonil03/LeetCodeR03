# 98. Validate Binary Search Tree
## Question Level: Medium
### Description:
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg"><br>
Input: root = `[2,1,3]`<br>
Output: true<br>
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg"><br>
Input: root = `[5,1,4,null,null,3,6]`<br>
Output: false<br>
Explanation: The root node's value is 5 but its right child's value is 4.<br>


### Constraints:

- The number of nodes in the tree is in the range `[1, 10^4]`.
- -2^31 <= `Node.val` <= 2^31 - 1


### <i>Concepts Used:
- Tree
- Depth-First Search
- Binary Search Tree
- Binary Tree </i>