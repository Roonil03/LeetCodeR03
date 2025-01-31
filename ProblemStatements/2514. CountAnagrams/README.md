# 2514. Count Anagrams
## Question Level: Hard
### Description:
You are given a string s containing one or more words. Every consecutive pair of words is separated by a single space `' '`.

A string t is an anagram of string s if the ith word of t is a permutation of the ith word of s.

For example, "`acb dfe`" is an anagram of "`abc def`", but "`def cab`" and "`adc bef`" are not.<br>
Return the number of distinct anagrams of s. Since the answer may be very large, return it modulo 10^9 + 7.

### Examples:
#### Example 1:

Input: s = "`too hot`"<br>
Output: 18<br>
Explanation: Some of the anagrams of the given string are "`too hot`", "`oot hot`", "`oto toh`", "`too toh`", and "`too oht`".<br>
#### Example 2:

Input: s = "`aa`"<br>
Output: 1<br>
Explanation: There is only one anagram possible for the given string.<br>

### Constraints:

- 1 <= `s.length` <= 10^5
- s consists of lowercase English letters and spaces `' '`.
- There is single space between consecutive words.

### <i>Concepts Used:
- Hash Table
- Math
- String
- Combinatorics
- Counting </i>