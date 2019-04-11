# Two Sum

```English
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

```Chinese
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

---

20190411 update

今天看剑指Offer，第41题与此题类似。

*题干*

```Chinese
输入一个递增排序的数组与一个数字s，在数组中查找两个数，使得他们的和，
正好是s，如果有多对和为s则输出任意一对。
```

`步骤`

1. 数组排序 (要求输出下标，排序会破会原有结构)
2. 两个指针，指向头尾位置
3. 相加小于s，则头向后移动
4. 相加等于s，则返回
5. 相加大于s，则尾部向前移
6. 直到头>尾

__注意这两个题的区别__