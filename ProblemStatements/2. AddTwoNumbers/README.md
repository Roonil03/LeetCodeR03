# 2. Add Two Numbers
## Question Level: Medium
### <b>Description</b>
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

### <b>Examples</b>
<u>Example 1:</u><br>
Input: ``l1 = [2,4,3], l2 = [5,6,4]``<br>
Output: ``[7,0,8]``<br>
Explanation: 342 + 465 = 807.<br>

<u>Example 2:</u><br>
Input: ``l1 = [0], l2 = [0]``<br>
Output: ``[0]``<br>

<u>Example 3:</u><br>
Input: ``l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]``<br>
Output: ``[8,9,9,9,0,0,0,1]``<br>

### <b>Constraints:</b>

The number of nodes in each linked list is in the range     ``[1, 100]``.<br>
``0 <= Node.val <= 9``<br>
It is guaranteed that the list represents a number that does not have leading zeros.