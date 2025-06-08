# 1657. Determine if Two Strings Are Close
## Question Level: Medium
### Description:
Two strings are considered close if you can attain one from the other using the following operations:
- Operation 1: Swap any two existing characters.
    - For example, a<b>b</b>cd<b>e</b> -> a<b>e</b>cd<b>b</b>
- Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
    - For example, <b>aa</b>c<b>abb</b> -> <b>bb</b>c<b>baa</b> (all `a`'s turn into `b`'s, and all `b`'s turn into `a`'s)

You can use the operations on either string as many times as necessary.

Given two strings, `word1` and `word2`, return true if `word1` and `word2` are close, and false otherwise.

### Examples:
#### Example 1:

Input: word1 = "`abc`", word2 = "`bca`"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "a<b>bc</b>" -> "a<b>cb</b>"
Apply Operation 1: "<b>a</b>c<b>b</b>" -> "<b>b</b>c<b>a</b>"
#### Example 2:

Input: word1 = "`a`", word2 = "`aa`"  
Output: false  
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
#### Example 3:

Input: word1 = "`cabbba`", word2 = "`abbccc`"  
Output: true  
Explanation: You can attain word2 from word1 in 3 operations.  
Apply Operation 1: "ca<b>b</b>bb<b>a</b>" -> "ca<b>a</b>bb<b>b</b>"  
Apply Operation 2: "<b>c</b>aa<b>bbb</b>" -> "<b>b</b>aa<b>ccc</b>"  
Apply Operation 2: "<b>baa</b>ccc" -> "<b>abb</b>ccc"  

### Constraints:

- 1 <= `word1.length`, `word2.length` <= 10<sup>5</sup>
- `word1` and `word2` contain only lowercase English letters.

### <i>Concepts Used:
- Hash Table
- String
- Sorting
- Counting</i>