# 1466. Reorder Routes to Make All Paths Lead to the City Zero
## Question Level: Medium
### Description:
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.

Roads are represented by connections where `connections[i]` = `[a`<sub>`i`</sub>`, b`<sub>`i`</sub>`]` represents a road from city a<sub>i</sub> to city b<sub>i</sub>.

This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.

It's guaranteed that each city can reach city 0 after reorder.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/05/13/sample_1_1819.png"><br>
Input: n = 6, connections = `[[0,1],[1,3],[2,3],[4,0],[4,5]]`  
Output: 3  
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2020/05/13/sample_2_1819.png"><br>
Input: n = 5, connections = `[[1,0],[1,2],[3,2],[3,4]]`  
Output: 2  
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).  
#### Example 3:

Input: n = 3, connections = `[[1,0],[2,0]]`  
Output: 0  

### Constraints:

- 2 <= n <= 5 * 10<sup>4</sup>
- `connections.length` == n - 1
- `connections[i].length` == 2
- 0 <= a<sub>i</sub>, b<sub>i</sub> - <= n - 1
- a<sub>i</sub> != b<sub>i</sub>


### <i>Concepts Used:
- Depth-First Search
- Breadth-First Search
- Graph</i>