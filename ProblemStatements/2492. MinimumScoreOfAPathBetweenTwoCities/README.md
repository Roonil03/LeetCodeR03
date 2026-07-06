# 2492. Minimum Score of a Path Between Two Cities
## Question Level: Medium
### Description:
You are given a positive integer n representing n cities numbered from 1 to n. You are also given a 2D array roads where `roads[i]` = `[a`<sub>`i`</sub>`, b`<sub>`i`</sub>`, distance`<sub>`i`</sub>`]` indicates that there is a bidirectional road between cities `a`<sub>`i`</sub> and `b`<sub>`i`</sub> with a distance equal to `distance`<sub>`i`</sub>. The cities graph is not necessarily connected.

The score of a path between two cities is defined as the minimum distance of a road in this path.

Return the minimum possible score of a path between cities 1 and n.

Note:
- A path is a sequence of roads between two cities.
- It is allowed for a path to contain the same road multiple times, and you can visit cities 1 and n multiple times along the path.
- The test cases are generated such that there is at least one path between 1 and n.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2022/10/12/graph11.png"><br>
Input: n = 4, roads = `[[1,2,9],[2,3,6],[2,4,5],[1,4,7]]`  
Output: 5  
Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 4. The score of this path is min(9,5) = 5.
It can be shown that no other path has less score.  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2022/10/12/graph22.png"><br>
Input: n = 4, roads = `[[1,2,2],[1,3,4],[3,4,7]]`  
Output: 2  
Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 1 -> 3 -> 4. The score of this path is min(2,2,4,7) = 2.  

### Constraints:

- 2 <= `n` <= 10<sup>5</sup>
- 1 <= `roads.length` <= 10<sup>5</sup>
- `roads[i].length` == 3
- 1 <= `a`<sub>`i`</sub>, `b`<sub>`i`</sub> <= `n`
- `a`<sub>`i`</sub> != `b`<sub>`i`</sub>
- 1 <= `distance`<sub>`i`</sub> <= 10<sup>4</sup>
- There are no repeated edges.
- There is at least one path between 1 and n.

### <i>Concepts Used:
- Depth-First Search
- Breadth-First Search
- Union-Find
- Graph Theory</i>