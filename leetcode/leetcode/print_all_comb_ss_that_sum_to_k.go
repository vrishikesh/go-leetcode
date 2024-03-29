package leetcode

import (
	"container/list"
	"fmt"
)

func PrintAllCombSubsequencesThatSumToK() {
	printAllCombSubsequencesThatSumToK([]int{2, 3, 6, 7}, 7)
}

func printAllCombSubsequencesThatSumToK(s []int, k int) {
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
		if sum+s[i] <= k {
			f(i, sum+s[i])
		}

		q.Remove(q.Front())
		f(i+1, sum)
	}

	f(0, 0)
}
