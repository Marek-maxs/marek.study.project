# 4Sum

## Tags

## Companies

Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

 

## Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
## Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
 

### Constraints:

- 1 <= nums.length <= 200
- 109 <= nums[i] <= 109
- 109 <= target <= 109

## 基本思路

nSum 系列问题的核心思路就是排序 + 双指针。

先给数组从小到大排序，然后双指针 lo 和 hi 分别在数组开头和结尾，这样就可以控制 nums[lo] 和 nums[hi] 这两数之和的大小：

如果你想让它俩的和大一些，就让 lo++，如果你想让它俩的和小一些，就让 hi--。

基于两数之和可以得到一个万能函数 nSumTarget，扩展出 n 数之和的解法，具体分析见详细题解。