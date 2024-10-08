# 2696. Minimum String Length After Removing Substrings
## Question Level: Easy
### Description:
You are given a string s consisting only of uppercase English letters.

You can apply some operations to this string where, in one operation, you can remove any occurrence of one of the substrings "AB" or "CD" from s.

Return the minimum possible length of the resulting string that you can obtain.

Note that the string concatenates after removing the substring and could produce new "AB" or "CD" substrings.

### Examples:
<b>Example 1:</b><br>
Input: s = "ABFCACDB"<br>
Output: 2<br>
Explanation: We can do the following operations:
- Remove the substring "ABFCACDB", so s = "FCACDB".<br>
- Remove the substring "FCACDB", so s = "FCAB".<br>
- Remove the substring "FCAB", so s = "FC".<br>
So the resulting length of the string is 2.<br>
It can be shown that it is the minimum length that we can obtain.<br>

<b>Example 2:</b><br>
Input: s = "ACBBD"<br>
Output: 5<br>
Explanation: We cannot do any operations on the string so the length remains the same.<br>


### Constraints:
- 1 <= s.length <= 100
- s consists only of uppercase English letters.