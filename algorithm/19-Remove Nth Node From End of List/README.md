# Remove Nth Node From End of List

## Tag

## Companies

Given the head of a linked list, remove the n^th node from the end of the list end return its head.

### Example 1:

```go
Input: head = [1,2,3,4,5],n = 2
Output:[1,2,3,5]
```

### Example 2:

```go
Input: head = [1], n = 1
Output: []
```

### Example 3:

```go
Input: head = [1,2], n = 1
Output: [1]
```

### Constraints:
- The number of nodesin the list is sz.
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

Follow up: Could you do this in one pass?

### 基本思路

要删除倒数第 n 个节点，就得获得倒数第 n + 1 个节点的引用。

获取单链表的倒数第 k 个节点，就是想考察 双指针技巧 中快慢指针的运用，一般都会要求你只遍历一次链表，就算出倒数第 k 个节点。

第一步，我们先让一个指针 p1 指向链表的头节点 head，然后走 k 步：

第二步，用一个指针 p2 指向链表头节点 head：

第三步，让 p1 和 p2 同时向前走，p1 走到链表末尾的空指针时走了 n - k 步，p2 也走了 n - k 步，也就是链表的倒数第 k 个节点：

这样，只遍历了一次链表，就获得了倒数第 k 个节点 p2。

解法中在链表头部接一个虚拟节点 dummy 是为了避免删除倒数第一个元素时出现空指针异常，在头部加入 dummy 节点并不影响尾部倒数第 k 个元素是什么。