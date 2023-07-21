package leetcode

import (
	"container/list"
	"fmt"
)

func PrintAllSSThatSumToK(s []int, k int) {
	var f func(i int, sum int)
	q := list.New()

	f = func(i int, sum int) {
		if i >= len(s) {
			if sum == k {
				for e := q.Front(); e != nil; e = e.Next() {
					fmt.Printf("%v ", e.Value)
				}
				fmt.Println()
			}
			return
		}

		q.PushBack(s[i])
		f(i+1, sum+s[i])

		q.Remove(q.Front())
		f(i+1, sum)
	}

	f(0, 0)
}
