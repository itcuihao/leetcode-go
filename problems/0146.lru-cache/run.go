package cache

import "fmt"

func Run() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	i := cache.Get(1) // 返回  1
	fmt.Println(i)
	cache.Put(3, 3)  // 该操作会使得密钥 2 作废
	i = cache.Get(2) // 返回 -1 (未找到)
	fmt.Println(i)
	cache.Put(4, 4)  // 该操作会使得密钥 1 作废
	i = cache.Get(1) // 返回 -1 (未找到)
	fmt.Println(i)
	i = cache.Get(3) // 返回  3
	fmt.Println(i)
	i = cache.Get(4) // 返回  4
	fmt.Println(i)
}

type LRUCache struct {
	lcap  int
	param map[int]int
	cir   map[int]int
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		lcap:  capacity,
		param: make(map[int]int, capacity),
		cir:   make(map[int]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.param[key]
	if !ok {
		return -1
	}
	this.cir[key]++
	return v
}

func (this *LRUCache) Put(key int, value int) {
	if this.lcap == len(this.param) {
		old := 0
		for k, v := range this.cir {
			if v == 0 {
				old = k
			}
		}
		delete(this.param, old)
		delete(this.cir, old)
	}
	this.param[key] = value
	this.cir[key] = 0
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
