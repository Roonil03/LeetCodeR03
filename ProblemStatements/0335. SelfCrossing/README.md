# 335. Self Crossing
## Question Level: Hard
### Description:
You are given an array of integers distance.

You start at the point `(0, 0)` on an X-Y plane, and you move `distance[0]` meters to the north, then `distance[1]` meters to the west, `distance[2]` meters to the south, `distance[3]` meters to the east, and so on. In other words, after each move, your direction changes counter-clockwise.

Return true if your path crosses itself or false if it does not.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2022/12/21/11.jpg"><br>
Input: distance = `[2,1,1,2]`  
Output: true  
Explanation: The path crosses itself at the point `(0, 1)`.  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2022/12/21/22.jpg"><br>
Input: distance = `[1,2,3,4]`  
Output: false  
Explanation: The path does not cross itself at any point.  
#### Example 3:

<img src="https://assets.leetcode.com/uploads/2022/12/21/33.jpg"><br> 
Input: distance = `[1,1,1,2,1]`  
Output: true  
Explanation: The path crosses itself at the point `(0, 0)`.  

### Constraints:

- 1 <= `distance.length` <= 10<sup>5</sup>
- 1 <= `distance[i]` <= 10<sup>5</sup>

### <i>Concepts Used:
- Array
- Math
- Geometry</i>