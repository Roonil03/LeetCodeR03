# 115. Distinct Subsequences
## Question Level: Hard
### Description:
Given two strings s and t, return the number of distinct subsequences of s which equals t.

The test cases are generated so that the answer fits on a 32-bit signed integer.

### Examples:
#### xample 1:

Input: s = "`rabbbit`", t = "`rabbit`"  
Output: 3  
Explanation:  
As shown below, there are 3 ways you can generate "`rabbit`" from s.  
- <b>rabb</b>b<b>it</b>
- <b>ra</b>b<b>bbit</b>
- <b>rab</b>b<b>bit</b>
#### Example 2:

Input: s = "`babgbag`", t = "`bag`"  
Output: 5  
Explanation:  
As shown below, there are 5 ways you can generate "`bag`" from s.
- <b>ba</b>b<b>g</b>bag
- <b>ba</b>bgba<b>g</b>
- <b>b</b>abgb<b>ag</b>
- ba<b>b</b>gb<b>ag</b>
- babg<b>bag</b>

### Constraints:

- 1 <= `s.length`, `t.length` <= 1000
- s and t consist of English letters.

### <i>Concepts Used:
- String
- Dynamic Programming </i>