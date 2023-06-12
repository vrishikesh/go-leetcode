package leetcode

import (
	"fmt"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

func PrintAllSubsequences(s []int) {
	var q queue.Queue[int]
	var f func(int)

	f = func(i int) {
		if i >= len(s) {
			fmt.Println(q)
			return
		}

		q.Push(s[i])
		f(i + 1)

		q.Pop()
		f(i + 1)
	}

	f(0)
}
