package leetcode

import (
	"container/list"
	"fmt"
)

func PalindromePartitioning() {
	palindromePartitioning("aab")
}

func palindromePartitioning(s string) {
	var ans []*list.List
	var f func(int)
	q := list.New()

	f = func(index int) {
		if index == len(s) {
			ans = append(ans, q)
			return
		}

		for i := index; i < len(s); i++ {
			if isValidPalindrome(s[index : i+1]) {
				q.PushBack(s[index : i+1])
				f(i + 1)
				q.Remove(q.Front())
			}
		}
	}

	f(0)

	for _, l := range ans {
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value)
		}
		fmt.Println()
	}
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
