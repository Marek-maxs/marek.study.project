# Longest Substring Without Repeating Characters

## Tags

hash-table | two-pointers | string | sliding-window

## Companies

Given a string s, find the length of the longest substring without repeating characters.


## Example 1:
```go
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

## Example 2:
```go
Input: s = "bbbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

## Example 3:
```go
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke, with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
	
```

## Constraints:
- 0 <= s.lenth <= 5 * 10^4
- s consists of english letters, digits, symbols and spaces.

# 基本思路

这题比其他滑动窗口的题目简单，连 need 和vaild 都不需要， 而且更新窗口内数据也只需要简单的更新计数器 window 即可。
当 window[c] 值大于1时， 说明窗口中存在重复字符，不符合条件，就该移动 left 缩小窗口了。
另外， 要在收缩窗口完成后更新 res , 因为窗口收缩的 while 条件是存在重复元素， 换句话说收缩完成后一定保证窗口中没有重复。
