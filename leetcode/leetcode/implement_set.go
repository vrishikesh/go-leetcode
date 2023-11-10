package leetcode

import "fmt"

type set[T comparable] struct {
	m    map[T]*DoublyListNode[T]
	head *DoublyListNode[T]
	tail *DoublyListNode[T]
}

func NewSet[T comparable](xx ...T) *set[T] {
	m := map[T]*DoublyListNode[T]{}
	var head *DoublyListNode[T]
	var tail *DoublyListNode[T]
	for _, x := range xx {
		node := &DoublyListNode[T]{Val: x}
		m[x] = node
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			node.Prev = tail
			tail = tail.Next
		}
	}
	return &set[T]{m, head, tail}
}

func (s *set[T]) Has(x T) bool {
	_, ok := s.m[x]
	return ok
}

func (s *set[T]) Add(x T) {
	if s.Has(x) {
		return
	}

	node := &DoublyListNode[T]{Val: x}
	s.m[x] = node
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.Next = node
		node.Prev = s.tail
		s.tail = s.tail.Next
	}
}

func (s *set[T]) Remove(x T) {
	if !s.Has(x) {
		return
	}

	if len(s.m) == 1 {
		s.head = nil
		s.tail = nil
		delete(s.m, x)
		return
	}

	node := s.m[x]
	if node == s.head {
		s.head = s.head.Next
		s.head.Prev = nil
	} else if node == s.tail {
		s.tail = s.tail.Prev
		s.tail.Next = nil
	} else {
		prev := node.Prev
		next := node.Next
		prev.Next = next
		next.Prev = prev
	}

	node = nil
	delete(s.m, x)
}

func (s *set[T]) Chan() chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		curr := s.head
		for curr != nil {
			ch <- curr.Val
			curr = curr.Next
		}
	}()
	return ch
}

func CreateSet() {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(1)
	s.Add(3)
	// s.Remove(1)
	// s.Remove(3)
	// s.Remove(2)
	for x := range s.Chan() {
		fmt.Println(x)
	}
}
