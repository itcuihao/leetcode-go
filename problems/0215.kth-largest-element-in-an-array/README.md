
# 0215.kth-largest-element-in-an-array

```English
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note: 
You may assume k is always valid, 1 ≤ k ≤ array's length.
```

```Chinese
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
```

--- 

不同的 partition

```
func partion(a []int, l int, r int)int{
    i, j := l, r
    v := a[r]

    for {
            for i < r && a[i] <= v {
                    i++
            }
            for j >= l && a[j] >= v {
                    j--
            }

            if i > j {
                    break
            }
            a[i], a[j] = a[j], a[i]
            i++
            j++
    }

    a[i], a[r] = a[r], a[i]
    return i

}
```