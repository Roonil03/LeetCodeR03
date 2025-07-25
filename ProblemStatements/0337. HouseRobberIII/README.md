# 337. House Robber III
## Question Level: Medium
### Description:
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.

Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night.

Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/03/10/rob1-tree.jpg"><br>
Input: root = `[3,2,3,null,3,null,1]`  
Output: 7  
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/03/10/rob2-tree.jpg"><br>
Input: root = `[3,4,5,1,3,null,1]`  
Output: 9  
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.  

### Constraints:

- The number of nodes in the tree is in the range `[1, 10`<sup>`4`</sup>`]`.
- 0 <= `Node.val` <= 10<sup>4</sup>

### <i>Concepts Used:
- Dynamic Programming
- Tree
- Depth-First Search
- Binary Tree</i>