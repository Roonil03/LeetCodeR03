# 1129. Shortest Path with Alternating Colors
## Question Level: Medium
### Description:
You are given an integer n, the number of nodes in a directed graph where the nodes are labeled from 0 to n - 1. Each edge is red or blue in this graph, and there could be self-edges and parallel edges.

You are given two arrays redEdges and blueEdges where:
- `redEdges[i]` = `[a`<sub>`i`</sub>`, b`<sub>`i`</sub>`]` indicates that there is a directed red edge from node a<sub>i</sub> to node b<sub>i</sub> in the graph, and
- `blueEdges[j]` = `[u`<sub>`j`</sub>`, v`<sub>`j`</sub>`]` indicates that there is a directed blue edge from node u<sub>j</sub> to node v<sub>j</sub> in the graph.

Return an array answer of length n, where each `answer[x]` is the length of the shortest path from node 0 to node x such that the edge colors alternate along the path, or -1 if such a path does not exist.

### Examples:
#### Example 1:

Input: n = 3, redEdges = `[[0,1],[1,2]]`, blueEdges = `[]`  
Output: `[0,1,-1]`  
#### Example 2:

Input: n = 3, redEdges = `[[0,1]]`, blueEdges = `[[2,1]]`  
Output: `[0,1,-1]`  

### Constraints:

- 1 <= `n` <= 100
- 0 <= `redEdges.length`, `blueEdges.length` <= 400
- `redEdges[i].length `== `blueEdges[j].length` == 2
- 0 <= `a`<sub>`i`</sub>, `b`<sub>`i`</sub>, `u`<sub>`j`</sub>, `v`<sub>`j`</sub> < n

### <i>Concepts Used:
- Breadth-First Search
- Graph</i>