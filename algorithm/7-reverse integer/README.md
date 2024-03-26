# Reverse Integer

## tags
math

## Companies
Given a signed 32-bit integer x, return x with its digits reversed.If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned)

## Example 1:
```go
Input: x = 123
Output: 321
```

## Example 2:
```go
Input: x = -123
Output: -321
```

## Example 3:
```go
Input: x = 120
Output: 21
```

## Constraints:
- -2^31 <= x <= 2^31 -1

## 整数反转- 思路
记 rew 为翻转后的数字，为完成翻转，我们可以重复「弹出」 x 的末尾数字，将其「推入」rew 的末尾， 直至 x 为 0
要在没有辅助栈或数组的帮助下 弹出 和 推入 数字， 我们可以使用如下数学方法：
```go
// 弹出 x 的末尾数字 digit
digit = x % 10
x /= 10

// 将数字 digit 推入 rev 末尾
rev = rev * 10 + digit
```

题目需要判断反转后的数字是否超过 323232 位有符号整数的范围 [−231,231−1][-2^{31},2^{31}-1][−2 
31
 ,2 
31
 −1]，例如 x=2123456789x=2123456789x=2123456789 反转后的 rev=9876543212>231−1=2147483647\textit{rev}=9876543212>2^{31}-1=2147483647rev=9876543212>2 
31
 −1=2147483647，超过了 323232 位有符号整数的范围。

因此我们需要在「推入」数字之前，判断是否满足

−231≤rev⋅10+digit≤231−1-2^{31}\le\textit{rev}\cdot10+\textit{digit}\le2^{31}-1−2 
31
 ≤rev⋅10+digit≤2 
31
 −1

若该不等式不成立则返回 000。

但是题目要求不允许使用 646464 位整数，即运算过程中的数字必须在 323232 位有符号整数的范围内，因此我们不能直接按照上述式子计算，需要另寻他路。

考虑 x>0x>0x>0 的情况，记 INT_MAX=231−1=2147483647\textit{INT\_MAX}=2^{31}-1=2147483647INT_MAX=2 
31
 −1=2147483647，由于

INT_MAX=⌊INT_MAX10⌋⋅10+(INT_MAX mod 10)=⌊INT_MAX10⌋⋅10+7\begin{aligned} \textit{INT\_MAX}&=\lfloor\dfrac{\textit{INT\_MAX}}{10}\rfloor\cdot10+(\textit{INT\_MAX}\bmod10)\\ &=\lfloor\dfrac{\textit{INT\_MAX}}{10}\rfloor\cdot10+7 \end{aligned}
INT_MAX
​	
  
=⌊ 
10
INT_MAX
​	
 ⌋⋅10+(INT_MAXmod10)
=⌊ 
10
INT_MAX
​	
 ⌋⋅10+7
​	
 
则不等式

rev⋅10+digit≤INT_MAX\textit{rev}\cdot10+\textit{digit}\le\textit{INT\_MAX}rev⋅10+digit≤INT_MAX

等价于

rev⋅10+digit≤⌊INT_MAX10⌋⋅10+7\textit{rev}\cdot10+\textit{digit}\le\lfloor\dfrac{\textit{INT\_MAX}}{10}\rfloor\cdot10+7rev⋅10+digit≤⌊ 
10
INT_MAX
​	
 ⌋⋅10+7

移项得

(rev−⌊INT_MAX10⌋)⋅10≤7−digit(\textit{rev}-\lfloor\dfrac{\textit{INT\_MAX}}{10}\rfloor)\cdot10\le7-\textit{digit}(rev−⌊ 
10
INT_MAX
​	
 ⌋)⋅10≤7−digit

讨论该不等式成立的条件：

若 rev>⌊INT_MAX10⌋\textit{rev}>\lfloor\cfrac{\textit{INT\_MAX}}{10}\rfloorrev>⌊ 
10
INT_MAX
​	
 ⌋，由于 digit≥0\textit{digit}\ge0digit≥0，不等式不成立。 若 rev=⌊INT_MAX10⌋\textit{rev}=\lfloor\cfrac{\textit{INT\_MAX}}{10}\rfloorrev=⌊ 
10
INT_MAX
​	
 ⌋，当且仅当 digit≤7\textit{digit}\le7digit≤7 时，不等式成立。 若 rev<⌊INT_MAX10⌋\textit{rev}<\lfloor\cfrac{\textit{INT\_MAX}}{10}\rfloorrev<⌊ 
10
INT_MAX
​	
 ⌋，由于 digit≤9\textit{digit}\le9digit≤9，不等式成立。

注意到当 rev=⌊INT_MAX10⌋\textit{rev}=\lfloor\cfrac{\textit{INT\_MAX}}{10}\rfloorrev=⌊ 
10
INT_MAX
​	
 ⌋ 时若还能推入数字，则说明 xxx 的位数与 INT_MAX\textit{INT\_MAX}INT_MAX 的位数相同，且要推入的数字 digit\textit{digit}digit 为 xxx 的最高位。由于 xxx 不超过 INT_MAX\textit{INT\_MAX}INT_MAX，因此 digit\textit{digit}digit 不会超过 INT_MAX\textit{INT\_MAX}INT_MAX 的最高位，即 digit≤2\textit{digit}\le2digit≤2。所以实际上当 rev=⌊INT_MAX10⌋\textit{rev}=\lfloor\cfrac{\textit{INT\_MAX}}{10}\rfloorrev=⌊ 
10
INT_MAX
​	
 ⌋ 时不等式必定成立。

因此判定条件可简化为：当且仅当 rev≤⌊INT_MAX10⌋\textit{rev}\le\lfloor\cfrac{\textit{INT\_MAX}}{10}\rfloorrev≤⌊ 
10
INT_MAX
​	
 ⌋ 时，不等式成立。

x<0x<0x<0 的情况类似，留给读者自证，此处不再赘述。

综上所述，判断不等式

−231≤rev⋅10+digit≤231−1-2^{31}\le\textit{rev}\cdot10+\textit{digit}\le2^{31}-1−2 
31
 ≤rev⋅10+digit≤2 
31
 −1

是否成立，可改为判断不等式

⌈−23110⌉≤rev≤⌊231−110⌋\lceil\cfrac{-2^{31}}{10}\rceil\le\textit{rev}\le\lfloor\dfrac{2^{31}-1}{10}\rfloor⌈ 
10
−2 
31
 
​	
 ⌉≤rev≤⌊ 
10
2 
31
 −1
​	
 ⌋

是否成立，若不成立则返回 000。

作者：力扣官方题解
链接：https://leetcode.cn/problems/reverse-integer/solutions/755611/zheng-shu-fan-zhuan-by-leetcode-solution-bccn/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


