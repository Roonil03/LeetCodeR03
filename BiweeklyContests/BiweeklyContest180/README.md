# Biweekly Contest 180
## Ranking: 
## Question 1:
*Traffic Signal Color*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer timer representing the remaining time (in seconds) on a traffic signal.

The signal follows these rules:
- If `timer == 0`, the signal is "`Green`"
- If `timer == 30`, the signal is "`Orange`"
- If 30 < timer <= 90, the signal is "`Red`"

Return the current state of the signal. If none of the above conditions are met, return "`Invalid`".

### Examples:
#### Example 1:

Input: timer = 60

Output: "`Red`"

Explanation:

Since timer = 60, and 30 < timer <= 90, the answer is "`Red`".

#### Example 2:

Input: timer = 5

Output: "`Invalid`"

Explanation:

Since timer = 5, it does not satisfy any of the given conditions, the answer is "`Invalid`".

### Constraints:

- 0 <= `timer` <= 1000

## Question 2:
*Count Digit Appearances*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums and an integer digit.

Return the total number of times digit appears in the decimal representation of all elements in nums.

### Examples:
#### Example 1:

Input: nums = `[12,54,32,22]`, digit = 2

Output: 4

Explanation:

The digit 2 appears once in 12 and 32, and twice in 22. Thus, the total number of times digit 2 appears is 4.

#### Example 2:

Input: nums = `[1,34,7]`, digit = 9

Output: 0

Explanation:

The digit 9 does not appear in the decimal representation of any element in nums, so the total number of times digit 9 appears is 0.

### Constraints:

- 1 <= `nums.length` <= 1000
- 1 <= `nums[i]` <= 10<sup>6​​​​​​​</sup>
- 0 <= `digit` <= 9

## Question 3:
*Minimum Operations to Transform Array into Alternating Prime*
### Description:
You are given an integer array nums.

An array is considered alternating prime if:
- Elements at even indices (0-based) are prime numbers.
- Elements at odd indices are non-prime numbers.

In one operation, you may increment any element by 1.

Return the minimum number of operations required to transform nums into an alternating prime array.

A prime number is a natural number greater than 1 with only two factors, 1 and itself.

### Examples:
#### Example 1:

Input: nums = `[1,2,3,4]`

Output: 3

Explanation:

- The element at index 0 must be prime. Increment `nums[0] = 1` to 2, using 1 operation.
- The element at index 1 must be non-prime. Increment `nums[1] = 2` to 4, using 2 operations.
- The element at index 2 is already prime.
- The element at index 3 is already non-prime.

Total operations = 1 + 2 = 3.

#### Example 2:

Input: nums = `[5,6,7,8]`

Output: 0

Explanation:
- The elements at indices 0 and 2 are already prime.
- The elements at indices 1 and 3 are already non-prime.

No operations are needed.

#### Example 3:

Input: nums = `[4,4]`

Output: 1

Explanation:
- The element at index 0 must be prime. Increment `nums[0] = 4` to 5, using 1 operation.
- The element at index 1 is already non-prime.

Total operations = 1.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>5</sup>

## Question 4:
*Maximum Value of Concatenated Binary Segments*
### Difficulty: Hard
### Points: 6
### Description:
You are given two integer arrays nums1 and nums0, each of size n.

- `nums1[i]` represents the number of '`1`'s in the ith segment.
- `nums0[i]` represents the number of '`0`'s in the ith segment.

For each index i, construct a binary segment consisting of:
- `nums1[i]` occurrences of '`1`' followed by
- `nums0[i]` occurrences of '`0`'.

You may rearrange the order of these segments in any way. After rearranging, concatenate all segments to form a single binary string.

Return the maximum possible integer value of the concatenated binary string.

Since the result can be very large, return the answer modulo 10<sup>9</sup> + 7.

### Examples:
#### Example 1:

Input: nums1 = `[1,2]`, nums0 = `[1,0]`

Output: 14

Explanation:
- At index 0, `nums1[0]` = 1 and `nums0[0]` = 1, so the segment formed is "`10`".
- At index 1, `nums1[1]` = 2 and `nums0[1]` = 0, so the segment formed is "`11`".
- Reordering the segments as "`11`" followed by "10" produces the binary string "`1110`".
- The binary number "`1110`" has value 14 which is the maximum possible value.
#### Example 2:

Input: nums1 = `[3,1]`, nums0 = `[0,3]`

Output: 120

Explanation:
- At index 0, `nums1[0]` = 3 and `nums0[0]` = 0, so the segment formed is "`111`".
- At index 1, `nums1[1]` = 1 and `nums0[1]` = 3, so the segment formed is "`1000`".
- Reordering the segments as "`111`" followed by "`1000`" produces the binary string "`1111000`".
- The binary number "`1111000`" has value 120 which is the maximum possible value.

### Constraints:

- 1 <= `n` == `nums1.length` == `nums0.length` <= 10<sup>5</sup>
- 0 <= `nums1[i]`, `nums0[i]` <= 10<sup>4</sup>
- `nums1[i] + nums0[i]` > 0
- The total sum of all elements in `nums1` and `nums0` does not exceed 2 * 10<sup>5</sup>.

