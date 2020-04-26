
# 0155.min-stack

```
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
 

示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
 

提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

```text
type MinStack struct {
	Data     []int
	MinIndex int
}
该实现的执行结果

执行用时 :
20 ms
, 在所有 Go 提交中击败了
78.89%
的用户
内存消耗 :
7.5 MB
, 在所有 Go 提交中击败了
75.00%
的用户
```

执行8ms的例子，利用一个数组存取对应的最小值，空间换时间，消除循环。
```text
type MinStack struct {
    data []int
    minData []int
    size int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        data: make([]int, 0),
        minData: make([]int, 0),
        size: 0,
    }
}


func (this *MinStack) Push(x int)  {
    if this.size == 0 {
        this.minData = append(this.minData, x)
    } else if this.minData[this.size-1] > x {
        this.minData = append(this.minData, x)
    } else {
        this.minData = append(this.minData, this.minData[this.size-1])
    }
    this.data = append(this.data, x)
    this.size++
}


func (this *MinStack) Pop()  {
    this.data = this.data[:this.size-1]
    this.minData = this.minData[:this.size-1]
    this.size--
}


func (this *MinStack) Top() int {
    return this.data[this.size-1]
}


func (this *MinStack) GetMin() int {
    return this.minData[this.size-1]
}

```