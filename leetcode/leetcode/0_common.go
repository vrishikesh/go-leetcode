package leetcode

type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

type DoublyListNode[T comparable] struct {
	Val  T
	Next *DoublyListNode[T]
	Prev *DoublyListNode[T]
}

type AnaryNode[T comparable] struct {
	Val      T
	Children []*AnaryNode[T]
}

type Queue[T comparable] []T

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Enqueue(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Dequeue() T {
	cp := *q
	x := cp[0]
	*q = cp[1:]
	return x
}

func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

type Stack[T comparable] []T

func (q *Stack[T]) Len() int {
	return len(*q)
}

func (q *Stack[T]) Push(x T) {
	*q = append(*q, x)
}

func (q *Stack[T]) Pop() T {
	cp := *q
	n := q.Len()
	x := cp[n-1]
	*q = cp[:n-1]
	return x
}

func (q *Stack[T]) Peek() T {
	n := q.Len()
	return (*q)[n-1]
}
