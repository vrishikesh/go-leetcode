package leetcode

import (
	"container/list"
	"fmt"
)

func PrintAllSubsequences() {
	printAllSubsequences([]int{1, 2, 1})
}

func printAllSubsequences(s []int) {
	var f func(int)
	q := list.New()

	f = func(i int) {
		if i >= len(s) {
			for e := q.Front(); e != nil; e = e.Next() {
				fmt.Printf("%v ", e.Value)
			}
			fmt.Println()
			return
		}

		q.PushBack(s[i])
		f(i + 1)

		q.Remove(q.Front())
		f(i + 1)
	}

	f(0)
}
