# 题目
Given a string, find the length of the longest substring without repeating characters.

Examples:

```
Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

# 分析

```
最长的不重复子字符串，只要有一个重复的字符，便要重新计算。

用 map[rune]int 记录遍历过的字符与数目，
当读取到重复的字符，记录不改变 max

```