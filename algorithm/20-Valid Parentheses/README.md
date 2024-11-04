# Valid Parentheses

## Tags

### Companies
Give a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1.Open brackets must be closed by the same type of brackets.
2.Open brackets must be closed in the correct order.
3.Every close bracket hasacorresponding open bracket of the same type.

### Example 1:

Input: s = "()"
Output : true

### Example 2:

Input: s = "()[]{}"
Output: true

### Example 3:

Input: s = "(]"
Output: false

### Constraints:
- 1 <= s.length <= 10^4
- s consists of parentheses only '()[]{}'

#### 基本思路

栈是一种先进后出的数据结构，处理括号问题的时候尤其有用。
遇到左括号就入栈，遇到右括号就去栈中寻找最近的左括号，看是否匹配。