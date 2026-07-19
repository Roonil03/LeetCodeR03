# Biweekly Contest 187
## Ranking:
## Question 1:
*Rearrange String to Avoid Character Pair*
### Difficulty: Easy
### Points: 3
### Description:
You are given a string s and two distinct lowercase English letters x and y.

Rearrange the characters of `s` to construct a new string `t` such that:
- `t` is a permutation of `s`.
- Every occurrence of y appears before every occurrence of `x` in `t`.

Return any valid string `t`.

### Examples:
#### Example 1:

Input: s = "`aabc`", x = "`a`", y = "`c`"

Output: "`cbaa`"

Explanation:

The string "`cbaa`" is a permutation of "`aabc`", and every occurrence of '`c`' appears before every occurrence of '`a`'.

#### Example 2:

Input: s = "`dcab`", x = "`d`", y = "`b`"

Output: "`cabd`"

Explanation:

The string "`cabd`" is a permutation of "`dcab`", and every occurrence of '`b`' appears before every occurrence of '`d`'.

#### Example 3:

Input: s = "`axe`", x = "`o`", y = "`x`"

Output: "`axe`"

Explanation:

The string "`axe`" is already valid. Since '`o`' does not occur in the string, the required condition is automatically satisfied.

### Constraints:

- 1 <= `s.length` <= 100
- `s` consists of lowercase English letters.
- `x` and `y` are lowercase English letters.
- `x` != `y`

## Question 2:
*Maximum Value of an Alternating Sequence*
### Difficulty: Medium
### Points: 4
### Description:
You are given three integers n, s, and m.

A sequence seq of integers of length n is considered valid if:
- `seq[0]` = s.
- The sequence is alternating, meaning that either:
    - `seq[0]` > `seq[1]` < `seq[2]` > ..., or
    - `seq[0]` < `seq[1]` > `seq[2]` < ....
- For every adjacent pair, `|seq[i] - seq[i - 1]|` <= `m`.

A sequence of length 1 is considered alternating.

Return the maximum possible element that can appear in any valid sequence.

### Examples:
#### Example 1:

Input: n = 4, s = 3, m = 5

Output: 12

Explanation:

One valid sequence is `[3, 8, 7, 12]`.  
The maximum element in the sequence is 12.
#### Example 2:

Input: n = 2, s = 4, m = 3

Output: 7

Explanation:

One valid sequence is `[4, 7]`.  
The maximum element in the sequence is 7.

### Constraints:

- 1 <= `n`, `s` <= 10<sup>9</sup>
- 1 <= `m` <= 10<sup>5</sup>

## Question 3: 
*Minimum Adjacent Swaps to Partition Array*
### Difficulty: Medium
### Points: 5
### Description:
You are given an integer array nums and two integers a and b such that a < b.

An array is called good if it can be split into three contiguous parts, in this order, such that:
- Every element in the first part is less than a.
- Every element in the second part is in the range `[a, b]` inclusive.
- Every element in the third part is greater than b.

Any of the three parts may be empty.

In one adjacent swap, you may swap two neighboring elements of `nums`.

Return the minimum number of adjacent swaps required to make nums good. Since the answer may be very large, return it modulo 10<sup>9</sup> + 7.

### Examples:
#### Example 1:

Input: nums = `[1,3,2,4,5,6]`, a = 3, b = 4

Output: 1

Explanation:

- Swap `nums[1]` and `nums[2]`. The array becomes `[1, 2, 3, 4, 5, 6]`.
- This array is good because it can be split into `[1, 2]`, `[3, 4]`, and `[5, 6]`.
#### Example 2:

Input: nums = `[9,7,5,3]`, a = 4, b = 8

Output: 5

Explanation:

One sequence of optimal swaps is as follows:

- Swap `nums[2]` and `nums[3]`. The array becomes `[9, 7, 3, 5]`.
- Swap `nums[1]` and `nums[2]`. The array becomes `[9, 3, 7, 5]`.
- Swap `nums[0]` and `nums[1]`. The array becomes `[3, 9, 7, 5]`.
- Swap `nums[1]` and `nums[2]`. The array becomes `[3, 7, 9, 5]`.
- Swap `nums[2]` and `nums[3]`. The array becomes `[3, 7, 5, 9]`.  
This array is good because it can be split into `[3]`, `[7, 5]`, and `[9]`.
#### Example 3:

Input: nums = `[3,7,5,9]`, a = 4, b = 8

Output: 0

Explanation:

The array is already good. No swaps are needed.

### Constraints:

- 1 <= `nums.length` <= 10<sup>5</sup>
- ​​​​​​​1 <= `nums[i]` <= 10<sup>9</sup>
- 1 <= `a` < `b` <= 10<sup>9</sup>

## Question 4: 
*Minimum Cost to Convert String III*
### Difficulty: Hard
### Points: 6
### Description:
You are given two strings, source and target.

You are also given a 2D string array rules, where `rules[i]` = `[pattern`<sub>`i`</sub>`, replacement`<sub>`i`</sub>`]`, and an integer array costs, where `costs[i]` is the base cost of applying `rules[i]`. Both arrays have the same length. Additionally, pattern<sub>i</sub> and replacement<sub>i</sub> have the same length.

You may apply any rule any number of times. Each rule application works as follows:
- Choose an index l such that the range of positions from l to l + pattern<sub>i</sub>>.length - 1 exists in the current string and none of these positions has been used in a previous rule application.  
- For each index j, the character `pattern`<sub>`i`</sub>`[j]` must either be equal to the current character at position l + j, or be '`*`'.
- Replace the characters in this range with `replacement`<sub>`i`</sub>. The replacement is used exactly as given and does not contain wildcards.
- The cost of this rule application is `costs[i]` plus the number of '`*`' characters in patterni.
- Once a character position has been used in a rule application, it cannot be used in any later rule application.

Since every pattern<sub>i</sub> and replacement<sub>i</sub> have the same length, character positions are preserved after every rule application.

Return the minimum total cost required to transform source into target. If it is impossible, return -1.

### Examples:
#### Example 1:

Input: source = "`hello`", target = "`world`", rules = `[["he","wo"],["llo","rld"]]`, costs = `[3,4]`

Output: 7

Explanation:

- Apply `rules[0]` to replace "`he`" with "`wo`" at cost 3, so the string becomes "`wollo`".
- Apply `rules[1]` to replace "`llo`" with "`rld`" at cost 4, so the string becomes "`world`".  
The total cost is 3 + 4 = 7.
#### Example 2:

Input: source = "`cat`", target = "`dog`", rules = `[["c*t","dog"]]`, costs = `[2]`

Output: 3

Explanation:

- Apply `rules[0]` to replace "`cat`" with "`dog`". The wildcard '`*`' matches '`a`', adding 1 to the base cost 2.
- The total cost is 2 + 1 = 3.
#### Example 3:

Input: source = "`test`", target = "`next`", rules = `[["*e*t","next"]]`, costs = `[4]`

Output: 6

Explanation:

- Apply `rules[0]` to replace "`test`" with "`next`". The first wildcard matches '`t`' and the second wildcard matches '`s`', adding 2 to the base cost 4.
- The total cost is 4 + 2 = 6.
#### Example 4:

Input: source = "`ab`", target = "`bc`", rules = `[["a*","bd"]]`, costs = `[9]`

Output: -1

Explanation:

No sequence of rule applications can transform source into target, so the answer is -1.

### Constraints:

- 1 <= `source.length`, `target.length` <= 5000
- `source` and `target` consist of lowercase English letters.
- 1 <= `rules.length` == `costs.length` <= 200
- `rules[i]` = `[pattern`<sub>`i`</sub>`, replacement`<sub>`i`</sub>`]`
- 1 <= `pattern`<sub>`i`</sub>`.length` == `replacement`<sub>`i`</sub>`.length` <= 20
- `pattern`<sub>`i`</sub> contains at least one lowercase English letter and at most 5 '`*`' characters.
- `replacement`<sub>`i`</sub> contains only lowercase English letters.
- 1 <= `costs[i]` <= 1000

