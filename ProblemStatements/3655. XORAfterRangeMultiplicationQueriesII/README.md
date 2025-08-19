# 3655. XOR After Range Multiplication Queries II
## Question Level: Hard
### Description:
You are given an integer array nums of length n and a 2D integer array queries of size q, where `queries[i]` = `[l`<sub>`i`</sub>`, r`<sub>`i`</sub>`, k`<sub>`i`</sub>`, v`<sub>`i`</sub>`]`.

For each query, you must apply the following operations in order:
- Set idx = l<sub>i</sub>.
While idx <= r<sub>i</sub>:
    - Update: `nums[idx]` = `(nums[idx] * vi) % (10`<sup>`9`</sup>` + 7)`.
    - Set `idx += k`<sub>`i`</sub>.

Return the bitwise `XOR` of all elements in nums after processing all queries.

### Examples:
#### Example 1:

Input: nums = `[1,1,1]`, queries = `[[0,2,1,4]]`

Output: 4

Explanation:

- A single query `[0, 2, 1, 4]` multiplies every element from index 0 through index 2 by 4.
- The array changes from `[1, 1, 1]` to `[4, 4, 4]`.
- The XOR of all elements is `4 ^ 4 ^ 4` = `4`.
#### Example 2:

Input: nums = `[2,3,1,5,4]`, queries = `[[1,4,2,3],[0,2,1,2]]`

Output: 31

Explanation:

- The first query `[1, 4, 2, 3]` multiplies the elements at indices 1 and 3 by 3, transforming the array to `[2, 9, 1, 15, 4]`.
- The second query `[0, 2, 1, 2]` multiplies the elements at indices 0, 1, and 2 by 2, resulting in `[4, 18, 2, 15, 4]`.
- Finally, the XOR of all elements is `4 ^ 18 ^ 2 ^ 15 ^ 4` = `31`.​​​​​​​

### Constraints:

- 1 <= `n` == `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>9</sup>
- 1 <= `q` == `queries.length` <= 10<sup>5​​​​​​​</sup>
- `queries[i]` = `[l`<sub>`i`</sub>`, r`<sub>`i`</sub>`, k`<sub>`i`</sub>`, v`<sub>`i`</sub>`]`
- 0 <= `l`<sub>`i`</sub> <= `r`<sub>`i`</sub> < `n`
- 1 <= `k`<sub>`i`</sub> <= `n`
- 1 <= `v`<sub>`i`</sub> <= 10<sup>5</sup>

### <i>Concepts Used:
- Array
- Divide and Conquer</i>