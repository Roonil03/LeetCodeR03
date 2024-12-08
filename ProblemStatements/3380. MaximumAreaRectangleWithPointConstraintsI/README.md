# 3380. Maximum Area Rectangle With Point Constraints I
## Question Level: Medium
### Description:
You are given an array points where `points[i]` = `[xi, yi]` represents the coordinates of a point on an infinite plane.

Your task is to find the maximum area of a rectangle that:
- Can be formed using four of these points as its corners.
- Does not contain any other point inside or on its border.
- Has its edges parallel to the axes.

Return the maximum area that you can obtain or -1 if no such rectangle is possible.

### Examples:
#### Example 1:
Input: points = `[[1,1],[1,3],[3,1],[3,3]]`<br>
Output: 4<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png"><br>
We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.

#### Example 2:
Input: points = `[[1,1],[1,3],[3,1],[3,3],[2,2]]`<br>
Output: -1<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png"><br>
There is only one rectangle possible is with points `[1,1]`, `[1,3]`, `[3,1]` and `[3,3]` but `[2,2]` will always lie inside it. Hence, returning -1.

#### Example 3:
Input: points = `[[1,1],[1,3],[3,1],[3,3],[1,2],[3,2]]`<br>
Output: 2<br>
Explanation:<br>
<img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png"><br>
The maximum area rectangle is formed by the points `[1,3]`, `[1,2]`, `[3,2]`, `[3,3]`, which has an area of 2. Additionally, the points `[1,1]`, `[1,2]`, `[3,1]`, `[3,2]` also form a valid rectangle with the same area.

### Constraints:
- 1 <= `points.length` <= 10
- `points[i].length` == 2
- 0 <= `xi`, `yi` <= 100
- All the given points are unique.

### <i>Weekly Contest Question</i>