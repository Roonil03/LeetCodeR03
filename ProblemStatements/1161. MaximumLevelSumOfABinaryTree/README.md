# 1161. Maximum Level Sum of a Binary Tree
## Question Level: Medium
### Description:
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG"><br>
Input: root = `[1,7,0,7,-8,null,null]`  
Output: 2  
Explanation:   
Level 1 sum = 1.  
Level 2 sum = 7 + 0 = 7.  
Level 3 sum = 7 + -8 = -1.  
So we return the level with the maximum sum which is level 2.  
#### Example 2:

Input: root = `[989,null,10250,98693,-89388,null,null,null,-32127]`  
Output: 2  

### Constraints:

- The number of nodes in the tree is in the range `[1, 10`<sup>`4`</sup>`]`.
- -10<sup>5</sup> <= `Node.val` <= 10<sup>5</sup>

### <i>Concepts Used:
- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree</i>