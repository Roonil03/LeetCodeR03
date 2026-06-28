# Weekly Contest 508
## Ranking:
## Question 1:
*Maximum Total Sum of K Selected Elements*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums and two integers k and mul.

Select exactly k elements from nums. Process these elements one by one in any order you choose.

For each selected element, independently choose one of the following:
- Add the element's value to the total sum, or
- Multiply the element by the current value of mul and add the result to the total sum.

After processing each selected element, mul decreases by 1, regardless of which option was chosen. The current value of mul may become 0 or negative.

Return an integer denoting the maximum possible total sum.

### Examples:
#### Example 1:

Input: nums = `[6,1,2,9]`, k = 3, mul = 2

Output: 26

Explanation:

One optimal way:
- One optimal selection is `nums[3]` = 9, `nums[0]` = 6, and `nums[2]` = 2.
- Process `nums[3]` = 9 first: choose multiplication, so it contributes 9 * 2 = 18. Now, mul becomes 1.
- Process `nums[0]` = 6 next: choose multiplication, so it contributes 6 * 1 = 6. Now, mul becomes 0.
- Process `nums[2]` = 2 last: choose addition, so it contributes 2.  
- The total sum is 18 + 6 + 2 = 26.
#### Example 2:

Input: nums = `[3,7,5,2]`, k = 2, mul = 4

Output: 43

Explanation:

One optimal way:
- One optimal selection is `nums[1]` = 7 and `nums[2]` = 5.
- Process `nums[1]` = 7 first: choose multiplication, so it contributes 7 * 4 = 28. Now, mul becomes 3.
- Process nums[2] = 5 next: choose multiplication, so it contributes 5 * 3 = 15.  
- The total sum is 28 + 15 = 43.
#### Example 3:

Input: nums = `[4,4]`, k = 1, mul = 1

Output: 4

Explanation:

One optimal way:

- One optimal selection is `nums[0]` = 4.
- Process `nums[0]` = 4: choose multiplication, so it contributes 4 * 1 = 4.
- The total sum is 4.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>5</sup>
- 1 <= `k` <= `nums.length`
- 1 <= `mul` <= 10<sup>5</sup>

## Question 2:
*Filter Occupied Intervals*
### Difficulty: Medium
### Points: 4
### Description:
You are given a 2D integer array `occupiedIntervals`, where `occupiedIntervals[i]` = `[start`<sub>`i`</sub>`, end`<sub>`i`</sub>`]` represents a time interval during which you are occupied. Each interval starts at start<sub>i</sub> and ends at end<sub>i</sub>, inclusive. These intervals may overlap.

You are also given two integers freeStart and freeEnd, which define a free time interval from freeStart to freeEnd, inclusive.

Your task is to merge all occupied intervals that overlap or touch, then remove all integer points in the free interval from the merged occupied intervals.

Two intervals touch if the second interval starts immediately after the first one ends. For example, `[1, 1]` and `[2, 2]` touch and should be merged into `[1, 2]`.

Return the remaining occupied intervals in sorted order. The returned intervals must be non-overlapping and must contain the minimum number of intervals possible. If there are no remaining occupied points, return an empty list.

### Examples:
#### Example 1:

Input: occupiedIntervals = `[[2,6],[4,8],[10,10],[10,12],[14,16]]`, freeStart = 7, freeEnd = 11

Output: `[[2,6],[12,12],[14,16]]`

Explanation:

- After merging, the occupied intervals are `[2, 8]`, `[10, 12]`, and `[14, 16]`.
- Excluding the free interval `[7, 11]` results in `[2, 6]`, `[12, 12]`, and `[14, 16]`.
#### Example 2:

Input: occupiedIntervals = `[[1,5],[2,3]]`, freeStart = 3, freeEnd = 8

Output: `[[1,2]]`

Explanation:
- After merging, the occupied interval is `[1, 5]`.
- Excluding the free interval `[3, 8]` results in `[1, 2]`.

### Constraints:

- 1 <= `occupiedIntervals.length` <= 5 * 10<sup>4</sup>
- `occupiedIntervals[i].length` == 2
- 1 <= `start`<sub>`i`</sub> <= `end`<sub>`i`</sub> <= 10<sup>9</sup>
- 1 <= `freeStart` <= `freeEnd` <= 10<sup>9</sup>

## Question 3:
*Maximum Subarray Sum After Multiplier*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer array nums and a positive integer k.

You must choose exactly one subarray of nums and perform exactly one of the following operations:
1. Multiply each number in the chosen subarray by k.
2. Divide each number in the chosen subarray by k.
    - When dividing a positive number by k, use the floor value of the division result.
    - When dividing a negative number by k, use the ceiling value of the division result.

Return the maximum possible sum of a non-empty subarray in the resulting array.

Note that the subarray chosen for the operation and the subarray chosen for the sum may be different.

### Examples:
#### Example 1:

Input: nums = `[1,-2,3,4,-5]`, k = 2

Output: 14

Explanation:
- Multiply each number in the subarray `[3, 4]` by 2.
- This results in nums = `[1, -2, 6, 8, -5]`.  
- The subarray with the largest sum is `[6, 8]`, so the output is 6 + 8 = 14.  
#### Example 2:

Input: nums = `[-5,-4,-3]`, k = 2

Output: -1

Explanation:
- Divide each number in the subarray `[-3]` by 2.
- This results in nums = `[-5, -4, -1]`.
- The subarray with the largest sum is `[-1]`, so the output is -1.


### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- -10<sup>5</sup> <= `nums[i]` <= 10<sup>5</sup>?
- 1 <= `k` <= 10<sup>5</sup>

## Question 4:
*Minimum Time to Reach Target With Limited Power*
### Difficulty: Hard
### Points: 6
### Description:
You are given a directed weighted graph with n nodes labeled from 0 to n - 1.

The graph is represented by a 2D integer array edges, where `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`, t`<sub>`i`</sub>`]` indicates a directed edge from node u<sub>i</sub> to node v<sub>i</sub> that takes t<sub>i</sub> seconds to traverse.

You are also given an integer power representing the initial available power, and an integer array cost of length n, where `cost[u]` represents the power required to forward the signal from node u through any one of its outgoing edges.

You are given two integers source and target.

The signal starts at source at time 0 with power units of power and follows these rules:
- The signal may traverse a directed edge from node u only if the remaining power is at least `cost[u].`
- No power is consumed when the signal arrives at a node, unless it later leaves that node by traversing another edge.
- When the signal is forwarded from node u, the remaining power is decreased by `cost[u]` units.
- Traversing an edge `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`, t`<sub>`i`</sub>`]` increases the total time by t<sub>i</sub> seconds.

Return an integer array answer of size 2, where:
- `answer[0]` is the minimum time required for the signal to reach node target.
- `answer[1]` is the maximum remaining power among all paths that achieve `answer[0]`.

If the signal cannot reach target, return `[-1, -1]`.

### Examples:
#### Example 1:

<img src="https://assets.leetcode.com/uploads/2026/02/09/g1.png"><br>
Input: n = 5, edges = `[[0,1,1],[1,4,1],[0,2,1],[2,3,1],[3,4,1]]`, power = 4, cost = `[2,3,1,1,1]`, source = 0, target = 4

Output: `[3,0]`

Explanation:
- The signal starts at node 0 with 4 units of power.
- The path `0 -> 1 -> 4` is not valid, because after leaving node 0, the signal has 2 units of power remaining, which is less than `cost[1]` = 3.
- The valid path `0 -> 2 -> 3 -> 4` takes a total time of 3.
- The total power consumed along this path is `cost[0] + cost[2] + cost[3] = 4`, leaving 0 remaining power.
- Hence, the answer is `[3, 0]`.

#### Example 2:

<img src="https://assets.leetcode.com/uploads/2026/02/09/g22.png"><br>

Input: n = 3, edges = `[[0,1,2],[1,2,2],[2,0,2]]`, power = 3, cost = `[1,1,1]`, source = 1, target = 1

Output: `[0,3]`

Explanation:
- Since the source and target are the same node, no traversal is required.
- Hence, the minimum total time taken is 0, and no power is consumed.
- Therefore, the answer is `[0, 3]`.

#### Example 3:

​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/02/09/g23.png"><br>

Input: n = 4, edges = `[[0,1,3],[2,3,4]]`, power = 3, cost = `[1,1,1,1]`, source = 0, target = 3

Output: `[-1,-1]`

Explanation:

There is no valid path from source to target, therefore return `[-1, -1]`.

### Constraints:

- 1 <= `n` <= 1000
- 0 <= `edges.length` <= 1000
- `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`, t`<sub>`i`</sub>`]`
- 0 <= `u`<sub>`i`</sub>, `v`<sub>`i`</sub> <= `n - 1`
- 1 <= `t`<sub>`i`</sub> <= 109
- 1 <= `power` <= 1000
- `cost.length` == `n`
- 1 <= `cost[i]` <= 2000
- 0 <= `source`, `target` <= `n - 1`

