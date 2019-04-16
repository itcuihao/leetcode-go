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

// //  别人家的：
// 这种写法看的比价凌乱
type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type LRUCache struct {
	capacity int
	head     *node
	mMap     map[int]*node
}

func Constructor(capacity int) LRUCache {
	// 初始化
	cache := LRUCache{capacity: capacity}
	cache.mMap = make(map[int]*node)
	cache.head = &node{key: -1, val: -1}

	// 初始化链表
	cur := cache.head
	for i := 0; i < capacity-1; i++ {
		node := &node{key: -1, val: -1}
		cur.next = node
		node.prev = cur
		cur = node
	}

	cur.next = cache.head
	cache.head.prev = cur

	return cache
}

func (cache *LRUCache) move2Head(cur *node) {

	// 当前节点 == head节点
	if cur == cache.head {
		// head节点 = head的前一个节点
		cache.head = cache.head.prev
		return
	}

	cur.prev.next = cur.next
	cur.next.prev = cur.prev

	cur.next = cache.head.next
	cur.next.prev = cur
	cache.head.next = cur
	cur.prev = cache.head
}

func (cache *LRUCache) Get(key int) int {
	// 取值
	node, ok := cache.mMap[key]
	// 没有值，return -1
	if !ok {
		return -1
	}

	// 移动到head
	cache.move2Head(node)
	// 返回值
	return node.val
}

func (cache *LRUCache) Put(key int, value int) {
	// 判断key是否存在
	node, ok := cache.mMap[key]
	if ok {
		// 存在更新value
		node.val = value
		// 更新head节点
		cache.move2Head(node)
	} else {
		// 不存在
		// 获取head old key
		oldKey := cache.head.key
		// 更新head节点key,value
		cache.head.key = key
		cache.head.val = value
		// 删除 old key
		delete(cache.mMap, oldKey)
		// 更新map
		cache.mMap[key] = cache.head
		// 更新head=head.prev
		cache.head = cache.head.prev
	}
}

// type LRUCache struct {
// 	lcap  int
// 	param map[int]int
// 	cir   []int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		lcap:  capacity,
// 		param: make(map[int]int, capacity),
// 		cir:   make([]int, 0, capacity),
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	v, ok := this.param[key]
// 	if !ok {
// 		return -1
// 	}
// 	this.ext(key)

// 	return v
// }

// func (this *LRUCache) Put(key int, value int) {
// 	_, ok := this.param[key]
// 	if ok {
// 		this.param[key] = value
// 		this.ext(key)
// 		return
// 	}

// 	if this.lcap == len(this.cir) {
// 		old := this.cir[len(this.cir)-1]
// 		delete(this.param, old)
// 		this.cir = this.cir[:len(this.cir)-1]
// 	}
// 	this.param[key] = value
// 	this.cir = append(this.cir, key)
// 	this.ext(key)

// }

// func (this *LRUCache) ext(key int) {

// 	list := make([]int, 0, len(this.cir))
// 	index := 0

// 	for i, v := range this.cir {
// 		if key == v {
// 			index = i
// 		}
// 	}

// 	list = append(list, this.cir[index])

// 	this.cutc(index)

// 	list = append(list, this.cir...)

// 	this.cir = list

// }

// func (this *LRUCache) cutc(index int) {
// 	this.cir = append(this.cir[:index], this.cir[index+1:]...)
// }
