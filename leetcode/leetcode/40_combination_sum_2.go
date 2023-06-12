package leetcode

import (
	"fmt"
	"sort"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

// Print all unique combinations of subsequences that sum to k and in lexicographically sorted order
func CombinationSum2(s []int, k int) {
	var q queue.Queue[int]
	var f func(i int, sum int)

	sort.Ints(s)

	f = func(index int, target int) {
		if target == 0 {
			fmt.Println(q)
			return
		}

		for i := index; i < len(s); i++ {
			if i > index && s[i] == s[i-1] {
				continue
			}
			if s[i] > target {
				break
			}
			q.Push(s[i])
			f(i+1, target-s[i])
			q.Pop()
		}
	}

	f(0, k)
}
