# 587. Erect the Fence
## Question Level: Hard
### Description:
You are given an array trees where `trees[i]` = `[xi, yi]` represents the location of a tree in the garden.

Fence the entire garden using the minimum length of rope, as it is expensive. The garden is well-fenced only if all the trees are enclosed.

Return the coordinates of trees that are exactly located on the fence perimeter. You may return the answer in any order.


### Examples:
#### Example 1:
<img src="https://assets.leetcode.com/uploads/2021/04/24/erect2-plane.jpg"><br>
Input: trees = `[[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]`<br>
Output: `[[1,1],[2,0],[4,2],[3,3],[2,4]]`<br>
Explanation: All the trees will be on the perimeter of the fence except the tree at `[2, 2]`, which will be inside the fence.<br>
#### Example 2:
<img src="https://assets.leetcode.com/uploads/2021/04/24/erect1-plane.jpg"><br>
Input: trees = `[[1,2],[2,2],[4,2]]`<br>
Output:` [[4,2],[2,2],[1,2]]`<br>
Explanation: The fence forms a line that passes through all the trees.<br>


### Constraints:

- 1 <= `trees.length` <= 3000
- `trees[i].length` == 2
- 0 <= `xi`, `yi` <= 100
- All the given positions are unique.

### <i>Concepts Used:
- Array
- Math
- Geometry </i>