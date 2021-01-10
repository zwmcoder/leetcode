## Longest Palindromic Substring
Given a string s, return the longest palindromic substring in s.

 

Example 1:
```code
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example 2:
```code
Input: s = "cbbd"
Output: "bb"
```
Example 3:
```code
Input: s = "a"
Output: "a"
```
Example 4:
```code
Input: s = "ac"
Output: "a"
 ```

Constraints:
```code
1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case),
```


Optimize the original code at main.go:
```code 
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    max := 0
    start := 0
    for i:= 0; i<len(s); i++ {
        lo, so := extendPalindrome(s,i,i)
        if max < lo {
            max = lo
            start = so
        }
        le, se := extendPalindrome(s,i,i+1)
        if max < le {
            max = le
            start = se
        }
    }
    return s[start:start+max]
}
func extendPalindrome(s string, i, j int) (length, start int) {
    for ; i>=0 && j<len(s) && s[i] == s[j]; {
        i -= 1
        j += 1
    }
    return j-i-1, i+1
}
```