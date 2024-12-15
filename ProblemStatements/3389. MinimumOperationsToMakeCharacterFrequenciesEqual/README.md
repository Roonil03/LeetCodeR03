# 3389. Minimum Operations to Make Character Frequencies Equal
## Question Level: Hard
### Description:
You are given a string s.

A string `t` is called good if all characters of `t` occur the same number of times.

You can perform the following operations any number of times:
- Delete a character from s.
- Insert a character in s.
- Change a character in s to its next letter in the alphabet.<br>
<i>Note that you cannot change 'z' to 'a' using the third operation.</i>

Return the minimum number of operations required to make s good.

### Examples:
#### Example 1:
Input: s = `"acab"`<br>
Output: 1<br>
Explanation:<br>
- We can make s good by deleting one occurrence of character `'a'`.

#### Example 2:
Input: s = `"wddw"`<br>
Output: 0<br>
Explanation:
- We do not need to perform any operations since `s` is initially good.

#### Example 3:
Input: s = `"aaabc"`<br>
Output: 2<br>
Explanation:<br>
We can make s good by applying these operations:
- Change one occurrence of `'a'` to `'b'`
- Insert one occurrence of `'c'` into `s`

### Constraints:

- 3 <= `s.length` <= 2 * 10^4
- `s` contains only lowercase English letters.

### <i>Weekly Contest Question </i>