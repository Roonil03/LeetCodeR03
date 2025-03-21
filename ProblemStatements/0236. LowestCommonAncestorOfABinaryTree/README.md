# 236. Lowest Common Ancestor of a Binary Tree
## Question Level: Medium
### Description:
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “*The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).*”

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"><br>
Input: root = `[3,5,1,6,2,0,8,null,null,7,4]`, p = 5, q = 1<br>
Output: 3<br>
Explanation: The LCA of nodes 5 and 1 is 3.<br>
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"><br>
Input: root = `[3,5,1,6,2,0,8,null,null,7,4]`, p = 5, q = 4<br>
Output: 5<br>
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.<br>
#### Example 3:

Input: root = `[1,2]`, p = 1, q = 2<br>
Output: 1<br>

### Constraints:

- The number of nodes in the tree is in the range [2, 10<sup>5</sup>].
- -10<sup>9</sup> <= `Node.val` <= 10<sup>9</sup>
- All `Node.val` are unique.
- p != q
- p and q will exist in the tree.

### <i>Concepts Used:
- Tree
- Depth-First Search
- Binary Tree </i>