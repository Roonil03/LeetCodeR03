# 310. Minimum Height Trees
## Question Level: Medium
### Description:
A tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.

Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where `edges[i]` = `[a`<sub>`i`</sub>`, b`<sub>`i`</sub>`]` indicates that there is an undirected edge between the two nodes a<sub>i</sub> and b<sub>i</sub> in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).

Return a list of all MHTs' root labels. You can return the answer in any order.

The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2020/09/01/e1.jpg"><br>
Input: n = 4, edges = `[[1,0],[1,2],[1,3]]`  
Output: `[1]`  
Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2020/09/01/e2.jpg"><br>
Input: n = 6, edges = `[[3,0],[3,1],[3,2],[3,4],[5,4]]`  
Output: `[3,4]`  

### Constraints:

- 1 <= `n` <= 2 * 10<sup>4</sup>
- `edges.length` == `n - 1`
- 0 <= `a`<sub>`i`</sub>, `b`<sub>`i`</sub> < `n`
- `a`<sub>`i`</sub> != `b`<sub>`i`</sub>
- All the pairs `(a`<sub>`i`</sub>`, b`<sub>`i`</sub>`)` are distinct.
- The given input is guaranteed to be a tree and there will be no repeated edges.

### <i>Concepts Used:
- Depth-First Search
- Breadth-First Search
- Graph
- Topological Sort</i>