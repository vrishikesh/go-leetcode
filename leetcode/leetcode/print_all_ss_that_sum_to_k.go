package leetcode

import (
	"fmt"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

func PrintAllSSThatSumToK(s []int, k int) {
	var q queue.Queue[int]
	var f func(i int, sum int)

	f = func(i int, sum int) {
		if i >= len(s) {
			if sum == k {
				fmt.Println(q)
			}
			return
		}

		q.Push(s[i])
		f(i+1, sum+s[i])

		q.Pop()
		f(i+1, sum)
	}

	f(0, 0)
}
