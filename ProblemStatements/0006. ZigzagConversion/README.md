# 6. Zigzag Conversion
## Question Level: Medium
### Description:
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

``P   A   H   N``<br>
``A P L S I I G``<br>
``Y   I   R``<br>
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

### Example:
Example 1:<br>
Input: s = "PAYPALISHIRING", numRows = 3<br>
Output: "PAHNAPLSIIGYIR"<br>
<br>
Example 2:<br>
Input: s = "PAYPALISHIRING", numRows = 4<br>
Output: "PINALSIGYAHRPI"<br>
Explanation:<br>
``P     I    N``<br>
``A   L S  I G``<br>
``Y A   H R``<br>
``P     I``<br>

Example 3:<br>
Input: s = "A", numRows = 1<br>
Output: "A"<br>


### Constraints:
- 1 <= s.length <= 1000
- s consists of English letters (lower-case and upper-case), ``','`` and ``'.'``.
- 1 <= numRows <= 1000