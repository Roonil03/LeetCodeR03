# 1106. Parsing A Boolean Expression
## Question Level: Hard
### Description:
A boolean expression is an expression that evaluates to either true or false. It can be in one of the following shapes:

- ``'t'`` that evaluates to true.
- ``'f'`` that evaluates to false.
- ``'!(subExpr)'`` that evaluates to the logical NOT of the inner expression subExpr.
- ``'&(subExpr1, subExpr2, ..., subExprn)'`` that evaluates to the logical AND of the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.
- ``'|(subExpr1, subExpr2, ..., subExprn)'`` that evaluates to the logical OR of the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.

Given a string expression that represents a boolean expression, return the evaluation of that expression.

It is guaranteed that the given expression is valid and follows the given rules.


### Examples:
<b>Example 1:</b><br>
Input: expression = ``"&(|(f))"``<br>
Output: false<br>
Explanation: 
- First, evaluate ``|(f)`` --> ``f``. The expression is now ``"&(f)"``.
- Then, evaluate ``&(f)`` --> ``f``. The expression is now ``"f"``.
- Finally, return false.

<b>Example 2:</b><br>
Input: expression = ``"|(f,f,f,t)"``<br>
Output: true<br>
Explanation: 
- The evaluation of ``(false OR false OR false OR true)`` is true.

<b>Example 3:</b><br>
Input: expression = ``"!(&(f,t))"``<br>
Output: true<br>
Explanation: 
- First, evaluate ``&(f,t)`` --> ``(false AND true)`` --> ``false`` --> ``f``. The expression is now ``"!(f)"``.
- Then, evaluate ``!(f)`` --> ``NOT false`` --> ``true``. We return true.

### Constraints:
- 1 <= expression.length <= 2 * 10^4
- expression[i] is one following characters: ``'('``, ``')'``, ``'&'``, ``'|'``, ``'!'``, ``'t'``, ``'f'`, and ``','``.

### <i> Concepts Used:
- String
- Stack
- Recursion </i>