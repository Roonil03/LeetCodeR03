# 1372. Longest ZigZag Path in a Binary Tree
## Question Level: Medium
### Description:
You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:
- Choose any node in the binary tree and a direction (right or left).
- If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
- Change the direction from right to left or from left to right.
- Repeat the second and third steps until you can't move in the tree.

Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/01/22/sample_1_1702.png"><br>
Input: root = `[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]`
Output: 3
Explanation: Longest ZigZag path in blue nodes `(right -> left -> right)`.
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2020/01/22/sample_2_1702.png"><br>
Input: root = `[1,1,1,null,1,null,null,1,1,null,1]`
Output: 4
Explanation: Longest ZigZag path in blue nodes `(left -> right -> left -> right)`.
#### Example 3:

Input: root = `[1]`
Output: 0

### Constraints:

- The number of nodes in the tree is in the range `[1, 5 * 10`<sup>`4`</sup>`]`.
- 1 <= `Node.val` <= 100

### <i>Concepts Used:
- Dynamic Programming
- Tree
- Depth-First Search
- Binary Tree</i>