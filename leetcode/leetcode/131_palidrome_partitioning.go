package leetcode

import (
	"fmt"

	"github.com/vrishikesh/go-leetcode/queue/queue"
)

func PalindromePartitioning(s string) {
	var ans [][]string
	var q queue.Queue[string]
	var f func(int)

	f = func(index int) {
		if index == len(s) {
			ans = append(ans, q)
			return
		}

		for i := index; i < len(s); i++ {
			if isValidPalindrome(s[index : i+1]) {
				q.Push(s[index : i+1])
				f(i + 1)
				q.Pop()
			}
		}
	}

	f(0)

	fmt.Println(ans)
}

func isValidPalindrome(s string) bool {
	N := len(s)
	for i := 0; i < N/2; i++ {
		if s[i] != s[N-i-1] {
			return false
		}
	}
	return true
}
