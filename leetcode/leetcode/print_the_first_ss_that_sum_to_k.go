package leetcode

import (
	"container/list"
	"fmt"
)

func PrintTheFirstSSThatSumToK(s []int, k int) {
	var f func(int, int) bool
	q := list.New()

	f = func(i, sum int) bool {
		if i >= len(s) {
			if sum == k {
				for e := q.Front(); e != nil; e = e.Next() {
					fmt.Printf("%v ", e.Value)
				}
				fmt.Println()
				return true
			}
			return false
		}

		q.PushBack(s[i])
		if f(i+1, sum+s[i]) {
			return true
		}

		q.Remove(q.Front())
		return f(i+1, sum)
	}

	f(0, 0)
}
