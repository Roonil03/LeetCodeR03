# 3243. Shortest Distance After Road Addition Queries I
## Question Level: Medium
### Description:
You are given an integer n and a 2D integer array queries.

There are n cities numbered from 0 to `n - 1`. Initially, there is a unidirectional road from city `i` to city `i + 1` for all 0 <= `i` < `n - 1`.

`queries[i]` = `[ui, vi]` represents the addition of a new unidirectional road from city `ui` to city `vi`. After each query, you need to find the length of the shortest path from city 0 to city `n - 1`.

Return an array answer where for each i in the range `[0, queries.length - 1]`, `answer[i]` is the length of the shortest path from city 0 to city `n - 1` after processing the first `i + 1` queries.

### Examples:
#### Example 1:
Input: n = 5, queries = `[[2,4],[0,2],[0,4]]`<br>
Output: `[3,2,1]`<br>
Explanation:
<img src="https://assets.leetcode.com/uploads/2024/06/28/image8.jpg"><br>
```
After the addition of the road from 2 to 4, the length of the shortest path from 0 to 4 is 3.
```
-<img src="https://assets.leetcode.com/uploads/2024/06/28/image9.jpg"><br>
```
After the addition of the road from 0 to 2, the length of the shortest path from 0 to 4 is 2.
```
<img src="https://assets.leetcode.com/uploads/2024/06/28/image10.jpg"><br>
```
After the addition of the road from 0 to 4, the length of the shortest path from 0 to 4 is 1.
```
#### Example 2:
Input: n = 4, queries = `[[0,3],[0,2]]`<br>
Output: `[1,1]`<br>
Explanation:
<img src="https://assets.leetcode.com/uploads/2024/06/28/image11.jpg"><br>
```
After the addition of the road from 0 to 3, the length of the shortest path from 0 to 3 is 1.
```
<img src="https://assets.leetcode.com/uploads/2024/06/28/image12.jpg"><br>
```
After the addition of the road from 0 to 2, the length of the shortest path remains 1.
```

### Constraints:
- 3 <= `n` <= 500
- 1 <= `queries.length` <= 500
- `queries[i].length` == 2
- 0 <= `queries[i][0]` < `queries[i][1]` < n
- 1 < `queries[i][1]` - `queries[i][0]`
- There are no repeated roads among the queries.

### <i>Concepts Used:
- Array
- Breadth-First Search
- Graph </i>