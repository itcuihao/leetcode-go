# 题干

Given a linked list, remove the n-th node from the end of list and return its head.

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

## Example:

```
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
```

## Note:

Given n will always be valid.

给定的 n 保证是有效的。

## Follow up:

Could you do this in one pass?

你能尝试使用一趟扫描实现吗？

## 思路

- 循环一次找到节点返回，p.next=p.next.next