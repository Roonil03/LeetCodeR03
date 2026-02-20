# Biweekly Contest 176
## Ranking: 395 / 41954
## Question 1:
*Weighted Word Mapping*
### Difficulty: Easy
### Points: 3
### Description:
You are given an array of strings words, where each string represents a word containing lowercase English letters.

You are also given an integer array `weights` of length 26, where `weights[i]` represents the weight of the i<sup>th</sup> lowercase English letter.

The weight of a word is defined as the sum of the weights of its characters.

For each word, take its weight modulo 26 and map the result to a lowercase English letter using reverse alphabetical order `(0 -> 'z', 1 -> 'y', ..., 25 -> 'a')`.

Return a string formed by concatenating the mapped characters for all words in order.

### Examples:
#### Example 1:

Input: words = `["abcd","def","xyz"]`, weights = `[5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]`

Output: "`rij`"

Explanation:
- The weight of "`abcd`" is 5 + 3 + 12 + 14 = 34. The `result modulo 26 is 34 % 26 = 8`, which maps to '`r`'.
- The weight of "`def`" is 14 + 1 + 2 = 17. The `result modulo 26 is 17 % 26 = 17`, which maps to '`i`'.
- The weight of "`xyz`" is 7 + 7 + 2 = 16. The `result modulo 26 is 16 % 26 = 16`, which maps to '`j`'.  
Thus, the string formed by concatenating the mapped characters is "`rij`".

#### Example 2:

Input: words = `["a","b","c"]`, weights = `[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]`

Output: "`yyy`"

Explanation:
- Each word has weight 1. The `result modulo 26 is 1 % 26 = 1`, which maps to '`y`'.  
Thus, the string formed by concatenating the mapped characters is "`yyy`".

#### Example 3:

Input: words = `["abcd"]`, weights = `[7,5,3,4,3,5,4,9,4,2,2,7,10,2,5,10,6,1,2,2,4,1,3,4,4,5]`

Output: "`g`"

Explanation:​​​​​​​

- The weight of "`abcd`" is 7 + 5 + 3 + 4 = 19. The `result modulo 26 is 19 % 26 = 19`, which maps to '`g`'.  
Thus, the string formed by concatenating the mapped characters is "`g`".

### Constraints:

- 1 <= `words.length` <= 100
- 1 <= `words[i].length` <= 10
- `weights.length` == 26
- 1 <= `weights[i]` <= 100
- `words[i]` consists of lowercase English letters.

## Question 2:
*Number of Prefix Connected Groups*
### Difficulty: Medium
### Points: 4
### Description:
You are given an array of strings words and an integer k.

Two words a and b at distinct indices are prefix-connected if `a[0..k-1]` == `b[0..k-1]`.

A connected group is a set of words such that each pair of words is prefix-connected.

Return the number of connected groups that contain at least two words, formed from the given words.

**Note:**
- Words with length less than k cannot join any group and are ignored.
- Duplicate strings are treated as separate words.
- A prefix of a string is a substring that starts from the beginning of the string and extends to any point within it.

### Examples:
#### Example 1:

Input: words = `["apple","apply","banana","bandit"]`, k = 2

Output: 2

Explanation:

Words sharing the same first k = 2 letters are grouped together:
- `words[0]` = "`apple`" and `words[1]` = "`apply`" share prefix "`ap`".
- `words[2]` = "`banana`" and `words[3]` = "`bandit`" share prefix "`ba`".  
Thus, there are 2 connected groups, each containing at least two words.

#### Example 2:

Input: words = `["car","cat","cartoon"]`, k = 3

Output: 1

Explanation:

Words are evaluated for a prefix of length k = 3:
- `words[0]` = "`car`" and `words[2]` = "`cartoon`" share prefix "`car`".
- `words[1]` = "`cat`" does not share a 3-length prefix with any other word.  
Thus, there is 1 connected group.

#### Example 3:

Input: words = `["bat","dog","dog","doggy","bat"]`, k = 3

Output: 2

Explanation:

Words are evaluated for a prefix of length k = 3:
- `words[0]` = "`bat`" and `words[4]` = "`bat`" form a group.
- `words[1]` = "`dog`", `words[2]` = "`dog`" and `words[3]` = "`doggy`" share prefix "`dog`".  
Thus, there are 2 connected groups, each containing at least two words.

### Constraints:

- 1 <= `words.length` <= 5000
- 1 <= `words[i].length` <= 100
- 1 <= `k` <= 100
- All strings in words consist of lowercase English letters.

## Question 3:
*House Robber V*
### Difficulty: Medium
### Points: 5
### Description:
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed and is protected by a security system with a color code.

You are given two integer arrays nums and colors, both of length n, where `nums[i]` is the amount of money in the ith house and `colors[i]` is the color code of that house.

You cannot rob two adjacent houses if they share the same color code.

Return the maximum amount of money you can rob.

### Examples:
#### Example 1:

Input: nums = `[1,4,3,5]`, colors = `[1,1,2,2]`

Output: 9

Explanation:

Choose houses i = 1 with `nums[1]` = 4 and i = 3 with `nums[3]` = 5 because they are non-adjacent.  
Thus, the total amount robbed is 4 + 5 = 9.
#### Example 2:

Input: nums = `[3,1,2,4]`, colors = `[2,3,2,2]`

Output: 8

Explanation:

Choose houses i = 0 with `nums[0]` = 3, i = 1 with `nums[1]` = 1, and i = 3 with `nums[3]` = 4.  
This selection is valid because houses i = 0 and i = 1 have different colors, and house i = 3 is non-adjacent to i = 1.  
Thus, the total amount robbed is 3 + 1 + 4 = 8.  
#### Example 3:

Input: nums = `[10,1,3,9]`, colors = `[1,1,1,2]`

Output: 22

Explanation:

Choose houses i = 0 with `nums[0]` = 10, i = 2 with `nums[2]` = 3, and i = 3 with `nums[3]` = 9.  
This selection is valid because houses i = 0 and i = 2 are non-adjacent, and houses i = 2 and i = 3 have different colors.  
Thus, the total amount robbed is 10 + 3 + 9 = 22.  

### Constraints:

- 1 <= `n` == `nums.length` == `colors.length` <= 10<sup>5</sup>
- 1 <= `nums[i]`, `colors[i]` <= 10<sup>5</sup>

## Question 4
*Palindromic Path Queries in a Tree*
### Difficulty: Hard
### Points: 7
### Description:
You are given an undirected tree with n nodes labeled 0 to `n - 1`. This is represented by a 2D array edges of length `n - 1`, where `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`]` indicates an undirected edge between nodes `u`<sub>`i`</sub> and `v`<sub>`i`</sub>.

You are also given a string s of length n consisting of lowercase English letters, where `s[i]` represents the character assigned to node i.

You are also given a string array queries, where each `queries[i]` is either:
- "`update u`<sub>`i`</sub>` c`": Change the character at node `u`<sub>`i`</sub> to `c`. Formally, update `s[u`<sub>`i`</sub>`] = c`.
- "`query u`<sub>`i`</sub>` v`<sub>`i`</sub>": Determine whether the string formed by the characters on the unique path from `u`<sub>`i`</sub> to `v`<sub>`i`</sub> (inclusive) can be rearranged into a palindrome.

Return a boolean array `answer`, where `answer[j]` is true if the j<sup>th</sup> query of type "`query u`<sub>`i`</sub>` v`<sub>`i`</sub>"​​​​​​​ can be rearranged into a palindrome, and false otherwise.

A palindrome is a string that reads the same forward and backward.
 
### Examples:
#### Example 1:

Input: n = 3, edges = `[[0,1],[1,2]]`, s = "`aac`", queries = `["query 0 2","update 1 b","query 0 2"]`

Output: `[true,false]`

Explanation:

- "`query 0 2`": Path `0 → 1 → 2` gives "`aac`", which can be rearranged to form "`aca`", a palindrome. Thus, `answer[0] = true`.
- "`update 1 b`": Update node 1 to '`b`', now s = "`abc`".
- "`query 0 2`": Path characters are "`abc`", which cannot be rearranged to form a palindrome. Thus, `answer[1] = false`.  
Thus, answer = `[true, false]`.

#### Example 2:

Input: n = 4, edges = `[[0,1],[0,2],[0,3]]`, s = "`abca`", queries = `["query 1 2","update 0 b","query 2 3","update 3 a","query 1 3"]`

Output: `[false,false,true]`

Explanation:

- "`query 1 2`": Path `1 → 0 → 2` gives "`bac`", which cannot be rearranged to form a palindrome. Thus, `answer[0] = false`.
- "`update 0 b`": Update node 0 to '`b`', now s = "`bbca`".
- "`query 2 3`": Path `2 → 0 → 3` gives "`cba`", which cannot be rearranged to form a palindrome. Thus, `answer[1] = false`.
- "`update 3 a`": Update node 3 to '`a`', s = "`bbca`".
- "`query 1 3`": Path `1 → 0 → 3` gives "`bba`", which can be rearranged to form "`bab`", a palindrome. Thus, `answer[2] = true`.  
Thus, answer = `[false, false, true]`.

### Constraints:

- 1 <= `n` == `s.length` <= `5 * 10`<sup>`4`</sup>
- `edges.length` == `n - 1`
- `edges[i]` = `[u`<sub>`i`</sub>`, v`<sub>`i`</sub>`]`
- 0 <= `u`<sub>`i`</sub>, `v`<sub>`i`</sub> <= `n - 1`
- `s` consists of lowercase English letters.
- The input is generated such that `edges` represents a valid tree.
- 1 <= `queries.length` <= `5 * 10`<sup>`4`</sup>​​​​​​​
    - `queries[i]` = "`update u`<sub>`i`</sub>` c`" or
    - `queries[i]` = "`query u`<sub>`i`</sub>` v`<sub>`i`</sub>"
    - 0 <= `u`<sub>`i`</sub>, `v`<sub>`i`</sub> <= `n - 1`
    - `c` is a lowercase English letter.