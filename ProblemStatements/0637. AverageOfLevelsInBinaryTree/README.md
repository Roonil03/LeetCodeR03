# 637. Average of Levels in Binary Tree
## Question Level: Easy
### Description:
Given the `root` of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10^-5 of the actual answer will be accepted.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2021/03/09/avg1-tree.jpg"><br>
Input: root = `[3,9,20,null,null,15,7]`<br>
Output: `[3.00000,14.50000,11.00000]`<br>
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.<br>
Hence return `[3, 14.5, 11]`.<br>
#### Example 2:
<img src-"https://assets.leetcode.com/uploads/2021/03/09/avg2-tree.jpg"><br>
Input: root = `[3,9,20,15,7]`<br>
Output: `[3.00000,14.50000,11.00000]`<br>

### Constraints:

- The number of nodes in the tree is in the range `[1, 10^4]`.
- -2^31 <= Node.val <= 2^31 - 1

### <i>Concepts Used:
- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree </i>