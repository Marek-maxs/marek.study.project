# 3Sum

## Tags
array two-pointers

## Companies
Given an integer array nums, return all the trplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the soluation set mush not contain duplicate trplets.

## Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

## Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

## Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
 

### Constraints:

- 3 <= nums.length <= 3000
- 105 <= nums[i] <= 105


### 基本思路

nSum 系列问题的核心思路就是排序 + 双指针。

先给数组从小到大排序，然后双指针 lo 和 hi 分别在数组开头和结尾，这样就可以控制 nums[lo] 和 nums[hi] 这两数之和的大小：

如果你想让它俩的和大一些，就让 lo++，如果你想让它俩的和小一些，就让 hi--。

基于两数之和可以得到一个万能函数 nSumTarget, 扩展出 n 数之和的解法， 具体分析详细题解。
