# Biweekly Contest 186
## Ranking: 3966 / 39050
## Question 1:
*Unique Middle Element*
### Difficulty: Easy
### Points: 3
### Description:
You are given an integer array nums of odd length n.

Return true if the middle element of nums appears exactly once in the array. Otherwise return false.

### Examples:
#### Example 1:

Input: nums = `[1,2,3]`

Output: true

Explanation:

The middle element of nums is 2, which appears exactly once.

Thus, the answer is true.

#### Example 2:

Input: nums = `[1,2,2]`

Output: false

Explanation:

The middle element of nums is 2, which appears twice.

Thus, the answer is false.

### Constraints:

- 1 <= `n` == `nums.length` <= 100
- `n` is odd.
- 1 <= `nums[i]` <= 100

## Question 2:
*Maximum Valid Pair Sum*
### Difficulty: Medium
### Points: 4
### Description:
You are given an integer array nums of length n and an integer k.

A pair of indices `(i, j)` is called valid if:
- 0 <= `i` < j < n
- j - i >= k

Return the maximum value of `nums[i] + nums[j]` among all valid pairs.

### Examples:
#### Example 1:

Input: nums = `[1,3,5,2,8]`, k = 2

Output: 13

Explanation:

The valid pairs are:
- `(0, 2)`: `nums[0] + nums[2]` = 6
- `(0, 3)`: `nums[0] + nums[3]` = 3
- `(0, 4)`: `nums[0] + nums[4]` = 9
- `(1, 3)`: `nums[1] + nums[3]` = 5
- `(1, 4)`: `nums[1] + nums[4]` = 11
- `(2, 4)`: `nums[2] + nums[4]` = 13  
Thus, the answer is 13.​​​​​​​

#### Example 2:

Input: nums = `[5,1,9]`, k = 1

Output: 14

Explanation:
- Since k = 1, every pair is valid.
- The maximum value is obtained from a pair `(0, 2)`​​​​​​​, which is `nums[0] + nums[2]` = 5 + 9 = 14.  
- Thus, the answer is 14.

### Constraints:

- 2 <= `n` == `nums.length` <= 10<sup>5</sup>
- 1 <= `nums[i]` <= 10<sup>9</sup>
- -1 <= `k` <= `n - 1`

## Question 3:
*Minimum Operations to Transform Binary String*
### Difficulty: Medium
### Points: 5
### Description:
You are given two binary strings s1 and s2 of the same length n.

You can perform the following operations on s1 any number of times, in any order:
- Choose an index i such that `s1[i]` == '`0`', and change it to '`1`'.
- Choose an index i such that 0 <= i < n - 1, and both `s1[i]` and `s1[i + 1]` are '`1`'. Change both characters to '`0`'.

Return the minimum number of operations required to make s1 equal to s2. If it is impossible, return -1.

### Examples:
#### Example 1:

Input: s1 = "`11`", s2 = "`00`"

Output: 1

Explanation:

Change indices 0 and 1 from '`1`' to '`0`' in one operation, so "`11`" becomes "`00`". Thus, the answer is 1.

#### Example 2:

Input: s1 = "`01`", s2 = "`10`"

Output: 3

Explanation:
- Change index 0 from '`0`' to '`1`', so "`01`" becomes "`11`".
- Change indices 0 and 1 from '`1`' to '`0`', so "`11`" becomes "`00`".
- Change index 0 from '`0`' to '`1`', so "`00`" becomes "`10`".
- Thus, the answer is 3.
#### Example 3:

Input: s1 = "`1`", s2 = "`0`"

Output: -1

Explanation:

The first operation cannot change '`1`' to '`0`', and the second operation requires two adjacent characters. Therefore, it is impossible.

### Constraints:

- 1 <= `n` == `s1.length` == `s2.length` <= 10<sup>5</sup>
- `s1` and `s2` consist only of '`0`' and '`1`'.

## Question 4:
*Count Distinct Ways to Form Target from Two Strings*
### Difficulty: Hard
### Points: 6
### Description:
You are given three strings `word1`, `word2`, and `target`.

Your task is to count the number of ways to form `target` by choosing characters from `word1` and `word2` under the following conditions:
- For each character of `target`, choose one matching character from either word1 or word2.
- The chosen indices from `word1` must be strictly increasing.
- The chosen indices from `word2` must be strictly increasing.
- At least one character must be chosen from both `word1` and `word2`.

Two ways are considered different if, for at least one position in target, the chosen character comes from a different string or a different index.

Return the number of ways. Since the answer may be very large, return it modulo `10`<sup>`9`</sup>` + 7`.

### Examples:
#### Example 1:

Input: word1 = "`abc`", word2 = "`bac`", target = "`abc`"

Output: 5

Explanation:

There are 5 ways to form `target`:
- `word1[0]` = '`a`', `word1[1]` = '`b`', `word2[2]` = '`c`'
- `word1[0]` = '`a`', `word2[0]` = '`b`', `word1[2]` = '`c`'
- `word1[0]` = '`a`', `word2[0]` = '`b`', `word2[2]` = '`c`'
- `word2[1]` = '`a`', `word1[1]` = '`b`', `word1[2]` = '`c`'
- `word2[1]` = '`a`', `word1[1]` = '`b`', `word2[2]` = '`c`'  
All ways preserve the increasing index order inside each string and choose at least one character from each string.

#### Example 2:

Input: word1 = "`cd`", word2 = "`cd`", target = "`ccd`"

Output: 4

Explanation:

There are 4 ways to form `target`:
- `word1[0]` = '`c`', `word2[0]` = '`c`', `word1[1]` = '`d`'
- `word1[0]` = '`c`', `word2[0]` = '`c`', `word2[1]` = '`d`'
- `word2[0]` = '`c`', `word1[0]` = '`c`', `word1[1] `= '`d`'
- `word2[0]` = '`c`', `word1[0]` = '`c`', `word2[1]` = '`d`'  
The first two '`c`' characters in target must come one from each string. The final '`d`' can be chosen from either string.

#### Example 3:

Input: word1 = "`xy`", word2 = "`xy`", target = "`xyxy`"

Output: 2

Explanation:

There are 2 ways to form target:
- `word1[0]` = '`x`', `word1[1]` = '`y`', `word2[0]` = '`x`', `word2[1]` = '`y`'
- `word2[0]` = '`x`', `word2[1]` = '`y`', `word1[0]` = '`x`', `word1[1]` = '`y`'  
Each "`xy`" part in target comes entirely from one string.

#### Example 4:

Input: word1 = "`ab`", word2 = "`cde`", target = "`ace`"

Output: 1

Explanation:

The only way is to choose `word1[0]` = '`a`', `word2[0]` = '`c`', and `word2[2]` = '`e`'. Thus, the answer is 1.

### Constraints:

- 1 <= `word1.length`, `word2.length`, `target.length` <= 100
- `word1`, `word2`, and `target` consist of lowercase English letters only.