package stack

import "testing"

func TestRun(t *testing.T) {
	minStack := Constructor()
	//minStack.Push(-2)
	//minStack.Push(0)
	//minStack.Push(-1)
	a := minStack.GetMin()
	t.Log(a)

	a = minStack.Top()
	t.Log(a)
	minStack.Pop()

	a = minStack.GetMin()
	t.Log(a)
}
