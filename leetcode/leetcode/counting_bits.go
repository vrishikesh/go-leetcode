package leetcode

import (
	"fmt"
	"math"
)

func CountBits() {
	fmt.Println(countBits(5))
	fmt.Println(countBits(2))
	fmt.Println(countBits(1))
	fmt.Println(countBits(0))
}

// bottom up
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	DP := make([]int, n+1)
	DP[1] = 1
	var bitCounter float64 = 1
	for i := 2; i <= n; {
		base := int(math.Pow(2, bitCounter))
		for j := 0; j < base && i <= n; {
			DP[i] = 1 + DP[j]
			i++
			j++
		}
		bitCounter += 1
	}
	return DP
}

// naive approach
// func countBits(n int) []int {
// 	ans := make([]int, n+1)
// 	for i := 0; i <= n; i++ {
// 		count := 0
// 		for j := i; j > 0; j = j >> 1 {
// 			if j&1 == 1 {
// 				count += 1
// 			}
// 		}
// 		ans[i] = count
// 	}
// 	return ans
// }
