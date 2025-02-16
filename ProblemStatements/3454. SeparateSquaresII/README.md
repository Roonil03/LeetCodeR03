# 3454. Separate Squares II
## Question Level: Hard### Description:
You are given a 2D integer array squares. Each `squares[i]` = `[xi, yi, li]` represents the coordinates of the bottom-left point and the side length of a square parallel to the x-axis.

Find the minimum y-coordinate value of a horizontal line such that the total area covered by squares above the line equals the total area covered by squares below the line.

Answers within 10^-5 of the actual answer will be accepted.

Note: Squares may overlap. Overlapping areas should be counted only once in this version.

### Examples:
#### Example 1:

Input: squares = `[[0,0,1],[2,2,1]]`<br>
Output: 1.00000<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2025/01/15/4065example1drawio.png"><br>
Any horizontal line between y = 1 and y = 2 results in an equal split, with 1 square unit above and 1 square unit below. The minimum y-value is 1.

#### Example 2:

Input: squares = `[[0,0,2],[1,1,1]]`<br>
Output: 1.00000<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2025/01/15/4065example2drawio.png"><br>
Since the blue square overlaps with the red square, it will not be counted again. Thus, the line y = 1 splits the squares into two equal parts.

### Constraints:

- 1 <= `squares.length` <= 5 * 10^4
- `squares[i]` = `[xi, yi, li]`
- `squares[i].length` == 3
- 0 <= `xi`, `yi` <= 10^9
- 1 <= `li` <= 10^9

### <i>Biweekly Contest Question</i>