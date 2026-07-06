# Weekly Contest 509
## Ranking:
## Question 1:
*Sum of Integers with Maximum Digit Range*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array `nums`.

The digit range of an integer is defined as the difference between its largest digit and smallest digit.

For example, the digit range of `5724` is `7 - 2` = `5`.

Return the sum of all integers in nums whose digit range is equal to the maximum digit range among all integers in the array.

### Examples:
#### Example 1:

Input: nums = `[5724,111,350]`

Output: 6074

Explanation:  
| i | `nums[i]` | Largest | Smallest | Digit Range |
| -- | -- | -- | -- | -- |
| 0 | 5724 | 7 | 2 | 5 |
| 1 | 111 | 1 | 1 | 0 |
| 2 | 350 | 5 | 0 | 5 |  

The maximum digit range is 5. The integers with this digit range are 5724 and 350, so the answer is 5724 + 350 = 6074.

#### Example 2:

Input: nums = `[90,900]`

Output: 990

Explanation:

| i | `nums[i]` | Largest | Smallest | Digit Range |
| -- | -- | -- | -- | -- |
| 0 | 90 | 9 | 0 | 9 |
| 1 | 900 | 9 | 0 | 9 |

The maximum digit range is 9. Both integers have this digit range, so the answer is 90 + 900 = 990.

### Constraints:

- 1 <= `nums.length` <= 100
- 10 <= `nums[i]` <= 10<sup>5</sup>

## Question 2:
*Subsequence After One Replacement*
### Difficulty: Medium
### Points: 5
### Description:
You are given two strings s and t consisting of lowercase English letters.

You may choose at most one index in s and replace the character at that index with any lowercase English letter.

Return true if it is possible to make s a subsequence of t; otherwise, return false.

### Examples:
#### Example 1:

Input: s = "`cat`", t = "`chat`"

Output: true

Explanation:
- Replace `s[1]` from '`a`' to '`h`'. The resulting string is "`cht`".
- "`cht`" is a subsequence of "`chat`" because we can match '`c`', '`h`', and '`t`' in order.
#### Example 2:

Input: s = "`plane`", t = "`apple`"

Output: false

Explanation:
- The characters '`p`', '`l`', and '`e`' can be matched in t, but the remaining characters cannot be matched while preserving the required order.
- Even after replacing any one character in `s`, it is impossible to make `s` a subsequence of `t`.

### Constraints:

- 1 <= `s.length`, `t.length` <= 10<sup>5</sup>
- `s` and `t` consist only of lowercase English letters.

## Question 3:
*Divisible Game*
### Difficulty: Medium
### Points: 3
### Description:
You are given an integer array nums of length n.

Alice and Bob are playing a game. Alice chooses:
- An integer k such that` k > 1`.
- Two integers l and r such that `0 <= l <= r < n`.

Initially, both Alice's and Bob's scores are 0.

For each index i in the range `[l, r]` (inclusive):
- If `nums[i]` is divisible by k, Alice's score increases by `nums[i]`.
- Otherwise, Bob's score increases by `nums[i]`.

The score difference is Alice's score minus Bob's score.

Alice wants to maximize the score difference. If there are multiple values of k that achieve the maximum score difference, she chooses the smallest such k.

Return the product of the maximum score difference and the chosen value of k. Since the result can be large, return it modulo 10<sup>9</sup> + 7.

### Examples:
#### Example 1:

Input: nums = `[1,4,6,8]`

Output: 36

Explanation:
- Alice can choose k = 2, l = 1, and r = 3.
- All values in `nums[1..3]` are divisible by 2, so Alice's score is 4 + 6 + 8 = 18, while Bob's score is 0.
- The score difference is 18, which is the maximum possible. Among all values of k that achieve this score difference, the smallest is 2.
- Therefore, the answer is 18 * 2 = 36.
#### Example 2:

Input: nums = `[2,1,2]`

Output: 6

Explanation:
- Alice can choose `k = 2`, `l = 0`, and `r = 2`.
- The values `nums[0]` and `nums[2]` are divisible by 2, so Alice's score is 2 + 2 = 4. The value `nums[1]` is not divisible by 2, so Bob's score is 1.
- The score difference is 4 - 1 = 3, which is the maximum possible. Among all values of k that achieve this score difference, the smallest is 2.
- Therefore, the answer is `3 * 2` = 6.
#### Example 3:

Input: nums = `[1]`

Output: 1000000005

Explanation:
- Alice must choose some k > 1. The smallest possible choice is k = 2.
- Since `nums[0]` is not divisible by 2, Alice's score is 0, while Bob's score is 1.
- The score difference is -1, which is the maximum possible.
- Therefore, the answer is -1 * 2 = -2. Modulo 10<sup>9</sup> + 7, this equals 1000000005.

### Constraints:

- 1 <= `nums.length` <= 1000
- 1 <= `nums[i]` <= 10<sup>6</sup>

## Question 4:
*Palindromic Subarray Sum*
### Difficulty: Hard
### Points: 6
### Description:
You are given an integer array nums.

Return the maximum possible sum of a subarray of nums that is a palindrome.

### Examples:
#### Example 1:

Input: nums = `[10,10]`

Output: 20

Explanation:  
The whole array `[10,10]` is a palindrome. Therefore, the maximum sum is 10 + 10 = 20.

#### Example 2:

Input: nums = `[1,2,3,2,1,5,6]`

Output: 9

Explanation:  
The contiguous subarray `[1,2,3,2,1]` is a palindrome. Its sum is 1 + 2 + 3 + 2 + 1 = 9 and it is the maximum sum.

#### Example 3:

Input: nums = `[7,1,2,1,7,3,4,3,4]`

Output: 18

Explanation:  
The contiguous subarray `[7,1,2,1,7] `is a palindrome. Its sum is 7 + 1 + 2 + 1 + 7 = 18 and it is the maximum sum.

#### Example 4:

Input: nums = `[1,2,3,4,5]`

Output: 5

Explanation:  
No subarray with length greater than 1 is a palindrome. The largest element in the array is 5. Therefore, the answer is 5.

#### Example 5:

Input: nums = `[1000]`

Output: 1000

Explanation:  
The subarray with only one element is a palindrome. Therefore, the answer is 1000.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10​​​​​​​<sup>9</sup>
