## Remove Nth Node From End of List
Given the head of a linked list, remove the nth node from the end of the list and return its head.

**Follow up**: Could you do this in one pass?

 

Example 1:
![avatar](./resource/remove_ex1.jpg)
```code
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```
Example 2:
```code
Input: head = [1], n = 1
Output: []
```
Example 3:
```code
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:

+ The number of nodes in the list is sz.
+ 1 <= sz <= 30
+ 0 <= Node.val <= 100
+ 1 <= n <= sz
