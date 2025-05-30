# 2672. Number of Adjacent Elements With the Same Color
## Question Level: Medium
### Description:
You are given an integer n representing an array colors of length n where all elements are set to 0's meaning uncolored. You are also given a 2D integer array queries where `queries[i]` = `[index`<sub>`i`</sub>`, color`<sub>`i`</sub>`]`. For the ith query:
- Set `colors[index`<sub>`i`</sub>`]` to color<sub>i</sub>.
- Count adjacent pairs in colors set to the same color (regardless of color<sub>i</sub>).

Return an array answer of the same length as queries where `answer[i]` is the answer to the ith query.

### Examples:
#### Example 1:

Input: n = 4, queries = `[[0,2],[1,2],[3,1],[1,1],[2,1]]`  
Output: `[0,1,1,0,2]`  
Explanation:  
- Initially array colors = `[0,0,0,0]`, where 0 denotes uncolored elements of the array.
- After the 1st query colors = `[2,0,0,0]`. The count of adjacent pairs with the same color is 0.
- After the 2nd query colors = `[2,2,0,0]`. The count of adjacent pairs with the same color is 1.
- After the 3rd query colors = `[2,2,0,1]`. The count of adjacent pairs with the same color is 1.
- After the 4th query colors = `[2,1,0,1]`. The count of adjacent pairs with the same color is 0.
- After the 5th query colors = `[2,1,1,1]`. The count of adjacent pairs with the same color is 2.

#### Example 2:

Input: n = 1, queries = `[[0,100000]]`
Output: `[0]`
Explanation:
- After the 1st query colors = `[100000]`. The count of adjacent pairs with the same color is 0.

### Constraints:

- 1 <= n <= 10<sup>5</sup>
- 1 <= `queries.length` <= 10<sup>5</sup>
- `queries[i].length` == 2
- 0 <= index<sub>i</sub> <= n - 1
- 1 <=  color<sub>i <= 10<sup>5</sup>

### <i>Concepts Used:
- Array</i>