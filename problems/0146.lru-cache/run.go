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
	cir   []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lcap:  capacity,
		param: make(map[int]int, capacity),
		cir:   make([]int, 0, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.param[key]
	if !ok {
		return -1
	}
	this.ext(key)
	fmt.Printf("c:%v, p:%+v, r:%+v\n", this.lcap, this.param, this.cir)
	return v
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.param[key]
	if ok {
		this.param[key] = value
		this.ext(key)
		return
	}

	if this.lcap == len(this.cir) {
		old := this.cir[0]
		delete(this.param, old)
		this.cir = this.cir[1:]
	}
	this.param[key] = value
	this.cir = append(this.cir, key)
	this.ext(key)

	fmt.Printf("c:%v, p:%+v, r:%+v\n", this.lcap, this.param, this.cir)

}

func (this *LRUCache) ext(key int) {
	index := 0
	end := len(this.cir) - 1
	for i, v := range this.cir {
		if key == v {
			index = i
		}
	}
	this.cir[0], this.cir[index] = this.cir[index], this.cir[0]
	this.cir[end], this.cir[index] = this.cir[index], this.cir[end]
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
