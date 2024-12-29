# 1639. Number of Ways to Form a Target String Given a Dictionary
## Question Level: Hard
### Description:
You are given a list of strings of the same length words and a string target.

Your task is to form target using the given words under the following rules:
- target should be formed from left to right.
- To form the ith character (0-indexed) of target, you can choose the kth character of the jth string in words if `target[i]` = `words[j][k]`.
- Once you use the kth character of the jth string of words, you can no longer use the xth character of any string in words where x <= k. In other words, all characters to the left of or at index k become unusuable for every string.
- Repeat the process until you form the string target.

Notice that you can use multiple characters from the same string in words provided the conditions above are met.

Return the number of ways to form target from words. Since the answer may be too large, return it modulo `10^9 + 7`.

### Examples:
#### Example 1:

Input: words = `["acca","bbbb","caca"]`, target = "`aba`"<br>
Output: 6<br>
Explanation: There are 6 ways to form target.
- "`aba`" -> index 0 ("<b>a</b>cca"), index 1 ("b<b>b</b>bb"), index 3 ("cac<b>a</b>")
- "`aba`" -> index 0 ("<b>a</b>cca"), index 2 ("bb<b>b</b>b"), index 3 ("cac<b>a</b>")
- "`aba`" -> index 0 ("<b>a</b>cca"), index 1 ("b<b>b</b>bb"), index 3 ("acc<b>a</b>")
- "`aba`" -> index 0 ("<b>a</b>cca"), index 2 ("bb<b>b</b>b"), index 3 ("acc<b>a</b>")
- "`aba`" -> index 1 ("c<b>a</b>ca"), index 2 ("bb<b>b</b>b"), index 3 ("acc<b>a</b>")
- "`aba`" -> index 1 ("c<b>a</b>ca"), index 2 ("bb<b>b</b>b"), index 3 ("cac<b>a</b>")
#### Example 2:

Input: words = `["abba","baab"]`, target = "`bab`"<br>
Output: 4<br>
Explanation: There are 4 ways to form target.
- "`bab`" -> index 0 ("<b>b</b>aab"), index 1 ("b<b>a</b>ab"), index 2 ("ab<b>b</b>a")
- "`bab`" -> index 0 ("<b>b</b>aab"), index 1 ("b<b>a</b>ab"), index 3 ("baa<b>b</b>")
- "`bab`" -> index 0 ("<b>b</b>aab"), index 2 ("ba<b>a</b>b"), index 3 ("baa<b>b</b>")
- "`bab`" -> index 1 ("a<b>b</b>ba"), index 2 ("ba<b>a</b>b"), index 3 ("baa<b>b</b>")

### Constraints:

- 1 <= `words.length` <= 1000
- 1 <= `words[i].length` <= 1000
- All strings in words have the same length.
- 1 <= `target.length` <= 1000
- `words[i]` and `target` contain only lowercase English letters.

### <i>Concepts Used:
- Array
- String
- Dynamic Programming
</i>