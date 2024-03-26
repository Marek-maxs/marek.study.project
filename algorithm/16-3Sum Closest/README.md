# 3Sum Closest

## Tags


## Companies
bloomberg

Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.

## Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
## Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
 

### Constraints:

- 3 <= nums.length <= 500
- 1000 <= nums[i] <= 1000
- 104 <= target <= 104

### 基本思路

只要你做过 259. 较小的三数之和，这道题稍微改一下就应该能搞定了。

一样是先排序，然后固定第一个数，再去 nums[start..] 中寻找最接近 target - delta 的两数之和。

我写的解法其实可以合并成一个函数，但考虑到和前文 一个函数秒杀 nSum 问题 中代码的一致性，我还是把解法分成了两个函数来写。