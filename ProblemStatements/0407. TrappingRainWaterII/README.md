# 407. Trapping Rain Water II
## Question Level: Hard
### Description:
Given an `m x n` integer matrix heightMap representing the height of each unit cell in a 2D elevation map, return the volume of water it can trap after raining.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2021/04/08/trap1-3d.jpg"><br>
Input: heightMap = `[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]`<br>
Output: 4<br>
Explanation: After the rain, water is trapped between the blocks.<br>
We have two small ponds 1 and 3 units trapped.<br>
The total volume of water trapped is 4.
#### Example 2:
<img src="https://assets.leetcode.com/uploads/2021/04/08/trap2-3d.jpg"><br>
Input: heightMap = `[[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]`<br>
Output: 10<br>