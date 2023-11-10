package leetcode

import (
	"cmp"
	"fmt"
)

type priorityQueue[T cmp.Ordered] struct {
	comparator func(a, b T) bool
	heap       []T
}

func NewPriorityQueue[T cmp.Ordered](
	comparator func(a, b T) bool,
	xx ...T,
) *priorityQueue[T] {
	if comparator == nil {
		comparator = func(a, b T) bool {
			return a > b
		}
	}

	size := max(len(xx)*2, 8)
	heap := make([]T, 0, size)
	pq := &priorityQueue[T]{comparator, heap}
	for i := 0; i < len(xx); i++ {
		pq.Push(xx[i])
	}
	return pq
}

func (pq *priorityQueue[T]) Size() int {
	return len(pq.heap)
}

func (pq *priorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *priorityQueue[T]) Peek() T {
	return pq.heap[0]
}

func (pq *priorityQueue[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (pq *priorityQueue[T]) leftChild(idx int) int {
	return idx*2 + 1
}

func (pq *priorityQueue[T]) rightChild(idx int) int {
	return idx*2 + 2
}

func (pq *priorityQueue[T]) swap(x, y int) {
	pq.heap[x], pq.heap[y] = pq.heap[y], pq.heap[x]
}

func (pq *priorityQueue[T]) compare(x, y int) bool {
	return pq.comparator(pq.heap[x], pq.heap[y])
}

func (pq *priorityQueue[T]) Push(x T) {
	pq.heap = append(pq.heap, x)
	idx := pq.Size() - 1

	for idx > 0 && pq.compare(idx, pq.parent(idx)) {
		pq.swap(idx, pq.parent(idx))
		idx = pq.parent(idx)
	}
}

func (pq *priorityQueue[T]) Pop() T {
	n := pq.Size()
	x := pq.Peek()
	pq.swap(0, n-1)
	pq.heap = pq.heap[:n-1]
	idx := 0

	for (pq.leftChild(idx) < pq.Size() && pq.compare(pq.leftChild(idx), idx)) ||
		(pq.rightChild(idx) < pq.Size() && pq.compare(pq.rightChild(idx), idx)) {
		greaterIdx := pq.leftChild(idx)
		if pq.rightChild(idx) < pq.Size() && pq.compare(pq.rightChild(idx), pq.leftChild(idx)) {
			greaterIdx = pq.rightChild(idx)
		}
		pq.swap(greaterIdx, idx)
		idx = greaterIdx
	}

	return x
}

func ImplementPriorityQueue() {
	pq := NewPriorityQueue[int](
		func(a, b int) bool { return a > b },
		50, 25, 35, 10, 15, 20, 40,
	)
	pq.Push(45)
	pq.Push(75)

	for !pq.IsEmpty() {
		fmt.Println(pq.Pop())
	}
}
