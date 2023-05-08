package leetcode

import "fmt"

func PrintTheCountOfSsThatSumToK(s []int, k int) {
	var f func(int, int) int

	f = func(i, sum int) int {
		if i >= len(s) {
			if sum == k {
				return 1
			}
			return 0
		}

		l := f(i+1, sum+s[i])
		r := f(i+1, sum)
		return l + r
	}

	fmt.Println(f(0, 0))
}
