# 93. Restore IP Addresses
## Question Level: Medium
### Description:
A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

For example, "`0.1.2.201`" and "`192.168.1.1`" are valid IP addresses, but "`0.011.255.245`", "`192.168.1.312`" and "`192.168@1.1`" are invalid IP addresses.<br>
Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.

### Examples:
#### Example 1:

Input: s = "`25525511135`"<br>
Output: `["255.255.11.135","255.255.111.35"]`<br>
#### Example 2:

Input: s = "`0000`"<br>
Output: `["0.0.0.0"]`<br>
#### Example 3:

Input: s = "`101023`"<br>
Output: `["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]`<br>


### Constraints:

- 1 <= `s.length` <= 20
- s consists of digits only.

### <i>Concepts Used:
- String
- Backtracking </i>