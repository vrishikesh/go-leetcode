package leetcode

import "fmt"

func CountSubstrings() {
	fmt.Println(countSubstrings("aaaaa"))
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("aa"))
	fmt.Println(countSubstrings("a"))
}

// top down
func countSubstrings(s string) int {
	cache := make(map[[2]any]int, len(s))
	var isPalindrome = func(s string) int {
		n := len(s)
		for i := 0; i < n/2; i++ {
			if s[i] != s[n-1-i] {
				return 0
			}
		}
		return 1
	}
	var recurse func(str string, idx int) int
	recurse = func(str string, idx int) int {
		n := len(str)
		if n == 0 {
			return 0
		}
		cacheIdx := [2]any{str, idx}
		if _, ok := cache[cacheIdx]; ok {
			return 0
		}
		cache[cacheIdx] = isPalindrome(str) + recurse(str[:n-1], idx) + recurse(str[1:], idx+1)
		return cache[cacheIdx]
	}
	return recurse(s, 0)
}
