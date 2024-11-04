# Add Two Numbers

## Tags

linked-list | math

## Companies
You are given two non-empty linked lists representing two non-ngeative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.


## Example 1:
2 > 4 > 3
5 > 6 > 4
---------
7 > 0 > 8

```go
Input: l1 = [2,4,3], l2 = [5,6,4]
Output :[7,0,8]
Explanation: 342 + 465 = 807.
```

## Example 2:
```go
Input: l1 = [0], l2 = [0]
Output: [0]
```

## Example 3:
```go
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

## Constraints:
- The number of nodes in each linked list is in the range [1, 100]
- 0 <=Node.val <= 9
- It is guaranteed that the list represents a number that does not have leading zeros.

## labuladong 思路

### 基本思路
逆序存储很友好了，直接遍历链表就是从个位开始，符合我们计算加法的习惯顺序。如果是正序存储，那倒要费点脑筋了，可能需要`翻转链表` 或者使用栈来辅助。
这道题主要考察 `链表双指针技巧` 和加法运算过程中对进位的处理。注意这个 carry 变量的处理，在我们手动模拟加法过程的时候经常用倒。
代码中还用到一个链表的算法题中是很常见的[虚拟头结点] 技巧，也就是 dummy 节点。 你可以试试，如果不使用 dummy 虚拟节点， 代码会稍显复杂，而有了 dummy 节点这个占位符，可以避免处理初始的空指针情况，降低代码的复杂性。
标签： 数据结构， 链表双指针
