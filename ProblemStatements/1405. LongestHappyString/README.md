# 1405. Longest Happy String
## Question Level: Medium
### Description:
A string s is called happy if it satisfies the following conditions:

- s only contains the letters ``'a'``, ``'b'``, and ``'c'``.
- s does not contain any of ``"aaa"``, ``"bbb"``, or ``"ccc"`` as a substring.
- s contains at most a occurrences of the letter ``'a'``.
- s contains at most b occurrences of the letter ``'b'``.
- s contains at most c occurrences of the letter ``'c'``.

Given three integers a, b, and c, return the longest possible happy string. If there are multiple longest happy strings, return any of them. If there is no such string, return the empty string "".

A substring is a contiguous sequence of characters within a string.

### Examples:
<b>Example 1:</b><br>
Input: ``a = 1``, ``b = 1``, ``c = 7``<br>
Output: ``"ccaccbcc"``<br>
Explanation: ``"ccbccacc"`` would also be a correct answer.<br>

<b>Example 2:</b><br>
Input: ``a = 7``, ``b = 1``, ``c = 0``<br>
Output: ``"aabaa"``<br>
Explanation: It is the only correct answer in this case.<br>

### Constraints:

- ``0 <= a, b, c <= 100``
- ``a + b + c > 0``

### <i> Concepts Used:
- String
- Greedy
- Heap (Priority Queue)</i>