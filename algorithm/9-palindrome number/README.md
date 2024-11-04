# Palindrome Number

## Tags
math

## Companies
Unknow

Give an integer x, return true if x is a palindrome, and false otherwise.

### Example 1:
```go
Input: x = 121
Output: true
Explanation: 121 readsas 121 from left to right and from right to left.
	
```

### Example 2:
```go
Input: x = -121
OutPut: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

### Example 3:
```go
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Thereform it is not a palindrome.
```

### Constraints:
- -2^31 <= x <= 2^31 -1

## 基本思路
如果让你判断回文串应该很简单，我在 数组双指针技巧汇总 中讲过。
操作数字没办法像操作字符串那么简单粗暴，但只要你在知道我在 Rabin Karp 算法详解 中讲到的从最高位开始生成数字的技巧，就能轻松解决这个问题：

```
string s = "8264"
int number = 0;
for (int i = 0; i < s.size(); i++) {
    // 将字符串转化成数字
    number = 10 * number + (s[i] - '0');
    print(number);
}

// 打印输出：
// 8
// 82
// 826
// 8264
```

你从后往前把 x 的每一位拿出来，用这个技巧生成一个数字 y，如果 y 和 x 相等，则说明 x 是回文数字。

如何从后往前拿出一个数字的每一位？和 10 求余数就行了呗。看代码吧。