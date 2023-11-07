package leetcode

import "fmt"

type MyQueue struct {
	st1 Stack[int]
	st2 Stack[int]
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.st1.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.st2.Len() == 0 {
		for q.st1.Len() > 0 {
			q.st2.Push(q.st1.Pop())
		}
	}

	return q.st2.Pop()
}

func (q *MyQueue) Peek() int {
	if q.st2.Len() == 0 {
		for q.st1.Len() > 0 {
			q.st2.Push(q.st1.Pop())
		}
	}

	return q.st2.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.st1.Len() == 0 && q.st2.Len() == 0
}

func ImplementQueueUsingStacks() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
