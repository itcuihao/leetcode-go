package elements

import (
	"container/heap"
	"fmt"
)

func Run() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n]++
	}
	pq := PriorityQueue{}
	for k, v := range m {
		heap.Push(&pq, &Item{key: k, count: v})
	}
	result := make([]int, 0, k)
	for len(result) < k {
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.key)
	}
	return result
}

type Item struct {
	key   int
	count int
}

// 优先队列
type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].count > p[j].count
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	n := len(*p)
	item := (*p)[n-1]
	*p = (*p)[:n-1]
	return item
}
