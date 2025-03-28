# 2583. Kth Largest Sum in a Binary Tree
## Question Level: Medium
### Description:
You are given the root of a binary tree and a positive integer k.

The level sum in the tree is the sum of the values of the nodes that are on the same level.

Return the kth largest level sum in the tree (not necessarily distinct). If there are fewer than k levels in the tree, return -1.

Note that two nodes are on the same level if they have the same distance from the root.

### Examples:
<b>Example 1:</b><br>
<img src="https://assets.leetcode.com/uploads/2022/12/14/binaryytreeedrawio-2.png"><br>
Input: root = [5,8,9,2,1,3,7,4,6], k = 2<br>
Output: 13<br>
Explanation: The level sums are the following:<br>
- Level 1: 5.
- Level 2: 8 + 9 = 17.
- Level 3: 2 + 1 + 3 + 7 = 13.
- Level 4: 4 + 6 = 10.
The 2nd largest level sum is 13.<br>

<b>Example 2:</b><br>
<img src="https://assets.leetcode.com/uploads/2022/12/14/treedrawio-3.png"><br>
Input: root = [1,2,null,3], k = 1<br>
Output: 3<br>
Explanation: The largest level sum is 3.<br>