# Weekly Contest 490
## Ranking: 202 / 43027
## Question 1:
*Find the Score Difference in a Game*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums, where `nums[i]` represents the points scored in the i<sup>th</sup> game.

There are exactly two players. Initially, the first player is active and the second player is inactive.

The following rules apply sequentially for each game i:
- If `nums[i]` is odd, the active and inactive players swap roles.
- In every 6th game (that is, game indices 5, 11, 17, ...), the active and inactive players swap roles.
- The active player plays the ith game and gains `nums[i]` points.

Return the score difference, defined as the first player's total score minus the second player's total score.

### Examples:
#### Example 1:

Input: nums = `[1,2,3]`

Output: 0

Explanation:​​​​​​​

Game 0: Since the points are odd, the second player becomes active and gains `nums[0]` = 1 point.  
Game 1: No swap occurs. The second player gains `nums[1]` = 2 points.  
Game 2: Since the points are odd, the first player becomes active and gains `nums[2]` = 3 points.  
The score difference is 3 - 3 = 0.   
#### Example 2:

Input: nums = `[2,4,2,1,2,1]`

Output: 4

Explanation:

Games 0 to 2: The first player gains 2 + 4 + 2 = 8 points.  
Game 3: Since the points are odd, the second player is now active and gains `nums[3]` = 1 point.  
Game 4: The second player gains `nums[4]` = 2 points.  
Game 5: Since the points are odd, the players swap roles. Then, because this is the 6th game, the players swap again. The second player gains `nums[5]` = 1 point.  
The score difference is 8 - 4 = 4.  
#### Example 3:

Input: nums = `[1]`

Output: -1

Explanation:

Game 0: Since the points are odd, the second player is now active and gains `nums[0]` = 1 point.  
The score difference is 0 - 1 = -1.  

### Constraints:

- 1 <= `nums.length` <= 1000
- 1 <= `nums[i]` <= 1000
 
## Question 2:
*Check Digitorial Permutation*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer n.

A number is called digitorial if the sum of the factorials of its digits is equal to the number itself.

Determine whether any permutation of n (including the original order) forms a digitorial number.

Return true if such a permutation exists, otherwise return false.

Note:
- The factorial of a non-negative integer x, denoted as x!, is the product of all positive integers less than or equal to x, and 0! = 1.
- A permutation is a rearrangement of all the digits of a number that does not start with zero. Any arrangement starting with zero is invalid.

### Examples:
#### Example 1:

Input: n = 145

Output: true

Explanation:

The number 145 itself is digitorial since 1! + 4! + 5! = 1 + 24 + 120 = 145. Thus, the answer is true.

#### Example 2:

Input: n = 10

Output: false

Explanation:​​​​​​​

10 is not digitorial since 1! + 0! = 2 is not equal to 10, and the permutation "01" is invalid because it starts with zero.


### Constraints:

- 1 <= `n` <= 10<sup>9</sup>

## Question 3:
*Maximum Bitwise XOR After Rearrangement*
### Difficulty: Medium
### Points: 5
### Description:
You are given two binary strings s and t​​​​​​​, each of length n.

You may rearrange the characters of t in any order, but s must remain unchanged.

Return a binary string of length n representing the maximum integer value obtainable by taking the bitwise `XOR` of s and rearranged t.

### Examples:
#### Example 1:

Input: s = "`101`", t = "`011`"

Output: "`110`"

Explanation:

One optimal rearrangement of t is "`011`".  
The bitwise XOR of s and rearranged t is "`101`" XOR "`011`" = "`110`", which is the maximum possible.  
#### Example 2:

Input: s = "`0110`", t = "`1110`"

Output: "`1101`"  

Explanation:

One optimal rearrangement of t is "`1011`".  
The bitwise XOR of s and rearranged t is "`0110`" XOR "`1011`" = "`1101`", which is the maximum possible.  
#### Example 3:

Input: s = "`0101`", t = "`1001`"

Output: "`1111`"

Explanation:

One optimal rearrangement of t is "`1010`".   
The bitwise XOR of s and rearranged t is "`0101`" `XOR` "`1010`" = "`1111`", which is the maximum possible.  
 
### Constraints:

- 1 <= `n` == `s.length` == `t.length` <= 2 * 10<sup>5</sup>
- `s[i]` and `t[i]` are either '`0`' or '`1`'.

## Question 4:
*Count Sequences to K*
### Difficulty: Hard
### Points: 6
### Description:
You are given an integer array `nums`, and an integer k.

Start with an initial value val = 1 and process nums from left to right. At each index i, you must choose exactly one of the following actions:
- Multiply `val` by `nums[i]`.
- Divide `val` by `nums[i]`.
- Leave `val` unchanged.

After processing all elements, val is considered equal to k only if its final rational value exactly equals k.

Return the count of distinct sequences of choices that result in val == k.

Note: Division is rational (exact), not integer division. For example, `2 / 4` = `1 / 2`.

### Examples:
### Example 1:

Input: nums = `[2,3,2]`, k = 6

Output: 2

Explanation:

The following 2 distinct sequences of choices result in val == k:
| Sequence | Operation on `nums[0]` | Operation on `nums[1]` | Operation on `nums[2] `| Final val |
| --- | --- | --- | --- | --- |
| 1 | Multiply: val = `1 * 2 = 2` | Multiply: val = `2 * 3 = 6` | Leave val unchanged | 6 |
| 2 | Leave val unchanged | Multiply: val = `1 * 3 = 3` | Multiply: val = `3 * 2 = 6` | 6 |

#### Example 2:

Input: nums = `[4,6,3]`, k = 2

Output: 2

Explanation:

The following 2 distinct sequences of choices result in val == k:
| Sequence | Operation on `nums[0]` | Operation on `nums[1]` | Operation on `nums[2]` | Final val |
| --- | --- | --- | --- | --- |
| 1 | Multiply: val = `1 * 4 = 4` | Divide: val = `4 / 6` = `2 / 3` | Multiply: val = `(2 / 3) * 3` = `2` | 2 |
| 2 | Leave val unchanged | Multiply: val = `1 * 6 = 6` | Divide: val = `6 / 3 = 2` | 2 |

#### Example 3:

Input: nums = `[1,5]`, k = 1

Output: 3

Explanation:

The following 3 distinct sequences of choices result in val == k:
| Sequence | Operation on `nums[0]` | Operation on `nums[1]` | Final val |
| --- | --- | --- | --- |
| 1 | Multiply: val = `1 * 1 = 1` | Leave val unchanged | 1 |
| 2 | Divide: val = `1 / 1 = 1` | Leave val unchanged | 1 |
| 3 | Leave val unchanged | Leave val unchanged | 1 |

### Constraints:

- 1 <= `nums.length` <= 19
- 1 <= `nums[i]` <= 6
- 1 <= `k` <= 10<sup>15</sup>