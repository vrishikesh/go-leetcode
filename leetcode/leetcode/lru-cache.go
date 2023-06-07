package leetcode

import "fmt"

type LRUCache struct {
	Capacity int
	Recent   []int
	Values   map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Recent:   make([]int, capacity),
		Values:   make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	for i := 0; i < this.Capacity; i++ {
		if this.Recent[i] == key {
			return this.Values[key]
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.Recent = append([]int{key}, this.Recent...)
	if len(this.Recent) == this.Capacity {
		this.Recent = this.Recent[:this.Capacity]
	}
	this.Values[key] = value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj .Get(key);
 * obj .Put(key,value);
 */
func main() {
	obj := Constructor(2)
	obj.Put(2, 2)                                    // nil
	obj.Put(1, 1)                                    // nil
	fmt.Printf("should return 1: %d\n", obj.Get(1))  // returns 1
	obj.Put(3, 3)                                    // evicts key 2
	fmt.Printf("should return -1: %d\n", obj.Get(2)) // returns -1 (not found)
	obj.Put(4, 4)                                    // evicts key 1
	fmt.Printf("should return -1: %d\n", obj.Get(1)) // -1
	fmt.Printf("should return 3: %d\n", obj.Get(3))  // 3
	fmt.Printf("should return 4: %d\n", obj.Get(4))  // 4
}
