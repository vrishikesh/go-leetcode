package leetcode

import (
	"fmt"
	"queue/queue"
)

// print all permutations of a string/array using map
// O(n! * n) S(n)
func PrintAllPermutationsMap(s []int) {
	m := make(map[int]bool, len(s))
	var q queue.Queue[int]
	var f func(int)

	f = func(index int) {
		if index == len(s) {
			fmt.Println(q)
			return
		}
		for i := 0; i < len(s); i++ {
			if !m[i] {
				q.Push(s[i])
				m[i] = true
				f(index + 1)
				q.Pop()
				m[i] = false
			}
		}
	}

	f(0)
}

// print all permutations of a string/array using swap
// O(n! * n) S(1)
func PrintAllPermutationsSwap(a [3]int) {
	var f func(int, [3]int)

	f = func(index int, a [3]int) {
		if index == len(a) {
			fmt.Println(a)
			return
		}
		for i := index; i < len(a); i++ {
			a[index], a[i] = a[i], a[index]
			f(index+1, a)
			a[index], a[i] = a[i], a[index]
		}
	}

	f(0, a)
}
