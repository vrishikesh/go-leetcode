package leetcode

import "fmt"

func IsSubsequence(s string, t string) bool {
	T := len(t)
	S := len(s)

	if T < S {
		return false
	}
	if S == 0 {
		return true
	}

	p1 := 0

	for i := 0; i < T; i++ {
		if s[p1] == t[i] {
			p1 += 1
		}
		if p1 == S {
			// return true
			break
		}
	}

	fmt.Println(p1)
	fmt.Println(p1 == len(s))
	return p1 == S
}
