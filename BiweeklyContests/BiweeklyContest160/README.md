# Biweekly Contest 160
## Ranking:
## Question 1:
*Hexadecimal and Hexatrigesimal Conversion*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer n.

Return the concatenation of the hexadecimal representation of n<sup>2</sup> and the hexatrigesimal representation of n<sup>3</sup>.

A hexadecimal number is defined as a base-16 numeral system that uses the digits 0 – 9 and the uppercase letters A - F to represent values from 0 to 15.

A hexatrigesimal number is defined as a base-36 numeral system that uses the digits 0 – 9 and the uppercase letters A - Z to represent values from 0 to 35.

### Testcases:
#### Example 1:

Input: n = 13

Output: "`A91P1`"

Explanation:

- n<sup>2</sup> = 13 * 13 = 169. In hexadecimal, it converts to (10 * 16) + 9 = 169, which corresponds to "`A9`".
- n<sup>3</sup> = 13 * 13 * 13 = 2197. In hexatrigesimal, it converts to (1 * 36<sup>2</sup>) + (25 * 36) + 1 = 2197, which corresponds to "`1P1`".
- Concatenating both results gives "`A9`" + "`1P1`" = "`A91P1`".
#### Example 2:

Input: n = 36

Output: "`5101000`"

Explanation:

- n<sup>2</sup> = 36 * 36 = 1296. In hexadecimal, it converts to (5 * 16<sup>2</sup>) + (1 * 16) + 0 = 1296, which corresponds to "`510`".  
- n<sup>3</sup> = 36 * 36 * 36 = 46656. In hexatrigesimal, it converts to (1 * 36<sup>3</sup>) + (0 * 36<sup>2</sup>) + (0 * 36) + 0 = 46656, which corresponds to "`1000`".  
- Concatenating both results gives "`510`" + "`1000`" = "`5101000`".

### Constraints:

- 1 <= n <= 1000

## Question 2:
*Minimum Cost Path with Alternating Directions II*
### Difficulty: Medium
### Points: 5
### Description:
You are given two integers m and n representing the number of rows and columns of a grid, respectively.

The cost to enter cell `(i, j)` is defined as `(i + 1) * (j + 1)`.

You are also given a 2D integer array waitCost where `waitCost[i][j]` defines the cost to wait on that cell.

You start at cell (0, 0) at second 1.

At each step, you follow an alternating pattern:
- On odd-numbered seconds, you must move right or down to an adjacent cell, paying its entry cost.
- On even-numbered seconds, you must wait in place, paying `waitCost[i][j]`.

Return the minimum total cost required to reach `(m - 1, n - 1)`.

### Testcases:
#### Example 1:

Input: m = 1, n = 2, waitCost = `[[1,2]]`

Output: 3

Explanation:

The optimal path is:

- Start at cell `(0, 0)` at second 1 with entry cost `(0 + 1) * (0 + 1)` = 1.
- Second 1: Move right to cell `(0, 1)` with entry cost `(0 + 1) * (1 + 1)` = 2.  
Thus, the total cost is 1 + 2 = 3.

#### Example 2:

Input: m = 2, n = 2, waitCost = `[[3,5],[2,4]]`

Output: 9

Explanation:

The optimal path is:

- Start at cell `(0, 0)` at second 1 with entry cost `(0 + 1) * (0 + 1)` = 1.
- Second 1: Move down to cell `(1, 0)` with entry cost `(1 + 1) * (0 + 1)` = 2.
- Second 2: Wait at cell `(1, 0)`, paying `waitCost[1][0]` = 2.
- Second 3: Move right to cell `(1, 1)` with entry cost `(1 + 1) * (1 + 1)` = 4.  
Thus, the total cost is 1 + 2 + 2 + 4 = 9.

#### Example 3:

Input: m = 2, n = 3, waitCost = `[[6,1,4],[3,2,5]]`

Output: 16

Explanation:

The optimal path is:

- Start at cell `(0, 0)` at second 1 with entry cost `(0 + 1) * (0 + 1)` = 1.
- Second 1: Move right to cell `(0, 1)` with entry cost `(0 + 1) * (1 + 1)` = 2.
- Second 2: Wait at cell `(0, 1)`, paying `waitCost[0][1]` = 1.
- Second 3: Move down to cell `(1, 1)` with entry cost `(1 + 1) * (1 + 1)` = 4.
- Second 4: Wait at cell `(1, 1)`, paying `waitCost[1][1]` = 2.
- Second 5: Move right to cell `(1, 2)` with entry cost `(1 + 1) * (2 + 1)` = 6.  
Thus, the total cost is 1 + 2 + 1 + 4 + 2 + 6 = 16.

### Constraints:

- 1 <= `m`, `n` <= 10<sup>5</sup>
- 2 <= `m * n` <= 10<sup>5</sup>
- `waitCost.length` == `m`
- `waitCost[0].length` == `n`
- 0 <= `waitCost[i][j]` <= 10<sup>5</sup>

## Question 3:
*Minimum Time to Reach Destination in Directed Graph*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer n and a directed graph with n nodes labeled from 0 to n - 1. This is represented by a 2D array edges, where `edges[i]` = `[u`<sup>`i`</sup>`, v`<sup>`i`</sup>`, start`<sup>`i`</sup>`, end`<sup>`i`</sup>`]` indicates an edge from node u<sup>i</sup> to v<sup>i</sup> that can only be used at any integer time t such that start<sup>i</sup> <= t <= end<sup>i</sup>.

You start at node 0 at time 0.

In one unit of time, you can either:
- Wait at your current node without moving, or
- Travel along an outgoing edge from your current node if the current time t satisfies start<sup>i</sup> <= t <= end<sup>i</sup>.

Return the minimum time required to reach node n - 1. If it is impossible, return -1.

### Testcases:
#### Example 1:

Input: n = 3, edges = `[[0,1,0,1],[1,2,2,5]]`

Output: 3

Explanation:

<img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004535.png"><br>

The optimal path is:

At time t = 0, take the edge `(0 → 1)` which is available from 0 to 1. You arrive at node 1 at time t = 1, then wait until t = 2.
At time t = 2, take the edge `(1 → 2)` which is available from 2 to 5. You arrive at node 2 at time 3.
Hence, the minimum time to reach node 2 is 3.

Example 2:

Input: n = 4, edges = `[[0,1,0,3],[1,3,7,8],[0,2,1,5],[2,3,4,7]]`

Output: 5

Explanation:

<img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004757.png"><br>

The optimal path is:

Wait at node 0 until time t = 1, then take the edge `(0 → 2)` which is available from 1 to 5. You arrive at node 2 at t = 2.
Wait at node 2 until time t = 4, then take the edge `(2 → 3)` which is available from 4 to 7. You arrive at node 3 at t = 5.
Hence, the minimum time to reach node 3 is 5.

Example 3:

Input: n = 3, edges = `[[1,0,1,3],[1,2,3,5]]`

Output: -1

Explanation:

<img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004914.png"><br>

Since there is no outgoing edge from node 0, it is impossible to reach node 2. Hence, the output is -1.

### Constraints:

- 1 <= `n` <= 10<sup>5</sup>
- 0 <= `edges.length` <= 10<sup>5</sup>
- `edges[i]` == `[u`<sup>`i`</sup>`, v`<sup>`i`</sup>`, start`<sup>`i`</sup>`, end`<sup>`i`</sup>`]`
- 0 <= `u`<sup>`i`</sup>`, v`<sup>`i`</sup>` <= n - 1
- `u`<sup>`i`</sup> != `v`<sup>`i`</sup>
- 0 <= `start`<sup>`i`</sup> <= `end`<sup>`i`</sup> <= 10<sp>9</sup>

## Question 4:
*Minimum Stability Factor of Array*
### Difficulty: Hard
### Points: 7
### Description:
You are given an integer array nums and an integer maxC.

A subarray is called stable if the highest common factor `(HCF)` of all its elements is greater than or equal to 2.

The stability factor of an array is defined as the length of its longest stable subarray.

You may modify at most `maxC` elements of the array to any integer.

Return the minimum possible stability factor of the array after at most `maxC` modifications. If no stable subarray remains, return 0.

Note:

- The highest common factor (HCF) of an array is the largest integer that evenly divides all the array elements.
- A subarray of length 1 is stable if its only element is greater than or equal to 2, since `HCF([x])` = x.


### Testcases:
#### Example 1:

Input: nums = `[3,5,10]`, maxC = 1

Output: 1

Explanation:

- The stable subarray `[5, 10]` has HCF = 5, which has a stability factor of 2.
- Since maxC = 1, one optimal strategy is to change `nums[1]` to 7, resulting in nums = `[3, 7, 10]`.
- Now, no subarray of length greater than 1 has HCF >= 2. Thus, the minimum possible stability factor is 1.
### Example 2:

Input: nums = `[2,6,8]`, maxC = 2

Output: 1

Explanation:

- The subarray `[2, 6, 8]` has HCF = 2, which has a stability factor of 3.
- Since maxC = 2, one optimal strategy is to change `nums[1]` to 3 and `nums[2]` to 5, resulting in nums = `[2, 3, 5]`.
- Now, no subarray of length greater than 1 has HCF >= 2. Thus, the minimum possible stability factor is 1.
#### Example 3:

Input: nums = `[2,4,9,6]`, maxC = 1

Output: 2

Explanation:

- The stable subarrays are:
    - `[2, 4]` with HCF = 2 and stability factor of 2.
    - `[9, 6]` with HCF = 3 and stability factor of 2.
- Since maxC = 1, the stability factor of 2 cannot be reduced due to two separate stable subarrays. Thus, the minimum possible stability factor is 2.


### Constraints:

- 1 <= `n` == `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>9</sup>
- 0 <= `maxC` <= `n`
