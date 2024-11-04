# Generate Parentheses

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

 

## Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
## Example 2:

Input: n = 1
Output: ["()"]
 

Constraints:

1 <= n <= 8
通知：数据结构精品课 和 递归算法专题课 限时附赠网站会员，全新纸质书《labuladong 的算法笔记》 出版，签名版限时半价！

⭐️labuladong 题解
labuladong思路
### 基本思路
PS：这道题在《算法小抄》 的第 306 页。

本题可以改写为：

现在有 2n 个位置，每个位置可以放置字符 ( 或者 )，组成的所有括号组合中，有多少个是合法的？

这就是典型的回溯算法提醒，暴力穷举就行了。

不过为了减少不必要的穷举，我们要知道合法括号串有以下性质：

1、一个「合法」括号组合的左括号数量一定等于右括号数量，这个很好理解。

2、对于一个「合法」的括号字符串组合 p，必然对于任何  0 <= i < len(p) 都有：子串 p[0..i] 中左括号的数量都大于或等于右括号的数量。

因为从左往右算的话，肯定是左括号多嘛，到最后左右括号数量相等，说明这个括号组合是合法的。

用 left 记录还可以使用多少个左括号，用 right 记录还可以使用多少个右括号，就可以直接套用 回溯算法套路框架 了。