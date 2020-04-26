package stack

type MinStack struct {
	Data     []int
	MinIndex int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	length := len(this.Data)
	if length > 0 && length >= this.MinIndex && this.Data[this.MinIndex] > x {
		this.MinIndex = length
	}
	this.Data = append(this.Data, x)
}

func (this *MinStack) Pop() {
	length := len(this.Data)
	if length == 0 {
		return
	}
	this.Data = this.Data[:length-1]
	if length-1 == this.MinIndex {
		this.MinIndex = 0
		for i := 0; i < len(this.Data); i++ {
			if this.Data[this.MinIndex] > this.Data[i] {
				this.MinIndex = i
			}
		}
	}

}

func (this *MinStack) Top() int {
	length := len(this.Data)
	if length == 0 || length < this.MinIndex {
		return 0
	}
	return this.Data[length-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Data) == 0 || len(this.Data) < this.MinIndex {
		return 0
	}
	return this.Data[this.MinIndex]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
