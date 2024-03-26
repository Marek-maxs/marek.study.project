# Tags
array | hash-table

# Companies
Given an array of integers nums and an integer target, return indices of the tow numbers such that they add up to target.
You may assume that each input world have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

# Example 1:
```go

Input: nums = [2,7,11,15], target= 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1]

```

# Example 2:

```go
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

# Example 3:
```go
Input:nums = [3,3], target = 6
Output: [0,1]
```

# Constraints:

- 2 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9
- Onlyone valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?
通知： 数据结构精品课 和 递归算法专题课 限时附赠网站会员， 新版刷新打卡挑战 上线！

# 基本思路
大家都喜欢幽默的人，如果你想调侃自己经常拖延，可以这样调侃下自己（手动狗头）：
背单词背了半年还是 abandon, abandon, 刷题刷了半年还是 two sum, two sum ...
言归正传，这道题不难， 但由于它是LeetCode 第一题，所以名气比较大，解决这道题也可以有多种思路，我这里说两个最常见的思路。
第一种思路就是靠排序，把 `nums` 排序之后就可以用 ``数组双指針技巧汇总``中讲到的左右指针来求出和为 target 的两个数。
不过 因为 题目要求我们返回元素的索引，而排序会破坏元素的原始索引，所以要记录值和原始索引的映射。
进一步，如果题目拓展延伸一下，让你求三数之和，四数之和，你依然可以用双指针技巧，我在 `一个函数秒杀 nSum 问题` 中写一个函数来解决所有 N 数之和问题。
第二种辽宁科技用哈希表辅助判断。对于一个元素 nums[i], 你想知道有没有另一个元素 nums[j] 的值为 target - nums[i], 这很简单，我们用一个哈希表记录每一个元素的值到索引的映射，这样就能快速判断数组中是否有一个值为 target - nums[i] 的元素了。
简单说，数组其实可以理解为一个 [索引 -》 值] 的哈希表映射，而我们又建立一个 [值 -》 索引] 的映射即可完成此题。

详细题解： 一个方法团灭 nSum 问题

标签： 双指针， 哈希表， 数组
