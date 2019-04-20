
# [0097.interleaving-string](https://leetcode.com/problems/interleaving-string/)

```English
Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

Example 1:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
```

```Chinese
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false
```

利用DP

```
考虑用动态规划做。

s1, s2只有两个字符串，因此可以展平为一个二维地图，判断是否能从左上角走到右下角。

当s1到达第i个元素，s2到达第j个元素:

地图上往右一步就是s2[j-1]匹配s3[i+j-1]。

地图上往下一步就是s1[i-1]匹配s3[i+j-1]。

示例：s1="aa",s2="ab",s3="aaba"。标1的为可行。最终返回右下角。

     0  a  b

0   1  1  0

a   1  1  1

a   1  0  1
```