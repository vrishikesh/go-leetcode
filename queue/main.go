package main

import (
	"fmt"

	"queue/queue"
)

func main() {
	var q queue.Queue[int]
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q)
	fmt.Println(q.Len())
	fmt.Println(q.IsEmpty())
}
