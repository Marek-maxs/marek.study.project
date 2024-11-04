# Regular Expression Matching

## Tags
string dynamic-programming backtracking

## Companies
Give an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where：

- '.' Matches any single character.
- '*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

## Example 1:
```go
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

## Example 2:
```go
Input: s = "aa", p = "a*"
Output : true
Explanation: '*' means zero or more of the preceding, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

```

## Example 3:
```go
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
	
```

### Constraints:
- 1 <= s.length <= 20
- 1 <= p.length <= 20
- s contains only lowercase English letters.
- p contains only lowercase English letters, '.', and '*'
- It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.


## 基本思路

s 和 p 相互匹配的过程大致是，两个指针 i, j 分别在 s 和 p 上移动， 如果最后两个指针都能移动到字符串的末尾， 那么就匹配成功，反之则匹配失败。
正则表达算法问题只需要把住一个基本点： 看 s[i] 和 p[j] 两个字符是否匹配，一切逻辑围绕匹配/不匹配两种情况展开即可。
动态规则算法的核心就是【状态】和【选择】，【状态】无非就是 i 和 j 两个指针的位置， 【选择】就是模式串的p[j] 选择匹配几个字符。
dp 函数的定义如下：
**若 dp(s, i, p, j) = true, 则表示s[i..] 可以匹配p[j..]; 若 dp(s,i,p,j) = false, 则表示 s[i..] 无法匹配**. 
 
