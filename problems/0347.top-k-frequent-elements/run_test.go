package elements

import (
	"container/heap"
	"testing"
)

func TestRun(t *testing.T) {
	Run()
}

func TestPriorityQueue(t *testing.T) {
	p := PriorityQueue{}
	p.Push(&Item{key: 1, count: 1})
	p.Push(&Item{key: 1, count: 3})
	p.Push(&Item{key: 2, count: 2})
	p.Push(&Item{key: 2, count: 1})
	a := p.Pop()
	t.Log(a)
	heap.Push(&p, &Item{key: 3, count: 1})
	heap.Push(&p, &Item{key: 4, count: 5})
	heap.Push(&p, &Item{key: 5, count: 2})
	i := heap.Pop(&p)
	t.Log(i)
}
