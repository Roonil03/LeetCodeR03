# 1192. Critical Connections in a Network
## Question Level: Hard
### Description:
There are n servers numbered from 0 to n - 1 connected by undirected server-to-server connections forming a network where `connections[i]` = `[a`<sub>`i`</sub>`, b`<sub>`i`</sub>`]` represents a connection between servers a<sub>i</sub> and b<sub>i</sub>. Any server can reach other servers directly or indirectly through the network.

A critical connection is a connection that, if removed, will make some servers unable to reach some other server.

Return all critical connections in the network in any order.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2019/09/03/1537_ex1_2.png"><br>
Input: n = 4, connections = `[[0,1],[1,2],[2,0],[1,3]]`  
Output: `[[1,3]]`  
Explanation: `[[3,1]]` is also accepted.  
#### Example 2:

Input: n = 2, connections = `[[0,1]]`  
Output: `[[0,1]]`  

### Constraints:

- 2 <= `n` <= 10<sup>5</sup>
- `n - 1` <= `connections.length` <= 10<sup>5</sup>
- 0 <= a<sub>i</sub>, b<sub>i</sub> <= `n - 1`
- a<sub>i</sub> != b<sub>i</sub>
- There are no repeated connections.

### <i>Concepts Used:
- Depth-First Search
- Graph
- Biconnected Component</i>