package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
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
