# Merge Two Sorted Lists

## Tags

### Companies

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list.The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list

### Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

### Example 2:
Input: list1 = [], list2 = []
Output: []

### Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

### Constraints:
- The number of nodes in both lists is in the range [0, 50]
- -100 <= Node.val <= 100
- Both list1 and list2 are sorted in non-decreasing order.


### 基本思路

这个算法的逻辑类似于 [拉拉链] , l1,l2 类似于拉链两侧的锯齿，指针 p 就好像拉链的拉索， 将两个有序链表合并。
代码中还用到一个链表的算法题中是很常见的【虚拟头结点】技巧， 也就是 dummy 节点，它相当于是个占位符，可以避免处理空指针的情况，降低代码的复杂性。
