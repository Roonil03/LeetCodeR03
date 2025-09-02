# 391. Perfect Rectangle
## Question Level: Hard
### Description:
Given an array rectangles where `rectangles[i]` = `[x`<sub>`i</`sub>`, y`<sub>`i</`sub>`, a`<sub>`i</`sub>`, b`<sub>`i</`sub>`]` represents an axis-aligned rectangle. The bottom-left point of the rectangle is `(x`<sub>`i</`sub>`, y`<sub>`i</`sub>`)` and the top-right point of it is `(a`<sub>`i</`sub>`, b`<sub>`i</`sub>`)`.

Return true if all the rectangles together form an exact cover of a rectangular region.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/03/27/perectrec1-plane.jpg"><br>
Input: rectangles = `[[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]`  
Output: true  
Explanation: All 5 rectangles together form an exact cover of a rectangular region.  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/03/27/perfectrec2-plane.jpg"><br>
Input: rectangles = `[[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]`  
Output: false  
Explanation: Because there is a gap between the two rectangular regions.  
#### Example 3:

<img src="https://assets.leetcode.com/uploads/2021/03/27/perfectrec2-plane.jpg"><br>
Input: rectangles = `[[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]`  
Output: false  
Explanation: Because two of the rectangles overlap with each other.  
 
### Constraints:

- 1 <= `rectangles.length` <= 2 * 10<sup>4</sup>
- `rectangles[i].length` == 4
- -10<sup>5</sup> <= `x`<sub>`i`</sub> < `a`<sub>`i`</sub> <= 10<sup>5</sup>
- -10<sup>5</sup> <= `y`<sub>`i`</sub> < `b`<sub>`i`</sub> <= 10<sup>5</sup>

### <i>Concepts Used:
- Array
- Hash Table
- Math
- Geometry
- Line Sweep</i>