# 3650. Minimum Cost Path with Edge Reversals
## Question Level: Medium
### Description:
You are given a directed, weighted graph with n nodes labeled from `0` to `n - 1`, and an array edges where `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`, w`<sub>`i`</sub>`]` represents a directed edge from node `u`<sub>`i`</sub> to node `v`<sub>`i`</sub> with cost `w`<sub>`i`</sub>.

Each node `u`<sub>`i`</sub> has a switch that can be used at most once: when you arrive at `u`<sub>`i`</sub> and have not yet used its switch, you may activate it on one of its incoming edges `v`<sub>`i`</sub>` → u`<sub>`i`</sub> reverse that edge to `u`<sub>`i`</sub>` → v`<sub>`i`</sub> and immediately traverse it.

The reversal is only valid for that single move, and using a reversed edge costs `2 * w`<sub>`i`</sub>.

Return the minimum total cost to travel from node `0` to node `n - 1`. If it is not possible, return `-1`.

### Examples:
#### Example 1:

Input: n = 4, edges = `[[0,1,3],[3,1,1],[2,3,4],[0,2,2]]`

Output: 5

Explanation:

<img src="https://assets.leetcode.com/uploads/2025/05/07/e1drawio.png"><br>

    - Use the path 0 → 1 (cost 3).
    - At node 1 reverse the original edge 3 → 1 into 1 → 3 and traverse it at cost 2 * 1 = 2.
    - Total cost is 3 + 2 = 5.
#### Example 2:

Input: n = 4, edges = `[[0,2,1],[2,1,1],[1,3,1],[2,3,3]]`

Output: 3

Explanation:
    - No reversal is needed. Take the path 0 → 2 (cost 1), then 2 → 1 (cost 1), then 1 → 3 (cost 1).
    - Total cost is 1 + 1 + 1 = 3.

### Constraints:

- 2 <= `n` <= 5 * 10<sup>4</sup>
- 1 <= `edges.length` <= 10<sup>5</sup>
- `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`, w`<sub>`i`</sub>`]`
- 0 <= `u`<sub>`i`</sub>, `v`<sub>`i`</sub> <= n - 1
- 1 <= `w`<sub>`i`</sub> <= 1000

### <i>Concepts Used:
- Graph Theory
- Heap (Priority Queue)
- Shortest Path</i>