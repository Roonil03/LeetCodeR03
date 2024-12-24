# 963. Minimum Area Rectangle II
## Question Level: Medium
### Description:
You are given an array of points in the X-Y plane points where `points[i]` = `[xi, yi]`.

Return the minimum area of any rectangle formed from these points, with sides not necessarily parallel to the X and Y axes. If there is not any such rectangle, return 0.

Answers within 10^-5 of the actual answer will be accepted.

### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2018/12/21/1a.png"><br>
Input: points = `[[1,2],[2,1],[1,0],[0,1]]`<br>
Output: 2.00000<br>
Explanation: The minimum area rectangle occurs at `[1,2]`,`[2,1]`,`[1,0]`,`[0,1]`, with an area of 2.
#### Example 2:
<img src="https://assets.leetcode.com/uploads/2018/12/22/2.png"><br>
Input: points = `[[0,1],[2,1],[1,1],[1,0],[2,0]]`<br>
Output: 1.00000<br>
Explanation: The minimum area rectangle occurs at `[1,0]`,`[1,1]`,`[2,1]`,`[2,0]`, with an area of 1.
#### Example 3:
<img src="https://assets.leetcode.com/uploads/2018/12/22/3.png"><br>
Input: points = `[[0,3],[1,2],[3,1],[1,3],[2,1]]`<br>
Output: 0<br>
Explanation: There is no possible rectangle to form from these points.


### Constraints:

- 1 <= `points.length` <= 50
- `points[i].length` == 2
- 0 <= `xi`, `yi` <= 4 * 10^4
- All the given points are unique.


### <i>Concepts Used:
- Array
- Math
- Geometry </i>
 