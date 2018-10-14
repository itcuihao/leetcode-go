# 29. Divide Two Integers

## 题干

- English

```
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
```

- 中文

```
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
```

## 题目分析

```
1. 被除数/除数=商 （忽略余数）=> 被除数=除数*商。

2. 商（任意整数）可以表示为：a_0*2^0+a_1*2^1+...a_i*2^i+...+a_n*2^n.

3. 在Java中左移操作<<相当于对一个数乘以2，右移操作相当于除以2.

4.我们让除数左移直到大于被除数前的的一个数，例如计算28/3，我们进行三次左移操作，使3*2*2*2=3*8=24<28(注意四次左移操作得到3*2^4=48>28).记录下2*2*2=2^3=8.

5.我们让28减去24得到4，然后像第四步一样计算4/3,则3*2^0=3<4.记录下2^0=1.

6.由于4-3=1小于除数3，停止计算，并将每轮得到的值相加，在本例中8+1=9，记得到商（即28/3=9）。

至此，程序的主题思想已经介绍完了，接下来要注意数据左移和求整数绝对值的边界问题。

7. 将输入的Int(32位)型数字转换为long(64位)型。

8. 考虑MIN_VALUE/-1的特殊情况。
```
