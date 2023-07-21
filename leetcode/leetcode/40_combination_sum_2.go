package leetcode

import (
	"container/list"
	"fmt"
	"sort"
)

// Print all unique combinations of subsequences that sum to k and in lexicographically sorted order
func CombinationSum2(s []int, k int) {
	var f func(i int, sum int)
	q := list.New()

	sort.Ints(s)

	f = func(index int, target int) {
		if target == 0 {
			for e := q.Front(); e != nil; e = e.Next() {
				fmt.Printf("%v ", e.Value)
			}
			fmt.Println()
			return
		}

		for i := index; i < len(s); i++ {
			if i > index && s[i] == s[i-1] {
				continue
			}
			if s[i] > target {
				break
			}
			q.PushBack(s[i])
			f(i+1, target-s[i])
			q.Remove(q.Front())
		}
	}

	f(0, k)
}
