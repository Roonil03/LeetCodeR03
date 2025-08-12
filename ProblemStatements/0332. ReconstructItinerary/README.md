# 332. Reconstruct Itinerary
## Question Level: Hard
### Description:
You are given a list of airline tickets where `tickets[i]` = `[from`<sub>`i`</sub>`, to`<sub>`i`</sub>`]` represent the departure and the arrival airports of one flight. Reconstruct the itinerary in order and return it.

All of the tickets belong to a man who departs from "`JFK`", thus, the itinerary must begin with "`JFK`". If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.
- For example, the itinerary `["JFK", "LGA"]` has a smaller lexical order than `["JFK", "LGB"]`.

You may assume all tickets form at least one valid itinerary. You must use all the tickets once and only once.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2021/03/14/itinerary1-graph.jpg"><br>
Input: tickets = `[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]`  
Output: `["JFK","MUC","LHR","SFO","SJC"]`  
#### Example 2:

<img src="https://assets.leetcode.com/uploads/2021/03/14/itinerary2-graph.jpg"><br>
Input: tickets = `[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]`  
Output: `["JFK","ATL","JFK","SFO","ATL","SFO"]`  
Explanation: Another possible reconstruction is `["JFK","SFO","ATL","JFK","ATL","SFO"]` but it is larger in lexical order.  

### Constraints:

- 1 <= `tickets.length` <= 300
- `tickets[i].length` == 2
- `from`<sub>`i`<sub>.length` == 3
- `to`<sub>`i`<sub>.length` == 3
- `from`<sub>`i`<sub> and `to`<sub>`i`<sub> consist of uppercase English letters.
- `from`<sub>`i`<sub> != `to`<sub>`i`<sub>

### <i>Concepts Used:
- Depth-First Search
- Graph
- Eulerian Circuit</i>