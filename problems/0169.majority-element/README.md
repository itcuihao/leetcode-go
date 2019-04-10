
# 0169.majority-element

```English
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
```

```Chinese
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
```

**分析**

1. 利用map，循环时重复的key，将value++ ，再次循环map，将value>len(n)/2的key返回。
2. 既然众数是超过数组一半的数，那就将数组排序，返回中间值，该值就是众数。

**总结**

遇到数组的题目可以考虑排序后有序数组可不可以有效的解决问题

例如：136题也是类似