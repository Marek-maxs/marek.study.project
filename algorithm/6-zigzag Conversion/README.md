# Zigzag Conversion

## Tags
string

## Companies
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:(you may want to display this pattern in a fixed font for better legibility)

```shell
P   A   H   N
A P L S I I G
Y I R
```
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
```go
string convert(string s, int numRows);
```

Example 1:
```go
Input: s = "PAYPALISHIRING",  numRows = 3
Output: "PAHNAPLSIIGYIR"
```

Example 2:
```go
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINGLSIGYAHRPI"
Explanation:
	P   I   N
A   L S I G
Y A H R
P   I
```

Example 3:
```go
Input: s = "A", numRows = 1
Output: "A"
```

## Constraints:
- 1 <= s.length <= 10000
- s consists of English letters (lower-case and upper-case), "," and "."
- 1 <= numRows <= 1000

## 方法一： 利用二维矩阵模拟
设 n 为字字符串 s 的长度， r = numRows. 对于 r = 1 (只有一行) 或者 r ≥ n (只有一列) 的情况，
答案与 s 相同， 我们可以直接返回 s . 对于其余情况，考虑创建一个二维矩阵，然后在矩阵上按 Z 字形填写字符串 s , 最后逐行扫描辞职 的非空字符，组成答案。

根据题意，当我们在矩阵上填写字符时，会向下填写 r 个字符， 然后向右上继续填写 r - 2 个字符， 最后回到第一行， 因此 Z 字形变换的周期 t = r + r - 2 = 2r -2, 每个周期会占用
矩阵上的 1 + r - 2 = r -1 列.

因此我们有 ⌈ t/n⌉ 个 周期（最后一个周期视作完整周期）， 乘上每个周期的列数，得到矩阵的列数 c = [n/t].(r - 1)
创建一个 r 行 c 列的矩阵，然后遍历字符串s 并按Z字形填写。具体来说，设当前填写的位置为（x,y）, 即矩阵的 x 行 y 列。 初始 （x,y） = (0,0), 即矩阵左上角。
若当前字符下标 i 满足 i mod t < r - 1, 则向下移动， 否则向右上移动。

填写完成后， 逐行扫描矩阵中的非空字符，组成答案。

