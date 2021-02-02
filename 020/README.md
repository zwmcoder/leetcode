## Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

Example 1:
```code
Input: s = "()"
Output: true
```
Example 2:
```code
Input: s = "()[]{}"
Output: true
```
Example 3:
```code
Input: s = "(]"
Output: false
```
Example 4:
```code
Input: s = "([)]"
Output: false
```
Example 5:
```code
Input: s = "{[]}"
Output: true
```

Constraints:
```code
1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
```