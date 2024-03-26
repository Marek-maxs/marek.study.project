# Container With Most Water

## Tags
array two-pointers

## Companies
You are given an integer array height of length n.There are n vertical lines drawn such that the tow endpoints of the i^th line are (i, 0) and (i, height[i]). 
Find two lines that together with the x-asix form form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

 
## Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

## Example 2:
Input: height = [1,1]
Output: 1

### Constraints:
- n == height.length
- 2 <= n <= 10^5
- 0 <= height[i] <= 10^4

## 基本思路
区别在于：接雨水问题给出的类似一幅直方图，每个横坐标都有宽度，而本题给出的每个坐标是一条竖线，没有宽度。

接雨水问题更难一些，每个坐标对应的矩形通过左右的最大高度的最小值推算自己能装多少水：

本题可完全套用接雨水问题的思路，相对还更简单：

用 left 和 right 两个指针从两端向中心收缩，一边收缩一边计算 [left, right] 之间的矩形面积，取最大的面积值即是答案。

不过肯定有读者会问，下面这段 if 语句为什么要移动较低的一边：

```
// 双指针技巧，移动较低的一边
if (height[left] < height[right]) {
    left++;
} else {
    right--;
}
```

其实也好理解，因为矩形的高度是由 min(height[left], height[right]) 即较低的一边决定的：

你如果移动较低的那一边，那条边可能会变高，使得矩形的高度变大，进而就「有可能」使得矩形的面积变大；相反，如果你去移动较高的那一边，矩形的高度是无论如何都不会变大的，所以不可能使矩形的面积变得更大。

详细题解：如何高效解决接雨水问题