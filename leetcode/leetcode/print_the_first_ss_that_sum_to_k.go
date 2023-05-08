package leetcode

import (
	"fmt"
	"queue/queue"
)

func PrintTheFirstSSThatSumToK(s []int, k int) {
	var q queue.Queue[int]
	var f func(int, int) bool

	f = func(i, sum int) bool {
		if i >= len(s) {
			if sum == k {
				fmt.Println(q)
				return true
			}
			return false
		}

		q.Push(s[i])
		if f(i+1, sum+s[i]) {
			return true
		}

		q.Pop()
		return f(i+1, sum)
	}

	f(0, 0)
}
