package leetcode

import (
	"fmt"
	"math"
)

func NumSquares() {
	// fmt.Println(numSquares(12)) // 3 - 4+4+4
	fmt.Println(numSquares(13)) // 2 - 4+9
	// fmt.Println(numSquares(1))  // 1
	// fmt.Println(numSquares(4)) // 1
}

func numSquares(n int) int {
	cache := make([]int, 10000)
	for i := 1; i < 10000; i++ {
		cache[i] = i
	}
	for i := 4; i <= n; i++ {
		for j := 0; j <= i; j++ {
			rt := int(math.Sqrt(float64(j)))
			if rt*rt == j {
				cache[j] = min(cache[j], cache[i-j]+1)
			}
		}
	}
	return cache[n]
}

// top down
// func numSquares(n int) int {
// 	cache := make([]int, 10000)
// 	var dfs func(nn int) int
// 	dfs = func(nn int) int {
// 		if nn <= 0 {
// 			return 0
// 		}
// 		if v := cache[nn]; v == 0 {
// 			count := 10000
// 			for i := nn; i > 0; i-- {
// 				rt := int(math.Sqrt(float64(i)))
// 				if rt*rt == i {
// 					count = min(count, dfs(nn-i)+1)
// 				}
// 			}
// 			cache[nn] = count
// 		}
// 		return cache[nn]
// 	}
// 	return dfs(n)
// }
