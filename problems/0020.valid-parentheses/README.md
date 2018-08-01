# 题干

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

An input string is valid if:
有效字符串需满足：

```
Open brackets must be closed by the same type of brackets.
左括号必须用相同类型的右括号闭合。

Open brackets must be closed in the correct order.
左括号必须以正确的顺序闭合。
```

Note that an empty string is also considered valid.
注意空字符串可被认为是有效字符串。

Example 1:

```
Input: "()"
Output: true
```

Example 2:

```
Input: "()[]{}"
Output: true
```

Example 3:

```
Input: "(]"
Output: false
```

Example 4:

```
Input: "([)]"
Output: false
```

Example 5:

```
Input: "{[]}"
Output: true
```