package queue

import (
	"fmt"
)

type Queue[T comparable] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	N := len(*q)
	v := (*q)[N-1]
	*q = (*q)[:N-1]
	return v
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", *q)
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
