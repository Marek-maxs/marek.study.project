# Longest Palindromic Substring

## Tags 
string | dynamic-programing

## Companies

Give a string s, return the longest palindromic substring in s.

## Example 1:
```go
Input : s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
	
```

## Example 2:
```go
Input: 2 = "cbbd"
Output: "bb"
```

## Constraints:

- 1 <= s.length <= 10000
- s consist of only digits and English letters.

## 基本思路

ps: 这道题在《算法小抄》 的第 373 页。
寻找回文串的问题心思想是： 从中间开始向两边扩散来判断回文串，对于最长回文子串，就是这个意思：

```go
for 0 <= i < len(s):
	找到以s[i]为中心的回文串
    更新答案

```
找回文串的关键技巧是传入两个指针 l 和 r 向两边扩散， 因为这样实现可以同时处理回文串长度为奇数和偶数的情况。

```go
for 0 <= i < len(s):
	# 找到以 s[i] 为中心的回文串
    palindrome(s, i, i)
    # 找到以 s[i] 和 s[i+1] 为中心的回文串
    palindrome(s, i, i+1)
    更新答案