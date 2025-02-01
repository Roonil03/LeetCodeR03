# 3127. Make a Square with the Same Color
## Question Level: Easy
### Description:
You are given a 2D matrix grid of size `3 x 3` consisting only of characters '`B`' and '`W`'. Character '`W`' represents the white color, and character '`B`' represents the black color.

Your task is to change the color of at most one cell so that the matrix has a `2 x 2` square where all cells are of the same color.

Return true if it is possible to create a `2 x 2` square of the same color, otherwise, return false.

### Examples:
#### Example 1:
 
Input: grid = `[["B","W","B"],["B","W","W"],["B","W","B"]]`<br>
Output: true<br>
Explanation:<br>
It can be done by changing the color of the `grid[0][2]`.<br>

#### Example 2: 
 
Input: grid = `[["B","W","B"],["W","B","W"],["B","W","B"]]`<br>
Output: false<br>
Explanation:<br>
It cannot be done by changing at most one cell.<br>

#### Example 3: 
 
Input: grid = `[["B","W","B"],["B","W","W"],["B","W","W"]]`<br>
Output: true<br>
Explanation:<br>
The grid already contains a 2 x 2 square of the same color.<br>

### Constraints:

- `grid.length` == 3
- `grid[i].length` == 3
- `grid[i][j]` is either 'W' or 'B'.


### <i>Concepts Used:
- Array
- Matrix
- Enumeration </i>