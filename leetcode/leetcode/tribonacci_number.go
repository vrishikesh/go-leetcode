package leetcode

import "fmt"

func Tribonacci() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(3))
	fmt.Println(tribonacci(2))
	fmt.Println(tribonacci(1))
	fmt.Println(tribonacci(0))
}

// bottom up
func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	DP := make([]int, n+1)
	DP[1] = 1
	DP[2] = 1
	for i := 3; i <= n; i++ {
		DP[i] = DP[i-3] + DP[i-2] + DP[i-1]
	}
	return DP[n]
}

// top down
// func tribonacci(n int) int {
// 	DP := make([]int, max(n+1, 3))
// 	DP[1] = 1
// 	DP[2] = 1
// 	var recurse func(n int) int
// 	recurse = func(n int) int {
// 		if n == 0 || n == 1 {
// 			return n
// 		}
// 		if n == 2 {
// 			return 1
// 		}
// 		if v := DP[n]; v == 0 {
// 			DP[n] = recurse(n-3) + recurse(n-2) + recurse(n-1)
// 		}
// 		return DP[n]
// 	}
// 	return recurse(n)
// }
